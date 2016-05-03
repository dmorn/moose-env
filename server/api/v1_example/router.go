package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	//display server header
	fmt.Println("\033[H\033[2J") //clear screen

	//MOOSE ENVIRONMENT SERVER - in hexadecimal lol
	green := color.New(color.FgGreen).Add(color.Bold)
	green.Printf("4D 4F 4F 53 45  45 4E 56 49 52 4F 4E 4D 45 4E 54  53 45 52 56 45 52")
	fmt.Println("\n")

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
