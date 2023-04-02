package structures

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func createMatrices() map[string]Matrix {
	return map[string]Matrix{
		"array matrix":  NewArrayMatrix(3, 3, 255),
		"sparse matrix": NewSparseMatrix(255),
	}
}

func TestArrayMatrixNew(t *testing.T) {
	for name, matrix := range createMatrices() {
		t.Run(name, func(t *testing.T) {
			require.NotNil(t, matrix)
		})
	}
}

func TestArrayMatrixGet(t *testing.T) {
	t.Run("default", func(t *testing.T) {

		for name, matrix := range createMatrices() {
			t.Run(name, func(t *testing.T) {
				ret := matrix.Get(1, 1)

				require.Equal(t, 255, ret)
			})
		}
	})

	t.Run("set", func(t *testing.T) {
		for name, matrix := range createMatrices() {
			t.Run(name, func(t *testing.T) {
				matrix.Set(1, 1, 5)

				ret := matrix.Get(1, 1)

				require.Equal(t, 5, ret)
			})
		}
	})
}
