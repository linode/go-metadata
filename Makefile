# Whether API requests and responses should be displayed in stderr.
LINODE_DEBUG ?= 0

# The path to the pubkey to configure the E2E testing instance with.
TEST_PUBKEY ?= ~/.ssh/id_rsa.pub

SKIP_LINT         ?= 0
SKIP_DOCKER       ?= 0

GOLANGCILINT      := golangci-lint
GOLANGCILINT_IMG  := golangci/golangci-lint:latest
GOLANGCILINT_ARGS := run

lint:
ifeq ($(SKIP_LINT), 1)
	@echo Skipping lint stage
else ifeq ($(SKIP_DOCKER), 1)
	$(GOLANGCILINT) $(GOLANGCILINT_ARGS)
else
	docker run --rm -v $(shell pwd):/app -w /app $(GOLANGCILINT_IMG) $(GOLANGCILINT) $(GOLANGCILINT_ARGS)
endif

fmt:
	gofumpt -w -l .

fix-lint: fmt
	$(GOLANGCILINT) $(GOLANGCILINT_ARGS) --fix

# Installs dependencies required to run the remote E2E suite.
test-deps:
	pip3 install --upgrade ansible -r https://raw.githubusercontent.com/linode/ansible_linode/main/requirements.txt
	ansible-galaxy collection install linode.cloud

# Runs the E2E test suite on a host provisioned by Ansible.
e2e:
	ANSIBLE_HOST_KEY_CHECKING=False ANSIBLE_STDOUT_CALLBACK=debug ansible-playbook -v --extra-vars="debug=${LINODE_DEBUG} ssh_pubkey_path=${TEST_PUBKEY}" ./hack/run-e2e.yml

# Runs the E2E test suite locally.
# NOTE: E2E tests must be run from within a Linode.
e2e-local:
	cd test/integration && make e2e-local