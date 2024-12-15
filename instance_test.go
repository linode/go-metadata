package metadata

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock client for testing purposes
type InstanceMockclient struct {
	Resp *InstanceData
	Err  error
}

func (m *InstanceMockclient) GetInstance(ctx context.Context) (*InstanceData, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resp, nil
}

func TestGetInstance_Success(t *testing.T) {
	// Create a mock client with a successful response
	mockClient := &InstanceMockclient{
		Resp: &InstanceData{
			ID:       1,
			Label:    "test-instance",
			Region:   "us-west",
			Type:     "standard",
			HostUUID: "abc123",
			Tags:     []string{"tag1", "tag2"},
			Specs: InstanceSpecsData{
				VCPUs:    2,
				Memory:   4096,
				GPUs:     0,
				Transfer: 2000,
				Disk:     50,
			},
			Backups: InstanceBackupsData{
				Enabled: true,
				Status:  String("active"),
			},
			Image: InstanceImageData{
				ID:    "linode/ubuntu24.04",
				Label: "Ubuntu 24.04 LTS",
			},
			AccountEUIID: "ABCD1234-ABCD-1234-ABCD1234ABCD1234",
		},
	}

	instance, err := mockClient.GetInstance(context.Background())

	// Assert the result
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, instance, "Expected non-nil instance")
	assert.Equal(t, "test-instance", instance.Label, "Unexpected instance label")
}

func TestGetInstance_Error(t *testing.T) {
	// Create a mock client with an error response
	mockClient := &InstanceMockclient{
		Err: errors.New("mock error"),
	}

	instance, err := mockClient.GetInstance(context.Background())

	assert.Error(t, err, "Expected an error")
	assert.Nil(t, instance, "Expected nil instance")
	assert.EqualError(t, err, "mock error", "Unexpected error message")
}

// Helper function to create a string pointer
func String(s string) *string {
	return &s
}
