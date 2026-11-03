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
	// app.Use(func(c *folcon.Context) {
	// 	// fmt.Println("[Global Logger] Incoming Request")
	// 	c.Next()
	// })

	// API Group with a specific prefix
	api := app.Group("/api/v1")

	api.GET("/users", func(c *folcon.Context) {
		users := []string{"Alice", "Bob", "Charlie"}

		response := res.SuccessResponse(http.StatusOK, "data fetched successfully", users, nil)

		c.JSON(200, response)

	})

	api.POST("/users", func(c *folcon.Context) {
		c.String(201, "User created successfully!")
	})

	app.GET("/", func(c *folcon.Context) {

		response := res.SuccessResponse(http.StatusOK, "Welcome to Folcon Framework!", nil, nil)

		c.JSON(200, response)

	})

	app.GET("/test", func(c *folcon.Context) {

		response := res.SuccessResponse(http.StatusOK, "test", "test data", nil)

		c.JSON(200, response)

	})

	api.GET("/products", func(c *folcon.Context) {
		users := []string{"Alice", "Bob", "Charlie"}

		// page = 1, limit = 10, total items = 50
		response := res.SuccessWithPagination(200, "Users fetched successfully", users, 1, 10, 50)

		c.JSON(200, response)
	})

	errorApi := app.Group("/error")

	// error middleware
	errorApi.Use(folcon.ErrorHandler())

	errorApi.GET("/panic", func(c *folcon.Context) {
		var ptr *string
		*ptr = "cause panic" // Triggers nil pointer dereference
	})

	fmt.Println("Folcon server running on http://localhost:8080...")
	app.Run(":8080")
}
