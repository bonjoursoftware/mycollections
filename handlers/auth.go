package handlers

import (
	"errors"
	"net/http"

	"mycollections/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func HandleRegister(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	err := models.CreateUser(email, name, password)
	if err != nil {
		msg := "Failed to register"
		if errors.Is(err, models.ErrEmailExists) {
			msg = "Email already exists"
		}
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"Error": msg})
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{"Success": "Registration successful. Please login."})
}

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func HandleLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := models.AuthenticateUser(email, password)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error": "Invalid email or password"})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Set("user_email", user.Email)
	session.Set("user_name", user.Name)
	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"Error": "Failed to save session"})
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func HandleLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}