package main

import (
	"fmt"
	"github.com/mmussett/mashcli/cli/app/areas"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/urfave/cli"
)

func doBeforeAreaBackup(c *cli.Context) {
	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli area backup --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionAreaBackup(c *cli.Context) {

	var backupSetName = c.Args().Get(0)

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

	err = areas.BackupArea(accessToken, backupSetName)
	if err != nil {
		fmt.Printf("can't backup area: %v", err)
		cli.OsExiter(-1)
		return
	}
}

func doBeforeAreaRestore(c *cli.Context) {
	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli area restore --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionAreaRestore(c *cli.Context) {

	var backupSetName = c.Args().Get(0)

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

	err = areas.RestoreArea(accessToken, backupSetName)
	if err != nil {
		fmt.Printf("can't restore area: %v", err)
		cli.OsExiter(-1)
		return
	}
}

