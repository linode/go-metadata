package metadata

import (
	"context"
	"errors"
	"net/netip"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock client for testing purposes
type NetworkMockclient struct {
	Resp *NetworkData
	Err  error
}

func (m *NetworkMockclient) GetNetwork(ctx context.Context) (*NetworkData, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resp, nil
}

func TestGetNetwork_Success(t *testing.T) {
	mockClient := &NetworkMockclient{
		Resp: &NetworkData{
			Interfaces: []InterfaceData{
				{Label: "eth0", Purpose: "public", IPAMAddress: netip.MustParsePrefix("203.0.113.0/24")},
				{Label: "eth1", Purpose: "private", IPAMAddress: netip.MustParsePrefix("192.168.1.0/24")},
			},
			IPv4: IPv4Data{
				Public:  []netip.Prefix{netip.MustParsePrefix("203.0.113.0/24")},
				Private: []netip.Prefix{netip.MustParsePrefix("192.168.1.0/24")},
				Shared:  []netip.Prefix{netip.MustParsePrefix("198.51.100.0/24")},
			},
			IPv6: IPv6Data{
				SLAAC:        netip.MustParsePrefix("2001:db8::/64"),
				LinkLocal:    netip.MustParsePrefix("fe80::/64"),
				Ranges:       []netip.Prefix{netip.MustParsePrefix("2001:db8::/64")},
				SharedRanges: []netip.Prefix{netip.MustParsePrefix("fd00::/64")},
			},
		},
	}

	network, err := mockClient.GetNetwork(context.Background())

	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, network, "Expected non-nil network")
	assert.Len(t, network.Interfaces, 2, "Unexpected number of interfaces")
	assert.Len(t, network.IPv4.Public, 1, "Unexpected number of public IPv4 prefixes")
	assert.Len(t, network.IPv6.Ranges, 1, "Unexpected number of IPv6 ranges")
}

func TestGetNetwork_Error(t *testing.T) {
	mockClient := &NetworkMockclient{
		Err: errors.New("mock error"),
	}

	network, err := mockClient.GetNetwork(context.Background())

	assert.Error(t, err, "Expected an error")
	assert.Nil(t, network, "Expected nil network")
	assert.EqualError(t, err, "mock error", "Unexpected error message")
}
