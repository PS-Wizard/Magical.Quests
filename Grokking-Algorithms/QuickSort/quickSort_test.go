package quicksort_test

import (
	"quicksort"
	"reflect"
	"testing"
)

type Input struct {
	inputs []int
	wants  []int
}

func TestQuickSort(t *testing.T) {
	tests := []Input{
		{
			inputs: []int{1, 5, 3, 2, 4},
			wants:  []int{1, 2, 3, 4, 5},
		},
		{
			inputs: []int{10, 7, 8, 9, 1, 5},
			wants:  []int{1, 5, 7, 8, 9, 10},
		},
		{
			inputs: []int{0, -5, 3, 99, 21, 4, -1},
			wants:  []int{-5, -1, 0, 3, 4, 21, 99},
		},
		{
			inputs: []int{4},
			wants:  []int{4},
		},
		{
			inputs: []int{5, 5, 5, 5},
			wants:  []int{5, 5, 5, 5},
		},
		{
			inputs: []int{100, 99, 98, 97, 96, 95, 94},
			wants:  []int{94, 95, 96, 97, 98, 99, 100},
		},
		{
			inputs: []int{},
			wants:  []int{},
		},
	}
	for _, val := range tests {
		got := quicksort.QuickSort(val.inputs)
		if !reflect.DeepEqual(got, val.wants) {
			t.Errorf("got %v, want %v", got, val.wants)
		}
	}
}

