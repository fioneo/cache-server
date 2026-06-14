package alerts_client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

const alertsAPIBaseURL = "https://api.alerts.in.ua"

func getAlertsAPI(path string) ([]byte, error) {
	token := os.Getenv("TOKEN")
	if token == "" {
		return nil, fmt.Errorf("TOKEN env variable is empty")
	}

	requestURL := fmt.Sprintf("%s%s?token=%s", alertsAPIBaseURL, path, url.QueryEscape(token))
	resp, err := client.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("unexpected upstream status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
