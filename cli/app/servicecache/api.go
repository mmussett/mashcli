package servicecache

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

const (
	resourcePath = "v3/rest/services/%s/cache"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*ServiceCache, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId)
	e := new(mashcli.MasheryError)
	s := new(ServiceCache)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(s, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("servicecache: unable to get service cache : GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return s, nil
}

func (m *ServiceCache) Create(accessToken string, mp *MethodParams) (*ServiceCache, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Set("Accept", "application/json").Post(path).Receive(m, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("servicecache: unable to create service cache: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func (m *ServiceCache) Update(accessToken string, mp *MethodParams) (*ServiceCache, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Put(path).Receive(m, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("servicecache: unable to update service cache: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	path := fmt.Sprintf("v3/rest/services/%s", mp.ServiceId)
	e := new(mashcli.MasheryError)

	m := new(DeleteServiceCacheResponse)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Body(strings.NewReader("{\"cache\":{\"cacheTtl\":0}}")).Set("Content-Type", "application/json").Put(path).Receive(m, e)

	if err != nil || resp.StatusCode != 200 {
		return errors.Errorf("servicecache: unable to update service cache: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return nil

}