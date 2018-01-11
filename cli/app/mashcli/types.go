package mashcli

type MasheryError struct {
	ErrorCode    int         `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
	Errors       interface{} `json:"errors"`
}

type EmptyResponse struct {
}

type Params struct {
	Fields string `url:"fields,omitempty"`
}

type MasheryCreateResponse struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Version string `json:"version"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

type MasheryDeleteServiceResponse struct {
	Status string `json:"status"`
}

type Filter struct {
	Filter string `url:"filter,omitempty"`
}

type Config struct {
	// Mashery User ID
	UserId string `json:"userid""`
	// Mashery User Password
	Password string `json:"password"`
	// API Key
	ApiKey string `json:"apikey"`
	// API Key Secret
	ApiKeySecret string `json:"apikeysecret"`
	// Name of the configuration
	Name string `json:"name"`
	// Area UUID
	Area string `json:"area"`
	// Traffic Manager Host
	Tm string `json:"tm"`
	// Control Centre URL
	CcUrl string `json:"ccurl"`
}
