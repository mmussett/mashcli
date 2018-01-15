package planservices

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mmussett/mashcli/cli/app"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"

	"io/ioutil"
	"os"
)


func ShowPlanService(accessToken, packageId, planId, serviceId, format  string) error {
	m, err := Get(accessToken, &MethodParams{PackageId:packageId,PlanId:planId,ServiceId:serviceId}, &mashcli.Params{Fields: PLANSERVICES_ALL_FIELDS})

	if err != nil {
		return err
	}

	if format=="table" {
		m.PrettyPrint()
	} else {
		fmt.Println(m.Marshall())
	}

	return nil
}

func ShowAllPlanServices(accessToken, packageId, planId, format, filter string) error {

	pc, err := GetCollection(accessToken,  &MethodParams{PackageId:packageId,PlanId:planId}, &mashcli.Params{Fields: PLANSERVICES_ALL_FIELDS}, &mashcli.Filter{Filter:filter})

	if err != nil {
		return err
	}

	if format=="table" {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"API Name", "Service ID", "Endpoint Name", "Endpoint ID", "Method Name", "Method ID"})
		table.SetCaption(true, "Package ID : "+packageId)
		table.SetAutoMergeCells(true)
		table.SetRowLine(true)
		for _, p := range *pc {
			data := []string{p.Name, p.Id, "", "", "", ""}
			table.Append(data)

			for _, e := range p.Endpoints {
				data := []string{p.Name, p.Id, e.Name, e.Id, "", ""}
				table.Append(data)

				for _, m := range e.Methods {
					data := []string{p.Name, p.Id, e.Name, e.Id, m.Name, m.Id}
					table.Append(data)

				}
			}

		}

		table.Render()
	} else {
		b, err := json.MarshalIndent(pc, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(b))
	}
	return nil

}


func (p *PlanServices) PrettyPrint() {

	data := []string{p.Id, p.Name, p.Created[:19], p.Updated[:19]}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Plan ID", "Plan Name", "Created", "Updated"})
	table.SetAutoFormatHeaders(true)
	table.Append(data)
	table.Render()
	return

}


func Import(accessToken, filename string, mp *MethodParams) (*PlanServices, error) {


	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(PlanServices)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	a.Id = ""

	m, err := a.Create(accessToken,mp)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func ExportAll(accessToken, path string, mp *MethodParams) error {

	m, err := GetCollection(accessToken, mp, &mashcli.Params{Fields: PLANSERVICES_ALL_FIELDS},&mashcli.Filter{Filter:""})

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	filename := path + "/planservices.json"

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil

}

func ImportAll(accessToken, filename string, mp *MethodParams) error {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	a := new([]PlanServices)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	for _,p := range *a {
		p.Id = ""
		_, err := p.Create(accessToken,mp)
		if err != nil {
			return err
		}

	}

	return nil
}

func Export(accessToken, path string, mp *MethodParams) error {

	valid, err := app.DirExists(path)
	if err != nil {
		return err
	}

	if !valid {
		return errors.New("Directory " + path + " does not exist")
	}

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: PLANSERVICES_ALL_FIELDS})
	if err != nil {
		return err
	}

	filename := path + "/planservices-" + m.Id + "-" + m.Name + ".json"

	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (a *PlanServices) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}



