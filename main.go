package main

import (
	"log"
	"open-telemetry-test/app"
)

func main() {
	if err := app.New(); err != nil {
		log.Fatal("failed to start application %w", err)
	}
}
