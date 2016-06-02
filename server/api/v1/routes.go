package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Objects",
		"GET",
		"/objects",
		ObjectsHandler,
	},
	Route{
		"Users",
		"GET",
		"/users",
		UsersHandler,
	},
	Route{
		"Categories",
		"GET",
		"/categories",
		CategoriesHandler,
	},
	Route{
		"Groups",
		"GET",
		"/groups",
		GroupsHandler,
	},
	Route{
		"Stocks",
		"GET",
		"/stocks",
		StocksHandler,
	},
	Route{
		"Items",
		"GET",
		"/items",
		ItemsHandler,
	},

	//specific routes
	Route{
		"Object",
		"GET",
		"/objects/id={object_id}",
		ObjectHandler,
	},
	Route{
		"User",
		"GET",
		"/users/id={user_id}",
		UserHandler,
	},
	Route{
		"Object",
		"GET",
		"/objects/cat={category_id}",
		ObjectHandler,
	},
	Route{
		"Object",
		"GET",
		"/objects/start_cat_id={category_id}",
		ObjectsWithCategoriesAndSubcategoriesHandler,
	},
	Route{
		"Item",
		"GET",
		"/items/id={item_id}",
		ItemHandler,
	},
	Route{
		"Item",
		"GET",
		"/items/cat={category_id}",
		ItemHandler,
	},
	Route{
		"Item",
		"GET",
		"/items/start_cat_id={category_id}",
		ItemsWithCategoriesAndSubcategoriesHandler,
	},
	Route{
		"Category",
		"GET",
		"/categories/id={category_id}",
		CategoryHandler,
	},
	Route{
		"Categories",
		"GET",
		"/categories/start_id={category_id}",
		CategoriesWithSubcategoriesHandeler,
	},
	Route{
		"Categories",
		"GET",
		"/categories/parent_id={parent_id}",
		CategoriesWithParentHandler,
	},

	//posts
	//login related stuff
	Route{
		"Login",
		"POST",
		"/login",
		LoginHandler,
	},
	Route{
		"Register",
		"POST",
		"/register",
		RegistrationHandler,
	},

	Route{
		"Object",
		"POST",
		"/object",
		PostObjectHandler,
	},
	Route{
		"Item",
		"POST",
		"/item",
		PostItemHandler,
	},
}
