package service

import (
	"errors"

	"idp-jobs-go/model"
	"idp-jobs-go/repository"
)

type JobService struct {
	repo repository.JobRepository
}

func NewJobService(r repository.JobRepository) *JobService {
	return &JobService{repo: r}
}

func (s *JobService) GetJob(id string) (*model.Job, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	return s.repo.GetByID(id)
}