package methods

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

const (
	resourcePath           = "v3/rest/services/%s/endpoints/%s/methods/%s"
	collectionResourcePath = "v3/rest/services/%s/endpoints/%s/methods"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*Methods, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId, mp.MethodId)
	e := new(mashcli.MasheryError)
	s := new(Methods)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(s, e)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.Errorf("methods: unable to get method: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)

	}

	return s, nil
}

func GetCollection(accessToken string, mp *MethodParams, params *mashcli.Params) (*[]Methods, error) {

	path := fmt.Sprintf(collectionResourcePath, mp.ServiceId, mp.EndpointId)
	e := new(mashcli.MasheryError)
	p := new([]Methods)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(p, e)
	if err != nil {
		return nil, err
	}

	if e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, fmt.Errorf("methods: unable to get methods collection: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func (m *Methods) Create(accessToken string, mp *MethodParams) (*Methods, error) {

	path := fmt.Sprintf(collectionResourcePath, mp.ServiceId, mp.EndpointId)
	e := new(mashcli.MasheryError)

	m.Id = ""
	m.Created = ""
	m.Updated = ""

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Set("Accept", "application/json").Post(path).Receive(m, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("methods: unable to create method: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func (m *Methods) Update(accessToken string, mp *MethodParams) (*Methods, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId, m.Id)
	e := new(mashcli.MasheryError)

	m.Id = ""
	m.Created = ""
	m.Updated = ""

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Put(path).Receive(m, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("methods: unable to update method id=%s: PUT %s -> (%s %s)", m.Id, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	r := new(DeleteMethodsResponse)

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId, mp.MethodId)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if resp.StatusCode == 200 || resp.StatusCode == 404 {
		return nil
	}

	if err != io.EOF {
		return errors.Errorf("method: unable to delete method: DELETE %s", path)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		return errors.Errorf("method: unable to delete method: DELETE %s -> (%s %s)", path, resp.StatusCode, resp.Status)
	}

	return nil

}
