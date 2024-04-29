package main

import (
	"fmt"

	controller "github.com/KelvinYou/kelvinyou-server/controllers"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware ...
// CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	port := "8080"

	// Create a new Gin router
	r := gin.Default()

	r.Use(CORSMiddleware())

	api := r.Group("/api")
	{
		v0 := api.Group("/v0")
		{
			demo := new(controller.DemoController)

			v0.GET("/demo", demo.All)
			v0.GET("/demo/:id", demo.One)

		}
	}

	// Start the server
	r.Run(":" + port)
}
