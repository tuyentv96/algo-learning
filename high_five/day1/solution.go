package day1

import "sort"

func highFive(items [][]int) [][]int {
	m := make(map[int][]int)
	for _, item := range items {
		id, score := item[0], item[1]
		scores, ok := m[id]
		if !ok {
			m[id] = make([]int, 5)
		}

		j := -1
		for i := 0; i < len(scores); i++ {
			if score > scores[i] {
				j = i
			}
		}

		if j != -1 {
			scores[j] = score
		}
	}

	ids := make([]int, 0)
	for id, _ := range m {
		ids = append(ids, id)
	}

	sort.Ints(ids)

	res := make([][]int, 0)
	for _, id := range ids {
		sum := 0
		for _, score := range m[id] {
			sum += score
		}

		res = append(res, []int{id, sum / 5})
	}

	return res
}
