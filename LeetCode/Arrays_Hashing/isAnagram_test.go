package arrayshashing

import "testing"

type TestCases struct {
	S     []string
	T     []string
	Wants []bool
}

func TestIsAnagram(t *testing.T) {
	testCases := TestCases{
		S:     []string{"anagram", "rat"},
		T:     []string{"nagaram", "car"},
		Wants: []bool{true, false},
	}
	for i, v := range testCases.S {
		want, got := testCases.Wants[i], isAnagram(v, testCases.T[i])
		if want != got {
			t.Errorf("Wanted %+v, Got %+v for the strings: %s, %s", want, got, v, testCases.T[i])
		}
	}
}
