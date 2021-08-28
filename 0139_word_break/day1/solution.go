package day1

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}
