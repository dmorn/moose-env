#include<iostream>
#include<restclient-cpp/restclient.h>

using namespace std;
int main() {
    
    RestClient::Response r = RestClient::get("http://localhost:8080/users");
    cout << r.body << endl;
    
    return 0;
}

