package applicationpackagekeys

type ApplicationPackageKeys struct {
	Apikey  string `json:"apikey,omitempty"`
	Created string `json:"created,omitempty"`
	Id      string `json:"id,omitempty"`
	Package struct {
		Created                     string `json:"created,omitempty"`
		Description                 string `json:"description,omitempty"`
		Id                          string `json:"id,omitempty"`
		IsUsingSharedSecret         bool   `json:"isUsingSharedSecret,omitempty"`
		KeyAdapter                  string `json:"keyAdapter,omitempty"`
		KeyLength                   int    `json:"keyLength,omitempty"`
		Name                        string `json:"name,omitempty"`
		NearQuotaThreshold          int    `json:"nearQuotaThreshold,omitempty"`
		NotifyAdminEmails           string `json:"notifyAdminEmails,omitempty"`
		NotifyAdminNearQuota        bool   `json:"notifyAdminNearQuota,omitempty"`
		NotifyAdminOverQuota        bool   `json:"notifyAdminOverQuota,omitempty"`
		NotifyAdminOverThrottle     bool   `json:"notifyAdminOverThrottle,omitempty"`
		NotifyAdminPeriod           string `json:"notifyAdminPeriod,omitempty"`
		NotifyDeveloperNearQuota    bool   `json:"notifyDeveloperNearQuota,omitempty"`
		NotifyDeveloperOverQuota    bool   `json:"notifyDeveloperOverQuota,omitempty"`
		NotifyDeveloperOverThrottle bool   `json:"notifyDeveloperOverThrottle,omitempty"`
		NotifyDeveloperPeriod       string `json:"notifyDeveloperPeriod,omitempty"`
		SharedSecretLength          int    `json:"sharedSecretLength,omitempty"`
		Status                      string `json:"status,omitempty"`
		Updated                     string `json:"updated,omitempty"`
	} `json:"package"`
	Plan struct {
		Created                            string `json:"created,omitempty"`
		Description                        string `json:"description,omitempty"`
		Id                                 string `json:"id,omitempty"`
		IsModerated                        bool   `json:"isModerated,omitempty"`
		IsPublic                           bool   `json:"isPublic,omitempty"`
		Name                               string `json:"name,omitempty"`
		Notes                              string `json:"notes,omitempty"`
		PackageKeyMax                      int    `json:"packageKeyMax,omitempty"`
		PackageKeyModerationThreshold      int    `json:"packageKeyModerationThreshold,omitempty"`
		QPSLimitCeiling                    int    `json:"qpsLimitCeiling,omitempty"`
		QPSLimitExempt                     bool   `json:"qpsLimitExempt,omitempty"`
		QPSLimitPackageKeyOverrideAllowed  bool   `json:"qpsLimitPackageKeyOverrideAllowed,omitempty"`
		RateLimitCeiling                   int    `json:"rateLimitCeiling,omitempty"`
		RateLimitExempt                    bool   `json:"rateLimitExempt,omitempty"`
		RateLimitPackageKeyOverrideAllowed bool   `json:"rateLimitPackageKeyOverrideAllowed,omitempty"`
		RateLimitPeriod                    string `json:"rateLimitPeriod,omitempty"`
		ResponseFilterOverrideAllowed      bool   `json:"responseFilterOverrideAllowed,omitempty"`
		Status                             string `json:"status,omitempty"`
		Updated                            string `json:"updated,omitempty"`
	} `json:"plan"`
	QPSLimitCeiling  int    `json:"qpsLimitCeiling,omitempty"`
	QPSLimitExempt   bool   `json:"qpsLimitExempt,omitempty"`
	RateLimitCeiling int    `json:"rateLimitCeiling,omitempty"`
	RateLimitExempt  bool   `json:"rateLimitExempt,omitempty"`
	Secret           string `json:"secret,omitempty"`
	Status           string `json:"status,omitempty"`
	Updated          string `json:"updated,omitempty"`
}

