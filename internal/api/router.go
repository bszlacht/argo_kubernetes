package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func MakeRouter(ms matrixService, h Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/matmul", h.HandleMatMul).Methods(http.MethodGet)
	return r
}
