package day1

// Given a sorted array of consecutive numbers from left most,
// find the first missing number, i.e., the hole.
func find(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		missing := missingNumber(nums, mid)
		if missing < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left-1] + k - missingNumber(nums, left-1)
}

func missingNumber(nums []int, i int) int {
	return nums[i] - nums[0] - i
}
