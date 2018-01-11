package systemdomainauthentication

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
	SERVICE_ID  = "x2dmsgggz5gdmesg5n5s8a8t"
	ENDPOINT_ID = "zx22ky4m2m47dpkw2njxrg95"
)

var mp = &MethodParams{ServiceId: SERVICE_ID, EndpointId: ENDPOINT_ID}

func TestGet(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	m, err := Get(tok, mp, &mashcli.Params{Fields: SYSTEMDOMAINAUTHENTICATION_ALL_FIELDS})
	spew.Dump(err)

	a.Nil(err)

	spew.Dump(m)

}

func TestCreate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	m := new(SystemDomainAuthentication)

	m.Username = "tibco"
	m.Password = "tibco123"
	m.Type = type_httpBasic
	m.Certificate = "test"
	m, err = m.Create(tok, mp)
	a.Nil(err)

	err = ShowSystemDomainAuthentication(tok, mp)

}

func TestUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var m *SystemDomainAuthentication

	m, err = Get(tok, mp, &mashcli.Params{Fields: SYSTEMDOMAINAUTHENTICATION_ALL_FIELDS})
	a.Nil(err)

	m.PrettyPrint(mp)

	m.Username = "tibco2"
	m.Password = "tibco123"
	m.Type = "clientSslCert"
	m.Certificate = "test"
	m, err = m.Update(tok, mp)
	a.Nil(err)

	m, err = Get(tok, mp, &mashcli.Params{Fields: SYSTEMDOMAINAUTHENTICATION_ALL_FIELDS})
	a.Nil(err)

	m.PrettyPrint(mp)

	a.True(m.Username == "tibco2", "update failed")

}

func TestCleardown(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	Delete(tok, &MethodParams{ServiceId: SERVICE_ID, EndpointId: ENDPOINT_ID})

	ShowSystemDomainAuthentication(tok, mp)
}

func TestDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	ShowSystemDomainAuthentication(tok, mp)

	err = Delete(tok, mp)
	a.Nil(err)

	_, err = Get(tok, mp, &mashcli.Params{Fields: SYSTEMDOMAINAUTHENTICATION_ALL_FIELDS})
	a.NotNil(err)

	ShowSystemDomainAuthentication(tok, mp)

}

func TestShowSystemDomainAuthentication(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	err = ShowSystemDomainAuthentication(tok, mp)

	a.Nil(err)

}
