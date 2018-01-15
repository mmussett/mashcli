package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/mmussett/mashcli/cli/app/packagekeys"
	"github.com/urfave/cli"
	"github.com/Songmu/prompter"
	"github.com/fatih/color"
)

func doBeforePackageKeySetRates(c *cli.Context) {
	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli packagekeys setrates --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionPackageKeySetRates(c *cli.Context) {

	var packageKeyId = c.Args().Get(0)

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

	if c.IsSet("throttle") {
		var argVal = c.String("throttle")
		switch argVal {
		case "unlimited":
			err = packagekeys.SetThrottleRateUnlimited(accessToken, packageKeyId)
			if err != nil {
				fmt.Printf("can't set package key rates: %v", err)
				cli.OsExiter(-1)
				return
			}
		case "default":
			err = packagekeys.SetThrottleRateDefault(accessToken, packageKeyId)
			if err != nil {
				fmt.Printf("can't set package key rates: %v", err)
				cli.OsExiter(-1)
				return
			}
		default:
			rate, err := strconv.Atoi(argVal)
			if err != nil {
				fmt.Printf("can't set package key rates: %v", err)
				cli.OsExiter(-1)
				return
			}

			err = packagekeys.SetThrottleRate(accessToken, packageKeyId, rate)
			if err != nil {
				fmt.Printf("can't set package key rates: %v", err)
				cli.OsExiter(-1)
				return
			}
		}
	}

	if c.IsSet("quota") {
		var argVal = c.String("quota")
		switch argVal {
		case "unlimited":
			err = packagekeys.SetQuotaRateUnlimited(accessToken, packageKeyId)
			if err != nil {
				fmt.Printf("can't set package key rates: %v", err)
				cli.OsExiter(-1)
				return
			}
		case "default":
			err = packagekeys.SetQuotaRateDefault(accessToken, packageKeyId)
			if err != nil {
				fmt.Printf("can't set package key rates: %v", err)
				cli.OsExiter(-1)
				return
			}
		default:
			rate, err := strconv.Atoi(argVal)
			if err != nil {
				fmt.Printf("can't set package key rates: %v", err)
				cli.OsExiter(-1)
				return
			}
			err = packagekeys.SetQuotaRate(accessToken, packageKeyId, rate)
			if err != nil {
				fmt.Printf("can't set package key rates: %v", err)
				cli.OsExiter(-1)
				return
			}
		}
	}

	return

}

func doBeforePackageKeySetStatus(c *cli.Context) {
	if len(c.Args()) != 2 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli packagekeys setstatus --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionPackageKeySetStatus(c *cli.Context) {

	var packageKeyId = c.Args().Get(0)
	var status = c.Args().Get(1)

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

	switch strings.ToLower(status) {
	case "active":
		status = "active"
	case "disabled":
		status = "disabled"
	default:
		fmt.Println("status must be 'active' or 'disabled'")
		cli.OsExiter(-1)
		return
	}

	err = packagekeys.SetStatus(accessToken, packageKeyId, status)
	if err != nil {
		fmt.Printf("can't set package key status: %v", err)
		cli.OsExiter(-1)
		return
	}
	return

}

func doBeforePackageKeysDelete(c *cli.Context) {

	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli packagekeys delete --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionPackageKeysDelete(c *cli.Context) {

	var packageKeyId = c.Args().Get(0)

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

	err = packagekeys.DeletePackageKey(accessToken, packageKeyId)
	if err != nil {
		fmt.Printf("can't delete package key: %v", err)
		cli.OsExiter(-1)
		return
	}

}

func doBeforePackageKeysShow(c *cli.Context) {
	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli packagekeys show --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionPackageKeysShow(c *cli.Context) {

	// Arg0 = Package ID
	var packageKeyId = c.Args().Get(0)

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

	err = packagekeys.ShowPackageKeys(accessToken, packageKeyId,format)
	if err != nil {
		fmt.Printf("can't show package keys: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforePackageKeysShowAll(c *cli.Context) {
}

func doActionPackageKeysShowAll(c *cli.Context) {

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

	err = packagekeys.ShowAllPackageKeys(accessToken,format,filter)
	if err != nil {
		fmt.Printf("can't show all package keys: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforePackageKeysExport(c *cli.Context) {

	if len(c.Args()) < 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli packagekeys export --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionPackageKeysExport(c *cli.Context) {

	var packageKeyId = c.Args().Get(0)
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

	err = packagekeys.Export(accessToken, filename, packageKeyId)
	if err != nil {
		fmt.Printf("can't export package key: %v", err)
		cli.OsExiter(-1)
		return
	}

}

func doBeforePackageKeysImport(c *cli.Context) {
}

func doActionPackageKeysImport(c *cli.Context) {

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

	_, err = packagekeys.Import(accessToken, filename)
	if err != nil {
		fmt.Printf("can't import package key: %v", err)
		cli.OsExiter(-1)
		return
	}
}


func doBeforePackageKeysNuke(c *cli.Context) {

	if !c.BoolT("force") {
		red := color.New(color.FgRed)
		boldRed := red.Add(color.Bold)

		confirm := prompter.YN(boldRed.Sprint("WARNING: Do you really want to nuke all package keys?"), false)
		if !confirm {
			cli.OsExiter(-1)
			return
		}

		confirm = prompter.YN(boldRed.Sprint("WARNING: Terrible things will happen. Do you really-really want to nuke all package keys?"), false)
		if !confirm {
			cli.OsExiter(-1)
			return
		}
	}

}

func doActionPackageKeysNuke(c *cli.Context) {

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

	err = packagekeys.Nuke(accessToken)
	if err != nil {
		fmt.Printf("can't nuke package keys: %v", err)
		cli.OsExiter(-1)
		return
	}


}
