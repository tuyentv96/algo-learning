package day1

// Given a sorted array of consecutive numbers from 11 to NN,
// find the first missing number, i.e., the hole.
func find(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] != mid+1 && nums[mid-1] == mid {
			return mid + 1
		}

		if nums[mid] == mid+1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
