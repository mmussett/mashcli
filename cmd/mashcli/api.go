package main

import (
	"fmt"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/mmussett/mashcli/cli/app/transform"
	"github.com/urfave/cli"
)


func doBeforeSwaggerImport(c *cli.Context) {

	if len(c.Args()) != 2 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli api import --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionSwaggerImport(c *cli.Context) {

	var filename = c.Args().Get(0)
	var publicDomain = c.Args().Get(1)

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

	err = transform.ImportSwagger(accessToken,filename,publicDomain)
	if err != nil {
		fmt.Printf("can't import swagger : %v", err)
		cli.OsExiter(-1)
		return
	}
}
