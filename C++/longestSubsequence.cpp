/* Given an array, find the length of the longest subsequence

EXAMPLE
[2,1,5,9,4,3] --> 4 (because [1,2,3,4])

*/

#include <map>
using namespace std;

int longestSubsequence(int* givenArray)  {
    map<int, bool> dictionary;
	int length = 0;
	for(int i = 0; i < sizeof(givenArray); i++){
		dictionary[givenArray[i]] = true;
	}

	for(int i = 0; i < sizeof(givenArray); i++) {
		int j = givenArray[i];
		int tempLength = 0;
		if(!dictionary[givenArray[i]-1]) {
			while(dictionary[j]){
				tempLength++;
				j++;
				if(tempLength > length){
					int length = tempLength;
				}
				if(!dictionary[j]){
					break;
				}
			}
		}
	}
	return length;
}

// Test call

int main(){
	int* arr = new int[1,2,3,5,4,8,9];
	int result = longestSubsequence(arr);
}

