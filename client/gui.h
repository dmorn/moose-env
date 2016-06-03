#ifndef gui_h
#define gui_h
#include <iostream>
#include <vector>
#include <cpr/cpr.h>
#include <json.hpp>
#include "item.h"
#include <termios.h>
#include <unistd.h>
#include "user.h"
//#include "category.h"
//#include "menuItem.h"

using namespace std;
using json = nlohmann::json;
class Gui{
	
	#define MENU_ITEMS 15
	#define DISPLAY_WIDTH 66
	#define POPUP_WIDTH 48
	


	#define LOGIN "LGI"
	#define MAIN_MENU "SMM"
	#define ITEM_LIST "SIL"
	#define CATEGORY_LIST "SCL"
	#define STOCK_LIST "SSL"
	#define ITEM_PAGE "SIP"
	#define ADD_ITEM_PAGE "AIP"
	#define ADD_STOCK_PAGE "ASP"
	#define OBJ_BY_CAT_LIST "OBC"
	#define BUY_ITEM_PAGE "BIP"
	#define PROFILE "PRO"
	#define WISHLIST "PRO"
	#define TEXT_POPUP "TPP"

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
		void addItemPage(Item item);
		void itemPage(Item item);
		void updateScrollPos();
		bool hasResult(string query);
		bool popupYesNo(string text);
		int popupNumber(string text);
		string popupInput(string text);
		void popupMessage(string text);
		string centerText(string text, int width);
		string fillWithSpace(int cnt);
		string limitText(string text);
		bool isNumber(string s);
		json getJson(string content);
		json postJson(string content, json data);
		json postJsonNoToken(string content, json data);
		string title, footer;
		int selectedMenuItem, tmpSelectedMenuItem;
		int scrollPos;
		string currMenu;
		int currCategoryId;
		std::vector<Item> items;
		Item selectedItem;
		bool addItem, addItemToStock;
		User user;
};

#endif
