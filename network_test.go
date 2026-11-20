package metadata

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/netip"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestGetNetworkVPCResponse(t *testing.T) {
	payload := `{
		"interfaces": [{"id":101,"label":null,"purpose":"vpc","ipam_address":null,"vpc":{
			"id":202,"label":"vpc-label","subnet":{"id":303,"label":"subnet-label",
			"ipv4":{"instance_address":"10.0.0.2","instance_ranges":["10.0.1.0/24"],"nat_1_1":"192.0.2.2","subnet_range":"10.0.0.0/16"},
			"ipv6":{"instance_ranges":["2001:db8:1::/64"],"is_public":false,"slaac":[{"instance_address":"2001:db8::2","instance_range":"2001:db8::/64"}],"subnet_range":"2001:db8::/48"}
		}}}],
		"ipv4":{"public":[],"private":[],"shared":[],"vpc":["10.0.0.2"]},
		"ipv6":{"slaac":null,"link_local":null,"ranges":[],"shared_ranges":[],"vpc":["2001:db8::2"]}
	}`
	httpClient := &http.Client{}
	httpmock.ActivateNonDefault(httpClient)
	t.Cleanup(httpmock.DeactivateAndReset)
	httpmock.RegisterResponder("GET", mockMetadataHost+"/network", httpmock.NewStringResponder(200, payload))
	client, err := NewClient(context.Background(), ClientWithHTTPClient(httpClient), ClientWithoutManagedToken(), ClientWithToken("test-token"))
	require.NoError(t, err)
	response, err := client.GetNetwork(context.Background())
	require.NoError(t, err)
	require.Len(t, response.Interfaces, 1)
	iface := response.Interfaces[0]
	assert.Equal(t, 101, iface.ID)
	require.NotNil(t, iface.VPC)
	subnet := iface.VPC.Subnet
	assert.Equal(t, 303, subnet.ID)
	require.NotNil(t, subnet.IPv4)
	assert.Equal(t, "10.0.0.2", subnet.IPv4.InstanceAddress)
	require.NotNil(t, subnet.IPv6)
	require.Len(t, subnet.IPv6.SLAAC, 1)
	assert.Equal(t, "2001:db8::2", subnet.IPv6.SLAAC[0].InstanceAddress)
	assert.Equal(t, []string{"10.0.0.2"}, response.IPv4.VPC)
	assert.Equal(t, []string{"2001:db8::2"}, response.IPv6.VPC)
}

func TestNetworkVPCNulls(t *testing.T) {
	var iface InterfaceData
	require.NoError(t, json.Unmarshal([]byte(`{"vpc":null}`), &iface))
	assert.Nil(t, iface.VPC)
	var subnet InterfaceVPCSubnetData
	require.NoError(t, json.Unmarshal([]byte(`{"ipv4":null,"ipv6":null}`), &subnet))
	assert.Nil(t, subnet.IPv4)
	assert.Nil(t, subnet.IPv6)
	var ipv4 InterfaceVPCIPv4Data
	require.NoError(t, json.Unmarshal([]byte(`{"nat_1_1":null}`), &ipv4))
	assert.Nil(t, ipv4.NAT1To1)
}
