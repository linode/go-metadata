package integration

import (
	"context"
	"net"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)

	ipv4List := mdNet.IPv4.Public

	for _, ip := range ipv4List {
		assert.True(t, isValidIPv4(ip.String()))
	}

	ipv6Slaac := mdNet.IPv6.SLAAC
	ipv6LinkLocal := mdNet.IPv6.LinkLocal

	assert.True(t, isValidIPv6(ipv6Slaac.String()))
	assert.True(t, isValidIPv6(ipv6LinkLocal.String()))
	assert.Contains(t, ipv6LinkLocal.String(), "fe80")
}
