#include "gui.h"

Gui::Gui(){
	selectedMenuItem=0;
	tmpSelectedMenuItem=-1;
    scrollPos=0;
	title="Test";
	currMenu = LOGIN;
    currCategoryId=0;
	addItem=false;
	addItemToStock=false;
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

	
	getJson("user");


	if(currMenu != LOGIN){

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
				auto cat = getJson("categories/id="+to_string(items.at(selectedMenuItem).getParentId()));
		        currCategoryId = cat["parent_id"];		
				list();
			}
			else if(currMenu == OBJ_BY_CAT_LIST)
			{
				clearMenu();
				list(CATEGORY_LIST);
			}
			else{
				selectedMenuItem=0;
				mainMenu();
			}
		}
		else if(keycode == 10 && items.at(selectedMenuItem).getFunction() != "nil")
		{
		   	currMenu = items.at(selectedMenuItem).getFunction();

			if(currMenu == ITEM_LIST || currMenu == PROFILE) 
				list();	

		    if(currMenu == CATEGORY_LIST) {

				if(hasResult("categories/parent_id="+to_string(items.at(selectedMenuItem).getId())))
				{
					currCategoryId = items.at(selectedMenuItem).getId();
					list();
				}
				else{
					bool ok = popupYesNo("No subcategories, search in this category? (y/n)");
					if(ok){
						currCategoryId = items.at(selectedMenuItem).getId();
						if(addItem)
							list(OBJ_BY_CAT_LIST);
						else
							list(ITEM_LIST);
					}
				}	

				if(items.at(selectedMenuItem).getId() == 0) {

			    	currCategoryId = items.at(selectedMenuItem).getId();
					list();	
				}

				/*
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
				else {
					if(hasResult("items/start_cat_id="+to_string(items.at(selectedMenuItem).getId()))) {
						currCategoryId = items.at(selectedMenuItem).getId();
						currMenu = ITEM_LIST;
						list();	
					}
					else{
						popupMessage("No such items.");
					}	
				}*/
			}

			else if(currMenu == ADD_ITEM_PAGE || currMenu == ADD_STOCK_PAGE){
				addItem=true;
				if(currMenu == ADD_STOCK_PAGE)
					addItemToStock=true;
				
				list(CATEGORY_LIST);
			}
			else if(currMenu == BUY_ITEM_PAGE){
				int quantity = popupNumber("Quantity: ");
				bool ok = popupYesNo("Order " + to_string(quantity) + "x " +selectedItem.getName()+ " for " + 
									to_string(selectedItem.getCoins() * quantity) + " coins?");
			
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
			if(currMenu == ITEM_LIST) {
				currCategoryId=0;
				currMenu=CATEGORY_LIST;
				list();
			}
			else if (currMenu == CATEGORY_LIST){
				if(addItem)
				{	
					if(currCategoryId==0)
						popupMessage("Please specify the category.");

					else
						list(OBJ_BY_CAT_LIST);
				}
				else
					list(ITEM_LIST);				
			}
		}

		print();
	}

	else {
		
    	std::system("clear");
		string username, password;
		cout << "please log in:\nusername: ";
		cin >> username;
		cout << "password: ";

		termios oldt;
		tcgetattr(STDIN_FILENO, &oldt);
		termios newt = oldt;
		newt.c_lflag &= ~ECHO;
		tcsetattr(STDIN_FILENO, TCSANOW, &newt);
		getline(cin, password);
		cin >> password;
		tcsetattr(STDIN_FILENO, TCSANOW, &oldt);

		json userData = { {"username", username}, {"password", password} };
		
		json res = postJsonNoToken("login",userData);

		if(!res.is_null()){
			string t = res["token"];
			user = User(t);
			json uJ = getJson("user");
			user = User((int)uJ["id"], uJ["username"], uJ["email"], uJ["name"], uJ["surname"], (int)uJ["balance"], (int)uJ["type"], (int)uJ["group_id"], t);
			mainMenu();
		}
		else
			cout << "\nlogin incorrect.\n";

		update(0);
	}
}

void Gui::mainMenu(){
	currMenu = MAIN_MENU;
    clearMenu();
	addItem=false;
	addItemToStock=false;
	title = "Welcome " + user.getName() + " to Moose env.";
   	addMenuItem(Item("Item List", ITEM_LIST));
   	addMenuItem(Item("Add item to stock",ADD_STOCK_PAGE));
   	addMenuItem(Item("Add item to wishlist",ADD_ITEM_PAGE));
   	addMenuItem(Item("Stock list",STOCK_LIST));
   	addMenuItem(Item("View Profile",PROFILE));
}

void Gui::addItemPage(Item item) {
	
	title = "Add Item";
    clearMenu();
	addMenuItem(Item(item.getName(),CATEGORY_LIST));

    std::system("clear");	
    int quantity = popupNumber("Quantity:");
    int coins = popupNumber("Coins:");
	bool add = popupYesNo("Add: "+to_string(quantity) +"x " +item.getName() + " for " + to_string(coins) + " coins? (y/n)");
	
	int stock_id=0;
	Item newItem = Item(item.getName(),-1,item.getDescription(), coins, quantity, stock_id, item.getId());

	json newItemJson = {
	  {"status", (addItemToStock)?1:3},
	  {"coins", coins},
	  {"quantity", quantity},
	  {"stock_id", stock_id},
	  {"object_id", item.getId()},
	};

	
	addItem=false;
	addItemToStock=false;
	int newId = postJson("item",newItemJson).back()["id"];
	newItem.setId(newId);
	if(add)
		itemPage(newItem);

	else mainMenu();
	currCategoryId=0;
}

