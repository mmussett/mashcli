package domains

type Domains struct {
	Id      string `json:"id,omitempty"`
	Created string `json:"created,omitempty"`
	Domain  string `json:"domain,omitempty"`
	Status  string `json:"status,omitempty"`
}

type MethodParams struct {
	DomainId string
}
