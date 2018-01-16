package packagekeys

import (
	"encoding/json"
	"fmt"

	"github.com/mmussett/mashcli/cli/app"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"

	"io/ioutil"
	"os"
	"strconv"
)

func Nuke(accessToken string, preview bool) error {

	pkc := new([]PackageKeys)

	pkc, err := GetCollection(accessToken, &mashcli.Params{Fields: PACKAGEKEYS_ALL_FIELDS}, &mashcli.Filter{Filter:""})
	if err != nil {
		return err
	}


		for _, pk := range *pkc {
			if !preview {
				err := DeletePackageKey(accessToken, pk.Id)
				if err != nil {
					return err
				}
			} else {
				fmt.Println("Preview Deleting Package Key "+pk.Apikey)
			}
		}

	return nil

}

func SetStatus(accessToken, packageKeyId, status string ) error {

	_, err := Get(accessToken, &MethodParams{PackageKeyId:packageKeyId},&mashcli.Params{Fields:PACKAGEKEYS_ALL_FIELDS})
	if err != nil {
		return err
	}

	p := new(PackageKeys)
	p.Status = status

	p, err = p.Update(accessToken,&MethodParams{PackageKeyId:packageKeyId})
	if err != nil {
		return err
	}

	return nil

}

func SetThrottleRateDefault(accessToken, packageKeyId string) error {

	c, err := Get(accessToken, &MethodParams{PackageKeyId:packageKeyId},&mashcli.Params{Fields:PACKAGEKEYS_ALL_FIELDS})
	if err != nil {
		return err
	}

	p := new(PackageKeys)
	p.RateLimitCeiling = c.RateLimitCeiling
	p.RateLimitExempt = c.RateLimitExempt
	p.QPSLimitCeiling = 0
	p.QPSLimitExempt = false

	p, err = p.Update(accessToken,&MethodParams{PackageKeyId:packageKeyId})
	if err != nil {
		return err
	}

	return nil

}

func SetThrottleRate(accessToken, packageKeyId string, rate int) error {

	c, err := Get(accessToken, &MethodParams{PackageKeyId:packageKeyId},&mashcli.Params{Fields:PACKAGEKEYS_ALL_FIELDS})
	if err != nil {
		return err
	}

	p := new(PackageKeys)
	p.RateLimitCeiling = c.RateLimitCeiling
	p.RateLimitExempt = c.RateLimitExempt
	p.QPSLimitExempt = false
  p.QPSLimitCeiling = rate

	p, err = p.Update(accessToken,&MethodParams{PackageKeyId:packageKeyId})
	if err != nil {
		return err
	}

	return nil

}


func SetQuotaRateDefault(accessToken, packageKeyId string) error {

	c, err := Get(accessToken, &MethodParams{PackageKeyId:packageKeyId},&mashcli.Params{Fields:PACKAGEKEYS_ALL_FIELDS})
	if err != nil {
		return err
	}

	p := new(PackageKeys)
	p.RateLimitCeiling = 0
	p.RateLimitExempt = false
	p.QPSLimitExempt = c.QPSLimitExempt
	p.QPSLimitCeiling = c.QPSLimitCeiling

	p, err = p.Update(accessToken,&MethodParams{PackageKeyId:packageKeyId})
	if err != nil {
		return err
	}

	return nil

}

func SetQuotaRate(accessToken, packageKeyId string, rate int) error {

	c, err := Get(accessToken, &MethodParams{PackageKeyId:packageKeyId},&mashcli.Params{Fields:PACKAGEKEYS_ALL_FIELDS})
	if err != nil {
		return err
	}

	p := new(PackageKeys)
	p.QPSLimitCeiling = c.QPSLimitCeiling
	p.QPSLimitExempt = c.QPSLimitExempt
	p.RateLimitExempt = false
	p.RateLimitCeiling = rate

	p, err = p.Update(accessToken,&MethodParams{PackageKeyId:packageKeyId})
	if err != nil {
		return err
	}

	return nil

}

func SetThrottleRateUnlimited(accessToken, packageKeyId string) error {

	c, err := Get(accessToken, &MethodParams{PackageKeyId:packageKeyId},&mashcli.Params{Fields:PACKAGEKEYS_ALL_FIELDS})
	if err != nil {
		return err
	}

	p := new(PackageKeys)
	p.RateLimitExempt = c.RateLimitExempt
	p.RateLimitCeiling = c.RateLimitCeiling
	p.QPSLimitExempt = true
	p.QPSLimitCeiling = 0

	p, err = p.Update(accessToken,&MethodParams{PackageKeyId:packageKeyId})
	if err != nil {
		return err
	}

	return nil

}

