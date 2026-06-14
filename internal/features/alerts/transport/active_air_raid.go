package alerts_transport

import (
	"net/http"

	alerts_service "alerts.api/internal/features/alerts/service"
)

func AlertsHandler(w http.ResponseWriter, r *http.Request) {
	ServeCachedResponse(w, r, alerts_service.ActiveAlertsKey, "alerts")
}
