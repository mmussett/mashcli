package plandesigner

import (
	"encoding/json"
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
	PACKAGE_ID = "0f5d46bb-235f-4810-aebb-64e2669479cd"  // mashclitest_package1
	PLAN_ID = "8db24ee2-1ac0-4176-848a-213efefc31d4"  // mashclitest_plan1
	ACME_SERVICE_ID = "vx6f3xyyr99sahfwkbvw82dm" // ACME API
)



func Test(t *testing.T) {

	a := assert.New(t)
	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
  a.Nil(err)

	err = AddServiceToPackagePlan(tok, "vx6f3xyyr99sahfwkbvw82dm", PACKAGE_ID, PLAN_ID)
  a.Nil(err)

  err = AddServiceToPackagePlan(tok,"gykda32cdzsu257fyju9x52m",PACKAGE_ID, PLAN_ID)
  a.Nil(err)
}
