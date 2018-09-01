# Write a function that given 3 ordered arrays, find the intersection 

# EXAMPLE
# Arr1 = [2,6,9,11,13,17]
# Arr2 = [3,6,7,10,13,17]
# Arr3 = [4,5,6,9,11,13]

# Output: [6,13]

def arrayIntersection(arr1, arr2, arr3):
    
    result = list()
    dictionary = {}


    for number in arr1:
        dictionary[number] = 1
    for number in arr2:
        if number in dictionary.keys():
            dictionary[number] += 1
    for number in arr3:
        if number in dictionary.keys():
            dictionary[number] += 1
            if(dictionary[number] >= 3):
                result.append(number)
              

    return result

# Test call
arr1 = [2,8,9,10,15]
arr2 = [2,9,11,17,15]
arr3 = [2,3,9,17,15]
result = arrayIntersection(arr1,arr2,arr3)
print(str(result))
              
    


    
        