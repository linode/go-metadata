package metadata

import (
	"context"
	"math"
	"strconv"
	"time"
)

type GenerateTokenOptions struct {
	ExpirySeconds int
}

type TokenData struct {
	Token         string `json:"token"`
	ExpirySeconds int
	Created       time.Time
}

func (t *TokenData) IsExpired() bool {
	return int(math.Ceil(time.Since(t.Created).Seconds())) > t.ExpirySeconds
}

func (c *Client) GenerateToken(ctx context.Context, opts GenerateTokenOptions) (*TokenData, error) {
	// Temporary override so things don't break
	req := c.R(ctx).
		ExpectContentType("text/plain").
		SetHeader("Content-Type", "text/plain")

	tokenExpirySeconds := 3600
	if opts.ExpirySeconds != 0 {
		tokenExpirySeconds = opts.ExpirySeconds
	}

	req.SetHeader("Metadata-Token-Expiry-Seconds", strconv.Itoa(tokenExpirySeconds))

	resp, err := req.Put("token")
	if err != nil {
		return nil, err
	}

	return &TokenData{
		Token:         resp.String(),
		ExpirySeconds: tokenExpirySeconds,
		Created:       time.Now(),
	}, nil
}
