package endpoints

const ENDPOINTS_ALL_FIELDS = `id, allowMissingApiKey, apiKeyValueLocationKey, apiKeyValueLocations, apiMethodDetectionKey, apiMethodDetectionLocations, cache
, connectionTimeoutForSystemDomainRequest, connectionTimeoutForSystemDomainResponse, cookiesDuringHttpRedirectsEnabled, cors, created, customRequestAuthenticationAdapter
, dropApiKeyFromIncomingCall, forceGzipOfBackendCall, gzipPassthroughSupportEnabled, headersToExcludeFromIncomingCall, highSecurity, hostPassthroughIncludedInBackendCallHeader
, inboundSslRequired, jsonpCallbackParameter, jsonpCallbackParameterValue, scheduledMaintenanceEvent, forwardedHeaders, returnedHeaders, methods, name, numberOfHttpRedirectsToFollow
, outboundRequestTargetPath, outboundRequestTargetQueryParameters, outboundTransportProtocol, processor, publicDomains, requestAuthenticationType, requestPathAlias, requestProtocol
, oauthGrantTypes, stringsToTrimFromApiKey, supportedHttpMethods, systemDomainAuthentication, systemDomains, trafficManagerDomain, updated, useSystemDomainCredentials
, systemDomainCredentialKey, systemDomainCredentialSecret`

/* All the fields for a Service
id, name, created, updated,editorHandle, revisionNumber, robotsPolicy, crossdomainPolicy,description, errorSets, errorSets.name, errorSets.type, errorSets.jsonp
, errorSets.jsonpType, errorSets.errorMessages, qpsLimitOverall
, rfc3986Encode, securityProfile, version, cache, roles, roles.id
, roles.created, roles.updates, roles.name, roles.action
, endpoints.allowMissingApiKey, endpoints.apiKeyValueLocationKey
, endpoints.apiKeyValueLocations, endpoints.apiMethodDetectionKey
, endpoints.apiMethodDetectionLocations
, endpoints.cache.clientSurrogateControlEnabled
, endpoints.cache.contentCacheKeyHeaders
, endpoints.connectionTimeoutForSystemDomainRequest
, endpoints.connectionTimeoutForSystemDomainResponse
, endpoints.cookiesDuringHttpRedirectsEnabled, endpoints.cors
, endpoints.cors.allDomainsEnabled, endpoints.cors.maxAge
, endpoints.customRequestAuthenticationAdapter
, endpoints.dropApiKeyFromIncomingCall, endpoints.forceGzipOfBackendCallid
, name, created, updated, editorHandle, revisionNumber, robotsPolicy
, crossdomainPolicy, description, errorSets, errorSets.name, errorSets.type
, errorSets.jsonp, errorSets.jsonpType, errorSets.errorMessages
, qpsLimitOverall, rfc3986Encode, securityProfile, version, cache, roles
, roles.id, roles.created, roles.updates, roles.name, roles.action
, endpoints.allowMissingApiKey, endpoints.apiKeyValueLocationKey
, endpoints.apiKeyValueLocations, endpoints.apiMethodDetectionKey
, endpoints.apiMethodDetectionLocations
, endpoints.cache.clientSurrogateControlEnabled
, endpoints.cache.contentCacheKeyHeaders
, endpoints.connectionTimeoutForSystemDomainRequest
, endpoints.connectionTimeoutForSystemDomainResponse
, endpoints.cookiesDuringHttpRedirectsEnabled, endpoints.cors
, endpoints.cors.allDomainsEnabled, endpoints.cors.maxAge
, endpoints.customRequestAuthenticationAdapter
, endpoints.dropApiKeyFromIncomingCall, endpoints.forceGzipOfBackendCall
, endpoints.gzipPassthroughSupportEnabled
, endpoints.headersToExcludeFromIncomingCall, endpoints.highSecurity
, endpoints.hostPassthroughIncludedInBackendCallHeader
, endpoints.inboundSslRequired, endpoints.jsonpCallbackParameter
, endpoints.jsonpCallbackParameterValue, endpoints.scheduledMaintenanceEvent
, endpoints.forwardedHeaders, endpoints.returnedHeaders, endpoints.methods
, endpoints.methods.name, endpoints.methods.sampleJsonResponse
, endpoints.methods.sampleXmlResponse, endpoints.methods.responseFilters
, endpoints.methods.responseFilters.id, endpoints.methods.responseFilters.name
, endpoints.methods.responseFilters.created
, endpoints.methods.responseFilters.updated
, endpoints.methods.responseFilters.notes
, endpoints.methods.responseFilters.xmlFilterFields
, endpoints.methods.responseFilters.jsonFilterFields, endpoints.name
, endpoints.numberOfHttpRedirectsToFollow, endpoints.outboundRequestTargetPath
, endpoints.outboundRequestTargetQueryParameters
, endpoints.outboundTransportProtocol, endpoints.processor
, endpoints.publicDomains, endpoints.requestAuthenticationType
, endpoints.scheduledMaintenanceEvent, endpoints.scheduledMaintenanceEvent.id
, endpoints.scheduledMaintenanceEvent.name
, endpoints.scheduledMaintenanceEvent.startDateTime
, endpoints.scheduledMaintenanceEvent.endDateTime
, endpoints.scheduledMaintenanceEvent.endpoints, endpoints.requestPathAlias
, endpoints.requestProtocol, endpoints.oauthGrantTypes
, endpoints.stringsToTrimFromApiKey, endpoints.supportedHttpMethods
, endpoints.systemDomainAuthentication
, endpoints.systemDomainAuthentication.type
, endpoints.systemDomainAuthentication.username
, endpoints.systemDomainAuthentication.certificate
, endpoints.systemDomainAuthentication.password, endpoints.systemDomains
, endpoints.trafficManagerDomain, endpoints.useSystemDomainCredentials
*/

