package day1

func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]bool, numCourses)
	graph := make([][]int, numCourses)

	for i := 0; i < len(prerequisites); i++ {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(graph, visited, i) {
			return false
		}
	}

	return true
}

func dfs(graph [][]int, visited []bool, course int) bool {
	if visited[course] {
		return false
	}

	visited[course] = true
	for i := 0; i < len(graph[course]); i++ {
		if !dfs(graph, visited, graph[course][i]) {
			return false
		}
	}

	visited[course] = false
	return true
}
