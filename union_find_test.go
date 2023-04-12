package structures

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnionFindNew(t *testing.T) {
	uf := NewUnionFind[int]()
	require.NotNil(t, uf)
}

func TestUnionFindRoot(t *testing.T) {
	t.Run("", func(t *testing.T) {
		uf := NewUnionFind[int]()
		ret := uf.Root(1)
		require.Equal(t, 1, ret)
	})
	t.Run("", func(t *testing.T) {
		uf := NewUnionFind[int]()
		uf.Merge(2, 1)
		ret := uf.Root(1)
		require.Equal(t, 2, ret)
	})
}

func TestUnionFindIsSame(t *testing.T) {
	t.Run("", func(t *testing.T) {
		uf := NewUnionFind[int]()
		ret := uf.IsSame(1, 2)
		require.False(t, ret)
	})
	t.Run("", func(t *testing.T) {
		uf := NewUnionFind[int]()
		uf.Merge(2, 1)

		ret := uf.IsSame(1, 2)

		require.True(t, ret)
	})
}
