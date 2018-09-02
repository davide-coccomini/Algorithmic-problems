/* Given a tree of integers and return the number of non-empty unival trees
A unival tree is a tree which contain all nodes with the same number.
*/

package main

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

//DUMB SOLUTION

//O(n)
func isUnivalDumb(root *Tree) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Value != root.Value {
		return false
	}
	if root.Right != nil && root.Right.Value != root.Value {
		return false
	}
	if isUnivalDumb(root.Left) && isUnivalDumb(root.Right) {
		return true
	}
	return false
}

// O(n^2)
func countUnivalsDumb(root *Tree) int {
	if root == nil {
		return 0
	}

	count := countUnivalsDumb(root.Left) + countUnivalsDumb(root.Right)

	if isUnivalDumb(root) {
		count++
	}
	return count
}

// GOOD SOLUTION O(n)

func exploreTree(root *Tree) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftCount, isLeftUnival := exploreTree(root.Left)
	rightCount, isRightUnival := exploreTree(root.Right)

	isUnival := true
	if !isLeftUnival || !isRightUnival {
		isUnival = false
	}

	if root.Left != nil && root.Left.Value != root.Value {
		isUnival = false
	}

	if root.Right != nil && root.Right.Value != root.Value {
		isUnival = false
	}

	if isUnival {
		return leftCount + rightCount + 1, true
	} else {
		return leftCount + rightCount, false
	}
}

func countUnivals(root *Tree) int {
	count, isUnival := exploreTree(root)
	return count
}

// Asked by Google
