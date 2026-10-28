package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"mini-ozon/handlers"
	"mini-ozon/intern/models/orders"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubService struct{}

func (s *stubService) ChangeStatus(ctx context.Context, id int, status orders.Status) error {
	return nil
}
func TestUpdateStatus_success(t *testing.T) {
	reqBody := handlers.Req{
		OrderID: 1,
		Status:  "Ready",
	}
	data, err := json.Marshal(reqBody)
	require.NoError(t, err)
	req := httptest.NewRequest(
		http.MethodPatch,
		"/orders",
		bytes.NewBuffer(data),
	)
	req.Header.Set("Content-Type", "application/json")
	stub := &stubService{}
	handler := handlers.NewChangeHandler(stub)
	rec := httptest.NewRecorder()
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", handler.ChangeStatus)

	mux.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Result().StatusCode)
	assert.JSONEq(t, `{"message": "successfull update"}`, rec.Body.String())
}

func TestUpdateStatus_invalidMethod(t *testing.T) {
	reqBody := handlers.Req{
		OrderID: 1,
		Status:  "pending",
	}
	data, err := json.Marshal(reqBody)
	require.NoError(t, err)
	req := httptest.NewRequest(
		http.MethodPut,
		"/orders",
		bytes.NewBuffer(data),
	)
	rec := httptest.NewRecorder()
	mux := http.NewServeMux()
	h := handlers.NewChangeHandler(&stubService{})
	mux.HandleFunc("/orders", h.ChangeStatus)

	mux.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rec.Result().StatusCode)
}
