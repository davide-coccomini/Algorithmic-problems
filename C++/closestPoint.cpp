// Given an array of points (x,y), find the K closest points to the origin (0,0)

#include <math.h>
#include <string>
#include <map>
#include <iostream>

using namespace std;
float calculateDistance(int x,int y)  {
	float p = pow(x, 2) + pow(y, 2);
	return sqrt(p);
}

void closestPoint(int** givenArray, int k) {
	map<string, float> dictionary;

	for(int i=0; i<sizeof(givenArray); i++){
		float distance = calculateDistance(givenArray[i][0], givenArray[i][1]);
		string label = to_string(givenArray[i][0]) + "," + to_string(givenArray[i][1]);
		dictionary[label] = distance;
	}

	dictionary = mapSort(dictionary); // Should implement a function

	for(int i=0; i<k; i++){
		cout<<dictionary[to_string(i)]<<endl;
	}
}

int main(){

}
// Asked by Amazon
