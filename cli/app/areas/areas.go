package areas

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/applications"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/mmussett/mashcli/cli/app/members"
	"github.com/mmussett/mashcli/cli/app/packagekeys"
	"github.com/mmussett/mashcli/cli/app/roles"
	"github.com/mmussett/mashcli/cli/app/services"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	servicesResourcePath           = "v3/rest/services/%s"
	servicesResourceCollectionPath = "v3/rest/services"
	iodocsResourcePath             = "v3/rest/iodocs/services/%s"
)

func Nuke(accessToken string, preview bool) error {

	err := packagekeys.Nuke(accessToken, preview)
	if err != nil {
		return err
	}

	err = applications.Nuke(accessToken, preview)
	if err != nil {
		return err
	}

	err = members.Nuke(accessToken, preview)
	if err != nil {
		return err
	}

	err = services.Nuke(accessToken, preview)
	if err != nil {
		return err
	}


	err = roles.Nuke(accessToken, preview)
	if err != nil {
		return err
	}

	return nil


}


func deleteService(accessToken, serviceId string) error {

	path := fmt.Sprintf(servicesResourcePath, serviceId)
	r := new(DeleteServiceResponse)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if resp.StatusCode == 200 || resp.StatusCode == 404 {
		return nil
	}

	if err != io.EOF && (resp.StatusCode != 200 || resp.StatusCode != 404) {
		return errors.Errorf("areas: unable to delete service id=%s: DELETE %s", serviceId, path)
	}

	return nil

}
func getServicesCollection(accessToken string, params *mashcli.Params) (*[]ServicesCollection, error) {

	path := servicesResourceCollectionPath
	e := new(mashcli.MasheryError)
	p := new([]ServicesCollection)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(p, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, fmt.Errorf("area: unable to get service collection: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func getService(accessToken string, mp *MethodParams, params *mashcli.Params) (*Services, error) {

	path := fmt.Sprintf(servicesResourcePath, mp.ServiceId)
	e := new(mashcli.MasheryError)
	p := new(Services)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(p, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, fmt.Errorf("areas: unable to get service: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func getIoDocs(accessToken string, mp *MethodParams, params *mashcli.Params) (*IoDocs, error) {

	path := fmt.Sprintf(iodocsResourcePath, mp.ServiceId)
	e := new(mashcli.MasheryError)
	p := new(IoDocs)

	resp, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(p, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, fmt.Errorf("areas: unable to get iodocs: GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return p, nil
}

func BackupArea(accessToken string, backupName string) error {

	servicesPath := fmt.Sprintf("%s/services", backupName)
	os.MkdirAll(servicesPath, 0700)

	iodocsPath := fmt.Sprintf("%s/iodocs", backupName)
	os.MkdirAll(iodocsPath, 0700)

	servicesCollection, err := getServicesCollection(accessToken, &mashcli.Params{Fields: SERVICES_COLLECTION_FIELDS})
	if err != nil {
		return err
	}

	for _, s := range *servicesCollection {

		service, err := getService(accessToken, &MethodParams{ServiceId: s.Id}, &mashcli.Params{Fields: SERVICE_EXPORT_FIELDS})
		if err != nil {
			return err
		}

		serviceAsString, _ := json.MarshalIndent(service, "", "  ")
		err = ioutil.WriteFile(backupName+"/services/"+s.Id+"-"+strings.Replace(s.Name, " ", "_", -1)+".json", serviceAsString, 0644)
		if err != nil {
			return err
		}

		iodoc, err := getIoDocs(accessToken, &MethodParams{ServiceId: s.Id}, &mashcli.Params{Fields: IODOCS_EXPORT_FIELDS})
		iodocAsString, _ := json.MarshalIndent(iodoc, "", "  ")
		err = ioutil.WriteFile(backupName+"/iodocs/"+s.Id+"-"+strings.Replace(s.Name, " ", "_", -1)+".json", iodocAsString, 0644)
		if err != nil {
			return err
		}

	}

	return nil

}

func RestoreArea(accessToken string, backupName string) error {

	servicesPath := fmt.Sprintf("%s/services", backupName)

	files, err := ioutil.ReadDir(servicesPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		filename := fmt.Sprintf("%s/%s", servicesPath, file.Name())
		if strings.HasSuffix(filename,".json") {
			err = importService(accessToken, filename)
			if err != nil {
				return fmt.Errorf("area: failed to import service definition %s : %v", filename, err)
			}
		}
	}

	return nil

}

func importService(accessToken string, filename string) error {

	e := new(mashcli.MasheryError)
	service := new(Services)

	fmt.Println(filename)
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return fmt.Errorf("area: file not found %s", filename)
	}

	err = json.Unmarshal(data, &service)
	if err != nil {
		return err
	}

	createServiceRequest := new(CreateServiceRequest)

	createServiceRequest.Name = service.Name
	createServiceRequest.Version = service.Version
	createServiceRequest.CrossDomainPolicy = service.CrossdomainPolicy
	createServiceRequest.RFC3986Encode = service.Rfc3986Encode
	createServiceRequest.Description = service.Description
	createServiceRequest.RobotsPolicy = service.RobotsPolicy
	createServiceRequest.QpsLimitOverall = service.QpsLimitOverall
	createServiceRequest.Id = service.Id
	createServiceRequest.Created = service.Created
	createServiceRequest.Updated = service.Updated
	createServiceRequest.EditorHandle = service.EditorHandle
	createServiceRequest.RevisionNumber = service.RevisionNumber

	createServiceResponse := new(CreateServiceResponse)

	path := servicesResourceCollectionPath
	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(createServiceRequest).Set("Content-Type", "application/json").Set("Accept", "application/json").Post(path).Receive(createServiceResponse, e)

	if err != nil {
		return err
	}

	if e.ErrorCode != 0 || resp.StatusCode != 200 {
		return fmt.Errorf("area: unable to create service: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	if service.SecurityProfile.Oauth.GrantTypes != nil {

		// Create oath security profile
		createOauthRequest := new(CreateOauthRequest)
		createOauthRequest.ForwardedHeaders = service.SecurityProfile.Oauth.ForwardedHeaders
		createOauthRequest.AccessTokenTTL = service.SecurityProfile.Oauth.AccessTokenTtl
		createOauthRequest.AccessTokenTTLEnabled = service.SecurityProfile.Oauth.AccessTokenTtlEnabled
		createOauthRequest.AccessTokenType = service.SecurityProfile.Oauth.AccessTokenType
		createOauthRequest.AllowMultipleToken = service.SecurityProfile.Oauth.AllowMultipleToken
		createOauthRequest.AuthorizationCodeTTL = service.SecurityProfile.Oauth.AuthorizationCodeTtl
		createOauthRequest.EnableRefreshTokenTTL = service.SecurityProfile.Oauth.EnableRefreshTokenTtl
		createOauthRequest.ForceOauthRedirectURL = service.SecurityProfile.Oauth.ForceOauthRedirectURL
		createOauthRequest.ForceSslRedirectURLEnabled = service.SecurityProfile.Oauth.ForceSslRedirectURLEnabled
		createOauthRequest.ForwardedHeaders = service.SecurityProfile.Oauth.ForwardedHeaders
		createOauthRequest.GrantTypes = service.SecurityProfile.Oauth.GrantTypes
		createOauthRequest.MacAlgorithm = service.SecurityProfile.Oauth.MacAlgorithm
		createOauthRequest.MasheryTokenAPIEnabled = service.SecurityProfile.Oauth.MasheryTokenApiEnabled
		createOauthRequest.QPSLimitCeiling = service.SecurityProfile.Oauth.QpsLimitCeiling
		createOauthRequest.RateLimitCeiling = service.SecurityProfile.Oauth.RateLimitCeiling
		createOauthRequest.RefreshTokenEnabled = service.SecurityProfile.Oauth.RefreshTokenEnabled
		createOauthRequest.RefreshTokenTTL = service.SecurityProfile.Oauth.RefreshTokenTtl
		createOauthRequest.SecureTokensEnabled = service.SecurityProfile.Oauth.SecureTokensEnabled
		createOauthRequest.TokenBasedRateLimitsEnabled = service.SecurityProfile.Oauth.TokenBasedRateLimitsEnabled

		createOauthResponse := new(CreateOauthResponse)

		path := fmt.Sprintf("v3/rest/services/%s/securityProfile/oauth", createServiceResponse.Id)
		resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(createOauthRequest).Set("Content-Type", "application/json").Post(path).Receive(createOauthResponse, e)
		if err != nil {
			return err
		}

		if e.ErrorCode != 0 || resp.StatusCode != 200 {
			return fmt.Errorf("area: unable to create security profile: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
		}

	}

	// Create endpoints

	createEndpointRequest := new(CreateEndpointRequest)
	for _, serviceEndpoint := range service.Endpoints {

		createEndpointRequest.TrafficManagerDomain = serviceEndpoint.TrafficManagerDomain
		createEndpointRequest.RequestProtocol = serviceEndpoint.RequestProtocol
		createEndpointRequest.SupportedHttpMethods = serviceEndpoint.SupportedHttpMethods
		createEndpointRequest.NumberOfHttpRedirectsToFollow = serviceEndpoint.NumberOfHttpRedirectsToFollow
		createEndpointRequest.UseSystemDomainCredentials = serviceEndpoint.UseSystemDomainCredentials
		createEndpointRequest.RequestPathAlias = serviceEndpoint.RequestPathAlias
		createEndpointRequest.RequestAuthenticationType = serviceEndpoint.RequestAuthenticationType
		createEndpointRequest.ApiMethodDetectionLocations = serviceEndpoint.ApiMethodDetectionLocations
		createEndpointRequest.OutboundRequestTargetPath = serviceEndpoint.OutboundRequestTargetPath
		createEndpointRequest.ApiMethodDetectionKey = serviceEndpoint.ApiMethodDetectionKey
		createEndpointRequest.CookiesDuringHttpRedirectsEnabled = serviceEndpoint.CookiesDuringHttpRedirectsEnabled
		createEndpointRequest.ForwardedHeaders = serviceEndpoint.ForwardedHeaders
		createEndpointRequest.Id = ""
		createEndpointRequest.Updated = ""
		createEndpointRequest.Created = ""
		createEndpointRequest.GzipPassthroughSupportEnabled = serviceEndpoint.GzipPassthroughSupportEnabled
		createEndpointRequest.InboundSslRequired = serviceEndpoint.InboundSslRequired
		createEndpointRequest.Name = serviceEndpoint.Name
		createEndpointRequest.ApiKeyValueLocations = serviceEndpoint.ApiKeyValueLocations
		createEndpointRequest.CustomRequestAuthenticationAdapter = serviceEndpoint.CustomRequestAuthenticationAdapter
		createEndpointRequest.OauthGrantTypes = serviceEndpoint.OauthGrantTypes
		createEndpointRequest.AllowMissingApiKey = serviceEndpoint.AllowMissingApiKey
		createEndpointRequest.DropApiKeyFromIncomingCall = serviceEndpoint.DropApiKeyFromIncomingCall
		createEndpointRequest.OutboundRequestTargetQueryParameters = serviceEndpoint.OutboundRequestTargetQueryParameters
		createEndpointRequest.HighSecurity = serviceEndpoint.HighSecurity
		createEndpointRequest.ForceGzipOfBackendCall = serviceEndpoint.ForceGzipOfBackendCall
		createEndpointRequest.ConnectionTimeoutForSystemDomainResponse = serviceEndpoint.ConnectionTimeoutForSystemDomainResponse
		createEndpointRequest.ApiKeyValueLocationKey = serviceEndpoint.ApiKeyValueLocationKey
		createEndpointRequest.ConnectionTimeoutForSystemDomainRequest = serviceEndpoint.ConnectionTimeoutForSystemDomainRequest
		createEndpointRequest.JsonpCallbackParameter = serviceEndpoint.JsonpCallbackParameter
		createEndpointRequest.JsonpCallbackParameterValue = serviceEndpoint.JsonpCallbackParameterValue
		createEndpointRequest.ReturnedHeaders = serviceEndpoint.ReturnedHeaders
		createEndpointRequest.HostPassthroughIncludedInBackendCallHeader = serviceEndpoint.HostPassthroughIncludedInBackendCallHeader
		createEndpointRequest.OutboundTransportProtocol = serviceEndpoint.OutboundTransportProtocol
		createEndpointRequest.SystemDomains = serviceEndpoint.SystemDomains

		// Add the endpoint

		createEndpointResponse := new(CreateEndpointResponse)

		path := fmt.Sprintf("v3/rest/services/%s/endpoints", createServiceResponse.Id)

		resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(createEndpointRequest).Set("Content-Type", "application/json").Set("Accept", "application/json").Post(path).Receive(createEndpointResponse, e)

		if err != nil || resp.StatusCode != 200 {
			deleteService(accessToken, createServiceResponse.Id)
			return errors.Errorf("area: unable to create endpoint: POST %s%s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
		}


		// Add CORS to endpoint

		if !(serviceEndpoint.Cors.DomainsAllowed == nil && serviceEndpoint.Cors.AllDomainsEnabled == false) {

			createCorsRequest := new(CreateCorsRequest)

			createCorsRequest.CookiesAllowed = serviceEndpoint.Cors.CookiesAllowed
			createCorsRequest.DomainsAllowed = serviceEndpoint.Cors.DomainsAllowed
			createCorsRequest.HeadersAllowed = serviceEndpoint.Cors.HeadersAllowed
			createCorsRequest.HeadersExposed = serviceEndpoint.Cors.HeadersExposed
			createCorsRequest.MaxAge = serviceEndpoint.Cors.MaxAge
			createCorsRequest.SubDomainMatchingAllowed = serviceEndpoint.Cors.SubDomainMatchingAllowed
			createCorsRequest.AllDomainsEnabled = serviceEndpoint.Cors.AllDomainsEnabled

			createCorsResponse := new(CreateCorsResponse)

			path := fmt.Sprintf("v3/rest/services/%s/endpoints/%s/cors", createServiceResponse.Id, createEndpointResponse.Id)

			resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(createCorsRequest).Set("Content-Type", "application/json").Set("Accept", "application/json").Post(path).Receive(createCorsResponse, e)
			if err != nil {
				deleteService(accessToken, createServiceResponse.Id)
				return err
			}

			if e.ErrorCode != 0 || resp.StatusCode != 200 {
				deleteService(accessToken, createServiceResponse.Id)
				return fmt.Errorf("area: unable to create endpoint: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
			}

		}

		// Add Methods to endpoint
		createMethodRequest := new(CreateMethodRequest)
		for _, method := range serviceEndpoint.Methods {
			createMethodRequest.ID = method.ID
			createMethodRequest.Name = method.Name
			createMethodRequest.SampleXMLResponse = method.SampleXmlResponse
			createMethodRequest.SampleJSONResponse = method.SampleJsonResponse
			createMethodRequest.Created = method.Created
			createMethodRequest.Updated = method.Updated

			createMethodResponse := new(CreateMethodResponse)

			path := fmt.Sprintf("v3/rest/services/%s/endpoints/%s/methods", createServiceResponse.Id, createEndpointResponse.Id)

			resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).BodyJSON(createMethodRequest).Set("Content-Type", "application/json").Set("Accept", "application/json").Post(path).Receive(createMethodResponse, e)

			if err != nil {
				deleteService(accessToken, createServiceResponse.Id)
				return err
			}

			if e.ErrorCode != 0 || resp.StatusCode != 200 {
				deleteService(accessToken, createServiceResponse.Id)
				return fmt.Errorf("area: unable to create endpoint: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
			}

		}

	} // end-for

	return nil

}
