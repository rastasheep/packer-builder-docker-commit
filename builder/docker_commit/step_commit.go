package docker_commit

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

// StepCommit commits the container to a image.
type StepCommit struct {
	imageId string
}

func (s *StepCommit) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*Config)
	driver := state.Get("driver").(Driver)
	containerId := state.Get("container_id").(string)
	ui := state.Get("ui").(packer.Ui)

	if config.Export {
		return multistep.ActionContinue
	}

	ui.Say("Committing the container")
	imageId, err := driver.Commit(containerId, config.CommitTag)
	if err != nil {
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	// Save the container ID
	s.imageId = imageId
	state.Put("image_id", s.imageId)
	ui.Message(fmt.Sprintf("Image ID: %s", s.imageId))

	return multistep.ActionContinue
}

func (s *StepCommit) Cleanup(state multistep.StateBag) {}
