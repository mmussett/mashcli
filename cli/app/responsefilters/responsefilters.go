package responsefilters

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

func ShowResponseFilters(accessToken string, mp *MethodParams) error {

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: RESPONSEFILTERS_ALL_FIELDS})

	if err != nil {
		return err
	}

	m.PrettyPrint(mp)

	return nil

}

func ShowAllResponseFilters(accessToken string, mp *MethodParams) error {

	mc, err := GetCollection(accessToken, mp, &mashcli.Params{Fields: RESPONSEFILTERS_ALL_FIELDS})

	caption := fmt.Sprintf("Service ID: %s - Endpoint ID: %s - Method ID: %s\n", mp.ServiceId, mp.EndpointId, mp.MethodId)

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Response Filter ID", "Name", "JSON Filter Fields", "XML Filter Fields", "Notes", "Created", "Updated"})
	table.SetCaption(true, caption)

	for _, m := range *mc {
		data := []string{m.Id, m.Name, m.JsonFilterFields, m.XmlFilterFields, m.Notes, m.Created[:19], m.Updated[:19]}
		table.Append(data)
	}
	table.Render()

	return nil

}

func (m *ResponseFilters) PrettyPrint(mp *MethodParams) {

	caption := fmt.Sprintf("Service ID: %s - Endpoint ID: %s - Method ID: %s\n", mp.ServiceId, mp.EndpointId, mp.MethodId)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Response Filter ID", "Name", "JSON Filter Fields", "XML Filter Fields", "Notes", "Created", "Updated"})
	table.SetCaption(true, caption)
	data := []string{m.Id, m.Name, m.JsonFilterFields, m.XmlFilterFields, m.Notes, m.Created[:19], m.Updated[:19]}
	table.Append(data)
	table.Render()
	return

}

func Import(accessToken, filename string, mp *MethodParams) (*ResponseFilters, error) {

	a, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a.Id = ""

	m, err := a.Create(accessToken, mp)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func ExportAll(accessToken, path string, mp *MethodParams) error {

	m, err := GetCollection(accessToken, mp, &mashcli.Params{Fields: RESPONSEFILTERS_ALL_FIELDS})

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	filename := path + "/methods.json"

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
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

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: RESPONSEFILTERS_ALL_FIELDS})
	if err != nil {
		return err
	}

	filename := path + "/methods-" + m.Id + "-" + m.Name + ".json"
	err = m.WriteFile(filename)
	if err != nil {
		return err
	}

	return nil
}

func (a *ResponseFilters) WriteFile(filename string) error {

	data, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil

}

func ReadFile(filename string) (*ResponseFilters, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(ResponseFilters)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (a *ResponseFilters) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
