package main

// Some examples
func main() {
	// allSubsets
	givenArray := []int{1, 2}
	printSubset(givenArray)

	// arrayCycle
	givenArray = []int{1, 2, 1, 3, 4, 8}
	findCycle(givenArray)

	// arrayIntersection
	arr1 := []int{2, 8, 9, 10, 15}
	arr2 := []int{2, 9, 11, 17, 15}
	arr3 := []int{2, 3, 9, 17, 15}
	result0 := arrayIntersection(arr1, arr2, arr3)
	for _, number := range result0 {
		print(string(number))
	}

	// firstRecurring
	result1 := getRecurringHash("ABCA")
	print(result1)

	// longestSubsequence
	arr := []int{1, 2, 3, 5, 4, 8, 9}
	result2 := longestSubsequence(arr)
	print(string(result2))

	// negativeIntegers
	n := 3
	m := 4
	matrix := make([][]int, n)
	matrix[0] = []int{-3, -2, -1, 1}
	matrix[1] = []int{-2, 2, 3, 4}
	matrix[2] = []int{4, 5, 7, 8}
	result3 := negativeIntegers(matrix, n, m)
	print(string(result3))

	// steps
	n = 2
	x := []int{1, 2}
	result4 := numWays(n, x)
	print(string(result4))

}
