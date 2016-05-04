package main

import "net/http"

type Route struct {
	Name         string
	Method       string
	Pattern      string
	HandlerFunc  http.HandlerFunc
	AuthRequired bool
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
		true,
	},
}
