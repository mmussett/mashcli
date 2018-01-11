package cors

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

const (
	resourcePath = "v3/rest/services/%s/endpoints/%s/cors"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*Cors, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)
	e := new(mashcli.MasheryError)
	s := new(Cors)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(s, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("cors: unable to get cors : GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return s, nil
}

func (m *Cors) Create(accessToken string, mp *MethodParams) (*Cors, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Set("Accept", "application/json").Post(path).Receive(m, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("cors: unable to create cors: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func (m *Cors) Update(accessToken string, mp *MethodParams) (*Cors, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Put(path).Receive(m, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("cors: unable to update cors: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	r := new(DeleteCorsResponse)

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if resp.StatusCode == 200 || resp.StatusCode == 404 {
		return nil
	}

	if err != io.EOF {
		return errors.Errorf("cors: unable to delete cors: DELETE %s", path)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		return errors.Errorf("cors: unable to delete cors: DELETE %s -> (%s %s)", path, resp.StatusCode, resp.Status)
	}

	return nil

}
