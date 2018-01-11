package systemdomains

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"strconv"
)

const (
	resourceCollectionPath = "v3/rest/domains/system"
)

func GetCollection(accessToken string, params *mashcli.Params) (*[]SystemDomains, error) {

	path := resourceCollectionPath
	e := new(mashcli.MasheryError)
	p := new([]SystemDomains)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(resourceCollectionPath).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(p, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, fmt.Errorf("systemdomains: unable to get system domains collection: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}
