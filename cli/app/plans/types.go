package plans

type Plan struct {
	Name                              string      `json:"name,omitempty"`
	Id                                string      `json:"id,omitempty"`
	Status                            string      `json:"status,omitempty"`
	Description                       string      `json:"description,omitempty"`
	Created                           string      `json:"created,omitempty"`
	Updated                           string      `json:"updated,omitempty"`
	MaxNumKeysAllowed                 int         `json:"maxNumKeysAllowed,omitempty"`
	NumKeysBeforeReview               int         `json:"numKeysBeforeReview,omitempty"`
	QPSLimitCeiling                   int         `json:"qpsLimitCeiling,omitempty"`
	RateLimitCeiling                  int         `json:"rateLimitCeiling,omitempty"`
	RateLimitPeriod                   string      `json:"rateLimitPeriod,omitempty"`
	AdminKeyProvisioningEnabled       bool        `json:"adminKeyProvisioningEnabled"`
	QPSLimitExempt                    bool        `json:"qpsLimitExempt,omitempty"`
	QPSLimitKeyOverrideAllowed        bool        `json:"qpsLimitKeyOverrideAllowed"`
	RateLimitExempt                   bool        `json:"rateLimitExempt,omitempty"`
	RateLimitKeyOverrideAllowed       bool        `json:"rateLimitKeyOverrideAllowed"`
	ResponseFilterOverrideAllowed     bool        `json:"responseFilterOverrideAllowed,omitempty"`
	SelfServiceKeyProvisioningEnabled bool        `json:"selfServiceKeyProvisioningEnabled"`
	Services                          interface{} `json:"services,omitempty"`
}

type DeletePlanResponse struct {
	Status string `json:"status"`
}

type MethodParams struct {
	PackageId string
	PlanId    string
}
