package commands

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestUpdateCommand(t *testing.T) {
	var _ cli.Command = &UpdateCommand{}
}
