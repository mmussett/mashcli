package packages

import (
	"encoding/json"
	"github.com/mmussett/mashcli/cli/app"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"

	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)


func Nuke(accessToken string) error {

	pc := new([]Package)

	pc, err := GetCollection(accessToken, &mashcli.Params{Fields: PACKAGE_ALL_FIELDS}, &mashcli.Filter{Filter:""})
	if err != nil {
		return err
	}


	for _, p := range *pc {
		err := DeletePackage(accessToken, p.Id)
		if err != nil {
			return err
		}
	}

	return nil

}

func AddPackage(accessToken, name, description string) error {

	var p = new(Package)

	p.Name = name
	p.Description = description

	p, err := p.Create(accessToken)
	if err != nil {
		return err
	}

	packageAsString, err := p.Marshall()
	if err != nil {
		return err
	}

	fmt.Println(packageAsString)

	return nil

}

func DeletePackage(accessToken, packageId string) error {

	err := Delete(accessToken,&MethodParams{PackageId:packageId})
	if err != nil {
		return err
	}

	return nil
}

func ClonePackage(accessToken, packageId, name, description string) error {

	p, err := Get(accessToken, &MethodParams{PackageId:packageId},&mashcli.Params{Fields:PACKAGE_ALL_FIELDS})
	if err != nil {
		return err
	}

	if len(name) > 0 {
		p.Name = name
	}

	if len(description) > 0 {
		p.Description = description
	}

	p, err = p.Create(accessToken)
	if err != nil {
		return err
	}

	packageAsString, err := p.Marshall()
	if err != nil {
		return err
	}

	fmt.Println(packageAsString)

	return nil
}

func ShowPackage(accessToken, packageId, format string) error {
	m, err := Get(accessToken, &MethodParams{PackageId:packageId}, &mashcli.Params{Fields: PACKAGE_ALL_FIELDS})

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

func ShowAllPackages(accessToken, format, filter string) error {

	pc, err := GetCollection(accessToken, &mashcli.Params{Fields: PACKAGE_ALL_FIELDS},&mashcli.Filter{Filter:filter})
	if err != nil {
		return err
	}

	if format=="table" {
		// Work out maximum width for description
		var widthDescription = 16
		for _, p := range *pc {
			l := len(p.Description)
			if l > widthDescription {
				widthDescription = l
			}
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name", "Description", "Created", "Updated"})
		table.SetColMinWidth(2, widthDescription)

		for _, p := range *pc {
			data := []string{p.Id, p.Name, p.Description, p.Created[:19], p.Updated[:19]}
			table.Append(data)

		}
		table.Render()
	} else {
		b, err := json.MarshalIndent(pc, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(b))
	}


	return nil

}


func (p *Package) PrettyPrint() {


	var data []string
	var header []string

	if p.SharedSecretLength > 0 {
		data = []string{p.Id, p.Name, strconv.Itoa(p.NearQuotaThreshold), p.NotifyDeveloperPeriod, app.FormatBool(p.NotifyDeveloperNearQuota), app.FormatBool(p.NotifyDeveloperOverQuota), app.FormatBool(p.NotifyDeveloperOverThrottle), p.NotifyAdminPeriod, app.FormatBool(p.NotifyAdminNearQuota), app.FormatBool(p.NotifyAdminOverQuota), app.FormatBool(p.NotifyAdminOverThrottle), p.NotifyAdminEmails, strconv.Itoa(p.KeyLength), "ON", strconv.Itoa(p.SharedSecretLength),  p.Created[:19], p.Updated[:19]}
    header = []string{"Package ID", "Name", "NQT", "DP", "DNQ", "DOQ", "DOT", "AP", "ANQ", "AOQ", "AOT", "AE", "KL", "GSS", "SSL", "Created", "Updated"}
	} else {
		data = []string{p.Id, p.Name, strconv.Itoa(p.NearQuotaThreshold), p.NotifyDeveloperPeriod, app.FormatBool(p.NotifyDeveloperNearQuota), app.FormatBool(p.NotifyDeveloperOverQuota), app.FormatBool(p.NotifyDeveloperOverThrottle), p.NotifyAdminPeriod, app.FormatBool(p.NotifyAdminNearQuota), app.FormatBool(p.NotifyAdminOverQuota), app.FormatBool(p.NotifyAdminOverThrottle), p.NotifyAdminEmails, strconv.Itoa(p.KeyLength), "OFF", p.Created[:19], p.Updated[:19]}
		header = []string{"Package ID", "Name", "NQT", "DP", "DNQ", "DOQ", "DOT", "AP", "ANQ", "AOQ", "AOT", "AE", "KL", "GSS", "Created", "Updated"}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetAutoFormatHeaders(true)
	table.Append(data)
	table.Render()

	fmt.Println("NQT   = Near Quota Theshold")
	fmt.Println("DP    = Developer Period")
	fmt.Println("DNQ   = Developer Near Quota")
	fmt.Println("DOQ   = Developer Over Quota")
	fmt.Println("DOT   = Developer Over Throttle")
	fmt.Println("AP    = Administrator Period")
	fmt.Println("ANQ   = Administrator Near Quota")
	fmt.Println("AOQ   = Administrtor Over Quota")
	fmt.Println("AOT   = Administrator Over Throttle")
	fmt.Println("AE    = Administrator Emails")
	fmt.Println("KL    = Key Length")
	fmt.Println("GSS   = Generated Shared Secrets")
	fmt.Println("SSL   = Shared Secret Length")
	return

}

func Import(accessToken, filename string)  (*Package, error) {

	if len(filename) != 0 {
		p, err := ReadFile(filename)
		if err != nil {
			return nil, err
		}
		p.Id = ""
		p.Created = ""
		p.Updated = ""

		p, err = p.Create(accessToken)
		if err != nil {
			return nil, err
		}

		p.WriteStdOut()
		return p, nil

	} else {
		p, err := ReadStdIn()
		if err != nil {
			return nil, err
		}
		p.Id = ""
		p.Created = ""
		p.Updated = ""

		p, err = p.Create(accessToken)
		if err != nil {
			return nil, err
		}

		p.WriteStdOut()
		return p, nil
	}

}

func ExportAll(accessToken, path string) error {

	m, err := GetCollection(accessToken, &mashcli.Params{Fields: PACKAGE_ALL_FIELDS}, &mashcli.Filter{Filter:""})

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	filename := path + "/packages.json"

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil

}

func ImportAll(accessToken, filename string) error {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	a := new([]Package)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	for _,p := range *a {
		p.Id = ""
		_, err := p.Create(accessToken)
		if err != nil {
			return err
		}

	}

	return nil
}

func Export(accessToken, packageId, filename string) error {

	s, err := Get(accessToken, &MethodParams{PackageId: packageId}, &mashcli.Params{Fields: PACKAGE_ALL_FIELDS})
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

func (p *Package) WriteStdOut() error {

	file := os.Stdout

	b, err := json.MarshalIndent(p, "", " ")
	if err == nil {
		s := string(b)
		file.WriteString(s)
		file.Sync()
		return nil
	} else {
		return err
	}
}

func (p *Package) WriteFile(filename string) error {

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil

}

func ReadStdIn() (*Package, error) {

	var data []byte

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	s := new(Package)
	json.Unmarshal(data, &s)
	return s, nil

}

func ReadFile(filename string) (*Package, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(Package)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (a *Package) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
