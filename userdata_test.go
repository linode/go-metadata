package metadata

import (
	"context"
	"encoding/base64"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type UserdataMockClient struct {
	UserData         string
	GetUserDataError error
}

func (m *UserdataMockClient) GetUserData(ctx context.Context) (string, error) {
	if m.GetUserDataError != nil {
		return "", m.GetUserDataError
	}
	return m.UserData, nil
}

func TestGetUserData_Success(t *testing.T) {
	mockClient := &UserdataMockClient{
		UserData: base64.StdEncoding.EncodeToString([]byte("mock-user-data")),
	}

	userData, err := mockClient.GetUserData(context.Background())

	assert.NoError(t, err, "Expected no error")
	// Note "bW9jay11c2VyLWRhdGE=" is the encoded value
	assert.Equal(t, "bW9jay11c2VyLWRhdGE=", userData, "Unexpected user data")
}

func TestGetUserData_Error(t *testing.T) {
	mockClient := &UserdataMockClient{
		GetUserDataError: errors.New("mock error"),
	}

	userData, err := mockClient.GetUserData(context.Background())

	assert.Error(t, err, "Expected an error")
	assert.Equal(t, "", userData, "Expected empty user data")
	assert.EqualError(t, err, "mock error", "Unexpected error message")
}
