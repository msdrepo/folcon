package main

import (
	"fmt"

	folcon "github.com/msdrepo/framework"
)

func main() {
	app := folcon.New()

	// Global Middleware
	app.Use(func(c *folcon.Context) {
		fmt.Println("[Global Logger] Incoming Request")
		c.Next()
	})

	// Home Route
	app.GET("/", func(c *folcon.Context) {
		c.String(200, "Welcome to Folcon Framework!")
	})

	// API Group with a specific prefix
	api := app.Group("/api/v1")
	{
		api.GET("/users", func(c *folcon.Context) {
			c.String(200, "List of all users from Folcon!")
		})

		api.POST("/users", func(c *folcon.Context) {
			c.String(201, "User created successfully!")
		})
	}

	fmt.Println("Folcon server running on http://localhost:8080...")
	app.Run(":8080")
}
