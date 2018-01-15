package main

import (
	"fmt"
	"strings"

	"github.com/Songmu/prompter"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/mmussett/mashcli/cli/app/members"
	"github.com/urfave/cli"
	"github.com/fatih/color"
)

func doBeforeMemberSetStatus(c *cli.Context) {
	if len(c.Args()) != 2 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli members setstatus --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionMemberSetStatus(c *cli.Context) {

	var memberId = c.Args().Get(0)
	var status = c.Args().Get(1)

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

	switch strings.ToLower(status) {
	case "active":
		status = "active"
	case "disabled":
		status = "disabled"
	default:
		fmt.Println("status must be 'active' or 'disabled'")
		cli.OsExiter(-1)
		return
	}

	err = members.SetStatus(accessToken, memberId, status)
	if err != nil {
		fmt.Printf("can't set status for member: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforeMembersShow(c *cli.Context) {
	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli members show --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionMembersShow(c *cli.Context) {

	// Arg0 = Package ID
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

	err = members.ShowMember(accessToken, memberId,format)
	if err != nil {
		fmt.Printf("can't show members: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforeMembersShowAll(c *cli.Context) {
}

func doActionMembersShowAll(c *cli.Context) {

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

	err = members.ShowAllMembers(accessToken,format,filter)
	if err != nil {
		fmt.Printf("can't show all members: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforeMembersAdd(c *cli.Context) {

	if len(c.Args()) != 3 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli members add --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionMembersAdd(c *cli.Context) {

	email := c.Args().Get(0)
	username := c.Args().Get(1)
	displayname := c.Args().Get(2)

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

	err = members.AddMember(accessToken,email,username,displayname)
	if err != nil {
		fmt.Printf("can't add member: %v", err)
		cli.OsExiter(-1)
		return
	}

	return
}


func doBeforeMembersExport(c *cli.Context) {

	if len(c.Args()) < 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli members export --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionMembersExport(c *cli.Context) {

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

	err = members.Export(accessToken, memberId, filename)
	if err != nil {
		fmt.Printf("can't export members: %v", err)
		cli.OsExiter(-1)
		return
	}

}

func doBeforeMembersImport(c *cli.Context) {
}

func doActionMembersImport(c *cli.Context) {

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

	_, err = members.Import(accessToken, filename)
	if err != nil {
		fmt.Printf("can't import members: %v", err)
		cli.OsExiter(-1)
		return
	}
}


func doBeforeMembersDelete(c *cli.Context) {

	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli members delete --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionMembersDelete(c *cli.Context) {

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

	err = members.DeleteMember(accessToken,memberId)
	if err != nil {
		fmt.Printf("can't delete members: %v", err)
		cli.OsExiter(-1)
		return
	}

}

func doBeforeMembersNuke(c *cli.Context) {

	if !c.BoolT("force") {

		red := color.New(color.FgRed)
		boldRed := red.Add(color.Bold)

		confirm := prompter.YN(boldRed.Sprint("WARNING: Do you really want to nuke all members?"), false)
		if !confirm {
			cli.OsExiter(-1)
			return
		}

		confirm = prompter.YN(boldRed.Sprint("WARNING: Terrible things will happen. Do you really-really want to nuke all members?"), false)
		if !confirm {
			cli.OsExiter(-1)
			return
		}
	}

}
func doActionMembersNuke(c *cli.Context) {

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

	err = members.Nuke(accessToken)
	if err != nil {
		fmt.Printf("can't nuke members: %v", err)
		cli.OsExiter(-1)
		return
	}

}