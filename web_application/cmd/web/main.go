package main

import (
	"fmt"
	"log"
	"mywebapp/pkg/config"
	"mywebapp/pkg/handlers"
	"mywebapp/pkg/render"
	"net/http"
)

const portnumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc

	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("starting application on %s\n", portnumber)
	http.ListenAndServe(portnumber, nil)
}
