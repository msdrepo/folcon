package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "hello! this is go web server.")
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "friend"
	}
	fmt.Fprintf(w, "welcome, %s!", name)
}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/user", userHandler)

	fmt.Println("the server is running on port 8080. - http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("error : failed to start the server : %v\n", err)
	}
}
