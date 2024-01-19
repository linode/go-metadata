package metadata

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SshkeysMockclient struct {
	Resp *SSHKeysData
	Err  error
}

func (m *SshkeysMockclient) GetSSHKeys(ctx context.Context) (*SSHKeysData, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resp, nil
}

func TestGetSSHKeys_Success(t *testing.T) {
	// Create a mock client with a successful response
	mockClient := &SshkeysMockclient{
		Resp: &SSHKeysData{
			Users: SSHKeysUserData{
				Root: []string{"ssh-randomkeyforunittestas;ldkjfqweeru", "ssh-randomkeyforunittestas;ldkjfqweerutwo"},
			},
		},
	}

	sshKeys, err := mockClient.GetSSHKeys(context.Background())

	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, sshKeys, "Expected non-nil SSHKeysData")
	assert.Len(t, sshKeys.Users.Root, 2, "Unexpected number of root SSH keys")
}

func TestGetSSHKeys_Error(t *testing.T) {
	// Create a mock client with an error response
	mockClient := &SshkeysMockclient{
		Err: errors.New("mock error"),
	}

	// Call the GetSSHKeys method
	sshKeys, err := mockClient.GetSSHKeys(context.Background())

	// Assert the result
	assert.Error(t, err, "Expected an error")
	assert.Nil(t, sshKeys, "Expected nil SSHKeysData")
	assert.EqualError(t, err, "mock error", "Unexpected error message")
}
