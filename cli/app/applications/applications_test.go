package applications

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

var testData = []byte(`{
        "ads": true,
        "adsSystem": "nullam porttitor",
        "commercial": false,
        "created": "2015-04-02T13:11:48.000+0000",
        "description": "consequat in consequat ut nulla sed accumsan felis ut at dolor quis odio consequat varius integer ac leo pellentesque ultrices",
        "externalId": "morbi vel lectus",
        "howDidYouHear": "tempus sit amet sem",
        "id": "de8eeec8-1687-428b-865b-cb048b808d2d",
        "name": "cras mi",
        "notes": "aenean lectus pellentesque",
        "oauthRedirectUri": "hac habitasse platea dictumst",
        "preferredOutput": "volutpat convallis",
        "preferredProtocol": "in hac",
        "tags": "dui nec nisi volutpat",
        "type": "enim blandit",
        "updated": "2015-03-25T06:02:14.000+0000",
        "uri": "potenti",
        "usageModel": "molestie",
        "username": "justo morbi ut odio"
    }`)

const (
	APPLICATION_ID = "d985067d-1568-4152-94f7-82acdbf4537d"
)

var mp = &MethodParams{ApplicationId: APPLICATION_ID}

func TestGet(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	p, err := Get(tok, mp, &mashcli.Params{Fields: APPLICATIONS_ALL_FIELDS})
	spew.Dump(err)

	a.Nil(err)

	a.True(p.Id == APPLICATION_ID, "incorrect application returned")

	spew.Dump(p)

}

func TestGetCollection(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var app *[]Applications

	app, err = GetCollection(tok, &mashcli.Params{Fields: APPLICATIONS_ALL_FIELDS})

	a.Nil(err)

	spew.Dump(app)

}

func TestUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	var app *Applications

	ShowAllApplications(tok)

	app, err = Get(tok, mp, &mashcli.Params{Fields: APPLICATIONS_ALL_FIELDS})
	a.Nil(err)

	app.Description = "Updated"
	app.Id = ""
	app.Created = ""
	app.Updated = ""
	app, err = app.Update(tok, mp)
	a.Nil(err)

	app, err = Get(tok, mp, &mashcli.Params{Fields: APPLICATIONS_ALL_FIELDS})
	a.Nil(err)

	a.True(app.Description == "Updated", "update failed")

	ShowAllApplications(tok)
}

func TestDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	_, err = Get(tok, &MethodParams{ApplicationId:"8a11f8bd-bd6c-47ec-b5a0-209b41eb111f"}, &mashcli.Params{Fields: APPLICATIONS_ALL_FIELDS})
	a.Nil(err)

	ShowAllApplications(tok)

	err = Delete(tok,&MethodParams{ApplicationId:"8a11f8bd-bd6c-47ec-b5a0-209b41eb111f"})
	a.Nil(err)

	ShowAllApplications(tok)

}

func TestShowAllApplications(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ShowAllApplications(tok)
	a.Nil(err)

}

func TestShowApplications(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	err = ShowApplication(tok, mp)

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

	err = ShowAllApplications(tok)

	err = Import(tok, "/Users/markmussett/Desktop/applications-d985067d-1568-4152-94f7-82acdbf4537d.json", mp)
	a.Nil(err)

	err = ShowAllApplications(tok)
	a.Nil(err)

}
