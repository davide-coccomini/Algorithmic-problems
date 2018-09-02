/* Take a string and return the first recurring characeter in it

EXAMPLE:
"ABCA" ->  "A"
"BCABA" -> "B"
"ABC" -> null

*/

// DUMP SOLUTION O(n^2)
package main

func getRecurring(givenString string) string {
	for i := 0; i < len(givenString); i++ {
		for j := 0; j < len(givenString); j++ {
			if givenString[i] == givenString[j] {
				return string(givenString[i])
			}
		}
	}

	return ""
}

// GOOD SOLUTION O(n)
// We use a structure, such as an hash table or a dictionary, which count the occurences of letters

func getRecurringHash(givenString string) string {
	occurences := make(map[string]int) // hash table or dictionary
	for i := 0; i < len(givenString); i++ {
		if occurences[string(givenString[i])] > 0 {
			return string(givenString[i])
		}
		occurences[string(i)] = 1
	}
	return ""
}

// Asked by Google
