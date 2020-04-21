package controllers

import (
	"net/http"

	"github.com/anshika97/weatherapi/api"
	"github.com/gorilla/mux"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := vars["city"]

	data, err := api.GetData(city)
	if err != nil {
		w.Write([]byte("404 not found"))
	}

	w.Write([]byte(data))
}
