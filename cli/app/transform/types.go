package transform

type Transform struct {
	Document struct {
		Endpoints []struct {
			InboundSslRequired        bool   `json:"inboundSslRequired"`
			Name                      string `json:"name"`
			OutboundRequestTargetPath string `json:"outboundRequestTargetPath"`
			OutboundTransportProtocol string `json:"outboundTransportProtocol"`
			PublicDomains             []struct {
				Address string `json:"address"`
			} `json:"publicDomains"`
			RequestAuthenticationType string   `json:"requestAuthenticationType"`
			RequestPathAlias          string   `json:"requestPathAlias"`
			RequestProtocol           string   `json:"requestProtocol"`
			SupportedHTTPMethods      []string `json:"supportedHttpMethods"`
			SystemDomains             []struct {
				Address string `json:"address"`
			} `json:"systemDomains"`
			TrafficManagerDomain string `json:"trafficManagerDomain"`
		} `json:"endpoints"`
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"document"`
	FeasibilityErrors []interface{} `json:"feasibilityErrors"`
	ValidationErrors  []interface{} `json:"validationErrors"`
}


type Services struct {
	Id        string      `json:"id,omitempty"`
	Name      string      `json:"name"`
	Version   string      `json:"version"`
	Endpoints []Endpoints `json:"endpoints,omitempty"`
	Updated   string      `json:"updated,omitempty"`
	Created   string      `json:"created,omitempty"`
}

type Endpoints struct {
	Name                      string          `json:"name"`
	OutboundRequestTargetPath string          `json:"outboundRequestTargetPath"`
	OutboundTransportProtocol string          `json:"outboundTransportProtocol"`
	PublicDomains             []PublicDomains `json:"publicDomains"`
	RequestAuthenticationType string          `json:"requestAuthenticationType"`
	RequestPathAlias          string          `json:"requestPathAlias"`
	RequestProtocol           string          `json:"requestProtocol"`
	SupportedHttpMethods      []string        `json:"supportedHttpMethods"`
	SystemDomains             []SystemDomains `json:"systemDomains"`
	TrafficManagerDomain      string          `json:"trafficManagerDomain"`
	InboundSslRequired        bool            `json:"inboundSslRequired"`
}

type PublicDomains struct {
	Address string `json:"address"`
}

type SystemDomains struct {
	Address string `json:"address"`
}

type Params struct {
	SourceFormat string `url:"sourceFormat,omitempty"`
	TargetFormat string `url:"targetFormat,omitempty"`
	PublicDomain string `url:"publicDomain,omitempty"`
}

type CreateServiceResponse struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Version string `json:"version"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}


type MethodParams struct {
	ServiceId string
}

type IoDocs struct {
	IoDocs interface{}
}
