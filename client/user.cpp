#include "user.h"

User::User() { 
	id=-1;	
};

User::User(string token) { 
	this->token=token;
};

User::User(int id, string username, string email, string name, string surname, int balance, int type, int group_id, string token) {
	this->id=id;
	this->username=username;
	this->email=email;
	this->name=name;
	this->surname=surname;
	this->balance=balance;
	this->type=type;
	this->group_id=group_id;
	this->token=token;
}

bool User::isValid() {
	return id!=-1;
}

int User::getId(){
	return id;
}

string User::getUsername(){
	return username;
}

string User::getEmail(){
	return email;
}

string User::getName(){
	return name;
}

string User::getSurname(){
	return surname;
}

int User::getBalance(){
	return balance;
}

int User::getType(){
	return type;
}

int User::getGroupId(){
	return group_id;
}

string User::getToken() {
	return token;
}
