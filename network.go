package metadata

import (
	"context"
	"net/netip"
)

type InterfaceVPCIPv4Data struct {
	InstanceAddress string   `json:"instance_address"`
	InstanceRanges  []string `json:"instance_ranges"`
	NAT1To1         *string  `json:"nat_1_1"`
	SubnetRange     string   `json:"subnet_range"`
}

type InterfaceVPCIPv6SLAACData struct {
	InstanceAddress string `json:"instance_address"`
	InstanceRange   string `json:"instance_range"`
}

type InterfaceVPCIPv6Data struct {
	InstanceRanges []string                    `json:"instance_ranges"`
	IsPublic       bool                        `json:"is_public"`
	SLAAC          []InterfaceVPCIPv6SLAACData `json:"slaac"`
	SubnetRange    string                      `json:"subnet_range"`
}

type InterfaceVPCSubnetData struct {
	ID    int                   `json:"id"`
	Label string                `json:"label"`
	IPv4  *InterfaceVPCIPv4Data `json:"ipv4"`
	IPv6  *InterfaceVPCIPv6Data `json:"ipv6"`
}

type InterfaceVPCData struct {
	ID     int                    `json:"id"`
	Label  string                 `json:"label"`
	Subnet InterfaceVPCSubnetData `json:"subnet"`
}

type InterfaceData struct {
	ID          int               `json:"id"`
	Label       string            `json:"label"`
	Purpose     string            `json:"purpose"`
	IPAMAddress netip.Prefix      `json:"ipam_address"`
	VPC         *InterfaceVPCData `json:"vpc"`
}

type IPv4Data struct {
	Public  []netip.Prefix `json:"public"`
	Private []netip.Prefix `json:"private"`
	Shared  []netip.Prefix `json:"shared"`
	VPC     []string       `json:"vpc"`
}

type IPv6Data struct {
	SLAAC        netip.Prefix   `json:"slaac"`
	LinkLocal    netip.Prefix   `json:"link_local"`
	Ranges       []netip.Prefix `json:"ranges"`
	SharedRanges []netip.Prefix `json:"shared_ranges"`
	VPC          []string       `json:"vpc"`
}

type NetworkData struct {
	Interfaces []InterfaceData `json:"interfaces"`
	IPv4       IPv4Data        `json:"ipv4"`
	IPv6       IPv6Data        `json:"ipv6"`
}

// GetNetwork gets networking information about the current Linode instance.
func (c *Client) GetNetwork(ctx context.Context) (*NetworkData, error) {
	req := c.R(ctx).SetResult(&NetworkData{})

	resp, err := coupleAPIErrors(req.Get("network"))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*NetworkData), nil
}
