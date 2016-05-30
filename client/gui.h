#ifndef gui_h
#define gui_h
#include <iostream>
#include <vector>
#include <cpr/cpr.h>
#include <json.hpp>
#include "item.h"
#include "category.h"
#include "menuItem.h"

using namespace std;
class Gui{
	
	#define MENU_ITEMS 10
	#define MAIN_MENU "SMM"
	#define ITEM_LIST "SIL"
	#define CATEGORY_LIST "SCL"
	#define STOCK_LIST "SSL"
	#define ITEM_PAGE "SIP"
	#define ADD_ITEM_PAGE "AIP"
	#define OBJ_BY_CAT_LIST "OBC"

	public:
		Gui();
		void print();
		void addMenuItem(Item item);
		void clearMenu();
		void update(int keycode);
		void mainMenu();
		
	private:
		void list();
		void list(string list_type);
		void addItemPage();
		void addItemPage(Item item);
		void itemPage(Item item);
		void updateScrollPos();
		bool hasResult(string query);
		string popupInput(string text);
		void popupMessage(string text);
		string centerText(string text, int width);
		bool isNumber(string s);

		string title;
		int selectedMenuItem, tmpSelectedMenuItem;
		int scrollPos;
		string currMenu;
		int currCategoryId;
		std::vector<Item> items;
		Item * selectedItem;
		bool addItem;
};

#endif
