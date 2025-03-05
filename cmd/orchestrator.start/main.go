package main

import (
	"log"

	"calc_service/internal/orchestrator"
)

func main() {
	app := orchestrator.NewOrchestrator()
	log.Println("Starting Orchestrator on port", app.Config.Addr)
	if err := app.RunServer(); err != nil {
		log.Fatal(err)
	}
}
