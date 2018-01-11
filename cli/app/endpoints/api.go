package endpoints

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

const (
	resourcePath           = "v3/rest/services/%s/endpoints/%s"
	collectionResourcePath = "v3/rest/services/%s/endpoints"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*Endpoints, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)
	e := new(mashcli.MasheryError)
	s := new(Endpoints)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(s, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, errors.Errorf("endpoints: unable to get endpoint id=%s: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return s, nil

}

func GetCollection(accessToken string, mp *MethodParams, params *mashcli.Params) (*[]Endpoints, error) {

	path := fmt.Sprintf(collectionResourcePath, mp.ServiceId)
	e := new(mashcli.MasheryError)
	p := new([]Endpoints)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(p, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, fmt.Errorf("endpoints: unable to get endpoint collection: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func (a *Endpoints) Create(accessToken string, mp *MethodParams) (*Endpoints, error) {

	path := fmt.Sprintf(collectionResourcePath, mp.ServiceId)
	e := new(mashcli.MasheryError)

	a.Id = ""
	a.Updated = ""
	a.Created = ""

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(a).Set("Content-Type", "application/json").Set("Accept", "application/json").Post(path).Receive(a, e)

	if err != nil || resp.StatusCode == 400 {
		return nil, errors.Errorf("endpoints: unable to create endpoint: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return a, nil

}

func (a *Endpoints) Update(accessToken string, mp *MethodParams) (*Endpoints, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)
	e := new(mashcli.MasheryError)

	a.Updated = ""
	a.Created = ""
	a.Id = ""

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(a).Set("Content-Type", "application/json").Put(path).Receive(a, e)

	if err != nil || resp.StatusCode == 404 {
		return nil, errors.Errorf("endpoints: unable to update endpoint %s: PUT %s -> (%s %s)", a.Id, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return a, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	r := new(DeleteEndpointsResponse)

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if err != io.EOF {
		return errors.Errorf("endpoints: unable to delete endpoint: DELETE %s", path)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		return errors.Errorf("endpoints: unable to delete endpoint: DELETE %s -> (%s %s)", path, resp.StatusCode, resp.Status)
	}

	return nil

}
