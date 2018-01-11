package responsefilters

type ResponseFilters struct {
	Id               string `json:"id,omitempty"`
	Created          string `json:"created,omitempty"`
	Updated          string `json:"updated,omitempty"`
	Name             string `json:"name,omitempty"`
	XmlFilterFields  string `json:"xmlFilterFields,omitempty"`
	Notes            string `json:"nodes,omitempty"`
	JsonFilterFields string `json:"jsonFilterFields,omitempty"`
}

type DeleteResponseFiltersResponse struct {
	Status string `json:"status"`
}

type MethodParams struct {
	ServiceId        string
	EndpointId       string
	MethodId         string
	ResponseFilterId string
}
