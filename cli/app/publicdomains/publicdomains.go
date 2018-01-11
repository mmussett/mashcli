package publicdomains

import (
	"encoding/json"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"os"
)

func ShowAllPublicDomains(accessToken string) error {

	a, err := GetCollection(accessToken, &mashcli.Params{Fields: PUBLICDOMAINS_ALL_FIELDS})

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Endpoint ID", "Name", "Method", "Domain", "Path", "Created", "Updated"})

	for _, m := range *a {
		data := []string{m.Id, m.Name, m.Method, m.Domain, m.Path, m.Created[:19], m.Updated[:19]}
		table.Append(data)
	}
	table.Render()

	return nil

}

func (m *PublicDomains) PrettyPrint() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Endpoint ID", "Name", "Method", "Domain", "Path", "Created", "Updated"})
	data := []string{m.Id, m.Name, m.Method, m.Domain, m.Path, m.Created[:19], m.Updated[:19]}
	table.Append(data)
	table.Render()

	return
}

func ExportAll(accessToken string, path string) error {

	phc, err := GetCollection(accessToken, &mashcli.Params{Fields: PUBLICDOMAINS_ALL_FIELDS})

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(phc, "", "  ")
	if err != nil {
		return err
	}
	filename := path + "/publicdomains.json"

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil

}

func (a *PublicDomains) WriteFile(filename string) error {

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

func ReadFile(filename string) (*PublicDomains, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(PublicDomains)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (a *PublicDomains) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