const PACKAGE_BASE_FIELDS = "id,name,created,updated,organization"

const PACKAGE_ALL_FIELDS = `id,name,created,updated,organization,description,notifyDeveloperPeriod,notifyDeveloperNearQuota,
notifyDeveloperOverQuota,notifyDeveloperOverThrottle,notifyAdminPeriod,notifyAdminNearQuota,notifyAdminOverQuota,
notifyAdminOverThrottle,notifyAdminEmails,nearQuotaThreshold,eav,keyAdapter,keyLength,sharedSecretLength,
plans.id,plans.created,plans.updated,plans.name,plans.description,plans.selfServiceKeyProvisioningEnabled,
plans.adminKeyProvisioningEnabled,plans.notes,plans.maxNumKeysAllowed,plans.numKeysBeforeReview,plans.qpsLimitCeiling,
plans.qpsLimitExempt,plans.qpsLimitKeyOverrideAllowed,plans.rateLimitCeiling, plans.rateLimitExempt, plans.rateLimitKeyOverrideAllowed,
plans.rateLimitPeriod,plans.responseFilterOverrideAllowed, plans.status, plans.emailTemplateSetId`

const PLAN_BASE_FIELDS = "id,name,status,selfServiceKeyProvisioningEnabled,adminKeyProvisioningEnabled,created,updated"

const PLAN_ALL_FIELDS = `id,name,created,updated,description,eav,selfServiceKeyProvisioningEnabled,adminKeyProvisioningEnabled,
notes,maxNumKeysAllowed,numKeysBeforeReview,qpsLimitCeiling,qpsLimitExempt,qpsLimitKeyOverrideAllowed,rateLimitCeiling,
rateLimitExempt,rateLimitKeyOverrideAllowed,rateLimitPeriod,responseFilterOverrideAllowed,status,emailTemplateSetId,services`

const SERVICE_EXPORT_FIELDS = `name,robotsPolicy,crossdomainPolicy,description,errorSets,
qpsLimitOverall,rfc3986Encode,securityProfile,version,endpoints.inboundSslRequired,
endpoints.jsonpCallbackParameter,endpoints.jsonpCallbackParameterValue,
endpoints.scheduledMaintenanceEvent,endpoints.forwardedHeaders,endpoints.returnedHeaders,
endpoints.methods.name,endpoints.methods.sampleXmlResponse,endpoints.methods.sampleJsonResponse,
endpoints.methods.responseFilters,endpoints.name,endpoints.numberOfHttpRedirectsToFollow,
endpoints.outboundRequestTargetPath,endpoints.outboundRequestTargetQueryParameters,
endpoints.outboundTransportProtocol,endpoints.processor,endpoints.publicDomains,
endpoints.requestAuthenticationType,endpoints.requestPathAlias,endpoints.requestProtocol,
endpoints.oauthGrantTypes,endpoints.stringsToTrimFromApiKey,endpoints.supportedHttpMethods,
endpoints.systemDomainAuthentication,endpoints.systemDomains,endpoints.trafficManagerDomain,
endpoints.updated,endpoints.useSystemDomainCredentials,endpoints.systemDomainCredentialKey,
endpoints.systemDomainCredentialSecret`

const SERVICE_BASE_FIELDS = "name, id, version, created, updated, description, qpsLimitOverall"

