package main

import (
	"fmt"

	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/urfave/cli"
)


func doBeforeConfigAdd(c *cli.Context) {

}

func doActionConfigAdd(c *cli.Context) {

	mashcli.Add()
	return

}


func doBeforeConfigShow(c *cli.Context) {

}

func doActionConfigShow(c *cli.Context) {

	m, err := mashcli.Load(c.String("area"))
	if err != nil {
		fmt.Printf("unable to load area config: %v", err)
		cli.OsExiter(-1)
		return
	}


	m.PrettyPrint()
	return

}
