package main

import (
	"context"
	"log"
	"time"

	metadata "github.com/linode/go-metadata"
)

func main() {
	ctx := context.Background()

	// Create a new client
	client, err := metadata.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new network watcher
	networkWatcher := client.NewNetworkWatcher(
		metadata.WatcherWithInterval(10 * time.Second),
	)

	// Start the network watcher in a goroutine.
	go networkWatcher.Start(ctx)

	// Wait for changes
	for {
		select {
		case data := <-networkWatcher.Updates:
			log.Printf(
				"Change to network configuration detected.\nNew data: %v\n",
				data,
			)
		case err := <-networkWatcher.Errors:
			log.Fatalf("Got error from network watcher: %s", err)
		}
	}
}
