package metadata

import "context"

// MaintenanceEventData describes maintenance affecting the current instance.
type MaintenanceEventData struct {
	MaintenanceID             int    `json:"maintenance_id"`
	AppliedMaintenancePolicy  string `json:"applied_maintenance_policy"`
	SelectedMaintenancePolicy string `json:"selected_maintenance_policy"`
	ScheduleType              string `json:"schedule_type"`
	Description               string `json:"description"`
	Source                    string `json:"source"`
	NotBefore                 string `json:"not_before"`
	EventStatus               string `json:"event_status"`
}

// MaintenanceEventsData is a page of maintenance events and pagination metadata.
type MaintenanceEventsData struct {
	Data    []MaintenanceEventData `json:"data"`
	Page    int                    `json:"page"`
	Pages   int                    `json:"pages"`
	Results int                    `json:"results"`
}

// GetMaintenanceEvents returns maintenance events for the current instance.
func (c *Client) GetMaintenanceEvents(ctx context.Context) (*MaintenanceEventsData, error) {
	req := c.R(ctx).SetResult(&MaintenanceEventsData{})
	resp, err := coupleAPIErrors(req.Get("maintenance/events"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*MaintenanceEventsData), nil
}
