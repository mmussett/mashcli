package securityprofiles

type SecurityProfile struct {
	SecurityProfile struct {
		Oauth Oauth `json:"oauth,omitempty"`
	} `json:"securityProfile,omitempty"`
}

type Oauth struct {
	AccessTokenTTL              int      `json:"accessTokenTtl,omitempty"`
	AccessTokenTTLEnabled       bool     `json:"accessTokenTtlEnabled,omitempty"`
	AccessTokenType             string   `json:"accessTokenType,omitempty"`
	AllowMultipleToken          bool     `json:"allowMultipleToken,omitempty"`
	AuthorizationCodeTTL        int      `json:"authorizationCodeTtl,omitempty"`
	EnableRefreshTokenTTL       bool     `json:"enableRefreshTokenTtl,omitempty"`
	ForceOauthRedirectURL       bool     `json:"forceOauthRedirectUrl,omitempty"`
	ForceSslRedirectURLEnabled  bool     `json:"forceSslRedirectUrlEnabled,omitempty"`
	ForwardedHeaders            []string `json:"forwardedHeaders,omitempty"`
	GrantTypes                  []string `json:"grantTypes,omitempty"`
	MacAlgorithm                string   `json:"macAlgorithm,omitempty"`
	MasheryTokenAPIEnabled      bool     `json:"masheryTokenApiEnabled,omitempty"`
	QPSLimitCeiling             int      `json:"qpsLimitCeiling,omitempty"`
	RateLimitCeiling            int      `json:"rateLimitCeiling,omitempty"`
	RefreshTokenEnabled         bool     `json:"refreshTokenEnabled,omitempty"`
	RefreshTokenTTL             int      `json:"refreshTokenTtl,omitempty"`
	SecureTokensEnabled         bool     `json:"secureTokensEnabled,omitempty"`
	TokenBasedRateLimitsEnabled bool     `json:"tokenBasedRateLimitsEnabled,omitempty"`
}

type MethodParams struct {
	ServiceId string
}
