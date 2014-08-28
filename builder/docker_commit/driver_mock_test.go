package docker_commit

import "testing"

func TestMockDriver_impl(t *testing.T) {
	var _ Driver = new(MockDriver)
}
