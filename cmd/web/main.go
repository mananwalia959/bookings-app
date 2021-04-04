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
	var config config.AppConfig
	tc, err := render.BuildTemplateCache()
	if err != nil {
		log.Fatal("Unable to load Cache")
	}

	config.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Application Started on Port %s ", portnumber)

	_ = http.ListenAndServe(portnumber, nil)
}
