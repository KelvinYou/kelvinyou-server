package controller

import (
	"net/http"
	"strconv"

	"github.com/KelvinYou/kelvinyou-server/models"
	"github.com/gin-gonic/gin"
)

func All(c *gin.Context) {
	// Assuming db is your database instance or a package-level variable.
	// Retrieve all dummy data of DemoItem from the database
	items := []models.DemoItem{
		{ID: 1, Name: "Item 1"},
		{ID: 2, Name: "Item 2"},
		{ID: 3, Name: "Item 3"},
	}

	// Return the items as JSON response
	c.JSON(http.StatusOK, items)
}

func One(c *gin.Context) {
	// Get the ID from the URL parameter
	id := c.Param("id")

	// Convert the ID to an integer
	itemID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	// Define the dummy data
	items := []models.DemoItem{
		{ID: 1, Name: "Item 1"},
		{ID: 2, Name: "Item 2"},
		{ID: 3, Name: "Item 3"},
	}

	// Convert the ID to int64
	itemID64 := int64(itemID)

	// Search for the item with the specified ID in the dummy data
	var item *models.DemoItem
	for _, i := range items {
		if i.ID == itemID64 {
			item = &i
			break
		}
	}

	// Check if the item exists
	if item == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// Return the item as a JSON response
	c.JSON(http.StatusOK, item)
}
