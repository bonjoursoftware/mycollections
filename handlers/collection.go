package handlers

import (
	"errors"
	"net/http"

	"mycollections/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowCollections(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	uid := userID.(int)

	collections, err := models.GetCollections(uid)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "collections.html", gin.H{"Error": "Failed to load collections"})
		return
	}

	c.HTML(http.StatusOK, "collections.html", gin.H{
		"Collections": collections,
	})
}

func HandleCreateCollection(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	uid := userID.(int)

	name := c.PostForm("name")

	// Get existing collections to re-render page if there's an error
	collections, err := models.GetCollections(uid)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "collections.html", gin.H{"Error": "Failed to load collections"})
		return
	}

	err = models.CreateCollection(uid, name)
	if err != nil {
		errMsg := "Failed to create collection"
		if errors.Is(err, models.ErrCollectionExists) {
			errMsg = "A collection with that name already exists"
		} else if errors.Is(err, models.ErrEmptyCollectionName) {
			errMsg = "Collection name cannot be empty"
		} else if errors.Is(err, models.ErrInvalidCollectionName) {
			errMsg = "Collection name must be alphanumeric and spaces"
		}

		c.HTML(http.StatusBadRequest, "collections.html", gin.H{
			"Collections": collections,
			"Error":       errMsg,
			"Name":        name,
		})
		return
	}

	c.Redirect(http.StatusFound, "/collections")
}
