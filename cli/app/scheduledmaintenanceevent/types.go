package scheduledmaintenanceevent

type ScheduledMaintenanceEvent struct {
	Id            string   `json:"id"`
	Name          string   `json:"name,omitempty"`
	StartDateTime string   `json:"startDateTime,omitempty"`
	EndDateTime   string   `json:"endDateTime,omitempty"`
	Endpoints     []string `json:"endpoints,omitempty"`
}

type DeleteScheduledMaintenanceEventResponse struct {
	Status string `json:"status"`
}

type MethodParams struct {
	ServiceId  string
	EndpointId string
}
