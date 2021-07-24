package day1

func islandPerimeter(grid [][]int) int {
	perimiter := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				perimiter += 4
				if i-1 >= 0 && grid[i-1][j] == 1 {
					perimiter -= 2
				}

				if j-1 >= 0 && grid[i][j-1] == 1 {
					perimiter -= 2
				}
			}
		}
	}

	return perimiter
}
