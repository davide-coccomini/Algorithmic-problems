/* Given an array, I want to print all the subset in any order you wish

EXAMPLE:

[1,2] -> {}, {1,2}, {1}, {2}
Output:
> _
> 1,
> 2,
> 1,2,

*/
package main

func printSubset(givenArray []int) {
	size := len(givenArray)
	subset := make([]int, size)
	explore(givenArray, subset, 0)
}

func explore(givenArray []int, subset []int, i int) {
	if i == len(givenArray) {
		print(subset)
	} else {
		// First exploration, without a number
		subset[i] = 0
		explore(givenArray, subset, i+1)
		// Second exploration, with the sequel number
		subset[i] = givenArray[i]
		explore(givenArray, subset, i+1)
	}
}

func main() {
	givenArray := []int{1, 2}
	printSubset(givenArray)
}

// Asked by Facebook
