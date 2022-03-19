package day4

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)

	dp := make([]int, m+1)
	prev := make([]int, m+1)

	for col := n - 1; col >= 0; col-- {
		for row := m - 1; row >= 0; row-- {
			if text2[col] == text1[row] {
				dp[row] = 1 + prev[row+1]
			} else {
				dp[row] = max(dp[row+1], prev[row])
			}
		}

		tmp := prev
		prev = dp
		dp = tmp
	}

	return prev[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
