package iodocs

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"strconv"
)

const (
	resourcePath           = "v3/rest/iodocs/services/%s"
)


func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*IoDocs, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId)
	e := new(mashcli.MasheryError)
	s := new(IoDocs)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(s, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, errors.Errorf("iodocs: unable to get iodocs id=%s: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return s, nil

}

