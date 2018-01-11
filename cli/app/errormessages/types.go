package errormessages


type ErrorMessages struct {
	ErrorMessage []ErrorMessage `json:"errorMessages"`
}

type ErrorMessage struct {
Code         int    `json:"code"`
DetailHeader string `json:"detailHeader"`
ResponseBody string `json:"responseBody"`
Status       string `json:"status"`
}

type MethodParams struct {
	ServiceId  string
	errorSetId string
}

