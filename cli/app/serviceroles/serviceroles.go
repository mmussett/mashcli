package serviceroles

import (
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"
	"os"
)

func ShowServiceRoles(accessToken string, mp *MethodParams) error {

	sc, err := Get(accessToken, mp, &mashcli.Params{Fields: SERVICEROLES_ALL_FIELDS})

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Service Role ID", "Name", "Action", "Created", "Updated"})

	for _, s := range *sc {
		data := []string{s.Id, s.Name, s.Action, s.Created[:19], s.Updated[:19]}
		table.Append(data)
	}
	table.Render()

	return nil

}

func (s *ServiceRoles) PrettyPrint() {

	data := []string{s.Id, s.Name, s.Action, s.Created[:19], s.Updated[:19]}
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Service Role ID", "Name", "Action", "Created", "Updated"})
	table.Append(data)
	table.Render()

	return

}


