package endpoints

type Endpoints struct {
	Id                                         string        `json:"id,omitempty"`
	Created                                    string        `json:"created,omitempty"`
	Updated                                    string        `json:"updated,omitempty"`
	AllowMissingApiKey                         bool          `json:"allowMissingApiKey,omitempty"`
	ApiKeyValueLocationKey                     string        `json:"apiKeyValueLocationKey,omitempty"`
	ApiKeyValueLocations                       []string      `json:"apiKeyValueLocations,omitempty"`
	ApiMethodDetectionKey                      string        `json:"apiMethodDetectionKey,omitempty"`
	ApiMethodDetectionLocations                []string      `json:"apiMethodDetectionLocations,omitempty"`
	ConnectionTimeoutForSystemDomainRequest    int64         `json:"connectionTimeoutForSystemDomainRequest,omitempty"`
	ConnectionTimeoutForSystemDomainResponse   int64         `json:"connectionTimeoutForSystemDomainResponse,omitempty"`
	CookiesDuringHttpRedirectsEnabled          bool          `json:"cookiesDuringHttpRedirectsEnabled,omitempty"`
	CustomRequestAuthenticationAdapter         string        `json:"customRequestAuthenticationAdapter,omitempty"`
	DropApiKeyFromIncomingCall                 bool          `json:"dropApiKeyFromIncomingCall,omitempty"`
	ForceGzipOfBackendCall                     bool          `json:"forceGzipOfBackendCall,omitempty"`
	ForwardedHeaders                           []string      `json:"forwardedHeaders,omitempty"`
	GzipPassthroughSupportEnabled              bool          `json:"gzipPassthroughSupportEnabled,omitempty"`
	HeadersToExcludeFromIncomingCall           []string      `json:"headersToExcludeFromIncomingCall,omitempty"`
	HighSecurity                               bool          `json:"highSecurity,omitempty"`
	HostPassthroughIncludedInBackendCallHeader bool          `json:"hostPassthroughIncludedInBackendCallHeader,omitempty"`
	InboundSslRequired                         bool          `json:"inboundSslRequired,omitempty"`
	JsonpCallbackParameter                     string        `json:"jsonpCallbackParameter,omitempty"`
	JsonpCallbackParameterValue                string        `json:"jsonpCallbackParameterValue,omitempty"`
	Name                                       string        `json:"name,omitempty"`
	NumberOfHttpRedirectsToFollow              int64         `json:"numberOfHttpRedirectsToFollow,omitempty"`
	OauthGrantTypes                            []string      `json:"oauthGrantTypes,omitempty"`
	OutboundRequestTargetPath                  string        `json:"outboundRequestTargetPath,omitempty"`
	OutboundRequestTargetQueryParameters       string        `json:"outboundRequestTargetQueryParameters,omitempty"`
	OutboundTransportProtocol                  string        `json:"outboundTransportProtocol,omitempty"`
	Processor                                  interface{}   `json:"processor,omitempty"`
	PublicDomains                              PublicDomains `json:"publicDomains,omitempty"`
	RequestAuthenticationType                  string        `json:"requestAuthenticationType,omitempty"`
	RequestPathAlias                           string        `json:"requestPathAlias,omitempty"`
	RequestProtocol                            string        `json:"requestProtocol,omitempty"`
	ReturnedHeaders                            []string      `json:"returnedHeaders,omitempty"`
	ScheduledMaintenanceEvent                  string        `json:"scheduledMaintenanceEvent,omitempty"`
	SupportedHttpMethods                       []string      `json:"supportedHttpMethods,omitempty"`
	SystemDomains                              SystemDomains `json:"systemDomains,omitempty"`
	TrafficManagerDomain                       string        `json:"trafficManagerDomain,omitempty"`
	UseSystemDomainCredentials                 bool          `json:"useSystemDomainCredentials,omitempty"`
	HttpsClientProfile                         interface{}   `json:"httpsClientProfile,omitempty"`
}

type Processor struct {
	Adapter            string      `json:"adapter,omitEmpty"`
	PostInputs         interface{} `json:"postInputs,omitEmpty"`
	PostProcessEnabled bool        `json:"postProcessEnabled,omitEmpty"`
	PreInputs          interface{} `json:"preInputs,omitEmpty"`
	PreProcessEnabled  bool        `json:"preProcessEnabled,omitEmpty"`
}

type SystemDomains []struct {
	Address string `json:"address"`
}

type PublicDomains []struct {
	Address string `json:"address"`
}

type DeleteEndpointsResponse struct {
	Status string `json:"status"`
}

type MethodParams struct {
	ServiceId  string
	EndpointId string
}
