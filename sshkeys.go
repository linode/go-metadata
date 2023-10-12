package metadata

import "context"

type SSHKeysUserData struct {
	Root []string `json:"root"`
}

type SSHKeysData struct {
	Users SSHKeysUserData `json:"users"`
}

func (c *Client) GetSSHKeys(ctx context.Context) (*SSHKeysData, error) {
	req := c.R(ctx).SetResult(&SSHKeysData{})

	resp, err := req.Get("ssh-keys")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*SSHKeysData), nil
}
