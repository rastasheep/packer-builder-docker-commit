package docker_commit

import (
	"github.com/mitchellh/packer/packer"
	"testing"
)

func TestCommunicator_impl(t *testing.T) {
	var _ packer.Communicator = new(Communicator)
}
