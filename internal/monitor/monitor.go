package monitor

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type MonitorResult struct {
	StatusCode int
	Latency    time.Duration
	Health     bool
	Error      error
}

func Check(resource string) MonitorResult {

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	start := time.Now()

	u, err := url.ParseRequestURI(resource)

	if err != nil {

		return MonitorResult{
			Error: err,
		}

	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return MonitorResult{
			Error: fmt.Errorf("unsupported URL scheme: %s", u.Scheme),
		}
	}

	if u.Hostname() == "" {
		return MonitorResult{
			Error: fmt.Errorf("URL hostname is missing"),
		}
	}

	res, err := client.Get(resource)
	status := true

	if err != nil {
		status = false
		return MonitorResult{
			Error: err,
		}
	}

	defer res.Body.Close()

	elapsed := time.Since(start)

	return MonitorResult{

		StatusCode: res.StatusCode,
		Latency:    elapsed,
		Health:     status,
		Error:      nil,
	}

}
