package structures

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArrayMatrixGetPanic(t *testing.T) {
	matrix := NewArrayMatrix(3, 3, 255)
	require.Panics(t, func() {
		matrix.Get(5, 5)
	})
}

func TestArrayMatrixSetPanic(t *testing.T) {
	matrix := NewArrayMatrix(3, 3, 255)
	require.Panics(t, func() {
		matrix.Set(5, 5, 7)
	})
}
