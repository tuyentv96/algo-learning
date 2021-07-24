package day1

// 0,1,2 => Red, White, Blue
func sortColors(nums []int) {
	redIndex := 0
	blueIndex := len(nums) - 1
	for i := 0; i <= blueIndex; i++ {
		if nums[i] == 0 {
			tmp := nums[redIndex]
			nums[redIndex] = nums[i]
			nums[i] = tmp
			redIndex++
		} else if nums[i] == 2 {
			tmp := nums[blueIndex]
			nums[blueIndex] = nums[i]
			nums[i] = tmp
			blueIndex--
			i--
		}
	}
}