func SetQuotaRateUnlimited(accessToken, packageKeyId string) error {

	c, err := Get(accessToken, &MethodParams{PackageKeyId:packageKeyId},&mashcli.Params{Fields:PACKAGEKEYS_ALL_FIELDS})
	if err != nil {
		return err
	}

	p := new(PackageKeys)
	p.QPSLimitExempt = c.QPSLimitExempt
	p.QPSLimitCeiling = c.QPSLimitCeiling
	p.RateLimitExempt = true
	p.RateLimitCeiling = 0

	p, err = p.Update(accessToken,&MethodParams{PackageKeyId:packageKeyId})
	if err != nil {
		return err
	}

	return nil

}


func DeletePackageKey(accessToken, packageKeyId string) error {

	err := Delete(accessToken,&MethodParams{PackageKeyId:packageKeyId})
	if err != nil {
		return err
	}

	return nil
}

func ShowPackageKeys(accessToken, packageKeyId, format  string) error {


	var a, err = Get(accessToken, &MethodParams{PackageKeyId:packageKeyId}, &mashcli.Params{Fields: PACKAGEKEYS_ALL_FIELDS})

	if err != nil {
		return err
	}

	if format=="table" {
		a.PrettyPrint()
	} else {
		fmt.Println(a.Marshall())
	}

	return nil
}

func ShowAllPackageKeys(accessToken, format, filter string) error {

	a := new([]PackageKeys)

	a, err := GetCollection(accessToken, &mashcli.Params{Fields: PACKAGEKEYS_ALL_FIELDS}, &mashcli.Filter{Filter:filter})

	if err != nil {
		return err
	}

	if format=="table" {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Package Key ID", "Package Key", "QPS Limit Ceiling", "QPS Limit Exempt", "Rate Limit Ceiling", "Rate Limit Exempt", "Status", "Created", "Updated"})

		for _, s := range *a {
			data := []string{s.Id, s.Apikey, strconv.Itoa(s.QPSLimitCeiling), app.FormatBool(s.QPSLimitExempt), strconv.Itoa(s.RateLimitCeiling), app.FormatBool(s.RateLimitExempt), s.Status, s.Created[:19], s.Updated[:19]}
			table.Append(data)

		}
		table.Render()
	} else {
		b, err := json.MarshalIndent(a, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(b))
	}

	return nil

}

func (a *PackageKeys) PrettyPrint() {


	data := []string{a.Apikey, strconv.Itoa(a.QPSLimitCeiling), app.FormatBool(a.QPSLimitExempt), strconv.Itoa(a.RateLimitCeiling), app.FormatBool(a.RateLimitExempt), a.Status, a.Created[:19], a.Updated[:19]}
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Package Key", "QPS Limit Ceiling", "QPS Limit Exempt", "Rate Limit Ceiling", "Rate Limit Exempt", "Status", "Created", "Updated"})
	table.Append(data)
	table.Render()

	return

}


func ExportAll(accessToken, path string) error {

	m, err := GetCollection(accessToken, &mashcli.Params{Fields: PACKAGEKEYS_ALL_FIELDS},&mashcli.Filter{Filter:""})

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	filename := path + "/packagekeys.json"

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

	a := new([]PackageKeys)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	for _,p := range *a {
		var mp = &MethodParams{PackageKeyId:p.Id}
		p.Id = ""
		_, err := p.Update(accessToken, mp)
		if err != nil {
			return err
		}

	}

	return nil
}

func Import(accessToken, filename string)  (*PackageKeys, error) {

	if len(filename) != 0 {
		p, err := ReadFile(filename)
		if err != nil {
			return nil, err
		}

		var packageKeyId = p.Id
		p.Id = ""
		p, err = p.Update(accessToken,&MethodParams{PackageKeyId:packageKeyId})
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
		var packageKeyId = p.Id
		p.Id = ""
		p, err = p.Update(accessToken,&MethodParams{PackageKeyId:packageKeyId})
		if err != nil {
			return nil, err
		}

		p.WriteStdOut()
		return p, nil
	}

}

func Export(accessToken, filename, packageKeyId string) error {

	s, err := Get(accessToken, &MethodParams{PackageKeyId:packageKeyId}, &mashcli.Params{Fields: PACKAGEKEYS_ALL_FIELDS})
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

func (p *PackageKeys) WriteStdOut() error {

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

func (p *PackageKeys) WriteFile(filename string) error {

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

func ReadStdIn() (*PackageKeys, error) {

	var data []byte

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	s := new(PackageKeys)
	json.Unmarshal(data, &s)
	return s, nil

}

func ReadFile(filename string) (*PackageKeys, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(PackageKeys)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}


func (a *PackageKeys) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}