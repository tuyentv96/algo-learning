package day1

func maximalSquare(matrix [][]byte) int {
	rows, cols := len(matrix), len(matrix[0])
	max := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				sq := 1
				flag := true
				for flag && sq+i < rows && sq+j < cols {
					for r := i; r <= sq+i; r++ {
						if matrix[r][j+sq] == '0' {
							flag = false
							break
						}
					}

					for c := j; c <= sq+j; c++ {
						if matrix[i+sq][c] == '0' {
							flag = false
							break
						}
					}

					if flag {
						sq++
					}
				}

				if sq > max {
					max = sq
				}
			}
		}
	}

	return max * max
}
