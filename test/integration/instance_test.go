package integration

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstance(t *testing.T) {
	t.Parallel()

	mdInst, err := metadataClient.GetInstance(context.Background())
	assert.NoError(t, err)

	assert.NotEqual(t, t, mdInst.ID, 0)

	apiInst, err := linodeClient.GetInstance(context.Background(), mdInst.ID)
	assert.NoError(t, err)

	assert.Equal(t, apiInst.Label, mdInst.Label)
	assert.Equal(t, apiInst.Region, mdInst.Region)
	assert.Equal(t, apiInst.Type, mdInst.Type)
	assert.Equal(t, apiInst.Tags, mdInst.Tags)

	assert.Equal(t, apiInst.Specs.Disk, mdInst.Specs.Disk)
	assert.Equal(t, apiInst.Specs.Memory, mdInst.Specs.Memory)
	assert.Equal(t, apiInst.Specs.VCPUs, mdInst.Specs.VCPUs)
	assert.Equal(t, apiInst.Specs.GPUs, mdInst.Specs.GPUs)
	assert.Equal(t, apiInst.Specs.Transfer, mdInst.Specs.Transfer)
}
