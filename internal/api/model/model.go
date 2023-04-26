package model

import (
	mathmodel "github.com/bszlacht/argo_kubernetes/internal/matrix_service/model"
)

type Error struct {
	Message string
}

type MatMulRequest struct {
	A mathmodel.Matrix
	B mathmodel.Matrix
}

type MatMulResponse struct {
	C mathmodel.Matrix
}
