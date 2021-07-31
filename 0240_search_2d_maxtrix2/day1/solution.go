package day1

func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	i, j := 0, cols-1
	for i < rows && j >= 0 {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}

	return false
}
