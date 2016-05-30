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

//GET
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
			return nil, err
		}
		obj.Category, _ = GetCategory(obj.CategoryId)
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
			return nil, err
		}
		cats = append(cats, cat)
	}
	return &cats, nil
}

func GetGroups() (*Groups, error) {

	const query = "select * from `group`"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	groups := Groups{}
	for rows.Next() {
		group := Group{}

		if err := rows.Scan(&group.Id, &group.Name, &group.Description); err != nil {
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
			return nil, err
		}
		stocks = append(stocks, stock)
	}
	return &stocks, nil
}

func GetItems() (*Items, error) {

	const query = `select * from item;`
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	items := Items{}
	for rows.Next() {
		item := Item{}

		if err := rows.Scan(&item.Id, &item.Coins, &item.Status, &item.Quantity, &item.ObjectId, &item.StockId); err != nil {
			return nil, err
		}
		item.Object, _ = GetObject(item.ObjectId)
		items = append(items, item)
	}
	return &items, nil
}

//specific getters
func GetObject(id int) (*Object, error) {

	query := fmt.Sprintf("select object_id, name, description, category_id from object where object_id = %d", id)

	object := Object{}
	err := db.QueryRow(query).
		Scan(&object.Id, &object.Name, &object.Description, &object.CategoryId)

	if err != nil {
		return nil, err
	}
	object.Category, err = GetCategory(object.CategoryId)
	return &object, err
}

func GetObjectByCategory(categoryID int) (*Objects, error) {

	query := fmt.Sprintf("select object_id, name, description, category_id from object where category_id = %d", categoryID)

	objects := Objects{}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		object := Object{}
		if err := rows.Scan(&object.Id, &object.Name, &object.Description, &object.CategoryId); err != nil {
			return nil, err
		}
		object.Category, _ = GetCategory(object.CategoryId)
		objects = append(objects, object)
	}
	return &objects, err
}

func GetCategory(id int) (*Category, error) {

	query := fmt.Sprintf("select category_id, parent_id, name, description from category where category_id = %d", id)

	cat := Category{}
	err := db.QueryRow(query).
		Scan(&cat.Id, &cat.ParentId, &cat.Name, &cat.Description)

	return &cat, err
}

func GetCategoriesWithSubcategories(id int) (*Categories, error) {

	query := fmt.Sprintf("select * from category where isSubCategoryOf(category_id, %d) = 1", id)
	categories := Categories{}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		category := Category{}

		if err := rows.Scan(&category.Id, &category.ParentId, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return &categories, err
}

func GetCategoriesWithParent(id int) (*Categories, error) {

	var query string
	if id == 0 {
		query = "select * from category where parent_id IS NULL"
	} else {
		query = fmt.Sprintf("select * from category where parent_id = %d", id)
	}

	categories := Categories{}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		category := Category{}

		if err := rows.Scan(&category.Id, &category.ParentId, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return &categories, err
}

func GetItem(id int) (*Item, error) {

	query := fmt.Sprintf("select item_id, coins, status, quantity, object_id, stock_id from item where item_id = %d", id)

	item := Item{}

	err := db.QueryRow(query).
		Scan(&item.Id, &item.Coins, &item.Status, &item.Quantity, &item.ObjectId, &item.StockId)

	if err != nil {
		return nil, err
	}

	item.Object, err = GetObject(item.ObjectId)
	return &item, err
}

func GetItemByCategory(categoryID int) (*Items, error) {

	query := fmt.Sprintf("select item_id from group_items where category_id = %d", categoryID)
	items := Items{}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		//take every id
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		item, _ := GetItem(id)
		items = append(items, *item)
	}
	return &items, err
}

//GetItemsWithCategoriesAndSubcategories gets every item with the requested category_id and every item each subcategory recursevly
func GetItemsWithCategoriesAndSubcategories(id int) (*Items, error) {

	//means that the category is a super category.
	if id == 0 {
		return GetItems()
	}

	query := fmt.Sprintf("select item_id from group_items where isSubCategoryOf(category_id, %d) = 1", id)
	items := Items{}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		//take every id
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		item, _ := GetItem(id)
		items = append(items, *item)
	}
	return &items, err
}

//POST

//for testing
//curl -H "Content-Type: application/json" -X POST -d '{"description":"test object", "name": "yolo", "category_id":2}' http://localhost:8080/object
func PostObject(object *Object) error {

	query := fmt.
		Sprintf("INSERT INTO `object` (`name`, `description`, `category_id`) VALUES ('%s', '%s', %d);",
		object.Name, object.Description, object.CategoryId)

	fmt.Println(query)
	_, err := db.Query(query)
	return err
}

func PostItem(item *Item) error {

	query := fmt.
		Sprintf("INSERT INTO `item` (`item_id`, `coins`, `status`, `quantity`, `object_id`, `stock_id`) VALUES (%d,%d,%d,%d,%d,%d);",
		item.Id, item.Coins, item.Status, item.Quantity, item.ObjectId, item.StockId)

	_, err := db.Query(query)
	return err
}
