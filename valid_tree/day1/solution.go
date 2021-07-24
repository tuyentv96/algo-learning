package day1

func validTree(n int, edges [][2]int) bool {
	graph := make([][]int, n)
	for i := range edges {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}

	visited := make([]bool, n)

	if hasCycle(visited, graph, 0, -1) {
		return false
	}

	// check graph connected
	for _, v := range visited {
		if !v {
			return false
		}
	}

	return true
}

func hasCycle(visited []bool, graph [][]int, node int, parent int) bool {
	visited[node] = true

	for _, next := range graph[node] {
		// if next node is visited but it is not parent of cur node => has cycle
		if visited[node] && parent != next {
			return false
		}

		// if next node is not visited, visit next node
		if !visited[node] && hasCycle(visited, graph, next, node) {
			return false
		}
	}

	return false
}
