package main

import (
	// "fmt"
	// "encoding/json"
	"log"
	"net/http"

	"tryGorm/database"

	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	database.User()
	handleRequest()
}

const handleRequest() {

	router := mux.NewRouter()

	router.HandleFunc("/cars", database.GetCars).Methods("GET")
	router.HandleFunc("/cars/{id}", database.GetCar).Methods("GET")
	router.HandleFunc("/drivers/{id}", database.GetDriver).Methods("GET")
	router.HandleFunc("/cars/{id}", database.DeleteCar).Methods("DELETE")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))

}