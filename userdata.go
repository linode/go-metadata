package metadata

import (
	"context"
	"encoding/base64"
	"fmt"
)

// GetUserData returns the user data for the current instance.
// NOTE: The result of this endpoint is automatically decoded from base64.
func (c *Client) GetUserData(ctx context.Context) (string, error) {
	req := c.R(ctx)

	resp, err := coupleAPIErrors(req.Get("user-data"))
	if err != nil {
		return "", err
	}

	// user-data is returned as a raw string
	decodedBytes, err := base64.StdEncoding.DecodeString(resp.String())
	if err != nil {
		return "", fmt.Errorf("failed to decode user-data: %w", err)
	}

	return string(decodedBytes), nil
}
