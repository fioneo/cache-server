package alerts_service

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	core_cache "alerts.api/internal/core/cache"
	alerts_client "alerts.api/internal/features/alerts/client"
)

const (
	ActiveAlertsKey     = "active_alerts"
	RegionAlertTypesKey = "region_alert_types"
)

var Cache = core_cache.NewCache()

func RefreshData(key string) error {
	var freshData []byte
	var err error

	switch key {
	case ActiveAlertsKey:
		freshData, err = alerts_client.GetActiveAlerts()
	case RegionAlertTypesKey:
		freshData, err = alerts_client.GetActiveRegions()
	default:
		err = fmt.Errorf("unknown refresh request: %s", key)
	}
	if err != nil {
		return err
	}

	Cache.Set(key, freshData)
	slog.Debug("data refreshed successfully", "key", key)
	return nil
}

func StartUpdater(ctx context.Context, interval time.Duration, key string) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := RefreshData(key); err != nil {
					slog.Error("refresh data failed", "key", key, "error", err)
				}
			}
		}
	}()
}
