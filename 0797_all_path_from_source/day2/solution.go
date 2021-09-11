package day2

type Queue struct {
	arr []int
}

func NewQueue() Queue {
	return Queue{
		arr: make([]int, 0),
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.arr) == 0
}

func (q *Queue) Enqueue(i int) {
	q.arr = append(q.arr, i)
}

func (q *Queue) Dequeue() int {
	i := q.arr[0]
	q.arr = q.arr[1:]
	return i
}

func allPathsSourceTarget(graph [][]int) [][]int {
	var result [][]int
	var tmp []int
	n := len(graph)

	tmp = append(tmp, 0)
	dfs(graph, &result, tmp, 0, n-1)

	return result
}

func dfs(graph [][]int, result *[][]int, cur []int, src, dest int) {
	if src == dest {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}

	for i := 0; i < len(graph[src]); i++ {
		cur = append(cur, graph[src][i])
		dfs(graph, result, cur, graph[src][i], dest)
		cur = cur[:len(cur)-1]
	}

	return
}
