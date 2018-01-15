package services

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

const (
	resourcePath           = "v3/rest/services/%s"
	resourceCollectionPath = "v3/rest/services"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*Services, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId)
	e := new(mashcli.MasheryError)
	s := new(Services)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(s, e)

	if err != nil || resp.StatusCode == 404 {
		return nil, errors.Errorf("service: unable to get service id=%s: GET %s -> (%s %s)", mp.ServiceId, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)

	}

	return s, nil
}

func GetCollection(accessToken string, params *mashcli.Params, filter *mashcli.Filter) (*[]Services, error) {

	path := resourceCollectionPath
	e := new(mashcli.MasheryError)
	p := new([]Services)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).QueryStruct(filter).Receive(p, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, fmt.Errorf("service: unable to get service collection: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func (s *Services) Create(accessToken string) (*Services, error) {

	path := fmt.Sprintf(resourcePath, "")
	e := new(mashcli.MasheryError)

	s.Id = ""
	s.Updated = ""
	s.Created = ""

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(s).Set("Content-Type", "application/json").Set("Accept", "application/json").Post(path).Receive(s, e)

	if err != nil || resp.StatusCode == 400 {
		return nil, errors.Errorf("services: unable to create service: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return s, nil

}

func (s *Services) Update(accessToken string, mp *MethodParams) (*Services, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(s).Set("Content-Type", "application/json; charset=utf-8").Put(path).Receive(s, e)

	fmt.Printf("%v\n", resp.Status)

	if err != nil || resp.StatusCode == 404 {
		return nil, errors.Errorf("services: unable to update service id=%s: PUT %s -> (%s %s)", mp.ServiceId, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return s, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	path := fmt.Sprintf(resourcePath, mp.ServiceId)
	r := new(DeleteServiceResponse)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if resp.StatusCode == 200 || resp.StatusCode == 404 {
		return nil
	}

	if err != io.EOF && (resp.StatusCode != 200 || resp.StatusCode != 404) {
		return errors.Errorf("services: unable to delete service id=%s: DELETE %s", mp.ServiceId, path)
	}

	return nil

}
