package oauth

import (
	"errors"
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"
	"github.com/op/go-logging"
	"os"
	"strconv"
)

func FetchOAuth(accessToken string, ServiceId string) (*OAuth, error) {

	type ErrorResponse struct {
		ErrorCode    int    `json:"errorCode"`
		ErrorMessage string `json:"errorMessage"`
	}

	params := &mashcli.Params{
		Fields: "",
	}

	oAuth := new(OAuth)
	masheryError := new(mashcli.MasheryError)

	resp, err := sling.New().Base(mashcli.BaseURL).Path("/v3/rest/services/"+ServiceId+"/securityProfile/oauth").Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(oAuth, masheryError)

	if mashcli.Log.IsEnabledFor(logging.DEBUG) {
		mashcli.Log.Debug(resp)
	}

	if err == nil {
		if masheryError.ErrorCode != 0 {
			mashcli.Log.Error("GET " + mashcli.BaseURL + "v3/rest/services/" + ServiceId + "/securityProfile/oauth -> (" + strconv.Itoa(masheryError.ErrorCode) + " " + masheryError.ErrorMessage + ")")
			return nil, errors.New("Security Profile not found")
		} else {

			return oAuth, nil
		}
	} else {
		mashcli.Log.Error(err)
		return nil, err
	}
}

func (o *OAuth) PrettyPrint() {

	if o != nil {

		data := []string{strconv.FormatInt(o.QpsLimitCeiling, 10), strconv.FormatInt(o.RateLimitCeiling, 10), strconv.FormatInt(o.AccessTokenTtl, 10), o.AccessTokenType, strconv.FormatInt(o.AuthorizationCodeTtl, 10), app.FlattenStringArray(o.ForwardedHeaders), app.FlattenStringArray(o.GrantTypes), o.MacAlgorithm}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"QPS Limit Ceiling", "Rate Limit Ceiling", "Access Token TTL", "Access Token Type", "Authorization Code TTL", "Forwarded Headers", "Grant Types", "MAC Algorithm"})
		table.Append(data)
		table.Render()

	}
}

func PrintOAuth(oAuth *OAuth) {

	if oAuth == nil {
	} else {
		fmt.Println("  OAuth")
		fmt.Println("  -----")
		fmt.Printf("    %-31s  :  %d\n", "QPS Limit Ceiling", oAuth.QpsLimitCeiling)
		fmt.Printf("    %-31s  :  %d\n", "Rate Limit Ceiling", oAuth.RateLimitCeiling)
		fmt.Printf("    %-31s  :  %d\n", "Access Token TTL", oAuth.AccessTokenTtl)
		fmt.Printf("    %-31s  :  %s\n", "Access Token Type", oAuth.AccessTokenType)
		fmt.Printf("    %-31s  :  %d\n", "Authorization Code TTL", oAuth.AuthorizationCodeTtl)
		fmt.Printf("    %-31s  :  %s\n", "Forwarded Headers", app.FlattenStringArray(oAuth.ForwardedHeaders))
		fmt.Printf("    %-31s  :  %s\n", "Oauth.GrantTypes", app.FlattenStringArray(oAuth.GrantTypes))
		fmt.Printf("    %-31s  :  %s\n", "MAC Algorithm", oAuth.MacAlgorithm)
		fmt.Printf("    %-31s  :  %d\n", "Refresh Token TTL", oAuth.RefreshTokenTtl)
		fmt.Printf("    %-31s  :  %t\n", "Refresh Token Enabled", oAuth.RefreshTokenEnabled)
		fmt.Printf("    %-31s  :  %t\n", "Access Token TTL Enabled", oAuth.AccessTokenTtlEnabled)
		fmt.Printf("    %-31s  :  %t\n", "Allow Multiple Token", oAuth.AllowMultipleToken)
		fmt.Printf("    %-31s  :  %t\n", "Enable Refresh Token TTL", oAuth.EnableRefreshTokenTtl)
		fmt.Printf("    %-31s  :  %t\n", "Force Oauth Redirect URL", oAuth.ForceOauthRedirectURL)
		fmt.Printf("    %-31s  :  %t\n", "Mashery Token API Enabled", oAuth.MasheryTokenApiEnabled)
		fmt.Printf("    %-31s  :  %t\n", "Refresh Token Enabled", oAuth.RefreshTokenEnabled)
		fmt.Printf("    %-31s  :  %t\n", "Secure Tokens Enabled", oAuth.SecureTokensEnabled)
		fmt.Printf("    %-31s  :  %t\n", "Token Based Rate Limits Enabled", oAuth.TokenBasedRateLimitsEnabled)
	}

}
