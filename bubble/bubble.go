package bubble

/*
Bubble Sort works by iterating over an array multiple times and swapping adjacent elements with one another. When
considering time complexity, it is a very naive solution n^2.

Example:
given an array [3,1,2]

Step 1:
	OUTER LOOP:
	i = 0
	arr[i]	--> 3
		INNER LOOP:
		j = 0
		arr[j] --> 3
		arr[j+1] --> 1
		arr[j] > arr[j+1]
		swap(arr, j, j+1)
		INNER LOOP:
		j = 1
		arr[j] --> 3
		arr[j+1] --> 2
		arr[j] > arr[j+1] is true
		swap(arr, j, j+1)
		DONE (j = len(arr) -1)
	OUTER LOOP:
	i = 1
		INNER LOOP:



 */
func Sort(arr []int ) []int {
	// Iterate for each element - 1, don't need to do run over the entire array because the last item is already sorted
	for i := 0; i < len(arr) -1; i++ {
		for j := 0; j < len(arr) - 1 - i; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
	return arr
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}