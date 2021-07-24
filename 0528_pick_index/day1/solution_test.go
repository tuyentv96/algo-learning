package day1

import (
	"testing"
)

func Test(t *testing.T) {
	arr := []int{3, 14, 1, 7}
	r := Constructor(arr)
	r.PickIndex()
}
