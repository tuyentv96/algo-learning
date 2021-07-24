package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(3)
	assert.Equal(t, -2, stack.GetMin())
	stack.Pop()
	assert.Equal(t, -2, stack.GetMin())
}
