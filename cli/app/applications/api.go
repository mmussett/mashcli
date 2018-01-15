package applications

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

const (
	resourcePath           = "v3/rest/applications/%s"
	collectionResourcePath = "v3/rest/applications"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*Applications, error) {

	path := fmt.Sprintf(resourcePath, mp.ApplicationId)

	e := new(mashcli.MasheryError)
	a := new(Applications)

	_, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(a, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode == 404 {
		return nil, errors.Errorf("applications: unable to get application id=%s: GET %s%s -> (%s %s)", mp.ApplicationId, mashcli.BaseURL, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return a, nil
}

func GetCollection(accessToken string, params *mashcli.Params, filter *mashcli.Filter) (*[]Applications, error) {

	path := collectionResourcePath

	e := new(mashcli.MasheryError)
	a := new([]Applications)

	_, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).QueryStruct(filter).Receive(a, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode != 0 {
		return nil, fmt.Errorf("applications: unable to get applications collection: GET %s%s -> (%s %s)", mashcli.BaseURL, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return a, nil
}

func (a *Applications) Update(accessToken string, mp *MethodParams) (*Applications, error) {

	path := fmt.Sprintf(resourcePath, mp.ApplicationId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(a).Set("Content-Type", "application/json").Put(path).Receive(a, e)

	if err != nil {
		return nil, errors.Errorf("applications: unable to update application: PUT %s%s -> (%s %s)", mashcli.BaseURL, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	if resp.StatusCode != 200 {
		return nil, errors.Errorf("applications: unable to update applications id=%s: PUT %s -> (%s %s)", mp.ApplicationId, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return a, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	path := fmt.Sprintf(resourcePath, mp.ApplicationId)
	r := new(DeleteApplicationsResponse)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if resp.StatusCode == 200 || resp.StatusCode == 404 {
		return nil
	}

	if err != io.EOF {
		return errors.Errorf("applications: unable to delete application id=%s: DELETE %s%s", mp.ApplicationId, mashcli.BaseURL, path)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		return errors.Errorf("applications: unable to delete application id=%s: DELETE %s%s -> (%s %s)", mp.ApplicationId, mashcli.BaseURL, path, resp.StatusCode, resp.Status)
	}

	return nil

}
