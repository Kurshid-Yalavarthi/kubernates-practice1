package main

import (
	"log"
	"net/http"

	"bitbucket.org/arcusnext/weather-api/src/handlers"
)

func main() {

	http.HandleFunc("/weatherByLatLongAndDate/", handlers.WeatherByLatLongAndDate)
	log.Println("Server listening on port 3002")
	log.Println("\tRoutes:")
	log.Println("\t\tGET /weatherByLatLongAndDate")
	log.Fatal(http.ListenAndServe(":3002", nil))
}
