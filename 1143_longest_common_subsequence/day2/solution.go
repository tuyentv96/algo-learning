package day2

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	return helper(text1, text2, 0, 0, dp)
}

func helper(text1, text2 string, p1, p2 int, dp [][]int) int {
	if dp[p1][p2] != -1 {
		return dp[p1][p2]
	}

	result := 0
	if text1[p1] == text2[p2] {
		result = 1 + helper(text1, text2, p1+1, p2+1, dp)
	} else {
		result = max(helper(text1, text2, p1+1, p2, dp), helper(text1, text2, p1, p2+1, dp))
	}

	dp[p1][p2] = result
	return dp[p1][p2]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
