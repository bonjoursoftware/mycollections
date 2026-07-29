package handlers

import (
	"net/http"
	"strconv"

	"mycollections/models"

	"github.com/gin-gonic/gin"
)

func ShowTagItems(c *gin.Context) {
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

	// Verify tag belongs to collection
	tags, err := models.GetTags(collectionID)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to load tags")
		return
	}
	
	var currentTag *models.Tag
	for _, t := range tags {
		if t.ID == tagID {
			currentTag = &t
			break
		}
	}
	
	if currentTag == nil {
		c.String(http.StatusNotFound, "Tag not found")
		return
	}

	items, err := models.GetItemsByTag(tagID)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to load items")
		return
	}

	c.HTML(http.StatusOK, "items.html", gin.H{
		"Collection": collection,
		"Tag":        currentTag,
		"AllTags":    tags,
		"Items":      items,
	})
}

func parseTagIDs(c *gin.Context) []int {
	tagIDStrings := c.PostFormArray("tags")
	var tagIDs []int
	for _, idStr := range tagIDStrings {
		if id, err := strconv.Atoi(idStr); err == nil {
			tagIDs = append(tagIDs, id)
		}
	}
	return tagIDs
}

func HandleCreateItem(c *gin.Context) {
	_, collectionID, collection, err := getAndVerifyCollection(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	tagIDStr := c.Param("tag_id")

	name := c.PostForm("name")
	description := c.PostForm("description")
	tagIDs := parseTagIDs(c)

	err = models.CreateItem(collectionID, name, description, tagIDs)
	if err != nil {
		// Provide error back
		tags, _ := models.GetTags(collectionID)
		var currentTag *models.Tag
		tagID, _ := strconv.Atoi(tagIDStr)
		for _, t := range tags {
			if t.ID == tagID {
				currentTag = &t
				break
			}
		}

		items, _ := models.GetItemsByTag(tagID)
		c.HTML(http.StatusBadRequest, "items.html", gin.H{
			"Collection": collection,
			"Tag":        currentTag,
			"AllTags":    tags,
			"Items":      items,
			"Error":      err.Error(),
			"Name":       name,
			"Description": description,
		})
		return
	}

	c.Redirect(http.StatusFound, "/collections/"+strconv.Itoa(collectionID)+"/tags/"+tagIDStr)
}

func HandleEditItem(c *gin.Context) {
	_, collectionID, collection, err := getAndVerifyCollection(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	tagIDStr := c.Param("tag_id")
	itemIDStr := c.Param("item_id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid item ID")
		return
	}

	name := c.PostForm("name")
	description := c.PostForm("description")
	tagIDs := parseTagIDs(c)

	err = models.UpdateItem(itemID, name, description, tagIDs)
	if err != nil {
		tags, _ := models.GetTags(collectionID)
		var currentTag *models.Tag
		tagID, _ := strconv.Atoi(tagIDStr)
		for _, t := range tags {
			if t.ID == tagID {
				currentTag = &t
				break
			}
		}
		items, _ := models.GetItemsByTag(tagID)
		c.HTML(http.StatusBadRequest, "items.html", gin.H{
			"Collection": collection,
			"Tag":        currentTag,
			"AllTags":    tags,
			"Items":      items,
			"Error":      "Failed to update item: " + err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/collections/"+strconv.Itoa(collectionID)+"/tags/"+tagIDStr)
}

func HandleDeleteItem(c *gin.Context) {
	_, collectionID, _, err := getAndVerifyCollection(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	tagIDStr := c.Param("tag_id")
	itemIDStr := c.Param("item_id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid item ID")
		return
	}

	err = models.DeleteItem(itemID)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to delete item")
		return
	}

	c.Redirect(http.StatusFound, "/collections/"+strconv.Itoa(collectionID)+"/tags/"+tagIDStr)
}
