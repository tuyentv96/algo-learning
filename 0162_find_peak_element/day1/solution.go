package day1

//func findPeakElement(nums []int) int {
//	for i:=0;i<len(nums)-1;i++{
//		if nums[i]>nums[i+1]{
//			return i
//		}
//	}
//
//	return len(nums)-1
//}

func findPeakElement(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
