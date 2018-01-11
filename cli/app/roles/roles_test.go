package roles

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
	ROLE_ID = "79a1a2f6-13e0-433a-87f9-0c50fd1ce2dc"
)

var mp = &MethodParams{RoleId: ROLE_ID}

func TestGet(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	m, err := Get(tok, mp, &mashcli.Params{Fields: ROLES_ALL_FIELDS})
	a.Nil(err)

	a.True(m.Id == ROLE_ID, "incorrect role returned")

	spew.Dump(m)

}

func TestGetCollection(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var m *[]Roles

	m, err = GetCollection(tok, &mashcli.Params{Fields: ROLES_ALL_FIELDS})

	a.Nil(err)

	spew.Dump(m)

}

func TestCreate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)


	ShowAllRoles(tok)

	m := new(Roles)
	m.Id = ""
	m.Name = "mashclirole"
	m, err = m.Create(tok)
	a.Nil(err)

	ShowAllRoles(tok)

	//var mp = &MethodParams{RoleId: m.Id}
	//err = Delete(tok, mp)
	//a.Nil(err)

}

func TestUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	r, err := Get(tok, mp, &mashcli.Params{Fields: ROLES_ALL_FIELDS})

	r.Name = "mashclirole-updated"

	r, err = r.Update(tok, mp)

	r, err = Get(tok, mp, &mashcli.Params{Fields: ROLES_ALL_FIELDS})
	a.Nil(err)

	a.True(r.Name == "mashclirole-updated", "update failed")

	ShowAllRoles(tok)

	r.Name = "mashclirole"
	r, err = r.Update(tok, mp)

}

func TestDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var m *Roles

	ShowAllRoles(tok)
	m, err = Get(tok, mp, &mashcli.Params{Fields: ROLES_ALL_FIELDS})
	a.Nil(err)

	m.Id = ""
	m.Name = "DeleteRole"
	m, err = m.Create(tok)
	a.Nil(err)

	ShowAllRoles(tok)

	var mp = &MethodParams{RoleId: m.Id}

	err = Delete(tok, mp)
	a.Nil(err)

	_, err = Get(tok, mp, &mashcli.Params{Fields: ROLES_ALL_FIELDS})
	a.NotNil(err)

	ShowAllRoles(tok)

}

func TestCleardown(t *testing.T) {
	var c mashcli.Config
	json.Unmarshal(config, &c)
	tok, _ := c.FetchOAuthToken()

	ShowAllRoles(tok)
	Delete(tok, &MethodParams{RoleId: "eace66a3-e9b5-4e5a-b1e2-2272f6505240"})
	ShowAllRoles(tok)
}

func TestShowAllRoles(t *testing.T) {

	assert := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	if err != nil {
		assert.Error(err)
	}

	err = ShowAllRoles(tok)

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

func TestImportAll(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ImportAll(tok, "/Users/markmussett/Desktop/roles.json")
	a.Nil(err)
}

func TestImport(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ShowAllRoles(tok)

	_, err = Import(tok, "/Users/markmussett/Desktop/roles-79a1a2f6-13e0-433a-87f9-0c50fd1ce2dc.json", mp)
	a.Nil(err)

	err = ShowAllRoles(tok)
	a.Nil(err)

}

func TestShowRoles(t *testing.T) {

	assert := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	if err != nil {
		assert.Error(err)
	}

	err = ShowRoles(tok, mp)

	if err != nil {
		assert.Error(err)
	}

}
