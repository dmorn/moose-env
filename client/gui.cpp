#include "gui.h"

Gui::Gui(){
	selectedMenuItem=0;
	tmpSelectedMenuItem=-1;
    scrollPos=0;
	title="Test";
	footer="moose.env v.1";
	currMenu = LOGIN;
    currCategoryId=0;
	addItem=false;
	addItemToStock=false;
	currentItem=NULL;
	selectedStock=NULL;
}

void Gui::clearMenu()
{
	for(int i=0; i<elements.size(); i++)
	{
		delete (elements.at(i));
	}
	elements.clear();
	selectedMenuItem=0;
}


void Gui::updateScrollPos(){
	if(selectedMenuItem < scrollPos)
		scrollPos = selectedMenuItem;
	else if(selectedMenuItem >= scrollPos + MENU_ELEMENTS)
		scrollPos = selectedMenuItem - MENU_ELEMENTS+1;
}

void Gui::update(int keycode) {

	//65 U , 68 L , 67 R, 66 D, 127 BACK, 10 ENTER, 32 SPACE, 9 TAB

	
	getJson("user");


	if(currMenu != LOGIN){

		if(keycode == 65)
			selectedMenuItem = (selectedMenuItem > 0) ? --selectedMenuItem : elements.size() -1;

		else if(keycode == 66)
			selectedMenuItem = (selectedMenuItem < elements.size()-1) ? ++selectedMenuItem : 0; 

		else if(keycode == 67){
		    selectedMenuItem += MENU_ELEMENTS;
		    if(selectedMenuItem > elements.size() - 1)
		        selectedMenuItem=elements.size()-1;  
		}
		else if(keycode == 68){
		    selectedMenuItem -= MENU_ELEMENTS;
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
				if(currentItem->getStatus() == 3)
					list(WISH_LIST);
				else
					list(ITEM_LIST);
			}
			else if(currMenu == CATEGORY_LIST && currCategoryId != 0)
			{
				Category* category = (Category*)elements.at(selectedMenuItem);
				if(category==NULL)
					popupMessage("not a category.");
				else {
					auto cat = getJson("categories/id="+to_string(category->getParentId()));
				    currCategoryId = cat["parent_id"];		
					list();
				}
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
		else if(keycode == 10 && elements.at(selectedMenuItem)->getFunction() != NO_FUNCTION)
		{
			if(elements.at(selectedMenuItem)->getFunction() == TEXT_POPUP)	//temporary menu change, no need to change currMenu
			{
				popupMessage(elements.at(selectedMenuItem)->getText());
			}
			else {

			   	currMenu = elements.at(selectedMenuItem)->getFunction();

				if(currMenu == ITEM_LIST || currMenu == STOCK_LIST || currMenu == WISH_LIST || currMenu == PROFILE) 
					list();	

				else if(currMenu == CATEGORY_LIST) {

					Category* category = (Category*)elements.at(selectedMenuItem);
					if(category==NULL)
						popupMessage("not a category.");
					else 
					{
						if(hasResult("categories/parent_id="+to_string(category->getId())))
						{
							currCategoryId = category->getId();
							list();
						}
						else
						{
							if(popupYesNo("No subcategories, search in this category? (y/n)"))
							{
								currCategoryId = category->getId();
								if(addItem)
									list(OBJ_BY_CAT_LIST);
								else
									list(ITEM_LIST);
							}
						}	

						if(category->getId() == 0) {
							currCategoryId = category->getId();
							list();	
						}

						/*
						else if(addItem) {
							if(hasResult("objects/cat="+to_string(elements.at(selectedMenuItem).getId())))
							{
								currCategoryId = elements.at(selectedMenuItem).getId();
								currMenu = OBJ_BY_CAT_LIST;
								list();		
							}
							else{
								popupMessage("No such objects.");
							}	
						}
						else {
							if(hasResult("items/start_cat_id="+to_string(elements.at(selectedMenuItem).getId()))) {
								currCategoryId = elements.at(selectedMenuItem).getId();
								currMenu = ITEM_LIST;
								list();	
							}
							else{
								popupMessage("No such items.");
							}	
						}*/
					}
				}

				else if(currMenu == ADD_ITEM_PAGE || currMenu == ADD_STOCK_PAGE){
					addItem=true;
					if(currMenu == ADD_STOCK_PAGE)
						addItemToStock=true;
				
					list(STOCK_LIST);
				}
				else if(currMenu == BUY_ITEM_PAGE){
					
					if(currentItem==NULL)
						popupMessage("No item selected.");
					else {
						int q = popupNumber("Quantity: ");
						while (currentItem->getQuantity() < q) {
							q = popupNumber("Not enough aviable. Quantity: ");
						}
						bool ok = popupYesNo("Buy " + to_string(q) + "x " +currentItem->getName()+ " for " + 
											to_string(currentItem->getCoins() * q) + " coins?");
						if(ok){
							json res = postJson("purchase/"+to_string(currentItem->getId()) +"/"+to_string(q));
							if(!res.empty()){
								string d = res["data"];
								popupMessage("Your reciept:\t"+d);
							}

							mainMenu();
						}
					}
				}				
				else if(currMenu == ORDER_ITEM_PAGE){
					
					if(currentItem==NULL)
						popupMessage("No item selected.");
					else {
						bool ok = popupYesNo("Order " + to_string(currentItem->getQuantity()) + "x " +currentItem->getName()+ " for " + 
											to_string(currentItem->getCoins() * currentItem->getQuantity()) + " coins?");
						if(ok) mainMenu();
					}
				}
				else if(currMenu == ADD_ITEM_SELECTED) {

					addItemPage(selectedMenuItem);
				}

				else if(currMenu == ITEM_PAGE) {
					tmpSelectedMenuItem = selectedMenuItem;
					itemPage(selectedMenuItem);
				}
				else if(currMenu == ITEM_BY_STOCK) {
					selectedStock = (Stock*) elements.at(selectedMenuItem);
					elements.erase(elements.begin() + selectedMenuItem);
					if(addItem){
						list(CATEGORY_LIST);
					}
					else {
						list(ITEM_LIST);
					}
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
	if(currentItem!=NULL);
		delete currentItem;
	currentItem=NULL;
	if(selectedStock!=NULL);
		delete selectedStock;
	selectedStock=NULL;
	title = "Welcome " + user.getName() + " to Moose env.";
	footer="moose.env v.1";
   	elements.push_back(new MenuItem("Item List", ITEM_LIST));
   	elements.push_back(new MenuItem("Wishlist", WISH_LIST));
   	elements.push_back(new MenuItem("Add item to stock",ADD_STOCK_PAGE));
   	elements.push_back(new MenuItem("Add item to wishlist",ADD_ITEM_PAGE));
   	elements.push_back(new MenuItem("Stock list",STOCK_LIST));
   	elements.push_back(new MenuItem("View Profile",PROFILE));
}

void Gui::addItemPage(int object_no) {
	
	Object* obj = (Object*)elements.at(object_no);
	elements.erase(elements.begin() + object_no);
	title = "Add Item";
    clearMenu();
	elements.push_back(new MenuItem(obj->getText(),CATEGORY_LIST));

    std::system("clear");	
    int quantity = popupNumber("Quantity:");
    int coins = popupNumber("Coins:");
	bool add = popupYesNo("Add: "+to_string(quantity) +"x " +obj->getName() + " for " + to_string(coins) + " coins? (y/n)");
	
	int stock_id=0;
	if(selectedStock != NULL)
		stock_id = selectedStock->getId();
	int status = (addItemToStock)?1:3;
	json newItemJson = {
	  {"status", status},
	  {"coins", coins},
	  {"quantity", quantity},
	  {"stock_id", stock_id},
	  {"object_id", obj->getId()},
	};

	
	addItem=false;
	addItemToStock=false;
	if(add) {
		int newId = postJson("item",newItemJson).back()["id"];
		popupMessage("Item added with id " + to_string(newId));
	}

	mainMenu();
	currCategoryId=0;
}

void Gui::itemPage(int item_no){
	currentItem = (Item*)elements.at(item_no);
	elements.erase(elements.begin() + item_no);
    clearMenu();
	title = "Item nr." + to_string(currentItem->getId()) + " - " + currentItem->getName();

	elements.push_back(new MenuItem(currentItem->getDescription(),TEXT_POPUP));
	elements.push_back(new MenuItem("Coins:\t" + to_string(currentItem->getCoins())));
	elements.push_back(new MenuItem("Quantity:\t" + to_string(currentItem->getQuantity())));
	elements.push_back(new MenuItem("Stock:\t" + currentItem->getStock()));
	if(currentItem->getQuantity() > 0){
		switch(currentItem->getStatus()) {
			case 1: elements.push_back(new MenuItem("Buy Item",BUY_ITEM_PAGE)); break;
			case 2: elements.push_back(new MenuItem("Item ordered but not yet in stock.")); break;
			case 3: elements.push_back(new MenuItem("Order and remove from wishlist",ORDER_ITEM_PAGE)); break;
		}
	}
	else
		elements.push_back(new MenuItem("Not available"));
}

void Gui::list(string list_type){
	currMenu = list_type;
    list();
}

void Gui::list(){

	footer="moose.env v.1";

	if(currMenu == WISH_LIST){
		title = "Wishlist";
		footer = "Use arrow keys to move cursor";
		auto res = getJson("items/wishlist");
		clearMenu();
		if(res.size() > 0)
		{
			for (auto& item : res) {
				json object = item["object"];
				Item* it = new Item(object["name"],(int)item["id"],object["description"],(int)item["coins"],(int)item["quantity"],item["stock"]["name"],(int)item["object_id"],(int)item["status"]);
				elements.push_back(it);
			}
			
			title = to_string(elements.size()) + " Items found in wishlist";
		}
		else {
			getchar();
			popupMessage("No items in wishlist.");
   			clearMenu();
			mainMenu();
		}
	}
		
	if(currMenu == ITEM_LIST){
		title = "Items";
		footer = "Use arrow keys to move cursor";
		string query = "items/start_cat_id="+to_string(currCategoryId);
		if(selectedStock != NULL)
			query = "items/1/"+to_string(selectedStock->getId())+"/"+to_string(currCategoryId);
		auto res = getJson(query);
		clearMenu();
		if(res.size() > 0)
		{
			for (auto& item : res) {
				json object = item["object"];
				Item* it = new Item(object["name"],(int)item["id"],object["description"],(int)item["coins"],(int)item["quantity"],item["stock"]["name"],(int)item["object_id"],(int)item["status"]);
				elements.push_back(it);
			}
			
			title = to_string(elements.size()) + " Items found";
			if(selectedStock != NULL)
				title += " @" + selectedStock->getText();
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
		footer = "Press TAB to select current category";
		if(addItem) title += " - SELECT ITEM CATEGORY";

	    clearMenu();
		for (auto& item : getJson("categories/parent_id="+to_string(currCategoryId))) {
			Category * c = new Category(item["name"],(int)item["id"],item["description"],currCategoryId);
			elements.push_back(c);
		}
	}	

	else if(currMenu == OBJ_BY_CAT_LIST){
		title = "Select Object type";   

	    clearMenu();
		for (auto& item : getJson("objects/start_cat_id="+to_string(currCategoryId))) {
			Object * o= new Object(item["name"],(int)item["id"],item["description"]);
			elements.push_back(o);
		}
	}
	else if(currMenu == STOCK_LIST){
		title = "Stocks";
		clearMenu();
		for (auto& item : getJson("stocks")) {
			Stock * s = new Stock(item["name"],(int)item["id"],item["location"]);
			elements.push_back(s);
		}
	}

	else if(currMenu == PROFILE){


		title = "Profile";
        
		clearMenu();
		elements.push_back(new MenuItem("Id:\t\t" + to_string(user.getId())));
		elements.push_back(new MenuItem("Username:\t" + user.getUsername()));
		elements.push_back(new MenuItem("Email:\t" + user.getEmail()));
		elements.push_back(new MenuItem("Name:\t" + user.getName()));
		elements.push_back(new MenuItem("Surname:\t" + user.getSurname()));
		elements.push_back(new MenuItem("Credits:\t" + to_string(user.getBalance())));
		elements.push_back(new MenuItem("Type:\t" + to_string(user.getType())));
		elements.push_back(new MenuItem("Group:\t" + to_string(user.getGroupId())));

	//REMOVE BEFORE PUBLISHMENT! ------------------------------------------------------------------------
		elements.push_back(new MenuItem("Token:\t" + user.getToken()));
	}

}

bool Gui::popupYesNo(string text) {
	string res = "";
	while(res != "y" && res != "n")
		res = popupInput(text);
	return res=="y";
}

void Gui::popupMessage(string text) {

    std::system("clear");	   
    cout << endl << endl;

	cout << centerText("+------------------------------------------------+",DISPLAY_WIDTH) << endl;
	for (unsigned i = 0; i < text.length(); i += POPUP_WIDTH - 4)
		cout << centerText("| "+centerText(text.substr(i, POPUP_WIDTH - 4),POPUP_WIDTH -2) +" |" ,DISPLAY_WIDTH) << endl;
	cout << centerText("+------------------------------------------------+",DISPLAY_WIDTH) << endl;
	cout << centerText("|"+centerText("Press Enter to continue.",POPUP_WIDTH) +"|" ,DISPLAY_WIDTH) << endl;
	cout << centerText("+------------------------------------------------+",DISPLAY_WIDTH) << endl;
	cin.ignore(); 	
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
	cout << centerText("|"+centerText(text,POPUP_WIDTH) +"|" ,DISPLAY_WIDTH) << endl;	
    cout << "\t+------------------------------------------------+" << endl;
	cout << "\tInput: ";
	string input;
	cin >> input;
    std::system("clear");	   
	list();
	return input;
}


/*
curl -H "Content-Type: application/json" -X POST -d '{"username":"matthias", "password": "test"}' http://localhost:8080/login
*/

json Gui::getJson(string content) {

	auto response = cpr::Get(cpr::Url{URL+content},
	cpr::Header{{"Authorization", "Bearer " +user.getToken()}});
	if(response.text == ("unauthorized")){
		currMenu = LOGIN;
		return NULL;
	}
	else
		return json::parse(response.text);
}

json Gui::postJsonNoToken(string content, json data) {

	auto r = cpr::Post(cpr::Url{URL+content},
	cpr::Body{data.dump()},
	cpr::Header{{"Content-Type", "application/json"}});
	if(r.status_code == 404) {
		json empty;	
		return empty;
	}
	return json::parse(r.text);
}

json Gui::postJson(string content) {

	auto r = cpr::Post(cpr::Url{URL+content},
	cpr::Body{},
	cpr::Header{{"Authorization", "Bearer " +user.getToken()},
				{"Content-Type", "application/json"}});
	if(r.text[0] == '{')
		return json::parse(r.text);
	
	else 
		popupMessage(r.text);
	json eJ;
	return eJ;

}

json Gui::postJson(string content, json data) {

	auto r = cpr::Post(cpr::Url{URL+content},
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
	int l = (w - t.size())/2;
	int r = w - t.size() - l;
	for(; l>0; l--) o+=" ";
	o+=t;
	for(; r>0; r--) o+=" ";
	return o;
}

string Gui::fillWithSpace(int cnt) {
	string spstr ="";
	for(int i=0; i<cnt; i++)
		spstr+=" ";
	return spstr;
}

void Gui::print() {
	
    updateScrollPos();

    std::system("clear");	   
	
    //cout << "+----------------------------------------------------------------+" << endl;
	cout << "\033[30;47m"+centerText(currMenu + " - " + title + " - " + to_string(currCategoryId),66) + "\033[0m" << endl;
    //cout << "+----------------------------------------------------------------+ " << endl;

	if(scrollPos > 0)
		cout << centerText("^",66);
	cout<<endl;
    for(int i=scrollPos; i<MENU_ELEMENTS + scrollPos; i++)
	{
        if(i<elements.size()){
			if(i==selectedMenuItem)
				cout << "\033[30;47m"+limitText(elements.at(i)->getText()) +"\033[0m" << endl;
			else
				cout << limitText(elements.at(i)->getText()) << endl;

/*

			if(i==selectedMenuItem)
				cout << "\033[30;47m"+limitText(to_string(i) + ": " + elements.at(i)->getText()) +"\033[0m" << endl;
			else
				cout << limitText(to_string(i) + ": " + elements.at(i)->getText()) << endl;
*/
		}
		else cout << endl;
	}
	if(scrollPos + MENU_ELEMENTS <= elements.size()-1)
		cout << centerText("v",66);

	cout << endl << "\033[30;47m"+centerText("-- " +footer + " --",66)+"\033[0m" << endl;
	
}

string Gui::limitText(string text) {
	if(text.size() <= DISPLAY_WIDTH)
		return text;
	
	return text.substr(0,DISPLAY_WIDTH-3)+"...";

}
