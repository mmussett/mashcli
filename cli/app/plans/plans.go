package plans

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mmussett/mashcli/cli/app"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"

	"io/ioutil"
	"os"
	"strconv"
)

func Nuke(accessToken, packageId string) error {

	pc := new([]Plan)

	pc, err := GetCollection(accessToken, &MethodParams{PackageId:packageId}, &mashcli.Params{Fields: PLAN_ALL_FIELDS}, &mashcli.Filter{Filter:""})
	if err != nil {
		return err
	}


	for _, p := range *pc {
		err := DeletePlan(accessToken, packageId, p.Id)
		if err != nil {
			return err
		}
	}

	return nil

}

func SetStatus(accessToken, packageId, planId string, status string ) error {

	c, err := Get(accessToken, &MethodParams{PackageId:packageId,PlanId:planId},&mashcli.Params{Fields:PLAN_ALL_FIELDS})
	if err != nil {
		return err
	}

	var p = new(Plan)

	// Workaround for boolean omitempty... must set else we reset to false
	p.QPSLimitKeyOverrideAllowed = c.QPSLimitKeyOverrideAllowed
	p.RateLimitKeyOverrideAllowed = c.RateLimitKeyOverrideAllowed
	p.AdminKeyProvisioningEnabled = c.AdminKeyProvisioningEnabled
	p.SelfServiceKeyProvisioningEnabled = c.SelfServiceKeyProvisioningEnabled

	p.Status = status

	p, err = p.Update(accessToken,&MethodParams{PackageId:packageId,PlanId:planId})
	if err != nil {
		return err
	}

	return nil

}

func SetKeyProperties(accessToken, packageId, planId, selfServiceKeyProvisioning, adminKeyProvisioning string, maxAllowableKeys, maxKeysUntilModeration  int ) error {

	c, err := Get(accessToken, &MethodParams{PackageId:packageId,PlanId:planId},&mashcli.Params{Fields:PLAN_ALL_FIELDS})
	if err != nil {
		return err
	}


	var p = new(Plan)

	// Workaround for boolean omitempty... must set else we reset to false
	p.QPSLimitKeyOverrideAllowed = c.QPSLimitKeyOverrideAllowed
	p.RateLimitKeyOverrideAllowed = c.RateLimitKeyOverrideAllowed

	if maxAllowableKeys >= 0 {
		p.MaxNumKeysAllowed = maxAllowableKeys
	}

	if maxKeysUntilModeration >= 0 {
		p.NumKeysBeforeReview = maxKeysUntilModeration
	}


	switch selfServiceKeyProvisioning {
	case "true":
		p.SelfServiceKeyProvisioningEnabled = true
	case "false":
		p.SelfServiceKeyProvisioningEnabled = false
	}


	switch adminKeyProvisioning {
	case "true":
		p.AdminKeyProvisioningEnabled = true
	case "false":
		p.AdminKeyProvisioningEnabled = false
	}


	p, err = p.Update(accessToken,&MethodParams{PackageId:packageId,PlanId:planId})
	if err != nil {
		return err
	}

	return nil

}



func SetRateLimits(accessToken, packageId, planId string, throttle int, throttleOverride string, quota int, quotaPeriod string, quotaOverride string) error {

	c, err := Get(accessToken, &MethodParams{PackageId:packageId,PlanId:planId},&mashcli.Params{Fields:PLAN_ALL_FIELDS})
	if err != nil {
		return err
	}

	var p = new(Plan)

	// Workaround for boolean omitempty... must set else we reset to false
	p.AdminKeyProvisioningEnabled = c.AdminKeyProvisioningEnabled
	p.SelfServiceKeyProvisioningEnabled = c.SelfServiceKeyProvisioningEnabled

	if throttle >= 0 {
		p.QPSLimitCeiling = throttle
	}

	if quota >= 0 {
		p.RateLimitCeiling = quota
	}

	switch throttleOverride {
	case "true":
		p.QPSLimitKeyOverrideAllowed = true
	case "false":
		p.QPSLimitKeyOverrideAllowed = false
	}

	switch quotaOverride {
	case "true":
		p.RateLimitKeyOverrideAllowed = true
	case "false":
		p.RateLimitKeyOverrideAllowed = false
	}

	switch strings.ToLower(quotaPeriod) {
	case "minute":
		p.RateLimitPeriod = "minute"
		fallthrough
	case "hour":
		p.RateLimitPeriod = "hour"
		fallthrough
	case "day":
		p.RateLimitPeriod = "day"
		fallthrough
	case "month":
		p.RateLimitPeriod = "month"
	}

	p, err = p.Update(accessToken,&MethodParams{PackageId:packageId,PlanId:planId})
	if err != nil {
		return err
	}

	return nil

}



func AddPlan(accessToken, packageId, name, description string) error {

	var p = new(Plan)

	p.Name = name
	p.Description = description

	p, err := p.Create(accessToken,&MethodParams{PackageId:packageId})
	if err != nil {
		return err
	}

	planAsString, err := p.Marshall()
	if err != nil {
		return err
	}

	fmt.Println(planAsString)

	return nil

}

func ClonePlan(accessToken, packageId, planId, name, description string) error {

	p, err := Get(accessToken, &MethodParams{PackageId:packageId,PlanId:planId},&mashcli.Params{Fields:PLAN_ALL_FIELDS})
	if err != nil {
		return err
	}

	if len(name) > 0 {
		p.Name = name
	}

	if len(description) > 0 {
		p.Description = description
	}

	p, err = p.Create(accessToken,&MethodParams{PackageId:packageId})
	if err != nil {
		return err
	}

	planAsString, err := p.Marshall()
	if err != nil {
		return err
	}

	fmt.Println(planAsString)

	return nil
}


