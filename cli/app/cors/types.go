package cors

type Cors struct {
	AllDomainsEnabled        bool     `json:"allDomainsEnabled,omitempty"`
	CookiesAllowed           bool     `json:"cookiesAllowed,omitempty"`
	DomainsAllowed           []string `json:"domainsAllowed,omitempty"`
	HeadersAllowed           []string `json:"headersAllowed,omitempty"`
	HeadersExposed           []string `json"headersExposed,omitempty"`
	MaxAge                   int64    `json:"maxAge,omitempty"`
	SubDomainMatchingAllowed bool     `json:"subDomainMatchingAllowed,omitempty"`
}

type DeleteCorsResponse struct {
	Status string `json:"status"`
}

type MethodParams struct {
	ServiceId  string
	EndpointId string
}
