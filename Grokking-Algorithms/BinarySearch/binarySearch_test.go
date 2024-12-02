package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	inputs := []Input{
		{
			array:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			numToSearch: 6,
			want:        5,
		},
		{
			array:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			numToSearch: 11,
			want:        0,
		},
	}
	for _, input := range inputs {
		got := input.BinarySearch()
		if got != input.want {
			t.Errorf("Wanted %v but got %v", input.want, got)
		}
	}

}
