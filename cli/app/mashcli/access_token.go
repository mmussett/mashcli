package mashcli

import (
	"github.com/dghubble/sling"
	"strings"
)

type token struct {
	TokenType    string `json:"id"`
	MAPI         string `json:"mapi"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func (c *Config) FetchOAuthToken() (string, error) {

	token := new(token)

	body := strings.NewReader("grant_type=password&username=" + c.UserId + "&password=" + c.Password + "&scope=" + c.Area)

	_, err := sling.New().Base(BaseURL).SetBasicAuth(c.ApiKey, c.ApiKeySecret).Body(body).Set("Content-Type", "application/x-www-form-urlencoded").Post("v3/token").ReceiveSuccess(token)

	if err != nil {
		return "",err
	}

	return token.AccessToken, nil


}
