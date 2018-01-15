package main

import (
	"github.com/op/go-logging"
	"github.com/urfave/cli"
	"os"
)

const (
	name      = "mashcli"
	usage     = "TIBCO Mashery - Command Line Interface"
	version   = "0.1.0"
	copyright = "2017 Mark Mussett."
	author    = "Mark Mussett (mmussett@me.com)"
)

var log = logging.MustGetLogger("mashcli")

func main() {

	a := cli.NewApp()
	a.Name = name
	a.Usage = usage
	a.Version = version
	a.Copyright = copyright
	a.Author = author


	//appConfig := mashcli.Load("")

	a.Commands = []cli.Command{
		{
			Name:  "api",
			Usage: "Manage application package keys-related operations. For additional help, use 'mashcli applicationpackagekeys --help'",
			Subcommands: []cli.Command{
				{
					Name:  "import",
					Aliases: []string{"i"},
					Usage: "Import Open-API / Swagger definition",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					ArgsUsage: "filename PublicDomain",
					Before: func(c *cli.Context) error {
						doBeforeSwaggerImport(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionSwaggerImport(c)
						return nil
					},
				},
			},
		},
		{
			Name:  "applications",
			Aliases: []string{"ap"},
			Usage: "Manage application-related operations for the current user. For additional help, use 'mashcli applications --help'",
			Subcommands: []cli.Command{
				{
					Name:      "delete",
					Aliases: []string{"d"},
					Usage:     "delete application",
					ArgsUsage: "ApplicationID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforeApplicationsDelete(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionApplicationsDelete(c)
						return nil
					},
				},
				{
					Name:  "export",
					Aliases: []string{"e"},
					Usage: "Export application",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "filename", Usage: "The export Filename for the Package Definition"},
					},
					ArgsUsage: "ApplicationID",
					Before: func(c *cli.Context) error {
						doBeforeApplicationsExport(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionApplicationsExport(c)
						return nil
					},
				},
				{
					Name:      "import",
					Aliases: []string{"i"},
					Usage:     "import application",
					ArgsUsage: "",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "filename", Usage: "The import Filename for the Package Definition"},
					},
					Before: func(c *cli.Context) error {
						doBeforeApplicationsImport(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionApplicationsImport(c)
						return nil
					},
				},
				{
					Name:      "nuke",
					Usage:     "nuke all applications",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.BoolFlag{Name: "force, f", Usage: "Ignore warnings and prompts"},
					},
					Before: func(c *cli.Context) error {
						doBeforeApplicationsNuke(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionApplicationsNuke(c)
						return nil
					},
				},
				{
					Name:  "show",
					Aliases: []string{"s"},
					Usage: "Show application",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					ArgsUsage: "ApplicationID",
					Before: func(c *cli.Context) error {
						doBeforeApplicationsShow(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionApplicationsShow(c)
						return nil
					},
				},
				{
					Name:  "showall",
					Aliases: []string{"sa"},
					Usage: "Show all applications",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					Before: func(c *cli.Context) error {
						doBeforeApplicationsShowAll(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionApplicationsShowAll(c)
						return nil
					},
				},
			},
		},
		{
			Name:  "applicationpackagekeys",
			Aliases: []string{"ak"},
			Usage: "Manage application package keys-related operations. For additional help, use 'mashcli applicationpackagekeys --help'",
			Subcommands: []cli.Command{
				{
					Name:  "showall",
					Aliases: []string{"sa"},
					Usage: "Show all package keys for an application",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					ArgsUsage: "ApplicationID",
					Before: func(c *cli.Context) error {
						doBeforeApplicationPackageKeysShowAll(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionApplicationPackageKeysShowAll(c)
						return nil
					},
				},
			},
		},
		{
			Name:  "area",
			Aliases: []string{"ar"},
			Usage: "Manage application package keys-related operations. For additional help, use 'mashcli area --help'",
			Subcommands: []cli.Command{
				{
					Name:  "backup",
					Aliases: []string{"b"},
					Usage: "Backup area configuration",
					ArgsUsage: "Name",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforeAreaBackup(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionAreaBackup(c)
						return nil
					},
				},
				{
					Name:  "restore",
					Aliases: []string{"r"},
					Usage: "Restore area configuration",
					ArgsUsage: "Name",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforeAreaRestore(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionAreaRestore(c)
						return nil
					},
				},
			},
		},

		{
			Name:  "members",
			Aliases: []string{"m"},
			Usage: "Manage member-related operations. For additional help, use 'mashcli members --help'",
			Subcommands: []cli.Command{
				{
					Name:      "add",
					Aliases: []string{"a"},
					Usage:     "add member",
					ArgsUsage: "Email Username DisplayName",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforeMembersAdd(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionMembersAdd(c)
						return nil
					},
				},
				{
					Name:      "delete",
					Aliases: []string{"d"},
					Usage:     "delete member",
					ArgsUsage: "MemberID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforeMembersDelete(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionMembersDelete(c)
						return nil
					},
				},
				{
					Name:  "export",
					Aliases: []string{"e"},
					Usage: "Export specific member",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "filename", Usage: "The export Filename for the Member Definition"},
					},
					ArgsUsage: "PackageID",
					Before: func(c *cli.Context) error {
						doBeforeMembersExport(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionMembersExport(c)
						return nil
					},
				},
				{
					Name:      "import",
					Aliases: []string{"i"},
					Usage:     "import package",
					ArgsUsage: "",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "filename", Usage: "The import Filename for the Member Definition"},
					},
					Before: func(c *cli.Context) error {
						doBeforeMembersImport(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionMembersImport(c)
						return nil
					},
				},
				{
					Name:      "nuke",
					Usage:     "nuke all members",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.BoolFlag{Name: "force, f", Usage: "Ignore warnings and prompts"},
					},
					Before: func(c *cli.Context) error {
						doBeforeMembersNuke(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionMembersNuke(c)
						return nil
					},
				},
				{
					Name:  "setstatus",
					Aliases: []string{"ss"},
					Usage: "Set member status",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					ArgsUsage: "MemberID Status",
					Before: func(c *cli.Context) error {
						doBeforeMemberSetStatus(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionMemberSetStatus(c)
						return nil
					},
				},
				{
					Name:  "show",
					Aliases: []string{"s"},
					Usage: "Show specific member",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					ArgsUsage: "MemberID",
					Before: func(c *cli.Context) error {
						doBeforeMembersShow(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionMembersShow(c)
						return nil
					},
				},
				{
					Name:  "showall",
					Aliases: []string{"sa"},
					Usage: "Show all members",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					Before: func(c *cli.Context) error {
						doBeforeMembersShowAll(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionMembersShowAll(c)
						return nil
					},
				},
			},
		},
		{
			Name:  "memberapplications",
			Aliases: []string{"ma"},
			Usage: "Manage application-related operations for the current user. For additional help, use 'mashcli memberapplications --help'",
			Subcommands: []cli.Command{
				{
					Name:  "export",
					Aliases: []string{"e"},
					Usage: "Export application",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "filename", Usage: "The export Filename for the Package Definition"},
					},
					ArgsUsage: "MemberID",
					Before: func(c *cli.Context) error {
						doBeforeMemberApplicationsExport(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionMemberApplicationsExport(c)
						return nil
					},
				},
				{
					Name:  "showall",
					Aliases: []string{"sa"},
					Usage: "Show member applications",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					ArgsUsage: "MemberID",
					Before: func(c *cli.Context) error {
						doBeforeMemberApplicationsShow(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionMemberApplicationsShow(c)
						return nil
					},
				},
			},
		},
		{
			Name:  "services",
			Aliases: []string{"s"},
			Usage: "Manage service-related operations for the current user. For additional help, use 'mashcli services --help'",
			Subcommands: []cli.Command{
				{
					Name:      "add",
					Aliases: []string{"a"},
					Usage:     "Add a service",
					ArgsUsage: "Name",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "description, d", Usage: "Describe your service"},
						cli.StringFlag{Name: "version, v", Usage: "Version identifier"},
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.IntFlag{Name: "qps, q", Usage: "Aggregate QPS"},
					},
					Before: func(c *cli.Context) error {
						doBeforeServiceAdd(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionServiceAdd(c)
						return nil
					},
				},
				{
					Name:      "clone",
					Aliases: []string{"c"},
					Usage:     "Clone a service",
					ArgsUsage: "ServiceID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforeServiceClone(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionServiceClone(c)
						return nil
					},
				},
				{
					Name:      "delete",
					Aliases: []string{"d"},
					Usage:     "Delete a service",
					ArgsUsage: "ServiceID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforeServiceDelete(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionServiceDelete(c)
						return nil
					},
				},
				{
					Name:      "export",
					Aliases: []string{"e"},
					Usage:     "Export a service",
					ArgsUsage: "ServiceID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "filename", Usage: "The export Filename for the Service Definition"},
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforeServiceExport(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionServiceExport(c)
						return nil
					},
				},
				{
					Name:  "import",
					Aliases: []string{"i"},
					Usage: "Import a service",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "filename", Usage: "The filename containing the Service Definition"},
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforeServiceImport(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionServiceImport(c)
						return nil
					},
				},
				{
					Name:      "nuke",
					Usage:     "nuke all services ",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.BoolFlag{Name: "force, f", Usage: "Ignore warnings and prompts"},
					},
					Before: func(c *cli.Context) error {
						doBeforeServiceNuke(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionServiceNuke(c)
						return nil
					},
				},
				{
					Name:      "show",
					Aliases: []string{"s"},
					Usage:     "Show specific service",
					ArgsUsage: "ID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					Before: func(c *cli.Context) error {
						doBeforeServiceShow(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionServiceShow(c)
						return nil
					},
				},
				{
					Name:  "showall",
					Aliases: []string{"sa"},
					Usage: "Show all services",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					Before: func(c *cli.Context) error {
						doBeforeServiceShowAll(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionServiceShowAll(c)
						return nil
					},
				},
			},
		},
		{
			Name:  "endpoints",
			Aliases: []string{"e"},
			Usage: "Manage endpoint-related operations for the current user. For additional help, use 'mashcli endpoints --help'",
			Subcommands: []cli.Command{
				{
					Name:      "add",
					Aliases: []string{"a"},
					Usage:     "add endpoint to a service",
					ArgsUsage: "ServiceID Name PublicEndpointAddress SystemEndpointAddress",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforeEndpointsAdd(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionEndpointsAdd(c)
						return nil
					},
				},
				{
					Name:      "clone",
					Aliases: []string{"c"},
					Usage:     "clone endpoint on a service",
					ArgsUsage: "ServiceID EndpointID Name PublicEndpointAddress SystemEndpointAddress",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforeEndpointsClone(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionEndpointsClone(c)
						return nil
					},
				},
				{
					Name:      "delete",
					Aliases: []string{"d"},
					Usage:     "delete endpoint for a service",
					ArgsUsage: "ServiceID EndpointID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforeEndpointsDelete(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionEndpointsDelete(c)
						return nil
					},
				},
				{
					Name:      "export",
					Aliases: []string{"e"},
					Usage:     "export endpoint from a service",
					ArgsUsage: "ServiceID EndpointID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "filename", Usage: "The export Filename for the Endpoint Definition"},
					},
					Before: func(c *cli.Context) error {
						doBeforeEndpointsExport(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionEndpointsExport(c)
						return nil
					},
				},
				{
					Name:      "import",
					Aliases: []string{"i"},
					Usage:     "import endpoint to a service",
					ArgsUsage: "ServiceID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "filename", Usage: "The import Filename for the Endpoint Definition"},
					},
					Before: func(c *cli.Context) error {
						doBeforeEndpointsImport(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionEndpointsImport(c)
						return nil
					},
				},
				{
					Name:      "show",
					Aliases: []string{"s"},
					Usage:     "Show endpoint for a service",
					ArgsUsage: "ServiceID EndpointID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					Before: func(c *cli.Context) error {
						doBeforeEndpointsShow(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionEndpointsShow(c)
						return nil
					},
				},
				{
					Name:      "showall",
					Aliases: []string{"sa"},
					Usage:     "Show all endpoints for a service",
					ArgsUsage: "ServiceID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					Before: func(c *cli.Context) error {
						doBeforeEndpointsShowAll(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionEndpointsShowAll(c)
						return nil
					},
				},
			},
		},
		{
			Name:  "packages",
			Aliases: []string{"pa"},
			Usage: "Manage package-related operations for the current user. For additional help, use 'mashcli packages --help'",
			Subcommands: []cli.Command{
				{
					Name:      "add",
					Aliases: []string{"a"},
					Usage:     "add package",
					ArgsUsage: "PackageName PackageDescription",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforePackageAdd(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPackageAdd(c)
						return nil
					},
				},
				{
					Name:      "clone",
					Aliases: []string{"c"},
					Usage:     "clone package",
					ArgsUsage: "PackageID Name Description",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforePackageClone(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPackageClone(c)
						return nil
					},
				},
				{
					Name:      "delete",
					Aliases: []string{"d"},
					Usage:     "delete package",
					ArgsUsage: "PackageID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforePackageDelete(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPackageDelete(c)
						return nil
					},
				},
				{
					Name:  "export",
					Aliases: []string{"e"},
					Usage: "Export specific package",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "filename", Usage: "The export Filename for the Package Definition"},
					},
					ArgsUsage: "PackageID",
					Before: func(c *cli.Context) error {
						doBeforePackageExport(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPackageExport(c)
						return nil
					},
				},
				{
					Name:      "import",
					Aliases: []string{"i"},
					Usage:     "import package",
					ArgsUsage: "",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "filename", Usage: "The import Filename for the Package Definition"},
					},
					Before: func(c *cli.Context) error {
						doBeforePackageImport(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPackageImport(c)
						return nil
					},
				},
				{
					Name:      "nuke",
					Usage:     "nuke all packages",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.BoolFlag{Name: "force, f", Usage: "Ignore warnings and prompts"},
					},
					Before: func(c *cli.Context) error {
						doBeforePackageNuke(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPackageNuke(c)
						return nil
					},
				},
				{
					Name:  "show",
					Aliases: []string{"s"},
					Usage: "Show specific package",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					ArgsUsage: "PackageID",
					Before: func(c *cli.Context) error {
						doBeforePackagesShow(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPackagesShow(c)
						return nil
					},
				},
				{
					Name:  "showall",
					Aliases: []string{"sa"},
					Usage: "Show all packages",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					Before: func(c *cli.Context) error {
						doBeforePackagesShowAll(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPackagesShowAll(c)
						return nil
					},
				},
			},
		},
		{
			Name:  "packagekeys",
			Aliases: []string{"k"},
			Usage: "Manage package key-related operations. For additional help, use 'mashcli packagekeys --help'",
			Subcommands: []cli.Command{
				{
					Name:      "delete",
					Aliases: []string{"d"},
					Usage:     "delete package key",
					ArgsUsage: "PackageKeyID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforePackageKeysDelete(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPackageKeysDelete(c)
						return nil
					},
				},
				{
					Name:  "export",
					Aliases: []string{"e"},
					Usage: "Export specific package key",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "filename", Usage: "The export Filename for the Package Key Definition"},
					},
					ArgsUsage: "PackageID",
					Before: func(c *cli.Context) error {
						doBeforePackageKeysExport(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPackageKeysExport(c)
						return nil
					},
				},
				{
					Name:      "import",
					Aliases: []string{"i"},
					Usage:     "import package key",
					ArgsUsage: "",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "filename", Usage: "The import Filename for the Package Key Definition"},
					},
					Before: func(c *cli.Context) error {
						doBeforePackageKeysImport(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPackageKeysImport(c)
						return nil
					},
				},
				{
					Name:      "nuke",
					Usage:     "nuke all package key",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.BoolFlag{Name: "force, f", Usage: "Ignore warnings and prompts"},
					},
					Before: func(c *cli.Context) error {
						doBeforePackageKeysNuke(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPackageKeysNuke(c)
						return nil
					},
				},
				{
					Name:  "setrates",
					Aliases: []string{"sr"},
					Usage: "Set package key rates",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "throttle, t", Usage: "Throttle for this key"},
						cli.StringFlag{Name: "quota, q", Usage: "Quota for this key"},
					},
					ArgsUsage: "PackageKeyID",
					Before: func(c *cli.Context) error {
						doBeforePackageKeySetRates(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPackageKeySetRates(c)
						return nil
					},
				},
				{
					Name:  "setstatus",
					Aliases: []string{"ss"},
					Usage: "Set package key status",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					ArgsUsage: "PackageKeyID Status",
					Before: func(c *cli.Context) error {
						doBeforePackageKeySetStatus(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPackageKeySetStatus(c)
						return nil
					},
				},
				{
					Name:  "show",
					Aliases: []string{"s"},
					Usage: "Show specific package key",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					ArgsUsage: "PackageKeyID",
					Before: func(c *cli.Context) error {
						doBeforePackageKeysShow(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPackageKeysShow(c)
						return nil
					},
				},
				{
					Name:  "showall",
					Aliases: []string{"sa"},
					Usage: "Show all packages keys",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					Before: func(c *cli.Context) error {
						doBeforePackageKeysShowAll(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPackageKeysShowAll(c)
						return nil
					},
				},
			},
		},
		{
			Name:  "plans",
			Aliases: []string{"pl"},
			Usage: "Manage plan-related operations for the current user. For additional help, use 'mashcli plans --help'",
			Subcommands: []cli.Command{
				{
					Name:      "add",
					Aliases: []string{"a"},
					Usage:     "add plan",
					ArgsUsage: "PackageID PlanName PlanDescription",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforePlanAdd(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPlanAdd(c)
						return nil
					},
				},
				{
					Name:      "clone",
					Aliases: []string{"c"},
					Usage:     "clone plan",
					ArgsUsage: "PackageID PlanID Name Description",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforePlanClone(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPlanClone(c)
						return nil
					},
				},
				{
					Name:      "delete",
					Aliases: []string{"d"},
					Usage:     "delete plan",
					ArgsUsage: "PackageID PlanID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforePlanDelete(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPlanDelete(c)
						return nil
					},
				},
				{
					Name:  "export",
					Aliases: []string{"e"},
					Usage: "Export specific plan",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "filename", Usage: "The export Filename for the Package Definition"},
					},
					ArgsUsage: "PackageID PlanID",
					Before: func(c *cli.Context) error {
						doBeforePlanExport(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPlanExport(c)
						return nil
					},
				},
				{
					Name:      "import",
					Aliases: []string{"i"},
					Usage:     "import plan",
					ArgsUsage: "PackageID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "filename", Usage: "The import Filename for the Package Definition"},
					},
					Before: func(c *cli.Context) error {
						doBeforePlanImport(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPlanImport(c)
						return nil
					},
				},
				{
					Name:      "nuke",
					Usage:     "nuke all plans",
					ArgsUsage: "PackageID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.BoolFlag{Name: "force, f", Usage: "Ignore warnings and prompts"},
					},
					Before: func(c *cli.Context) error {
						doBeforePlanNuke(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPlanNuke(c)
						return nil
					},
				},
				{
					Name:  "setkeyprops",
					Aliases: []string{"sk"},
					Usage: "Set plan key properties",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.IntFlag{Name: "maxkeys, mk", Usage: "Maximum allowable Keys"},
						cli.IntFlag{Name: "keysmoderated, km", Usage: "Number of Keys allowed until moderation"},
						cli.StringFlag{Name: "selfservicekeys, ssk", Usage: "Self-service Key provisioning"},
						cli.StringFlag{Name: "adminkeys, ak", Usage: "Quota override"},
					},
					ArgsUsage: "PackageID PlanID",
					Before: func(c *cli.Context) error {
						doBeforePlanSetKeyProperties(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPlanSetKeyProperties(c)
						return nil
					},
				},
				{
					Name:  "setratelimits",
					Aliases: []string{"sr"},
					Usage: "Set plan rate limits",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.IntFlag{Name: "throttle, t", Usage: "Throttle limit"},
						cli.StringFlag{Name: "throttleoverride, to", Usage: "Throttle override"},
						cli.IntFlag{Name: "quota, q", Usage: "Quota limit"},
						cli.StringFlag{Name: "quotaperiod, qp", Usage: "Quota period"},
						cli.StringFlag{Name: "quotaoverride, qo", Usage: "Quota override"},
					},
					ArgsUsage: "PackageID PlanID",
					Before: func(c *cli.Context) error {
						doBeforePlanSetRateLimits(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPlanSetRateLimits(c)
						return nil
					},
				},
				{
					Name:  "setstatus",
					Aliases: []string{"ss"},
					Usage: "Set plan status",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					ArgsUsage: "PackageID PlanID Status",
					Before: func(c *cli.Context) error {
						doBeforePlanSetStatus(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPlanSetStatus(c)
						return nil
					},
				},
				{
					Name:  "show",
					Aliases: []string{"s"},
					Usage: "Show plan",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					ArgsUsage: "PackageID PlanID",
					Before: func(c *cli.Context) error {
						doBeforePlanShow(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPlanShow(c)
						return nil
					},
				},
				{
					Name:  "showall",
					Aliases: []string{"sa"},
					Usage: "Show all package plans",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					ArgsUsage: "PackageID",
					Before: func(c *cli.Context) error {
						doBeforePlanShowAll(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPlanShowAll(c)
						return nil
					},
				},
			},
		},
		{
			Name:  "planservices",
			Aliases: []string{"ps"},
			Usage: "Manage planservice-related operations for the current user. For additional help, use 'mashcli planservices --help'",
			Subcommands: []cli.Command{
				{
					Name:  "show",
					Aliases: []string{"s"},
					Usage: "Show service on plan",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					ArgsUsage: "PackageID PlanID",
					Before: func(c *cli.Context) error {
						doBeforePlanServiceShow(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPlanServiceShow(c)
						return nil
					},
				},
				{
					Name:  "showall",
					Aliases: []string{"sa"},
					Usage: "Show all services on plan",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					ArgsUsage: "PackageID PlanID",
					Before: func(c *cli.Context) error {
						doBeforePlanServiceShowAll(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPlanServiceShowAll(c)
						return nil
					},
				},
			},
		},
		{
			Name:  "plandesigner",
			Aliases: []string{"pd"},
			Usage: "Manage plan designer operations for the current user. For additional help, use 'mashcli plandesigner --help'",
			Subcommands: []cli.Command{
				{
					Name:      "add",
					Aliases: []string{"a"},
					Usage:     "add service to plan",
					ArgsUsage: "PackageID PlanID ServiceID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforePlanDesignerAdd(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPlanDesignerAdd(c)
						return nil
					},
				},
				{
					Name:      "delete",
					Aliases: []string{"d"},
					Usage:     "delete service from plan",
					ArgsUsage: "PackageID PlanID ServiceID",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
					},
					Before: func(c *cli.Context) error {
						doBeforePlanDesignerDelete(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionPlanDesignerDelete(c)
						return nil
					},
				},
			},
		},
		{
			Name:  "config",
			Aliases: []string{"c"},
			Usage: "Configuration for the current user. For additional help, use 'mashcli config --help'",
			Subcommands: []cli.Command{
				{
					Name:  "show",
					Aliases: []string{"s"},
					Usage: "Show config",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "area, a", Usage: "Area Configuration Name"},
						cli.StringFlag{Name: "output, o", Usage: "Output format table or json)"},
					},
					Before: func(c *cli.Context) error {
						doBeforeConfigShow(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionConfigShow(c)
						return nil
					},
				},
				{
					Name:  "add",
					Aliases: []string{"a"},
					Usage: "Add a new configuration",
					Before: func(c *cli.Context) error {
						doBeforeConfigAdd(c)
						return nil
					},
					Action: func(c *cli.Context) error {
						doActionConfigAdd(c)
						return nil
					},
				},
			},
		},
	}

	a.Run(os.Args)
}
