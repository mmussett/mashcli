package errormessages

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"strconv"
)

const (
	resourcePath = "v3/rest/services/%s/errorSets/%s"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*ErrorMessages, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId,mp.errorSetId)
	e := new(mashcli.MasheryError)
	s := new(ErrorMessages)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(s, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("errormessages: unable to get error messages : GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return s, nil
}


func Update(accessToken string, mp *MethodParams, em *ErrorMessages) (*ErrorMessages, error) {

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.errorSetId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(em).Set("Content-Type", "application/json").Put(path).Receive(em, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("errormessages: unable to update error messages: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return em, nil

}

