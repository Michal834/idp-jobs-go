package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"idp-jobs-go/repository"
	"idp-jobs-go/service"
)

func TestGetJobEndpoint(t *testing.T) {
	repo := repository.NewInMemoryJobRepository()
	svc := service.NewJobService(repo)
	h := NewHandler(svc)

	req := httptest.NewRequest("GET", "/idp/jobs/1", nil)
	w := httptest.NewRecorder()

	h.Router().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}