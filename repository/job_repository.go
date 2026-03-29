package repository

import "idp-jobs-go/model"

type JobRepository interface {
	GetByID(id string) (*model.Job, error)
}