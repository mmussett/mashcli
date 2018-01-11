package iodocs

import (
	"errors"
	"github.com/franela/goreq"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"io/ioutil"
)

func FetchIODoc(accessToken string, ServiceId string) (string, error) {

	type Params struct {
		Fields string
	}

	params := Params{Fields: "definition, created, ServiceId, defaultApi"}

	request := goreq.Request{
		Uri:         mashcli.BaseURL + "/v3/rest/iodocs/services/" + ServiceId,
		Method:      "GET",
		ContentType: "application/json",
		QueryString: params,
	}

	request.AddHeader("Authorization", "Bearer "+accessToken)

	res, err := request.Do()

	if err != nil {
		mashcli.Log.Error(err)
		return "", err
	} else {

		if res.StatusCode != 200 {
			return "", errors.New("Iodoc not found")
		}

		result, err := res.Body.ToString()
		if err != nil {
			return "", err
		}

		return result, nil

	}

}

func Export(accessToken string, ServiceId string, filename string) error {

	iodocAsString, err := FetchIODoc(accessToken, ServiceId)

	if err == nil {
		ioutil.WriteFile(filename, []byte(iodocAsString), 0644)
		return nil
	} else {
		mashcli.Log.Error(err)
		return err
	}
}

func Import(accessToken string, ServiceId string, filename string) error {

	return nil
}
