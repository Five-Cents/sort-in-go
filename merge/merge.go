package merge

import (
	"math"
)

// Merge sort is a divide and conquer algorithm. It divides input arrays into two halves, calls
// itself for the two halves and then merges the two sorted halves.
func Sort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	var left, right []int
	var leftSize, rightSize int

	leftSize = int(math.Floor(float64(len(arr) / 2)))
	rightSize = leftSize

	for i := 0; i < leftSize; i++ {
		left = append(left, arr[i])
	}

	for i := rightSize; i < len(arr); i++ {
		right = append(right, arr[i])
	}

	left = Sort(left)
	right = Sort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	var merged []int

	// While the left and right array have elements sort them into the array 'merged' in order of smallest to largest.
	// Note, the array = append(array[:key], array[key+1:]...) is how elements are deleted from a slice in go :)
	for i := 0; i < len(left); i++ {
		for j := 0; j < len(right); j++ {
			if left[i] > right[j] {
				merged = append(merged, right[j])
				// We are removing an item from the slice, so we must also decrement j!
				right = append(right[:j], right[j+1:]...)
				j--
			} else {
				merged = append(merged, left[i])
				left = append(left[:i], left[i+1:]...)
				// We are removing an item from the left slice, so we decrement i and break out of the loop iterating
				// over the right slice.
				i--
				break
			}
		}
	}

	// Add the remaining elements from the other array.
	for _, val := range left {
		merged = append(merged, val)
	}
	for _, val := range right {
		merged = append(merged, val)
	}

	return merged
}
