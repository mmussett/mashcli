package servicecache

type ServiceCache struct {
	CacheTtl              int     `json:"cacheTtl,omitempty"`
}


type MethodParams struct {
	ServiceId  string
}


type DeleteServiceCacheResponse struct {
	Created string `json:"created"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Updated string `json:"updated"`
}