func DeletePlan(accessToken, packageId, planId string) error {

	err := Delete(accessToken,&MethodParams{PackageId:packageId,PlanId:planId})
	if err != nil {
		return err
	}

	return nil
}

func ShowPlan(accessToken, packageId, planId, format string) error {
	m, err := Get(accessToken, &MethodParams{PackageId:packageId,PlanId:planId}, &mashcli.Params{Fields: PLAN_ALL_FIELDS})

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

func ShowAllPlans(accessToken, packageId, format, filter  string) error {

	pc, err := GetCollection(accessToken, &MethodParams{PackageId:packageId}, &mashcli.Params{Fields: PLAN_ALL_FIELDS}, &mashcli.Filter{Filter:filter})
	if err != nil {
		return err
	}

	if format=="table" {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Plan ID", "Name", "Description", "Status", "SSK", "AK", "MK", "KM", "T", "APKTO", "Q", "QP", "APKQO", "Created", "Updated"})
		table.SetCaption(true, "Package ID : "+packageId)

		for _, p := range *pc {
			var d= ""
			if len(p.Description) > 20 {
				d = p.Description[:20] + "..."
			}
			data := []string{p.Id, p.Name, d, p.Status, app.FormatBool(p.SelfServiceKeyProvisioningEnabled), app.FormatBool(p.AdminKeyProvisioningEnabled), strconv.Itoa(p.MaxNumKeysAllowed), strconv.Itoa(p.NumKeysBeforeReview), strconv.Itoa(p.QPSLimitCeiling), app.FormatBool(p.RateLimitKeyOverrideAllowed), strconv.Itoa(p.RateLimitCeiling), p.RateLimitPeriod, app.FormatBool(p.QPSLimitKeyOverrideAllowed), p.Created[:19], p.Updated[:19]}
			table.Append(data)
		}
		table.Render()

		fmt.Println("\nSSK   = Self-Service Key Provisioning")
		fmt.Println("AK    = Admin Key Provisioning")
		fmt.Println("MK    = Maximum Allowable Keys")
		fmt.Println("KM    = Keys Until Moderation")
		fmt.Println("T     = Throttle")
		fmt.Println("APKTO = Allow Package Key Throttle Override")
		fmt.Println("Q     = Quota")
		fmt.Println("QP    = Quota Period")
		fmt.Println("APKQO = Allow Package Key Quota Override")
	} else {
		b, err := json.MarshalIndent(pc, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(b))
	}

	return nil

}


func (p *Plan) PrettyPrint() {

	var d = ""
	if len(p.Description)>20 {
		d = p.Description[:20] + "..."
	}
	data := []string{p.Id, p.Name, d, p.Status, app.FormatBool(p.SelfServiceKeyProvisioningEnabled), app.FormatBool(p.AdminKeyProvisioningEnabled), strconv.Itoa(p.MaxNumKeysAllowed), strconv.Itoa(p.NumKeysBeforeReview), strconv.Itoa(p.QPSLimitCeiling), app.FormatBool(p.RateLimitKeyOverrideAllowed), strconv.Itoa(p.RateLimitCeiling), p.RateLimitPeriod, app.FormatBool(p.QPSLimitKeyOverrideAllowed), p.Created[:19], p.Updated[:19]}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Plan ID", "Name", "Description", "Status", "SSK", "AK", "MK", "KM", "T", "APKTO", "Q", "QP", "APKQO", "Created", "Updated"})
	table.SetAutoFormatHeaders(true)
	table.Append(data)
	table.Render()

	fmt.Println("\nSSK   = Self-Service Key Provisioning")
	fmt.Println("AK    = Admin Key Provisioning")
	fmt.Println("MK    = Maximum Allowable Keys")
	fmt.Println("KM    = Keys Until Moderation")
	fmt.Println("T     = Throttle")
	fmt.Println("APKTO = Allow Package Key Throttle Override")
	fmt.Println("Q     = Quota")
	fmt.Println("QP    = Quota Period")
	fmt.Println("APKQO = Allow Package Key Quota Override")
	return

}


func Import(accessToken, packageId, filename string)  (*Plan, error) {

	if len(filename) != 0 {
		p, err := ReadFile(filename)
		if err != nil {
			return nil, err
		}

		p, err = p.Create(accessToken,&MethodParams{PackageId:packageId})
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
		p.Services = nil
		p, err = p.Create(accessToken,&MethodParams{PackageId:packageId})
		if err != nil {
			return nil, err
		}

		p.WriteStdOut()
		return p, nil
	}

}

func ExportAll(accessToken, path string, mp *MethodParams) error {

	m, err := GetCollection(accessToken, mp, &mashcli.Params{Fields: PLAN_ALL_FIELDS},&mashcli.Filter{Filter:""})

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	filename := path + "/plans.json"

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil

}

func ImportAll(accessToken, filename string, mp *MethodParams) error {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	a := new([]Plan)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	for _,p := range *a {
		p.Id = ""
		_, err := p.Create(accessToken,mp)
		if err != nil {
			return err
		}

	}

	return nil
}

func Export(accessToken, packageId, planId, filename string) error {

	s, err := Get(accessToken, &MethodParams{PackageId: packageId,PlanId:planId}, &mashcli.Params{Fields: PLAN_ALL_FIELDS})
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

func ReadStdIn() (*Plan, error) {

	var data []byte

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	s := new(Plan)
	json.Unmarshal(data, &s)
	return s, nil

}

func ReadFile(filename string) (*Plan, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(Plan)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (p *Plan) WriteStdOut() error {

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

func (p *Plan) WriteFile(filename string) error {

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

func (a *Plan) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}



