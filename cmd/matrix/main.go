package main

import (
	"log"
	"net/http"

	"github.com/bszlacht/argo_kubernetes/internal/api"
	"github.com/bszlacht/argo_kubernetes/internal/matrix_service"
)

func main() {
	ms := matrix_service.NewMatrixService()
	h := api.NewHandler(ms)

	r := api.MakeRouter(ms, h)

	log.Print("Running server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
