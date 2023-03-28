package structures

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

var _ Identifier[uint64] = &Number{}

type Number struct {
	id uint64
}

func (n *Number) Identify() uint64 {
	return n.id
}

func TestSetNew(t *testing.T) {
	set := NewSet[uint64, *Number]()
	require.NotNil(t, set)
}

func TestSetHas(t *testing.T) {
	set := NewSet[uint64, *Number]()
	set.Add(&Number{3})
	set.Delete(&Number{3})

	ret := set.Has(&Number{3})

	require.False(t, ret)
}

func TestSetHas2(t *testing.T) {
	set := NewSet[uint64, *Number]()
	set.Add(&Number{3})

	ret := set.Has(&Number{3})

	require.True(t, ret)
}

func TestSetSlice(t *testing.T) {
	set := NewSet[uint64, *Number]()

	slice := set.Slice()

	require.Nil(t, slice)
	require.Equal(t, 0, len(slice))
}

func TestSetSliceOne(t *testing.T) {
	set := NewSet[uint64, *Number]()
	set.Add(&Number{3})
	set.Add(&Number{4})

	slice := set.Slice()
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].id < slice[j].id
	})

	require.EqualValues(t, []*Number{{3}, {4}}, slice)
}
