package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	t.Parallel()

	result, err := metadataClient.GenerateToken(context.Background())
	assert.NoError(t, err)

	assert.NotNil(t, result)
}
