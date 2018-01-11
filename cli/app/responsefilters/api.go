package responsefilters

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

const (
	resourcePath           = "v3/rest/services/%s/endpoints/%s/methods/%s/responseFilters/%s"
	collectionResourcePath = "v3/rest/services/%s/endpoints/%s/methods/%s/responseFilters"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*ResponseFilters, error) {

	e := new(mashcli.MasheryError)
	s := new(ResponseFilters)

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId, mp.MethodId, mp.ResponseFilterId)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(s, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("responsefilters: unable to get response filter id=%s: GET %s -> (%s %s)", mp.MethodId, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return s, nil
}

func GetCollection(accessToken string, mp *MethodParams, params *mashcli.Params) (*[]ResponseFilters, error) {

	path := fmt.Sprintf(collectionResourcePath, mp.ServiceId, mp.EndpointId, mp.MethodId)
	e := new(mashcli.MasheryError)
	p := new([]ResponseFilters)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(p, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, fmt.Errorf("responsefilters: unable to get response filter collection: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func (m *ResponseFilters) Create(accessToken string, mp *MethodParams) (*ResponseFilters, error) {

	path := fmt.Sprintf(collectionResourcePath, mp.ServiceId, mp.EndpointId, mp.MethodId)
	e := new(mashcli.MasheryError)

	m.Id = ""
	m.Created = ""
	m.Updated = ""

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Set("Accept", "application/json").Post(path).Receive(m, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("responsefilters: unable to create response filter: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func (m *ResponseFilters) Update(accessToken string, mp *MethodParams) (*ResponseFilters, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId, mp.MethodId, mp.ResponseFilterId)
	e := new(mashcli.MasheryError)

	m.Id = ""
	m.Created = ""
	m.Updated = ""

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Put(path).Receive(m, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("responsefilters: unable to update response filter id=%s: PUT %s -> (%s %s)", m.Id, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	r := new(DeleteResponseFiltersResponse)

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId, mp.MethodId, mp.ResponseFilterId)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if resp.StatusCode == 200 || resp.StatusCode == 404 {
		return nil
	}

	if err != io.EOF {
		return errors.Errorf("responsefilters: unable to delete response filter id=%s: DELETE %s", mp.ResponseFilterId, path, mp.ServiceId)
	}

	if resp.StatusCode != 200 || resp.StatusCode != 404 {
		return errors.Errorf("responsefilters: unable to delete response filter id=%s: DELETE %s -> (%s %s)", mp.MethodId, path, resp.StatusCode, resp.Status)
	}

	return nil

}
