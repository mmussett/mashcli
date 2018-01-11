package securityprofiles

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	_ "github.com/davecgh/go-spew/spew"
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

	m, err := Get(tok, mp, &mashcli.Params{Fields: SECURITYPROFILES_ALL_FIELDS})
	a.Nil(err)

	fmt.Println(m.Marshall())

	spew.Dump(m)

}

func TestUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	_, err = Get(tok, mp, &mashcli.Params{Fields: SECURITYPROFILES_ALL_FIELDS})

	m2 := new(SecurityProfile)
	m2, err = m2.Update(tok, mp)
	a.Nil(err)
}

func TestShowSecurityProfile(t *testing.T) {

	assert := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	if err != nil {
		assert.Error(err)
	}

	err = ShowSecurityProfile(tok, mp)

	if err != nil {
		assert.Error(err)
	}

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

	ShowSecurityProfile(tok, mp)

	_, err = Import(tok, "/Users/markmussett/Desktop/securityprofile-x2dmsgggz5gdmesg5n5s8a8t.json", mp)
	a.Nil(err)

	ShowSecurityProfile(tok, mp)

}
