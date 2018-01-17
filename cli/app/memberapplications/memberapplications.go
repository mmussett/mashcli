package memberapplications

import (
	"encoding/json"
	"fmt"

	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/mmussett/mashcli/cli/app/members"
	"github.com/olekukonko/tablewriter"

	"io/ioutil"
	"os"
)

func AddMemberApplication(accessToken, username, name, description string) error {

	// Go and fetch member

	memberId, err := members.GetMemberIdFromUsername(accessToken,username)
	if err != nil {
		return err
	}

	var ma = new(MemberApplications)

	ma.Name = name
	ma.Description = description
	ma.Username = username


	ma, err = ma.Create(accessToken,&MethodParams{MemberId:memberId})
	if err != nil {
		return err
	}

	memberApplicationAsString, err := ma.Marshall()
	if err != nil {
		return err
	}

	fmt.Println(memberApplicationAsString)

	return nil
}

func ShowMemberApplications(accessToken, memberId, format, filter string) error {

	mac := new([]MemberApplications)

	mac, err := GetCollection(accessToken, &MethodParams{MemberId:memberId}, &mashcli.Params{Fields: MEMBERAPPLICATIONS_ALL_FIELDS}, &mashcli.Filter{Filter:filter})
	if err != nil {
		return err
	}

	if format=="table" {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Application ID", "Name", "Description", "Username", "Created", "Updated"})
		table.SetCaption(true, "Member ID : "+memberId)

		for _, a := range *mac {
			data := []string{a.Id, a.Name, a.Description, a.Username, a.Created[:19], a.Updated[:19]}
			table.Append(data)

		}
		table.Render()
	} else {
		b, err := json.MarshalIndent(mac, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(b))
	}


	return nil

}

func (m *MemberApplications) PrettyPrint() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Application ID", "Name", "Description", "Username", "Created", "Updated"})
	data := []string{m.Id, m.Name, m.Description, m.Username, m.Created[:19], m.Updated[:19]}
	table.Append(data)
	table.Render()

	return

}

func Export(accessToken, memberId, filename string) error {

	mac, err := GetCollection(accessToken, &MethodParams{MemberId: memberId}, &mashcli.Params{Fields: MEMBERAPPLICATIONS_ALL_FIELDS}, &mashcli.Filter{Filter:""})
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(mac, "", "  ")
	if err != nil {
		return err
	}

	if len(filename) != 0 {
		err = ioutil.WriteFile(filename, data, 0644)
		if err != nil {
			return err
		}
	} else {
		file := os.Stdout
		s := string(data)
		_, err = file.WriteString(s)
		if err != nil {
			return err
		}
		err = file.Sync()
		if err != nil {
			return err
		}
	}

	return nil
}


func Import(accessToken, memberId, filename string)  (*MemberApplications, error) {

	if len(filename) != 0 {
		ma, err := ReadFile(filename)
		if err != nil {
			return nil, err
		}

		ma.Id = ""
		ma.Created = ""
		ma.Updated = ""

		ma, err = ma.Create(accessToken,&MethodParams{MemberId:memberId})
		if err != nil {
			return nil, err
		}

		ma.WriteStdOut()
		return ma, nil

	} else {
		ma, err := ReadStdIn()
		if err != nil {
			return nil, err
		}

		ma.Id = ""
		ma.Created = ""
		ma.Updated = ""

		ma, err = ma.Create(accessToken,&MethodParams{MemberId:memberId})
		if err != nil {
			return nil, err
		}

		ma.WriteStdOut()
		return ma, nil
	}

}

func (ma *MemberApplications) WriteStdOut() error {

	file := os.Stdout

	b, err := json.MarshalIndent(ma, "", " ")
	if err == nil {
		s := string(b)
		file.WriteString(s)
		file.Sync()
		return nil
	} else {
		return err
	}
}

func (ma *MemberApplications) WriteFile(filename string) error {

	data, err := json.MarshalIndent(ma, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil

}

func ReadStdIn() (*MemberApplications, error) {

	var data []byte

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	ma := new(MemberApplications)
	json.Unmarshal(data, &ma)
	return ma, nil

}

func ReadFile(filename string) (*MemberApplications, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	ma := new(MemberApplications)
	err = json.Unmarshal(data, &ma)
	if err != nil {
		return nil, err
	}

	return ma, nil

}


func (a *MemberApplications) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
