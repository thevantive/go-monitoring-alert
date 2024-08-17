package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define the route for /monitoring-database/alert
	r.POST("/monitoring/alert", func(c *gin.Context) {

		// menyiapkan
		type RequestBody struct {
			Connections uint    `json:"connections" binding:"required"`
			MemoryUsage float32 `json:"memory_usage" binding:"required"`
		}

		// mengambil request body
		var requestBody RequestBody
		if err := c.ShouldBindBodyWithJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request body",
			})
		}

		// melakukan log response
		fmt.Printf("Received alert - Connections: %d, Memory Usage: %f%%\n", requestBody.Connections, requestBody.MemoryUsage)

		// response berhasil
		c.JSON(http.StatusOK, gin.H{
			"message": "Alert received successfully",
			"received": gin.H{
				"connections":  requestBody.Connections,
				"memory_usage": requestBody.MemoryUsage,
			},
		})
	})

	// Start the server on localhost:8081
	r.Run(":8081")
}
