# Take a string and return the first recurring characeter in it 

# EXAMPLE:
# "ABCA" ->  "A"
# "BCABA" -> "B"
# "ABC" -> null


# DUMP SOLUTION O(n^2)

def getRecurring(given_string):
    for c in given_string:
        for k in given_string:
            if(c==k):
                return c
            
    return null


# GOOD SOLUTION O(n)
# We use a structure, such as an hash table or a dictionary, which count the occurences of letters 

def getRecurring(given_string):
    occurences = {} # hash table or dictionary
    for char in given_string:
        if char in occurences:
            return char
        occurences[char] = 1
    return null   



# Asked by Google