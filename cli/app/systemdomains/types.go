package systemdomains

type SystemDomains struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty""`
	Path    string `json:"path,omitempty""`
	Domain  string `json:"domain,omitempty""`
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
}

type MethodParams struct {
}
