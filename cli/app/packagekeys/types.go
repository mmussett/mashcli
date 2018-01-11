package packagekeys

type PackageKeys struct {
	Apikey           string `json:"apikey,omitempty"`
	Created          string `json:"created,omitempty"`
	Id               string `json:"id,omitempty"`
	QPSLimitCeiling  int    `json:"qpsLimitCeiling"`
	QPSLimitExempt   bool   `json:"qpsLimitExempt"`
	RateLimitCeiling int    `json:"rateLimitCeiling"`
	RateLimitExempt  bool   `json:"rateLimitExempt"`
	Secret           string `json:"secret,omitempty"`
	Status           string `json:"status,omitempty"`
	Updated          string `json:"updated,omitempty"`
	Limits           []struct {
		Period string `json:"period,omitempty"`
		Source string  `json:"source,omitempty"`
		Ceiling int `json:"ceiling,omitempty"`
	} `json:"limits,omitempty"`
}

type DeletePackageKeysResponse struct {
	Status string `json:"status"`
}

type MethodParams struct {
	PackageKeyId string
}
