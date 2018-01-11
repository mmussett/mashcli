package systemhostnames

import (
	"encoding/json"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"os"
)

func ShowAllSystemHostnames(accessToken string) error {

	shc, err := GetCollection(accessToken, &mashcli.Params{Fields: SYSTEMHOSTNAMES_ALL_FIELDS})

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Address"})

	for _, sh := range *shc {
		data := []string{sh.Address}
		table.Append(data)
	}
	table.Render()

	return nil

}

func (m *SystemHostnames) PrettyPrint() {

	data := []string{m.Address}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Address"})
	table.Append(data)
	table.Render()

	return

}

func ExportAll(accessToken string, path string) error {

	phc, err := GetCollection(accessToken, &mashcli.Params{Fields: SYSTEMHOSTNAMES_ALL_FIELDS})

	if err != nil {
		return err
	}

	for _, ph := range *phc {
		filename := path + "/systemhostnames-" + ph.Address + ".json"
		err = ph.WriteFile(filename)
		if err != nil {
			return err
		}
	}

	return nil

}

func (a *SystemHostnames) WriteFile(filename string) error {

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

func ReadFile(filename string) (*SystemHostnames, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(SystemHostnames)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (a *SystemHostnames) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
