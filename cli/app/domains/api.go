package domains

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"strconv"
)

const (
	resourcePath           = "v3/rest/domains/%s"
	resourceCollectionPath = "v3/rest/domains"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*Domains, error) {

	path := fmt.Sprintf(resourcePath, mp.DomainId)
	e := new(mashcli.MasheryError)
	p := new(Domains)

	_, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(p, e)

	if err != nil || e.ErrorCode == 404 {
		return nil, errors.Errorf("domains: unable to get domains id=%s: GET %s -> (%s %s)", mp.DomainId, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)

	}

	return p, nil
}

func GetCollection(accessToken string, params *mashcli.Params) (*[]Domains, error) {

	path := resourceCollectionPath
	e := new(mashcli.MasheryError)
	p := new([]Domains)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(resourceCollectionPath).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(p, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, fmt.Errorf("domains: unable to get domains collection: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func (m *Domains) Create(accessToken string) (*Domains, error) {

	path := resourceCollectionPath
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Set("Accept", "application/json").Post(path).Receive(m, e)
	fmt.Printf("%v\n", resp.Status)

	if err != nil || resp.StatusCode == 400 {
		return nil, errors.Errorf("domains: unable to create domains: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}
