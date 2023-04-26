//go:build integration

package integration_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bszlacht/argo_kubernetes/internal/api"
	apimodel "github.com/bszlacht/argo_kubernetes/internal/api/model"
	"github.com/bszlacht/argo_kubernetes/internal/matrix_service"
	matmodel "github.com/bszlacht/argo_kubernetes/internal/matrix_service/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMatrixHandlers(t *testing.T) {
	ms := matrix_service.NewMatrixService()
	h := api.NewHandler(ms)
	r := api.MakeRouter(ms, h)

	t.Run("MatMul", func(t *testing.T) {
		makeMatMulRequest := func(body interface{}) *http.Response {
			var bodyReader bytes.Buffer
			err := json.NewEncoder(&bodyReader).Encode(body)
			require.NoError(t, err)

			req, _ := http.NewRequest("GET", "/matmul", &bodyReader)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			return w.Result()
		}

		t.Run("Bad request", func(t *testing.T) {
			req := map[string][]float64{
				"A": {1},
				"B": {2},
			}

			resp := makeMatMulRequest(req)
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		})

		t.Run("Success", func(t *testing.T) {
			req := apimodel.MatMulRequest{
				A: matmodel.Matrix{
					{1, 2},
					{3, 4},
				},
				B: matmodel.Matrix{
					{5, 6},
					{7, 8},
				},
			}
			expected := apimodel.MatMulResponse{
				C: matmodel.Matrix{
					{19, 22},
					{43, 50},
				},
			}

			resp := makeMatMulRequest(req)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			respBody := apimodel.MatMulResponse{}
			err := json.NewDecoder(resp.Body).Decode(&respBody)
			require.NoError(t, err)
			assert.Equal(t, expected, respBody)
		})
	})
}
