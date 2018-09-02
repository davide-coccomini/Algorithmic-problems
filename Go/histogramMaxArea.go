/* Given an array which represents an histogram, return the area of the largest rectangle. Consider each column wide 1 and tall the number in the array.

 EXAMPLE
 [1,2,2]


   __ __
__|  |  |
  |  |  |

In this case the greatest rectangle is the one made of the two colums (2 and 3) so the area is (1+1)*2 = 4
*/
package main

//import "github.com/golang-collections/collections/stack"

func histogramMaxArea(givenArray []int) int {
	stack := New() * Stack()
	max := 0
	i := 0

	for i < len(givenArray) {
		if stack.isEmpty() || stack.peek() <= givenArray[i] {
			stack.push(i)
			i++
		} else {
			tempMax := stack.pop()
			area := 0
			if stack.isEmpty {
				area = givenArray[tempMax] * (i - 1)
			} else {
				area = givenArray[tempMax] * (i - 1 - stack.peek())
			}
			if area > max {
				max = area
			}
		}
	}
	for stack.peek() {
		tempMax := stack.pop()
		area := 0
		if stack.isEmpty {
			area = givenArray[tempMax] * (i - 1)
		} else {
			area = givenArray[tempMax] * (i - 1 - stack.peek())
		}
		if area > max {
			max = area
		}
	}
	return max
}

/*
func main(){
	arr := []int{1,2,2,3}
	maxArea = histogramMaxArea(arr)
	print(str(maxArea))
}*/
