package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mananwalia959/bookings-app/pkg/config"
	"github.com/mananwalia959/bookings-app/pkg/handlers"
	"github.com/mananwalia959/bookings-app/pkg/render"
	"github.com/mananwalia959/bookings-app/pkg/routes"
)

var portnumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.BuildTemplateCache()
	if err != nil {
		log.Fatal("Unable to load Cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)

	handlers.SetRepo(repo)
	render.AddAppConfig(&app)

	fmt.Println(fmt.Sprintf("Application Started on Port %s ", portnumber))

	// _ = http.ListenAndServe(portnumber, routes)

	srv := &http.Server{
		Addr:    portnumber,
		Handler: routes.GetRoutes(&app),
	}

	err = srv.ListenAndServe()
	// since ListenAndServe blocks forever
	// log. Fatal will not be called unless an error comes
	log.Fatal(err)
}
