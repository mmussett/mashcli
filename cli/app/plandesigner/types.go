package plandesigner

type ServiceType struct {
	Id        string       `json:"id,omitempty"`
	Name      string       `json:"name,omitempty"`
	Created   string       `json:"created,omitempty"`
	Updated   string       `json:"updated,omitempty"`
	Endpoints []EndpointType `json:"endpoints,omitempty"`
}

type ResponseFilterType struct {
	Id               string `json:"id,omitempty"`
	Created          string `json:"created,omitempty"`
	Updated          string `json:"updated,omitempty"`
	Name             string `json:"name,omitempty"`
	XmlFilterFields  string `json:"xmlFilterFields,omitempty"`
	Notes            string `json:"nodes,omitempty"`
	JsonFilterFields string `json:"jsonFilterFields,omitempty"`
}

type MethodType struct {
	Id               string             `json:"id,omitempty"`
	Created          string             `json:"created"`
	Updated          string             `json:"updated"`
	Name             string             `json:"name,omitempty"`
	ResponseFilter   interface{}        `json:"responseFilter"`
	RateLimitCeiling int                `json:"rateLimitCeiling"`
	RateLimitExempt  bool               `json:"rateLimitExempt"`
	RateLimitPeriod  string             `json:"rateLimitPeriod"`
	QPSLimitCeiling  int                `json:"qpsLimitCeiling"`
	QPSLimitExempt   bool               `json:"qpsLimitExempt"`
}

type EndpointType struct {
	Id                      string       `json:"id,omitempty"`
	Created                 string       `json:"created,omitempty"`
	Updated                 string       `json:"updated,omitempty"`
	Name                    string       `json:"name,omitempty"`
	Methods                 []MethodType `json:"methods,omitempty"`
	UndefinedMethodsAllowed bool         `json:"undefinedMethodsAllowed,omitempty"`
}

type PlanDesigner struct {
	Services []ServiceType `json:"services"`
}

type PlanDesignerResponse struct {
	Name                              string      `json:"name"`
	Id                                string      `json:"id"`
	Status                            string      `json:"status"`
	Description                       string      `json:"description"`
	Created                           string      `json:"created"`
	Updated                           string      `json:"updated"`
	MaxNumKeysAllowed                 int         `json:"maxNumKeysAllowed"`
	NumKeysBeforeReview               int         `json:"numKeysBeforeReview"`
	QPSLimitCeiling                   int         `json:"qpsLimitCeiling"`
	RateLimitCeiling                  int         `json:"rateLimitCeiling"`
	RateLimitPeriod                   string      `json:"rateLimitPeriod"`
	AdminKeyProvisioningEnabled       bool        `json:"adminKeyProvisioningEnabled"`
	QPSLimitExempt                    bool        `json:"qpsLimitExempt"`
	QPSLimitKeyOverrideAllowed        bool        `json:"qpsLimitKeyOverrideAllowed"`
	RateLimitExempt                   bool        `json:"rateLimitExempt"`
	RateLimitKeyOverrideAllowed       bool        `json:"rateLimitKeyOverrideAllowed"`
	ResponseFilterOverrideAllowed     bool        `json:"responseFilterOverrideAllowed"`
	SelfServiceKeyProvisioningEnabled bool        `json:"selfServiceKeyProvisioningEnabled"`
}

type MethodParams struct {
	PackageId string
	PlanId    string
	ServiceId string
}
