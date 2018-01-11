package errorsets

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

func ShowErrorSets(accessToken string, mp *MethodParams) error {

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: ERRORSETS_ALL_FIELDS})

	if err != nil {
		return err
	}

	m.PrettyPrint(mp)

	return nil

}

func (m *ErrorSets) PrettyPrint(mp *MethodParams) {

	caption := fmt.Sprintf("Service ID: %s", mp.ServiceId)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Error Set ID", "Name", "Type", "JSONP", "JSONP Type"})
	table.SetCaption(true, caption)

	for _,v := range m.ErrorSet {
		data := []string{v.ID,v.Name,v.Type, app.FormatBool(v.Jsonp),v.JsonpType}
		table.Append(data)
	}
	table.Render()
	return

}

func Import(accessToken, filename string, mp *MethodParams) (*ErrorSets, error) {

	a, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: ERRORSETS_ALL_FIELDS})
	if err != nil {
		m, err = a.CreateCollection(accessToken, mp)
		if err != nil {
			return nil, err
		}
	} else {
		m, err = a.UpdateCollection(accessToken, mp)
		if err != nil {
			return nil, err
		}
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

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: ERRORSETS_ALL_FIELDS})
	if err != nil {
		return err
	}

	filename := path + "/errorsets-" + mp.ServiceId + ".json"
	err = m.WriteFile(filename)
	if err != nil {
		return err
	}

	return nil
}

func (a *ErrorSets) WriteFile(filename string) error {

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

func ReadFile(filename string) (*ErrorSets, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(ErrorSets)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (a *ErrorSets) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
