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
		false,
	},
	Route{
		"Objects",
		"GET",
		"/objects",
		ObjectsHandler,
		false,
	},
	Route{
		"Users",
		"GET",
		"/users",
		UsersHandler,
		false,
	},
	Route{
		"Categories",
		"GET",
		"/categories",
		CategoriesHandler,
		false,
	},
	Route{
		"Groups",
		"GET",
		"/groups",
		GroupsHandler,
		false,
	},
	Route{
		"Stocks",
		"GET",
		"/stocks",
		StocksHandler,
		false,
	},
	Route{
		"Items",
		"GET",
		"/items",
		ItemsHandler,
		false,
	},

	//specific routes
	Route{
		"Object",
		"GET",
		"/objects/id={object_id}",
		ObjectHandler,
		false,
	},
	Route{
		"Object",
		"GET",
		"/objects/cat={category_id}",
		ObjectHandler,
		false,
	},
	Route{
		"Item",
		"GET",
		"/items/id={item_id}",
		ItemHandler,
		false,
	},
}
