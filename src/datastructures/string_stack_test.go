package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfPushToStackWorks(t *testing.T) {
	stack := StringStack{}
	stack.push("Hello")
	assert.Equal(t, []string{"Hello"}, stack.data)
}

func TestIfPopFromStringStackWorks(t *testing.T) {
	stack := StringStack{}
	stack.push("Hello")
	stack.push("World")
	assert.Equal(t, "World", stack.pop())
	assert.Equal(t, []string{"Hello"}, stack.data)
}

func TestIfIsStackEmptyWorks(t *testing.T) {
	stack := StringStack{}
	assert.Equal(t, true, stack.isEmpty())
}

func TestIfPeekFromStringStackWorks(t *testing.T) {
	stack := StringStack{}
	stack.push("Hello")
	stack.push("World")
	assert.Equal(t, "World", stack.peek())
	assert.Equal(t, []string{"Hello", "World"}, stack.data)
}