type CreateApplicationPackageKeysRequest struct {
	Apikey  interface{} `json:"apikey"`
	Package struct {
		Created                     string        `json:"created"`
		Description                 string        `json:"description"`
		Eav                         interface{}   `json:"eav"`
		Id                          string        `json:"id"`
		KeyLength                   int           `json:"keyLength"`
		Name                        string        `json:"name"`
		NearQuotaThreshold          int           `json:"nearQuotaThreshold"`
		NotifyAdminEmails           string        `json:"notifyAdminEmails"`
		NotifyAdminNearQuota        bool          `json:"notifyAdminNearQuota"`
		NotifyAdminOverQuota        bool          `json:"notifyAdminOverQuota"`
		NotifyAdminOverThrottle     bool          `json:"notifyAdminOverThrottle"`
		NotifyAdminPeriod           string        `json:"notifyAdminPeriod"`
		NotifyDeveloperNearQuota    bool          `json:"notifyDeveloperNearQuota"`
		NotifyDeveloperOverQuota    bool          `json:"notifyDeveloperOverQuota"`
		NotifyDeveloperOverThrottle bool          `json:"notifyDeveloperOverThrottle"`
		NotifyDeveloperPeriod       string        `json:"notifyDeveloperPeriod"`
		Organization                interface{}   `json:"organization"`
		PackageKeyAdapter           string        `json:"packageKeyAdapter"`
		PackageKeyStringLength      int           `json:"packageKeyStringLength"`
		Plans                       []interface{} `json:"plans"`
		SharedSecretLength          int           `json:"sharedSecretLength"`
		Status                      string        `json:"status"`
		Updated                     string        `json:"updated"`
	} `json:"package"`
	Plan struct {
		AdminKeyProvisioningEnabled                    bool          `json:"adminKeyProvisioningEnabled"`
		AllowKeyOverridesForCallsPerDefinedQuotaPeriod bool          `json:"allowKeyOverridesForCallsPerDefinedQuotaPeriod"`
		AllowKeyOverridesForCallsPerSecond             bool          `json:"allowKeyOverridesForCallsPerSecond"`
		Created                                        string        `json:"created"`
		DefinedPeriodForQuota                          string        `json:"definedPeriodForQuota"`
		Description                                    string        `json:"description"`
		Eav                                            interface{}   `json:"eav"`
		EmailTemplateSetID                             int           `json:"emailTemplateSetId"`
		Id                                             string        `json:"id"`
		KeyReviewThreshold                             int           `json:"keyReviewThreshold"`
		MaxCallsPerDefinedQuotaPeriod                  int           `json:"maxCallsPerDefinedQuotaPeriod"`
		MaxCallsPerSecond                              int           `json:"maxCallsPerSecond"`
		MaxNumKeysAllowed                              int           `json:"maxNumKeysAllowed"`
		MaximumAllowedKeys                             int           `json:"maximumAllowedKeys"`
		Name                                           string        `json:"name"`
		Notes                                          string        `json:"notes"`
		NumKeysBeforeReview                            int           `json:"numKeysBeforeReview"`
		QPSLimitCeiling                                int           `json:"qpsLimitCeiling"`
		QPSLimitExempt                                 bool          `json:"qpsLimitExempt"`
		QPSLimitKeyOverrideAllowed                     bool          `json:"qpsLimitKeyOverrideAllowed"`
		RateLimitCeiling                               int           `json:"rateLimitCeiling"`
		RateLimitExempt                                bool          `json:"rateLimitExempt"`
		RateLimitKeyOverrideAllowed                    bool          `json:"rateLimitKeyOverrideAllowed"`
		RateLimitPeriod                                string        `json:"rateLimitPeriod"`
		ResponseFilterOverrideAllowed                  bool          `json:"responseFilterOverrideAllowed"`
		Roles                                          []interface{} `json:"roles"`
		SelfServiceKeyProvisioningEnabled              bool          `json:"selfServiceKeyProvisioningEnabled"`
		Services                                       []interface{} `json:"services"`
		Status                                         string        `json:"status"`
		UnlimitedCallsPerDefinedQuotaPeriod            bool          `json:"unlimitedCallsPerDefinedQuotaPeriod"`
		UnlimitedCallsPerSecond                        bool          `json:"unlimitedCallsPerSecond"`
		Updated                                        string        `json:"updated"`
	} `json:"plan"`
	Secret interface{} `json:"secret"`
	Status string      `json:"status"`
}

type UpdateStatusRequest struct {
	Status string `json:"status"`
}
type DeleteApplicationPackageKeysResponse struct {
	Status string `json:"status"`
}

type MethodParams struct {
	ApplicationId string
	PackageKeyId  string
}
