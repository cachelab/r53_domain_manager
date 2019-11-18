package commands

import (
	"strconv"
	"strings"

	"github.com/mitchellh/cli"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53domains"
	"github.com/aws/aws-sdk-go/service/route53domains/route53domainsiface"
	"github.com/olekukonko/tablewriter"
)

type ListCommand struct {
	Ui     cli.Ui
	Client route53domainsiface.Route53DomainsAPI
}

func (c *ListCommand) Help() string {
	helpText := `
Usage: r53dm list

  This will list all the domains you have purchased using Route 53.
`

	return strings.TrimSpace(helpText)
}

func (c *ListCommand) Synopsis() string {
	return "Lists domains that have purchased"
}

func (c *ListCommand) Run(args []string) int {
	var marker *string
	var allDomains []*route53domains.DomainSummary

	more := true

	for more {
		domains, err := c.Client.ListDomains(&route53domains.ListDomainsInput{
			MaxItems: aws.Int64(20),
			Marker:   marker,
		})
		if err != nil {
			c.Ui.Output(err.Error())
			return 1
		}

		allDomains = append(allDomains, domains.Domains...)

		if domains.NextPageMarker == nil {
			more = false
		} else {
			marker = domains.NextPageMarker
		}
	}

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader([]string{"Domain", "Auto Renew", "Transfer Lock", "Expiry"})

	for _, domain := range allDomains {
		table.Append([]string{
			*domain.DomainName,
			strconv.FormatBool(*domain.AutoRenew),
			strconv.FormatBool(*domain.TransferLock),
			domain.Expiry.Format("2006-01-02"),
		})
	}

	table.Render()

	c.Ui.Output(tableString.String())

	return 0
}
