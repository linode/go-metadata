package integration

import (
	"context"
	"testing"
	"time"

	"github.com/linode/go-metadata"
	"github.com/stretchr/testify/assert"
)

func TestClient_UnmanagedTokenExpired(t *testing.T) {
	mdClient, err := metadata.NewClient(
		context.Background(),
		metadata.ClientWithoutManagedToken(),
	)
	assert.NoError(t, err)

	_, err = mdClient.RefreshToken(
		context.Background(),
		metadata.TokenWithExpiry(1),
	)
	assert.NoError(t, err)

	// Hack to wait for token expiry
	time.Sleep(time.Second * 2)

	// We expect this to fail because the token has expired
	_, err = mdClient.GetInstance(context.Background())
	assert.Error(t, err)
}

func TestClient_ManagedTokenRefresh(t *testing.T) {
	mdClient, err := metadata.NewClient(context.Background(), metadata.ClientWithManagedToken(
		metadata.TokenWithExpiry(1),
	))
	assert.NoError(t, err)

	// Hack to wait for token expiry
	time.Sleep(time.Second * 2)

	// Token should have automatically refreshed
	_, err = mdClient.GetInstance(context.Background())
	assert.NoError(t, err)
}

func TestClient_DebugLogs(t *testing.T) {
	var logger testLogger

	mdClient, err := metadata.NewClient(
		context.Background(),
		metadata.ClientWithDebug(),
		metadata.ClientWithLogger(&logger),
	)
	assert.NoError(t, err)

	_, err = mdClient.GetInstance(context.Background())
	assert.NoError(t, err)

	debugOutput := logger.Data.String()

	// Validate a few arbitrary bits of the debug output
	assert.Contains(t, debugOutput, "User-Agent: go-metadata")
	assert.Contains(t, debugOutput, "/v1/token")
	assert.Contains(t, debugOutput, "/v1/instance")
}
