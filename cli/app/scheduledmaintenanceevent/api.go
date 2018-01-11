package scheduledmaintenanceevent

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/sling"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/pkg/errors"
	"io"
	"strconv"
	"strings"
)

const (
	resourcePath = "v3/rest/services/%s/endpoints/%s/scheduledMaintenanceEvent"
)

func Get(accessToken string, mp *MethodParams, params *mashcli.Params) (*ScheduledMaintenanceEvent, error) {

	e := new(mashcli.MasheryError)
	s := new(ScheduledMaintenanceEvent)

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)

	_, err := sling.New().Base(mashcli.BaseURL).Path(path).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").QueryStruct(params).Receive(s, e)

	if err != nil {
		return nil, err
	}

	if e.ErrorCode == 404 {
		return nil, errors.Errorf("scheduledmaintenanceevent: unable to get scheduled maintenance event : GET %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return s, nil
}

func (m *ScheduledMaintenanceEvent) Create(accessToken string, mp *MethodParams) (*ScheduledMaintenanceEvent, error) {

	e := new(mashcli.MasheryError)

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Set("Content-Type", "application/json").Set("Accept", "application/json").BodyJSON(m).Post(path).Receive(m, e)

	if err != nil {
		return nil, errors.Errorf("scheduledmaintenanceevent: unable to create scheduled maintenance event: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	if resp.StatusCode != 200 {
		return nil, errors.Errorf("scheduledmaintenanceevent: unable to create scheduled maintenance event: POST %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func (m *ScheduledMaintenanceEvent) Update(accessToken string, mp *MethodParams) (*ScheduledMaintenanceEvent, error) {

	e := new(mashcli.MasheryError)

	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, errors.Errorf("methods: unable to marshall contents")
	}

	body := strings.NewReader(string(bytes))

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Body(body).Set("Content-Type", "application/json").Put(path).Receive(m, e)

	if err != nil {
		return nil, errors.Errorf("scheduledmaintenanceevent: unable to update scheduled maintenance event: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	if resp.StatusCode != 200 {
		return nil, errors.Errorf("scheduledmaintenanceevent: unable to update scheduled maintenance event: PUT %s -> (%s %s)", path, strconv.Itoa(e.ErrorCode), e.ErrorMessage)
	}

	return m, nil

}

func Delete(accessToken string, mp *MethodParams) error {

	r := new(DeleteScheduledMaintenanceEventResponse)

	path := fmt.Sprintf(resourcePath, mp.ServiceId, mp.EndpointId)

	resp, err := sling.New().Base(mashcli.BaseURL).Set("Authorization", "Bearer "+accessToken).Delete(path).ReceiveSuccess(r)

	if resp.StatusCode == 200 || resp.StatusCode == 404 {
		return nil
	}

	if err != io.EOF {
		return errors.Errorf("scheduledmaintenanceevent: unable to delete scheduled maintenance event: DELETE %s", path, mp.ServiceId)
	}

	if resp.StatusCode != 200 || resp.StatusCode != 404 {
		return errors.Errorf("scheduledmaintenanceevent: unable to delete rscheduled maintenance event: DELETE %s -> (%s %s)", path, resp.StatusCode, resp.Status)
	}

	return nil

}
