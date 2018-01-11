package servicecache

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

func ShowServiceCache(accessToken string, mp *MethodParams) error {

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: SERVICECACHE_ALL_FIELDS})

	if err != nil {
		return err
	}

	m.PrettyPrint(mp)

	return nil

}

func (m *ServiceCache) PrettyPrint(mp *MethodParams) {

	caption := fmt.Sprintf("Service ID: %s", mp.ServiceId)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cache TTL"})
	table.SetCaption(true, caption)
	data := []string{strconv.Itoa(m.CacheTtl)}
	table.Append(data)
	table.Render()
	return

}

func Import(accessToken, filename string, mp *MethodParams) (*ServiceCache, error) {

	a, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: SERVICECACHE_ALL_FIELDS})
	if err != nil {
		m, err = a.Create(accessToken, mp)
		if err != nil {
			return nil, err
		}
	} else {
		m, err = a.Update(accessToken, mp)
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

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: SERVICECACHE_ALL_FIELDS})
	if err != nil {
		return err
	}

	filename := path + "/servicecache-" + mp.ServiceId + ".json"
	err = m.WriteFile(filename)
	if err != nil {
		return err
	}

	return nil
}

func (a *ServiceCache) WriteFile(filename string) error {

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

func ReadFile(filename string) (*ServiceCache, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(ServiceCache)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (a *ServiceCache) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
