# Whether API requests and responses should be displayed in stderr
LINODE_DEBUG ?= 0

# Installs dependencies required to run the remote E2E suite
test-deps:
	pip3 install ansible -r pip install -r https://raw.githubusercontent.com/linode/ansible_linode/main/requirements.txt
	ansible-galaxy collection install linode.cloud

# Runs the E2E test suite on a host provisioned by Ansible.
e2e:
	ANSIBLE_HOST_KEY_CHECKING=False ANSIBLE_STDOUT_CALLBACK=debug ansible-playbook -v --extra-vars="debug=${LINODE_DEBUG}" ./hack/run-e2e.yml

# Runs the E2E test suite locally.
# NOTE: E2E tests must be run from within a Linode.
e2e-local:
	cd test/integration && make e2e-local