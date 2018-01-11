package memberapplications

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

var sample = []byte(`
{
				"name": "Test Application",
        "description": "Test Application",
        "username": "mmussett"
    }
`)

const (
	MEMBER_ID      = "0150c6b5-0f7c-4a50-9876-e2048abda1c6"
	APPLICATION_ID = "f60a07e7-d3e0-4fec-bb4b-8fb5225a7820"
)

var mp = &MethodParams{memberId: MEMBER_ID, ApplicationId: APPLICATION_ID}

func TestGetCollection(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var mac *[]MemberApplications

	mac, err = GetCollection(tok, mp, &mashcli.Params{Fields: MEMBERAPPLICATIONS_ALL_FIELDS})
	a.Nil(err)

	spew.Dump(mac)

}

func TestCreate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var ma = new(MemberApplications)

	err = json.Unmarshal(sample, &ma)
	a.Nil(err)

	ShowMemberApplications(tok, mp)

	ma, err = ma.Create(tok, mp)
	a.Nil(err)

	ShowMemberApplications(tok, mp)

	spew.Dump(ma)
}

func TestUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	var ma = new(MemberApplications)
	err = json.Unmarshal(sample, &ma)
	ma, err = ma.Create(tok, mp)
	a.Nil(err)

	ShowMemberApplications(tok, &MethodParams{memberId: MEMBER_ID, ApplicationId: ma.Id})

	ma.Description = "Updated Description"
	ma, err = ma.Update(tok, &MethodParams{memberId: MEMBER_ID, ApplicationId: ma.Id})
	a.Nil(err)

	ShowMemberApplications(tok, &MethodParams{memberId: MEMBER_ID, ApplicationId: ma.Id})

	Delete(tok, &MethodParams{memberId: MEMBER_ID, ApplicationId: ma.Id})

}

func TestDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	// Create
	var ma = new(MemberApplications)
	err = json.Unmarshal(sample, &ma)
	ma, err = ma.Create(tok, mp)
	a.Nil(err)

	ShowMemberApplications(tok, mp)
	var mp = &MethodParams{memberId: MEMBER_ID, ApplicationId: ma.Id}
	ShowMemberApplications(tok, mp)

	err = Delete(tok, mp)
	a.Nil(err)

	ShowMemberApplications(tok, mp)
}

func TestCleardown(t *testing.T) {
	var c mashcli.Config
	json.Unmarshal(config, &c)
	tok, _ := c.FetchOAuthToken()
	ShowMemberApplications(tok, mp)
	Delete(tok, &MethodParams{memberId: "0150c6b5-0f7c-4a50-9876-e2048abda1c6", ApplicationId: "d26ede9c-21a8-4ae9-acca-0e604f9683d5"})
	ShowMemberApplications(tok, mp)
}

func TestShowAllApplications(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ShowMemberApplications(tok, mp)
	a.Nil(err)

}

func TestExportAll(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ExportAll(tok, "/Users/markmussett/Desktop", mp)
	a.Nil(err)
}

func TestImport(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	ShowMemberApplications(tok, mp)

	_, err = Import(tok, "/Users/markmussett/Desktop/0150c6b5-0f7c-4a50-9876-e2048abda1c6-e36fc044-cf71-4bb5-a0e8-a16068b4fb11-ACME Application.json", mp)
	a.Nil(err)

	ShowMemberApplications(tok, mp)

}
