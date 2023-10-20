# Linode Metadata Service Client for Go

This package allows Go projects to easily interact with the [Linode Metadata Service](https://www.linode.com/docs/products/compute/compute-instances/guides/metadata/?tabs=linode-api).

## Documentation

See [godoc](https://pkg.go.dev/github.com/linode/go-metadata) for a complete documentation reference.

## Testing

Before running tests on this project, please ensure you have a 
[Linode Personal Access Token](https://www.linode.com/docs/products/tools/api/guides/manage-api-tokens/)
exported under the `LINODE_TOKEN` environment variable.

### End-to-End Testing Using Ansible

This project contains an Ansible playbook to automatically deploy the necessary infrastructure 
and run end-to-end tests on it.

To install the dependencies for this playbook, ensure you have Python 3 installed and run the following:

```bash
make test-deps
```

After all dependencies have been installed, you can run the end-to-end test suite by running the following:

```bash
make e2e
```

If your local SSH public key is stored in a location other than `~/.ssh/id_rsa.pub`, 
you may need to override the `TEST_PUBKEY` argument:

```bash
make TEST_PUBKEY=/path/to/my/pubkey e2e
```

**NOTE: To speed up subsequent test runs, the infrastructure provisioned for testing will persist after the test run is complete. 
This infrastructure is safe to manually remove.**

### Manual End-to-End Testing

End-to-end tests can also be manually run using the `make e2e-local` target.
This test suite is expected to run from within a Linode instance and will likely 
fail in other environments.

## Contribution Guidelines

Want to improve metadata-go? Please start [here](CONTRIBUTING.md).