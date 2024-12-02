package selectionsort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	inputs := []Input{
		{
			array: []int{5, 2, 1, 3, 4, 6},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			array: []int{1, 1, 2, 6, 3, 4},
			want:  []int{1, 1, 2, 3, 4, 6},
		},
	}
	for _, input := range inputs {
		fmt.Printf("For %v\n", input.array)
		input.selectionSort()
		fmt.Printf("Got Back: %v\n", input.array)
		if !reflect.DeepEqual(input.array, input.want) {
			t.Errorf("Got %v, Want %v", input.array, input.want)
		}
	}

}
