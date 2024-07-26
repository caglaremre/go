package main

import (
	"fmt"
	"mywebapp/pkg/handlers"
	"net/http"
)

const portnumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf("starting application on %s\n", portnumber)
	http.ListenAndServe(portnumber, nil)
}
