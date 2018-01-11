package methods

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
	SERVICE_ID  = "gykda32cdzsu257fyju9x52m"
	ENDPOINT_ID = "mkftqkvmjtum5ru73gsv8rh7"
	METHOD_ID   = "e4ad0b1e-cecd-49ce-90df-8899b9c9c1ec"
)

var mp = &MethodParams{ServiceId: SERVICE_ID, EndpointId: ENDPOINT_ID, MethodId: METHOD_ID}

func TestGet(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	m, err := Get(tok, mp, &mashcli.Params{Fields: METHODS_ALL_FIELDS})
	spew.Dump(err)

	a.Nil(err)

	a.True(m.Id == METHOD_ID, "incorrect method returned")

	spew.Dump(m)

}

func TestGetCollection(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var mc *[]Methods

	mc, err = GetCollection(tok, mp, &mashcli.Params{Fields: METHODS_ALL_FIELDS})

	a.Nil(err)

	spew.Dump(mc)

}

func TestCreate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var m *Methods

	ShowAllMethods(tok, mp)

	m, err = Get(tok, mp, &mashcli.Params{Fields: METHODS_ALL_FIELDS})

	a.Nil(err)

	m.Name = "clone3" + m.Name
	m.Id = ""
	m, err = m.Create(tok, mp)
	a.Nil(err)

	ShowAllMethods(tok, mp)

}

func TestUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var m *Methods

	m, err = Get(tok, mp, &mashcli.Params{Fields: METHODS_ALL_FIELDS})
	a.Nil(err)

	m.PrettyPrint()

	newName := m.Name + "_clone"
	m.Name = newName
	m, err = m.Update(tok, mp)
	a.Nil(err)

	m, err = Get(tok, &MethodParams{ServiceId: SERVICE_ID, EndpointId: ENDPOINT_ID, MethodId: "93a1e45b-2199-411f-8a4e-c7ee9d7fb0c5"}, &mashcli.Params{Fields: METHODS_ALL_FIELDS})
	a.Nil(err)

	m.PrettyPrint()

	a.True(m.Name == newName, "update failed")

}

func TestCleardown(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	Delete(tok, &MethodParams{ServiceId: SERVICE_ID, EndpointId: ENDPOINT_ID, MethodId: "93a1e45b-2199-411f-8a4e-c7ee9d7fb0c5"})

	ShowAllMethods(tok, mp)
}

func TestDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var m *Methods

	ShowAllMethods(tok, mp)

	m, err = Get(tok, mp, &mashcli.Params{Fields: METHODS_ALL_FIELDS})

	a.Nil(err)

	m.Name = "clone" + m.Name
	m.Id = ""
	m, err = m.Create(tok, mp)
	a.Nil(err)

	ShowAllMethods(tok, mp)

	nmp := &MethodParams{ServiceId: SERVICE_ID, EndpointId: ENDPOINT_ID, MethodId: m.Id}
	err = Delete(tok, nmp)
	a.Nil(err)

	m, err = Get(tok, nmp, &mashcli.Params{Fields: METHODS_ALL_FIELDS})
	a.NotNil(err)

	ShowAllMethods(tok, mp)

}

func TestShowAllMethods(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ShowAllMethods(tok, mp)
	a.Nil(err)

}

func TestShowMethods(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	err = ShowMethods(tok, mp)

	a.Nil(err)

}
