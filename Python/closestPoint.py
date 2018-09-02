# Given an array of points (x,y), find the K closest points to the origin (0,0)
import math

def calculateDistance(x, y):
    return math.sqrt(x**2 + y**2)



def closestPoint(given_array, k):
    dictionary = {}
    for point in given_array:
        distance = calculateDistance(point[0],point[1])
        label = str(point[0])+","+str(point[1])
        dictionary[label] = distance

    dictionary = sorted(dictionary)

    for i in range(0,k):
        print(dictionary[i])
        


        
# Asked by Amazon