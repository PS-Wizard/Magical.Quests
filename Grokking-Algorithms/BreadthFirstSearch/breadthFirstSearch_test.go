package breadthfirstsearch_test

import (
	breadthfirstsearch "breadthFirstSearch"
	"fmt"
	"testing"
)

type Input struct {
	graph map[string][]string
	want  string
}

func TestBreadthFirstSearch(t *testing.T) {
	tests := []Input{
		{
			graph: map[string][]string{
				"you":    {"alice", "bob", "claire"},
				"bob":    {"anuj", "peggy"},
				"alice":  {"peggy"},
				"claire": {"thom", "jhonny"},
				"peggy":  {},
				"thom":   {},
				"jhonny": {},
			},
			want: "jhonny",
		},
		{
			graph: map[string][]string{
				"you":    {"alice", "bob", "claire"},
				"bob":    {"anuj", "peggy"},
				"alice":  {"peggy"},
				"claire": {"thom", "jhonny"},
				"peggy":  {},
				"thom":   {},
				"jhonny": {},
			},
			want: "",
		},
	}

	for idx, val := range tests {
		name, _ := breadthfirstsearch.BreadthFirstSearch(val.graph, len(val.graph), val.want)
		if name != val.want {
			t.Errorf("Got %v, Want %v", name, val.want)
		}
        fmt.Println("------ End of Test: ---------", idx)
	}
}
