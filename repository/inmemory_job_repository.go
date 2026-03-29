package repository

import (
	"errors"
	"sync"
	"time"

	"idp-jobs-go/model"
)

type InMemoryJobRepository struct {
	data map[string]*model.Job
	mu   sync.RWMutex
}

func NewInMemoryJobRepository() *InMemoryJobRepository {
	repo := &InMemoryJobRepository{
		data: make(map[string]*model.Job),
	}

	repo.data["1"] = &model.Job{
		JobID:     "1",
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	return repo
}

func (r *InMemoryJobRepository) GetByID(id string) (*model.Job, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	job, ok := r.data[id]
	if !ok {
		return nil, errors.New("job not found")
	}

	return job, nil
}