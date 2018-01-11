package errorsets

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
	ERRORSET_ID = "b4511c3a-2663-433b-92d0-1fa8c2f9b84a"
)

var mp = &MethodParams{ServiceId: SERVICE_ID, errorSetId:ERRORSET_ID}

func TestGet(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	m, err := Get(tok, mp, &mashcli.Params{Fields: ERRORSETS_ALL_FIELDS})
	spew.Dump(err)

	a.Nil(err)

	spew.Dump(m)

}

func TestErrorSets_Create(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	ShowErrorSets(tok, mp)

	es := ErrorSet{Type:"application/json",Name:"Test4",Jsonp:false}

	_, err = es.Create(tok,mp)
	a.Nil(err)


	ShowErrorSets(tok, mp)
}


func TestErrorSets_Update(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)


	ShowErrorSets(tok, mp)

	es := new(ErrorSet)
	es.Type = "application/json"
	es.Jsonp = true
	es.JsonpType = "application/json"
	es, err = es.Update(tok,mp)
	a.Nil(err)


	ShowErrorSets(tok, mp)
	//a.True(m.ClientSurrogateControlEnabled != true, "update failed")

}

func TestShowErrorSets(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ShowErrorSets(tok, mp)

	a.Nil(err)

}

func TestDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	ShowErrorSets(tok, mp)

	err = Delete(tok, &MethodParams{ServiceId:SERVICE_ID,errorSetId:"1f5b5835-1ca1-4edb-96e2-61c958190058"})
	a.Nil(err)
	ShowErrorSets(tok, mp)

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

	ShowErrorSets(tok, mp)
	m, err := Import(tok, "/Users/markmussett/Desktop/errorsets-x2dmsgggz5gdmesg5n5s8a8t.json", mp)
	a.Nil(err)

	m.PrettyPrint(mp)
}
