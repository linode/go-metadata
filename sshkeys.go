package metadata

import "context"

type SSHKeysUserData struct {
	Root []string `json:"root"`
}

type SSHKeysData struct {
	Users SSHKeysUserData `json:"users"`
}

// GetSSHKeys gets all SSH keys for the current instance.
func (c *Client) GetSSHKeys(ctx context.Context) (*SSHKeysData, error) {
	req := c.R(ctx).SetResult(&SSHKeysData{})

	resp, err := coupleAPIErrors(req.Get("ssh-keys"))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*SSHKeysData), nil
}
