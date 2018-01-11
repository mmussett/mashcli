package systemdomainauthentication

import (
	"fmt"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"
	"os"
)

func ShowSystemDomainAuthentication(accessToken string, mp *MethodParams) error {

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: SYSTEMDOMAINAUTHENTICATION_ALL_FIELDS})

	if err != nil {
		return err
	}

	m.PrettyPrint(mp)

	return nil

}

func (m *SystemDomainAuthentication) PrettyPrint(mp *MethodParams) {

	caption := fmt.Sprintf("Service ID: %s - Endpoint ID: %s \n", mp.ServiceId, mp.EndpointId)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Type", "Username", "Certificate", "Password"})
	table.SetCaption(true, caption)
	data := []string{m.Type, m.Username, m.Certificate, m.Password}
	table.Append(data)
	table.Render()
	return

}
