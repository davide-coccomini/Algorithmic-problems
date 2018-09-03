/* Given an array, I want to print all the subset in any order you wish

 EXAMPLE:

 [1,2] -> {}, {1,2}, {1}, {2}
 Output:
 > _
 > 1,
 > 2,
 > 1,2,

*/
#include <iostream>
using namespace std;


void explore(int* given_array,int* subset, int i){
    if(i == sizeof(given_array)){
        cout << subset << endl;
    }else{
        // First exploration, without a number
        subset[i] = 0;
        explore(given_array, subset, i+1);
        //  Second exploration, with the sequel number
        subset[i] = given_array[i];
        explore(given_array, subset, i+1);
    }
}


void printSubset(int* given_array){
    int* subset = new int[sizeof(given_array)];
    explore(given_array, subset, 0);
}

int main(){
    
}
// Asked by Facebook