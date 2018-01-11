package oauth

type OAuth struct {
	AccessTokenTtl              int64    `json:"accessTokenTtl"`
	AccessTokenTtlEnabled       bool     `json:"accessTokenTtlEnabled"`
	AccessTokenType             string   `json:"accessTokenType"`
	AllowMultipleToken          bool     `json:"allowMultipleToken"`
	AuthorizationCodeTtl        int64    `json:"authorizationCodeTtl"`
	EnableRefreshTokenTtl       bool     `json:"enableRefreshTokenTtl"`
	ForceOauthRedirectURL       bool     `json:"forceOauthRedirectUrl"`
	ForceSslRedirectURLEnabled  bool     `json:"forceSslRedirectUrlEnabled"`
	ForwardedHeaders            []string `json:"forwardedHeaders"`
	GrantTypes                  []string `json:"grantTypes"`
	MacAlgorithm                string   `json:"macAlgorithm"`
	MasheryTokenApiEnabled      bool     `json:"masheryTokenApiEnabled"`
	QpsLimitCeiling             int64    `json:"qpsLimitCeiling"`
	RateLimitCeiling            int64    `json:"rateLimitCeiling"`
	RefreshTokenEnabled         bool     `json:"refreshTokenEnabled"`
	RefreshTokenTtl             int64    `json:"refreshTokenTtl"`
	SecureTokensEnabled         bool     `json:"secureTokensEnabled"`
	TokenBasedRateLimitsEnabled bool     `json:"tokenBasedRateLimitsEnabled"`
}
