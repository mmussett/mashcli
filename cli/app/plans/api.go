package plans

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

const (
	resourcePath = "v3/rest/packages/%s/plans/%s"
	collectionResourcePath = "v3/rest/packages/%s/plans"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*Plan, error) {

	path := fmt.Sprintf(resourcePath,mp.PackageId,mp.PlanId)
	e := new(mashcli.MasheryError)
	p := new(Plan)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(p, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("plans: unable to get plan: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func GetCollection(accessToken string, mp *MethodParams, params *mashcli.Params, filter *mashcli.Filter) (*[]Plan, error) {

	path := fmt.Sprintf(collectionResourcePath, mp.PackageId)
	e := new(mashcli.MasheryError)
	p := new([]Plan)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).QueryStruct(filter).Receive(p, e)

	if err != nil || e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, fmt.Errorf("plans: unable to get plan collection: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func (p *Plan) Create(accessToken string,mp *MethodParams) (*Plan, error) {

	path := fmt.Sprintf(collectionResourcePath, mp.PackageId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(p).Set("Content-Type", "application/json").Post(path).Receive(p, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("plans: unable to create plan: POST %s%s -> (%s %s)", mashcli.BaseURL, resourcePath, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil

}

func (p *Plan) Update(accessToken string, mp *MethodParams) (*Plan, error) {

	path := fmt.Sprintf(resourcePath,mp.PackageId,mp.PlanId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(p).Set("Content-Type", "application/json").Put(path).Receive(p, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("plans: unable to update plan: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	path := fmt.Sprintf(resourcePath,mp.PackageId,mp.PlanId)
	r := new(DeletePlanResponse)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if err != io.EOF {
		return errors.Errorf("plans: unable to delete plan: DELETE %s", path)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		return errors.Errorf("plans: unable to delete plan: DELETE %s -> (%s %s)", path, resp.StatusCode, resp.Status)
	}

	return nil

}
