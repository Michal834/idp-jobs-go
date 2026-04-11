package service

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"idp-jobs-go/model"
)

func loadJobFromFile(t *testing.T, path string) *model.Job {
	t.Helper()

	data, err := os.ReadFile(path)
	require.NoError(t, err)

	var job model.Job
	require.NoError(t, json.Unmarshal(data, &job))

	return &job
}
