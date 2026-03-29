package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"idp-jobs-go/service"
)

type Handler struct {
	svc *service.JobService
}

func NewHandler(s *service.JobService) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/idp/jobs/{jobId}", h.getJob)
	return r
}

func (h *Handler) getJob(w http.ResponseWriter, r *http.Request) {
	jobID := chi.URLParam(r, "jobId")

	job, err := h.svc.GetJob(jobID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}