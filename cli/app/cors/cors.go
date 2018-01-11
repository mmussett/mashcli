package cors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mmussett/mashcli/cli/app"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"os"
	"strconv"
)

func ShowCors(accessToken, serviceId, endpointId, format string) error {

	m, err := Get(accessToken, &MethodParams{ServiceId:serviceId,EndpointId:endpointId}, &mashcli.Params{Fields: CORS_ALL_FIELDS})

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

func (m *Cors) PrettyPrint() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Allow any domain", "Sub-domain Matching", "Cookies Allowed", "Max Age", "Domains Allowed", "Headers Allowed", "Headers Exposed"})
	data := []string{app.FormatBool(m.AllDomainsEnabled), app.FormatBool(m.SubDomainMatchingAllowed), app.FormatBool(m.CookiesAllowed), strconv.FormatInt(m.MaxAge, 10), app.FlattenStringArray(m.DomainsAllowed), app.FlattenStringArray(m.HeadersAllowed), app.FlattenStringArray(m.HeadersExposed)}
	table.Append(data)
	table.Render()
	return

}

func Import(accessToken, filename string, mp *MethodParams) (*Cors, error) {

	a, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var m *Cors
	m, err = a.Create(accessToken, mp)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func Export(accessToken, path string, mp *MethodParams) error {

	valid, err := app.DirExists(path)
	if err != nil {
		return err
	}

	if !valid {
		return errors.New("Directory " + path + " does not exist")
	}

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: CORS_ALL_FIELDS})
	if err != nil {
		return err
	}

	filename := path + "/cors-" + mp.ServiceId + "-" + mp.EndpointId + ".json"
	err = m.WriteFile(filename)
	if err != nil {
		return err
	}

	return nil
}

func (a *Cors) WriteFile(filename string) error {

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

func ReadFile(filename string) (*Cors, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(Cors)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (a *Cors) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
