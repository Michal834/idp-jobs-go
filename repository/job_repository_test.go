package repository

import (
	"testing"
)

func TestInMemoryJobRepository_GetByID(t *testing.T) {
	// Inicjalizacja repozytorium
	repo := NewInMemoryJobRepository()

	t.Run("should return existing job", func(t *testing.T) {
		job, err := repo.GetByID("1")

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if job == nil {
			t.Fatal("expected job to be found, got nil")
		}
		if job.JobID != "1" {
			t.Errorf("expected job ID 1, got %s", job.JobID)
		}
	})

	t.Run("should return error for non-existing job", func(t *testing.T) {
		_, err := repo.GetByID("999")

		if err == nil {
			t.Error("expected error for non-existing job, got nil")
		}
		if err.Error() != "job not found" {
			t.Errorf("expected 'job not found' error, got '%v'", err)
		}
	})
}
