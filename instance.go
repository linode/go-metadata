package main

import (
	"context"
)

type InstanceBackupsData struct {
	Enabled bool    `json:"enabled"`
	Status  *string `json:"status"`
}

type InstanceData struct {
	LocalHostname string              `json:"local-hostname"`
	Region        string              `json:"region"`
	Type          string              `json:"type"`
	Machine       string              `json:"string"`
	ID            int                 `json:"id"`
	InstanceID    int                 `json:"instance-id"`
	CPUs          int                 `json:"cpus"`
	Memory        int                 `json:"memory"`
	Disk          int                 `json:"disk"`
	Backups       InstanceBackupsData `json:"backups"`
}

func (c *Client) GetInstance(ctx context.Context) (*InstanceData, error) {
	req := c.R(ctx).SetResult(&InstanceData{})

	resp, err := req.Get("instance")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*InstanceData), nil
}
