package insertion

/*
Insertion Sort-Insertion sort is a simple sorting algorithm that builds the final sorted array (or list) one item at a
time. It is much less efficient on large lists than more advanced algorithms such as quicksort, heapsort, or merge sort.

However, insertion sort provides several advantages:

1) 	Simple to implement.
2) 	Efficient for (quite) small data sets, much like other quadratic sorting algorithms
3) 	More efficient in practice than most other simple quadratic (i.e., O(n2)) algorithms such as selection sort or bubble sort
4) 	Adaptive, i.e., efficient for data sets that are already substantially sorted: the time complexity is O(nk) when each
		element in the input is no more than k places away from its sorted position
5) 	Stable; i.e., does not change the relative order of elements with equal keys
6) 	In-place; i.e., only requires a constant amount O(1) of additional memory space
7) 	Online; i.e., can sort a list as it receives it

source: https://everythingcomputerscience.com/algorithms/CSSorting_Algorithms.html
*/
func Sort(arr []int) []int {
	// Iterate over the array.
	for i := 1; i < len(arr); i++ {
		// Keep track of the 'key', the int that 'i' is currently iterating on.
		// We are unsure of how many times we need to iterate, so let's use a while loop :)
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = key
	}
	return arr
}
