/* Write a function that given 3 ordered arrays, find the intersection 

 EXAMPLE
 Arr1 = [2,6,9,11,13,17]
 Arr2 = [3,6,7,10,13,17]
 Arr3 = [4,5,6,9,11,13]

 Output: [6,13]
*/
#include <iostream>
#include <vector>
#include <map>

using namespace std;

vector<int> arrayIntersection(int* arr1,int* arr2,int* arr3){
    
    vector<int> result;
    map<int, int> dictionary;


    for(int i=0; i<sizeof(arr1); i++)
        dictionary[arr1[i]] = 1;
    for(int i=0; i<sizeof(arr2); i++){
        if(dictionary.count(arr2[i]))
            dictionary[arr2[i]]++;
    }
    for(int i=0; i<sizeof(arr3); i++){
        if(dictionary.count(arr3[i])){
            dictionary[arr3[i]]++;
            if(dictionary[arr3[i]] >= 3)
                result.push_back(arr3[i]);
        }
    }

    return result;
}
// Test call
int main(){
    int* arr1 = new int[2,8,9,10,15];
    int* arr2 = new int[2,9,11,17,15];
    int* arr3 = new int[2,3,9,17,15];
    vector<int> result = arrayIntersection(arr1,arr2,arr3);
}   
    


    
        