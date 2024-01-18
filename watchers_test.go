package metadata

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWatcherWithInterval(t *testing.T) {
	config := watcherConfig{}
	WatcherWithInterval(10 * time.Minute)(&config)

	assert.Equal(t, 10*time.Minute, config.Interval, "Unexpected interval duration")
}
