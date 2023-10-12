package metadata

import (
	"context"
	"net/netip"
)

type VLANData struct {
	Label       string       `json:"label"`
	Purpose     string       `json:"purpose"`
	IPAMAddress netip.Prefix `json:"ipam_address"`
}

type IPv4Data struct {
	Public  []netip.Prefix `json:"public"`
	Private []netip.Prefix `json:"private"`
	Shared  []netip.Prefix `json:"shared"`
}

type IPv6Data struct {
	SLAAC        netip.Prefix   `json:"slaac"`
	LinkLocal    netip.Prefix   `json:"link-local"`
	Ranges       []netip.Prefix `json:"ranges"`
	SharedRanges []netip.Prefix `json:"shared-ranges"`
}

type NetworkData struct {
	VLANs []VLANData `json:"vlans"`
	IPv4  IPv4Data   `json:"ipv4"`
	IPv6  IPv6Data   `json:"ipv6"`
}

func (c *Client) GetNetwork(ctx context.Context) (*NetworkData, error) {
	req := c.R(ctx).SetResult(&NetworkData{})

	resp, err := req.Get("network")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*NetworkData), nil
}
