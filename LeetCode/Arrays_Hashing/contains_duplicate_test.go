package arrayshashing

import "testing"

func TestContainsDuplicate(t *testing.T) {
	cases := []int{1, 2, 3, 4, 5, 1}
	want := true
	got := containsDuplicate(cases)
	if want != got {
		t.Errorf("Wanted %+v, Got %+v, Test: %+v", want, got, cases)
	}

}
