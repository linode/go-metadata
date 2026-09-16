// This module stores testing dependencies that we want to keep
// separate from the main package.

module github.com/linode/go-metadata/test/integration

go 1.25.0

require (
	github.com/jarcoal/httpmock v1.4.2
	github.com/linode/go-metadata v0.0.0
	github.com/linode/linodego v1.69.1
	github.com/stretchr/testify v1.12.1
)

require (
	github.com/go-resty/resty/v2 v2.17.2 // indirect
	github.com/google/go-querystring v1.2.0 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/net v0.55.0 // indirect
	golang.org/x/text v0.37.0 // indirect
	gopkg.in/ini.v1 v1.67.2 // indirect
)

replace github.com/linode/go-metadata => ../../
