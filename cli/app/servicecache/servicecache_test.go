package servicecache

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
)

var mp = &MethodParams{ServiceId: SERVICE_ID}

func TestGet(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	m, err := Get(tok, mp, &mashcli.Params{Fields: SERVICECACHE_ALL_FIELDS})
	spew.Dump(err)

	a.Nil(err)

	spew.Dump(m)

}

func TestServiceCache_Create(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	m := new(ServiceCache)
	m.CacheTtl = 73
	m,err = m.Create(tok,mp)
	a.Nil(err)

}


func TestUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var m *ServiceCache

	m, err = Get(tok, mp, &mashcli.Params{Fields: SERVICECACHE_ALL_FIELDS})
	a.Nil(err)

	m.PrettyPrint(mp)

	m.CacheTtl = 30

	m, err = m.Update(tok, mp)
	a.Nil(err)

	m, err = Get(tok, mp, &mashcli.Params{Fields: SERVICECACHE_ALL_FIELDS})
	a.Nil(err)

	m.PrettyPrint(mp)

	//a.True(m.ClientSurrogateControlEnabled != true, "update failed")

}

func TestShowServiceCache(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ShowServiceCache(tok, mp)

	a.Nil(err)

}

func TestDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	ShowServiceCache(tok, mp)

	err = Delete(tok, mp)
	a.Nil(err)
	ShowServiceCache(tok, mp)

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

	ShowServiceCache(tok, mp)
	m, err := Import(tok, "/Users/markmussett/Desktop/servicecache-x2dmsgggz5gdmesg5n5s8a8t.json", mp)
	a.Nil(err)

	m.PrettyPrint(mp)
}
