package services

type Services struct {
	Id                string `json:"id,omitempty"`
	Created           string `json:"created,omitempty"`
	CrossdomainPolicy string `json:"crossdomainPolicy,omitempty"`
	Description       string `json:"description,omitempty"`
	EditorHandle      string `json:"editorHandle,omitempty"`
	Name              string `json:"name,omitempty"`
	QpsLimitOverall   int64  `json:"qpsLimitOverall,omitempty"`
	RevisionNumber    int64  `json:"revisionNumber,omitempty"`
	Rfc3986Encode     bool   `json:"rfc3986Encode,omitempty"`
	RobotsPolicy      string `json:"robotsPolicy,omitempty"`
	Roles             string `json:"roles,omitempty"`
	Updated           string `json:"updated,omitempty"`
	Version           string `json:"version,omitempty"`
}

type DeleteServiceResponse struct {
	Status string `json:"status"`
}

type MethodParams struct {
	ServiceId string
}
