package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mycollections/db"
	"mycollections/handlers"
	"mycollections/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	dbPath := "app.db"
	if os.Getenv("DB_PATH") != "" {
		dbPath = os.Getenv("DB_PATH")
	}

	if err := db.InitDB(dbPath); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	r := gin.Default()

	// Initialize Session Store
	sessionSecret := os.Getenv("SESSION_SECRET")
	if sessionSecret == "" {
		log.Fatal("FATAL: SESSION_SECRET environment variable is required")
	}
	store := cookie.NewStore([]byte(sessionSecret))
	r.Use(sessions.Sessions("mysession", store))

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Public Routes
	r.GET("/register", handlers.ShowRegister)
	r.POST("/register", handlers.HandleRegister)
	r.GET("/login", handlers.ShowLogin)
	r.POST("/login", handlers.HandleLogin)
	r.POST("/logout", handlers.HandleLogout)

	// Protected Routes
	authorized := r.Group("/")
	authorized.Use(middleware.RequireAuth())
	{
		authorized.GET("/", handlers.ShowHome)
		authorized.GET("/collections", handlers.ShowCollections)
		authorized.POST("/collections", handlers.HandleCreateCollection)
		
		authorized.GET("/collections/:id", handlers.ShowCollectionTags)
		authorized.POST("/collections/:id/tags", handlers.HandleCreateTag)
		authorized.POST("/collections/:id/tags/:tag_id/edit", handlers.HandleEditTag)
		authorized.POST("/collections/:id/tags/:tag_id/delete", handlers.HandleDeleteTag)
		authorized.GET("/collections/:id/tags/:tag_id", handlers.ShowTagItems)
		authorized.POST("/collections/:id/tags/:tag_id/items", handlers.HandleCreateItem)
		authorized.POST("/collections/:id/tags/:tag_id/items/:item_id/edit", handlers.HandleEditItem)
		authorized.POST("/collections/:id/tags/:tag_id/items/:item_id/delete", handlers.HandleDeleteItem)
	}

	// Start server
	port := ":8080"
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		log.Printf("Starting server on %s...\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT (Ctrl+C)
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	// Close the database connection
	if err := db.CloseDB(); err != nil {
		log.Fatalf("Database forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}