package main

import (
	"GO/trevor/bookings-31/pkg/config"
	"GO/trevor/bookings-31/pkg/handlers"
	"GO/trevor/bookings-31/pkg/render"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":3000"

// main is the main function
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false // false: (DEV mode)read cache everytime.

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	tmp := fmt.Sprintf("Staring application on port %s", portNumber)
	fmt.Println(tmp)
	// _ = http.ListenAndServe(portNumber, nil)

	svr := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = svr.ListenAndServe()
	log.Fatal(err)
}
