package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"


	"io/ioutil"
	"net/url"
	"os"
)

func AddEndpoint(accessToken, serviceId, endpointName, publicEndpoint, systemEndpoint string) error {

	e := new(Endpoints)
	publicUrl, err := url.Parse(publicEndpoint)
	if err != nil {
		return err
	}

	systemUrl, err := url.Parse(systemEndpoint)
	if err != nil {
		return err
	}


	e.Name = endpointName

	e.RequestPathAlias = publicUrl.Path
	if publicUrl.Scheme == "http" {
		e.InboundSslRequired = false
	} else {
		e.InboundSslRequired = true
	}

	e.InboundSslRequired = false

	e.OutboundRequestTargetPath = systemUrl.Path
	e.OutboundTransportProtocol = systemUrl.Scheme

	e.InboundSslRequired = false
	e.Name = "Test"
	e.OutboundRequestTargetPath = systemUrl.Path
	e.OutboundRequestTargetQueryParameters = ""
	e.OutboundTransportProtocol = "use-inbound"
	e.PublicDomains = make(PublicDomains,1)
	e.PublicDomains[0].Address = publicUrl.Host
	e.RequestPathAlias = publicUrl.Path
	e.SystemDomains = make(SystemDomains,1)
	e.SystemDomains[0].Address = publicUrl.Host
	e.TrafficManagerDomain = publicUrl.Host
	e.HttpsClientProfile = nil
	e.Processor = nil
	e.ApiKeyValueLocationKey = "api_key"
	e.ApiKeyValueLocations = []string{"request-parameters","request-body"}

	e, err = e.Create(accessToken,&MethodParams{ServiceId:serviceId})
	if err != nil {
		return err
	}

	endpointAsString, err := e.Marshall()
	if err != nil {
		return err
	}

	fmt.Println(endpointAsString)

	return nil

}

func CloneEndpoint(accessToken, serviceId, endpointId, endpointName, publicEndpoint, systemEndpoint string) error {

	e, err := Get(accessToken, &MethodParams{ServiceId:serviceId,EndpointId:endpointId},&mashcli.Params{Fields:ENDPOINTS_ALL_FIELDS})
	if err != nil {
		return err
	}

	publicUrl, err := url.Parse(publicEndpoint)
	if err != nil {
		return err
	}

	systemUrl, err := url.Parse(systemEndpoint)
	if err != nil {
		return err
	}


	e.Name = endpointName
	e.RequestPathAlias = publicUrl.Path
	if publicUrl.Scheme == "http" {
		e.InboundSslRequired = false
	} else {
		e.InboundSslRequired = true
	}
	e.PublicDomains[0].Address = publicUrl.Host
	e.OutboundRequestTargetPath = systemUrl.Path
	e.OutboundTransportProtocol = systemUrl.Scheme
	e.SystemDomains[0].Address = systemUrl.Host


	e, err = e.Create(accessToken,&MethodParams{ServiceId:serviceId})
	if err != nil {
		return err
	}

	endpointAsString, err := e.Marshall()
	if err != nil {
		return err
	}

	fmt.Println(endpointAsString)

	return nil
}

func DeleteEndpoint(accessToken, serviceId, endpointId string) error {

	err := Delete(accessToken,&MethodParams{ServiceId:serviceId,EndpointId:endpointId})
	if err != nil {
		return err
	}

	return nil
}


func ShowEndpoints(accessToken, serviceId, endpointId, format string) error {

	m, err := Get(accessToken, &MethodParams{ServiceId:serviceId,EndpointId:endpointId}, &mashcli.Params{Fields: ENDPOINTS_ALL_FIELDS})

	if err != nil {
		return err
	}

	if format=="table" {
		m.PrettyPrint()
	} else {
		fmt.Println(m.Marshall())
	}

	return nil

}

