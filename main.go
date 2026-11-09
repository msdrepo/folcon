package main

import (
	"fmt"

	"github.com/msdrepo/folcon/routes"

	folcon "github.com/msdrepo/framework"
)

func main() {

	app := routes.Route()

	if err := folcon.ConnectMySql(); err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}

	fmt.Println("Folcon server running on http://localhost:8080...")
	app.Run(":8080")
}
