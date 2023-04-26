package api

import (
	"net/http"

	matmodel "github.com/bszlacht/argo_kubernetes/internal/matrix_service/model"
)

type matrixService interface {
	MatMul(a, b matmodel.Matrix) (matmodel.Matrix, error)
}

type Handler interface {
	HandleMatMul(http.ResponseWriter, *http.Request)
}

type handler struct {
	matrixService matrixService
}

func NewHandler(ms matrixService) Handler {
	return handler{
		matrixService: ms,
	}
}
