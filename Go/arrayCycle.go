/* Write a function which, given an array of indexes, returns if there are cycles.
The number in the array are indexes of the array itself.

EXAMPLE:
[1,2,1,3,4,8] --> True
   ^ ^
In this example there's a cycle!

[1,2,3,4,5,6] --> False
*/

package main

func findCycle(givenArray []int) bool {
	p := 0
	q := 0
	l := len(givenArray)

	for true {

		p = givenArray[p]
		if p == q {
			return true
		}

		if p < 0 || p >= l { // Out of bounds
			return false
		}

		p = givenArray[p]

		if p == q {
			return true
		}

		q = givenArray[q]

		if p == q {
			return true
		}
	}
	return false
}

// Asked by Google
