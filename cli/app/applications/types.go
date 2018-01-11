package applications

type Applications struct {
	Id                string `json:"id,omitempty"`
	Created           string `json:"created,omitempty"`
	Updated           string `json:"updated,omitempty"`
	Username          string `json:"username,omitempty"`
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	Type              string `json:"type,omitempty"`
	Commercial        bool   `json:"commercial,omitempty"`
	Ads               bool   `json:"ads,omitempty"`
	AdsSystem         string `json:"adsSystem,omitempty"`
	UsageModel        string `json:"usageModel,omitempty"`
	Tags              string `json:"tags,omitempty"`
	Notes             string `json:"notes,omitempty"`
	HowDidYouHear     string `json:"howDidYouHear,omitempty"`
	PreferredProtocol string `json:"preferredProtocol,omitempty"`
	PreferredOutput   string `json:"preferredOutput,omitempty"`
	ExternalID        string `json:"externalId,omitempty"`
	URI               string `json:"uri,omitempty"`
	Status            string `json:"status,omitempty"`
	IsPackaged        bool   `json:"isPackaged,omitempty"`
	OauthRedirectURI  string `json:"oauthRedirectUri,omitempty"`
}

type DeleteApplicationsResponse struct {
	Status string `json:"status"`
}

type MethodParams struct {
	ApplicationId string
}
