package model

import "time"

type Job struct {
	JobID     string    `json:"jobId"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}