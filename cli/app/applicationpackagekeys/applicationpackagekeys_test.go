package applicationpackagekeys

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
	APPLICATION_ID = "d985067d-1568-4152-94f7-82acdbf4537d"
	PACKAGEKEY_ID  = "477b6d79-7139-4cd2-ade0-e80ebf072b46"
)

var mp = &MethodParams{ApplicationId: APPLICATION_ID, PackageKeyId: PACKAGEKEY_ID}

func TestGetCollection(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var apk *[]ApplicationPackageKeys

	apk, err = GetCollection(tok, mp, &mashcli.Params{Fields: APPLICATIONPACKAGEKEYS_ALL_FIELDS})

	a.Nil(err)

	spew.Dump(apk)

	ShowAllApplicationPackageKeys(tok,mp)

}

func TestUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)


	ShowAllApplicationPackageKeys(tok, mp)


}

func TestDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	ShowAllApplicationPackageKeys(tok, mp)

	err = Delete(tok, &MethodParams{ApplicationId:APPLICATION_ID,PackageKeyId:"933f124d-44c7-4cfb-b0fa-40182875cdf4"})
	a.Nil(err)

	ShowAllApplicationPackageKeys(tok, mp)

}

func TestShowAllApplicationPackageKeys(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ShowAllApplicationPackageKeys(tok, mp)
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

	err = Import(tok, "/Users/markmussett/Desktop/applicationpackagekeys-d985067d-1568-4152-94f7-82acdbf4537d.json", mp)
	a.Nil(err)

	err = ShowAllApplicationPackageKeys(tok, mp)
	a.Nil(err)

}
