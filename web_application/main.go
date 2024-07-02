package main

import (
	"fmt"
	"net/http"
)

const portnumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Printf("starting application on %s\n", portnumber)
	http.ListenAndServe(portnumber, nil)
}
