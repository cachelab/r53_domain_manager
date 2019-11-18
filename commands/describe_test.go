package commands

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestDescribeCommand(t *testing.T) {
	var _ cli.Command = &DescribeCommand{}
}
