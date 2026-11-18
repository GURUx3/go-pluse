package main

import (
	"fmt"

	"github.com/GURUx3/pulse/internal/monitor"
)

func main() {
	// hii
	// target := "https://httpbin.org/status/200"

	targets := []string{
		"https://api.github.com",
		"https://google.com",
		"https://httpbin.org/status/500",
	}

	for _, target := range targets {
		result := monitor.Check(target)

		fmt.Println("\nPULSE MONITOR")
		fmt.Println("--------------------")
		fmt.Println("URL:", target)

		if result.Error != nil {
			fmt.Println("Health: UNHEALTHY")
			fmt.Println("Error:", result.Error)
			continue
		}

		fmt.Println("Status:", result.StatusCode)
		fmt.Printf("Latency: %dms\n", result.Latency.Milliseconds())

		if result.Health {
			fmt.Println("Health: HEALTHY")
		} else {
			fmt.Println("Health: UNHEALTHY")
		}
	}

}
