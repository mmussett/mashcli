package errorsets

type ErrorSets struct {
	ErrorSet []ErrorSet `json:"errorSets"`
}

type ErrorSet struct {
	ID        string `json:"id,omitempty"`
	Jsonp     bool   `json:"jsonp,omitempty"`
	JsonpType string `json:"jsonpType,omitempty"`
	Name      string `json:"name"`
	Type      string `json:"type,omitempty"`
}

type MethodParams struct {
	ServiceId  string
	errorSetId string
}

type DeleteErrorSetsResponse struct {
	Status string `json:"Success"`
}
