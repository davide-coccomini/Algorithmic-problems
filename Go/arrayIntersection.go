/* Write a function that given 3 ordered arrays, find the intersection

EXAMPLE
Arr1 = [2,6,9,11,13,17]
Arr2 = [3,6,7,10,13,17]
Arr3 = [4,5,6,9,11,13]

Output: [6,13]
*/
package main

func arrayIntersection(arr1 []int, arr2 []int, arr3 []int) []int {

	result := []int{}
	dictionary := make(map[int]int)

	for _, number := range arr1 {
		dictionary[number] = 1
	}

	for _, number := range arr2 {
		if dictionary[number] > 0 {
			dictionary[number]++
		}
	}
	for _, number := range arr3 {
		if dictionary[number] > 0 {
			dictionary[number]++
			if dictionary[number] >= 3 {
				result = append(result, number)
			}
		}
	}

	return result

}

/*
// Test call
func main(){
	arr1 := int[]{2,8,9,10,15}
	arr2 := int[]{2,9,11,17,15}
	arr3 := int[]{2,3,9,17,15}
	result := arrayIntersection(arr1,arr2,arr3)
	print(str(result))
}
*/
