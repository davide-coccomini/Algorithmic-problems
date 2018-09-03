/* Given an array which represents an histogram, return the area of the largest rectangle. Consider each column wide 1 and tall the number in the array.

 EXAMPLE
 [1,2,2]


   __ __
__|  |  |
  |  |  |

In this case the greatest rectangle is the one made of the two colums (2 and 3) so the area is (1+1)*2 = 4
*/

#include <stack>  

using namespace std;



int histogramMaxArea(int* givenArray) {
	stack<int> s;
	int max = 0;
	int i = 0;

	while(i < sizeof(givenArray)) {
		if(s.empty() || s.peek() <= givenArray[i]){
			s.push(i);
			i++;
		} else {
			int tempMax = s.pop();
			int area = 0;
			if(s.empty()) {
				area = givenArray[tempMax] * (i - 1);
			} else {
				area = givenArray[tempMax] * (i - 1 - s.peek());
			}
			if(area > max) {
				max = area;
			}
		}
	}
	while(s.peek()) {
		int tempMax = s.pop();
		int area = 0;
		if(s.empty()) {
			area = givenArray[tempMax] * (i - 1);
		} else {
			area = givenArray[tempMax] * (i - 1 - s.peek());
		}
		if(area > max){
			max = area;
		}
	}
	return max;
}

int main(){


}
