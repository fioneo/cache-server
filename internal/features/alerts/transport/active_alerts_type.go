package alerts_transport

import (
	"net/http"

	alerts_service "alerts.api/internal/features/alerts/service"
)

func AlertsTypeHandler(w http.ResponseWriter, r *http.Request) {
	ServeCachedResponse(w, r, alerts_service.RegionAlertTypesKey, "region alert types")
}
