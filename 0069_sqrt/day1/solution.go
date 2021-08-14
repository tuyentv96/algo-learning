package day1

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	left, right := 1, x
	for left < right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		}

		if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
