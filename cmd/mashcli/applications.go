package main

import (
	"fmt"
	"github.com/Songmu/prompter"
	"github.com/mmussett/mashcli/cli/app/applications"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/urfave/cli"
	"github.com/fatih/color"
)

func doBeforeApplicationsShow(c *cli.Context) {
	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli applications show --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionApplicationsShow(c *cli.Context) {

	// Arg0 = Application ID
	var applicationId = c.Args().Get(0)

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

	err = applications.ShowApplication(accessToken, applicationId,format)
	if err != nil {
		fmt.Printf("can't show application: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforeApplicationsShowAll(c *cli.Context) {
}

func doActionApplicationsShowAll(c *cli.Context) {

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


	err = applications.ShowAllApplications(accessToken,format,filter)
	if err != nil {
		fmt.Printf("can't show all applications: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforeApplicationsExport(c *cli.Context) {

	if len(c.Args()) < 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli applications export --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionApplicationsExport(c *cli.Context) {

	var applicationId = c.Args().Get(0)
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

	err = applications.Export(accessToken, applicationId, filename)
	if err != nil {
		fmt.Printf("can't export application: %v", err)
		cli.OsExiter(-1)
		return
	}

}

func doBeforeApplicationsImport(c *cli.Context) {
}

func doActionApplicationsImport(c *cli.Context) {

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

	_, err = applications.Import(accessToken, filename)
	if err != nil {
		fmt.Printf("can't import application: %v", err)
		cli.OsExiter(-1)
		return
	}
}

func doBeforeApplicationsDelete(c *cli.Context) {

	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli applications delete --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionApplicationsDelete(c *cli.Context) {

	var applicationId = c.Args().Get(0)

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

	err = applications.DeleteApplication(accessToken, applicationId)
	if err != nil {
		fmt.Printf("can't delete application: %v", err)
		cli.OsExiter(-1)
		return
	}

}


func doBeforeApplicationsNuke(c *cli.Context) {

	if !c.BoolT("force") {

		red := color.New(color.FgRed)
		boldRed := red.Add(color.Bold)

		confirm := prompter.YN(boldRed.Sprint("WARNING: Do you really want to nuke all applications?"), false)
		if !confirm {
			cli.OsExiter(-1)
			return
		}

		confirm = prompter.YN(boldRed.Sprint("WARNING: Terrible things will happen. Do you really-really want to nuke all applications?"), false)
		if !confirm {
			cli.OsExiter(-1)
			return
		}
	}

}
func doActionApplicationsNuke(c *cli.Context) {

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

	err = applications.Nuke(accessToken)
	if err != nil {
		fmt.Printf("can't nuke applications: %v", err)
		cli.OsExiter(-1)
		return
	}

}