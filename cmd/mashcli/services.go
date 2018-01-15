package main

import (
	"fmt"

	"github.com/Songmu/prompter"
	// "fmt"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/mmussett/mashcli/cli/app/services"
	"github.com/urfave/cli"
	"github.com/fatih/color"
	// "os"
)

func doBeforeServiceShow(c *cli.Context) {
	if len(c.Args()) == 0 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli service show --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionServiceShow(c *cli.Context) {

	// Arg0 = Service ID
	var serviceId = c.Args().Get(0)

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

	err = services.ShowService(accessToken, serviceId, format)
	if err != nil {
		fmt.Printf("can't show service: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforeServiceShowAll(c *cli.Context) {
}

func doActionServiceShowAll(c *cli.Context) {

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

	err = services.ShowAllServices(accessToken, format, filter)
	if err != nil {
		fmt.Printf("can't show all services: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforeServiceAdd(c *cli.Context) {

	if len(c.Args()) == 0 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli service add --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionServiceAdd(c *cli.Context) {

	var name, version, description = "", "", ""
	var aggregateQps int64 = 0

	name = c.Args().Get(0)

	if c.IsSet("version") {
		version = c.String("version")
	}

	if c.IsSet("description") {
		description = c.String("description")
	}

	if c.IsSet("qps") {
		aggregateQps = c.Int64("qps")
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

	err = services.AddService(accessToken, name, version, description, aggregateQps)
	if err != nil {
		fmt.Printf("can't add service: %v", err)
		cli.OsExiter(-1)
		return
	}

	return
}

func doBeforeServiceExport(c *cli.Context) {

	if len(c.Args()) == 0 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli service export --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionServiceExport(c *cli.Context) {

	var serviceId = c.Args().Get(0)
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

	err = services.Export(accessToken, serviceId, filename)
	if err != nil {
		fmt.Printf("can't export service: %v", err)
		cli.OsExiter(-1)
		return
	}

}

func doBeforeServiceImport(c *cli.Context) {
}

func doActionServiceImport(c *cli.Context) {

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

	_, err = services.Import(accessToken, filename)
	if err != nil {
		fmt.Printf("can't import service: %v", err)
		cli.OsExiter(-1)
		return
	}
}

func doBeforeServiceDelete(c *cli.Context) {

	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli service delete --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionServiceDelete(c *cli.Context) {

	var serviceId = c.Args().Get(0)

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

	err = services.DeleteService(accessToken, serviceId)
	if err != nil {
		fmt.Printf("can't delete service: %v", err)
		cli.OsExiter(-1)
		return
	}

}

func doBeforeServiceClone(c *cli.Context) {

	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli service clone --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionServiceClone(c *cli.Context) {

	var serviceId = c.Args().Get(0)

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

	err = services.CloneService(accessToken, serviceId)
	if err != nil {
		fmt.Printf("can't clone service: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforeServiceNuke(c *cli.Context) {


	if !c.BoolT("force") {
		red := color.New(color.FgRed)
		boldRed := red.Add(color.Bold)

		confirm := prompter.YN(boldRed.Sprint("WARNING: Do you really want to nuke all services?"), false)
		if !confirm {
			cli.OsExiter(-1)
			return
		}

		confirm = prompter.YN(boldRed.Sprint("WARNING: Terrible things will happen. Do you really-really want to nuke all services?"), false)
		if !confirm {
			cli.OsExiter(-1)
			return
		}
	}

}
func doActionServiceNuke(c *cli.Context) {

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

	err = services.Nuke(accessToken)
	if err != nil {
		fmt.Printf("can't nuke services: %v", err)
		cli.OsExiter(-1)
		return
	}

}