package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type car struct {
	Model string `json:"model,omitempty"`
}

var cars []car

func getCars(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(cars)
}

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	cars = append(cars, car{Model: "Ferrari"})
	router.HandleFunc("/cars", getCars).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
