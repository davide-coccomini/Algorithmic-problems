# Given an array which represents an histogram, return the area of the largest rectangle. Consider each column wide 1 and tall the number in the array.

# EXAMPLE
# [1,2,2]

#
#   __ __
#__|  |  |
#  |  |  |

# In this case the greatest rectangle is the one made of the two colums (2 and 3) so the area is (1+1)*2 = 4


class Stack:
     def __init__(self):
         self.items = []

     def isEmpty(self):
         return self.items == []

     def push(self, item):
         self.items.append(item)

     def pop(self):
         return self.items.pop()

     def peek(self):
         return self.items[len(self.items)-1]

     def size(self):
         return len(self.items)


def histogramMaxArea(given_array):
    stack = Stack()
    max = 0
    i = 0

    while(i<len(given_array)):
        if(stack.isEmpty() or stack.peek() <= given_array[i]):
            stack.push(i)
            i += 1
        else:
            temp_max = stack.pop()
            area = 0
            if(stack.isEmpty):
                area = given_array[temp_max] * (i-1)
            else:
                area = given_array[temp_max] * (i-1-stack.peek())
            if(area>max): 
                max = area
    while(stack.peek()):
        temp_max = stack.pop()
        if(stack.isEmpty):
            area = given_array[temp_max] * (i-1)
        else:
            area = given_array[temp_max] * (i-1-stack.peek())
        if(area>max): 
                max = area           
        
    return max

arr = [1,2,2,3]
maxArea = histogramMaxArea(arr)
print(str(maxArea))
