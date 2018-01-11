package securityprofiles

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/mmussett/mashcli/cli/app"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"

	"io/ioutil"

	"encoding/json"
)

func ShowSecurityProfile(accessToken string, mp *MethodParams) error {

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: SECURITYPROFILES_ALL_FIELDS})

	if err != nil {
		return err
	}

	m.PrettyPrint(mp)

	return nil

}

func (m *SecurityProfile) PrettyPrint(mp *MethodParams) {

	var o= &m.SecurityProfile.Oauth

	table := tablewriter.NewWriter(os.Stdout)
	caption := fmt.Sprintf("Service ID: %s\n", mp.ServiceId)
	var data []string
	var header []string

	header = append(header, "Grant Types", "Tok RateLim", "Calls/Sec", "RateLim Ceil", "Enable A/Token TTL", "A/Token TTL", "A/Token Type")
	data = append(data, app.FlattenStringArray(o.GrantTypes), app.FormatBool(o.TokenBasedRateLimitsEnabled), strconv.Itoa(o.QPSLimitCeiling), strconv.Itoa(o.RateLimitCeiling), app.FormatBool(o.TokenBasedRateLimitsEnabled), strconv.Itoa(o.AccessTokenTTL), o.AccessTokenType)
	if o.AccessTokenType == "mac" {
		header = append(header, "Mac Algo")
		data = append(data, o.MacAlgorithm)
	}

	header = append(header, "R/Token")
	data = append(data, app.FormatBool(o.RefreshTokenEnabled))

	if o.EnableRefreshTokenTTL {
		header = append(header, "R/Token TTL")
		data = append(data, strconv.Itoa(o.RefreshTokenTTL))
	}

	header = append(header,"Mash token API","Secure Token","Forward Headers","Force SSL Redirect", "Validate against Pre-Reg URL", "Allow Multiple Tokens", "Auth Code TTL")
	data = append(data,app.FormatBool(o.MasheryTokenAPIEnabled), app.FormatBool(o.SecureTokensEnabled), app.FlattenStringArray(o.ForwardedHeaders), app.FormatBool(o.ForceOauthRedirectURL), app.FormatBool(o.AllowMultipleToken), strconv.Itoa(o.AuthorizationCodeTTL))

	table.SetHeader(header)
	table.SetCaption(true, caption)
	table.Append(data)
	table.Render()

	return

}

func Import(accessToken, filename string, mp *MethodParams) (*SecurityProfile, error) {

	a, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}

	m, err := a.Update(accessToken, mp)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func Export(accessToken, path string, mp *MethodParams) error {

	valid, err := app.DirExists(path)
	if err != nil {
		return err
	}

	if !valid {
		return errors.New("Directory " + path + " does not exist")
	}

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: SECURITYPROFILES_ALL_FIELDS})
	if err != nil {
		return err
	}

	filename := path + "/securityprofile-" + mp.ServiceId + ".json"
	err = m.WriteFile(filename)
	if err != nil {
		return err
	}

	return nil
}

func (a *SecurityProfile) WriteFile(filename string) error {

	data, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil

}

func ReadFile(filename string) (*SecurityProfile, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(SecurityProfile)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (a *SecurityProfile) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
