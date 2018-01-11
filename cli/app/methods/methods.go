package methods

import (
	"encoding/json"
	"errors"
	"github.com/mmussett/mashcli/cli/app"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"

	"fmt"
	"io/ioutil"
	"os"
)

func ShowMethods(accessToken, serviceId, endpointId, methodId, format string) error {

	m, err := Get(accessToken, &MethodParams{ServiceId:serviceId,EndpointId:endpointId, MethodId:methodId}, &mashcli.Params{Fields: METHODS_ALL_FIELDS})

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

func ShowAllMethods(accessToken, serviceId, endpointId, format string) error {

	mc, err := GetCollection(accessToken, &MethodParams{ServiceId:serviceId,EndpointId:endpointId}, &mashcli.Params{Fields: METHODS_ALL_FIELDS})
	if err != nil {
		return err
	}

	if format=="table" {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Method ID", "Name", "Sample JSON Response", "Sample XML Response", "Created", "Updated"})
		table.SetCaption(true, "Service ID : "+serviceId)

		for _, m := range *mc {
			var sampleJsonPlaceholder, sampleXmlPlaceholder string
			if len(m.SampleJsonResponse) > 0 {
				sampleJsonPlaceholder = "{...}"
			}
			if len(m.SampleXmlResponse) > 0 {
				sampleXmlPlaceholder = "<>...</>"
			}
			data := []string{m.Id, m.Name, sampleJsonPlaceholder, sampleXmlPlaceholder, m.Created[:19], m.Updated[:19]}
			table.Append(data)
		}
		table.Render()
	} else {
		b, err := json.MarshalIndent(mc, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(b))
	}

	return nil

}

func (m *Methods) PrettyPrint() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Method ID", "Name", "sample JSON Response", "sample XML Response", "Created", "Updated"})
	var sampleJsonPlaceholder, sampleXmlPlaceholder string
	if len(m.SampleJsonResponse) > 0 {
		sampleJsonPlaceholder = "{...}"
	}
	if len(m.SampleXmlResponse) > 0 {
		sampleXmlPlaceholder = "<>...</>"
	}
	data := []string{m.Id, m.Name, sampleJsonPlaceholder, sampleXmlPlaceholder, m.Created[:19], m.Updated[:19]}
	table.Append(data)
	table.Render()
	return

}

func Import(accessToken, filename string, mp *MethodParams) (*Methods, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(Methods)
	err = json.Unmarshal(data, &a)
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

	m, err := GetCollection(accessToken, mp, &mashcli.Params{Fields: METHODS_ALL_FIELDS})

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

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: METHODS_ALL_FIELDS})
	if err != nil {
		return err
	}

	filename := path + "/methods-" + m.Id + "-" + m.Name + ".json"

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


func (a *Methods) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
