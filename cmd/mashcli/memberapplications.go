package main

import (
	"fmt"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/mmussett/mashcli/cli/app/memberapplications"
	"github.com/urfave/cli"
)

func doBeforeMemberApplicationsShow(c *cli.Context) {
	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli applications show --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionMemberApplicationsShow(c *cli.Context) {

	var memberId = c.Args().Get(0)

	m, err := mashcli.Load(c.String("area"))
	if err != nil {
		fmt.Printf("unable to load area config: %v", err)
		cli.OsExiter(-1)
		return
	}

	accessToken, err := m.FetchOAuthToken()
	if err != nil {
		fmt.Printf("unable to fetch oauth token: %v", err)
		cli.OsExiter(-1)
		return
	}

	var format = "table"
	if c.IsSet("output") {
		format = c.String("output")
	}

	var filter = ""
	if c.IsSet("filter") {
		filter = c.String("filter")
	}

	err = memberapplications.ShowMemberApplications(accessToken, memberId, format, filter)
	if err != nil {
		fmt.Printf("can't show member application: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}


func doBeforeMemberApplicationsExport(c *cli.Context) {

	if len(c.Args()) < 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli memberapplications export --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionMemberApplicationsExport(c *cli.Context) {

	var memberId = c.Args().Get(0)
	var filename = ""

	if c.IsSet("filename") {
		filename = c.String("filename")
	}

	m, err := mashcli.Load(c.String("area"))
	if err != nil {
		fmt.Printf("unable to load area config: %v", err)
		cli.OsExiter(-1)
		return
	}

	accessToken, err := m.FetchOAuthToken()
	if err != nil {
		fmt.Printf("unable to fetch oauth token: %v", err)
		cli.OsExiter(-1)
		return
	}

	err = memberapplications.Export(accessToken, memberId, filename)
	if err != nil {
		fmt.Printf("can't export member application: %v", err)
		cli.OsExiter(-1)
		return
	}

}
