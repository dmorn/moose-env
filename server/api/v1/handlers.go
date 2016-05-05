package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
