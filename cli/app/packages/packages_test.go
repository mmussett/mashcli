package packages

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
	PACKAGE_ID = "0f5d46bb-235f-4810-aebb-64e2669479cd"
)

var mp = &MethodParams{PackageId: PACKAGE_ID}


func TestGet(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	p, err := Get(tok, mp, &mashcli.Params{Fields: PACKAGE_ALL_FIELDS})

	a.Nil(err)

	a.True(p.Id == PACKAGE_ID, "incorrect package returned")

	spew.Dump(p)

}

func TestGetCollection(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var p *[]Package

	p, err = GetCollection(tok, &mashcli.Params{Fields: PACKAGE_ALL_FIELDS})

	a.Nil(err)

	spew.Dump(p)

}

func TestCreate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var p *Package

	ShowAllPackages(tok)

	p, err = Get(tok, mp, &mashcli.Params{Fields: PACKAGE_ALL_FIELDS})

	a.Nil(err)

	newPackageName := "Clone " + p.Name
	p.Id = ""
	p.Name = newPackageName
	p, err = p.Create(tok)
	a.Nil(err)

	ShowAllPackages(tok)

}

func TestUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var p *Package

	ShowAllPackages(tok)

	p, err = Get(tok, mp, &mashcli.Params{Fields: PACKAGE_ALL_FIELDS})
	a.Nil(err)

	p.Description = "Updated"
	p, err = p.Update(tok,mp)
	a.Nil(err)

	p, err = Get(tok, mp, &mashcli.Params{Fields: PACKAGE_ALL_FIELDS})
	a.Nil(err)

	a.True(p.Description == "Updated", "update failed")

	ShowAllPackages(tok)

}

func TestDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var p *Package

	ShowAllPackages(tok)

	p, err = Get(tok, mp, &mashcli.Params{Fields: PACKAGE_ALL_FIELDS})

	a.Nil(err)

	newPackageName := "Delete Me"
	p.Id = ""
	p.Name = newPackageName
	p, err = p.Create(tok)
	a.Nil(err)

	ShowAllPackages(tok)

	err = Delete(tok, &MethodParams{PackageId:p.Id})
	a.Nil(err)

	_, err = Get(tok, &MethodParams{PackageId:p.Id}, &mashcli.Params{Fields: PACKAGE_ALL_FIELDS})
	a.NotNil(err)

	ShowAllPackages(tok)

}

func TestCleardown(t *testing.T) {
	var c mashcli.Config

	json.Unmarshal(config, &c)
	tok, _ := c.FetchOAuthToken()

	ShowAllPackages(tok)
	Delete(tok, &MethodParams{PackageId:PACKAGE_ID})
	ShowAllPackages(tok)
}

func TestShowAllPackages(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ShowAllPackages(tok)
	a.Nil(err)

}

func TestShowPackage(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	err = ShowPackage(tok, mp)

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

	_, err = Import(tok, "/Users/markmussett/Desktop/packages-0f5d46bb-235f-4810-aebb-64e2669479cd-mashclitest.json")
	a.Nil(err)
	ShowAllPackages(tok)
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

	err = ImportAll(tok, "/Users/markmussett/Desktop/packages.json")
	a.Nil(err)
}
