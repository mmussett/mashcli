package transform

import (
	"encoding/json"
	"fmt"

	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"

	"io/ioutil"
	"strings"
)

const (
	transformResourcePath = "v3/rest/transform"
	servicesResourcePath  = "v3/rest/services"
)

func TransformSwaggerToIodocs(accessToken, swagger, publicDomain  string) (*Transform, error) {

	e := new(mashcli.MasheryError)
	t := new(Transform)

	params := &Params{PublicDomain:publicDomain,SourceFormat:"swagger2",TargetFormat:"iodocsv1"}

	body := strings.NewReader(swagger)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").Body(body).QueryStruct(params).Post(transformResourcePath).Receive(t, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("transform: unable to transform swagger: POST %s -> (%s)", transformResourcePath, resp.Status)
	}

	return t,nil

}


func TransformSwaggerToService(accessToken, swagger, publicDomain  string) (*Services, error) {

	e := new(mashcli.MasheryError)
	t := new(Transform)
	s := new(Services)

	params := &Params{PublicDomain:publicDomain,SourceFormat:"swagger2",TargetFormat:"masheryapi"}

	body := strings.NewReader(swagger)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").Body(body).QueryStruct(params).Post(transformResourcePath).Receive(t, e)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.Errorf("transform: unable to transform swagger: POST %s -> (%s)", transformResourcePath, resp.Status)
	}

	s.Id = ""
	s.Created = ""
	s.Updated = ""
	s.Name = t.Document.Name
	s.Version = t.Document.Version

	var endpointSlice []Endpoints

	for _, te := range t.Document.Endpoints {
		var publicDomainSlice []PublicDomains
		var systemDomainSlice []SystemDomains
		var endpoint Endpoints
		endpoint.Name = te.Name
		endpoint.OutboundRequestTargetPath = te.OutboundRequestTargetPath
		endpoint.OutboundTransportProtocol = te.OutboundTransportProtocol
		endpoint.RequestAuthenticationType = te.RequestAuthenticationType
		endpoint.RequestPathAlias = te.RequestPathAlias
		endpoint.RequestProtocol = te.RequestProtocol
		endpoint.SupportedHttpMethods = te.SupportedHTTPMethods
		endpoint.TrafficManagerDomain = te.TrafficManagerDomain
		endpoint.InboundSslRequired = te.InboundSslRequired

		publicDomainSlice = append(publicDomainSlice, PublicDomains{Address: te.PublicDomains[0].Address})
		systemDomainSlice = append(systemDomainSlice, SystemDomains{Address: te.SystemDomains[0].Address})
		endpoint.PublicDomains = publicDomainSlice
		endpoint.SystemDomains = systemDomainSlice

		endpointSlice = append(endpointSlice, endpoint)

	}

	s.Endpoints = endpointSlice

	return s,nil

}

func ReadSwaggerFile(filename string) (string,error) {

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.Errorf("transform: unable to read swagger file:  %s", filename)
	}

	return string(bytes), nil
}

func ImportSwagger(accessToken, filename, publicDomain string) error {

	e := new(mashcli.MasheryError)
	serviceResponse := new(CreateServiceResponse)

	swagger, err := ReadSwaggerFile(filename)
	if err != nil {
		return err
	}

	service, err := TransformSwaggerToService(accessToken,swagger,publicDomain)
	if err != nil {
		return err
	}

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").BodyJSON(service).Post(servicesResourcePath).Receive(serviceResponse, e)
	if err != nil {
  	return err
  }

	if resp.StatusCode != 200 {
		return errors.Errorf("transform: unable to create service from swagger: POST %s -> (%s)", servicesResourcePath, resp.Status)
	}

	b, err := json.MarshalIndent(serviceResponse,"","    ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))

	return nil
}

func (a *Services) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (a *Transform) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
