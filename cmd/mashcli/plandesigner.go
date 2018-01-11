package main

import (
	"fmt"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/mmussett/mashcli/cli/app/plandesigner"
	"github.com/urfave/cli"
)


func doBeforePlanDesignerAdd(c *cli.Context) {

	if len(c.Args()) < 2 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli plandesigner add --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionPlanDesignerAdd(c *cli.Context) {

	packageId := c.Args().Get(0)
	planId := c.Args().Get(1)
	serviceId := c.Args().Get(2)

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

	err = plandesigner.AddServiceToPackagePlan(accessToken,serviceId,packageId,planId)
	if err != nil {
		fmt.Printf("can't add service to plan: %v", err)
		cli.OsExiter(-1)
		return
	}

	return
}

func doBeforePlanDesignerDelete(c *cli.Context) {

	if len(c.Args()) < 2 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli plandesigner delete --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionPlanDesignerDelete(c *cli.Context) {

	packageId := c.Args().Get(0)
	planId := c.Args().Get(1)
	serviceId := c.Args().Get(2)

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

	err = plandesigner.DeleteServiceFromPackagePlan(accessToken,serviceId,packageId,planId)
	if err != nil {
		fmt.Printf("can't delete service from plan: %v", err)
		cli.OsExiter(-1)
		return
	}

	return
}

