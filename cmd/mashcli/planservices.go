package main

import (
	"fmt"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/mmussett/mashcli/cli/app/planservices"
	"github.com/urfave/cli"
)

func doBeforePlanServiceShow(c *cli.Context) {
	if len(c.Args()) != 3 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli planservice show --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionPlanServiceShow(c *cli.Context) {

	var packageId = c.Args().Get(0)
	var planId = c.Args().Get(1)
	var serviceId = c.Args().Get(2)

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


	err = planservices.ShowPlanService(accessToken, packageId, planId,serviceId, format)
	if err != nil {
		fmt.Printf("can't show planservice: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforePlanServiceShowAll(c *cli.Context) {
	if len(c.Args()) != 2 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli planservice showall --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionPlanServiceShowAll(c *cli.Context) {

	var packageId = c.Args().Get(0)
	var planId = c.Args().Get(1)

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

	err = planservices.ShowAllPlanServices(accessToken, packageId, planId, format, filter)
	if err != nil {
		fmt.Printf("can't show all services for plans: %v", err)
		cli.OsExiter(-1)
		return
	}

	return
}
