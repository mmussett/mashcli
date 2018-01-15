package applicationpackagekeys

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

const (
	resourcePath           = "v3/rest/applications/%s/packageKeys/%s"
	collectionResourcePath = "v3/rest/applications/%s/packageKeys"
)


func GetCollection(accessToken string, mp *MethodParams, params *mashcli.Params, filter *mashcli.Filter) (*[]ApplicationPackageKeys, error) {

	path := fmt.Sprintf(collectionResourcePath, mp.ApplicationId)

	e := new(mashcli.MasheryError)
	apk := new([]ApplicationPackageKeys)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).QueryStruct(filter).Receive(apk, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, fmt.Errorf("applicationpackagekey: unable to get application package key collection: GET %s%s -> (%s %s)", mashcli.BaseURL, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return apk, nil
}

func (a *ApplicationPackageKeys) Update(accessToken string, mp *MethodParams) (*ApplicationPackageKeys, error) {

	path := fmt.Sprintf(resourcePath, mp.ApplicationId, mp.PackageKeyId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(a).Set("Accept", "application/json").Set("Content-Type", "application/json").Put(path).Receive(a, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("applicationpackagekeys: unable to update application package key id=%s: PUT %s%s -> (%s %s)", mp.PackageKeyId, mashcli.BaseURL, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return a, nil

}

func (a *CreateApplicationPackageKeysRequest) Create(accessToken string, mp *MethodParams) (*ApplicationPackageKeys, error) {

	path := fmt.Sprintf(collectionResourcePath, mp.ApplicationId)
	e := new(mashcli.MasheryError)
	s := new(ApplicationPackageKeys)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(a).Set("Accept", "application/json").Set("Content-Type", "application/json").Post(path).Receive(s, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("applicationpackagekeys: unable to create application package key: POST %s%s -> (%s %s)", mashcli.BaseURL, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return s, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	path := fmt.Sprintf(resourcePath, mp.ApplicationId, mp.PackageKeyId)

	r := new(DeleteApplicationPackageKeysResponse)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if resp.StatusCode == 200 || resp.StatusCode == 404 {
		return nil
	}

	if err != io.EOF {
		return errors.Errorf("applications: unable to delete application package key id=%s: DELETE %s%s/%s/packageKeys/%s", mp.PackageKeyId, mashcli.BaseURL, path)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		return errors.Errorf("applications: unable to delete application package key id=%s: DELETE %s%s -> (%s %s)", mp.PackageKeyId, mashcli.BaseURL, path, resp.StatusCode, resp.Status)
	}

	return nil

}
