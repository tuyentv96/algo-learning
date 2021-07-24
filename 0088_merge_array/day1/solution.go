package day1

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := len(nums1)-1, m-1, n-1
	for j >= 0 && k >= 0 {
		if nums1[j] >= nums2[k] {
			nums1[i] = nums1[j]
			j--
		} else {
			nums1[i] = nums2[k]
			k--
		}

		i--
	}

	for j >= 0 {
		nums1[i] = nums1[j]
		j--
		i--
	}

	for k >= 0 {
		nums1[i] = nums1[k]
		k--
		i--
	}

	return
}
