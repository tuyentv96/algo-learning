package day3

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for col := n - 1; col >= 0; col-- {
		for row := m - 1; row >= 0; row-- {
			if text2[col] == text1[row] {
				dp[row][col] = 1 + dp[row+1][col+1]
			} else {
				dp[row][col] = max(dp[row+1][col], dp[row][col+1])
			}
		}
	}

	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
