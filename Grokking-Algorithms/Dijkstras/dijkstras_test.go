package dijkstras_test

import (
	"dijkstras"
	"testing"
)

type Input struct {
	graph map[string]map[string]int
	start string
	want  map[string]int
}

func TestDijkstra(t *testing.T) {
	tests := []Input{
		{
			graph: map[string]map[string]int{
				"start": {
					"a": 6,
					"b": 2,
				},
				"a": {
					"fin": 1,
				},
				"b": {
					"a":   3,
					"fin": 5,
				},
				"fin": {},
			},
			start: "start",
            want:  map[string]int{"a": 5, "b": 2, "fin": 6,"start":0},
		},
		{
			graph: map[string]map[string]int{
				"start": {
					"x": 4,
					"y": 2,
				},
				"x": {
					"y": 5,
					"z": 10,
				},
				"y": {
					"z": 3,
				},
				"z": {},
			},
			start: "start",
            want:  map[string]int{"x": 4, "y": 2, "z": 5,"start":0},
		},
	}

	for _, test := range tests {
		result := dijkstras.Dijkstra(test.graph, test.start)
		if !equal(result, test.want) {
			t.Errorf("got %v, want %v", result, test.want)
		}
	}
}

// Helper function to compare maps
func equal(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

