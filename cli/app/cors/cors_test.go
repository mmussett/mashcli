package cors

import (
	"encoding/json"
	"fmt"
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

	m, err := Get(tok, mp, &mashcli.Params{Fields: CORS_ALL_FIELDS})
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

	m := new(Cors)

	var allowedDomains []string
	allowedDomains = append(allowedDomains, "tibco.com")
	m.AllDomainsEnabled = true
	m.CookiesAllowed = true
	m.SubDomainMatchingAllowed = true
	m.DomainsAllowed = allowedDomains
	m.MaxAge = 64000
	m, err = m.Create(tok, mp)
	a.Nil(err)

	err = ShowCors(tok, mp)

}

func TestUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var m *Cors

	m, err = Get(tok, mp, &mashcli.Params{Fields: CORS_ALL_FIELDS})
	a.Nil(err)

	m.PrettyPrint(mp)

	m.CookiesAllowed = true
	m.SubDomainMatchingAllowed = true
	m.AllDomainsEnabled = true
	m.MaxAge = 500
	m, err = m.Update(tok, mp)
	a.Nil(err)

	m, err = Get(tok, mp, &mashcli.Params{Fields: CORS_ALL_FIELDS})
	a.Nil(err)

	m.PrettyPrint(mp)

	a.True(m.CookiesAllowed == true, "update failed")

}

func TestCleardown(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	Delete(tok, &MethodParams{ServiceId: SERVICE_ID, EndpointId: ENDPOINT_ID})

	ShowCors(tok, mp)
}

func TestDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	ShowCors(tok, mp)
	err = Delete(tok, mp)
	a.Nil(err)

	ShowCors(tok, mp)

}

func TestShowCors(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	err = ShowCors(tok, mp)

	a.Nil(err)

}

func TestCors_Marshall(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	apk, err := Get(tok, mp, &mashcli.Params{Fields: CORS_ALL_FIELDS})
	a.Nil(err)

	fmt.Println(apk.Marshall())

}

func TestCors_WriteFile(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	apk, err := Get(tok, mp, &mashcli.Params{Fields: CORS_ALL_FIELDS})
	a.Nil(err)

	apk.WriteFile("/Users/markmussett/Desktop/cors.json")
}

func TestCors_ReadFile(t *testing.T) {

	a := assert.New(t)

	apk, err := ReadFile("/Users/markmussett/Desktop/cors.json")
	a.Nil(err)

	fmt.Println(apk.Marshall())
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

	_, err = Import(tok, "/Users/markmussett/Desktop/cors-x2dmsgggz5gdmesg5n5s8a8t-zx22ky4m2m47dpkw2njxrg95.json", mp)
	a.Nil(err)

	err = ShowCors(tok, mp)
	a.Nil(err)

}
