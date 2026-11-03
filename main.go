package main

import (
	"fmt"
	"net/http"

	res "github.com/msdrepo/folcon/responses"
	folcon "github.com/msdrepo/framework"
)

func main() {
	app := folcon.New()

	// Global Middleware
	app.Use(func(c *folcon.Context) {
		// fmt.Println("[Global Logger] Incoming Request")
		c.Next()
	})

	// API Group with a specific prefix
	api := app.Group("/api/v1")
	{
		api.GET("/users", func(c *folcon.Context) {
			users := []string{"Alice", "Bob", "Charlie"}

			response := res.SuccessResponse{
				Status:     "success",
				StatusCode: http.StatusOK,
				Message:    "data fetched successfully",
				Data:       users,
			}

			c.JSON(200, response)

		})

		api.POST("/users", func(c *folcon.Context) {
			c.String(201, "User created successfully!")
		})
	}

	app.GET("/", func(c *folcon.Context) {

		response := res.SuccessResponse{
			Status:     "success",
			StatusCode: http.StatusOK,
			Message:    "Welcome to Folcon Framework!",
			Data:       nil,
		}

		c.JSON(200, response)

	})

	app.GET("/test", func(c *folcon.Context) {

		response := res.SuccessResponse{
			Status:     "success",
			StatusCode: http.StatusOK,
			Message:    "test",
			Data:       "test data",
		}

		c.JSON(200, response)

	})

	fmt.Println("Folcon server running on http://localhost:8080...")
	app.Run(":8080")
}
