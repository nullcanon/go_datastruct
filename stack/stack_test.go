package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestStack(t *testing.T) {
	stack := New()
	assert.Equal(t, uint(0), stack.Len())

	stack.Push(1)
	assert.Equal(t, uint(1), stack.Len())
	assert.Equal(t, 1, stack.Top())

	assert.Equal(t, 1, stack.Pop())

}