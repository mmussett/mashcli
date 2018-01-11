package serviceroles


type ServiceRoles struct {
	Action  string `json:"action,omitempty"`
	Created string `json:"created,omitempty"`
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Updated string `json:"updated,omitempty"`
}

type DeleteServiceRolesResponse struct {
	Status string `json:"status"`
}

type MethodParams struct {
	ServiceId string
}
