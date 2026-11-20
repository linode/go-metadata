package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetMaintenanceEvents(t *testing.T) {
	t.Parallel()

	events, err := metadataClient.GetMaintenanceEvents(context.Background())
	require.NoError(t, err)
	require.NotNil(t, events)

	assert.GreaterOrEqual(t, events.Page, 1)
	assert.GreaterOrEqual(t, events.Pages, 1)
	assert.GreaterOrEqual(t, events.Results, len(events.Data))

	for _, event := range events.Data {
		assert.Positive(t, event.MaintenanceID)
		assert.NotEmpty(t, event.EventStatus)
	}
}
