package main

import (
	"fmt"
	"os"

	"github.com/ona-se/onaflix-data-pipeline/internal/processor"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://gitpod:gitpod@localhost:5432/onaflix?sslmode=disable"
	}

	p, err := processor.New(dbURL)
	if err != nil {
		log.Fatalf("Failed to create processor: %v", err)
	}
	defer p.Close()

	log.Info("Starting OnaFlix data pipeline")

	stats, err := p.Run()
	if err != nil {
		log.Fatalf("Pipeline failed: %v", err)
	}

	fmt.Printf("Pipeline complete: %d movies processed, %d enriched, %d errors\n",
		stats.Processed, stats.Enriched, stats.Errors)
}
