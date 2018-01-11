package roles

type Roles struct {
	Id           string        `json:"id,omitempty"`
	Name           string        `json:"name,omitempty"`
	Created            string        `json:"created,omitempty"`
	Updated            string        `json:"updated,omitempty"`
}

type DeleteRolesResponse struct {
	Status string `json:"status"`
}

type MethodParams struct {
	RoleId string
}
