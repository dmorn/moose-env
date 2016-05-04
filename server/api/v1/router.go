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

	green := color.New(color.FgRed)

	const title = `   
   __  _______  ____  ________  _____  ___   __
  /  |/  / __ \/ __ \/ __/ __/ / __/ |/ / | / /
 / /|_/ / /_/ / /_/ /\ \/ _/  / _//    /| |/ / 
/_/  /_/\____/\____/___/___/ /___/_/|_/ |___(_)
                                               
	`
	green.Printf("%s\n", title)

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name, route.AuthRequired)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
