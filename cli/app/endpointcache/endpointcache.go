package endpointcache

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

func ShowEndpointCache(accessToken, serviceId, endpointId, format string) error {

	m, err := Get(accessToken, &MethodParams{ServiceId:serviceId,EndpointId:endpointId}, &mashcli.Params{Fields: ENDPOINTCACHE_ALL_FIELDS})

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

func (m *EndpointCache) PrettyPrint() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cache TTL", "Surrogate-Control Headers", "Respond From Stale Cache", "Vary Header", "Cache-Control Headers", "Developer Key Cache Key", "Content Cache Key Headers"})
	data := []string{strconv.Itoa(m.Cache.CacheTTLOverride), app.FormatBool(m.Cache.RespondFromStaleCacheEnabled), app.FormatBool(m.Cache.VaryHeaderEnabled), app.FormatBool(m.Cache.ClientSurrogateControlEnabled), app.FormatBool(m.Cache.ResponseCacheControlEnabled), app.FormatBool(m.Cache.IncludeAPIKeyInContentCacheKey), app.FlattenStringArray(m.Cache.ContentCacheKeyHeaders)}
	table.Append(data)
	table.Render()
	return

}

func Import(accessToken, filename string, mp *MethodParams) (*EndpointCache, error) {

	a, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}

	m, err := a.Create(accessToken, mp)
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

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: ENDPOINTCACHE_ALL_FIELDS})
	if err != nil {
		return err
	}

	filename := path + "/endpointcache-" + mp.ServiceId + "-" + mp.EndpointId + ".json"
	err = m.WriteFile(filename)
	if err != nil {
		return err
	}

	return nil
}

func (a *EndpointCache) WriteFile(filename string) error {

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

func ReadFile(filename string) (*EndpointCache, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(EndpointCache)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (a *EndpointCache) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
