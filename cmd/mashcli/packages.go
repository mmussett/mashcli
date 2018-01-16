package main

import (
	"fmt"
	"github.com/Songmu/prompter"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/mmussett/mashcli/cli/app/packages"
	"github.com/urfave/cli"
	"github.com/fatih/color"
)

func doBeforePackagesShow(c *cli.Context) {
	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli packages show --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionPackagesShow(c *cli.Context) {

	// Arg0 = Package ID
	var packageId = c.Args().Get(0)

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

	err = packages.ShowPackage(accessToken,  packageId, format)
	if err != nil {
		fmt.Printf("can't show package: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforePackagesShowAll(c *cli.Context) {
}

func doActionPackagesShowAll(c *cli.Context) {

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

	err = packages.ShowAllPackages(accessToken,format,filter,nameglob)
	if err != nil {
		fmt.Printf("can't show all packages: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforePackageAdd(c *cli.Context) {

	if len(c.Args()) == 0 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli packages add --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionPackageAdd(c *cli.Context) {

	packageName := c.Args().Get(0)
	packageDesc := c.Args().Get(1)

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

	err = packages.AddPackage(accessToken, packageName, packageDesc)
	if err != nil {
		fmt.Printf("can't add package: %v", err)
		cli.OsExiter(-1)
		return
	}

	return
}

func doBeforePackageExport(c *cli.Context) {

	if len(c.Args()) < 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli packages export --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionPackageExport(c *cli.Context) {

	var packageId = c.Args().Get(0)
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

	err = packages.Export(accessToken, packageId, filename)
	if err != nil {
		fmt.Printf("can't export package: %v", err)
		cli.OsExiter(-1)
		return
	}

}

func doBeforePackageImport(c *cli.Context) {
}

func doActionPackageImport(c *cli.Context) {

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

	_, err = packages.Import(accessToken, filename)
	if err != nil {
		fmt.Printf("can't import package: %v", err)
		cli.OsExiter(-1)
		return
	}
}

func doBeforePackageDelete(c *cli.Context) {

	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli packages delete --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionPackageDelete(c *cli.Context) {

	var packageId = c.Args().Get(0)

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

	err = packages.DeletePackage(accessToken, packageId)
	if err != nil {
		fmt.Printf("can't delete package: %v", err)
		cli.OsExiter(-1)
		return
	}

}

func doBeforePackageClone(c *cli.Context) {

	if len(c.Args()) == 0 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli packages clone --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionPackageClone(c *cli.Context) {

	var packageId = c.Args().Get(0)
	var packageName = c.Args().Get(1)
	var packageDesc = c.Args().Get(2)

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

	err = packages.ClonePackage(accessToken, packageId, packageName, packageDesc)
	if err != nil {
		fmt.Printf("can't clone package: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforePackageNuke(c *cli.Context) {

	if !c.BoolT("force") {

		red := color.New(color.FgRed)
		boldRed := red.Add(color.Bold)

		confirm := prompter.YN(boldRed.Sprint("WARNING: Do you really want to nuke all packages?"), false)
		if !confirm {
			cli.OsExiter(-1)
			return
		}

		confirm = prompter.YN(boldRed.Sprint("WARNING: Terrible things will happen. Do you really-really want to nuke all packages?"), false)
		if !confirm {
			cli.OsExiter(-1)
			return
		}
	}

}
func doActionPackageNuke(c *cli.Context) {

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

	err = packages.Nuke(accessToken, preview)
	if err != nil {
		fmt.Printf("can't nuke packages: %v", err)
		cli.OsExiter(-1)
		return
	}

}