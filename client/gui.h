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
	
	#define MAX_MENU_ITEMS 10
	#define MAIN_MENU "SMM"
	#define ITEM_LIST "SIL"
	#define CATEGORY_LIST "SCL"
	#define STOCK_LIST "SSL"
	#define ITEM_PAGE "SIP"

	public:
		Gui();
		void print();
		void addMenuItem(MenuItem item);
		void clearMenu();
		void update(int keycode);
		void mainMenu();
		
	private:
		void list();
		void list(string list_type);
		void itemPage(Item item);
		string title;
		MenuItem menuItems[MAX_MENU_ITEMS];
		int menuItemCnt;
		int selectedMenuItem, tmpSelectedMenuItem;
		int page;
		string currMenu;
		int categoryParentId;
		std::vector<Item> items;
		Item * selectedItem;
};

#endif
