package endpointcache

import (
	"io"
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"strconv"
)

const (
	resourcePath = "v3/rest/services/%s/endpoints/%s"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*EndpointCache, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)
	e := new(mashcli.MasheryError)
	s := new(EndpointCache)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(s, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("endpointcache: unable to get endpoint cache : GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return s, nil
}

func (m *EndpointCache) Create(accessToken string, mp *MethodParams) (*EndpointCache, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Set("Accept", "application/json").Put(path).Receive(m, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("endpointcache: unable to create endpoint cache: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func (m *EndpointCache) Update(accessToken string, mp *MethodParams) (*EndpointCache, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)
	e := new(mashcli.MasheryError)
	res := new(UpdateCacheResponse)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Put(path).Receive(res, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("endpointcache: unable to update endpoint cache: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)
	e := new(mashcli.MasheryError)
	res := new(UpdateCacheResponse)

	m := new(EndpointCache)
	m.Cache.CacheTTLOverride = 0
	m.Cache.VaryHeaderEnabled = false
	m.Cache.RespondFromStaleCacheEnabled = false
	m.Cache.ClientSurrogateControlEnabled = false
	m.Cache.IncludeAPIKeyInContentCacheKey = false
	m.Cache.ResponseCacheControlEnabled = false

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Put(path).Receive(res, e)

	if err != io.EOF {
		return errors.Errorf("endpointcache: unable to delete endpoint cache: DELETE %s", path)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		return errors.Errorf("endpointcache: unable to delete endpoint cache: DELETE %s -> (%s %s)", path, resp.StatusCode, resp.Status)
	}

	return nil

}
