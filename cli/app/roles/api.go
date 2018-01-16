package roles

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

const (
	resourcePath           = "v3/rest/roles/%s"
	resourceCollectionPath = "v3/rest/roles"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*Roles, error) {

	path := fmt.Sprintf(resourcePath, mp.RoleId)
	e := new(mashcli.MasheryError)
	p := new(Roles)

	_, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(p, e)

	if err != nil || e.ErrorCode == 404 {
		return nil, errors.Errorf("roles: unable to get role: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)

	}

	return p, nil
}

func GetCollection(accessToken string, params *mashcli.Params, filter *mashcli.Filter) (*[]Roles, error) {

	path := resourceCollectionPath
	e := new(mashcli.MasheryError)
	p := new([]Roles)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(resourceCollectionPath).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).QueryStruct(filter).Receive(p, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, fmt.Errorf("roles: unable to get roles collection: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func (m *Roles) Create(accessToken string) (*Roles, error) {

	path := resourceCollectionPath
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Set("Accept", "application/json").Post(path).Receive(m, e)

	if err != nil || resp.StatusCode == 400 {
		return nil, errors.Errorf("roles: unable to create role: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func (m *Roles) Update(accessToken string, mp *MethodParams) (*Roles, error) {

	path := fmt.Sprintf(resourcePath, mp.RoleId)
	e := new(mashcli.MasheryError)

	m.Id  = ""
	m.Created = ""
	m.Updated = ""

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json; charset=utf-8").Put(path).Receive(m, e)

	if err != nil || resp.StatusCode == 404 {
		return nil, errors.Errorf("roles: unable to update role: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	path := fmt.Sprintf(resourcePath, mp.RoleId)
	r := new(DeleteRolesResponse)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if err != io.EOF {
		return errors.Errorf("roles: unable to delete role: DELETE %s", path)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		return errors.Errorf("roles: unable to delete role: DELETE %s -> (%s %s)", path, resp.StatusCode, resp.Status)
	}

	return nil

}
