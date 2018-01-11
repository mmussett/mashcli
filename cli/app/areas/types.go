package areas

type ServicesCollection struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type Services struct {
	Id    string `json:"id"`
	Cache struct {
		CacheTtl int64 `json:"cacheTtl"`
	} `json:"cache"`
	Created           string `json:"created"`
	CrossdomainPolicy string `json:"crossdomainPolicy"`
	Description       string `json:"description"`
	EditorHandle      string `json:"editorHandle"`
	Endpoints         []struct {
		AllowMissingApiKey          bool     `json:"allowMissingApiKey"`
		ApiKeyValueLocationKey      string   `json:"apiKeyValueLocationKey"`
		ApiKeyValueLocations        []string `json:"apiKeyValueLocations"`
		ApiMethodDetectionKey       string   `json:"apiMethodDetectionKey"`
		ApiMethodDetectionLocations []string `json:"apiMethodDetectionLocations"`
		Cache                       struct {
			ClientSurrogateControlEnabled bool     `json:"clientSurrogateControlEnabled"`
			ContentCacheKeyHeaders        []string `json:"contentCacheKeyHeaders"`
		} `json:"cache"`
		ConnectionTimeoutForSystemDomainRequest  int64 `json:"connectionTimeoutForSystemDomainRequest"`
		ConnectionTimeoutForSystemDomainResponse int64 `json:"connectionTimeoutForSystemDomainResponse"`
		CookiesDuringHttpRedirectsEnabled        bool  `json:"cookiesDuringHttpRedirectsEnabled"`
		Cors                                     struct {
			AllDomainsEnabled        bool     `json:"allDomainsEnabled"`
			CookiesAllowed           bool     `json:"cookiesAllowed"`
			DomainsAllowed           []string `json:"domainsAllowed"`
			HeadersAllowed           []string `json:"headersAllowed"`
			HeadersExposed           []string `json:"headersExposed"`
			MaxAge                   int64    `json:"maxAge"`
			SubDomainMatchingAllowed bool     `json:"subDomainMatchingAllowed"`
		} `json:"cors"`
		CustomRequestAuthenticationAdapter         string   `json:"customRequestAuthenticationAdapter"`
		DropApiKeyFromIncomingCall                 bool     `json:"dropApiKeyFromIncomingCall"`
		ForceGzipOfBackendCall                     bool     `json:"forceGzipOfBackendCall"`
		ForwardedHeaders                           []string `json:"forwardedHeaders"`
		GzipPassthroughSupportEnabled              bool     `json:"gzipPassthroughSupportEnabled"`
		HeadersToExcludeFromIncomingCall           []string `json:"headersToExcludeFromIncomingCall"`
		HighSecurity                               bool     `json:"highSecurity"`
		HostPassthroughIncludedInBackendCallHeader bool     `json:"hostPassthroughIncludedInBackendCallHeader"`
		InboundSslRequired                         bool     `json:"inboundSslRequired"`
		JsonpCallbackParameter                     string   `json:"jsonpCallbackParameter"`
		JsonpCallbackParameterValue                string   `json:"jsonpCallbackParameterValue"`
		Methods                                    []struct {
			ID                 string `json:"id"`
			Created            string `json:"created"`
			Name               string `json:"name"`
			SampleJsonResponse string `json:"sampleJsonResponse"`
			SampleXmlResponse  string `json:"sampleXmlResponse"`
			Updated            string `json:"updated"`
		} `json:"methods"`
		Name                                 string   `json:"name"`
		NumberOfHttpRedirectsToFollow        int64    `json:"numberOfHttpRedirectsToFollow"`
		OauthGrantTypes                      []string `json:"oauthGrantTypes"`
		OutboundRequestTargetPath            string   `json:"outboundRequestTargetPath"`
		OutboundRequestTargetQueryParameters string   `json:"outboundRequestTargetQueryParameters"`
		OutboundTransportProtocol            string   `json:"outboundTransportProtocol"`
		Processor                            struct {
			Adapter    string `json:"adapter"`
			PostInputs struct {
			} `json:"postInputs"`
			PostProcessEnabled bool `json:"postProcessEnabled"`
			PreInputs          struct {
			} `json:"preInputs"`
			PreProcessEnabled bool `json:"preProcessEnabled"`
		} `json:"processor"`
		PublicDomains []struct {
			Address string `json:"address"`
		} `json:"publicDomains"`
		RequestAuthenticationType  string   `json:"requestAuthenticationType"`
		RequestPathAlias           string   `json:"requestPathAlias"`
		RequestProtocol            string   `json:"requestProtocol"`
		ReturnedHeaders            []string `json:"returnedHeaders"`
		ScheduledMaintenanceEvent  string   `json:"scheduledMaintenanceEvent"`
		SupportedHttpMethods       []string `json:"supportedHttpMethods"`
		SystemDomainAuthentication struct {
			Password string `json:"password"`
			Type     string `json:"type"`
			Username string `json:"username"`
		} `json:"systemDomainAuthentication"`
		SystemDomains []struct {
			Address string `json:"address"`
		} `json:"systemDomains"`
		TrafficManagerDomain       string `json:"trafficManagerDomain"`
		UseSystemDomainCredentials bool   `json:"useSystemDomainCredentials"`
	} `json:"endpoints"`
	ErrorSets []struct {
		ID            string `json:"id"`
		ErrorMessages []struct {
			ID           string `json:"id"`
			Code         int64  `json:"code"`
			DetailHeader string `json:"detailHeader"`
			ResponseBody string `json:"responseBody"`
			Status       string `json:"status"`
		} `json:"errorMessages"`
		Jsonp     bool   `json:"jsonp"`
		JsonpType string `json:"jsonpType"`
		Name      string `json:"name"`
		Type      string `json:"type"`
	} `json:"errorSets"`
	Name            string `json:"name"`
	QpsLimitOverall int64  `json:"qpsLimitOverall"`
	RevisionNumber  int64  `json:"revisionNumber"`
	Rfc3986Encode   bool   `json:"rfc3986Encode"`
	RobotsPolicy    string `json:"robotsPolicy"`
	Roles           string `json:"roles"`
	SecurityProfile struct {
		Oauth struct {
			AccessTokenTtl              int64    `json:"accessTokenTtl"`
			AccessTokenTtlEnabled       bool     `json:"accessTokenTtlEnabled"`
			AccessTokenType             string   `json:"accessTokenType"`
			AllowMultipleToken          bool     `json:"allowMultipleToken"`
			AuthorizationCodeTtl        int64    `json:"authorizationCodeTtl"`
			EnableRefreshTokenTtl       bool     `json:"enableRefreshTokenTtl"`
			ForceOauthRedirectURL       bool     `json:"forceOauthRedirectUrl"`
			ForceSslRedirectURLEnabled  bool     `json:"forceSslRedirectUrlEnabled"`
			ForwardedHeaders            []string `json:"forwardedHeaders"`
			GrantTypes                  []string `json:"grantTypes"`
			MacAlgorithm                string   `json:"macAlgorithm"`
			MasheryTokenApiEnabled      bool     `json:"masheryTokenApiEnabled"`
			QpsLimitCeiling             int64    `json:"qpsLimitCeiling"`
			RateLimitCeiling            int64    `json:"rateLimitCeiling"`
			RefreshTokenEnabled         bool     `json:"refreshTokenEnabled"`
			RefreshTokenTtl             int64    `json:"refreshTokenTtl"`
			SecureTokensEnabled         bool     `json:"secureTokensEnabled"`
			TokenBasedRateLimitsEnabled bool     `json:"tokenBasedRateLimitsEnabled"`
		} `json:"oauth"`
	} `json:"securityProfile"`
	Updated string `json:"updated"`
	Version string `json:"version"`
}

