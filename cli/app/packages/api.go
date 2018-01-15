package packages

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

const (
	resourcePath = "v3/rest/packages/%s"
	collectionResourcePath = "v3/rest/packages"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*Package, error) {

	path := fmt.Sprintf(resourcePath,mp.PackageId)
	e := new(mashcli.MasheryError)
	p := new(Package)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(p, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("package: unable to get package: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func GetCollection(accessToken string, params *mashcli.Params, filter *mashcli.Filter) (*[]Package, error) {

	path := collectionResourcePath
	e := new(mashcli.MasheryError)
	p := new([]Package)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).QueryStruct(filter).Receive(p, e)

	if err != nil || e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, fmt.Errorf("package: unable to get package collection: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func (p *Package) Create(accessToken string) (*Package, error) {

	path := collectionResourcePath
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(p).Set("Content-Type", "application/json").Post(path).Receive(p, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("package: unable to create package: POST %s%s -> (%s %s)", mashcli.BaseURL, resourcePath, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil

}

func (p *Package) Update(accessToken string, mp *MethodParams) (*Package, error) {

	path := fmt.Sprintf(resourcePath,mp.PackageId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(p).Set("Content-Type", "application/json").Put(path).Receive(p, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("package: unable to update package: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	path := fmt.Sprintf(resourcePath,mp.PackageId)
	r := new(DeletePackageResponse)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if err != io.EOF {
		return errors.Errorf("package: unable to delete package: DELETE %s", path)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		return errors.Errorf("package: unable to delete package: DELETE %s -> (%s %s)", path, resp.StatusCode, resp.Status)
	}

	return nil

}
