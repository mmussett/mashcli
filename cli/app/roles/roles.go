package roles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gobwas/glob"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"
)

func Nuke(accessToken string, preview bool) error {

	rc := new([]Roles)

	rc, err := GetCollection(accessToken, &mashcli.Params{Fields: ROLES_ALL_FIELDS}, &mashcli.Filter{Filter: ""})
	if err != nil {
		return err
	}

	for _, r := range *rc {
		if !((r.Name == "Administrator") || (r.Name == "API Manager") || (r.Name == "Call Inspector Administrator") || (r.Name == "Call Inspector User") || (r.Name == "Community Manager") || (r.Name == "Content Manager") || (r.Name == "Everyone") || (r.Name == "Member") || (r.Name == "Portal Manager") || (r.Name == "Program Manager") || (r.Name == "Reports User") || (r.Name == "Support User"))  {
			if !preview {
				err := DeleteRole(accessToken, r.Id)
				if err != nil {
					return err
				}
			} else {
				fmt.Println("Preview Delete Role "+r.Name)
			}

		}
	}

	return nil

}

func ShowRole(accessToken, roleId, format string) error {

	r, err := Get(accessToken, &MethodParams{RoleId: roleId}, &mashcli.Params{Fields: ROLES_ALL_FIELDS})
	if err != nil {
		return err
	}

	if format == "table" {
		r.PrettyPrint()
	} else {
		fmt.Println(r.Marshall())
	}
	return nil

}

func ShowAllRoles(accessToken, format, filter, nameglob string) error {

	rc, err := GetCollection(accessToken, &mashcli.Params{Fields: ROLES_ALL_FIELDS}, &mashcli.Filter{Filter: filter})

	if err != nil {
		return err
	}

	var g glob.Glob
	if nameglob == "" {
		g = glob.MustCompile("*")
	} else {
		g = glob.MustCompile(nameglob)
	}

	if format == "table" {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Role ID", "Name", "Created", "Updated"})

		for _, r := range *rc {

			if g.Match(r.Name) {
				data := []string{r.Id, r.Name, r.Created[:19], r.Updated[:19]}
				table.Append(data)
			}
		}
		table.Render()
	} else {
		b, err := json.MarshalIndent(rc, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(b))
	}

	return nil

}

func (m *Roles) PrettyPrint() {

	data := []string{m.Name, m.Created[:19], m.Updated[:19]}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Created", "Updated"})
	table.Append(data)
	table.Render()

	return

}

func DeleteRole(accessToken, roleId string) error {

	err := Delete(accessToken, &MethodParams{RoleId: roleId})
	if err != nil {
		return err
	}

	return nil
}

func AddRole(accessToken, name string) error {

	var r = new(Roles)

	r.Name = name

	r, err := r.Create(accessToken)
	if err != nil {
		return err
	}

	roleAsString, err := r.Marshall()
	if err != nil {
		return err
	}

	fmt.Println(roleAsString)

	return nil

}

func Import(accessToken, filename string) (*Roles, error) {

	if len(filename) != 0 {
		r, err := ReadFile(filename)
		if err != nil {
			return nil, err
		}
		r.Id = ""

		r, err = r.Create(accessToken)
		if err != nil {
			return nil, err
		}

		r.WriteStdOut()
		return r, nil

	} else {
		r, err := ReadStdIn()
		if err != nil {
			return nil, err
		}
		r.Id = ""

		r, err = r.Create(accessToken)
		if err != nil {
			return nil, err
		}

		r.WriteStdOut()
		return r, nil
	}

}

func Export(accessToken, roleId, filename string) error {

	r, err := Get(accessToken, &MethodParams{RoleId: roleId}, &mashcli.Params{Fields: ROLES_ALL_FIELDS})
	if err != nil {
		return err
	}

	if len(filename) != 0 {
		r.WriteFile(filename)
	} else {
		r.WriteStdOut()
	}

	return nil
}

func ExportAll(accessToken string, dirPath string) error {

	rc, err := GetCollection(accessToken, &mashcli.Params{Fields: ROLES_ALL_FIELDS}, &mashcli.Filter{Filter: ""})
	if err != nil {
		return err
	}

	for _, r := range *rc {
		filename := fmt.Sprintf("%s/roles-%s-%s.json", dirPath, r.Id, r.Name)
		err := r.WriteFile(filename)
		if err != nil {
			return err
		}
	}

	return nil

}

func ImportAll(accessToken, filename string) error {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	a := new([]Roles)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	for _, p := range *a {
		p.Id = ""
		_, err := p.Create(accessToken)
		if err != nil {
			return err
		}

	}

	return nil
}

func (r *Roles) WriteStdOut() error {

	file := os.Stdout

	b, err := json.MarshalIndent(r, "", " ")
	if err == nil {
		s := string(b)
		file.WriteString(s)
		file.Sync()
		return nil
	} else {
		return err
	}
}

func (r *Roles) WriteFile(filename string) error {

	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil

}

func ReadStdIn() (*Roles, error) {

	var data []byte

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	s := new(Roles)
	json.Unmarshal(data, &s)
	return s, nil

}

func ReadFile(filename string) (*Roles, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(Roles)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (r *Roles) Marshall() (string, error) {

	b, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