bool Gui::popupYesNo(string text) {
	string res = "";
	while(res != "y" && res != "n")
		res = popupInput(text);
	return res=="y";
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
}

int Gui::popupNumber(string text) {
	string res="";
	while(!isNumber(res))
		res = popupInput(text);
    return stoi(res);
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
	selectedItem = item;
	addMenuItem(Item(item.getDescription(),"asd"));
	addMenuItem(Item("Coins:\t" + to_string(item.getCoins()),"asd"));
	addMenuItem(Item("Quantity:\t" + to_string(item.getQuantity()),"asd"));
	addMenuItem(Item("Stock:\t" + to_string(item.getStockId()),"asd"));
	addMenuItem(Item("Buy Item",BUY_ITEM_PAGE));
}

void Gui::list(string list_type){
	currMenu = list_type;
    list();
}

void Gui::list(){
		
	if(currMenu == ITEM_LIST){
		title = "Items";


		auto res = getJson("items/start_cat_id="+to_string(currCategoryId));
		clearMenu();
		if(res.size() > 0)
		{
			for (auto& item : res) {
				json object = item["object"];
				items.push_back(Item(object["name"],(int)item["id"],object["description"],(int)item["coins"],(int)item["quantity"],(int)item["stock_id"],(int)item["object_id"]));
			}
			
			title = to_string(items.size()) + " Items found";
		}
		else if(currCategoryId != 0)
		{	
			getchar();
			popupMessage("No items found, removing category filter");
			currCategoryId=0;
			list();
		}	
		else {
			getchar();
			popupMessage("No items found.");
   			clearMenu();
			mainMenu();
		}
	}	

	else if(currMenu == CATEGORY_LIST){
		title = "Categories";
		if(addItem) title += " - SELECT ITEM CATEGORY";

	    clearMenu();
		for (auto& item : getJson("categories/parent_id="+to_string(currCategoryId))) {
			items.push_back(Item(item["name"],(int)item["id"],item["description"],currCategoryId));
		}
	}	

	else if(currMenu == OBJ_BY_CAT_LIST){
		title = "Select Object type";   

	    clearMenu();
		for (auto& item : getJson("objects/cat="+to_string(currCategoryId))) {
			items.push_back(Item(item["name"],(int)item["id"],item["description"]));
		}
	}
	else if(currMenu == STOCK_LIST){
		title = "Stocks";
        
		clearMenu();
	   	addMenuItem(Item("Stock"));
	   	addMenuItem(Item("Stock"));
	   	addMenuItem(Item("Stock"));
	   	addMenuItem(Item("Stock"));
	}

	else if(currMenu == PROFILE){


		title = "Profile";
        
		clearMenu();
		addMenuItem(Item("Id:\t\t" + to_string(user.getId())));
		addMenuItem(Item("Username:\t" + user.getUsername()));
		addMenuItem(Item("Email:\t" + user.getEmail()));
		addMenuItem(Item("Name:\t" + user.getName()));
		addMenuItem(Item("Surname:\t" + user.getSurname()));
		addMenuItem(Item("Credits:\t" + to_string(user.getBalance())));
		addMenuItem(Item("Type:\t" + to_string(user.getType())));
		addMenuItem(Item("Group:\t" + to_string(user.getGroupId())));
	//User((int)uJ["id"], uJ["username"], uJ["email"], uJ["name"], uJ["surname"], (int)uJ["balance"], (int)uJ["type"], (int)uJ["group_id"], t);

	}

}

/*
curl -H "Content-Type: application/json" -X POST -d '{"username":"matthias", "password": "test"}' http://localhost:8080/login
*/

json Gui::getJson(string content) {

	auto response = cpr::Get(cpr::Url{"http://localhost:8080/"+content},
	cpr::Header{{"Authorization", "Bearer " +user.getToken()}});
	if(response.text == ("unauthorized")){
		currMenu = LOGIN;
		return NULL;
	}
	else
		return json::parse(response.text);
}

json Gui::postJsonNoToken(string content, json data) {

	auto r = cpr::Post(cpr::Url{"http://localhost:8080/"+content},
	cpr::Body{data.dump()},
	cpr::Header{{"Content-Type", "application/json"}});
	if(r.status_code == 404) {
		json empty;	
		return empty;
	}
	return json::parse(r.text);
}

json Gui::postJson(string content, json data) {

	auto r = cpr::Post(cpr::Url{"http://localhost:8080/"+content},
	cpr::Body{data.dump()},
	cpr::Header{{"Authorization", "Bearer " +user.getToken()},
				{"Content-Type", "application/json"}});
	return json::parse(r.text);

}

bool Gui::hasResult(string query) {	
	return getJson(query).size() > 0;
}

bool Gui::isNumber(string s) {
	
	if(s.size() == 0) return false;
	for(int i=0; i < s.size(); i++)
		if((int)s[i] < 48 || (int)s[i] > 57)
			return false;
	return true;
}

string Gui::centerText(string t, int w) {
	string o;
	int l = w/2 - t.size()/2;
	int r = w - l - t.size();
	for(; l>0; l--) o+=" ";
	o+=t;
	for(; r>0; r--) o+=" ";
	return o;
}

void Gui::print() {
	
    updateScrollPos();

    std::system("clear");	   
	
    cout << "+----------------------------------------------------------------+" << endl;
	cout << "|"+centerText(currMenu + " - " + title + " - " + to_string(currCategoryId),64) + "|" << endl;
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
