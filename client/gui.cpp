#include "gui.h"

Gui::Gui(){
	
	elementCnt=0;
	selectedElement=0;
	title="Test";
}

void Gui::addElement(string elem, string function, string param_1)
{
	if(elementCnt < MAX_ELEMENTS){
		elements[elementCnt][0]=elem;
		elements[elementCnt][1]=function;
		elements[elementCnt][2]=param_1;
		elementCnt++;
	}
}


void Gui::clearElements()
{
	elementCnt=0;
}

void Gui::update(int keycode) {
	//65 U , 68 L , 67 R, 66 D, 127 BACK, 10 ENTER
	if(keycode == 65){
		selectedElement--;
			if(selectedElement < 0)
				selectedElement = elementCnt-1;
	}	
	else if(keycode == 66){
		selectedElement++;
			if(selectedElement >= elementCnt)
				selectedElement = 0;
	}
	else if(keycode == 127)
	{
		clearElements();
		mainMenu();
		selectedElement=0;
	}
	else if(keycode == 10)
	{
		clearElements();
		string func = elements[selectedElement][1];
		if(func == MAIN_MENU) 
			mainMenu();
		else if(func == ITEM_PAGE)
			itemPage(elements[selectedElement][2]);
		else 
			list(func,stoi(elements[selectedElement][2]));
		selectedElement=0;
	}
	print();
}

void Gui::mainMenu(){
	
	title = "Moose Environment";
   	addElement("Item list",ITEM_LIST,"0");
   	addElement("Stock list",STOCK_LIST,"0");
}

void Gui::itemPage(string id){
	
	title = "Item nr."+id;
	auto response = cpr::Get(cpr::Url{"http://localhost:8080/objects/id="+id});
	auto item = nlohmann::json::parse(response.text);
	addElement(item["name"],"nil","1");
	addElement(item["description"],"nil","1");
	addElement(item["description"],"nil","1");
}

void Gui::list(string type, int page){
	
	if(type == ITEM_LIST){
		title = "Items - Page "+to_string(page);
	 	auto response = cpr::Get(cpr::Url{"http://localhost:8080/items"});
		auto items = nlohmann::json::parse(response.text);
		
		int cnt=0;
		for (auto& item : items) {
			if(cnt>=(page*(MAX_ELEMENTS-2)))
	   			addElement(to_string((int)item["id"]),ITEM_PAGE,to_string((int)item["id"]));
			if((cnt+1)>=((page+1)*(MAX_ELEMENTS-2))) {
	   			addElement("...To page " + to_string(page+1) + "-->",ITEM_LIST,to_string(page+1));
				break;
			}
			cnt++;
			
		}
		if(page > 0){
			string text = "<--To page ";
			text += to_string(page-1);
			text += "...";
   			addElement(text,ITEM_LIST,to_string(page-1));
		}
	}	
	else if(type == STOCK_LIST){
		title = "Stocks - Page "+page;
	   	addElement("Stock",STOCK_LIST,"0");
	   	addElement("Stock",STOCK_LIST,"0");
	   	addElement("Stock",STOCK_LIST,"0");
	   	addElement("Stock",STOCK_LIST,"0");
	}
}

void Gui::print() {
	
    std::system("clear");	   

    cout << "------------------------------------------------------------------" << endl;
    cout << "\t" + title << endl;
    cout << "------------------------------------------------------------------" << endl;
    int i=0;
    for(i=0; i<elementCnt; i++)
	{
		if(i==selectedElement)
		{
			cout << "|X| " + elements[i][0] << endl;
		}
		else
		{
			cout << "| | " + elements[i][0] << endl;
		}
	}
	
}
