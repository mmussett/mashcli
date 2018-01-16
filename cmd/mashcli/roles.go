package main

import (
	"fmt"
	"github.com/Songmu/prompter"
	"github.com/mmussett/mashcli/cli/app/roles"

	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/urfave/cli"
	"github.com/fatih/color"

)

func doBeforeRolesShow(c *cli.Context) {
	if len(c.Args()) == 0 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli service show --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionRolesShow(c *cli.Context) {

	// Arg0 = Role ID
	var roleId = c.Args().Get(0)

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

	err = roles.ShowRole(accessToken, roleId, format)
	if err != nil {
		fmt.Printf("can't show role: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforeRolesShowAll(c *cli.Context) {
}

func doActionRolesShowAll(c *cli.Context) {

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

	var nameglob = ""
	if c.IsSet("name") {
		nameglob = c.String("name")
	}


	err = roles.ShowAllRoles(accessToken, format, filter, nameglob)
	if err != nil {
		fmt.Printf("can't show all roles: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforeRolesAdd(c *cli.Context) {

	if len(c.Args()) == 0 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli service add --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionRolesAdd(c *cli.Context) {

	var name =  ""

	name = c.Args().Get(0)


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

	err = roles.AddRole(accessToken, name)
	if err != nil {
		fmt.Printf("can't add role: %v", err)
		cli.OsExiter(-1)
		return
	}

	return
}

func doBeforeRolesExport(c *cli.Context) {

	if len(c.Args()) == 0 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli roles export --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionRolesExport(c *cli.Context) {

	var roleId = c.Args().Get(0)
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

	err = roles.Export(accessToken, roleId, filename)
	if err != nil {
		fmt.Printf("can't export roles: %v", err)
		cli.OsExiter(-1)
		return
	}

}

func doBeforeRolesImport(c *cli.Context) {
}

func doActionRolesImport(c *cli.Context) {

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

	_, err = roles.Import(accessToken, filename)
	if err != nil {
		fmt.Printf("can't import role: %v", err)
		cli.OsExiter(-1)
		return
	}
}

func doBeforeRolesDelete(c *cli.Context) {

	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli roles delete --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionRolesDelete(c *cli.Context) {

	var roleId = c.Args().Get(0)

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

	err = roles.DeleteRole(accessToken, roleId)
	if err != nil {
		fmt.Printf("can't delete role: %v", err)
		cli.OsExiter(-1)
		return
	}

}


func doBeforeRolesNuke(c *cli.Context) {


	if !c.BoolT("force") {
		red := color.New(color.FgRed)
		boldRed := red.Add(color.Bold)

		confirm := prompter.YN(boldRed.Sprint("WARNING: Do you really want to nuke all roles?"), false)
		if !confirm {
			cli.OsExiter(-1)
			return
		}

		confirm = prompter.YN(boldRed.Sprint("WARNING: Terrible things will happen. Do you really-really want to nuke all roles?"), false)
		if !confirm {
			cli.OsExiter(-1)
			return
		}
	}

}
func doActionRolesNuke(c *cli.Context) {

	var preview = false
	if c.IsSet("preview") {
		preview = c.Bool("preview")
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

	err = roles.Nuke(accessToken, preview)
	if err != nil {
		fmt.Printf("can't nuke roles: %v", err)
		cli.OsExiter(-1)
		return
	}

}