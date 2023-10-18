package integration

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserData(t *testing.T) {
	t.Parallel()

	_, err := metadataClient.GetUserData(context.Background())
	assert.NoError(t, err)
}
