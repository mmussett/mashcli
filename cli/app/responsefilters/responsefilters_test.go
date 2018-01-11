package responsefilters

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
	SERVICE_ID        = "x2dmsgggz5gdmesg5n5s8a8t"
	ENDPOINT_ID       = "zx22ky4m2m47dpkw2njxrg95"
	METHOD_ID         = "eb8e1a05-96ad-47ed-9a41-f0550fc1bb72"
	RESPONSEFILTER_ID = "7195ab59-a750-442b-81c3-810310eb3f35"
)

var mp = &MethodParams{ServiceId: SERVICE_ID, EndpointId: ENDPOINT_ID, MethodId: METHOD_ID, ResponseFilterId: RESPONSEFILTER_ID}

func TestGet(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	m, err := Get(tok, mp, &mashcli.Params{Fields: RESPONSEFILTERS_ALL_FIELDS})
	a.Nil(err)

	a.True(m.Id == RESPONSEFILTER_ID, "incorrect response filter returned")

	spew.Dump(m)

}

func TestGetCollection(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var mc *[]ResponseFilters

	mc, err = GetCollection(tok, mp, &mashcli.Params{Fields: RESPONSEFILTERS_ALL_FIELDS})

	a.Nil(err)

	spew.Dump(mc)

}

func TestCreate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var m *ResponseFilters

	ShowAllResponseFilters(tok, mp)

	m, err = Get(tok, mp, &mashcli.Params{Fields: RESPONSEFILTERS_ALL_FIELDS})

	a.Nil(err)

	m.Name = "clone2" + m.Name
	m.Id = ""
	m, err = m.Create(tok, mp)
	a.Nil(err)

	ShowAllResponseFilters(tok, mp)

}

func TestUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var m *ResponseFilters

	m, err = Get(tok, mp, &mashcli.Params{Fields: RESPONSEFILTERS_ALL_FIELDS})
	a.Nil(err)

	m.PrettyPrint(mp)

	newName := m.Name + "_clone"
	m.Name = newName
	m, err = m.Update(tok, mp)
	a.Nil(err)

	m, err = Get(tok, mp, &mashcli.Params{Fields: RESPONSEFILTERS_ALL_FIELDS})
	a.Nil(err)

	m.PrettyPrint(mp)

	a.True(m.Name == newName, "update failed")

}

func TestCleardown(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	Delete(tok, &MethodParams{ServiceId: SERVICE_ID, EndpointId: ENDPOINT_ID, MethodId: METHOD_ID, ResponseFilterId: "ca762882-4423-491c-afb1-aa13534d41a0"})

	ShowAllResponseFilters(tok, mp)
}

func TestDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var m *ResponseFilters

	ShowAllResponseFilters(tok, mp)

	m, err = Get(tok, mp, &mashcli.Params{Fields: RESPONSEFILTERS_ALL_FIELDS})

	a.Nil(err)

	m.Name = "clone" + m.Name
	m.Id = ""
	m, err = m.Create(tok, mp)
	a.Nil(err)

	ShowAllResponseFilters(tok, mp)

	nmp := &MethodParams{ServiceId: SERVICE_ID, EndpointId: ENDPOINT_ID, MethodId: METHOD_ID, ResponseFilterId: m.Id}
	err = Delete(tok, nmp)
	a.Nil(err)

	m, err = Get(tok, nmp, &mashcli.Params{Fields: RESPONSEFILTERS_ALL_FIELDS})
	a.NotNil(err)

	ShowAllResponseFilters(tok, mp)

}

func TestShowAllResponseFilters(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ShowAllResponseFilters(tok, mp)
	a.Nil(err)

}

func TestShowResponseFilters(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	err = ShowResponseFilters(tok, mp)

	a.Nil(err)

}
