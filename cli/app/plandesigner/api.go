package plandesigner

import (
	"fmt"
	"strconv"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
)

const (
	resourcePath = "v3/rest/packages/%s/plans/%s"
)

func (req *PlanDesigner) Update(accessToken string, mp *MethodParams) (*PlanDesignerResponse, error) {

	path := fmt.Sprintf(resourcePath,mp.PackageId,mp.PlanId)
	e := new(mashcli.MasheryError)
	res := new(PlanDesignerResponse)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(req).Set("Content-Type", "application/json").Put(path).Receive(res, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("plandesigner: unable to update plan: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return res, nil

}

func Delete() error {

	return nil
}