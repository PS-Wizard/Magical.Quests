package breadthfirstsearch

import "fmt"

type Queue struct {
	queue []string
}

func NewQueue(length int) *Queue {
	return &Queue{
		queue: []string{},
	}
}

func (q *Queue) Push(element string) {
	q.queue = append(q.queue, element)
}

func (q *Queue) Pop() string {
	element := q.queue[0]
	q.queue = q.queue[1:]
	return element
}

func (q *Queue) QueueMakerHelper(graph map[string][]string, name string) {
	for _, val := range graph[name] {
		fmt.Printf("Pushed %s into the queue\n", val)
		q.Push(val)
	}
}

func BreadthFirstSearch(graph map[string][]string, length int, mangoSeller string) (string, bool) {
	q := NewQueue(length)
	q.QueueMakerHelper(graph, "you")

	visited := make(map[string]bool) // Track visited nodes

	for len(q.queue) != 0 {
		fmt.Println(len(q.queue))
		person := q.Pop()

		if visited[person] {
			continue
		}

		visited[person] = true // Mark as visited

		fmt.Println(person)
		if person == mangoSeller {
			return person, true
		}
		fmt.Printf("%s Is Not The Mango Seller, Adding Their Neighbours into the queue\n", person)
		q.QueueMakerHelper(graph, person)
	}

	return "", false
}

