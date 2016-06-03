#ifndef user_h
#define user_h

#include <iostream>
using namespace std;
class User{

	public:
		User();
		User(string token);
		User(int id, string username, string email, string name, string surname, int balance, int type, int group_id, string token);
		int getId();
		string getUsername();
		string getEmail();
		string getName();
		string getSurname();
		int getBalance();
		int getType();
		int getGroupId();
		string getToken();
		bool isValid();
	
	private:
		string username, email, name, surname, token;
		int id, balance, type, group_id;
};

#endif
