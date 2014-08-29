package main

import (
	"github.com/mitchellh/packer/packer/plugin"
	"github.com/rastasheep/packer-builder-docker-commit/builder/docker_commit"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(docker_commit.Builder))
	server.Serve()

}
