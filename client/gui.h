#ifndef gui_h
#define gui_h
#include<iostream>

using namespace std;
class Gui{
	
	#define MAX_ELEMENTS 10
	#define ITEM_LIST 0
	#define STOCK_LIST 1

	public:
		Gui();
		void print();
		void addElement(string elem, string function);
		void clearElements();
		void update(int keycode);
		void mainMenu();
		
	private:
		void list(int type, int page);
		string title;
		string elements[MAX_ELEMENTS][2];
		int elementCnt;
		int selectedElement;

};

#endif
