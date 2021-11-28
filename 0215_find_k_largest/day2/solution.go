package day2

func findKthLargest(nums []int, k int) int {
	n := len(nums)

	if k > n {
		return 0
	}

	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, i, n)
	}

	for i := 0; i < k-1; i++ {
		nums[0], nums[n-1-i] = nums[n-1-i], nums[0]
		heapify(nums, 0, n-i-1)
	}

	return nums[0]
}

// max heap
func heapify(nums []int, i int, n int) {
	largest := i

	left, right := 2*i+1, 2*i+2

	if left < n && nums[left] > nums[largest] {
		largest = left
	}

	if right < n && nums[right] > nums[largest] {
		largest = right
	}

	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		heapify(nums, largest, n)
	}
}

func findKthLargest2(nums []int, k int) int {
	n := len(nums)

	if k > n {
		return 0
	}

	for i := k - 1; i >= 0; i-- {
		heapify2(nums, i, k)
	}

	for i := k; i < n; i++ {
		if nums[i] > nums[0] {
			nums[0] = nums[i]
			heapify2(nums, 0, k)
		}
	}

	return nums[0]
}

// min heap
func heapify2(nums []int, i int, n int) {
	smallest := i

	left, right := 2*i+1, 2*i+2

	if left < n && nums[left] < nums[smallest] {
		smallest = left
	}

	if right < n && nums[right] < nums[smallest] {
		smallest = right
	}

	if smallest != i {
		nums[smallest], nums[i] = nums[i], nums[smallest]
		heapify2(nums, smallest, n)
	}
}
