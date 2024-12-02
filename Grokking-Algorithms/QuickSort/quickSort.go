package quicksort

func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	pivot := array[len(array)/2]
	less, equal, greater := func() ([]int, []int, []int) {
		less := []int{}
		equal := []int{}
		greater := []int{}
		for _, val := range array {
			if val < pivot {
				less = append(less, val)
			} else if val == pivot {
				equal = append(equal, val)
			} else {
				greater = append(greater, val)
			}
		}
		return less, equal, greater
	}()

	appended := append(append(QuickSort(less), equal...), QuickSort(greater)...)
	return appended
}

