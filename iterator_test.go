package structures

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIterator(t *testing.T) {
	iterator := NewIterator([]uint64{1, 2, 3}...)

	require.True(t, iterator.HasNext())
	require.Equal(t, uint64(1), iterator.Next())
	require.Equal(t, uint64(2), iterator.Next())
	require.Equal(t, uint64(3), iterator.Next())
	require.False(t, iterator.HasNext())
	require.Panics(t, func() {
		iterator.Next()
	})
}
