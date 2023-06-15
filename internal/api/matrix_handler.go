package api

import (
	"encoding/json"
	"net/http"

	apimodel "github.com/bszlacht/argo_kubernetes/internal/api/model"
)

func jsonResponse(w http.ResponseWriter, status int, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}

func (h handler) HandleMatMul(w http.ResponseWriter, r *http.Request) {
	req := apimodel.MatMulRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, apimodel.Error{
			Message: "Failed to decode body",
		})
		return
	}

	C, err := h.matrixService.MatMul(req.A, req.B)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, apimodel.Error{
			Message: "Internal error",
		})
		return
	}

	jsonResponse(w, http.StatusOK, apimodel.MatMulResponse{
		C: C,
	})
}
