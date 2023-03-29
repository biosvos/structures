package structures

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUniqueIterator(t *testing.T) {
	iterator := NewUniqueIterator[uint64, *Number]([]*Number{{1}, {2}, {3}, {3}}...)

	require.True(t, iterator.HasNext())
	require.Equal(t, &Number{1}, iterator.Next())
	require.Equal(t, &Number{2}, iterator.Next())
	require.Equal(t, &Number{3}, iterator.Next())
	require.False(t, iterator.HasNext())
	require.Panics(t, func() {
		iterator.Next()
	})
}