const SERVICE_ALL_FIELDS = `id, name, created, updated,editorHandle, revisionNumber, robotsPolicy, crossdomainPolicy,description, errorSets, errorSets.name, errorSets.type, errorSets.jsonp
, errorSets.jsonpType, errorSets.errorMessages, qpsLimitOverall
, rfc3986Encode, securityProfile, version, cache, roles, roles.id
, roles.created, roles.updates, roles.name, roles.action
, endpoints.allowMissingApiKey, endpoints.apiKeyValueLocationKey
, endpoints.apiKeyValueLocations, endpoints.apiMethodDetectionKey
, endpoints.apiMethodDetectionLocations
, endpoints.cache.clientSurrogateControlEnabled
, endpoints.cache.contentCacheKeyHeaders
, endpoints.connectionTimeoutForSystemDomainRequest
, endpoints.connectionTimeoutForSystemDomainResponse
, endpoints.cookiesDuringHttpRedirectsEnabled, endpoints.cors
, endpoints.cors.allDomainsEnabled, endpoints.cors.maxAge
, endpoints.customRequestAuthenticationAdapter
, endpoints.dropApiKeyFromIncomingCall, endpoints.forceGzipOfBackendCallid
, name, created, updated, editorHandle, revisionNumber, robotsPolicy
, crossdomainPolicy, description, errorSets, errorSets.name, errorSets.type
, errorSets.jsonp, errorSets.jsonpType, errorSets.errorMessages
, qpsLimitOverall, rfc3986Encode, securityProfile, version, cache, roles
, roles.id, roles.created, roles.updates, roles.name, roles.action
, endpoints.allowMissingApiKey, endpoints.apiKeyValueLocationKey
, endpoints.apiKeyValueLocations, endpoints.apiMethodDetectionKey
, endpoints.apiMethodDetectionLocations
, endpoints.cache.clientSurrogateControlEnabled
, endpoints.cache.contentCacheKeyHeaders
, endpoints.connectionTimeoutForSystemDomainRequest
, endpoints.connectionTimeoutForSystemDomainResponse
, endpoints.cookiesDuringHttpRedirectsEnabled, endpoints.cors
, endpoints.cors.allDomainsEnabled, endpoints.cors.maxAge
, endpoints.customRequestAuthenticationAdapter
, endpoints.dropApiKeyFromIncomingCall, endpoints.forceGzipOfBackendCall
, endpoints.gzipPassthroughSupportEnabled
, endpoints.headersToExcludeFromIncomingCall, endpoints.highSecurity
, endpoints.hostPassthroughIncludedInBackendCallHeader
, endpoints.inboundSslRequired, endpoints.jsonpCallbackParameter
, endpoints.jsonpCallbackParameterValue, endpoints.scheduledMaintenanceEvent
, endpoints.forwardedHeaders, endpoints.returnedHeaders, endpoints.methods
, endpoints.methods.name, endpoints.methods.sampleJsonResponse
, endpoints.methods.sampleXmlResponse, endpoints.methods.responseFilters
, endpoints.methods.responseFilters.id, endpoints.methods.responseFilters.name
, endpoints.methods.responseFilters.created
, endpoints.methods.responseFilters.updated
, endpoints.methods.responseFilters.notes
, endpoints.methods.responseFilters.xmlFilterFields
, endpoints.methods.responseFilters.jsonFilterFields, endpoints.name
, endpoints.numberOfHttpRedirectsToFollow, endpoints.outboundRequestTargetPath
, endpoints.outboundRequestTargetQueryParameters
, endpoints.outboundTransportProtocol, endpoints.processor
, endpoints.publicDomains, endpoints.requestAuthenticationType
, endpoints.scheduledMaintenanceEvent, endpoints.scheduledMaintenanceEvent.id
, endpoints.scheduledMaintenanceEvent.name
, endpoints.scheduledMaintenanceEvent.startDateTime
, endpoints.scheduledMaintenanceEvent.endDateTime
, endpoints.scheduledMaintenanceEvent.endpoints, endpoints.requestPathAlias
, endpoints.requestProtocol, endpoints.oauthGrantTypes
, endpoints.stringsToTrimFromApiKey, endpoints.supportedHttpMethods
, endpoints.systemDomainAuthentication
, endpoints.systemDomainAuthentication.type
, endpoints.systemDomainAuthentication.username
, endpoints.systemDomainAuthentication.certificate
, endpoints.systemDomainAuthentication.password, endpoints.systemDomains
, endpoints.trafficManagerDomain, endpoints.useSystemDomainCredentials`

const APPLICATION_ALL_FIELDS = `id,name,username, description, type, commercial, ads, adsSystem, usageModel, tags, notes
, howDidyouHear, preferredProtocol, preferredOutput, externalId, uri,oauthRedirectUri, packageKeys created,updated`
