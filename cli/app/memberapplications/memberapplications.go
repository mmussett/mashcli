package memberapplications

import (
	"encoding/json"
	"fmt"

	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"

	"io/ioutil"
	"os"
)

func ShowMemberApplications(accessToken, memberId, format string) error {

	mac := new([]MemberApplications)

	mac, err := GetCollection(accessToken, &MethodParams{MemberId:memberId}, &mashcli.Params{Fields: MEMBERAPPLICATIONS_ALL_FIELDS})
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

	mac, err := GetCollection(accessToken, &MethodParams{MemberId: memberId}, &mashcli.Params{Fields: MEMBERAPPLICATIONS_ALL_FIELDS})
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


func (a *MemberApplications) Marshall() (string, error) {

	b, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
