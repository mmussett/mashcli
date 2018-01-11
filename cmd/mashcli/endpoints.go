package main

import (
	"fmt"
	"github.com/mmussett/mashcli/cli/app/endpoints"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/urfave/cli"
)

func doBeforeEndpointsShow(c *cli.Context) {
	if len(c.Args()) != 2 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli endpoints show --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionEndpointsShow(c *cli.Context) {

	// Arg0 = Service ID
	// Arg1 = Endpoint ID
	var serviceId = c.Args().Get(0)
	var endpointId = c.Args().Get(1)

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

	err = endpoints.ShowEndpoints(accessToken,serviceId,endpointId,format)
	if err != nil {
		fmt.Printf("can't show endpoint: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}

func doBeforeEndpointsShowAll(c *cli.Context) {
	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli endpoints show-all  --help' for usage")
		cli.OsExiter(-1)
	}
}

func doActionEndpointsShowAll(c *cli.Context) {

	// Arg0 = Service ID
	var serviceId = c.Args().Get(0)

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

	err = endpoints.ShowAllEndpoints(accessToken,serviceId,format)
	if err != nil {
		fmt.Printf("can't show all endpoints: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}


func doBeforeEndpointsExport(c *cli.Context) {

	if len(c.Args()) != 2 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli endpoints export --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionEndpointsExport(c *cli.Context) {

	var serviceId = c.Args().Get(0)
	var endpointId = c.Args().Get(1)
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

	err = endpoints.Export(accessToken, serviceId, endpointId, filename)
	if err != nil {
		fmt.Printf("can't export service: %v", err)
		cli.OsExiter(-1)
		return
	}

}

func doBeforeEndpointsImport(c *cli.Context) {

	if len(c.Args()) != 1 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli endpoints import --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionEndpointsImport(c *cli.Context) {

	var serviceId = c.Args().Get(0)
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

	_,err  = endpoints.Import(accessToken, filename, &endpoints.MethodParams{ServiceId:serviceId})
	if err != nil {
		fmt.Printf("can't import endpoints: %v", err)
		cli.OsExiter(-1)
		return
	}
}

func doBeforeEndpointsDelete(c *cli.Context) {

	if len(c.Args()) != 2 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli endpoints delete --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionEndpointsDelete(c *cli.Context) {

	var serviceId = c.Args().Get(0)
	var endpointId = c.Args().Get(1)

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

	err = endpoints.DeleteEndpoint(accessToken,serviceId,endpointId)
	if err != nil {
		fmt.Printf("can't delete endpoint: %v", err)
		cli.OsExiter(-1)
		return
	}

}

func doBeforeEndpointsAdd(c *cli.Context) {

	if len(c.Args()) != 4 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli endpoints add --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionEndpointsAdd(c *cli.Context) {

	var serviceId = c.Args().Get(0)
	var endpointName = c.Args().Get(1)
	var publicDomain = c.Args().Get(2)
	var privateDomain = c.Args().Get(3)

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

	err = endpoints.AddEndpoint(accessToken, serviceId, endpointName, publicDomain, privateDomain)
	if err != nil {
		fmt.Printf("can't add endpoint: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}


func doBeforeEndpointsClone(c *cli.Context) {

	if len(c.Args()) != 5 {
		fmt.Println("mashcli: argument mismatch")
		fmt.Println("Run 'mashcli endpoints clone --help' for usage")
		cli.OsExiter(-1)
	}

}

func doActionEndpointsClone(c *cli.Context) {

	var serviceId = c.Args().Get(0)
	var endpointId = c.Args().Get(1)
	var endpointName = c.Args().Get(2)
	var publicDomain = c.Args().Get(3)
	var privateDomain = c.Args().Get(4)

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

	err = endpoints.CloneEndpoint(accessToken, serviceId, endpointId, endpointName, publicDomain, privateDomain)
	if err != nil {
		fmt.Printf("can't clone endpoint: %v", err)
		cli.OsExiter(-1)
		return
	}

	return

}
