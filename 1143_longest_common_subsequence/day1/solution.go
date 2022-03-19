package day1

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

func helper(text1, text2 string, i, j int, dp [][]int) int {
	if dp[i][j] != -1 {
		return dp[i][j]
	}

	case1 := helper(text1, text2, i+1, j, dp)

	case2 := 0
	first := find(text2[j:], text1[i])
	if first != -1 {
		case2 = 1 + helper(text1, text2, i+1, first+1, dp)
	}

	dp[i][j] = max(case1, case2)

	return dp[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func find(s string, c byte) int {
	for i := range s {
		if s[i] == c {
			return i
		}
	}

	return -1
}
