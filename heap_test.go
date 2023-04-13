package structures

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func TestHeapIsEmpty(t *testing.T) {
	t.Run("", func(t *testing.T) {
		heap := NewHeap()
		heap.Push(1)

		ret := heap.IsEmpty()

		require.False(t, ret)
	})

	t.Run("", func(t *testing.T) {
		heap := NewHeap()

		ret := heap.IsEmpty()

		require.True(t, ret)
	})
}

func TestHeapPop(t *testing.T) {
	t.Run("", func(t *testing.T) {
		heap := NewHeap()
		heap.Push(1)

		ret := heap.Pop()

		require.Equal(t, 1, ret)
	})

	t.Run("", func(t *testing.T) {
		heap := NewHeap()
		heap.Push(2)

		ret := heap.Pop()

		require.Equal(t, 2, ret)
	})

	t.Run("random", func(t *testing.T) {
		var data []int
		for i := 0; i < 10; i++ {
			data = append(data, i)
		}
		rand.Shuffle(10, func(i, j int) {
			data[i], data[j] = data[j], data[i]
		})
		heap := NewHeap()
		for _, datum := range data {
			heap.Push(datum)
		}
		for i := 0; i < 10; i++ {
			ret := heap.Pop()
			require.Equalf(t, i, ret, "%v", data)
		}
	})

	t.Run("", func(t *testing.T) {
		data := []int{9, 7, 8, 6, 0, 5, 3, 1, 4, 2}
		heap := NewHeap()
		for _, datum := range data {
			heap.Push(datum)
		}
		for i := 0; i < 10; i++ {
			ret := heap.Pop()
			require.Equalf(t, i, ret, "%v", data)
		}
	})
}
