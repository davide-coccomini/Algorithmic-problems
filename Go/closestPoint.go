// Given an array of points (x,y), find the K closest points to the origin (0,0)

package main

import "math"

func calculateDistance(x int, y int) float64 {
	pow := math.Pow(float64(x), 2) + math.Pow(float64(y), 2)
	return math.Sqrt(pow)
}

func closestPoint(givenArray [][]int, k int) {
	dictionary := make(map[string]float64)
	for _, point := range givenArray {
		distance := calculateDistance(point[0], point[1])
		label := string(point[0]) + "," + string(point[1])
		dictionary[label] = distance
	}

	dictionary = mapSort(dictionary) // Should implement a function

	for i := 0; i < k; i++ {
		// Print the first N elements
	}
}

// Asked by Amazon
