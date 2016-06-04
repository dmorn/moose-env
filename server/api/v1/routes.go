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
		"Objects",
		"GET",
		"/objects",
		ObjectsHandler,
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
		"/items/{status}/{stock_id}/{start_cat_id}",
		ItemsHandlerStatusStockCat,
	},
	Route{
		"Items",
		"GET",
		"/items",
		ItemsHandler,
	},
	Route{
		"Items",
		"GET",
		"/items/wishlist",
		WishlistHandler,
	},
	Route{
		"Items",
		"GET",
		"/items/pending",
		PendingItemsHandler,
	},
	Route{
		"Items",
		"GET",
		"/items/stock",
		StockItemsHandler,
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
		"/user",
		UserHandler,
	},
	Route{
		"Users",
		"GET",
		"/users/group_id={group_id}",
		UsersHandler,
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
		"Stocks",
		"GET",
		"/stock/id={stock_id}",
		StockHandler,
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

	//others
	Route{
		"User",
		"POST",
		"/add_stock_taker/{username}/{stock_id}",
		AddStockTakerHandler,
	},
	Route{
		"User",
		"POST",
		"/balance/{username}/withdraw={amount}",
		UserWithdrawBalance,
	},
	Route{
		"User",
		"POST",
		"/balance/{username}/add={amount}",
		UserAddBalance,
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
		"/purchase/{item_id}/{quantity}",
		PurchaseItemHandler,
	},
	Route{
		"Item",
		"POST",
		"/item",
		PostItemHandler,
	},
	Route{
		"Item",
		"POST",
		"/purchase_wishlist/{stock_id}",
		PurchaseWishlistHandler,
	},
	Route{
		"Item",
		"POST",
		"/put_into_stock/{stock_id}",
		PutPurchasesIntoStockHandler,
	},
	Route{
		"Item",
		"POST",
		"/put_into_stock/{username}",
		PutNewItemIntoStockHandler,
	},

	Route{
		"Category",
		"POST",
		"/category",
		PostCategoryHandler,
	},
}
