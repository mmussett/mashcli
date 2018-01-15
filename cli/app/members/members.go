package members

import (
	"encoding/json"
	"strings"

	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"

	"fmt"
	"io/ioutil"
	"os"
)

func Nuke(accessToken string) error {

	mc := new([]Members)

	mc, err := GetCollection(accessToken, &mashcli.Params{Fields: MEMBERS_ALL_FIELDS})
	if err != nil {
		return err
	}


	for _, m := range *mc {
		if !(m.Username == "MasheryInternalOAuth2") && !(strings.Contains(strings.ToLower(m.LastName),"admin")) && !(strings.Contains(strings.ToLower(m.Email),"admin")){
			fmt.Println(m.Username)
			err := DeleteMember(accessToken, m.Id)
			if err != nil {
				return err
			}
		}
	}

	return nil

}
func SetStatus(accessToken, memberId, status string ) error {

	_, err := Get(accessToken, &MethodParams{MemberId:memberId},&mashcli.Params{Fields:MEMBERS_ALL_FIELDS})
	if err != nil {
		return err
	}

	var p = new(Members)

	p.AreaStatus = status

	p, err = p.Update(accessToken,&MethodParams{MemberId:memberId})
	if err != nil {
		return err
	}

	return nil

}

func AddMember(accessToken, email, username, displayname string) error {

	var m = new(Members)

	m.Email = email
	m.Username = username
	m.DisplayName = displayname

	p, err := m.Create(accessToken)
	if err != nil {
		return err
	}

	memberAsString, err := p.Marshall()
	if err != nil {
		return err
	}

	fmt.Println(memberAsString)

	return nil

}

func DeleteMember(accessToken, memberId string) error {

	err := Delete(accessToken,&MethodParams{MemberId:memberId})
	if err != nil {
		return err
	}

	return nil
}

func ShowMember(accessToken, memberId, format string) error {

	m, err := Get(accessToken, &MethodParams{MemberId:memberId}, &mashcli.Params{Fields: MEMBERS_ALL_FIELDS})

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

func ShowAllMembers(accessToken, format string) error {

	ac, err := GetCollection(accessToken, &mashcli.Params{Fields: MEMBERS_ALL_FIELDS})
	if err != nil {
		return err
	}

	if format=="table" {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Member ID", "Username", "Email", "Display Name", "Company", "First Name", "Last Name", "Status", "Created", "Updated"})

		for _, m := range *ac {
			data := []string{m.Id, m.Username, m.Email, m.DisplayName, m.Company, m.FirstName, m.LastName, m.AreaStatus, m.Created[:19], m.Updated[:19]}
			table.Append(data)
		}
		table.Render()
	} else {
		b, err := json.MarshalIndent(ac, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(b))
	}

	return nil

}

func (m *Members) PrettyPrint() {

	data := []string{m.Id, m.Username, m.Email, m.DisplayName, m.Company, m.FirstName, m.LastName, m.AreaStatus, m.Created[:19], m.Updated[:19]}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Member ID", "Username", "Email", "Display Name", "Company", "First Name", "Last Name", "Status", "Created", "Updated"})
	table.Append(data)
	table.Render()

	return

}

func ExportAll(accessToken string, filename string) error {

	m, err := GetCollection(accessToken, &mashcli.Params{Fields: MEMBERS_ALL_FIELDS})

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil

}


func Export(accessToken, memberId, filename string) error {

	s, err := Get(accessToken, &MethodParams{MemberId: memberId}, &mashcli.Params{Fields: MEMBERS_ALL_FIELDS})
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

func Import(accessToken, filename string)  (*Members, error) {

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

func (p *Members) WriteStdOut() error {

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

func (p *Members) WriteFile(filename string) error {

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

func ReadStdIn() (*Members, error) {

	var data []byte

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	s := new(Members)
	json.Unmarshal(data, &s)
	return s, nil

}

func ReadFile(filename string) (*Members, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	a := new(Members)
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	return a, nil

}

func (a *Members) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
