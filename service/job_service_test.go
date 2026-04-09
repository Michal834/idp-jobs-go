package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"idp-jobs-go/model"
)

type mockRepo struct {
	called bool
	job    *model.Job
	err    error
}

func (m *mockRepo) GetByID(jobID string) (*model.Job, error) {
	m.called = true
	return m.job, m.err
}

func TestGetJob(t *testing.T) {
	expectedJobFileNamePath := "testdata/Test_getJob/"
	expectedJobFileName := "expectedJob.json"

	tests := []struct {
		name           string
		jobID          string
		mockErr        error
		expectError    bool
		expectRepoCall bool
	}{
		{
			name:           "01.Happy_path_job_exists",
			jobID:          "123",
			expectError:    false,
			expectRepoCall: true,
		},
		{
			name:           "02.Empty_id",
			jobID:          "",
			expectError:    true,
			expectRepoCall: false,
		},
		{
			name:           "03.Job_not_found_nil_return",
			jobID:          "999",
			expectError:    false,
			expectRepoCall: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			var mockJob *model.Job

			// TODO: think about better abstraction here
			if tt.name == "01.Happy_path_job_exists" {
				mockJob = loadJobFromFile(t, expectedJobFileNamePath+tt.name+"/"+expectedJobFileName)
			}

			repo := &mockRepo{
				job: mockJob,
				err: tt.mockErr,
			}

			svc := NewJobService(repo)

			// When
			result, err := svc.GetJob(tt.jobID)

			// Then
			if tt.expectError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectRepoCall, repo.called)

			if tt.name == "01.Happy_path_job_exists" {
				expected := loadJobFromFile(t, expectedJobFileNamePath+tt.name+"/"+expectedJobFileName)

				require.NotNil(t, result)

				assert.Equal(t, expected.JobID, result.JobID)
				assert.Equal(t, expected.Status, result.Status)
				assert.True(t, expected.CreatedAt.Equal(result.CreatedAt))
			} else {
				assert.Nil(t, result)
			}
		})
	}
}
