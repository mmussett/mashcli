package packagekeys

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

const (
	resourcePath = "v3/rest/packageKeys/%s"
	collectionResourcePath = "v3/rest/packageKeys"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*PackageKeys, error) {

	path := fmt.Sprintf(resourcePath,mp.PackageKeyId)
	e := new(mashcli.MasheryError)
	a := new(PackageKeys)

	_, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(a, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode == 404 {
		return nil, errors.Errorf("packagekeys: unable to get package key: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return a, nil
}

func GetCollection(accessToken string, params *mashcli.Params, filter *mashcli.Filter) (*[]PackageKeys, error) {

	path := collectionResourcePath
	e := new(mashcli.MasheryError)
	apk := new([]PackageKeys)

	_, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).QueryStruct(filter).Receive(apk, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode != 0 {
		return nil, fmt.Errorf("packagekeys: unable to get package key collection: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return apk, nil
}

func (a *PackageKeys) Update(accessToken string, mp *MethodParams) (*PackageKeys, error) {

	path := fmt.Sprintf(resourcePath,mp.PackageKeyId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(a).Set("Accept", "application/json").Set("Content-Type", "application/json").Put(path).Receive(a, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("packagekeys: unable to update package key: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return a, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	path := fmt.Sprintf(resourcePath,mp.PackageKeyId)
	r := new(DeletePackageKeysResponse)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if resp.StatusCode == 200 || resp.StatusCode == 404 {
		return nil
	}

	if err != io.EOF {
		return errors.Errorf("applications: unable to delete package key: DELETE %s", path)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		return errors.Errorf("applications: unable to delete package key: DELETE %s -> (%s %s)", path, resp.StatusCode, resp.Status)
	}

	return nil

}
