// main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DispatchOrder struct {
	ID          int    `json:"id"`
	Customer    string `json:"customer"`
	Destination string `json:"destination"`
	Status      string `json:"status"`
}

var orders []DispatchOrder

func main() {
	router := gin.Default()

	// Create a new dispatch order
	router.POST("/dispatch", func(c *gin.Context) {
		var newOrder DispatchOrder
		if err := c.BindJSON(&newOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Simulate order ID generation 
		newOrder.ID = len(orders) + 1
		newOrder.Status = "Created"

		orders = append(orders, newOrder)

		c.JSON(http.StatusCreated, newOrder)
	})

	// Get all dispatch orders
	router.GET("/dispatch", func(c *gin.Context) {
		c.JSON(http.StatusOK, orders)
	})

	router.Run(":8080")
}
