package day2

func maximalSquare(matrix [][]byte) int {
	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows+1)
	max := 0

	for i := range dp {
		dp[i] = make([]int, cols+1)
	}

	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min3(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}

	return max * max
}

func min3(a, b, c int) int {
	result := a
	if result > b {
		result = b
	}

	if result > c {
		result = c
	}

	return result
}
