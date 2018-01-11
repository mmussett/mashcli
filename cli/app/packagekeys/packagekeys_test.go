package packagekeys

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
	PACKAGE_KEY_ID = "fb58764b-cbfd-4cbe-ba0f-626dec212160"
)

var mp = &MethodParams{PackageKeyId: PACKAGE_KEY_ID}


func TestGet(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	apk, err := Get(tok, mp, &mashcli.Params{Fields: PACKAGEKEYS_ALL_FIELDS})

	a.Nil(err)

	a.True(apk.Id == PACKAGE_KEY_ID, "incorrect application package key returned")

	spew.Dump(apk)

}

func TestGetCollection(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var apk *[]PackageKeys

	apk, err = GetCollection(tok, &mashcli.Params{Fields: PACKAGEKEYS_ALL_FIELDS})

	a.Nil(err)

	spew.Dump(apk)

}

func TestUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	apk, err := Get(tok, mp, &mashcli.Params{Fields: PACKAGEKEYS_ALL_FIELDS})
	a.Nil(err)

	apk.QPSLimitCeiling = 50

	ShowPackageKeys(tok, mp)

	apk, err = apk.Update(tok,mp)
	a.Nil(err)

	ShowPackageKeys(tok, mp)

}

func TestDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	ShowAllPackageKeys(tok)

	err = Delete(tok, &MethodParams{PackageKeyId:"fb58764b-cbfd-4cbe-ba0f-626dec212160"})
	a.Nil(err)

	ShowAllPackageKeys(tok)

}

func TestShowAllPackageKeys(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ShowAllPackageKeys(tok)
	a.Nil(err)

}

func TestShowPackageKeys(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	err = ShowPackageKeys(tok, mp)

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

	_, err = Import(tok, "/Users/markmussett/Desktop/packagekeys-fb58764b-cbfd-4cbe-ba0f-626dec212160.json",mp)
	a.Nil(err)
	ShowAllPackageKeys(tok)
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

func TestImportAll(t *testing.T) {

	a := assert.New(t)
	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ImportAll(tok, "/Users/markmussett/Desktop/packagekeys.json")
	a.Nil(err)
	ShowAllPackageKeys(tok)
}