package systemdomainauthentication

type SystemDomainAuthentication struct {
	Type        string `json:"type,omitempty"`
	Username    string `json:"username,omitempty"`
	Certificate string `json:"certificate,omitempty"`
	Password    string `json:"password,omitempty"`
}

type MethodParams struct {
	ServiceId  string
	EndpointId string
}

type DeleteSystemDomainAuthenticationResponse struct {
	Created string `json:"created"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Updated string `json:"updated"`
}
