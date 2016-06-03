#include "item.h"

Item::Item(){ }

Item::Item(vector<string> params){
	this->params=params;
	item_type=-1;
}

Item::Item(string name, int id, string description){
	this->params.push_back(name);
	this->params.push_back(to_string(id));
	this->params.push_back(description);
	item_type=ITEM_ITEM;
}

Item::Item(string name, int id, string description, int coins, int quantity, int stock_id, int object_id){
	this->params.push_back(name);
	this->params.push_back(to_string(id));
	this->params.push_back(description);
	this->params.push_back(to_string(coins));
	this->params.push_back(to_string(quantity));
	this->params.push_back(to_string(stock_id));
	this->params.push_back(to_string(object_id));
	item_type=ITEM_ITEM;
}

Item::Item(string name, int id, string description, int parent_id){
	this->params.push_back(name);
	this->params.push_back(to_string(id));
	this->params.push_back(description);
	this->params.push_back(to_string(parent_id));
	item_type=CATEGORY_ITEM;
}

Item::Item(string name){
	this->params.push_back(name);
	this->params.push_back("nil");
	item_type=MENU_ITEM;
}

Item::Item(string name, string function){
	this->params.push_back(name);
	this->params.push_back(function);
	item_type=MENU_ITEM;
}

Item::Item(string name, string function, string param1){
	this->params.push_back(name);
	this->params.push_back(function);
	this->params.push_back(param1);
	item_type=MENU_ITEM;
}

Item::Item(string name, vector<string> params){
	this->params=params;
	this->params.insert(this->params.begin(),name);
	item_type=MENU_ITEM;
}

vector<string> Item::getParams(){
	return params;
}

string Item::getName(){
	return params.front();
}

int Item::getId(){

	if(item_type == MENU_ITEM)
		return 0;
	return stoi(params.at(1));
}

void Item::setId(int id){
	params.at(1) = to_string(id);
}

string Item::getFunction(){
	if(item_type==ITEM_ITEM)
		return "SIP";
	else if(item_type==CATEGORY_ITEM)
		return "SCL";
	return params.at(1);
}

string Item::getParamAt(int i){
	if(i>=0 && i<=params.size())
		return params.at(i);
	else return NULL;
}

string Item::getDescription(){
	return params.at(2);
}

int Item::getParentId(){
	return stoi(params.at(3));
}

int Item::getCoins(){
	return stoi(params.at(3));
}
int Item::getQuantity(){
	return stoi(params.at(4));
}
int Item::getStockId(){
	return stoi(params.at(5));
}
int Item::getObjectId(){
	return stoi(params.at(6));
}







