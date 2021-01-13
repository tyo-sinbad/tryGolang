package database

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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

var db *gorm.DB
var err error

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

func User() {
	db, err = gorm.Open("postgres", "host=ziggy.db.elephantsql.com port=5432 user=klfmmivf dbname=klfmmivf sslmode=disable password=AjmwzYbNDjW4X0hbl0dryGuECZ1acBCn")

	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("running ...")
	}

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

func GetCars(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=ziggy.db.elephantsql.com port=5432 user=klfmmivf dbname=klfmmivf sslmode=disable password=AjmwzYbNDjW4X0hbl0dryGuECZ1acBCn")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var cars []Car
	db.Find(&cars)
	json.NewEncoder(w).Encode(&cars)
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=ziggy.db.elephantsql.com port=5432 user=klfmmivf dbname=klfmmivf sslmode=disable password=AjmwzYbNDjW4X0hbl0dryGuECZ1acBCn")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	params := mux.Vars(r)
	var car Car
	db.First(&car, params["id"])
	json.NewEncoder(w).Encode(&car)
}

func GetDriver(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=ziggy.db.elephantsql.com port=5432 user=klfmmivf dbname=klfmmivf sslmode=disable password=AjmwzYbNDjW4X0hbl0dryGuECZ1acBCn")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	params := mux.Vars(r)
	var driver Driver
	var cars []Car
	db.First(&driver, params["id"])
	db.Model(&driver).Related(&cars)
	driver.Cars = cars
	json.NewEncoder(w).Encode(&driver)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=ziggy.db.elephantsql.com port=5432 user=klfmmivf dbname=klfmmivf sslmode=disable password=AjmwzYbNDjW4X0hbl0dryGuECZ1acBCn")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	params := mux.Vars(r)
	var car Car
	db.First(&car, params["id"])
	db.Delete(&car)

	var cars []Car
	db.Find(&cars)
	json.NewEncoder(w).Encode(&cars)
}