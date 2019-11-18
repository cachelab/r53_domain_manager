package commands

import (
	"flag"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53domains"
	"github.com/aws/aws-sdk-go/service/route53domains/route53domainsiface"
	"github.com/mitchellh/cli"
	"github.com/olekukonko/tablewriter"
)

type DescribeCommand struct {
	Ui     cli.Ui
	Client route53domainsiface.Route53DomainsAPI
	Domain string
}

func (c *DescribeCommand) Help() string {
	helpText := `
Usage: r53dm describe [options]

  This will get more detailed information on the domain according to what
  you pass in as options.

Options:

  --domain=example.com   The domain you wish to get more information for.
`

	return strings.TrimSpace(helpText)
}

func (c *DescribeCommand) Synopsis() string {
	return "Describes domain passed in"
}

func (c *DescribeCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("agent", flag.ContinueOnError)
	cmdFlags.StringVar(&c.Domain, "domain", "example.com", "The domain you want to describe")
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

	domain, err := c.Client.GetDomainDetail(&route53domains.GetDomainDetailInput{
		DomainName: aws.String(c.Domain),
	})
	if err != nil {
		c.Ui.Output(err.Error())
		return 1
	}

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader([]string{"Contact Type", "Address Line 1", "City", "State", "Zip", "Email", "First", "Last", "Phone"})

	table.Append([]string{
		"Admin",
		*domain.AdminContact.AddressLine1,
		*domain.AdminContact.City,
		*domain.AdminContact.State,
		*domain.AdminContact.ZipCode,
		*domain.AdminContact.Email,
		*domain.AdminContact.FirstName,
		*domain.AdminContact.LastName,
		*domain.AdminContact.PhoneNumber,
	})

	table.Append([]string{
		"Registrant",
		*domain.RegistrantContact.AddressLine1,
		*domain.RegistrantContact.City,
		*domain.RegistrantContact.State,
		*domain.RegistrantContact.ZipCode,
		*domain.RegistrantContact.Email,
		*domain.RegistrantContact.FirstName,
		*domain.RegistrantContact.LastName,
		*domain.RegistrantContact.PhoneNumber,
	})

	table.Append([]string{
		"Admin",
		*domain.TechContact.AddressLine1,
		*domain.TechContact.City,
		*domain.TechContact.State,
		*domain.TechContact.ZipCode,
		*domain.TechContact.Email,
		*domain.TechContact.FirstName,
		*domain.TechContact.LastName,
		*domain.TechContact.PhoneNumber,
	})

	table.Render()

	c.Ui.Output(tableString.String())

	return 0
}
