package selectionsort

type Input struct {
	array []int
	want  []int
}

func (a *Input) selectionSort() {
	for j := 0; j < len(a.array); j++ {
		smallest := j
		for k := j + 1; k < len(a.array); k++ {
			if a.array[k] < a.array[smallest] {
				smallest = k
			}
		}
		a.array[j], a.array[smallest] = a.array[smallest], a.array[j]
	}
}
