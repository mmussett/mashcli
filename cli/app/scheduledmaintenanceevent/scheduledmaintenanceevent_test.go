package scheduledmaintenanceevent

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
	SERVICE_ID  = "gykda32cdzsu257fyju9x52m"
	ENDPOINT_ID = "mkftqkvmjtum5ru73gsv8rh7"
)

var mp = &MethodParams{ServiceId: SERVICE_ID, EndpointId: ENDPOINT_ID}

func TestGet(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	m, err := Get(tok, mp, &mashcli.Params{Fields: SCHEDULEDMAINTENANCEEVENT_ALL_FIELDS})
	spew.Dump(err)

	a.Nil(err)

	spew.Dump(m)

}

func TestCreate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	m := new(ScheduledMaintenanceEvent)
	m.StartDateTime = "2018-01-01T03:29:41.000+0000"
	m.EndDateTime = "2018-01-01T04:29:41.000+0000"
	m.Name = "Test"
	m, err = m.Create(tok, mp)
	a.Nil(err)

	err = ShowScheduledMaintenanceEvent(tok, mp)

}

func TestUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var m *ScheduledMaintenanceEvent

	m, err = Get(tok, mp, &mashcli.Params{Fields: SCHEDULEDMAINTENANCEEVENT_ALL_FIELDS})
	a.Nil(err)

	m.PrettyPrint(mp)

	newName := m.Name + "_clone"
	m.Name = newName
	m, err = m.Update(tok, mp)
	a.Nil(err)

	m, err = Get(tok, mp, &mashcli.Params{Fields: SCHEDULEDMAINTENANCEEVENT_ALL_FIELDS})
	a.Nil(err)

	m.PrettyPrint(mp)

	a.True(m.Name == newName, "update failed")

}

func TestCleardown(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	Delete(tok, &MethodParams{ServiceId: SERVICE_ID, EndpointId: ENDPOINT_ID})

	ShowScheduledMaintenanceEvent(tok, mp)
}

func TestDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var m *ScheduledMaintenanceEvent

	err = ShowScheduledMaintenanceEvent(tok, mp)

	m, err = Get(tok, mp, &mashcli.Params{Fields: SCHEDULEDMAINTENANCEEVENT_ALL_FIELDS})

	a.Nil(err)

	m.Name = "clone" + m.Name
	m.Id = ""
	m, err = m.Create(tok, mp)
	a.Nil(err)

	err = ShowScheduledMaintenanceEvent(tok, mp)

	nmp := &MethodParams{ServiceId: SERVICE_ID, EndpointId: ENDPOINT_ID}
	err = Delete(tok, nmp)
	a.Nil(err)

	m, err = Get(tok, nmp, &mashcli.Params{Fields: SCHEDULEDMAINTENANCEEVENT_ALL_FIELDS})
	a.NotNil(err)

	err = ShowScheduledMaintenanceEvent(tok, mp)

}

func TestShowScheduledMaintenanceEvent(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	err = ShowScheduledMaintenanceEvent(tok, mp)

	a.Nil(err)

}
