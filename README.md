# Linode Metadata Service Client for Go

This package allows Go projects to easily interact with the [Linode Metadata Service](https://www.linode.com/docs/products/compute/compute-instances/guides/metadata/?tabs=linode-api).

## Getting Started

### Prerequisites 

- Go >= 1.20
- A running [Linode Instance](https://www.linode.com/docs/api/linode-instances/)

### Installation

```bash
go get github.com/linode/go-metadata
```

### Basic Example

The follow sample shows a simple Go project that initializes a new metadata client and retrieves various information
about the current Linode.

```go
package main

import (
	"context"
	"fmt"
	"log"

	metadata "github.com/linode/go-metadata"
)

func main() {
	// Create a new client
	client, err := metadata.NewClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve metadata about the current instance from the metadata API
	instanceInfo, err := client.GetInstance(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Instance Label:", instanceInfo.Label)
}
```

### Without Token Management

By default, metadata API tokens are automatically generated and refreshed without any user intervention.
If you would like to manage API tokens yourself, this functionality can be disabled:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	metadata "github.com/linode/go-metadata"
)

func main() {
	// Get a token from the environment
	token := os.Getenv("LINODE_METADATA_TOKEN")

	// Create a new client
	client, err := metadata.NewClient(
		context.Background(), 
		metadata.ClientWithoutManagedToken(), 
		metadata.ClientWithToken(token),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve metadata about the current instance from the metadata API
	instanceInfo, err := client.GetInstance(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Instance Label:", instanceInfo.Label)
}
```

## Contribution Guidelines

Want to improve metadata-go? Please start [here](CONTRIBUTING.md).