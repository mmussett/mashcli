package applicationpackagekeys

import (
	"encoding/json"
	"errors"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"

	"fmt"
	"io/ioutil"
	"os"
	"strings"
)


func ShowAllApplicationPackageKeys(accessToken, applicationId, format, filter string) error {

	apkc := new([]ApplicationPackageKeys)

	apkc, err := GetCollection(accessToken, &MethodParams{ApplicationId:applicationId}, &mashcli.Params{Fields: APPLICATIONPACKAGEKEYS_ALL_FIELDS}, &mashcli.Filter{Filter:filter})

	if err != nil {
		return err
	}

	if format=="table" {
		table := tablewriter.NewWriter(os.Stdout)

		caption := fmt.Sprintf("Application ID: %s", applicationId)
		table.SetCaption(true, caption)
		table.SetHeader([]string{"Package Key ID", "Package Key", "Package", "Plan", "Status", "Created", "Updated"})

		for _, s := range *apkc {
			data := []string{s.Id, s.Apikey, s.Package.Name, s.Plan.Name, s.Status, s.Created[:19], s.Updated[:19]}
			table.Append(data)

		}
		table.Render()
	} else {
		b, err := json.MarshalIndent(apkc, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(b))
	}

	return nil

}


func Import(accessToken, filename string, mp *MethodParams) error {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	a := new([]ApplicationPackageKeys)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	for _,v := range *a {
		apk := new(CreateApplicationPackageKeysRequest)
		apk.Apikey = v.Apikey
		apk.Package.Id = v.Package.Id
		apk.Plan.Id =  v.Plan.Id
		apk.Status = v.Status
		res, err := apk.Create(accessToken,mp)
		if err != nil {
			return err
		}

		if v.Status == "active" {
			e := new(mashcli.MasheryError)
			s := new(ApplicationPackageKeys)
			path := fmt.Sprintf("v3/rest/applications/%s/packageKeys/%s", mp.ApplicationId, res.Id)
			sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Body(strings.NewReader("{\"status\":\"active\"}")).Set("Accept", "application/json").Set("Content-Type", "application/json").Put(path).Receive(s,e)
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

	m, err := GetCollection(accessToken, mp, &mashcli.Params{Fields: APPLICATIONPACKAGEKEYS_ALL_FIELDS}, &mashcli.Filter{Filter:""})
	if err != nil {
		return err
	}

	filename := path + "/applicationpackagekeys-" + mp.ApplicationId + ".json"

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

func ReadFile(filename string) (*ApplicationPackageKeys, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(ApplicationPackageKeys)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (a *ApplicationPackageKeys) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
