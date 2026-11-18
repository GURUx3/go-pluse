package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	client := &http.Client{
		Timeout: 5 * time.Second, // Total timeout for the request
	}

	startTime := time.Now()

	// fmt.Println(startTime)
	res, err := client.Get("https://api.github.com/zen/1")

	if err != nil {

		fmt.Println("Timeout error", err)
		return
	}

	defer res.Body.Close()

	// fmt.Println("Header :", res.Header)
	// body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("error", err)
		return
	}
	elapsed := time.Since(startTime)
	status := res.Status

	hel := "UNHEALTHY"

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		hel = "HEALTHY"
	}

	fmt.Println("\n\n")
	fmt.Println("PLUSE MONITOR")
	fmt.Println("\n")
	fmt.Println("Status : ", status)
	fmt.Printf("Latency: %dms\n", elapsed.Milliseconds())
	fmt.Println("Health:", hel)
	// fmt.Println("Body :\n")
	// fmt.Println(string(body))
}
