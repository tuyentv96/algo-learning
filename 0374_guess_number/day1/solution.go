package day1

func guess(n int) int {
	return 0
}

func guessNumber(n int) int {
	l, r := 1, n
	for l <= r {
		m := l + (r-l)/2
		g := guess(m)
		if g == 0 {
			return m
		}

		if g == -1 {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return 0
}
