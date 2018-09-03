
/* Write a function which, given an array of indexes, returns if there are cycles.
 The number in the array are indexes of the array itself.

 EXAMPLE:
 [1,2,1,3,4,8] --> True
    ^ ^
 In this example there's a cycle!

 [1,2,3,4,5,6] --> False
*/

bool findCycle(int* given_array){
    int p = 0;
    int q = 0;
    int l = sizeof(given_array);

    while(true){
        p = given_array[p];

        if(p==q)
            return true;
            
        if(p<0 || p>=l) // Out of bounds
            return false;

        p = given_array[p];

        if(p==q)
            return true;
        
        q = given_array[q];

        if(p==q)
            return true;
    }
 return false;
}

int main(){
    
}

                 
                     

// Asked by Google