#include<iostream>
#include <conio.h>
#include "Gui.h"
//#include<restclient-cpp/restclient.h>

using namespace std;

string title = "Title";


Gui gui = Gui();
int key_code;

int main() {
       
   // RestClient::Response r = RestClient::get("http://localhost:8080/users");
    //cout << r.body << endl;
    
   	gui.mainMenu();
    gui.print();
    
	while (1){
    	if ( kbhit() ){
    		int key = getch();
		    gui.update(key);
		}
	    else 
	        continue;
    }
    
    return 0;
}

