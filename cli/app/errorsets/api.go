package errorsets

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"strconv"
	"io"
)

const (
	resourcePath = "v3/rest/services/%s/errorSets/%s"
	collectionResourcePath = "v3/rest/services/%s"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*ErrorSets, error) {

	path := fmt.Sprintf(collectionResourcePath, mp.ServiceId)
	e := new(mashcli.MasheryError)
	s := new(ErrorSets)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(s, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("errorsets: unable to get service cache : GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return s, nil
}


func (m *ErrorSet) Create(accessToken string, mp *MethodParams) (*ErrorSet, error) {

	path := fmt.Sprintf("v3/rest/services/%s/errorSets", mp.ServiceId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Set("Accept", "application/json").Post(path).Receive(m, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("errorsets: unable to create error sets: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func (m *ErrorSets) CreateCollection(accessToken string, mp *MethodParams) (*ErrorSets, error) {

	path := fmt.Sprintf("v3/rest/services/%s/errorSets", mp.ServiceId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Set("Accept", "application/json").Post(path).Receive(m, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("errorsets: unable to create error sets: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}


func (m *ErrorSet) Update(accessToken string, mp *MethodParams) (*ErrorSet, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.errorSetId)
	e := new(mashcli.MasheryError)


	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Put(path).Receive(m, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("errorsets: unable to update error sets: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func (m *ErrorSets) UpdateCollection(accessToken string, mp *MethodParams) (*ErrorSets, error) {

	path := fmt.Sprintf(collectionResourcePath, mp.ServiceId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Put(path).Receive(m, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("errorsets: unable to update error sets: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	r := new(DeleteErrorSetsResponse)

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.errorSetId)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if resp.StatusCode == 200 || resp.StatusCode == 404 {
		return nil
	}

	if err != io.EOF {
		return errors.Errorf("errorsets: unable to delete error set: DELETE %s", path)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		return errors.Errorf("errorsets: unable to delete error set: DELETE %s -> (%s %s)", path, resp.StatusCode, resp.Status)
	}

	return nil

}