func ShowAllEndpoints(accessToken, serviceId, format string) error {

	ec, err := GetCollection(accessToken, &MethodParams{ServiceId:serviceId}, &mashcli.Params{Fields: ENDPOINTS_ALL_FIELDS})
	if err != nil {
		return err
	}

	if format=="table" {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Endpoint ID", "Name", "Public Endpoint Address", "System Endpoint Address", "Created", "Updated"})
		table.SetCaption(true, "Service ID : "+serviceId)

		for _, e := range *ec {

			var publicDomain string
			if e.InboundSslRequired {
				publicDomain = fmt.Sprintf("https://%s%s", e.PublicDomains[0].Address, e.RequestPathAlias)
			} else {
				publicDomain = fmt.Sprintf("http://%s%s", e.PublicDomains[0].Address, e.RequestPathAlias)
			}

			var systemDomain, outboundTransportProtocol string = "", ""
			if e.OutboundTransportProtocol == "use-inbound" {
				if e.InboundSslRequired {
					outboundTransportProtocol = "https"
				} else {
					outboundTransportProtocol = "http"
				}
			} else {
				outboundTransportProtocol = e.OutboundTransportProtocol
			}
			if len(e.OutboundRequestTargetQueryParameters) > 0 {
				systemDomain = fmt.Sprintf("%s://%s%s?%s", outboundTransportProtocol, e.SystemDomains[0].Address, e.OutboundRequestTargetPath, e.OutboundRequestTargetQueryParameters)
			} else {
				systemDomain = fmt.Sprintf("%s://%s%s", outboundTransportProtocol, e.SystemDomains[0].Address, e.OutboundRequestTargetPath)
			}

			data := []string{e.Id, e.Name, publicDomain, systemDomain, e.Created[:19], e.Updated[:19]}
			table.Append(data)

		}
		table.Render()
	} else {
		b, err := json.MarshalIndent(ec, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(b))
	}

	return nil

}

func (e *Endpoints) PrettyPrint() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Endpoint ID", "Name", "Public Endpoint Address", "System Endpoint Address", "Created", "Updated"})
	var publicDomain string
	if e.InboundSslRequired {
		publicDomain = fmt.Sprintf("https://%s%s", e.PublicDomains[0].Address, e.RequestPathAlias)
	} else {
		publicDomain = fmt.Sprintf("http://%s%s", e.PublicDomains[0].Address, e.RequestPathAlias)
	}
	var systemDomain, outboundTransportProtocol string = "", ""
	if e.OutboundTransportProtocol == "use-inbound" {
		if e.InboundSslRequired {
			outboundTransportProtocol = "https"
		} else {
			outboundTransportProtocol = "http"
		}
	} else {
		outboundTransportProtocol = e.OutboundTransportProtocol
	}
	if len(e.OutboundRequestTargetQueryParameters) > 0 {
		systemDomain = fmt.Sprintf("%s://%s%s?%s", outboundTransportProtocol, e.SystemDomains[0].Address, e.OutboundRequestTargetPath, e.OutboundRequestTargetQueryParameters)
	} else {
		systemDomain = fmt.Sprintf("%s://%s%s", outboundTransportProtocol, e.SystemDomains[0].Address, e.OutboundRequestTargetPath)
	}
	data := []string{e.Id, e.Name, publicDomain, systemDomain, e.Created[:19], e.Updated[:19]}
	table.Append(data)
	table.Render()

	return

}

func Import(accessToken, filename string, mp *MethodParams)  (*Endpoints, error) {

	if len(filename) != 0 {
		e, err := ReadFile(filename)
		if err != nil {
			return nil, err
		}

		e, err = e.Create(accessToken, mp)
		if err != nil {
			return nil, err
		}

		e.WriteStdOut()
		return e, nil

	} else {
		e, err := ReadStdIn()
		if err != nil {
			return nil, err
		}

		e, err = e.Create(accessToken, mp)
		if err != nil {
			return nil, err
		}

		e.WriteStdOut()
		return e, nil
	}

}

func ExportAll(accessToken, path string, mp *MethodParams) error {

	m, err := GetCollection(accessToken, mp, &mashcli.Params{Fields: ENDPOINTS_ALL_FIELDS})

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	filename := path + "/endpoints.json"

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil

}

func Export(accessToken, serviceId, endpointId, filename string) error {

	e, err := Get(accessToken, &MethodParams{ServiceId: serviceId, EndpointId: endpointId}, &mashcli.Params{Fields: ENDPOINTS_ALL_FIELDS})
	if err != nil {
		return err
	}

	if len(filename) != 0 {
		e.WriteFile(filename)
	} else {
		e.WriteStdOut()
	}

	return nil
}

func ReadStdIn() (*Endpoints, error) {

	var data []byte

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	s := new(Endpoints)
	json.Unmarshal(data, &s)
	return s, nil

}

func (s *Endpoints) WriteStdOut() error {

	file := os.Stdout

	b, err := json.MarshalIndent(s, "", " ")
	if err == nil {
		s := string(b)
		file.WriteString(s)
		file.Sync()
		return nil
	} else {
		return err
	}
}

func (a *Endpoints) WriteFile(filename string) error {

	data, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil

}

func ReadFile(filename string) (*Endpoints, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(Endpoints)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (a *Endpoints) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
