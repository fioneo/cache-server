package main

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"
	"time"

	core_logger "alerts.api/internal/core/logger"
	core_server "alerts.api/internal/core/server"
	alerts_service "alerts.api/internal/features/alerts/service"
	alerts_transport "alerts.api/internal/features/alerts/transport"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Warn(".env file not loaded", "error", err)
	}
	core_logger.Init()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := alerts_service.RefreshData(alerts_service.ActiveAlertsKey); err != nil {
		slog.Error("initial alerts refresh failed", "error", err)
	}
	if err := alerts_service.RefreshData(alerts_service.RegionAlertTypesKey); err != nil {
		slog.Error("initial regions type refresh failed", "error", err)
	}

	alerts_service.StartUpdater(ctx, 15*time.Second, alerts_service.ActiveAlertsKey)
	alerts_service.StartUpdater(ctx, 30*time.Second, alerts_service.RegionAlertTypesKey)

	server := core_server.NewHTTPServer()
	httpHandler := alerts_transport.NewHTTPHandler()
	apiRouter := core_server.NewApiRouter()
	apiRouter.RegisterRoutes(httpHandler.Routes()...)
	server.RegisterRoutes(apiRouter)

	if err := server.Start(ctx); err != nil {
		slog.Error("server stopped with error", "error", err)
	}
}
