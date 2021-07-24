package day2

func islandPerimeter(grid [][]int) int {
	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				res += 4

				if i > 0 && grid[i-1][j] == 1 {
					res -= 2
				}

				if j > 0 && grid[i][j-1] == 1 {
					res -= 2
				}
			}
		}
	}

	return res
}
