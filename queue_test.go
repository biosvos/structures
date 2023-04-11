package structures

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue(1)
	require.NotNil(t, queue)
}

func TestQueueDequeue(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		queue := NewQueue[int]()
		queue.Enqueue(1)

		ret := queue.Dequeue()

		require.Equal(t, 1, ret)
	})
	t.Run("2", func(t *testing.T) {
		queue := NewQueue[int]()
		queue.Enqueue(2)

		ret := queue.Dequeue()

		require.Equal(t, 2, ret)
	})
	t.Run("1,2", func(t *testing.T) {
		queue := NewQueue[int]()
		queue.Enqueue(1)
		queue.Enqueue(2)

		a := queue.Dequeue()
		b := queue.Dequeue()

		require.Equal(t, 1, a)
		require.Equal(t, 2, b)
	})
	t.Run("panic", func(t *testing.T) {
		queue := NewQueue[int]()

		require.Panics(t, func() {
			_ = queue.Dequeue()
		})
	})
}

func TestQueueIsEmpty(t *testing.T) {
	t.Run("", func(t *testing.T) {
		queue := NewQueue[int]()
		ret := queue.IsEmpty()
		require.True(t, ret)
	})

	t.Run("", func(t *testing.T) {
		queue := NewQueue[int]()
		queue.Enqueue(1)
		ret := queue.IsEmpty()
		require.False(t, ret)
	})
}
