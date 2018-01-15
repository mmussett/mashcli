package main

import (
	"fmt"
	"github.com/mmussett/mashcli/cli/app/applicationpackagekeys"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/urfave/cli"
)


func doBeforeApplicationPackageKeysShowAll(c *cli.Context) {
	if len(c.Args()) == 0 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli applicationpackagekeys showall --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionApplicationPackageKeysShowAll(c *cli.Context) {

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

	var filter = ""
	if c.IsSet("filter") {
		filter = c.String("filter")
	}

	err = applicationpackagekeys.ShowAllApplicationPackageKeys(accessToken,applicationId,format, filter)
	if err != nil {
		fmt.Printf("can't show all application package keys: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

