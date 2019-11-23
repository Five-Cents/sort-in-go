package quick

func Sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {

		// function variables
		var pivotVal = arr[len(arr)/2]
		var pivot []int
		var lesser []int
		var greater []int

		// allocate variables to their respective arrays
		for _, value := range arr {
			if value > pivotVal {
				greater = append(greater, value)
			} else if value < pivotVal {
				lesser = append(lesser, value)
			} else if value == pivotVal {
				pivot = append(pivot, value)
			}
		}

		// anonymous function to help with recursive part of the algorithm
		concatenate := func(lesser []int, greater []int, pivot []int) []int {
			var concatArr []int

			concatArr = append(concatArr, Sort(lesser)...)
			concatArr = append(concatArr, pivot...)
			concatArr = append(concatArr, Sort(greater)...)

			return concatArr
		}

		return concatenate(lesser, greater, pivot)
	}
}
