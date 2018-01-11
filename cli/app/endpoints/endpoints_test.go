package endpoints

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/mmussett/mashcli/cli/test"
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
	SERVICE_ID  = "x2dmsgggz5gdmesg5n5s8a8t"
	ENDPOINT_ID = "zx22ky4m2m47dpkw2njxrg95"
)

var testData = []byte(`
   {
    "id": "zx22ky4m2m47dpkw2njxrg95",
    "created": "2017-09-04T11:54:56.000+0000",
    "updated": "2017-09-05T14:57:51.000+0000",
    "allowMissingApiKey": false,
    "apiKeyValueLocationKey": "api_key",
    "apiKeyValueLocations": [
      "request-parameters",
      "request-body"
    ],
    "apiMethodDetectionKey": "0",
    "apiMethodDetectionLocations": [
      "request-path"
    ],
    "connectionTimeoutForSystemDomainRequest": 2,
    "connectionTimeoutForSystemDomainResponse": 2,
    "cookiesDuringHttpRedirectsEnabled": false,
    "customRequestAuthenticationAdapter": "",
    "dropApiKeyFromIncomingCall": false,
    "forceGzipOfBackendCall": false,
    "forwardedHeaders": [],
    "gzipPassthroughSupportEnabled": false,
    "headersToExcludeFromIncomingCall": [
      ""
    ],
    "highSecurity": false,
    "hostPassthroughIncludedInBackendCallHeader": false,
    "inboundSslRequired": false,
    "jsonpCallbackParameter": "",
    "jsonpCallbackParameterValue": "",
    "name": "URI2",
    "numberOfHttpRedirectsToFollow": 0,
    "oauthGrantTypes": [],
    "outboundRequestTargetPath": "/policy",
    "outboundRequestTargetQueryParameters": "",
    "outboundTransportProtocol": "http",
    "processor": {
      "adapter": "",
      "postInputs": {},
      "postProcessEnabled": false,
      "preInputs": {},
      "preProcessEnabled": false
    },
    "publicDomains": [
      {
        "address": "emealocal1.api.mashery.com"
      }
    ],
    "requestAuthenticationType": "apiKey",
    "requestPathAlias": "/test/mashcli",
    "requestProtocol": "rest",
    "returnedHeaders": [
      "mashery-responder"
    ],
    "scheduledMaintenanceEvent": "",
    "supportedHttpMethods": [
      "post",
      "get"
    ],
    "systemDomains": [
      {
        "address": "services.wideworldofacme.com"
      }
    ],
    "trafficManagerDomain": "emealocal1.api.mashery.com",
    "useSystemDomainCredentials": false
  }
`)

var mp = &MethodParams{ServiceId: SERVICE_ID, EndpointId: ENDPOINT_ID}

func TestGet(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	e, err := Get(tok, mp, &mashcli.Params{Fields: ENDPOINTS_ALL_FIELDS})
	spew.Dump(err)

	a.Nil(err)

	a.True(e.Id == ENDPOINT_ID, "incorrect endpoint returned")

	spew.Dump(e)

}

func TestGetCollection(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var app *[]Endpoints

	app, err = GetCollection(tok, mp, &mashcli.Params{Fields: ENDPOINTS_ALL_FIELDS})

	a.Nil(err)

	spew.Dump(app)

}

func TestShowAllEndpoints(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ShowAllEndpoints(tok, mp)
	a.Nil(err)

}

func TestShowEndpoints(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ShowEndpoints(tok, mp)

	a.Nil(err)

}

func TestEndpointsCreate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	var e Endpoints
	err = json.Unmarshal(testData, &e)
	a.Nil(err)
	e.PrettyPrint()

	ShowAllEndpoints(tok, mp)

	var f *Endpoints
	f, err = e.Create(tok, mp)
	a.Nil(err)
	f.PrettyPrint()

	ShowAllEndpoints(tok, mp)

}

func TestEndpointsUpdate(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	var e *Endpoints
	e, err = Get(tok, &MethodParams{ServiceId: SERVICE_ID, EndpointId: "sd3pbxzw7bpt2h8v5dveydee"}, &mashcli.Params{Fields: ENDPOINTS_ALL_FIELDS})
	e.PrettyPrint()

	e.Name = "Test_" + e.Name
	e, err = e.Update(tok, &MethodParams{ServiceId: SERVICE_ID, EndpointId: "sd3pbxzw7bpt2h8v5dveydee"})
	a.Nil(err)

	ShowAllEndpoints(tok, mp)
}

func TestEndpointsDelete(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	ShowAllEndpoints(tok, mp)

	err = Delete(tok, &MethodParams{ServiceId: SERVICE_ID, EndpointId: "hhakzc5ezha63w6wjkc6hey6"})
	a.Nil(err)

	ShowAllEndpoints(tok, mp)
}

func TestCleardown(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	Delete(tok, &MethodParams{ServiceId: SERVICE_ID, EndpointId: "hhakzc5ezha63w6wjkc6hey6"})
	Delete(tok, &MethodParams{ServiceId: SERVICE_ID, EndpointId: "kwvbmjxhtvry6q6npr8b3umq"})
	Delete(tok, &MethodParams{ServiceId: SERVICE_ID, EndpointId: "zjxsbc4nwhuney7psxvmutx6"})
	Delete(tok, &MethodParams{ServiceId: SERVICE_ID, EndpointId: "n394425cvnhgkqxst8kk9qxv"})
	Delete(tok, &MethodParams{ServiceId: SERVICE_ID, EndpointId: "cawajuwad4h7aq6bj9k26mm3"})
	Delete(tok, &MethodParams{ServiceId: SERVICE_ID, EndpointId: "juajpmju6bzxqyf4qds8uf4a"})
	Delete(tok, &MethodParams{ServiceId: SERVICE_ID, EndpointId: "cm8yd9usatq546kc8azv2e5a"})
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

	err = ExportAll(tok, "/Users/markmussett/Desktop", mp)
	a.Nil(err)
}

func TestImport(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	ShowAllEndpoints(tok, mp)

	var e *Endpoints
	e, err = Get(tok, mp, &mashcli.Params{Fields: ENDPOINTS_ALL_FIELDS})
	e.RequestPathAlias = "/" + testutils.SecureRandomAlphaString(10)
	e.Id = ""
	e.Created = ""
	e.Updated = ""
	e.WriteFile("/Users/markmussett/Desktop/endpoints-import-test.json")

	e, err = Import(tok, "/Users/markmussett/Desktop/endpoints-import-test.json", mp)
	a.Nil(err)

	ShowAllEndpoints(tok, mp)

	Delete(tok, &MethodParams{ServiceId: SERVICE_ID, EndpointId: e.Id})

	ShowAllEndpoints(tok, mp)

}
