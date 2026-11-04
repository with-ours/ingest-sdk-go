// Backfill historical events through /batch and resend only the rows that failed
// to queue.
//
//	OURS_PRIVACY_TOKEN=<source token> go run ./examples/batch-backfill
//
// Set OURS_PRIVACY_BASE_URL to send to a host other than production. All values
// are samples.
//
// The supported backfill window is 1 year. An older time is moved to the edge of
// the window, not rejected. The sample rows below are a few days old.
//
// A partial failure is HTTP 207, which the SDK returns as a response and not as
// an error, so check every row. Resend queue_failed rows unchanged. Correct
// invalid_event rows before sending them again.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	oursprivacy "github.com/with-ours/ingest-sdk-go"
)

const maxAttempts = 3

func main() {
	token := os.Getenv("OURS_PRIVACY_TOKEN")
	if token == "" {
		log.Fatal("OURS_PRIVACY_TOKEN is required: set it to your Source token")
	}
	client := oursprivacy.NewClient()

	// DistinctID is required on batch rows. Keep it stable so a resend deduplicates.
	occurredAt := time.Now().AddDate(0, 0, -3)
	var rows []oursprivacy.BatchNewParamsEvent
	for n := range 3 {
		rows = append(rows, oursprivacy.BatchNewParamsEvent{
			Event:      "Example Import", // must be an allowed event on your Source
			DistinctID: fmt.Sprintf("example-import-%d", n),
			ExternalID: oursprivacy.String("example-user-123"),
			Time:       oursprivacy.Float(float64(occurredAt.Add(time.Duration(n) * time.Minute).UnixMilli())),
		})
	}

	invalid := 0
	for attempt := 1; attempt <= maxAttempts && len(rows) > 0; attempt++ {
		if attempt > 1 {
			time.Sleep(time.Duration(attempt) * time.Second)
		}
		response, err := client.Batch.New(context.Background(), oursprivacy.BatchNewParams{Token: token, Events: rows})
		if err != nil {
			log.Fatal(err) // the whole request failed, after the SDK's own retries
		}
		fmt.Printf("attempt %d: accepted=%d failed=%d\n", attempt, response.Accepted, response.Failed)

		// Each result's Index points into the rows sent in this attempt.
		var retry []oursprivacy.BatchNewParamsEvent
		for _, result := range response.Results {
			switch {
			case result.Success:
			case result.Code == "queue_failed":
				retry = append(retry, rows[result.Index])
			default:
				invalid++
				log.Printf("%s needs correcting: %s: %s", rows[result.Index].DistinctID, result.Code, result.Message)
			}
		}
		rows = retry
	}
	if len(rows) > 0 || invalid > 0 {
		log.Fatalf("backfill incomplete: %d rows failed to queue, %d rows need correcting", len(rows), invalid)
	}
	fmt.Println("backfill complete")
}
