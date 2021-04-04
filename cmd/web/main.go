package main

import (
	"fmt"
	"net/http"

	"github.com/mananwalia959/bookings-app/pkg/handlers"
)

var portnumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Application Started on Port %s ", portnumber)

	_ = http.ListenAndServe(portnumber, nil)
}
