package members

import (
	"encoding/json"
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
	MEMBER_ID = "4c1af408-e7f9-43e5-a5a7-604220192bd3"
)

var mp = &MethodParams{memberId: MEMBER_ID}

func TestGet(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	m, err := Get(tok, mp, &mashcli.Params{Fields: MEMBERS_ALL_FIELDS})

	a.Nil(err)

	a.True(m.Id == MEMBER_ID, "incorrect member returned")

	spew.Dump(m)

}

func TestGetCollection(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var m *[]Members

	m, err = GetCollection(tok, &mashcli.Params{Fields: MEMBERS_ALL_FIELDS})

	a.Nil(err)

	spew.Dump(m)

}

func TestCreate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var m *Members

	ShowAllMembers(tok)

	m, err = Get(tok, mp, &mashcli.Params{Fields: MEMBERS_ALL_FIELDS})
	a.Nil(err)

	m.Username = "clone_" + m.Username
	m.Email = "clone." + m.Email
	m.DisplayName = "clone " + m.DisplayName
	m.Id = ""
	m, err = m.Create(tok)
	a.Nil(err)

	ShowAllMembers(tok)

	var mp = &MethodParams{memberId: m.Id}
	err = Delete(tok, mp)
	a.Nil(err)

}

func TestUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	_, err = Get(tok, mp, &mashcli.Params{Fields: MEMBERS_ALL_FIELDS})

	m2 := new(Members)
	m2.AreaStatus = "disabled"

	m2, err = m2.Update(tok, mp)
	a.True(m2.AreaStatus == "disabled", "update failed")

	m3 := new(Members)
	m3.AreaStatus = "active"

	m3, err = m3.Update(tok, mp)
	a.True(m3.AreaStatus == "active", "update failed")
}

func TestDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var m *Members

	ShowAllMembers(tok)
	m, err = Get(tok, mp, &mashcli.Params{Fields: MEMBERS_ALL_FIELDS})
	a.Nil(err)

	newUsernameName := "Clone " + m.Username
	m.Id = ""
	m.Username = newUsernameName
	m, err = m.Create(tok)
	a.Nil(err)

	ShowAllMembers(tok)

	var mp = &MethodParams{memberId: m.Id}

	err = Delete(tok, mp)
	a.Nil(err)

	_, err = Get(tok, mp, &mashcli.Params{Fields: MEMBERS_ALL_FIELDS})
	a.NotNil(err)

	ShowAllMembers(tok)

}

func TestCleardown(t *testing.T) {
	var c mashcli.Config
	var mp = &MethodParams{memberId: "eace66a3-e9b5-4e5a-b1e2-2272f6505240"}

	json.Unmarshal(config, &c)
	tok, _ := c.FetchOAuthToken()

	ShowAllMembers(tok)
	Delete(tok, mp)
	ShowAllMembers(tok)
}

func TestShowAllMembers(t *testing.T) {

	assert := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	if err != nil {
		assert.Error(err)
	}

	err = ShowAllMembers(tok)

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

func TestExportAll(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ExportAll(tok, "/Users/markmussett/Desktop")
	a.Nil(err)
}

func TestImport(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ShowAllMembers(tok)

	_, err = Import(tok, "/Users/markmussett/Desktop/members.json", mp)
	a.Nil(err)

	err = ShowAllMembers(tok)
	a.Nil(err)

}

func TestShowMembers(t *testing.T) {

	assert := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	if err != nil {
		assert.Error(err)
	}

	err = ShowMembers(tok, mp)

	if err != nil {
		assert.Error(err)
	}

}
