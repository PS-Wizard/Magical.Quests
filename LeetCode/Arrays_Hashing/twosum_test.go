package arrayshashing

import (
	"reflect"
	"testing"
)

type TC struct {
	nums   []int
	target int
	want   []int
}

func TestTwoSum(t *testing.T) {
	testCases := []TC{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
	}
	for _, v := range testCases {
		got := twoSum(v.nums, v.target)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("For %+v, and target %+v, wanted %+v, but got %+v", v.nums, v.target, v.want, got)
		}
	}
}
