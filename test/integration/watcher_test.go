package integration

import (
	"context"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/linode/go-metadata"
	"github.com/stretchr/testify/assert"
)

func TestNetworkWatcher(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	httpClient := &http.Client{}
	httpmock.ActivateNonDefault(httpClient)
	defer httpmock.DeactivateAndReset()
	// since we use a hacked httpClient, we need to mock all calls we make
	httpmock.RegisterResponder("PUT", "http://169.254.169.254/v1/token", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewJsonResponse(200, []string{
			"4fa1a6d669087162e7d65b36f8750c994ce4395b3e9cccea8924466819811004",
		})
	})

	httpmock.RegisterResponder("GET", "http://169.254.169.254/v1/network",
		func(req *http.Request) (*http.Response, error) {
			randomNumber := rand.Int()
			response := map[string]any{
				"interfaces": []string{},
				"ipv4": map[string]any{
					"public":  []string{"172.233.211.141/32"},
					"private": []string{},
					"shared":  []string{},
				},
				"ipv6": map[string]any{
					"slaac":         "2600:3c06::f03c:93ff:fe98:0e4c/128",
					"ranges":        []string{},
					"link_local":    "fe80::f03c:93ff:fe98:0e4c/128",
					"shared_ranges": []string{},
				},
			}
			if randomNumber%2 == 0 {
				response["ipv4"].(map[string]any)["public"] = []string{"172.233.211.142/32"}
				return httpmock.NewJsonResponse(200, response)
			} else {
				response["ipv4"].(map[string]any)["public"] = []string{"172.233.211.141/32"}
				return httpmock.NewJsonResponse(200, response)
			}
			return httpmock.NewJsonResponse(200, response)
		})

	metadataClient, err := metadata.NewClient(ctx, metadata.ClientWithHTTPClient(httpClient))
	assert.NoError(t, err)

	watcher := metadataClient.NewNetworkWatcher(metadata.WatcherWithInterval(1 * time.Second))
	go watcher.Start(ctx)
	numUpdates := 0
	for i := 1; i <= 5; i++ {
		updateData := <-watcher.Updates
		if updateData != nil {
			t.Logf("Changed IPv4: %s", updateData.IPv4.Public[0].String())
			numUpdates += 1
		}
		time.Sleep(1 * time.Second)
	}
	assert.GreaterOrEqual(t, numUpdates, 3) // interval is 1 sec
	watcher.Close()
}

func TestInstanceWatcher(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	httpClient := &http.Client{}
	httpmock.ActivateNonDefault(httpClient)
	defer httpmock.DeactivateAndReset()
	// since we use a hacked httpClient, we need to mock all calls we make
	httpmock.RegisterResponder("PUT", "http://169.254.169.254/v1/token", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewJsonResponse(200, []string{
			"4fa1a6d669087162e7d65b36f8750c994ce4395b3e9cccea8924466819811004",
		})
	})
	httpmock.RegisterResponder("GET", "http://169.254.169.254/v1/instance",
		func(req *http.Request) (*http.Response, error) {
			randomNumber := rand.Int()
			response := map[string]any{
				"backups": map[string]any{
					"enabled": true,
					"status":  "completed",
				},
				"host_uuid": "isthisauuid",
				"id":        51438702,
				"label":     "dev-us-ord",
				"region":    "us-ord",
				"specs": map[string]int{
					"disk":     327680,
					"gpus":     0,
					"memory":   16384,
					"transfer": 6000,
					"vcpus":    8,
				},
				"type": "g6-dedicated-8",
			}
			if randomNumber%2 == 0 {
				response["label"] = "even"
				return httpmock.NewJsonResponse(200, response)
			} else {
				response["label"] = "odd"
				return httpmock.NewJsonResponse(200, response)
			}
		})

	metadataClient, err := metadata.NewClient(ctx, metadata.ClientWithHTTPClient(httpClient))
	assert.NoError(t, err)

	watcher := metadataClient.NewInstanceWatcher(metadata.WatcherWithInterval(1 * time.Second))
	go watcher.Start(ctx)
	numUpdates := 0
	for i := 1; i <= 5; i++ {
		updateData := <-watcher.Updates
		if updateData != nil {
			t.Logf("Changed Label: %s", updateData.Label)
			numUpdates += 1
		}
		time.Sleep(1 * time.Second)
	}
	assert.GreaterOrEqual(t, numUpdates, 4) // interval is 1 sec
	watcher.Close()
}