type IoDocs struct {
	ServiceId string `json:"serviceId"`
	Definition interface{} `json:"definition"`
	Created    string `json:"created"`
	DefaultApi  interface{} `json:"defaultApi"`
}

type MethodParams struct {
	ServiceId string
}

type CreateServiceRequest struct {
	Id                string `json:"id"`
	RevisionNumber    int64  `json:"revisionNumber"`
	Created           string `json:"created"`
	Updated           string `json:"updated"`
	Name              string `json:"name"`
	RobotsPolicy      string `json:"robotsPolicy"`
	CrossDomainPolicy string `json:"crossdomainPolicy"`
	Description       string `json:"description"`
	QpsLimitOverall   int64  `json:"qpsLimitOverall"`
	RFC3986Encode     bool   `json:"rfc3986Encode"`
	EditorHandle      string `json:"editorHandle"`
	Version           string `json:"version"`
}

type CreateServiceResponse struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Version string `json:"version"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

type CreateOauthRequest struct {
	AccessTokenTTL              int64    `json:"accessTokenTtl"`
	AccessTokenTTLEnabled       bool     `json:"accessTokenTtlEnabled"`
	AccessTokenType             string   `json:"accessTokenType"`
	AllowMultipleToken          bool     `json:"allowMultipleToken"`
	AuthorizationCodeTTL        int64    `json:"authorizationCodeTtl"`
	EnableRefreshTokenTTL       bool     `json:"enableRefreshTokenTtl"`
	ForceOauthRedirectURL       bool     `json:"forceOauthRedirectUrl"`
	ForceSslRedirectURLEnabled  bool     `json:"forceSslRedirectUrlEnabled"`
	ForwardedHeaders            []string `json:"forwardedHeaders"`
	GrantTypes                  []string `json:"grantTypes"`
	MacAlgorithm                string   `json:"macAlgorithm"`
	MasheryTokenAPIEnabled      bool     `json:"masheryTokenApiEnabled"`
	QPSLimitCeiling             int64    `json:"qpsLimitCeiling"`
	RateLimitCeiling            int64    `json:"rateLimitCeiling"`
	RefreshTokenEnabled         bool     `json:"refreshTokenEnabled"`
	RefreshTokenTTL             int64    `json:"refreshTokenTtl"`
	SecureTokensEnabled         bool     `json:"secureTokensEnabled"`
	TokenBasedRateLimitsEnabled bool     `json:"tokenBasedRateLimitsEnabled"`
}

type CreateOauthResponse struct {
	QPSLimitCeiling             int      `json:"qpsLimitCeiling"`
	GrantTypes                  []string `json:"grantTypes,omitempty"`
	EnableRefreshTokenTTL       bool     `json:"enableRefreshTokenTtl"`
	ForceSslRedirectURLEnabled  bool     `json:"forceSslRedirectUrlEnabled"`
	AccessTokenTTL              int      `json:"accessTokenTtl"`
	MasheryTokenAPIEnabled      bool     `json:"masheryTokenApiEnabled"`
	ForceOauthRedirectURL       bool     `json:"forceOauthRedirectUrl"`
	AccessTokenTTLEnabled       bool     `json:"accessTokenTtlEnabled"`
	TokenBasedRateLimitsEnabled bool     `json:"tokenBasedRateLimitsEnabled"`
	MacAlgorithm                string   `json:"macAlgorithm"`
	AllowMultipleToken          bool     `json:"allowMultipleToken"`
	RefreshTokenEnabled         bool     `json:"refreshTokenEnabled"`
	ForwardedHeaders            []string `json:"forwardedHeaders,omitempty"`
	AuthorizationCodeTTL        int      `json:"authorizationCodeTtl"`
	RefreshTokenTTL             int      `json:"refreshTokenTtl"`
	SecureTokensEnabled         bool     `json:"secureTokensEnabled"`
	AccessTokenType             string   `json:"accessTokenType"`
	RateLimitCeiling            int      `json:"rateLimitCeiling"`
}

type CreateEndpointRequest struct {
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

type CreateEndpointResponse struct {
	Id      string `json:"id"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Name    string `json:"name"`
}

