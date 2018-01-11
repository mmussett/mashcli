package main

import (
	"fmt"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/mmussett/mashcli/cli/app/packages"
	"github.com/mmussett/mashcli/cli/app/plans"
	"github.com/urfave/cli"
	"strings"
)

func doBeforePlanSetStatus(c *cli.Context) {
	if len(c.Args()) != 3 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli plan setstatus --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionPlanSetStatus(c *cli.Context) {

	var packageId = c.Args().Get(0)
	var planId = c.Args().Get(1)
	var status = c.Args().Get(2)

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
	case "inactive":
		status = "inactive"
	default:
		fmt.Println("status must be 'active' or 'inactive'")
		cli.OsExiter(-1)
		return
	}

	err = plans.SetStatus(accessToken, packageId, planId, status)


	return

}

func doBeforePlanSetRateLimits(c *cli.Context) {
	if len(c.Args()) != 2 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli plan setratelimit --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionPlanSetRateLimits(c *cli.Context) {

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

	var throttle = -1
	var quota = -1
	var quotaPeriod, throttleOverride, quotaOverride = "", "", ""

	if c.IsSet("throttle") {
		throttle = c.Int("throttle")
	}

	if c.IsSet("quota") {
		quota = c.Int("quota")
	}

	if c.IsSet("throttleoverride") {
		switch strings.ToLower(c.String("throttleoverride")) {
		case "true":
			throttleOverride = "true"
		case "false":
			throttleOverride = "false"
		}
	}

	if c.IsSet("quotaoverride") {
		switch strings.ToLower(c.String("quotaoverride")) {
		case "true":
			quotaOverride = "true"
		case "false":
			quotaOverride = "false"
		}
	}

	if c.IsSet("quotaperiod") {
		switch strings.ToLower(c.String("quotaperiod")) {
		case "minute":
			quotaPeriod = "minute"
		case "hour":
			quotaPeriod = "hour"
		case "day":
			quotaPeriod = "day"
		case "month":
			quotaPeriod = "month"
		}
	}

	err = plans.SetRateLimits(accessToken, packageId, planId, throttle, throttleOverride, quota, quotaPeriod, quotaOverride)

	return

}

func doBeforePlanSetKeyProperties(c *cli.Context) {
	if len(c.Args()) != 2 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli plan setkeyproperties --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionPlanSetKeyProperties(c *cli.Context) {

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

	var maxKeysAllowed = -1
	var maxKeysAllowedModerated = -1
	var selfServiceKeyProvisioning, adminKeyProvisioning = "", ""

	if c.IsSet("maxkeys") {
		maxKeysAllowed = c.Int("maxkeys")
	}

	if c.IsSet("keysmoderated") {
		maxKeysAllowedModerated = c.Int("keysmoderated")
	}

	if c.IsSet("selfservicekeys") {
		switch strings.ToLower(c.String("selfservicekeys")) {
		case "true":
			selfServiceKeyProvisioning = "true"
		case "false":
			selfServiceKeyProvisioning = "false"
		}
	}

	if c.IsSet("adminkeys") {
		switch strings.ToLower(c.String("adminkeys")) {
		case "true":
			adminKeyProvisioning = "true"
		case "false":
			adminKeyProvisioning = "false"
		}
	}

	err = plans.SetKeyProperties(accessToken, packageId, planId, selfServiceKeyProvisioning, adminKeyProvisioning, maxKeysAllowed, maxKeysAllowedModerated)

	return

}

func doBeforePlanShow(c *cli.Context) {
	if len(c.Args()) != 2 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli plan show --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionPlanShow(c *cli.Context) {

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

	err = plans.ShowPlan(accessToken, packageId, planId,format)
	if err != nil {
		fmt.Printf("can't show package plan: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforePlanShowAll(c *cli.Context) {
	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli plan showall --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionPlanShowAll(c *cli.Context) {

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
	err = plans.ShowAllPlans(accessToken, packageId, format)
	if err != nil {
		fmt.Printf("can't show all package plans: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforePlanAdd(c *cli.Context) {

	if len(c.Args()) < 2 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli plans add --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionPlanAdd(c *cli.Context) {

	packageId := c.Args().Get(0)
	planName := c.Args().Get(1)
	planDesc := c.Args().Get(2)

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

	err = plans.AddPlan(accessToken, packageId, planName, planDesc)
	if err != nil {
		fmt.Printf("can't add plan: %v", err)
		cli.OsExiter(-1)
		return
	}

	return
}

func doBeforePlanExport(c *cli.Context) {

	if len(c.Args()) < 2 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli plan export --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionPlanExport(c *cli.Context) {

	var packageId = c.Args().Get(0)
	var planId = c.Args().Get(1)
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

	err = plans.Export(accessToken, packageId, planId, filename)
	if err != nil {
		fmt.Printf("can't export plan: %v", err)
		cli.OsExiter(-1)
		return
	}

}

func doBeforePlanImport(c *cli.Context) {
}

func doActionPlanImport(c *cli.Context) {

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
		fmt.Printf("can't import plan: %v", err)
		cli.OsExiter(-1)
		return
	}
}

func doBeforePlanDelete(c *cli.Context) {

	if len(c.Args()) != 2 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli plan delete --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionPlanDelete(c *cli.Context) {

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

	err = plans.DeletePlan(accessToken, packageId, planId)
	if err != nil {
		fmt.Printf("can't delete plan: %v", err)
		cli.OsExiter(-1)
		return
	}

}

func doBeforePlanClone(c *cli.Context) {

	if len(c.Args()) < 2 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli plan clone --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionPlanClone(c *cli.Context) {

	var packageId = c.Args().Get(0)
	var planId = c.Args().Get(1)
	var packageName = c.Args().Get(2)
	var packageDesc = c.Args().Get(3)

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

	err = plans.ClonePlan(accessToken, packageId, planId, packageName, packageDesc)
	if err != nil {
		fmt.Printf("can't clone plan: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}
