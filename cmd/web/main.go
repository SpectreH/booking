package main

import (
	"booking/pkg/config"
	"booking/pkg/handlers"
	"booking/pkg/render"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig
	app.UseCache = false

	var err error
	app.TemplateCache, err = render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	repo := handlers.SetNewRepo(&app)
	handlers.SetNewHandlers(repo)

	render.SetNewTemplates(&app)

	fmt.Println("Staring application on port", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
