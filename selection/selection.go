package selection

// In every iteration of selection sort, the minimum element from the unsorted subarray is picked up and moved into
// the sorted subarray.
func Sort(arr []int) []int {
	minIdx := 0
	for i := 0; i < len(arr); i++ {
		minIdx = i
		for j := 0; j < len(arr); j++ {
			if arr[j] > arr[minIdx] {
				minIdx = j
				arr[minIdx], arr[i] = arr[i], arr[minIdx]
			}
		}
	}
	return arr
}
