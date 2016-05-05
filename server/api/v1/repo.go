package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

// Initialise Database
func init() {
	connectToDatabse()
}

//connectToDatabse does what you expect it to do. Set enviroment variables first
func connectToDatabse() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to the database
	dbinfo := fmt.Sprintf("%s:%s@/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err = sql.Open("mysql", dbinfo)

	if err != nil {
		log.Fatal(err)
	}
}

//Getters
func GetObjects() (*Objects, error) {

	const query = `select * from object;`

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	objects := Objects{}

	for rows.Next() {
		obj := Object{}

		if err := rows.Scan(&obj.Id, &obj.Name, &obj.Description, &obj.CategoryId); err != nil {
			log.Fatal(err)
			return nil, err
		}
		objects = append(objects, obj)
	}
	return &objects, nil
}

func GetUsers() (*Users, error) {

	const query = `select * from user;`
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	users := Users{}
	for rows.Next() {
		user := User{}

		if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Name,
			&user.Surname, &user.Balance, &user.Type, &user.VerifyCode, &user.Salt, &user.GroupId); err != nil {
			log.Fatal(err)
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}

func GetCategories() (*Categories, error) {

	const query = `select * from category;`
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	cats := Categories{}
	for rows.Next() {
		cat := Category{}

		if err := rows.Scan(&cat.Id, &cat.ParentId, &cat.Name, &cat.Description); err != nil {
			log.Fatal(err)
			return nil, err
		}
		cats = append(cats, cat)
	}
	return &cats, nil
}

//TODO: query does not work
func GetGroups() (*Groups, error) {

	const query = `select * from group;`
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	groups := Groups{}
	for rows.Next() {
		group := Group{}

		if err := rows.Scan(&group.Id, &group.Name, &group.Description); err != nil {
			log.Fatal(err)
			return nil, err
		}
		groups = append(groups, group)
	}
	return &groups, nil
}

func GetStocks() (*Stocks, error) {

	const query = `select * from stock;`
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	stocks := Stocks{}
	for rows.Next() {
		stock := Stock{}

		if err := rows.Scan(&stock.Id, &stock.Name, &stock.Location); err != nil {
			log.Fatal(err)
			return nil, err
		}
		stocks = append(stocks, stock)
	}
	return &stocks, nil
}
