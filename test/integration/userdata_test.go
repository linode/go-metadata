package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserData(t *testing.T) {
	t.Parallel()

	_, err := metadataClient.GetUserData(context.Background())
	assert.NoError(t, err)
}
