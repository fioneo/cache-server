package alerts_client

import (
	"encoding/json"
	"fmt"
)

var alertTypeBits = map[string]uint8{
	"air_raid":           1,
	"artillery_shelling": 2,
	"urban_fights":       4,
	"chemical":           8,
	"nuclear":            16,
}

func CreateResponse(alerts []Alert) ([]byte, error) {
	encoded, err := EncodeActiveAlerts(alerts)
	if err != nil {
		return nil, err
	}

	return json.Marshal(encoded)
}

func EncodeActiveAlerts(alerts []Alert) (map[int]uint8, error) {
	masks := make(map[int]uint8, len(alerts))

	for _, alert := range alerts {
		bit, err := AlertTypeBit(alert.AlertType)
		if err != nil {
			return nil, err
		}

		masks[alert.OblastUID] |= bit
	}

	return masks, nil
}

func AlertTypeBit(alertType string) (uint8, error) {
	bit, ok := alertTypeBits[alertType]
	if !ok {
		return 0, fmt.Errorf("unknown alert type: %s", alertType)
	}

	return bit, nil
}
