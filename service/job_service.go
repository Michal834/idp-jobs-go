package service

import (
	"errors"

	"idp-jobs-go/model"
	"idp-jobs-go/repository"
)

// JobService provides operations for working with Job entities.
// It acts as a business logic layer between the API and the repository.
type JobService struct {
	repo repository.JobRepository
}

// NewJobService creates a new instance of JobService with the given repository.
//
// The repository is responsible for data access, while the service
// handles validation and business logic.
func NewJobService(r repository.JobRepository) *JobService {
	return &JobService{repo: r}
}

// GetJob retrieves a Job by its ID.
//
// It validates the input and delegates the retrieval to the repository.
//
// Returns:
//   - *model.Job: the job if found
//   - error: if the ID is empty or if the repository returns an error
//
// If id is empty, it returns an error without calling the repository.
func (s *JobService) GetJob(id string) (*model.Job, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	return s.repo.GetByID(id)
}
