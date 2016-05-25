#ifndef gui_h
#define gui_h
#include <iostream>
#include <vector>
#include <cpr/cpr.h>
#include <json.hpp>
#include "item.h"
#include "menuItem.h"

using namespace std;
class Gui{
	
	#define MAX_ELEMENTS 10
	#define MAIN_MENU "SMM"
	#define ITEM_LIST "SIL"
	#define STOCK_LIST "SSL"
	#define ITEM_PAGE "SIP"

	public:
		Gui();
		void print();
		void addElement(MenuItem elem);
		void clearElements();
		void update(int keycode);
		void mainMenu();
		
	private:
		void list();
		void list(string list_type);
		void itemPage(Item item);
		string title;
		MenuItem elements[MAX_ELEMENTS];
		int elementCnt;
		int selectedElement, tmpSelectedElement;
		int page;
		string currMenu;
		std::vector<Item> items;
		Item * selectedItem;
};

#endif
