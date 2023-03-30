package structures

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStackNew(t *testing.T) {
	stack := NewStack[string]()

	require.NotNil(t, stack)
}

func TestStackIsEmptyTrue(t *testing.T) {
	stack := NewStack[string]()

	ret := stack.IsEmpty()

	require.True(t, ret)
}

func TestStackIsEmptyFalse(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("abc")

	ret := stack.IsEmpty()

	require.False(t, ret)
}

func TestStackPop(t *testing.T) {
	stack := NewStack[string]()

	require.Panics(t, func() {
		stack.Pop()
	})
}

func TestStackPeek(t *testing.T) {
	stack := NewStack[string]()

	require.Panics(t, func() {
		stack.Peek()
	})
}

func TestStackPushPeekPop(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("abc")

	peekRet := stack.Peek()
	popRet := stack.Pop()

	require.Equal(t, "abc", peekRet)
	require.Equal(t, "abc", popRet)
}

func TestStackPushPopPop(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("abc")

	ret := stack.Pop()

	require.Equal(t, "abc", ret)
	require.Panics(t, func() {
		stack.Pop()
	})
}

func TestStackPushPushPopPop(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("abc")
	stack.Push("def")

	def := stack.Pop()
	abc := stack.Pop()

	require.Equal(t, "abc", abc)
	require.Equal(t, "def", def)
}

func TestStackSize0(t *testing.T) {
	stack := NewStack[string]()

	ret := stack.Size()

	require.Equal(t, 0, ret)
}

func TestStackSize1(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("abc")

	ret := stack.Size()

	require.Equal(t, 1, ret)
}
