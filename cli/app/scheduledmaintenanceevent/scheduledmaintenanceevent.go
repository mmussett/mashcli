package scheduledmaintenanceevent

import (
	"fmt"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/olekukonko/tablewriter"
	"os"
)

func ShowScheduledMaintenanceEvent(accessToken string, mp *MethodParams) error {

	m, err := Get(accessToken, mp, &mashcli.Params{Fields: SCHEDULEDMAINTENANCEEVENT_ALL_FIELDS})

	if err != nil {
		return err
	}

	m.PrettyPrint(mp)

	return nil

}

func (m *ScheduledMaintenanceEvent) PrettyPrint(mp *MethodParams) {

	caption := fmt.Sprintf("Service ID: %s - Endpoint ID: %s \n", mp.ServiceId, mp.EndpointId)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Scheduled Maintenance Event ID", "Name", "Start Date Time", "End Date Time"})
	table.SetCaption(true, caption)
	data := []string{m.Id, m.Name, m.StartDateTime[:19], m.EndDateTime[:19]}
	table.Append(data)
	table.Render()
	return

}
