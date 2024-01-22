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

	// Create a new instance watcher
	instanceWatcher := client.NewInstanceWatcher(
		metadata.WatcherWithInterval(10 * time.Second),
	)

	// Start the network watcher in a goroutine.
	go instanceWatcher.Start(ctx)

	// Wait for changes
	for {
		select {
		case data := <-instanceWatcher.Updates:
			log.Printf(
				"Change to instance detected.\nNew data: %v\n",
				data,
			)
		case err := <-instanceWatcher.Errors:
			log.Fatalf("Got error from instance watcher: %s", err)
		}
	}
}
