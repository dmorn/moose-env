package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//NewRouter constructor
func NewRouter() *mux.Router {

	//display server header
	fmt.Println("\033[H\033[2J") //clear screen

	const title = `   
   __  _______  ____  ________  _____  ___   __
  /  |/  / __ \/ __ \/ __/ __/ / __/ |/ / | / /
 / /|_/ / /_/ / /_/ /\ \/ _/  / _//    /| |/ / 
/_/  /_/\____/\____/___/___/ /___/_/|_/ |___(_)
                                               
	`
	fmt.Printf("%s\n", title)

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		if route.Name == "Login" || route.Name == "Test" {
			handler = MMiddleware(handler, route.Name)
		} else {
			handler = MAuthMiddleware(handler, route.Name)
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
