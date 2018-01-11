package domains

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

func ShowDomains(accessToken, domainId, format string) error {

	m, err := Get(accessToken, &MethodParams{DomainId:domainId}, &mashcli.Params{Fields: DOMAINS_ALL_FIELDS})

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

func ShowAllDomains(accessToken, format string) error {

	dc, err := GetCollection(accessToken, &mashcli.Params{Fields: DOMAINS_ALL_FIELDS})
	if err != nil {
		return err
	}

	if format=="table" {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Domain ID", "Domain", "Status", "Created"})

		for _, m := range *dc {
			data := []string{m.Id, m.Domain, m.Status, m.Created[:19]}
			table.Append(data)
		}
		table.Render()
	} else {
		b, err := json.MarshalIndent(dc, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(b))
	}

	return nil

}

func (m *Domains) PrettyPrint() {

	data := []string{m.Id, m.Domain, m.Status, m.Created[:19]}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Domain ID", "Domain", "Status", "Created"})
	table.Append(data)
	table.Render()

	return

}

func ExportAll(accessToken string, path string) error {

	ma, err := GetCollection(accessToken, &mashcli.Params{Fields: DOMAINS_ALL_FIELDS})

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(ma, "", "  ")
	if err != nil {
		return err
	}
	filename := path + "/domains.json"

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

	d, err := Get(accessToken, mp, &mashcli.Params{Fields: DOMAINS_ALL_FIELDS})
	if err != nil {
		return err
	}

	filename := path + "/domains-" + mp.DomainId + ".json"
	err = d.WriteFile(filename)
	if err != nil {
		return err
	}

	return nil
}

func Import(accessToken, filename string, mp *MethodParams) (*Domains, error) {

	a, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a.Id = ""
	a.PrettyPrint()

	m, err := a.Create(accessToken)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (a *Domains) WriteFile(filename string) error {

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

func ReadFile(filename string) (*Domains, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(Domains)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (a *Domains) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
