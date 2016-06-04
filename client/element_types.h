#ifndef element_types_h
#define element_types_h
#include <iostream>

using namespace std;

class Element{
	public:
		virtual const string& getName() const = 0;
		virtual const string& getFunction() const = 0;
};

class Category : public Element {
	public:
		Category (string name, int id, string description, int parent_id);

		virtual const string& getName() const;
		virtual const string& getFunction() const;
		const string& getDescription() const;
		const int getId() const;
		const int getParentId() const;

	private:
		string name, description, function;
		int id, parent_id;
};

class MenuItem : public Element {
	public:
		MenuItem (string name);
		MenuItem (string name, string function);

		virtual const string& getName() const;
		virtual const string& getFunction() const;
		//const string& getParam() const;

	private:
		string name, function;
};

class ItemItem : public Element {
	public:
		ItemItem();
		ItemItem (string name, int id, string description, int coins, int quantity, int stock_id, int object_id);

		virtual const string& getName() const;
		virtual const string& getFunction() const;
		const string& getDescription() const;
		const int getId() const;
		const int getCoins() const;
		const int getQuantity() const;
		const int getStockId() const;
		const int getObjectId() const;

	private:
		string name, function, description;
		int id, coins, quantity, stock_id, object_id;
};

class Object : public Element {
	public:
		Object();
		Object (string name, int id, string description);

		virtual const string& getName() const;
		virtual const string& getFunction() const;
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

		virtual const string& getName() const;
		virtual const string& getFunction() const;
		const string& getLocation() const;
		const int getId() const;

	private:
		string name, function, location;
		int id;
};

#endif
