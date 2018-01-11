package errormessages

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

func ShowErrorMessages(accessToken string, mp *MethodParams) error {

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: ERRORMESSAGES_ALL_FIELDS})

	if err != nil {
		return err
	}

	m.PrettyPrint(mp)

	return nil

}

func (m *ErrorMessages) PrettyPrint(mp *MethodParams) {

	caption := fmt.Sprintf("Service ID: %s - Errorset ID: %s", mp.ServiceId, mp.errorSetId)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Code", "Status", "Header", "Response Body"})

	// Calculate width of status
	var statusWidth  = 0
	for _,v := range m.ErrorMessage {
		if len(v.Status) > statusWidth {
			statusWidth = len(v.DetailHeader)
		}
	}

	// Calculate width of header
	var headerWidth  = 0
	for _,v := range m.ErrorMessage {
		if len(v.DetailHeader) > headerWidth {
			headerWidth = len(v.DetailHeader)
		}
	}

	table.SetColMinWidth(1,statusWidth)
	table.SetColMinWidth(2,headerWidth)
	table.SetCaption(true, caption)

	for _,v := range m.ErrorMessage {
		data := []string{strconv.Itoa(v.Code),v.Status,v.DetailHeader, v.ResponseBody}
		table.Append(data)
	}
	table.Render()
	return

}

func Import(accessToken, filename string, mp *MethodParams) (*ErrorMessages, error) {

	a, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}

	m, err := Update(accessToken, mp, a)

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

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: ERRORMESSAGES_ALL_FIELDS})
	if err != nil {
		return err
	}

	filename := path + "/errormessages-" + mp.ServiceId + "-" + mp.errorSetId+".json"
	err = m.WriteFile(filename)
	if err != nil {
		return err
	}

	return nil
}

func (a *ErrorMessages) WriteFile(filename string) error {

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

func ReadFile(filename string) (*ErrorMessages, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(ErrorMessages)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (a *ErrorMessages) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
