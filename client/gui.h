#ifndef gui_h
#define gui_h
#include<iostream>
#include <cpr/cpr.h>
#include <json.hpp>

using namespace std;
class Gui{
	
	#define MAX_ELEMENTS 12
	#define MAIN_MENU "SMM"
	#define ITEM_LIST "SIL"
	#define STOCK_LIST "SSL"
	#define ITEM_PAGE "SIP"

	public:
		Gui();
		void print();
		void addElement(string elem, string function, string param_1);
		void clearElements();
		void update(int keycode);
		void mainMenu();
		
	private:
		void list(string type, int page);
		void itemPage(string id);
		string title;
		string elements[MAX_ELEMENTS][10];
		int elementCnt;
		int selectedElement;

};

#endif
