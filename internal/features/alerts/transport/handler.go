package alerts_transport

import (
	"log/slog"
	"net/http"

	core_server "alerts.api/internal/core/server"
	alerts_service "alerts.api/internal/features/alerts/service"
)

type HTTPHandler struct{}

func NewHTTPHandler() *HTTPHandler {
	return &HTTPHandler{}
}

func (h *HTTPHandler) Routes() []core_server.Route {
	return []core_server.Route{
		{
			Path:    "/alerts",
			Method:  "GET",
			Handler: AlertsHandler,
		},
		{
			Path:    "/active",
			Method:  "GET",
			Handler: AlertsTypeHandler,
		},
	}
}

func ServeCachedResponse(w http.ResponseWriter, r *http.Request, key, name string) {
	data, etag, ok := alerts_service.Cache.Get(key)
	if !ok {
		http.Error(w, name+" data is not loaded yet", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("ETag", etag)

	if r.Header.Get("If-None-Match") == etag {
		w.WriteHeader(http.StatusNotModified)
		slog.Debug("response served from cache", "name", name, "status", http.StatusNotModified)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		slog.Error("write cached response", "name", name, "error", err)
		return
	}
	slog.Debug("response served", "name", name, "status", http.StatusOK)
}
