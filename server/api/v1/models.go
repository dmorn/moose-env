package main

//BaseInfo that most models implement
type BaseInfo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Balance  int    `json:"balance"`
	Type     int    `json:"type"`
	GroupId  int    `json:"group_id"`
}

type Users []User

type BaseUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type token struct {
	Token string `json:"token"`
}

type StockTaker struct {
	User  `json:"user"`
	Stock `json:"stock"`
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
	*Category  `json:"category"`
}

type Objects []Object

type Category struct {
	BaseInfo
	ParentId int `json:"parent_id"`
}

type Categories []Category

type Item struct {
	Id       int    `json:"id"`
	Status   int    `json:"status"`
	ObjectId int    `json:"object_id"`
	Link     string `json:"link"`
	Coins    int    `json:"coins"`
	Quantity int    `json:"quantity"`
	StockId  int    `json:"stock_id"`
	*Stock   `json:"stock"`
	*Object  `json:"object"`
}

type Items []Item

//receipt stuff

type Receipt struct {
	Data      string `json:"data"`
	Signature string `json:"signature"`
}

type ItemShorter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Coins    int    `json:"coins_payed"`
	Quantity int    `json:"quantity"`
	StockId  int    `json:"stock_id"`
	ObjectId int    `json:"object_id"`
}
