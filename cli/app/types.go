package app

type Methods []struct {
	ID                 string `json:"id"`
	Created            string `json:"created"`
	Name               string `json:"name"`
	SampleJsonResponse string `json:"sampleJsonResponse"`
	SampleXmlResponse  string `json:"sampleXmlResponse"`
	Updated            string `json:"updated"`
}

type Application struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
	ID       string `json:"id"`
}

type PackageKey struct {
	Status  string `json:"status"`
	Apikey  string `json:"apikey"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	ID      string `json:"id"`
	Package struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"package"`
	Plan struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"plan"`
}

type Service struct {
	ID    string `json:"id"`
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

type Endpoint struct {
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
	HeadersToExcludeFromIncomingCall           string   `json:"headersToExcludeFromIncomingCall"`
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

	publicDomains PublicDomains `json:"publicDomains"`
	//PublicDomains []struct {
	//	Address string `json:"address"`
	//} `json:"publicDomains"`
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
	//SystemDomains []struct {
	//	Address string `json:"address"`
	//} `json:"systemDomains"`

	systemDomains SystemDomains `json:"systemDomains"`

	TrafficManagerDomain       string `json:"trafficManagerDomain"`
	UseSystemDomainCredentials bool   `json:"useSystemDomainCredentials"`
}

type SystemDomains []struct {
	address string `json:"address"`
}

type PublicDomains []struct {
	address string `json:"address"`
}

type Endpoints []struct {
	ID                                       string   `json:"id"`
	AllowMissingApiKey                       bool     `json:"allowMissingApiKey"`
	ApiKeyValueLocationKey                   string   `json:"apiKeyValueLocationKey"`
	ApiKeyValueLocations                     []string `json:"apiKeyValueLocations"`
	ApiMethodDetectionKey                    string   `json:"apiMethodDetectionKey"`
	ApiMethodDetectionLocations              []string `json:"apiMethodDetectionLocations"`
	ConnectionTimeoutForSystemDomainRequest  int64    `json:"connectionTimeoutForSystemDomainRequest"`
	ConnectionTimeoutForSystemDomainResponse int64    `json:"connectionTimeoutForSystemDomainResponse"`
	CookiesDuringHttpRedirectsEnabled        bool     `json:"cookiesDuringHttpRedirectsEnabled"`
	Cors                                     struct {
		AllDomainsEnabled        bool     `json:"allDomainsEnabled"`
		CookiesAllowed           bool     `json:"cookiesAllowed"`
		DomainsAllowed           []string `json:"domainsAllowed"`
		HeadersAllowed           []string `json:"headersAllowed"`
		HeadersExposed           []string `json:"headersExposed"`
		MaxAge                   int64    `json:"maxAge"`
		SubDomainMatchingAllowed bool     `json:"subDomainMatchingAllowed"`
	} `json:"cors"`
	Created                                    string   `json:"created"`
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
		ID      string `json:"id"`
		Created string `json:"created"`
		Name    string `json:"name"`
		Updated string `json:"updated"`
	} `json:"methods"`
	Name                                 string `json:"name"`
	NumberOfHttpRedirectsToFollow        int64  `json:"numberOfHttpRedirectsToFollow"`
	OutboundRequestTargetPath            string `json:"outboundRequestTargetPath"`
	OutboundRequestTargetQueryParameters string `json:"outboundRequestTargetQueryParameters"`
	OutboundTransportProtocol            string `json:"outboundTransportProtocol"`
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
		PublicAddress string `json:"address"`
	} `json:"publicDomains"`
	SystemDomains []struct {
		SystemAddress string `json:"address"`
	} `json:"systemDomains"`
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

	TrafficManagerDomain       string `json:"trafficManagerDomain"`
	Updated                    string `json:"updated"`
	UseSystemDomainCredentials bool   `json:"useSystemDomainCredentials"`
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

type DeleteServiceResponse struct {
	Status string `json:"status"`
}

type CreateEndpointRequest struct {
	TrafficManagerDomain               string   `json:"trafficManagerDomain"`
	RequestProtocol                    string   `json:"requestProtocol"`
	SupportedHTTPMethods               []string `json:"supportedHttpMethods"`
	NumberOfHTTPRedirectsToFollow      int64    `json:"numberOfHttpRedirectsToFollow"`
	UseSystemDomainCredentials         bool     `json:"useSystemDomainCredentials"`
	RequestPathAlias                   string   `json:"requestPathAlias"`
	RequestAuthenticationType          string   `json:"requestAuthenticationType"`
	APIMethodDetectionLocations        []string `json:"apiMethodDetectionLocations"`
	OutboundRequestTargetPath          string   `json:"outboundRequestTargetPath"`
	APIMethodDetectionKey              string   `json:"apiMethodDetectionKey"`
	ID                                 string   `json:"id"`
	StringsToTrimFromAPIKey            string   `json:"stringsToTrimFromApiKey"`
	CookiesDuringHTTPRedirectsEnabled  bool     `json:"cookiesDuringHttpRedirectsEnabled"`
	ForwardedHeaders                   []string `json:"forwardedHeaders"`
	Updated                            string   `json:"updated"`
	Created                            string   `json:"created"`
	GzipPassthroughSupportEnabled      bool     `json:"gzipPassthroughSupportEnabled"`
	InboundSslRequired                 bool     `json:"inboundSslRequired"`
	Name                               string   `json:"name"`
	APIKeyValueLocations               []string `json:"apiKeyValueLocations"`
	CustomRequestAuthenticationAdapter string   `json:"customRequestAuthenticationAdapter,omitempty"`
	SystemDomainCredentialSecret       string   `json:"systemDomainCredentialSecret"`
	SystemDomains                      []struct {
		Address string `json:"address"`
	} `json:"systemDomains"`
	OauthGrantTypes                            []string `json:"oauthGrantTypes"`
	AllowMissingAPIKey                         bool     `json:"allowMissingApiKey"`
	DropAPIKeyFromIncomingCall                 bool     `json:"dropApiKeyFromIncomingCall"`
	SystemDomainCredentialKey                  string   `json:"systemDomainCredentialKey"`
	OutboundRequestTargetQueryParameters       string   `json:"outboundRequestTargetQueryParameters"`
	HighSecurity                               bool     `json:"highSecurity"`
	ForceGzipOfBackendCall                     bool     `json:"forceGzipOfBackendCall"`
	ConnectionTimeoutForSystemDomainResponse   int64    `json:"connectionTimeoutForSystemDomainResponse"`
	APIKeyValueLocationKey                     string   `json:"apiKeyValueLocationKey"`
	ConnectionTimeoutForSystemDomainRequest    int64    `json:"connectionTimeoutForSystemDomainRequest"`
	JsonpCallbackParameter                     string   `json:"jsonpCallbackParameter"`
	JsonpCallbackParameterValue                string   `json:"jsonpCallbackParameterValue"`
	ReturnedHeaders                            []string `json:"returnedHeaders"`
	HostPassthroughIncludedInBackendCallHeader bool     `json:"hostPassthroughIncludedInBackendCallHeader"`
	OutboundTransportProtocol                  string   `json:"outboundTransportProtocol"`
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

type CreateMethodRequest struct {
	ID                 string `json:"id"`
	SampleJSONResponse string `json:"sampleJsonResponse,omitempty"`
	Updated            string `json:"updated"`
	Created            string `json:"created"`
	Name               string `json:"name"`
	SampleXMLResponse  string `json:"sampleXmlResponse,omitempty"`
}

type CreateEndpointResponse struct {
	Id      string `json:"id"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Name    string `json:"name"`
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

type CreateMethodResponse struct {
	Id      string `json:"id"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Name    string `json:"name"`
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

type GetServiceResponse struct {
	Name            string `json:"name"`
	Id              string `json:"id"`
	Version         string `json:"version"`
	Created         string `json:"created"`
	Updated         string `json:"updated"`
	Description     string `json:"description"`
	QpsLimitOverall int64  `json:"qpsLimitOverall"`
}

type Params struct {
	Fields string `url:"fields,omitempty"`
}

type ErrorResponse struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type GetPackageResponse struct {
	Name         string `json:"name"`
	Id           string `json:"id"`
	Organization string `json:"organization"`
	Created      string `json:"created"`
	Updated      string `json:"updated"`
}

/*
const PACKAGE_ALL_FIELDS = `notifyDeveloperPeriod,notifyDeveloperNearQuota,
notifyDeveloperOverQuota,notifyDeveloperOverThrottle,notifyAdminPeriod,notifyAdminNearQuota,notifyAdminOverQuota,
notifyAdminOverThrottle,notifyAdminEmails,nearQuotaThreshold,eav,keyAdapter,keyLength,sharedSecretLength,
plans.id,
plans.created,
plans.updated,plans.name,
plans.description,
plans.selfServiceKeyProvisioningEnabled,
plans.adminKeyProvisioningEnabled,
plans.notes,
plans.maxNumKeysAllowed,
plans.numKeysBeforeReview,
plans.qpsLimitCeiling,
plans.qpsLimitExempt,
plans.qpsLimitKeyOverrideAllowed,
plans.rateLimitCeiling,
plans.rateLimitExempt,
plans.rateLimitKeyOverrideAllowed,
plans.rateLimitPeriod,
plans.responseFilterOverrideAllowed,
plans.status,
plans.emailTemplateSetId,
plans.services`
*/

type Plan struct {
	Name                              string    `json:"name"`
	Id                                string    `json:"id"`
	Status                            string    `json:"status"`
	Description                       string    `json:"description"`
	Created                           string    `json:"created"`
	Updated                           string    `json:"updated"`
	MaxNumKeysAllowed                 int       `json:"maxNumKeysAllowed"`
	NumKeysBeforeReview               int       `json:"numKeysBeforeReview"`
	QPSLimitCeiling                   int       `json:"qpsLimitCeiling"`
	RateLimitCeiling                  int       `json:"rateLimitCeiling"`
	RateLimitPeriod                   string    `json:"rateLimitPeriod"`
	AdminKeyProvisioningEnabled       bool      `json:"adminKeyProvisioningEnabled"`
	QPSLimitExempt                    bool      `json:"qpsLimitExempt"`
	QPSLimitKeyOverrideAllowed        bool      `json:"qpsLimitKeyOverrideAllowed"`
	RateLimitExempt                   bool      `json:"rateLimitExempt"`
	RateLimitKeyOverrideAllowed       bool      `json:"rateLimitKeyOverrideAllowed"`
	ResponseFilterOverrideAllowed     bool      `json:"responseFilterOverrideAllowed"`
	SelfServiceKeyProvisioningEnabled bool      `json:"selfServiceKeyProvisioningEnabled"`
	Services                          []Service `json:"services"`
}

type Package struct {
	Name                        string      `json:"name"`
	Id                          string      `json:"id"`
	Organization                interface{} `json:"organization"`
	Created                     string      `json:"created"`
	Updated                     string      `json:"updated"`
	Description                 string      `json:"description"`
	NotifyDeveloperPeriod       string      `json:"notifyDeveloperPeriod"`
	NotifyDeveloperNearQuota    bool        `json:"notifyDeveloperNearQuota"`
	NotifyDeveloperOverQuota    bool        `json:"notifyDeveloperOverQuota"`
	NotifyDeveloperOverThrottle bool        `json:"notifyDeveloperOverThrottle"`
	NotifyAdminPeriod           string      `json:"notifyAdminPeriod"`
	NotifyAdminNearQuota        bool        `json:"notifyAdminNearQuota"`
	NotifyAdminOverQuota        bool        `json:"notifyAdminOverQuota"`
	NotifyAdminOverThrottle     bool        `json:"notifyAdminOverThrottle"`
	NotifyAdminEmails           string      `json:"notifyAdminEmails"`
	NearQuotaThreshold          int         `json:"nearQuotaThreshold""`
	KeyAdapter                  string      `json:"keyAdapter"`
	KeyLength                   int         `json:"keyLength"`
	SharedSecretLength          int         `json:"sharedSecretLength"`
	Plans                       []Plan      `json:"plans"`
}

type SecurityProfile struct {
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
}
