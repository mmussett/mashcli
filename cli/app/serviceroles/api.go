package serviceroles

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

const (
	resourcePath = "v3/rest/services/%s/roles"

)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*[]ServiceRoles, error) {

	path := fmt.Sprintf(resourcePath,mp.ServiceId)
	e := new(mashcli.MasheryError)
	p := new([]ServiceRoles)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(p, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, fmt.Errorf("service: unable to get service collection: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func (s *ServiceRoles) Create(accessToken string, mp *MethodParams) (*ServiceRoles, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId)
	e := new(mashcli.MasheryError)

	s.Id = ""
	s.Updated = ""
	s.Created = ""

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(s).Set("Content-Type", "application/json").Set("Accept", "application/json").Post(path).Receive(s, e)
	fmt.Printf("%v\n", resp.Status)

	if err != nil || resp.StatusCode == 400 {
		return nil, errors.Errorf("servicesroles: unable to create service roles: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return s, nil

}

func (s *ServiceRoles) Update(accessToken string, mp *MethodParams) (*ServiceRoles, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(s).Set("Content-Type", "application/json; charset=utf-8").Put(path).Receive(s, e)

	fmt.Printf("%v\n", resp.Status)

	if err != nil || resp.StatusCode == 404 {
		return nil, errors.Errorf("serviceroles: unable to update service roles: PUT %s -> (%s %s)", mp.ServiceId, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return s, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	path := fmt.Sprintf(resourcePath, mp.ServiceId)
	r := new(DeleteServiceRolesResponse)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if resp.StatusCode == 200 || resp.StatusCode == 404 {
		return nil
	}

	if err != io.EOF && (resp.StatusCode != 200 || resp.StatusCode != 404) {
		return errors.Errorf("serviceroles: unable to delete service roles: DELETE %s", mp.ServiceId, path)
	}

	return nil

}
