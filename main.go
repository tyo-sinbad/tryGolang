package main

import (
	"fmt"
	"encoding/json"
	// "log"
	"net/http"

	// "tryGorm/database"

	// "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	// "github.com/rs/cors"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Driver struct {
	gorm.Model
	Name    string
	License string
	Cars    []Car
}

type Car struct {
	gorm.Model
	Year      int
	Make      string
	ModelName string
	DriverID  int
}

var (
	drivers = []Driver{
		{Name: "Jimmy Johnson", License: "ABC123"},
		{Name: "Howard Hills", License: "XYZ789"},
		{Name: "Craig Colbin", License: "DEF333"},
	}
	cars = []Car{
		{Year: 2000, Make: "Toyota", ModelName: "Tundra", DriverID: 1},
		{Year: 2001, Make: "Honda", ModelName: "Accord", DriverID: 1},
		{Year: 2002, Make: "Nissan", ModelName: "Sentra", DriverID: 2},
		{Year: 2003, Make: "Ford", ModelName: "F-150", DriverID: 3},
	}
)

var db *gorm.DB
var err error

func main() {
	// database.User()
	connection()
	// innitialMigrate()
	handleRequest()
}

func connection() {
	db, err = gorm.Open("postgres", "host=ziggy.db.elephantsql.com port=5432 user=klfmmivf dbname=klfmmivf sslmode=disable password=AjmwzYbNDjW4X0hbl0dryGuECZ1acBCn")
		if err != nil {
			panic("failed to connect database")
		}
		// defer db.Close()
}

func innitialMigrate() {
	connection()
	defer db.Close()

	db.AutoMigrate(&Driver{})
	db.AutoMigrate(&Car{})

	for index := range cars {
		db.Create(&cars[index])
	}

	for index := range drivers {
		db.Create(&drivers[index])
	}
}

func handleRequest() {
	port := ":8080"

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/cars", func(w http.ResponseWriter, r *http.Request) {
		connection()
		defer db.Close()
		var cars []Car
		db.Find(&cars)
		json.NewEncoder(w).Encode(&cars)
	})

	r.Get("/cars/{id}", func(w http.ResponseWriter, r *http.Request) {
		connection()
		defer db.Close()

		params := chi.URLParam(r, "id")
		// ctx := r.Context()
		// key := ctx.Value("key").(string)
		var car Car
		db.First(&car, params)
		json.NewEncoder(w).Encode(&car)
		// respond to the client
		// w.Write([]byte("hi %v, %v", params, key))
		// fmt.Println(key)
	})

	r.Get("/drivers/{id}", func(w http.ResponseWriter, r *http.Request) {
		connection()
		defer db.Close()

		params := chi.URLParam(r, "id")
		var driver Driver
		var cars []Car
		db.First(&driver, params)
		db.Model(&driver).Related(&cars)
		driver.Cars = cars
		json.NewEncoder(w).Encode(&driver)
	})

	r.Get("/cars/{id}", func(w http.ResponseWriter, r *http.Request) {
		connection()
		defer db.Close()

		params := chi.URLParam(r, "id")
		var car Car
		db.First(&car, params)
		db.Delete(&car)

		var cars []Car
		db.Find(&cars)
		json.NewEncoder(w).Encode(&cars)
	})

	fmt.Println("running in port" + port)

	http.ListenAndServe(port, r)

	// router := mux.NewRouter()
	// router.HandleFunc("/cars", database.GetCars).Methods("GET")
	// router.HandleFunc("/cars/{id}", database.GetCar).Methods("GET")
	// router.HandleFunc("/drivers/{id}", database.GetDriver).Methods("GET")
	// router.HandleFunc("/cars/{id}", database.DeleteCar).Methods("DELETE")

	// handler := cors.Default().Handler(router)
	// log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))

}
