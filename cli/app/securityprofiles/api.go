package securityprofiles

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"strconv"
)

const (
	resourcePath = "v3/rest/services/%s"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*SecurityProfile, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId)
	e := new(mashcli.MasheryError)
	p := new(SecurityProfile)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(p, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("securityprofiles: unable to get security profiles: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func (m *SecurityProfile) Update(accessToken string, mp *MethodParams) (*SecurityProfile, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId)
	e := new(mashcli.MasheryError)

	//dump, _ := m.Marshall()
	//fmt.Printf("%v\n",dump)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json; charset=utf-8").Put(path).Receive(m, e)

	fmt.Printf("%v\n", resp.Status)

	if err != nil || resp.StatusCode == 404 {
		return nil, errors.Errorf("securityprofiles: unable to update security profiles: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}
