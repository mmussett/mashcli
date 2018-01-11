package members

type Members struct {
	Address1           string        `json:"address1,omitempty"`
	Address2           string        `json:"address2,omitempty"`
	AreaStatus         string        `json:"areaStatus,omitempty"`
	Blog               string        `json:"blog,omitempty"`
	Company            string        `json:"company,omitempty"`
	CountryCode        string        `json:"countryCode,omitempty"`
	Created            string        `json:"created,omitempty"`
	DisplayName        string        `json:"displayName,omitempty"`
	Email              string        `json:"email,omitempty"`
	ExternalID         string        `json:"externalId,omitempty"`
	FirstName          string        `json:"firstName,omitempty"`
	Id                 string        `json:"id,omitempty"`
	Im                 string        `json:"im,omitempty"`
	Imsvc              string        `json:"imsvc,omitempty"`
	LastName           string        `json:"lastName,omitempty"`
	Locality           string        `json:"locality,omitempty"`
	PasswdNew          string        `json:"passwdNew,omitempty"`
	Phone              string        `json:"phone,omitempty"`
	PostalCode         string        `json:"postalCode,omitempty"`
	Region             string        `json:"region,omitempty"`
	RegistrationIpaddr string        `json:"registrationIpaddr,omitempty"`
	Roles              []interface{} `json:"roles,omitempty"`
	Updated            string        `json:"updated,omitempty"`
	URI                string        `json:"uri,omitempty"`
	Username           string        `json:"username,omitempty"`
}

type DeleteMemberResponse struct {
	Status string `json:"status"`
}

type MethodParams struct {
	MemberId string
}
