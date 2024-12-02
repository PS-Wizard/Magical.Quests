package binarysearch

type Input struct {
	array       []int
	numToSearch int
	want        int
}

func (i *Input) BinarySearch() int {
	lowerBound := 0
	upperBound := len(i.array) - 1
	for lowerBound <= upperBound {
		mid := (lowerBound + upperBound) / 2
		guess := i.array[mid]
		if guess == i.numToSearch {
			return mid
		}
		if guess < i.numToSearch {
			lowerBound = mid + 1
			continue
		}
		if guess > i.numToSearch {
			upperBound = mid - 1
			continue
		}
	}
	return 0
}
