package plans

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
	PLAN_ID = "8db24ee2-1ac0-4176-848a-213efefc31d4"
)

var mp = &MethodParams{PackageId: PACKAGE_ID,PlanId:PLAN_ID}


func TestGet(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	p, err := Get(tok, mp, &mashcli.Params{Fields: PLAN_ALL_FIELDS})

	a.Nil(err)

	a.True(p.Id == PLAN_ID, "incorrect plan returned")

	spew.Dump(p)

}

func TestGetCollection(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var p *[]Plan

	p, err = GetCollection(tok, mp, &mashcli.Params{Fields: PLAN_ALL_FIELDS})

	a.Nil(err)

	spew.Dump(p)

}

func TestCreate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)


	ShowAllPlans(tok, mp)

	p := new(Plan)
	p.Name = "mashclitest_plan3"
	p.Description = ""

	p, err = p.Create(tok,mp)
	a.Nil(err)

	ShowAllPlans(tok, mp)

}

func TestUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var p *Plan

	ShowAllPlans(tok, mp)

	p, err = Get(tok, mp, &mashcli.Params{Fields: PLAN_ALL_FIELDS})
	a.Nil(err)

	p.Description = "Updated"
	p, err = p.Update(tok,mp)
	a.Nil(err)

	p, err = Get(tok, mp, &mashcli.Params{Fields: PLAN_ALL_FIELDS})
	a.Nil(err)

	a.True(p.Description == "Updated", "update failed")

	ShowAllPlans(tok, mp)

}

func TestDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var p *Plan

	ShowAllPlans(tok, mp)

	p, err = Get(tok, mp, &mashcli.Params{Fields: PLAN_ALL_FIELDS})

	a.Nil(err)

	newPackageName := "Delete Me"
	p.Id = ""
	p.Name = newPackageName
	p, err = p.Create(tok,mp)
	a.Nil(err)

	ShowAllPlans(tok, mp)

	err = Delete(tok, &MethodParams{PackageId:PACKAGE_ID,PlanId:p.Id})
	a.Nil(err)

	_, err = Get(tok, &MethodParams{PackageId:PACKAGE_ID,PlanId:p.Id}, &mashcli.Params{Fields: PLAN_ALL_FIELDS})
	a.NotNil(err)

	ShowAllPlans(tok, mp)

}

func TestCleardown(t *testing.T) {
	var c mashcli.Config

	json.Unmarshal(config, &c)
	tok, _ := c.FetchOAuthToken()

	ShowAllPlans(tok, mp)
	Delete(tok, &MethodParams{PackageId:PACKAGE_ID,PlanId:""})
	ShowAllPlans(tok, mp)
}

func TestShowAllPlans(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ShowAllPlans(tok,mp)
	a.Nil(err)

}

func TestShowPlan(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	err = ShowPlan(tok, mp)
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

	_, err = Import(tok, "/Users/markmussett/Desktop/plans-69faacca-e627-4a49-836a-229264f5083c-mashclitest1.json",mp)
	a.Nil(err)
	ShowAllPlans(tok,mp)
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

	ShowAllPlans(tok,mp)
}
