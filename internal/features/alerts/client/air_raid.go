package alerts_client

func GetActiveAlerts() ([]byte, error) {
	return getAlertsAPI("/v1/iot/active_air_raid_alerts.json")
}
