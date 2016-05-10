package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

//Getter handlers
func ObjectsHandler(w http.ResponseWriter, r *http.Request) {

	if objects, err := GetObjects(); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		json.NewEncoder(w).Encode(objects)
	}
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {

	if users, err := GetUsers(); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		json.NewEncoder(w).Encode(users)
	}
}

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {

	if cats, err := GetCategories(); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		json.NewEncoder(w).Encode(cats)
	}
}

func GroupsHandler(w http.ResponseWriter, r *http.Request) {

	if groups, err := GetGroups(); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		json.NewEncoder(w).Encode(groups)
	}
}

func StocksHandler(w http.ResponseWriter, r *http.Request) {

	if stocks, err := GetStocks(); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		json.NewEncoder(w).Encode(stocks)
	}
}

func ItemsHandler(w http.ResponseWriter, r *http.Request) {

	if items, err := GetItems(); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		json.NewEncoder(w).Encode(items)
	}
}

//getter handlers specific

func ObjectHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var objectId int
	var err error
	if objectId, err = strconv.Atoi(vars["object_id"]); err != nil {
		log.Fatal(err)
		return
	}

	if object, err := GetObject(objectId); err != nil {

		log.Fatal(err)
		http.Error(w, err.Error(), 404)

	} else {

		if object.Id > 0 {
			json.NewEncoder(w).Encode(object)
		}
	}
}
