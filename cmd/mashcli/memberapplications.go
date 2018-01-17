package main

import (
	"fmt"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/mmussett/mashcli/cli/app/memberapplications"
	"github.com/urfave/cli"
)

func doBeforeMemberApplicationsAdd(c *cli.Context) {

	if len(c.Args()) != 3 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli memberapplications add --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionMemberApplicationsAdd(c *cli.Context) {

	var  name, description, username  = "", "", ""

	username = c.Args().Get(0)
	name = c.Args().Get(1)
	description = c.Args().Get(2)


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

	err = memberapplications.AddMemberApplication(accessToken, username, name, description)
	if err != nil {
		fmt.Printf("can't add member application: %v", err)
		cli.OsExiter(-1)
		return
	}

	return
}


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


func doBeforeMemberApplicationsImport(c *cli.Context) {
}

func doActionMemberApplicationsImport(c *cli.Context) {

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

	_, err = memberapplications.Import(accessToken, memberId, filename)
	if err != nil {
		fmt.Printf("can't import memberapplication: %v", err)
		cli.OsExiter(-1)
		return
	}
}