package main

import (
	"net/http"

	"github.com/anshika97/weatherapi/controllers"
	"github.com/gorilla/mux"
)

func main() {
	port := ":3000"

	println("Starting server at 3000")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/weather/{city}/", controllers.GetWeather)

	err := http.ListenAndServe(port, router)
	if err != nil {
		println("Error while listening to Port")
	}
}
