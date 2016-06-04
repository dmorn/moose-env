package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
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

func GetUsers(id int) (*Users, error) {

	query := fmt.Sprintf("select * from user where group_id=%d", id)
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	users := Users{}
	for rows.Next() {
		user := User{}

		if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Name,
			&user.Surname, &user.Balance, &user.Type, &user.GroupId); err != nil {
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

func GetCategoriesIDs() ([]int, error) {

	const query = `select category_id from category;`
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	cats := make([]int, 0)
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		cats = append(cats, id)
	}
	return cats, nil
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

		if err := rows.Scan(&item.Id, &item.Coins, &item.Status, &item.Quantity, &item.Link, &item.ObjectId, &item.StockId); err != nil {
			return nil, err
		}
		item.Object, _ = GetObject(item.ObjectId)
		item.Stock, _ = GetStock(item.StockId)
		items = append(items, item)
	}
	return &items, nil
}

func GetItemsWithStatus(id int) (*Items, error) {

	query := fmt.Sprintf("select * from item where status = %d", id)
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	items := Items{}
	for rows.Next() {
		item := Item{}

		if err := rows.Scan(&item.Id, &item.Coins, &item.Status, &item.Quantity, &item.Link, &item.ObjectId, &item.StockId); err != nil {
			return nil, err
		}
		item.Object, _ = GetObject(item.ObjectId)
		item.Stock, _ = GetStock(item.StockId)
		items = append(items, item)
	}
	return &items, nil
}

//specific getters

func GetUser(id int) (*User, error) {

	query := fmt.Sprintf("select * from user where user_id = %d", id)

	user := User{}
	err := db.QueryRow(query).Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Name,
		&user.Surname, &user.Balance, &user.Type, &user.GroupId)

	return &user, err
}

func GetUserByUsername(u string) (*User, error) {

	query := fmt.Sprintf("select * from `user` where username LIKE '%s'", u)

	user := User{}
	err := db.QueryRow(query).Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Name,
		&user.Surname, &user.Balance, &user.Type, &user.GroupId)

	return &user, err
}

func CheckUserIsStockTaker(user *User) ([]int, error) {

	query := fmt.Sprintf("select user_id from user_stock where user_id = %d", user.Id)

	if rows, err := db.Query(query); err != nil {
		return nil, err
	} else {
		list := make([]int, 0)
		for rows.Next() {
			var el int
			rows.Scan(&el)
			list = append(list, el)
		}

		if len(list) > 0 {
			return list, err
		}
		return nil, errors.New("You're not in the Fight Club boy")
	}
}

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

//GetItemsWithCategoriesAndSubcategories gets every item with the requested category_id and every item each subcategory recursevly
func GetObjectsWithCategoriesAndSubcategories(id int) (*Objects, error) {

	query := fmt.Sprintf("select * from object where isSubCategoryOf(category_id, %d) = 1", id)

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

	query := fmt.Sprintf("select * from category where parent_id = %d", id)
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

	query := fmt.Sprintf("select * from item where item_id = %d", id)

	item := Item{}

	err := db.QueryRow(query).
		Scan(&item.Id, &item.Coins, &item.Status, &item.Quantity, &item.Link, &item.ObjectId, &item.StockId)

	if err != nil {
		return nil, err
	}

	item.Object, err = GetObject(item.ObjectId)
	item.Stock, err = GetStock(item.StockId)
	return &item, err
}

func PurchaseItem(id int, quantity int, user *User) (*Item, error) {

	if item, err := GetItem(id); err != nil {
		return nil, err
	} else {

		//check that the user has enough money first
		if user.Balance-item.Coins < 0 {
			return nil, errors.New(fmt.Sprintf("Not Enough money bro, your balance is %d", user.Balance))

		}

		itemsAfterPossiblePurchase := item.Quantity - quantity

		//check that the quantity required is not too much
		if itemsAfterPossiblePurchase < 0 {
			return nil, errors.New(fmt.Sprintf("We don't have enough items to satisfy your request: %d", item.Quantity))
		}

		if itemsAfterPossiblePurchase == 0 {
			//we have the item, now delete it
			if err := DeleteItem(id); err != nil {
				return nil, err
			}
		} else {
			//just update the quantity
			if err := UpdateItemQuantity(id, itemsAfterPossiblePurchase); err != nil {
				return nil, err
			}
		}

		//withdraw money from user
		if err := WithdrawAmountToUserBalance(user, item.Coins); err != nil {
			//re insert the item or reupdate quantity
			if itemsAfterPossiblePurchase == 0 {
				PostItem(item, item.Status)
			} else {
				UpdateItemQuantity(id, item.Quantity)
			}
			return nil, err
		}

		//everything is fine
		return item, nil
	}
}

