// This module stores testing dependencies that we want to keep
// separate from the main package.

module github.com/linode/go-metadata/test/integration

go 1.23.0

toolchain go1.23.7

require (
	github.com/jarcoal/httpmock v1.4.0
	github.com/linode/go-metadata v0.0.0
	github.com/linode/linodego v1.48.1
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-resty/resty/v2 v2.16.5 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.36.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	gopkg.in/ini.v1 v1.66.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/linode/go-metadata => ../../
