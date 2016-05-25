package main

import (
	"encoding/json"
	"fmt"
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
	var objectID int
	var categoryID int

	var err1 error
	var err2 error

	objectID, err1 = strconv.Atoi(vars["object_id"])
	categoryID, err2 = strconv.Atoi(vars["category_id"])

	if err1 != nil && err2 != nil {
		http.Error(w, err1.Error(), 500)
		return
	}

	var object *Object
	var objects *Objects
	var err error

	if objectID > 0 {
		object, err = GetObject(objectID)
	} else if categoryID > 0 {
		objects, err = GetObjectByCategory(categoryID)
	}

	if err != nil {
		http.Error(w, err.Error(), 404)
	} else {
		if object != nil {
			json.NewEncoder(w).Encode(object)
		}
		if objects != nil {
			json.NewEncoder(w).Encode(objects)
		}
	}
}

func ItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var itemID int
	var item *Item
	var err error

	itemID, err = strconv.Atoi(vars["item_id"])

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if item, err = GetItem(itemID); err != nil {
		http.Error(w, err.Error(), 404)
	} else {
		json.NewEncoder(w).Encode(item)
	}
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var catID int
	var cat *Category
	var err error

	catID, err = strconv.Atoi(vars["category_id"])

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if cat, err = GetCategory(catID); err != nil {
		http.Error(w, err.Error(), 404)
	} else {
		json.NewEncoder(w).Encode(cat)
	}
}

func CategoriesWithSubcategoriesHandeler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var catID int
	var err error

	if catID, err = strconv.Atoi(vars["category_id"]); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if categories, err := GetCategoriesWithSubcategories(catID); err != nil {
		http.Error(w, err.Error(), 404)
	} else {
		json.NewEncoder(w).Encode(categories)
	}
}

func CategoriesWithParentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var catID int
	var err error

	if catID, err = strconv.Atoi(vars["parent_id"]); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if categories, err := GetCategoriesWithParent(catID); err != nil {
		http.Error(w, err.Error(), 404)
	} else {
		json.NewEncoder(w).Encode(categories)
	}

}

//post handlers

func PostItemHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var item *Item
	err := decoder.Decode(&item)
	if err != nil {
		fmt.Println("Error Decoding Form")
		http.Error(w, err.Error(), 500)
		return
	}

	if err = PostItem(item); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		items, _ := GetItems()
		json.NewEncoder(w).Encode(items) //should return 201
	}
}

func PostObjectHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var object *Object
	err := decoder.Decode(&object)
	if err != nil {
		fmt.Println("Error Decoding Form")
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(object)
	if err = PostObject(object); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		objects, _ := GetObjects()
		json.NewEncoder(w).Encode(objects) //should return 201
	}
}
