package service

import (
	"testing"

	"idp-jobs-go/model"
)

type mockRepo struct{}

func (m *mockRepo) GetByID(id string) (*model.Job, error) {
	return &model.Job{JobID: id, Status: "ok"}, nil
}

func TestGetJob(t *testing.T) {
	svc := NewJobService(&mockRepo{})

	job, err := svc.GetJob("123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if job.JobID != "123" {
		t.Fatalf("expected 123, got %s", job.JobID)
	}
}