type CreateCorsRequest struct {
	AllDomainsEnabled        bool     `json:"allDomainsEnabled"`
	CookiesAllowed           bool     `json:"cookiesAllowed"`
	DomainsAllowed           []string `json:"domainsAllowed"`
	HeadersAllowed           []string `json:"headersAllowed"`
	HeadersExposed           []string `json:"headersExposed"`
	MaxAge                   int64    `json:"maxAge"`
	SubDomainMatchingAllowed bool     `json:"subDomainMatchingAllowed"`
}

type CreateCorsResponse struct {
	MaxAge                   int           `json:"maxAge"`
	CookiesAllowed           bool          `json:"cookiesAllowed"`
	DomainsAllowed           []interface{} `json:"domainsAllowed"`
	HeadersAllowed           []string      `json:"headersAllowed"`
	HeadersExposed           []string      `json:"headersExposed"`
	SubDomainMatchingAllowed bool          `json:"subDomainMatchingAllowed"`
	AllDomainsEnabled        bool          `json:"allDomainsEnabled"`
}

type CreateMethodRequest struct {
	ID                 string `json:"id"`
	SampleJSONResponse string `json:"sampleJsonResponse,omitempty"`
	Updated            string `json:"updated"`
	Created            string `json:"created"`
	Name               string `json:"name"`
	SampleXMLResponse  string `json:"sampleXmlResponse,omitempty"`
}

type CreateMethodResponse struct {
	Id      string `json:"id"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Name    string `json:"name"`
}

type DeleteServiceResponse struct {
	Status string `json:"status"`
}
