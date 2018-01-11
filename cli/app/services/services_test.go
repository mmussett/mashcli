package services

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/stretchr/testify/assert"
	"testing"
)

var config = []byte(`
{
  "userid": "mmussett",
  "password": "f1rest0rm",
  "apikey": "7rn8vgpty6nywruhgc755qh6",
  "apikeysecret": "2q34GBG3nx",
  "name": "Emealocal1",
  "area": "c7e8e2d5-ff91-42eb-9885-10f2aa2cc3f5",
  "tm": "emealocal1.api.mashery.com",
  "ccurl": "https://emealocal1.admin.mashery.com/control-center"
}
`)

const (
	SERVICE_ID = "x2dmsgggz5gdmesg5n5s8a8t"
)

var mp = &MethodParams{ServiceId: SERVICE_ID}

func TestGet(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	s, err := Get(tok, mp, &mashcli.Params{Fields: SERVICES_ALL_FIELDS})

	a.Nil(err)

	a.True(s.Id == SERVICE_ID, "incorrect service returned")

	spew.Dump(s)

}

func TestGetCollection(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var p *[]Services

	p, err = GetCollection(tok, &mashcli.Params{Fields: SERVICES_ALL_FIELDS})

	a.Nil(err)

	spew.Dump(p)

}

func TestCreate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var s1 *Services
	var s2 = new(Services)

	ShowAllServices(tok)

	s1, err = Get(tok, mp, &mashcli.Params{Fields: SERVICES_ALL_FIELDS})
	a.Nil(err)

	s2.Name = "clone_" + s1.Name
	s2.Description = "clone_" + s1.Description

	s2, err = s2.Create(tok)
	a.Nil(err)

	ShowAllServices(tok)

	err = Delete(tok, &MethodParams{ServiceId: s2.Id})
	a.Nil(err)

}

func TestUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var s *Services

	ShowService(tok, mp)

	s, err = Get(tok, mp, &mashcli.Params{Fields: SERVICES_ALL_FIELDS})
	a.Nil(err)

	s.Description = "Updated"
	s, err = s.Update(tok, mp)
	a.Nil(err)

	s, err = Get(tok, mp, &mashcli.Params{Fields: SERVICES_ALL_FIELDS})
	a.Nil(err)

	a.True(s.Description == "Updated", "update failed")

	ShowService(tok, mp)

}

func TestDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var s *Services

	ShowAllServices(tok)

	s, err = Get(tok, mp, &mashcli.Params{Fields: SERVICES_ALL_FIELDS})

	a.Nil(err)

	newServiceName := "Clone " + s.Name
	s.Id = ""
	s.Name = newServiceName
	s, err = s.Create(tok)
	a.Nil(err)

	ShowAllServices(tok)

	err = Delete(tok, &MethodParams{ServiceId: s.Id})
	a.Nil(err)

	_, err = Get(tok, &MethodParams{ServiceId: s.Id}, &mashcli.Params{Fields: SERVICES_ALL_FIELDS})
	a.NotNil(err)

	ShowAllServices(tok)

}

func TestCleardown(t *testing.T) {
	var c mashcli.Config

	json.Unmarshal(config, &c)
	tok, _ := c.FetchOAuthToken()

	ShowAllServices(tok)
	Delete(tok, &MethodParams{ServiceId: ""})
	ShowAllServices(tok)
}

func TestShowAllServices(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ShowAllServices(tok)
	a.Nil(err)

}

func TestShowService(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ShowService(tok, mp)
	a.Nil(err)
}

func TestExport(t *testing.T) {

	a := assert.New(t)
	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = Export(tok, "/Users/markmussett/Desktop", mp)
	a.Nil(err)
}

func TestImport(t *testing.T) {

	a := assert.New(t)
	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	_, err = Import(tok, "/Users/markmussett/Desktop/services-x2dmsgggz5gdmesg5n5s8a8t-Mashcli Test API.json")
	a.Nil(err)
	ShowAllServices(tok)
}
func TestExportAll(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ExportAll(tok, "/Users/markmussett/Desktop")
	a.Nil(err)
}
