/* Take a string and return the first recurring characeter in it

EXAMPLE:
"ABCA" ->  "A"
"BCABA" -> "B"
"ABC" -> null

*/

// DUMP SOLUTION O(n^2)
#include <string>
#include <map>
using namespace std;

string getRecurring(string givenString) {
	for(int i = 0; i < sizeof(givenString); i++){
		for(int j = 0; j < sizeof(givenString); j++){
			if(givenString[i] == givenString[j]){
				return to_string(givenString[i]);
			}
		}
	}
	return "";
}

// GOOD SOLUTION O(n)
// We use a structure, such as an hash table or a dictionary, which count the occurences of letters

string getRecurringHash(string givenString){
    map<string, int> occurences; // hash table or dictionary
	for(int i = 0; i < sizeof(givenString); i++){
		if(occurences[to_string(givenString[i])] > 0){
			return to_string(givenString[i]);
		}
		occurences[to_string(i)] = 1;
	}
	return "";
}

int main(){
    
}
// Asked by Google
