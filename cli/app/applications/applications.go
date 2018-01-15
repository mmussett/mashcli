package applications

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

func Nuke(accessToken string) error {

	ac := new([]Applications)

	ac, err := GetCollection(accessToken, &mashcli.Params{Fields: APPLICATIONS_ALL_FIELDS})
	if err != nil {
		return err
	}


	for _, a := range *ac {
		err := DeleteApplication(accessToken, a.Id)
		if err != nil {
			return err
		}
	}

	return nil

}

func DeleteApplication(accessToken, applicationId string) error {

	err := Delete(accessToken,&MethodParams{ApplicationId:applicationId})
	if err != nil {
		return err
	}

	return nil
}


func ShowApplication(accessToken, applicationId, format string) error {

	m, err := Get(accessToken, &MethodParams{ApplicationId:applicationId}, &mashcli.Params{Fields: APPLICATIONS_ALL_FIELDS})

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

func ShowAllApplications(accessToken, format string) error {

	ac := new([]Applications)

	ac, err := GetCollection(accessToken, &mashcli.Params{Fields: APPLICATIONS_ALL_FIELDS})
	if err != nil {
		return err
	}

	if format=="table" {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "External ID", "Name", "Description", "Username", "Status", "Type", "Commercial", "Runs Adverts", "Usage", "Protocol", "Output", "Created", "Updated"})

		for _, s := range *ac {
			data := []string{s.Id, s.ExternalID, s.Name, s.Description, s.Username, s.Status, s.Type, app.FormatBool(s.Commercial), app.FormatBool(s.Ads), s.UsageModel, s.PreferredProtocol, s.PreferredOutput, s.Created[:19], s.Updated[:19]}
			table.Append(data)
		}

		table.Render()
	} else {
		b, err := json.MarshalIndent(ac, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(b))
	}

	return nil

}

func (a *Applications) PrettyPrint() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "External ID", "Name", "Description", "Username", "Status", "Type", "Commercial", "Runs Adverts", "Usage", "Protocol", "Output", "Created", "Updated"})
	data := []string{a.Id, a.ExternalID, a.Name, a.Description, a.Username, a.Status, a.Type, app.FormatBool(a.Commercial), app.FormatBool(a.Ads), a.UsageModel, a.PreferredProtocol, a.PreferredOutput, a.Created[:19], a.Updated[:19]}
	table.Append(data)
	table.Render()

	return
}



func ExportAll(accessToken, path string) error {

	valid, err := app.DirExists(path)
	if err != nil {
		return err
	}

	if !valid {
		return errors.New("Directory " + path + " does not exist")
	}

	m, err := GetCollection(accessToken, &mashcli.Params{Fields: APPLICATIONS_ALL_FIELDS})
	if err != nil {
		return err
	}

	filename := path + "/applications.json"

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

func Import(accessToken, filename string)  (*Applications, error) {

	if len(filename) != 0 {
		a, err := ReadFile(filename)
		if err != nil {
			return nil, err
		}

		var applicationId = a.Id
		a.Id = ""
		a.Created = ""
		a.Updated = ""

		a, err = a.Update(accessToken,&MethodParams{ApplicationId:applicationId})
		if err != nil {
			return nil, err
		}

		a.WriteStdOut()
		return a, nil

	} else {
		a, err := ReadStdIn()
		if err != nil {
			return nil, err
		}

		var applicationId = a.Id
		a.Id = ""
		a.Created = ""
		a.Updated = ""

		a, err = a.Update(accessToken,&MethodParams{ApplicationId:applicationId})
		if err != nil {
			return nil, err
		}

		a.WriteStdOut()
		return a, nil
	}

}

func Export(accessToken, applicationId, filename string) error {

	s, err := Get(accessToken, &MethodParams{ApplicationId: applicationId}, &mashcli.Params{Fields: APPLICATIONS_ALL_FIELDS})
	if err != nil {
		return err
	}

	if len(filename) != 0 {
		s.WriteFile(filename)
	} else {
		s.WriteStdOut()
	}

	return nil
}


func (p *Applications) WriteStdOut() error {

	file := os.Stdout

	b, err := json.MarshalIndent(p, "", " ")
	if err == nil {
		s := string(b)
		file.WriteString(s)
		file.Sync()
		return nil
	} else {
		return err
	}
}

func (p *Applications) WriteFile(filename string) error {

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil

}

func ReadStdIn() (*Applications, error) {

	var data []byte

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	s := new(Applications)
	json.Unmarshal(data, &s)
	return s, nil

}

func ReadFile(filename string) (*Applications, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(Applications)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}


func (a *Applications) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
