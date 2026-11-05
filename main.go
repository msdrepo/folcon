package main

import (
	"fmt"

	"github.com/msdrepo/folcon/routes"
)

func main() {

	app := routes.Route()

	fmt.Println("Folcon server running on http://localhost:8080...")
	app.Run(":8080")
}
