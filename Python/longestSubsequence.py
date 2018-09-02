# Given an array, find the length of the longest subsequence 

# EXAMPLE
# [2,1,5,9,4,3] --> 4 (because [1,2,3,4])

def longestSubsequence(given_array):
    dictionary = {}
    length = 0
    for i in given_array:
        dictionary[i] = True
    
    for i in given_array:
        j = i
        temp_length = 0
        if i-1 not in dictionary.keys(): # This condition reduce complexity
            while dictionary[j] == True:
                temp_length+=1
                j+=1
                if(temp_length>length):
                    length = temp_length

                if j not in dictionary.keys():
                    break

    return length


# Test call

arr = [1,2,3,5,4,8,9]
result = longestSubsequence(arr)
print(str(result))
