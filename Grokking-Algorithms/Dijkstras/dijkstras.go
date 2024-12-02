package dijkstras

func Dijkstra(graph map[string]map[string]int, start string) map[string]int {
	costs := map[string]int{}
	visited := map[string]bool{}

	for nodes := range graph {
		costs[nodes] = 9999
	}
	costs[start] = 0
	findLowestCostNode := func() string {
		lowestCostNode := ""
		for node, cost := range costs {
			if cost < 9999 && !visited[node] {
				lowestCostNode = node
			}
		}
		return lowestCostNode
	}

	for {
		nextNode := findLowestCostNode()
		if nextNode == "" {
			break
		}
		cost := costs[nextNode]
		for neighbours, weight := range graph[nextNode] {
			totalCost := cost + weight
			if totalCost < costs[neighbours] {
				costs[neighbours] = totalCost
			}
		}
		visited[nextNode] = true
	}
	return costs

}
