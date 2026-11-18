package main

import (
	"fmt"

	"github.com/GURUx3/pulse/internal/monitor"
)

func main() {
	target := "https://httpbin.org/status/200"

	result := monitor.Check(target)

	fmt.Println("\nPULSE MONITOR")
	fmt.Println("--------------------")

	if result.Error != nil {
		fmt.Println("Health: UNHEALTHY")
		fmt.Println("Error:", result.Error)
		return
	}

	fmt.Println("Status:", result.StatusCode)
	fmt.Printf("Latency: %dms\n", result.Latency.Milliseconds())

	if result.Health {
		fmt.Println("Health: HEALTHY")
	} else {
		fmt.Println("Health: UNHEALTHY")
	}
}
