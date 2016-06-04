#ifndef element_types_h
#define element_types_h
#include <iostream>

using namespace std;

class Element{
	public:
		virtual const string& getText() const = 0;
		virtual const string& getFunction() const = 0;
};

class Category : public Element {
	public:
		Category (string name, int id, string description, int parent_id);

		virtual const string& getText() const;
		virtual const string& getFunction() const;
		const string& getName() const;
		const string& getDescription() const;
		const int getId() const;
		const int getParentId() const;

	private:
		string name, description, function;
		int id, parent_id;
};

class MenuItem : public Element {
	public:
		MenuItem (string text);
		MenuItem (string text, string function);

		virtual const string& getText() const;
		virtual const string& getFunction() const;
		//const string& getParam() const;

	private:
		string text, function;
};

class Item : public Element {
	public:
		Item();
		Item (string name, int id, string description, int coins, int quantity, string stock, int object_id, int status);
		Item (string text, string name, int id, string description, int coins, int quantity, string stock, int object_id, int status);

		virtual const string& getText() const;
		virtual const string& getFunction() const;
		const string& getName() const;
		const string& getDescription() const;
		const string& getStock() const;
		const int getId() const;
		const int getCoins() const;
		const int getQuantity() const;
		const int getObjectId() const;
		const int getStatus() const;

	private:
		string text, name, function, description, stock;
		int id, coins, quantity, object_id, status;
};

class Object : public Element {
	public:
		Object();
		Object (string name, int id, string description);

		virtual const string& getText() const;
		virtual const string& getFunction() const;
		const string& getName() const;
		const string& getDescription() const;
		const int getId() const;

	private:
		string name, function, description;
		int id;
};

class Stock : public Element {
	public:
		Stock();
		Stock (string name, int id, string location);

		virtual const string& getText() const;
		virtual const string& getFunction() const;
		const string& getName() const;
		const string& getLocation() const;
		const int getId() const;

	private:
		string name, function, location;
		int id;
};

#endif
