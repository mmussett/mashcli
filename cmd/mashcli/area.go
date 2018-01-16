package main

import (
	"fmt"
	"github.com/Songmu/prompter"
	"github.com/mmussett/mashcli/cli/app/areas"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/mmussett/mashcli/cli/app/members"
	"github.com/urfave/cli"
	"github.com/fatih/color"
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

func doBeforeAreaNuke(c *cli.Context) {

	if !c.BoolT("force") {



		red := color.New(color.FgRed)
		boldRed := red.Add(color.Bold)

		confirm := prompter.YN(boldRed.Sprint("WARNING: Do you really want to nuke the area?"), false)
		if !confirm {
			cli.OsExiter(-1)
			return
		}

		confirm = prompter.YN(boldRed.Sprint("WARNING: Terrible things will happen. Do you really-really want to nuke the area?"), false)
		if !confirm {
			cli.OsExiter(-1)
			return
		}

		confirm = prompter.YN(boldRed.Sprint("WARNING: There's no going back now. Honestly are you sure?"), false)
		if !confirm {
			cli.OsExiter(-1)
			return
		}
	}

}
func doActionAreaNuke(c *cli.Context) {

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

	err = members.Nuke(accessToken, preview)
	if err != nil {
		fmt.Printf("can't nuke area: %v", err)
		cli.OsExiter(-1)
		return
	}

}
