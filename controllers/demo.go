package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DemoItem struct {
	ID   int64  `db:"id, primarykey, autoincrement" json:"id"`
	Name string `db:"name" json:"name"`
}

func All(c *gin.Context) {
	// Assuming db is your database instance or a package-level variable.
	// Retrieve all dummy data of DemoItem from the database
	items := []DemoItem{
		{ID: 1, Name: "Item 1"},
		{ID: 2, Name: "Item 2"},
		{ID: 3, Name: "Item 3"},
	}

	// Return the items as JSON response
	c.JSON(http.StatusOK, items)
}
