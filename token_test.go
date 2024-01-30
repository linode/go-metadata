package metadata

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TokenMockclient struct {
	Token            string
	GenerateTokenErr error
}

func (m *TokenMockclient) GenerateToken(ctx context.Context, opts ...TokenOption) (string, error) {
	if m.GenerateTokenErr != nil {
		return "", m.GenerateTokenErr
	}
	return m.Token, nil
}

func TestGenerateToken_Success(t *testing.T) {
	mockClient := &TokenMockclient{
		Token: "mock-token-value",
	}

	token, err := mockClient.GenerateToken(context.Background())

	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, "mock-token-value", token, "Unexpected token")
}

func TestGenerateToken_Error(t *testing.T) {
	mockClient := &TokenMockclient{
		GenerateTokenErr: errors.New("mock error"),
	}

	token, err := mockClient.GenerateToken(context.Background())

	assert.Error(t, err, "Expected an error")
	assert.Equal(t, "", token, "Expected empty token")
	assert.EqualError(t, err, "mock error", "Unexpected error message")
}
