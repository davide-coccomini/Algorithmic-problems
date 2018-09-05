/* Given an array, find the length of the longest subsequence

EXAMPLE
[2,1,5,9,4,3] --> 4 (because [1,2,3,4])

*/
package main

func longestSubsequence(givenArray []int) int {
	dictionary := make(map[int]bool)
	length := 0
	for i := 0; i < len(givenArray); i++ {
		dictionary[givenArray[i]] = true
	}

	for i := 0; i < len(givenArray); i++ {
		j := givenArray[i]
		tempLength := 0
		if !dictionary[givenArray[i]-1] {
			for dictionary[j] == true {
				tempLength++
				j++
				if tempLength > length {
					length := tempLength
				}
				if !dictionary[j] {
					break
				}
			}
		}
	}
	return length
}
