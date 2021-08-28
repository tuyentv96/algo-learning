package day1

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSoFar, minSoFar := nums[0], nums[0]
	result := maxSoFar

	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		tmp := max(cur, max(maxSoFar*cur, minSoFar*cur))
		minSoFar = min(cur, min(maxSoFar*cur, minSoFar*cur))

		maxSoFar = tmp
		result = max(result, maxSoFar)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
