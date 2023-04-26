package matrix_service

import (
	"testing"

	"github.com/bszlacht/argo_kubernetes/internal/matrix_service/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMatMul(t *testing.T) {
	ms := NewMatrixService()

	t.Run("Equal dimensions", func(t *testing.T) {
		A := model.Matrix{
			{1, 2},
			{3, 4},
		}
		B := model.Matrix{
			{5, 6},
			{7, 8},
		}
		expectedC := model.Matrix{
			{19, 22},
			{43, 50},
		}

		C, err := ms.MatMul(A, B)
		require.NoError(t, err)
		assert.Equal(t, expectedC, C)
	})

	t.Run("Unequal dimensions", func(t *testing.T) {
		A := model.Matrix{
			{1, 2},
			{3, 4},
			{5, 6},
		}
		B := model.Matrix{
			{7},
			{8},
		}
		expectedC := model.Matrix{
			{23},
			{53},
			{83},
		}

		C, err := ms.MatMul(A, B)
		require.NoError(t, err)
		assert.Equal(t, expectedC, C)
	})
}
