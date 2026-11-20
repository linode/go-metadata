package integration

import (
	"context"
	"net"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func isValidIPv4(ip string) bool {
	pattern := `^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/(3[0-2]|[1-2]?[0-9])$`
	matched, _ := regexp.MatchString(pattern, ip)
	return matched
}

func isValidIPv6(ip string) bool {
	// ParseCIDR returns the IP address
	ipAddr, _, err := net.ParseCIDR(ip)
	if err != nil {
		return false
	}

	// Check if the IP address is IPv6
	if ipAddr.To4() == nil {
		return true
	}

	return false
}

func TestGetNetwork(t *testing.T) {
	t.Parallel()

	mdNet, err := metadataClient.GetNetwork(context.Background())
	require.NoError(t, err)

	ipv4List := mdNet.IPv4.Public

	for _, ip := range ipv4List {
		assert.True(t, isValidIPv4(ip.String()))
	}

	ipv6Slaac := mdNet.IPv6.SLAAC
	ipv6LinkLocal := mdNet.IPv6.LinkLocal

	assert.True(t, isValidIPv6(ipv6Slaac.String()))
	assert.True(t, isValidIPv6(ipv6LinkLocal.String()))
	assert.Contains(t, ipv6LinkLocal.String(), "fe80")

	for _, ip := range mdNet.IPv4.VPC {
		assert.NotNil(t, net.ParseIP(ip))
	}
	for _, ip := range mdNet.IPv6.VPC {
		assert.NotNil(t, net.ParseIP(ip))
	}
	for _, iface := range mdNet.Interfaces {
		assert.Positive(t, iface.ID)
		if iface.VPC == nil {
			continue
		}

		assert.Positive(t, iface.VPC.ID)
		assert.Positive(t, iface.VPC.Subnet.ID)
		if iface.VPC.Subnet.IPv4 != nil {
			assert.NotNil(t, net.ParseIP(iface.VPC.Subnet.IPv4.InstanceAddress))
		}
		if iface.VPC.Subnet.IPv6 != nil {
			for _, address := range iface.VPC.Subnet.IPv6.SLAAC {
				assert.NotNil(t, net.ParseIP(address.InstanceAddress))
			}
		}
	}
}
