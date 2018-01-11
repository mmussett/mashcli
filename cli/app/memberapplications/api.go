package memberapplications

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

const (
	resourcePath           = "v3/rest/members/%s/applications/%s"
	collectionResourcePath = "v3/rest/members/%s/applications"
)

func GetCollection(accessToken string, mp *MethodParams, params *mashcli.Params) (*[]MemberApplications, error) {

	path := fmt.Sprintf(collectionResourcePath, mp.MemberId)
	e := new(mashcli.MasheryError)
	a := new([]MemberApplications)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(a, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, fmt.Errorf("memberapplications: unable to get member applications collection: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return a, nil
}

func (s *MemberApplications) Create(accessToken string, mp *MethodParams) (*MemberApplications, error) {

	path := fmt.Sprintf(collectionResourcePath, mp.MemberId)
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Post(path).BodyJSON(s).Set("Content-Type", "application/json").Set("Accept", "application/json").Receive(s, e)

	if err != nil || resp.StatusCode == 400 {
		return nil, errors.Errorf("memberapplications: unable to create member application: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return s, nil
}

func (a *MemberApplications) Update(accessToken string, mp *MethodParams) (*MemberApplications, error) {

	path := fmt.Sprintf(resourcePath, mp.MemberId, mp.ApplicationId)
	e := new(mashcli.MasheryError)

	a.Updated = ""
	a.Created = ""
	a.Id = ""

	//dump, _ := a.Marshall()
	//fmt.Printf("%v\n",dump)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(a).Set("Content-Type", "application/json").Put(path).Receive(a, e)

	if err != nil || resp.StatusCode == 404 {
		return nil, errors.Errorf("memberapplications: unable to update member application id=%s: PUT %s -> (%s %s)", mp.ApplicationId, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return a, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	path := fmt.Sprintf(resourcePath, mp.MemberId, mp.ApplicationId)

	r := new(DeleteApplicationsResponse)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if err != io.EOF {
		return errors.Errorf("memberapplications: unable to delete member applications: DELETE %s", path)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		return errors.Errorf("memberapplications: unable to delete member applications: DELETE %s -> (%s %s)", path, resp.StatusCode, resp.Status)
	}

	return nil

}
