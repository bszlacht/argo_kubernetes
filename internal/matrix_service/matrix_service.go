package matrix_service

import (
	"fmt"

	"github.com/bszlacht/argo_kubernetes/internal/matrix_service/model"
	"gonum.org/v1/gonum/mat"
)

type MatrixService interface {
	MatMul(a, b model.Matrix) (c model.Matrix, err error)
}

type matrixService struct {
}

func NewMatrixService() MatrixService {
	return matrixService{}
}

func matrixToDense(m model.Matrix) *mat.Dense {
	rows, cols := len(m), len(m[0])

	flattened := make([]float64, 0, rows*cols)
	for _, row := range m {
		flattened = append(flattened, row...)
	}

	return mat.NewDense(rows, cols, flattened)
}

func denseToMatrix(d *mat.Dense) model.Matrix {
	rows, _ := d.Dims()

	matrix := make(model.Matrix, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = d.RawRowView(i)
	}

	return matrix
}

func matmulResultDense(A, B *mat.Dense) *mat.Dense {
	rows, _ := A.Dims()
	_, cols := B.Dims()

	return mat.NewDense(rows, cols, make([]float64, rows*cols))
}

func (ms matrixService) MatMul(a, b model.Matrix) (c model.Matrix, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	A := matrixToDense(a)
	B := matrixToDense(b)
	C := matmulResultDense(A, B)

	C.Mul(A, B)

	return denseToMatrix(C), nil
}
