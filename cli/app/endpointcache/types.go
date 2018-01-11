package endpointcache

type DeleteEndpointCacheResponse struct {
	Status string `json:"status"`
}

type MethodParams struct {
	ServiceId  string
	EndpointId string
}

type EndpointCache struct {
	Cache struct {
		CacheTTLOverride               int      `json:"cacheTtlOverride"`
		ClientSurrogateControlEnabled  bool     `json:"clientSurrogateControlEnabled"`
		ContentCacheKeyHeaders         []string `json:"contentCacheKeyHeaders"`
		IncludeAPIKeyInContentCacheKey bool     `json:"includeApiKeyInContentCacheKey"`
		RespondFromStaleCacheEnabled   bool     `json:"respondFromStaleCacheEnabled"`
		ResponseCacheControlEnabled    bool     `json:"responseCacheControlEnabled"`
		VaryHeaderEnabled              bool     `json:"varyHeaderEnabled"`
	} `json:"cache"`
}

type UpdateCacheResponse struct {
	Created string `json:"created"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Updated string `json:"updated"`
}
