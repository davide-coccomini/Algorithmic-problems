
# Write a function which, given an array of indexes, returns if there are cycles.
# The number in the array are indexes of the array itself.

# EXAMPLE:
# [1,2,1,3,4,8] --> True
#    ^ ^
# In this example there's a cycle!

# [1,2,3,4,5,6] --> False

def findCycle(given_array):
    p = 0
    q = 0
    l = len(given_array)

    while(True):

        p = given_array[p] 
        if(p==q):
            return True
            
        if(p<0 or p>=l): # Out of bounds
            return False

        p = given_array[p]

        if(p==q):
            return True
        
        q = given_array[q]

        if(p==q):
            return True

    return False
                      
                     

# Asked by Google