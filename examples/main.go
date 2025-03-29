package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	flexprice "github.com/flexprice/go-sdk"
	"github.com/joho/godotenv"
	"github.com/samber/lo"
)

// This sample application demonstrates how to use the FlexPrice Go SDK
// to create and retrieve events, showing the basic patterns for API interaction.
// To run this example:
// 1. Copy this file to your project
// 2. Create a .env file with FLEXPRICE_API_KEY and FLEXPRICE_API_HOST
// 3. Run with: go run main.go

// Sample .env file:
// FLEXPRICE_API_KEY=your_api_key_here
// FLEXPRICE_API_HOST=api.flexprice.io

func RunSample() {
	// Load .env file if present
	godotenv.Load()

	// Get API credentials from environment
	apiKey := os.Getenv("FLEXPRICE_API_KEY")
	apiHost := os.Getenv("FLEXPRICE_API_HOST")

	if apiKey == "" || apiHost == "" {
		log.Fatal("Missing required environment variables: FLEXPRICE_API_KEY and FLEXPRICE_API_HOST")
	}

	// Initialize API client
	config := flexprice.NewConfiguration()
	config.Scheme = "https"
	config.Host = apiHost
	config.AddDefaultHeader("x-api-key", apiKey)

	client := flexprice.NewAPIClient(config)
	ctx := context.Background()

	// Generate a unique customer ID for this sample
	customerId := fmt.Sprintf("sample-customer-%d", time.Now().Unix())

	// Step 1: Create an event
	fmt.Println("Creating event...")
	eventRequest := flexprice.DtoIngestEventRequest{
		EventName:          "Sample Event",
		ExternalCustomerId: customerId,
		Properties: &map[string]string{
			"source":      "sample_app",
			"environment": "test",
			"timestamp":   time.Now().String(),
		},
		Source:    lo.ToPtr("sample_app"),
		Timestamp: lo.ToPtr(time.Now().Format(time.RFC3339)),
	}

	// Send the event creation request
	result, response, err := client.EventsAPI.EventsPost(ctx).
		Event(eventRequest).
		Execute()

	if err != nil {
		log.Fatalf("Error creating event: %v", err)
	}

	if response.StatusCode != 202 {
		log.Fatalf("Expected status code 202, got %d", response.StatusCode)
	}

	// The result is a map, so we need to use map access
	eventId := result["event_id"]
	fmt.Printf("Event created successfully! ID: %v\n\n", eventId)

	// Step 2: Retrieve events for this customer
	fmt.Println("Retrieving events for customer...")
	events, response, err := client.EventsAPI.EventsGet(ctx).
		ExternalCustomerId(customerId).
		Execute()

	if err != nil {
		log.Fatalf("Error retrieving events: %v", err)
	}

	if response.StatusCode != 200 {
		log.Fatalf("Expected status code 200, got %d", response.StatusCode)
	}

	// Process the events (the response is a map)
	fmt.Printf("Raw response: %+v\n\n", response)

	for i, event := range events.Events {
		fmt.Printf("Event %d: %v - %v\n", i+1, event.Id, event.EventName)
		fmt.Printf("Event properties: %v\n", event.Properties)
	}

	fmt.Println("Sample application completed successfully!")
}

// For direct execution of this file
func main() {
	RunSample()
}
