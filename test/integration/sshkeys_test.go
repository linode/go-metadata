package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSSHKeys(t *testing.T) {
	t.Parallel()

	sshKeys, err := metadataClient.GetSSHKeys(context.Background())
	assert.NoError(t, err)

	if len(sshKeys.Users) < 1 {
		t.Skip(
			"The current instance does not have any any SSH keys configured, skipping...")
	}

	for _, v := range sshKeys.Users {
		assert.Greater(t, len(v), 0)
	}
}
