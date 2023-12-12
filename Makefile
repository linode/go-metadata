# Whether API requests and responses should be displayed in stderr.
LINODE_DEBUG ?= 0

# The path to the pubkey to configure the E2E testing instance with.
TEST_PUBKEY ?= ~/.ssh/id_rsa.pub

# Whether to cleanup the Linode instance used in the testing
CLEANUP_TEST_LINODE_INSTANCE ?= false

# Installs dependencies required to run the remote E2E suite.
test-deps:
	pip3 install --upgrade ansible -r https://raw.githubusercontent.com/linode/ansible_linode/main/requirements.txt
	ansible-galaxy collection install linode.cloud

# Runs the E2E test suite on a host provisioned by Ansible.
e2e:
	ANSIBLE_HOST_KEY_CHECKING=False ANSIBLE_STDOUT_CALLBACK=debug ansible-playbook -v --extra-vars="debug=${LINODE_DEBUG} ssh_pubkey_path=${TEST_PUBKEY} cleanup_linode=${CLEANUP_TEST_LINODE_INSTANCE}" ./hack/run-e2e.yml

# Runs the E2E test suite locally.
# NOTE: E2E tests must be run from within a Linode.
e2e-local:
	cd test/integration && make e2e-local
