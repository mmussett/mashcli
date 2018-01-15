package planservices

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

const (
	resourcePath = "v3/rest/packages/%s/plans/%s/services/%s"
	collectionResourcePath = "v3/rest/packages/%s/plans/%s/services"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*PlanServices, error) {

	path := fmt.Sprintf(resourcePath,mp.PackageId,mp.PlanId,mp.ServiceId)
	e := new(mashcli.MasheryError)
	p := new(PlanServices)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(p, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("planservices: unable to get plan services: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func GetCollection(accessToken string, mp *MethodParams, params *mashcli.Params, filter *mashcli.Filter) (*[]PlanServices, error) {

	path := fmt.Sprintf(collectionResourcePath, mp.PackageId,mp.PlanId)
	e := new(mashcli.MasheryError)
	p := new([]PlanServices)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).QueryStruct(filter).Receive(p, e)

	if err != nil || e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, fmt.Errorf("planservices: unable to get plan services collection: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func (p *PlanServices) Create(accessToken string,mp *MethodParams) (*PlanServices, error) {

	path := fmt.Sprintf(collectionResourcePath, mp.PackageId,mp.PlanId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(p).Set("Content-Type", "application/json").Post(path).Receive(p, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("planservices: unable to create plan services: POST %s%s -> (%s %s)", mashcli.BaseURL, resourcePath, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil

}

func (p *PlanServices) Update(accessToken string, mp *MethodParams) (*PlanServices, error) {

	path := fmt.Sprintf(resourcePath,mp.PackageId,mp.PlanId,mp.ServiceId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(p).Set("Content-Type", "application/json").Put(path).Receive(p, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("planservices: unable to update plan services: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	path := fmt.Sprintf(resourcePath,mp.PackageId,mp.PlanId,mp.ServiceId)
	r := new(DeletePlanResponse)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if err != io.EOF {
		return errors.Errorf("planservices: unable to delete plan: DELETE %s", path)
	}

	if resp.StatusCode != 200 || resp.StatusCode != 404 {
		return errors.Errorf("planservices: unable to delete plan services: DELETE %s -> (%s %s)", path, resp.StatusCode, resp.Status)
	}

	return nil

}
