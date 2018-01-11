package packages

type Package struct {
	Name                        string      `json:"name,omitempty"`
	Id                          string      `json:"id,omitempty"`
	Organization                interface{} `json:"organization,omitempty"`
	Created                     string      `json:"created,omitempty"`
	Updated                     string      `json:"updated,omitempty"`
	Description                 string      `json:"description,omitempty"`
	NotifyDeveloperPeriod       string      `json:"notifyDeveloperPeriod,omitempty"`
	NotifyDeveloperNearQuota    bool        `json:"notifyDeveloperNearQuota,omitempty"`
	NotifyDeveloperOverQuota    bool        `json:"notifyDeveloperOverQuota,omitempty"`
	NotifyDeveloperOverThrottle bool        `json:"notifyDeveloperOverThrottle,omitempty"`
	NotifyAdminPeriod           string      `json:"notifyAdminPeriod,omitempty"`
	NotifyAdminNearQuota        bool        `json:"notifyAdminNearQuota,omitempty"`
	NotifyAdminOverQuota        bool        `json:"notifyAdminOverQuota,omitempty"`
	NotifyAdminOverThrottle     bool        `json:"notifyAdminOverThrottle,omitempty"`
	NotifyAdminEmails           string      `json:"notifyAdminEmails,omitempty"`
	NearQuotaThreshold          int         `json:"nearQuotaThreshold,omitempty"`
	Eav                         interface{} `json:"eav,omitempty"`
	KeyAdapter                  string      `json:"keyAdapter,omitempty"`
	KeyLength                   int         `json:"keyLength,omitempty"`
	SharedSecretLength          int         `json:"sharedSecretLength,omitempty"`
	Plans                       interface{} `json:"plans,omitempty"`
}

type CreatePackageResponse struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Version string `json:"version"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

type DeletePackageResponse struct {
	Status string `json:"status"`
}

type MethodParams struct {
	PackageId string
}
