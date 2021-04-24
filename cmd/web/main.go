package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mananwalia959/bookings-app/pkg/config"
	"github.com/mananwalia959/bookings-app/pkg/handlers"
	"github.com/mananwalia959/bookings-app/pkg/render"
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

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Application Started on Port %s ", portnumber))

	_ = http.ListenAndServe(portnumber, nil)
}
