#include "gui.h"

Gui::Gui(){
	
	menuItemCnt=0;
	selectedMenuItem=0;
	tmpSelectedMenuItem=-1;
	page=0;
	title="Test";
	currMenu = MAIN_MENU;
	selectedItem = NULL;
    categoryParentId=0;
}

void Gui::addMenuItem(MenuItem item)
{
	if(menuItemCnt < MAX_MENU_ITEMS)
		menuItems[menuItemCnt++]=item;
}


void Gui::clearMenu()
{
	menuItemCnt=0;
	selectedMenuItem=0;
}

void Gui::update(int keycode) {
	//65 U , 68 L , 67 R, 66 D, 127 BACK, 10 ENTER
	
	if(currMenu == ITEM_LIST && (keycode == 67 || keycode == 68)){
		clearMenu();
		selectedMenuItem = 0;
		
		if(keycode == 67 && items.size() > (page + 1) * MAX_MENU_ITEMS)
			page++;
		if(keycode == 68 && page > 0)
			page--;
		
		list();
	}	

	if(keycode == 65)
		selectedMenuItem = (selectedMenuItem <= 0) ? menuItemCnt-1 : --selectedMenuItem;

	else if(keycode == 66)
		selectedMenuItem = (selectedMenuItem >= menuItemCnt-1) ? 0 : ++selectedMenuItem;

	else if(keycode == 127)
	{
		clearMenu();
		if(currMenu == ITEM_PAGE){
			
			if(tmpSelectedMenuItem != -1) {
				selectedMenuItem = tmpSelectedMenuItem;
				tmpSelectedMenuItem = -1;
			}
			list(ITEM_LIST);
		}
		else{
			selectedMenuItem=0;
			if(currMenu == CATEGORY_LIST)
				selectedMenuItem=1;
			if(currMenu == STOCK_LIST)
				selectedMenuItem=2;
			currMenu = MAIN_MENU;
			mainMenu();
		}
	}
	else if(keycode == 10)
	{
        string newMenu = menuItems[selectedMenuItem].getFunction();
		if(newMenu == ITEM_LIST) 
		{
			clearMenu();
			list(newMenu);			
		}
        else if(newMenu == CATEGORY_LIST) {
            if(currMenu == newMenu)
                categoryParentId = items[selectedMenuItem].getId();
			clearMenu();
			list(newMenu);		
        }
		else if(newMenu == ITEM_PAGE)
		{
            currMenu = menuItems[selectedMenuItem].getFunction();
			tmpSelectedMenuItem = selectedMenuItem;
			clearMenu();
			itemPage(items.at(page*MAX_MENU_ITEMS+selectedMenuItem));
		}
	}
	print();
}

void Gui::mainMenu(){
	
   	addMenuItem(MenuItem("Item List", ITEM_LIST));
   	addMenuItem(MenuItem("Category list",CATEGORY_LIST));
   	addMenuItem(MenuItem("Stock list",STOCK_LIST));
}

void Gui::itemPage(Item item){
	
	title = "Item nr." + to_string(item.getId()) + " - " + item.getName();
	addMenuItem(MenuItem(item.getDescription(),"asd"));

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

		for(int i=page*MAX_MENU_ITEMS; i<page*MAX_MENU_ITEMS+MAX_MENU_ITEMS; i++)
		{
			if(i>=items.size()) break;
	   		addMenuItem(MenuItem(items.at(i).getName(),ITEM_PAGE));
		}
	}	

	if(currMenu == CATEGORY_LIST){
		title = "Categories - Page "+to_string(page);
        
		auto response = cpr::Get(cpr::Url{"http://localhost:8080/categories/parent_id="+to_string(categoryParentId)});
		auto json = nlohmann::json::parse(response.text);	
		items.clear();
		for (auto& item : json) {
			items.push_back(Item(categoryParentId,item["name"],item["description"],(int)item["id"]));
		}

		for(int i=page*MAX_MENU_ITEMS; i<page*MAX_MENU_ITEMS+MAX_MENU_ITEMS; i++)
		{
			if(i>=items.size()) break;
	   		addMenuItem(MenuItem(items.at(i).getName(),CATEGORY_LIST));
		}
	}	
	else if(currMenu == STOCK_LIST){
		title = "Stocks - Page "+page;
	   	addMenuItem(MenuItem("Stock",NULL));
	   	addMenuItem(MenuItem("Stock",NULL));
	   	addMenuItem(MenuItem("Stock",NULL));
	   	addMenuItem(MenuItem("Stock",NULL));
	}
}

void Gui::print() {
	
    std::system("clear");	   

    cout << "------------------------------------------------------------------" << endl;
    cout << "\t" + currMenu + " - " + title << endl;
    cout << "------------------------------------------------------------------" << endl;
    int i=0;
    for(i=0; i<menuItemCnt; i++)
	{
		if(i==selectedMenuItem)
			cout << "\033[30;47m"+ menuItems[i].getText() +"\033[0m" << endl;
		else
			cout << menuItems[i].getText() << endl;
	}
	for(i;i<10;i++)
	{
		cout << endl;
	}	
	cout << "-- Use arrow keys to move cursor/change page --" << endl;
	
}
