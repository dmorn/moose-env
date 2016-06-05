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
const string& Category::getText() const {
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
MenuItem::MenuItem (string text) :
	text(text),
	function("NIL")
{ }
MenuItem::MenuItem (string text, string function) :
	text(text),
	function(function)
{ }


const string& MenuItem::getText() const {
	return text;
}

const string& MenuItem::getFunction() const {
	return function;
}

// ITEM

Item::Item() { };

Item::Item(string name, int id, string description, int coins, int quantity, string stock, int object_id, int status, string link) :
	name(name),
	id(id),
	description(description),
	coins(coins),
	quantity(quantity),
	stock(stock),
	object_id(object_id),
	status(status),
	function("SIP"),
	link(link)
{
	text = to_string(id) + ": " + name + "\t" + to_string(coins) + " coins; " + to_string(quantity) + "x in ";
	switch(status){
		case 1: text += "stock"; break;
		case 2: text += "pending list"; break;
		case 3: text += "wishlist"; break;
	}
	if(stock != "")
		text+= " @"+stock;

}

Item::Item(string text, string name, int id, string description, int coins, int quantity, string stock, int object_id, int status, string link) :
	text(text),
	name(name),
	id(id),
	description(description),
	coins(coins),
	quantity(quantity),
	stock(stock),
	object_id(object_id),
	status(status),
	function("SIP"),
	link(link)
{ }

const string& Item::getText() const {
	return text;
}
const string& Item::getName() const {
	return name;
}
const string& Item::getFunction() const {
	return function;
}
const string& Item::getDescription() const{
	return description;
}

const int Item::getId() const{
	return id;
}

const int Item::getCoins() const{
	return coins;
}
const int Item::getQuantity() const{
	return quantity;
}
const string& Item::getStock() const{
	return stock;
}
const int Item::getObjectId() const{
	return object_id;
}
const int Item::getStatus() const{
	return status;
}
const string& Item::getLink() const{
	return link;
}

// OBJECT

Object::Object() { };

Object::Object(string name, int id, string description) :
	name(name),
	id(id),
	description(description),
	function("AIS")
{ }
const string& Object::getText() const {
	return name;
}
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
const string& Stock::getText() const {
	return name;
}
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