func DeleteItem(id int) error {

	query := fmt.
		Sprintf("DELETE FROM item WHERE item_id=%d LIMIT 1", id)
	_, err := db.Query(query)
	return err
}

func UpdateItemQuantity(id int, newQuantity int) error {
	query := fmt.
		Sprintf("UPDATE `item` SET quantity=%d WHERE item_id=%d", newQuantity, id)
	_, err := db.Query(query)
	return err
}

func GetStock(id int) (*Stock, error) {

	query := fmt.Sprintf("select * from stock where stock_id = %d", id)

	stock := Stock{}
	err := db.QueryRow(query).
		Scan(&stock.Id, &stock.Name, &stock.Location)

	return &stock, err
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

func GetItemsWithStatusStockCategory(status int, stock_id int, start_cat_id int) (*Items, error) {

	query := fmt.
		Sprintf("select item_id from group_items where isSubCategoryOf(category_id, %d) = 1 AND status=%d AND stock_id=%d",
		start_cat_id, status, stock_id)
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
func AddUserToStockTakers(user *User, id int) error {

	//update type into db
	query := fmt.
		Sprintf("UPDATE `user` SET type=2 WHERE user_id=%d", user.Id)

	_, err := db.Query(query)

	query = fmt.
		Sprintf("INSERT INTO `user_stock` (`user_id`, `stock_id`) VALUES (%d, %d)",
		user.Id, id)

	_, err = db.Query(query)
	return err
}

func PostUser(user *User) error {

	hasher := sha256.New()
	hasher.Write([]byte(user.Password))
	hash := hex.EncodeToString(hasher.Sum(nil))

	user.Type = 1 //is a normal user

	query := fmt.
		Sprintf("INSERT INTO `user` (`username`, `password`, `email`, `name`, `surname`, `balance`, `type`, `group_id`) VALUES ('%s', '%s', '%s', '%s', '%s', 0, 1, %d)",
		user.Username, hash, user.Email, user.Name, user.Surname, user.GroupId)

	fmt.Println(query)
	_, err := db.Query(query)
	return err
}

func PostObject(object *Object) error {

	query := fmt.
		Sprintf("INSERT INTO `object` (`name`, `description`, `category_id`) VALUES ('%s', '%s', %d);",
		object.Name, object.Description, object.CategoryId)

	fmt.Println(query)
	_, err := db.Query(query)
	return err
}

func PostItem(item *Item, status int) error {

	query := fmt.
		Sprintf("INSERT INTO `item` (`item_id`, `coins`, `status`, `quantity`,`link`, `object_id`, `stock_id`) VALUES (%d,%d,%d,%d,'%s',%d,%d);",
		item.Id, item.Coins, item.Status, item.Quantity, item.Link, item.ObjectId, item.StockId)

	_, err := db.Query(query)
	return err
}

func PostCategory(category *Category) error {

	var query string

	query = fmt.
		Sprintf("INSERT INTO `category` (`parent_id`, `name`, `description`) VALUES (%d, '%s', '%s');",
		category.ParentId, category.Name, category.Description)

	fmt.Println(query)
	_, err := db.Query(query)
	return err
}

//PATCH
func AddAmountToUserBalance(user *User, amount int) error {

	balance := amount + user.Balance
	query := fmt.
		Sprintf("UPDATE `user` SET balance=%d WHERE user_id=%d", balance, user.Id)

	_, err := db.Query(query)
	return err
}

func WithdrawAmountToUserBalance(user *User, amount int) error {

	balance := user.Balance - amount
	query := fmt.
		Sprintf("UPDATE `user` SET balance=%d WHERE user_id=%d", balance, user.Id)

	_, err := db.Query(query)
	return err
}

func UpdateItemsStatusToPending(stock_id int) error {

	query := fmt.
		Sprintf("UPDATE `item` SET status=2 WHERE status=3 AND stock_id=%d", stock_id)
	_, err := db.Query(query)
	return err
}

func PutItemsIntoStock(stock_id int) error {

	query := fmt.
		Sprintf("UPDATE `item` SET status=1 WHERE status=2 AND stock_id=%d", stock_id)
	_, err := db.Query(query)
	return err
}
