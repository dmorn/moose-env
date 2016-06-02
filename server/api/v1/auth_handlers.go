package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/freehaha/token-auth"
	"github.com/gorilla/mux"
)

//Getter handlers
func UsersHandler(w http.ResponseWriter, r *http.Request) {

	if users, err := GetUsers(); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		json.NewEncoder(w).Encode(users)
	}
}

func ObjectsHandler(w http.ResponseWriter, r *http.Request) {

	if objects, err := GetObjects(); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		json.NewEncoder(w).Encode(objects)
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
func UserHandler(w http.ResponseWriter, r *http.Request) {

	user, err := getUserFromToken(r)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(user)
}

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

func ObjectsWithCategoriesAndSubcategoriesHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var catID int
	var err error

	if catID, err = strconv.Atoi(vars["category_id"]); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if objects, err := GetObjectsWithCategoriesAndSubcategories(catID); err != nil {
		http.Error(w, err.Error(), 404)
	} else {
		json.NewEncoder(w).Encode(objects)
	}
}

func ItemHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var itemID int
	var categoryID int

	var err1 error
	var err2 error

	itemID, err1 = strconv.Atoi(vars["item_id"])
	categoryID, err2 = strconv.Atoi(vars["category_id"])

	if err1 != nil && err2 != nil {
		http.Error(w, err1.Error(), 500)
		return
	}

	var item *Item
	var items *Items
	var err error

	if itemID > 0 {
		item, err = GetItem(itemID)
	} else if categoryID > 0 {
		items, err = GetItemByCategory(categoryID)
	}

	if err != nil {
		http.Error(w, err.Error(), 404)
	} else {
		if item != nil {
			json.NewEncoder(w).Encode(item)
		}
		if items != nil {
			json.NewEncoder(w).Encode(items)
		}
	}
}

func ItemsWithCategoriesAndSubcategoriesHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var catID int
	var err error

	if catID, err = strconv.Atoi(vars["category_id"]); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if items, err := GetItemsWithCategoriesAndSubcategories(catID); err != nil {
		http.Error(w, err.Error(), 404)
	} else {
		json.NewEncoder(w).Encode(items)
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

//for testing
//curl -H "Content-Type: application/json" -X POST -d '{"description":"test object", "name": "yolo", "category_id":2}' http://localhost:8080/object
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

//curl -H "Content-Type: application/json" -X POST http://localhost:8080/add_stock_taker/daniel/1
func AddStockTakerHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var username string
	var stock_id int
	var err error
	var ok bool

	if stock_id, err = strconv.Atoi(vars["stock_id"]); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if username, ok = vars["username"]; ok == false {
		http.Error(w, err.Error(), 500)
		return
	}

	//check user state -> must be a stock taker himself
	if _, err := isUserStockTaker(r); err != nil {
		http.Error(w, err.Error(), 500)
	} else {

		if candidateStockTaker, err := GetUserByUsername(username); err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			if err := AddUserToStockTakers(candidateStockTaker, stock_id); err != nil {
				http.Error(w, err.Error(), 500)
			} else {
				//success
				json.NewEncoder(w).Encode(candidateStockTaker)
			}
		}
	}
}

//private helpers
func getUserFromToken(r *http.Request) (*User, error) {

	token := tauth.Get(r)
	username := token.Claims("id").(string)
	return GetUserByUsername(username)

}

func isUserStockTaker(r *http.Request) (*User, error) {

	token := tauth.Get(r)
	username := token.Claims("id").(string)

	if user, err := GetUserByUsername(username); err != nil {
		return nil, err
	} else {
		if err := CheckUserIsStockTaker(user); err != nil {
			return nil, err
		}
		return user, nil
	}
}
