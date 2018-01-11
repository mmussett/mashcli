package members

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

const (
	resourcePath           = "v3/rest/members/%s"
	resourceCollectionPath = "v3/rest/members"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*Members, error) {

	path := fmt.Sprintf(resourcePath, mp.MemberId)
	e := new(mashcli.MasheryError)
	p := new(Members)

	_, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(p, e)

	if err != nil || e.ErrorCode == 404 {
		return nil, errors.Errorf("members: unable to get member id=%s: GET %s -> (%s %s)", mp.MemberId, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)

	}

	return p, nil
}

func GetCollection(accessToken string, params *mashcli.Params) (*[]Members, error) {

	path := resourceCollectionPath
	e := new(mashcli.MasheryError)
	p := new([]Members)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(resourceCollectionPath).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(p, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, fmt.Errorf("members: unable to get members collection: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func (m *Members) Create(accessToken string) (*Members, error) {

	path := fmt.Sprintf(resourcePath, "")
	e := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json").Set("Accept", "application/json").Post(path).Receive(m, e)
	fmt.Printf("%v\n", resp.Status)

	if err != nil || resp.StatusCode == 400 {
		return nil, errors.Errorf("members: unable to create member: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func (m *Members) Update(accessToken string, mp *MethodParams) (*Members, error) {

	path := fmt.Sprintf(resourcePath, mp.MemberId)
	e := new(mashcli.MasheryError)

	//dump, _ := m.Marshall()
	//fmt.Printf("%v\n",dump)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(m).Set("Content-Type", "application/json; charset=utf-8").Put(path).Receive(m, e)

	if err != nil || resp.StatusCode == 404 {
		return nil, errors.Errorf("members: unable to update members id=%s: PUT %s -> (%s %s)", mp.MemberId, path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	path := fmt.Sprintf(resourcePath, mp.MemberId)
	r := new(DeleteMemberResponse)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if err != io.EOF {
		return errors.Errorf("members: unable to delete member: DELETE %s", path)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		return errors.Errorf("members: unable to delete member: DELETE %s -> (%s %s)", path, resp.StatusCode, resp.Status)
	}

	return nil

}
