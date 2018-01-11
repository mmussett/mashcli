package planservices

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
	PLAN_ID = "fa16ea8a-db42-4b61-9517-8c6bc1ad9bf2"
	SERVICE_ID = "vx6f3xyyr99sahfwkbvw82dm"
)

var mp = &MethodParams{PackageId: PACKAGE_ID,PlanId:PLAN_ID,ServiceId:SERVICE_ID}


func TestGet(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	p, err := Get(tok, mp, &mashcli.Params{Fields: PLANSERVICES_ALL_FIELDS})

	a.Nil(err)

	a.True(p.Id == SERVICE_ID, "incorrect plan returned")

	spew.Dump(p)

}

func TestGetCollection(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var p *[]PlanServices

	p, err = GetCollection(tok, mp, &mashcli.Params{Fields: PLANSERVICES_ALL_FIELDS})

	a.Nil(err)

	spew.Dump(p)

}

func TestCreate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var p *PlanServices

	ShowAllPlanServices(tok, mp)

	p, err = Get(tok, mp, &mashcli.Params{Fields: PLANSERVICES_ALL_FIELDS})

	a.Nil(err)

	newPackageName := "Clone " + p.Name
	p.Id = ""
	p.Name = newPackageName
	p, err = p.Create(tok,mp)
	a.Nil(err)

	ShowAllPlanServices(tok, mp)

}

func TestUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var p *PlanServices

	ShowAllPlanServices(tok, mp)

	p, err = Get(tok, mp, &mashcli.Params{Fields: PLANSERVICES_ALL_FIELDS})
	a.Nil(err)

	p, err = p.Update(tok,mp)

	p, err = Get(tok, mp, &mashcli.Params{Fields: PLANSERVICES_ALL_FIELDS})
	a.Nil(err)

	ShowAllPlanServices(tok, mp)

}

func TestDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var p *PlanServices

	ShowAllPlanServices(tok, mp)

	p, err = Get(tok, mp, &mashcli.Params{Fields: PLANSERVICES_ALL_FIELDS})

	a.Nil(err)

	newPackageName := "Delete Me"
	p.Id = ""
	p.Name = newPackageName
	p, err = p.Create(tok,mp)
	a.Nil(err)

	ShowAllPlanServices(tok, mp)

	err = Delete(tok, mp)
	a.Nil(err)

	_, err = Get(tok, mp, &mashcli.Params{Fields: PLANSERVICES_ALL_FIELDS})
	a.NotNil(err)

	ShowAllPlanServices(tok, mp)

}

func TestCleardown(t *testing.T) {
	var c mashcli.Config

	json.Unmarshal(config, &c)
	tok, _ := c.FetchOAuthToken()

	ShowAllPlanServices(tok, mp)
	Delete(tok, &MethodParams{PackageId:PACKAGE_ID,PlanId:""})
	ShowAllPlanServices(tok, mp)
}

func TestShowAllPlanServices(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ShowAllPlanServices(tok,mp)
	a.Nil(err)

}

func TestShowPlanService(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	err = ShowPlanService(tok, mp)
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

	_, err = Import(tok, "/Users/markmussett/Desktop/planservices-x2dmsgggz5gdmesg5n5s8a8t-Mashcli Test API.json",mp)
	a.Nil(err)
	ShowAllPlanServices(tok,mp)
}

func TestExportAll(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ExportAll(tok, "/Users/markmussett/Desktop",mp)
	a.Nil(err)
}

func TestImportAll(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ImportAll(tok, "/Users/markmussett/Desktop/plans.json",mp)
	a.Nil(err)

	ShowAllPlanServices(tok,mp)
}
