#include "gui.h"

Gui::Gui(){
	
	elementCnt=0;
	selectedElement=0;
	tmpSelectedElement=-1;
	page=0;
	title="Test";
	currMenu = MAIN_MENU;
	selectedItem = NULL;
}

void Gui::addElement(MenuItem elem)
{
	if(elementCnt < MAX_ELEMENTS)
		elements[elementCnt++]=elem;
}


void Gui::clearElements()
{
	elementCnt=0;
	selectedElement=0;
}

void Gui::update(int keycode) {
	//65 U , 68 L , 67 R, 66 D, 127 BACK, 10 ENTER
	
	if(currMenu == ITEM_LIST && (keycode == 67 || keycode == 68)){
		clearElements();
		selectedElement = 0;
		
		if(keycode == 67 && items.size() > (page + 1) * MAX_ELEMENTS)
			page++;
		if(keycode == 68 && page > 0)
			page--;
		
		list();
	}	

	if(keycode == 65)
		selectedElement = (selectedElement <= 0) ? elementCnt-1 : --selectedElement;

	else if(keycode == 66)
		selectedElement = (selectedElement >= elementCnt-1) ? 0 : ++selectedElement;

	else if(keycode == 127)
	{
		clearElements();
		if(currMenu == ITEM_PAGE){
			
			if(tmpSelectedElement != -1) {
				selectedElement = tmpSelectedElement;
				tmpSelectedElement = -1;
			}
			list(ITEM_LIST);
		}
		else{
			selectedElement=0;
			if(currMenu == STOCK_LIST)
				selectedElement=1;
			currMenu = MAIN_MENU;
			mainMenu();
		}
	}
	else if(keycode == 10)
	{
        currMenu = elements[selectedElement].getFunction();
		if(currMenu == ITEM_LIST) 
		{
			clearElements();
			list();			
		}
		else if(currMenu == ITEM_PAGE)
		{
			tmpSelectedElement = selectedElement;
			clearElements();
			itemPage(items.at(page*MAX_ELEMENTS+selectedElement));
		}
	}
	print();
}

void Gui::mainMenu(){
	
	title = "Moose Environment";
   	addElement(MenuItem("Item list",ITEM_LIST));
   	addElement(MenuItem("Stock list",STOCK_LIST));
}

void Gui::itemPage(Item item){
	
	title = "Item nr." + to_string(item.getId()) + " - " + item.getName();
	addElement(MenuItem(item.getDescription(),"asd"));

}

void Gui::list(string list_type){
	currMenu = list_type;
    list();
}

void Gui::list(){

	if(currMenu == ITEM_LIST){
		title = "Items - Page "+to_string(page);

		auto response = cpr::Get(cpr::Url{"http://localhost:8080/items"});
		auto json = nlohmann::json::parse(response.text);	
		items.clear();
		for (auto& item : json) {
			nlohmann::json object = item["object"];
			items.push_back(Item((int)object["id"],object["name"],object["description"]));
		}

		for(int i=page*MAX_ELEMENTS; i<page*MAX_ELEMENTS+MAX_ELEMENTS; i++)
		{
			if(i>=items.size()) break;
	   		addElement(MenuItem(items.at(i).getName(),ITEM_PAGE));
		}
	}	
	else if(currMenu == STOCK_LIST){
		title = "Stocks - Page "+page;
	   	addElement(MenuItem("Stock",NULL));
	   	addElement(MenuItem("Stock",NULL));
	   	addElement(MenuItem("Stock",NULL));
	   	addElement(MenuItem("Stock",NULL));
	}
}

void Gui::print() {
	
    std::system("clear");	   

    cout << "------------------------------------------------------------------" << endl;
    cout << "\t" + currMenu + " - " + title << endl;
    cout << "------------------------------------------------------------------" << endl;
    int i=0;
    for(i=0; i<elementCnt; i++)
	{
		if(i==selectedElement)
			cout << "\033[30;47m"+ elements[i].getText() +"\033[0m" << endl;
		else
			cout << elements[i].getText() << endl;
	}
	for(i;i<10;i++)
	{
		cout << endl;
	}	
	cout << "-- Use arrow keys to move cursor/change page --" << endl;
	
}
