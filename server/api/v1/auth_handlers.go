package main

import (
	"encoding/json"
	"errors"
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

//getter handlers specific
func UserHandler(w http.ResponseWriter, r *http.Request) {

	user, err := GetUserFromToken(r)

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

func StockHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var stockID int
	var err error

	if stockID, err = strconv.Atoi(vars["stock_id"]); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if stock, err := GetStock(stockID); err != nil {
		http.Error(w, err.Error(), 404)
	} else {
		json.NewEncoder(w).Encode(stock)
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

func UserWithdrawBalance(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var username string
	var amount int
	var err error
	var ok bool

	if _, _, err = IsUserStockTaker(r); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if amount, err = strconv.Atoi(vars["amount"]); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if username, ok = vars["username"]; ok == false {
		http.Error(w, err.Error(), 500)
		return
	}

	if user, err := GetUserByUsername(username); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		if err = WithdrawAmountToUserBalance(user, amount); err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			user.Balance = user.Balance - amount
			json.NewEncoder(w).Encode(user) //should return 201
		}
	}
}

//curl -H "Content-Type: application/json" -H "Authorization: Bearer dRBCRdLaRnNGTZGBxhVjYY8f9PM=" -X POST http://localhost:8080/balance/daniel/add=12
func UserAddBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var username string
	var amount int
	var err error
	var ok bool

	if _, _, err = IsUserStockTaker(r); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if amount, err = strconv.Atoi(vars["amount"]); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if username, ok = vars["username"]; ok == false {
		http.Error(w, err.Error(), 500)
		return
	}

	if user, err := GetUserByUsername(username); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		if err = AddAmountToUserBalance(user, amount); err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			user.Balance = user.Balance + amount
			json.NewEncoder(w).Encode(user) //should return 201
		}
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

	if err = PostObject(object); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		objects, _ := GetObjects()
		json.NewEncoder(w).Encode(objects) //should return 201
	}
}

//tests
//curl -H "Content-Type: application/json" -H "Authorization: Bearer Zz_vA49aKmcW_XdoM8A69FKAKS0=" -X POST -d '{"parent_id":{"Int64":1, "Valid":true}, "name": "testcateogyr", "description": "aksjjdas"}' http://localhost:8080/category
func PostCategoryHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var category *Category
	err := decoder.Decode(&category)
	if err != nil {
		fmt.Println("Error Decoding Form")
		http.Error(w, err.Error(), 500)
		return
	}

	//check that the parent category is correct
	if !isParentCategoryValid(category) && category.ParentId.Valid {
		http.Error(w, errors.New("Invalid Parent Id").Error(), http.StatusBadRequest)
		return
	}

	if err = PostCategory(category); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		json.NewEncoder(w).Encode(category) //should return 201
	}
}

//for tests. Remember to use an authenticated token
//curl -H "Content-Type: application/json" -H "Authorization: Bearer G-ibUdic9Zjd0bk3qS5DHQg5ZFs=" -X POST http://localhost:8080/add_stock_taker/daniel/1
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
	if _, _, err := IsUserStockTaker(r); err != nil {
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

//helpers -- usable only in tauth environment
func GetUserFromToken(r *http.Request) (*User, error) {

	token := tauth.Get(r)
	username := token.Claims("id").(string)
	return GetUserByUsername(username)

}

func IsUserStockTaker(r *http.Request) (*User, []int, error) {

	if user, err := GetUserFromToken(r); err != nil {
		return nil, nil, err
	} else {
		if list, err := CheckUserIsStockTaker(user); err != nil {
			return nil, nil, err
		} else {
			return user, list, nil
		}
	}
}

//private helpers

func isParentCategoryValid(category *Category) bool {

	id := category.ParentId
	categories, _ := GetCategoriesIDs()

	for _, val := range categories {
		if val.Int64 == id.Int64 {
			return true
		}
	}
	return false
}
