package commands

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestListCommand(t *testing.T) {
	var _ cli.Command = &ListCommand{}
}
