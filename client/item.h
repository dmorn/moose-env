#ifndef item_h
#define item_h
#include <iostream>
#include <vector>


//BASE: (vector<params>)
//ITEM: (name, id, description) (coins, quantity, stock, obj_id)
//CATEGORY: (name, id, description, parent_id)
//MENU ITEM: (name, function)
//MENU ITEM: (name, <vector>params)

using namespace std;
class Item{

	#define ITEM_ITEM 0
	#define CATEGORY_ITEM 1
	#define MENU_ITEM 2

	public:
		Item();
		Item(vector<string> params);
		Item(string name, int id, string description);
		Item(string name, int id, string description, int coins, int quantity, int stock_id, int object_id);
		Item(string name, int id, string description, int parent_id);
		Item(string name, string function);
		Item(string name, vector<string> params);
		int getId();
		void setId(int id);
		string getName();
		string getDescription();
		string getFunction();
		int getParentId();
		vector<string> getParams();

		int getCoins();
		int getQuantity();
		int getStockId();
		int getObjectId();

	private:
		vector<string> params;
		int item_type;
};

#endif
