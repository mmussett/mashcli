package services

import (
	"encoding/json"
	"fmt"
	"github.com/mmussett/mashcli/cli/app/mashcli"

	"github.com/olekukonko/tablewriter"

	"io/ioutil"
	"os"
	"strconv"
)


func Nuke(accessToken string) error {

	sc := new([]Services)

	sc, err := GetCollection(accessToken, &mashcli.Params{Fields: SERVICES_ALL_FIELDS}, &mashcli.Filter{Filter:""})
	if err != nil {
		return err
	}


	for _, s := range *sc {
		err := DeleteService(accessToken, s.Id)
		if err != nil {
			return err
		}
	}

	return nil

}

func DeleteService(accessToken, serviceId string) error {

	err := Delete(accessToken, &MethodParams{ServiceId: serviceId})
	if err != nil {
		return err
	}

	return nil
}

func AddService(accessToken, name, version, description string, aggregateQps int64) error {

	var service = new(Services)

	service.Name = name
	service.Version = version
	service.Description = description
	service.QpsLimitOverall = aggregateQps

	service, err := service.Create(accessToken)
	if err != nil {
		return err
	}

	serviceAsString, err := service.Marshall()
	if err != nil {
		return err
	}

	fmt.Println(serviceAsString)

	return nil

}

func CloneService(accessToken, serviceId string) error {

	s, err := Get(accessToken, &MethodParams{ServiceId: serviceId}, &mashcli.Params{Fields: SERVICES_ALL_FIELDS})
	if err != nil {
		return err
	}

	version, err := strconv.ParseFloat(s.Version, 32)
	if err != nil {
		s.Version = "1.0"
	} else {
		version++
		s.Version = strconv.FormatFloat(version, 'f', 1, 32)
	}

	s, err = s.Create(accessToken)
	if err != nil {
		return err
	}

	serviceAsString, err := s.Marshall()
	if err != nil {
		return err
	}

	fmt.Println(serviceAsString)

	return nil
}

func ShowService(accessToken,serviceId, format string) error {


	s, err := Get(accessToken, &MethodParams{ServiceId:serviceId}, &mashcli.Params{Fields: SERVICES_ALL_FIELDS})
	if err != nil {
		return err
	}

	if format=="table" {
		s.PrettyPrint()
	} else {
		fmt.Println(s.Marshall())
	}

	return nil

}

func ShowAllServices(accessToken, format, filter string) error {

	sc, err := GetCollection(accessToken, &mashcli.Params{Fields: SERVICES_ALL_FIELDS}, &mashcli.Filter{Filter:filter})

	if err != nil {
		return err
	}


	if format=="table" {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Service ID", "Name", "Description", "Agg. QPS", "Version", "Created", "Updated"})

		// Work out maximum width for description
		const maxWidthDescription= 32
		const minWidthDescription= 16
		const maxWidthName= 48
		const minWidthName= 16
		const maxWidthVersion= 16
		const minWidthVersion= 16

		var widthDescription= minWidthDescription
		for _, s := range *sc {
			l := len(s.Description)
			if l >= maxWidthDescription {
				widthDescription = maxWidthDescription
			} else if l < minWidthDescription {
				widthDescription = minWidthDescription
			} else {
				widthDescription = l
			}
		}

		table.SetColMinWidth(2, widthDescription)

		// Work out maximum width for name
		var widthName= minWidthName
		for _, s := range *sc {
			l := len(s.Name)
			if l >= maxWidthName {
				widthName = maxWidthName
			} else if l < minWidthName {
				widthName = minWidthName
			} else {
				widthName = l
			}
		}

		table.SetColMinWidth(1, widthName)

		// Work out maximum width for version
		var widthVersion= minWidthVersion
		for _, s := range *sc {
			l := len(s.Version)
			if l >= maxWidthVersion {
				widthVersion = maxWidthVersion
			} else if l < minWidthVersion {
				widthVersion = minWidthVersion
			} else {
				widthVersion = l
			}
		}

		table.SetColMinWidth(4, widthVersion)

		for _, s := range *sc {
			var desc= ""
			if len(s.Description) >= maxWidthDescription {
				desc = s.Description[:maxWidthDescription-2]
			} else {
				desc = s.Description
			}
			data := []string{s.Id, s.Name, desc, strconv.FormatInt(s.QpsLimitOverall, 10), s.Version, s.Created[:19], s.Updated[:19]}
			table.Append(data)
		}
		table.Render()
	} else {

		b, err := json.MarshalIndent(sc, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(b))
	}

	return nil

}

func (s *Services) PrettyPrint() {

	data := []string{s.Id, s.Name, s.Description, strconv.FormatInt(s.QpsLimitOverall, 10), s.Version, s.Created[:19], s.Updated[:19]}
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Service ID", "Name", "Description", "Agg. QPS", "Version", "Created", "Updated"})
	table.Append(data)
	table.Render()

	return

}

func Import(accessToken, filename string) (*Services, error) {

	if len(filename) != 0 {
		s, err := ReadFile(filename)
		if err != nil {
			return nil, err
		}
		s.Id = ""
		s.Created = ""
		s.Updated = ""

		s, err = s.Create(accessToken)
		if err != nil {
			return nil, err
		}

		s.WriteStdOut()
		return s, nil

	} else {
		s, err := ReadStdIn()
		if err != nil {
			return nil, err
		}
		s.Id = ""
		s.Created = ""
		s.Updated = ""

		s, err = s.Create(accessToken)
		if err != nil {
			return nil, err
		}

		s.WriteStdOut()
		return s, nil
	}

}

func Export(accessToken, serviceId, filename string) error {

	s, err := Get(accessToken, &MethodParams{ServiceId: serviceId}, &mashcli.Params{Fields: SERVICES_ALL_FIELDS})
	if err != nil {
		return err
	}

	if len(filename) != 0 {
		s.WriteFile(filename)
	} else {
		s.WriteStdOut()
	}

	return nil
}

func ExportAll(accessToken string, dirPath string) error {

	servicesCollection, err := GetCollection(accessToken, &mashcli.Params{Fields: SERVICES_ALL_FIELDS},&mashcli.Filter{Filter:""})
	if err != nil {
		return err
	}

	for _, service := range *servicesCollection {
		filename := fmt.Sprintf("%s/service-%s-%s.json", dirPath, service.Id, service.Name)
		err := service.WriteFile(filename)
		if err != nil {
			return err
		}
	}

	return nil

}

func (s *Services) WriteStdOut() error {

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

func (s *Services) WriteFile(filename string) error {

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil

}

func ReadStdIn() (*Services, error) {

	var data []byte

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	s := new(Services)
	json.Unmarshal(data, &s)
	return s, nil

}

func ReadFile(filename string) (*Services, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(Services)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (a *Services) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
