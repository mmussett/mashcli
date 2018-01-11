package methods

type Methods struct {
	Id                 string `json:"id,omitempty"`
	Created            string `json:"created,omitempty"`
	Name               string `json:"name,omitempty"`
	SampleJsonResponse string `json:"sampleJsonResponse,omitempty"`
	SampleXmlResponse  string `json:"sampleXmlResponse,omitempty"`
	Updated            string `json:"updated,omitempty"`
}

type DeleteMethodsResponse struct {
	Status string `json:"status"`
}

type MethodParams struct {
	ServiceId  string
	EndpointId string
	MethodId   string
}
