package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"mycollections/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// helper to get userID from session and collectionID from params, and verify ownership
func getAndVerifyCollection(c *gin.Context) (int, int, *models.Collection, error) {
	session := sessions.Default(c)
	userIDStr := session.Get("user_id")
	if userIDStr == nil {
		return 0, 0, nil, errors.New("unauthorized")
	}
	userID := userIDStr.(int)

	colIDStr := c.Param("id")
	collectionID, err := strconv.Atoi(colIDStr)
	if err != nil {
		return userID, 0, nil, errors.New("invalid collection id")
	}

	collection, err := models.GetCollectionByID(userID, collectionID)
	if err != nil {
		return userID, collectionID, nil, err
	}

	return userID, collectionID, collection, nil
}

func ShowCollectionTags(c *gin.Context) {
	_, collectionID, collection, err := getAndVerifyCollection(c)
	if err != nil {
		if errors.Is(err, models.ErrCollectionNotFound) {
			c.String(http.StatusNotFound, "Collection not found")
		} else {
			c.Redirect(http.StatusFound, "/login")
		}
		return
	}

	tags, err := models.GetTags(collectionID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "tags.html", gin.H{"Error": "Failed to load tags", "Collection": collection})
		return
	}

	c.HTML(http.StatusOK, "tags.html", gin.H{
		"Collection": collection,
		"Tags":       tags,
	})
}

func HandleCreateTag(c *gin.Context) {
	_, collectionID, collection, err := getAndVerifyCollection(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	name := c.PostForm("name")

	err = models.CreateTag(collectionID, name)
	if err != nil {
		tags, _ := models.GetTags(collectionID)
		errMsg := "Failed to create tag"
		if errors.Is(err, models.ErrTagExists) {
			errMsg = "A tag with that name already exists in this collection"
		} else if errors.Is(err, models.ErrEmptyTagName) {
			errMsg = "Tag name cannot be empty"
		} else if errors.Is(err, models.ErrInvalidTagName) {
			errMsg = "Tag name must be alphanumeric, hyphens, or underscores with no spaces"
		}

		c.HTML(http.StatusBadRequest, "tags.html", gin.H{
			"Collection": collection,
			"Tags":       tags,
			"Error":      errMsg,
			"Name":       name,
		})
		return
	}

	c.Redirect(http.StatusFound, "/collections/"+strconv.Itoa(collectionID))
}

func HandleEditTag(c *gin.Context) {
	_, collectionID, collection, err := getAndVerifyCollection(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	tagIDStr := c.Param("tag_id")
	tagID, err := strconv.Atoi(tagIDStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid tag ID")
		return
	}

	newName := c.PostForm("name")

	err = models.UpdateTag(tagID, newName)
	if err != nil {
		tags, _ := models.GetTags(collectionID)
		errMsg := "Failed to update tag"
		if errors.Is(err, models.ErrTagExists) {
			errMsg = "A tag with that name already exists in this collection"
		} else if errors.Is(err, models.ErrEmptyTagName) {
			errMsg = "Tag name cannot be empty"
		} else if errors.Is(err, models.ErrInvalidTagName) {
			errMsg = "Tag name must be alphanumeric, hyphens, or underscores with no spaces"
		} else if errors.Is(err, models.ErrTagNotFound) {
			errMsg = "Tag not found"
		}

		c.HTML(http.StatusBadRequest, "tags.html", gin.H{
			"Collection": collection,
			"Tags":       tags,
			"Error":      errMsg,
		})
		return
	}

	c.Redirect(http.StatusFound, "/collections/"+strconv.Itoa(collectionID))
}

func HandleDeleteTag(c *gin.Context) {
	_, collectionID, collection, err := getAndVerifyCollection(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	tagIDStr := c.Param("tag_id")
	tagID, err := strconv.Atoi(tagIDStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid tag ID")
		return
	}

	err = models.DeleteTag(tagID)
	if err != nil {
		tags, _ := models.GetTags(collectionID)
		c.HTML(http.StatusInternalServerError, "tags.html", gin.H{
			"Collection": collection,
			"Tags":       tags,
			"Error":      "Failed to delete tag",
		})
		return
	}

	c.Redirect(http.StatusFound, "/collections/"+strconv.Itoa(collectionID))
}
