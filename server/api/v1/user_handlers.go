package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//getters
func UsersHandler(w http.ResponseWriter, r *http.Request) {

	if users, err := GetUsers(); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		json.NewEncoder(w).Encode(users)
	}
}

//getter handlers specific
func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var userID int
	var user *User
	var err error

	userID, err = strconv.Atoi(vars["user_id"])

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if user, err = GetUser(userID); err != nil {
		http.Error(w, err.Error(), 404)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

//curl -H "Content-Type: application/json" -X POST -d '{"username":"matthias", "password": "test"}' http://localhost:8080/login
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var candidateUser *BaseUser
	err := decoder.Decode(&candidateUser)
	if err != nil {
		fmt.Println("Error Decoding Form")
		http.Error(w, err.Error(), 500)
		return
	}

	//we have the base user, now get User with that username
	user, err := GetUserByUsername(candidateUser.Username)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	byteCandidateUserPassword := []byte(candidateUser.Password)

	h256 := sha256.New()
	h256.Write(byteCandidateUserPassword)

	hashedByteCandidateUserPassword := h256.Sum(nil)

	//fmt.Printf("Password db: %s\n", user.Password)
	//fmt.Printf("Password sent: %x\n", hashedByteCandidateUserPassword)

	if user.Password == hex.EncodeToString(hashedByteCandidateUserPassword) {
		//user is authorized! generate and send token
		t := memStore.NewToken(user.Username)
		data := token{Token: t.String()}
		json.NewEncoder(w).Encode(data)

	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

//curl -H "Content-Type: application/json" -X POST -d '{"username":"matex", "password": "hello", "email": "ciao@ciao.com", "name": "phil", "surname": "hexx", "group_id": 1}' http://localhost:8080/register
func RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var user *User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println("Error Decoding Form")
		http.Error(w, err.Error(), 500)
		return
	}

	if err := PostUser(user); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		user, _ := GetUserByUsername(user.Username)
		json.NewEncoder(w).Encode(user) //should return 201
	}

}
