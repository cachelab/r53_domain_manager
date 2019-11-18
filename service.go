package main

import (
	"fmt"
	"os"

	"r53_domain_manager/commands"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53domains"
	"github.com/mitchellh/cli"
)

const version = "1.0.0"

type Service struct{}

func (svc *Service) init() {
	client := route53domains.New(session.New())

	ui := &cli.PrefixedUi{
		AskPrefix:    "",
		OutputPrefix: "",
		InfoPrefix:   "",
		ErrorPrefix:  "",
		Ui: &cli.BasicUi{
			Writer: os.Stdout,
			Reader: os.Stdin,
		},
	}

	c := cli.NewCLI("app", version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return &commands.ListCommand{
				Ui:     ui,
				Client: client,
			}, nil
		},
		"describe": func() (cli.Command, error) {
			return &commands.DescribeCommand{
				Ui:     ui,
				Client: client,
			}, nil
		},
		"update": func() (cli.Command, error) {
			return &commands.UpdateCommand{
				Ui:     ui,
				Client: client,
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(exitStatus)
}
