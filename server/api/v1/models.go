package main

import "database/sql"

//BaseInfo that most models implement
type BaseInfo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type User struct {
	Id         int            `json:"id"`
	Username   string         `json:"username"`
	Password   string         `json:"password"`
	Email      string         `json:"email"`
	Name       string         `json:"name"`
	Surname    string         `json:"surname"`
	Balance    int            `json:"balance"`
	Type       int            `json:"type"`
	VerifyCode sql.NullString `json:"verify_code"`
	GroupId    sql.NullString `json:"group_id"`
	Salt       string         `json:"salt"`
}

type Users []User

type StockTaker struct {
	User
	Stock
}

type Stock struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Stocks []Stock

type Group struct {
	BaseInfo
}

type Groups []Group

type JSonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

//Object is the abstraction of an Item
type Object struct {
	BaseInfo
	CategoryId int `json:"category_id"`
}

type Objects []Object

type Category struct {
	BaseInfo
	ParentId sql.NullInt64
}

type Categories []Category

/*


type Category struct {
	BaseInfo
	Parent_id int `json:"parent_id"`
}

type Group struct {
	BaseInfo
	Items []Item `json:"items"`
}

type Item struct {
	BaseInfo
	Category_id int    `json:"category_id"`
	Coins       int    `json:"coins"`
	Quantity    int    `json:"quantity"`
	Stock_name  string `json:"stock_name"`
	Group_id    int    `json:"group_id"`
	Stock_id    int    `json:"stock_id"`
}

type Request struct {
	Group_id int `json:"group_id"`
	User_id  int `json:"user_id"`
}
*/
