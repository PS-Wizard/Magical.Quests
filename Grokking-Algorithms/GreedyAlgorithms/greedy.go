package greedyalgorithms

import "fmt"

func returnIntersection(set1 []string, set2 map[string]struct{}) map[string]struct{} {
	intersections := make(map[string]struct{})
	for _, value := range set1 {
		if _, exists := set2[value]; exists {
			intersections[value] = struct{}{}
		}
	}

	return intersections
}
func Greedy(stations map[string][]string, stations_needed map[string]struct{}) {
	finalStations := make(map[string]struct{})
	for len(stations_needed) != 0 {
		stations_covered := map[string]struct{}{}
		best_station := ""
		covered := make(map[string]struct{})
		for stations, states := range stations {
			covered = returnIntersection(states, stations_needed)
			if len(covered) > len(stations_covered) {
				best_station = stations
				stations_covered = covered
			}
		}
		for key := range stations_covered {
			delete(stations_needed, key)
		}
		finalStations[best_station] = struct{}{}
	}

	fmt.Println(finalStations)
}
