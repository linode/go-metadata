package integration

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

// NOTE: This test assumes the host instance has
// a user configured with SSH keys. This should
// always be the case when using the `make e2e`
// target but may not always be the case when running
// the `make e2e-local` target.
func TestGetSSHKeys(t *testing.T) {
	t.Parallel()

	sshKeys, err := metadataClient.GetSSHKeys(context.Background())
	assert.NoError(t, err)

	assert.Greater(t, len(sshKeys.Users), 0)
	for _, v := range sshKeys.Users {
		assert.Greater(t, len(v), 0)
	}
}
