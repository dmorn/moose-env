#include "Gui.h"

Gui::Gui(){
	
	elementCnt=0;
	selectedElement=0;
	title="Test";
}

void Gui::addElement(string elem, string function)
{
	if(elementCnt < MAX_ELEMENTS){
		elements[elementCnt][0]=elem;
		elements[elementCnt++][1]=function;
	}
}

void Gui::clearElements()
{
	elementCnt=0;
}

void Gui::update(int keycode) {
	//72 U , 75 L , 77 R, 80 D, 8 BACK, 13 ENTER
	if(keycode == 72){
		selectedElement--;
			if(selectedElement < 0)
				selectedElement = elementCnt-1;
	}	
	else if(keycode == 80){
		selectedElement++;
			if(selectedElement >= elementCnt)
				selectedElement = 0;
	}
	else if(keycode == 8)
	{
		mainMenu();
		selectedElement=0;
	}
	else if(keycode == 13)
	{
		selectedElement=0;
		string func = elements[selectedElement][1];
		
		if(func == "show_main_menu") 
			mainMenu();
		else if(func == "show_item_list") 
			list(ITEM_LIST,0);
		else if(func == "show_stock_list")
			list(STOCK_LIST,0);
	}
	print();
}

void Gui::mainMenu(){
	
	clearElements();
   	addElement("Item list","show_item_list");
   	addElement("Stock list","show_stock_list");
   	print();
}

void Gui::list(int type, int page){
	
	clearElements();
	if(type == ITEM_LIST){
	   	addElement("Item","show_stock_list");
	   	addElement("Item","show_stock_list");
	   	addElement("Item","show_stock_list");
	   	addElement("Item","show_stock_list");
	}	
	else if(type == STOCK_LIST){
	   	addElement("Stock","show_stock_list");
	   	addElement("Stock","show_stock_list");
	   	addElement("Stock","show_stock_list");
	   	addElement("Stock","show_stock_list");
	}
   	print();
}

void Gui::print() {
	
  	system("cls");
    cout << "------------------------------------------------------------------" << endl;
    cout << "\t" + title << endl;
    cout << "------------------------------------------------------------------" << endl;
    int i=0;
    for(i=0; i<elementCnt; i++)
	{
		if(i==selectedElement)
		{
			cout << ">|-";
			cout << i;
    		cout << "-| " + elements[i][0] << endl;
		}
		else
		{
			cout << "| ";
			cout << i;
    		cout << " | " + elements[i][0] << endl;
		}
	}
	
}
