package greedyalgorithms_test

import (
	greedyalgorithms "greedyAlgorithms"
	"testing"
)

type Tests struct {
	states_needed map[string]struct{}
	stations      map[string][]string
}

func TestGreedy(t *testing.T) {
	tests := []Tests{
		{
			states_needed: map[string]struct{}{
				"mt": {},
				"wa": {},
				"or": {},
				"id": {},
				"nv": {},
				"ut": {},
				"ca": {},
				"az": {},
			},
			stations: map[string][]string{
				"kone":   {"id", "nv", "ut"},
				"ktwo":   {"wa", "id", "mt"},
				"kthree": {"or", "nv", "ca"},
				"kfour":  {"nv", "ut"},
				"kfive":  {"ca", "az"},
			},
		},
	}
	for _, val := range tests {
		greedyalgorithms.Greedy(val.stations, val.states_needed)
	}

}
