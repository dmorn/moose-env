#include "gui.h"

Gui::Gui(){
	
	selectedMenuItem=0;
	tmpSelectedMenuItem=-1;
    scrollPos=0;
	title="Test";
	currMenu = MAIN_MENU;
	selectedItem = NULL;
    currCategoryId=0;
	addItem=false;
}

void Gui::addMenuItem(Item item)
{
	if(items.size() < MENU_ITEMS)
		items.push_back(item);
}


void Gui::clearMenu()
{
	items.clear();
	selectedMenuItem=0;
}


void Gui::updateScrollPos(){
	if(selectedMenuItem < scrollPos)
		scrollPos = selectedMenuItem;
	else if(selectedMenuItem >= scrollPos + MENU_ITEMS)
		scrollPos = selectedMenuItem - MENU_ITEMS+1;
}

void Gui::update(int keycode) {

	//65 U , 68 L , 67 R, 66 D, 127 BACK, 10 ENTER, 32 SPACE, 9 TAB

	if(keycode == 65)
		selectedMenuItem = (selectedMenuItem > 0) ? --selectedMenuItem : items.size() -1;

	else if(keycode == 66)
		selectedMenuItem = (selectedMenuItem < items.size()-1) ? ++selectedMenuItem : 0; 

    else if(keycode == 67){
        selectedMenuItem += MENU_ITEMS;
        if(selectedMenuItem > items.size() - 1)
            selectedMenuItem=items.size()-1;  
    }
    else if(keycode == 68){
        selectedMenuItem -= MENU_ITEMS;
        if(selectedMenuItem < 0) 
			selectedMenuItem = 0;
	}


	else if(keycode == 127)
	{
		if(currMenu == ITEM_PAGE){
			
			if(tmpSelectedMenuItem != -1) {
				selectedMenuItem = tmpSelectedMenuItem;
				tmpSelectedMenuItem = -1;
			}
			clearMenu();
			list(ITEM_LIST);
		}
		else if(currMenu == CATEGORY_LIST && currCategoryId != 0)
		{
			auto response = cpr::Get(cpr::Url{"http://localhost:8080/categories/id="+to_string(items.at(selectedMenuItem).getParentId())});
			auto cat = nlohmann::json::parse(response.text);	
			auto p_id = cat["parent_id"];
            currCategoryId = p_id["Int64"];		
			list();
		}
		else if(currMenu == OBJ_BY_CAT_LIST)
		{
			currMenu  = CATEGORY_LIST;
			clearMenu();
			list();
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
        currMenu = items.at(selectedMenuItem).getFunction();

		if(currMenu == ITEM_LIST) 
			list();	

        else if(currMenu == CATEGORY_LIST) {
			if(items.at(selectedMenuItem).getId() == 0) {

	        	currCategoryId = items.at(selectedMenuItem).getId();
				list();	
			}
			else if(addItem) {
				if(hasResult("objects/cat="+to_string(items.at(selectedMenuItem).getId())))
				{
		        	currCategoryId = items.at(selectedMenuItem).getId();
					currMenu = OBJ_BY_CAT_LIST;
					list();		
				}
				else{
					popupMessage("No such items.");
				}	
			}
		}

		else if(currMenu == ADD_ITEM_PAGE){
			addItemPage();
		}

		else if(currMenu == ITEM_PAGE) {
			if(addItem) {
				addItemPage(items.at(selectedMenuItem));
			}
			else {
				tmpSelectedMenuItem = selectedMenuItem;
				itemPage(items.at(selectedMenuItem));
			}
		}
	}

	else if(keycode == 9)
	{
		currMenu=CATEGORY_LIST;
		if(hasResult("categories/parent_id="+to_string(items.at(selectedMenuItem).getId())))
		{
		    currCategoryId = items.at(selectedMenuItem).getId();
			list();
		}
		else{
			popupMessage("No further subcategories.");
		}	

	}
	print();
}

void Gui::mainMenu(){
	
    clearMenu();
	title = "Moose env.";
   	addMenuItem(Item("Item List", ITEM_LIST));
   	addMenuItem(Item("Add item",ADD_ITEM_PAGE));
   	addMenuItem(Item("Stock list",STOCK_LIST));
}
void Gui::addItemPage() {
	
	title = "Add Item";
	addItem=true;
    clearMenu();
	addMenuItem(Item("Select category",CATEGORY_LIST));

}

void Gui::addItemPage(Item item) {
	
	title = "Add Item";
	addItem=true;
    	clearMenu();
	addMenuItem(Item(item.getName(),CATEGORY_LIST));

    std::system("clear");	
	string res="";
	while(!isNumber(res))
		res = popupInput("Quantity:");
    int quantity = stoi(res);
	res="";
	while(!isNumber(res))
		res = popupInput("Coins:");
    int coins = stoi(res);
	string input;
	while(input != "y" && input != "n")
		input = popupInput("Add: "+to_string(quantity) +"x " +item.getName() + " for " + to_string(coins) + " coins? (y/n)");

	addItem=false;
	if(input == "y")
		itemPage(item);

	else mainMenu();
}

bool Gui::isNumber(string s) {
	
	if(s.size() == 0) return false;
	for(int i=0; i < s.size(); i++)
		if((int)s[i] < 48 || (int)s[i] > 57)
			return false;
	return true;
}

string Gui::centerText(string text, int width) {

	string out;
	int l_space = width/2 - text.size()/2;
	int r_space = width - text.size() - l_space;
	
	for(int i=0; i< l_space; i++)
		out += " ";
	out+= text;
	for(int i=0; i< r_space; i++)
		out += " ";
	return out;
}


void Gui::popupMessage(string text) {

    std::system("clear");	   
	
    cout << "\n\n\t+------------------------------------------------+" << endl;
	cout << "\t|"+centerText(text,48) +"|" << endl;
    cout << "\t+------------------------------------------------+" << endl;
	cout << "\t|"+centerText("Press Enter to continue.",48) +"|" << endl;
    cout << "\t+------------------------------------------------+" << endl;
	getchar();
    std::system("clear");	   
	list();
}

string Gui::popupInput(string text) {

    std::system("clear");	   
	cout << "\n\n\t+------------------------------------------------+" << endl;
	cout << "\t|"+centerText(text,48) +"|" << endl;
    cout << "\t+------------------------------------------------+" << endl;
	cout << "\tInput: ";
	string input;
	cin >> input;
    std::system("clear");	   
	list();
	return input;
}

void Gui::itemPage(Item item){
	
    clearMenu();
	title = "Item nr." + to_string(item.getId()) + " - " + item.getName();
	addMenuItem(Item(item.getDescription(),"asd"));

}

void Gui::list(string list_type){
	currMenu = list_type;
    list();
}

void Gui::list(){

	if(currMenu == ITEM_LIST){
		title = "Items";

		auto response = cpr::Get(cpr::Url{"http://localhost:8080/items"});
		auto json = nlohmann::json::parse(response.text);	

        clearMenu();
		for (auto& item : json) {
		    nlohmann::json object = item["object"];
		    items.push_back(Item(object["name"],(int)item["id"],object["description"]));
		}
	}	

	if(currMenu == CATEGORY_LIST){
		title = "Categories";
		if(addItem) title += " - SELECT ITEM CATEGORY";
        
		auto response = cpr::Get(cpr::Url{"http://localhost:8080/categories/parent_id="+to_string(currCategoryId)});
		auto json = nlohmann::json::parse(response.text);	

	    clearMenu();
		for (auto& item : json) {
			items.push_back(Item(item["name"],(int)item["id"],item["description"],currCategoryId));
		}
	}	

	if(currMenu == OBJ_BY_CAT_LIST){
		title = "Select Object type";
        
		auto response = cpr::Get(cpr::Url{"http://localhost:8080/objects/cat="+to_string(currCategoryId)});
		auto json = nlohmann::json::parse(response.text);	

	    clearMenu();
		for (auto& item : json) {
			items.push_back(Item(item["name"],(int)item["id"],item["description"]));
		}
	}
	else if(currMenu == STOCK_LIST){
		title = "Stocks";
        
		clearMenu();
	   	addMenuItem(Item("Stock",NULL));
	   	addMenuItem(Item("Stock",NULL));
	   	addMenuItem(Item("Stock",NULL));
	   	addMenuItem(Item("Stock",NULL));
	}
}

bool Gui::hasResult(string query) {	
	auto response = cpr::Get(cpr::Url{"http://localhost:8080/"+query});
	auto json = nlohmann::json::parse(response.text);	
	return json.size() > 0;
}

void Gui::print() {
	
    updateScrollPos();

    std::system("clear");	   
	
    cout << "+----------------------------------------------------------------+" << endl;
	cout << "|"+centerText(currMenu + " - " + title,64) + "|" << endl;
    cout << "+----------------------------------------------------------------+" << endl;

    for(int i=scrollPos; i<MENU_ITEMS + scrollPos; i++)
	{
        if(i<items.size()){
			if(i==selectedMenuItem)
				cout << "\033[30;47m"+to_string(i) + ": " + items.at(i).getName() +"\033[0m" << endl;
			else
				cout << to_string(i) + ": " + items.at(i).getName() << endl;
		}
		else cout << endl;
	}
	cout << "-- Use arrow keys to move cursor -- TAB to select categories --" << endl;
	
}
