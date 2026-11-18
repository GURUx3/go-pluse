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

	u, err := url.ParseRequestURI(resource)
	if err != nil {
		return MonitorResult{Error: err}
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

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	start := time.Now()

	res, err := client.Get(resource)
	elapsed := time.Since(start)

	if err != nil {
		return MonitorResult{
			Latency: elapsed,
			Error:   err,
		}
	}

	defer res.Body.Close()

	healthy := res.StatusCode >= 200 && res.StatusCode < 300

	return MonitorResult{
		StatusCode: res.StatusCode,
		Latency:    elapsed,
		Health:     healthy,
		Error:      nil,
	}
}
