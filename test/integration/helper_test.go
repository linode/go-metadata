package integration

import (
	"context"
	"log"
	"os"

	"github.com/linode/go-metadata"
	"github.com/linode/linodego"
)

var (
	testToken      = os.Getenv("LINODE_TOKEN")
	metadataClient *metadata.Client
	linodeClient   *linodego.Client
)

func init() {
	if testToken == "" {
		log.Fatal("LINODE_TOKEN must be specified to run the E2E test suite")
	}

	mdsClient, err := metadata.NewClient(
		context.Background(),
	)
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}

	apiClient := linodego.NewClient(nil)
	apiClient.SetToken(testToken)

	metadataClient = mdsClient
	linodeClient = &apiClient
}
