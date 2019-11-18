package commands

import (
	"flag"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53domains"
	"github.com/aws/aws-sdk-go/service/route53domains/route53domainsiface"
	"github.com/mitchellh/cli"
)

type UpdateCommand struct {
	Ui           cli.Ui
	Client       route53domainsiface.Route53DomainsAPI
	Domain       string
	AddressLine1 string
	City         string
	State        string
	Zip          string
	Email        string
	First        string
	Last         string
	Phone        string
}

func (c *UpdateCommand) Help() string {
	helpText := `
Usage: r53dm update [options]

  This will update the contact information on the domain according to what
  you pass in as options.

Options:

  --domain=example.com               The domain you wish to update.
  --addressline1="1234 Jones Drive"  The address line 1 you want to change.
  --city=Littleton                   The city you want to change.
  --state=CO                         The state you want to change.
  --zip=80127                        The zip you want to change.
  --email=hello@cachelab.co          The email you want to change.
  --first=Andrew                     The first name you want to change.
  --last=Puch                        The last name you want to change.
  --phone=Puch                       The phone number you want to change.
`

	return strings.TrimSpace(helpText)
}

func (c *UpdateCommand) Synopsis() string {
	return "Updates domain contact information"
}

func (c *UpdateCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("agent", flag.ContinueOnError)
	cmdFlags.StringVar(&c.Domain, "domain", "example.com", "The domain you want to update")
	cmdFlags.StringVar(&c.AddressLine1, "addressline1", "1234 Jones Drive", "The address line 1 you want to change")
	cmdFlags.StringVar(&c.City, "city", "CacheLab", "The city you want to change")
	cmdFlags.StringVar(&c.State, "state", "CacheLab", "The state you want to change")
	cmdFlags.StringVar(&c.Zip, "zip", "12345", "The zipcode you want to change")
	cmdFlags.StringVar(&c.Email, "email", "hello@example.com", "The email you want to change")
	cmdFlags.StringVar(&c.First, "first", "CacheLab", "The first name you want to change")
	cmdFlags.StringVar(&c.Last, "last", "CacheLab", "The last name you want to change")
	cmdFlags.StringVar(&c.Phone, "phone", "+1.4444444444", "The phone number you want to change")
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }

	if err := cmdFlags.Parse(args); err != nil {
		c.Ui.Output(err.Error())
		return 1
	}

	if c.Domain == "example.com" {
		c.Ui.Output("\nInvalid domain\n")
		cmdFlags.Usage()
		return 1
	}

	if c.AddressLine1 == "1234 Jones Drive" {
		c.Ui.Output("\nInvalid address line 1\n")
		cmdFlags.Usage()
		return 1
	}

	if c.City == "CacheLab" {
		c.Ui.Output("\nInvalid city\n")
		cmdFlags.Usage()
		return 1
	}

	if c.State == "CacheLab" {
		c.Ui.Output("\nInvalid state\n")
		cmdFlags.Usage()
		return 1
	}

	if c.Zip == "12345" {
		c.Ui.Output("\nInvalid zip\n")
		cmdFlags.Usage()
		return 1
	}

	if c.Email == "hello@example.com" {
		c.Ui.Output("\nInvalid email\n")
		cmdFlags.Usage()
		return 1
	}

	if c.First == "CacheLab" {
		c.Ui.Output("\nInvalid first\n")
		cmdFlags.Usage()
		return 1
	}

	if c.Last == "CacheLab" {
		c.Ui.Output("\nInvalid last\n")
		cmdFlags.Usage()
		return 1
	}

	if c.Phone == "+1.4444444444" {
		c.Ui.Output("\nInvalid phone\n")
		cmdFlags.Usage()
		return 1
	}

	contactDetail := &route53domains.ContactDetail{
		AddressLine1: aws.String(c.AddressLine1),
		City:         aws.String(c.City),
		Email:        aws.String(c.Email),
		FirstName:    aws.String(c.First),
		LastName:     aws.String(c.Last),
		PhoneNumber:  aws.String(c.Phone),
		State:        aws.String(c.State),
		ZipCode:      aws.String(c.Zip),
	}

	_, err := c.Client.UpdateDomainContact(&route53domains.UpdateDomainContactInput{
		DomainName:        aws.String(c.Domain),
		AdminContact:      contactDetail,
		RegistrantContact: contactDetail,
		TechContact:       contactDetail,
	})
	if err != nil {
		c.Ui.Output(err.Error())
		return 1
	}

	return 0
}
