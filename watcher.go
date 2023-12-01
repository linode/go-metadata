package metadata

import (
	"context"
	"reflect"
	"time"
)

const DefaultWatcherInterval = 5 * time.Minute

type NetworkWatcher struct {
	Updates  chan *NetworkData
	Errors   chan error
	cancel   chan struct{}
	client   *Client
	interval time.Duration
	ticker   *time.Ticker
}

func (watcher *NetworkWatcher) Start(ctx context.Context) {
	go func() {
		var oldNetworkData *NetworkData
		watcher.ticker = time.NewTicker(watcher.interval)
		defer watcher.ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-watcher.cancel:
				return
			case <-watcher.ticker.C:
				networkData, err := watcher.client.GetNetwork(ctx)
				if err != nil {
					watcher.Errors <- err
				}
				if !reflect.DeepEqual(networkData, oldNetworkData) {
					watcher.Updates <- networkData
					oldNetworkData = networkData
				}
			}
		}
	}()
}

func (watcher *NetworkWatcher) Close() error {
	close(watcher.cancel)
	close(watcher.Errors)
	close(watcher.Updates)
	watcher.ticker.Stop()
	return nil
}

type InstanceWatcher struct {
	Updates  chan *InstanceData
	Errors   chan error
	cancel   chan struct{}
	client   *Client
	interval time.Duration
	ticker   *time.Ticker
}

func (watcher *InstanceWatcher) Start(ctx context.Context) {
	go func() {
		var oldInstanceData *InstanceData
		watcher.ticker = time.NewTicker(watcher.interval)
		defer watcher.ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-watcher.cancel:
				return
			case <-watcher.ticker.C:
				instanceData, err := watcher.client.GetInstance(ctx)
				if err != nil {
					watcher.Errors <- err
				}
				if !reflect.DeepEqual(instanceData, oldInstanceData) { // Todo Testing
					watcher.Updates <- instanceData
					oldInstanceData = instanceData
				}
			}
		}
	}()
}

func (watcher *InstanceWatcher) Close() error {
	close(watcher.cancel)
	close(watcher.Errors)
	close(watcher.Updates)
	watcher.ticker.Stop()
	return nil
}

type WatcherOption func(options *watcherConfig)

type watcherConfig struct {
	Interval time.Duration
}

func (c *Client) NewInstanceWatcher(opts ...WatcherOption) *InstanceWatcher {
	watcherOpts := watcherConfig{
		Interval: DefaultWatcherInterval,
	}

	for _, opt := range opts {
		opt(&watcherOpts)
	}

	return &InstanceWatcher{
		Updates:  make(chan *InstanceData),
		Errors:   make(chan error),
		cancel:   make(chan struct{}),
		interval: watcherOpts.Interval,
		client:   c,
	}
}

func (c *Client) NewNetworkWatcher(opts ...WatcherOption) *NetworkWatcher {
	watcherOpts := watcherConfig{
		Interval: DefaultWatcherInterval,
	}

	for _, opt := range opts {
		opt(&watcherOpts)
	}

	return &NetworkWatcher{
		Updates:  make(chan *NetworkData),
		Errors:   make(chan error),
		cancel:   make(chan struct{}),
		interval: watcherOpts.Interval,
		client:   c,
	}
}

func WatcherWithInterval(duration time.Duration) WatcherOption {
	return func(options *watcherConfig) {
		options.Interval = duration
	}
}
