package metadata

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetMaintenanceEvents(t *testing.T) {
	for _, testCase := range []struct {
		name   string
		body   string
		status int
		events int
	}{
		{"event", `{"data":[{"maintenance_id":171,"event_status":"in_progress"}],"page":1,"pages":1,"results":1}`, 200, 1},
		{"empty", `{"data":[],"page":1,"pages":1,"results":0}`, 200, 0},
		{"error", `{"errors":[{"reason":"forbidden"}]}`, 403, 0},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			httpClient := &http.Client{}
			httpmock.ActivateNonDefault(httpClient)
			t.Cleanup(httpmock.DeactivateAndReset)
			httpmock.RegisterResponder("GET", mockMetadataHost+"/maintenance/events", httpmock.NewStringResponder(testCase.status, testCase.body))
			client, err := NewClient(context.Background(), ClientWithHTTPClient(httpClient), ClientWithoutManagedToken(), ClientWithToken("test-token"))
			require.NoError(t, err)
			response, err := client.GetMaintenanceEvents(context.Background())
			if testCase.status != 200 {
				require.Error(t, err)
				assert.Nil(t, response)
				return
			}
			require.NoError(t, err)
			require.Len(t, response.Data, testCase.events)
			assert.Equal(t, testCase.events, response.Results)
			if testCase.events > 0 {
				assert.Equal(t, 171, response.Data[0].MaintenanceID)
			}
		})
	}
}
