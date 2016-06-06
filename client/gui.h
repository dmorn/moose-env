#ifndef gui_h
#define gui_h
#include <iostream>
#include <vector>
#include <cpr/cpr.h>
#include <json.hpp>
#include <termios.h>
#include <unistd.h>
#include <fstream>  
#include "element_types.h"
#include "user.h"

using namespace std;
using json = nlohmann::json;
class Gui{
	
	#define MENU_ELEMENTS 15
	#define DISPLAY_WIDTH 66
	#define POPUP_WIDTH 48
	

	#define URL "http://localhost:8080/"
	#define LOGIN "LGI"
	#define MAIN_MENU "SMM"
	#define ITEM_LIST "SIL"
	#define WISH_LIST "SWL"
	#define PENDING_LIST "SPL"
	#define ITEM_BY_STOCK "IBS"
	#define CATEGORY_LIST "SCL"
	#define STOCK_LIST "SSL"
	#define MY_STOCK_LIST "MSL"
	#define ITEM_PAGE "SIP"
	#define ADD_ITEM_PAGE "AIP"
	#define ADD_ITEM_SELECTED "AIS"
	#define ADD_STOCK_ITEM_PAGE "ASP"
	#define OBJ_BY_CAT_LIST "OBC"
	#define BUY_ITEM_PAGE "BIP"
	#define ORDER_ITEM_PAGE "OIP"
	#define CONFIRM_ITEM_PAGE "CIP"
	#define PROFILE "PRO"
	#define ADD_USER "ADU"
	#define ADD_STOCK "ADS"
	#define ADD_CATEGORY "ADC"
	#define ADD_OBJECT "ADO"
	#define ADD_BALANCE "ADB"
	#define WISHLIST "PRO"
	#define TEXT_POPUP "TPP"
	#define NO_FUNCTION "NIL"

	public:
		Gui();
		void print();
		void clearMenu();
		void update(int keycode);
		void mainMenu();
		
	private:
		void list();
		void list(string list_type);
		void addItemPage(int obj_no);
		void itemPage(int item_no);
		void addCategoryPage();
		void addObjectPage();
		void updateScrollPos();
		void addBalance(string username, int amount);
		void withdrawBalance(string username, int amount);
		bool hasResult(string query);
		bool popupYesNo(string text);
		int popupNumber(string text);
		string popupInput(string text);
		void popupMessage(string text);
		string centerText(string text, int width);
		string fillWithSpace(int cnt);
		string limitText(string text);
		bool isNumber(string s);
		bool isJson(cpr::Response r);
		json getJson(string content);
		json postJson(string content);
		json postJson(string content, json data);
		json postJsonNoToken(string content, json data);
		string title, footer;
		int selectedMenuItem, tmpSelectedMenuItem;
		int scrollPos;
		string currMenu;
		int currCategoryId;
		std::vector<Element*> elements;
		Item* currentItem;
		bool addItem, addItemToStock, addCategory, addObject;
		Stock * selectedStock;
		User user;
};

#endif
