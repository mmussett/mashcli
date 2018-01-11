package systemdomainauthentication

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

const (
	resourcePath = "v3/rest/services/%s/endpoints/%s/systemDomainAuthentication"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*SystemDomainAuthentication, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)
	e := new(mashcli.MasheryError)
	s := new(SystemDomainAuthentication)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(s, e)

	if err != nil || e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, errors.Errorf("systemdomainauthentication: unable to get system domain authentication : GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return s, nil
}

func (m *SystemDomainAuthentication) Create(accessToken string, mp *MethodParams) (*SystemDomainAuthentication, error) {

	e := new(mashcli.MasheryError)

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Set("Accept", "application/json").Post(path).Receive(m, e)

	if err != nil || resp.StatusCode == 400 || e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, errors.Errorf("systemdomainauthentication: unable to create system domain authentication: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func (m *SystemDomainAuthentication) Update(accessToken string, mp *MethodParams) (*SystemDomainAuthentication, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Put(path).Receive(m, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("systemdomainauthentication: unable to update system domain authentication: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	path := fmt.Sprintf("v3/rest/services/%s/endpoints/%s", mp.ServiceId, mp.EndpointId)
	e := new(mashcli.MasheryError)

	m := new(DeleteSystemDomainAuthenticationResponse)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Body(strings.NewReader("{\"systemDomainAuthentication\":null}")).Set("Content-Type", "application/json").Put(path).Receive(m, e)

	if err != nil || resp.StatusCode != 200 {
		return errors.Errorf("systemdomainauthentication: unable to update system domain authentication: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return nil

}
