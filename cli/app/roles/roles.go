package roles

import (
	"encoding/json"
	"errors"
	"github.com/mmussett/mashcli/cli/app"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"os"
)

func ShowRoles(accessToken string, mp *MethodParams) error {

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: ROLES_ALL_FIELDS})

	if err != nil {
		return err
	}

	m.PrettyPrint()

	return nil

}

func ShowAllRoles(accessToken string) error {

	a, err := GetCollection(accessToken, &mashcli.Params{Fields: ROLES_ALL_FIELDS})

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Role ID", "Name", "Created", "Updated"})

	for _, m := range *a {
		data := []string{m.Id, m.Name, m.Created[:19], m.Updated[:19]}
		table.Append(data)
	}
	table.Render()

	return nil

}

func (m *Roles) PrettyPrint() {

	data := []string{m.Name, m.Created[:19], m.Updated[:19]}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Created", "Updated"})
	table.Append(data)
	table.Render()

	return

}


func Import(accessToken, filename string, mp *MethodParams) (*Roles, error) {


	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(Roles)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	a.Id = ""

	m, err := a.Update(accessToken, mp)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func ExportAll(accessToken, path string) error {

	m, err := GetCollection(accessToken, &mashcli.Params{Fields: ROLES_ALL_FIELDS})

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	filename := path + "/roles.json"

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil

}

func ImportAll(accessToken, filename string) error {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	a := new([]Roles)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	for _,p := range *a {
		p.Id = ""
		_, err := p.Create(accessToken)
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

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: ROLES_ALL_FIELDS})
	if err != nil {
		return err
	}

	filename := path + "/roles-" + m.Id + ".json"

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

func (a *Roles) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}