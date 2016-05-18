#include <iostream>

#include <cpr/cpr.h>
#include <json.hpp>


int main(int argc, char** argv) {
    auto response = cpr::Get(cpr::Url{"http://localhost:8080/objects"});
    auto json = nlohmann::json::parse(response.text);
    std::cout << json.dump(4) << std::endl;
}

