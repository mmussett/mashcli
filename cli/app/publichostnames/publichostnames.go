package publichostnames

import (
	"encoding/json"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"os"
)

func ShowAllPublicHostnames(accessToken string) error {

	phc, err := GetCollection(accessToken, &mashcli.Params{Fields: PUBLICHOSTNAMES_ALL_FIELDS})

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Address"})

	for _, ph := range *phc {
		data := []string{ph.Address}
		table.Append(data)
	}
	table.Render()

	return nil

}

func (m *PublicHostnames) PrettyPrint() {

	data := []string{m.Address}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Address"})
	table.Append(data)
	table.Render()

	return

}

func ExportAll(accessToken string, path string) error {

	phc, err := GetCollection(accessToken, &mashcli.Params{Fields: PUBLICHOSTNAMES_ALL_FIELDS})

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(phc, "", "  ")
	if err != nil {
		return err
	}
	filename := path + "/publichostnames.json"

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil

}

func (a *PublicHostnames) WriteFile(filename string) error {

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

func ReadFile(filename string) (*PublicHostnames, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(PublicHostnames)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (a *PublicHostnames) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
