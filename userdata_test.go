package metadata

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"

	"github.com/stretchr/testify/assert"
)

var mockMetadataHost = fmt.Sprintf("%s://%s/%s", APIProto, APIHost, APIVersion)

func SetupMockClient() *http.Client {
	// create mock client
	mockClient := http.DefaultClient
	httpmock.ActivateNonDefault(mockClient)

	// Mock out token request
	tokenResponder := httpmock.NewStringResponder(200, "[\"token\"]")
	httpmock.RegisterResponder("PUT", fmt.Sprintf("%s/token", mockMetadataHost), tokenResponder)
	return mockClient
}

func TestGetUserData_Success(t *testing.T) {
	mockClient := SetupMockClient()
	// Mock out user-data response with the encoded value for "mock-user-data"
	instanceResponder := httpmock.NewStringResponder(200, "bW9jay11c2VyLWRhdGE=")
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/user-data", mockMetadataHost), instanceResponder)

	newClient, err := NewClient(context.Background(), func(options *clientCreateConfig) {
		options.HTTPClient = mockClient
	})
	assert.NoError(t, err, "Expected no error")

	userData, err := newClient.GetUserData(context.Background())

	assert.NoError(t, err, "Expected no error")

	assert.Equal(t, "mock-user-data", userData, "Unexpected user data")
}

func TestGetUserDataGzip_Success(t *testing.T) {
	mockClient := SetupMockClient()
	// Mock out user-data response with the gzipped encoded value for "mock-user-data"
	instanceResponder := httpmock.NewStringResponder(200, "H4sIAO0n32YAA8vNT87WLS1OLdJNSSxJBACRtuznDgAAAA==")
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/user-data", mockMetadataHost), instanceResponder)

	newClient, err := NewClient(context.Background(), func(options *clientCreateConfig) {
		options.HTTPClient = mockClient
	})
	assert.NoError(t, err, "Expected no error")

	userData, err := newClient.GetUserData(context.Background())

	assert.NoError(t, err, "Expected no error")

	assert.Equal(t, "mock-user-data", userData, "Unexpected user data")
}

func TestGetUserDataGzip_Error(t *testing.T) {
	mockClient := SetupMockClient()
	// Mock out user-data response with the invalid gzip encoded value for "mock-user-data"
	invalidGzipData := []byte{0x1F, 0x8B, 0x08, 0x23}
	userDataResponse := base64.StdEncoding.EncodeToString(invalidGzipData)
	instanceResponder := httpmock.NewStringResponder(200, userDataResponse)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/user-data", mockMetadataHost), instanceResponder)

	newClient, err := NewClient(context.Background(), func(options *clientCreateConfig) {
		options.HTTPClient = mockClient
	})
	assert.NoError(t, err, "Expected no error")

	userData, err := newClient.GetUserData(context.Background())

	assert.EqualErrorf(t, err, "failed to ungzip user-data: unexpected EOF", "Unexpected error message")

	assert.Equal(t, "", userData, "expected Empty Userdata")
}

func TestGetUserData_Error(t *testing.T) {
	mockClient := SetupMockClient()

	instanceResponder := httpmock.NewStringResponder(500, "{\"errors\": [{\"reason\": \"failed to get metadata\"}]}")
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/user-data", mockMetadataHost), instanceResponder)
	newClient, err := NewClient(context.Background(), func(options *clientCreateConfig) {
		options.HTTPClient = mockClient
	})
	assert.NoError(t, err, "Expected no error")

	userData, err := newClient.GetUserData(context.Background())

	assert.Error(t, err, "Expected an error")
	assert.Equal(t, "", userData, "Expected empty user data")
	assert.EqualErrorf(t, err, "[500] failed to get metadata", "Unexpected error message")
}

func TestGetUserDataDecode_Error(t *testing.T) {
	mockClient := SetupMockClient()
	// Mock out user-data response with the gzipped encoded value for "mock-user-data"
	instanceResponder := httpmock.NewStringResponder(200, "invalid base64")
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/user-data", mockMetadataHost), instanceResponder)

	newClient, err := NewClient(context.Background(), func(options *clientCreateConfig) {
		options.HTTPClient = mockClient
	})
	assert.NoError(t, err, "Expected no error")

	userData, err := newClient.GetUserData(context.Background())

	assert.Error(t, err, "Expected an error")
	assert.Equal(t, "", userData, "Expected empty user data")
	assert.EqualErrorf(t, err, "failed to decode user-data: illegal base64 data at input byte 7", "Unexpected error message")
}
