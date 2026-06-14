package alerts_client

import (
	"encoding/json"
)

type ActiveAlertsResponse struct {
	Alerts []Alert `json:"alerts"`
}

type Alert struct {
	OblastUID int    `json:"location_oblast_uid"`
	AlertType string `json:"alert_type"`
}

func GetActiveRegions() ([]byte, error) {
	data, err := getAlertsAPI("/v1/alerts/active.json")
	if err != nil {
		return nil, err
	}

	var alertsResponse ActiveAlertsResponse
	if err := json.Unmarshal(data, &alertsResponse); err != nil {
		return nil, err
	}

	return CreateResponse(alertsResponse.Alerts)
}
