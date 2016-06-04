#include "element_types.h"

Category::Category (string name, int id, string description, int parent_id) :
	name(name),
	id(id),
	description(description),
	parent_id(parent_id),
	function("SCL")
{ }


const string& Category::getName() const {
	return name;
}

const string& Category::getFunction() const {
	return function;
}

const int Category::getId() const{
	return id;
}

const string& Category::getDescription() const{
	return description;
}

const int Category::getParentId() const{
	return parent_id;
}

// MENU_ITEM
MenuItem::MenuItem (string name) :
	name(name),
	function("NIL")
{ }
MenuItem::MenuItem (string name, string function) :
	name(name),
	function(function)
{ }


const string& MenuItem::getName() const {
	return name;
}

const string& MenuItem::getFunction() const {
	return function;
}

// ITEM_ITEM

ItemItem::ItemItem() { };

ItemItem::ItemItem(string name, int id, string description, int coins, int quantity, int stock_id, int object_id) :
	name(name),
	id(id),
	description(description),
	coins(coins),
	quantity(quantity),
	stock_id(stock_id),
	object_id(object_id),
	function("SIP")
{ }

const string& ItemItem::getName() const {
	return name;
}
const string& ItemItem::getFunction() const {
	return function;
}
const string& ItemItem::getDescription() const{
	return description;
}

const int ItemItem::getId() const{
	return id;
}

const int ItemItem::getCoins() const{
	return coins;
}
const int ItemItem::getQuantity() const{
	return quantity;
}
const int ItemItem::getStockId() const{
	return stock_id;
}
const int ItemItem::getObjectId() const{
	return object_id;
}

// OBJECT

Object::Object() { };

Object::Object(string name, int id, string description) :
	name(name),
	id(id),
	description(description),
	function("AIS")
{ }

const string& Object::getName() const {
	return name;
}
const string& Object::getFunction() const {
	return function;
}

const string& Object::getDescription() const{
	return description;
}

const int Object::getId() const{
	return id;
}

// STOCK

Stock::Stock() { };

Stock::Stock(string name, int id, string location) :
	name(name),
	id(id),
	location(location),
	function("IBS")
{ }

const string& Stock::getName() const {
	return name;
}
const string& Stock::getFunction() const {
	return function;
}

const string& Stock::getLocation() const{
	return location;
}

const int Stock::getId() const{
	return id;
}


