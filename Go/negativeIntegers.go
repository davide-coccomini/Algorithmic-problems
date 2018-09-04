/* Find the number of negative integers in a row-wise / column-wise sorted matrix

EXAMPLE

-3 -2 -1  1
-2  2  3  4
 4  5  7  8

Output: 4
*/

//A good solution could be to use the fact that the matrix is sorted, to find the first negative number in a row and understand that the next number will be negative too
package main

// O(n+m)

func negativeIntegers(matrix [][]int, n int, m int) int {
	count := 0
	i := 0
	j := m - 1
	for j >= 0 && i < n {
		if matrix[i][j] < 0 {
			i++
		} else {
			j--
		}
	}
	return count
}

// Asked by Amazon
