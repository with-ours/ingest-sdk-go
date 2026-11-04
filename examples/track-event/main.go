// Record one server-side event for a known user, with event properties.
//
//	OURS_PRIVACY_TOKEN=<source token> go run ./examples/track-event
//
// Set OURS_PRIVACY_BASE_URL to send to a host other than production. All values
// are samples. Never put credentials or health data in event properties.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	oursprivacy "github.com/with-ours/ingest-sdk-go"
)

func main() {
	token := os.Getenv("OURS_PRIVACY_TOKEN")
	if token == "" {
		log.Fatal("OURS_PRIVACY_TOKEN is required: set it to your Source token")
	}

	client := oursprivacy.NewClient()
	response, err := client.Track.Event(context.Background(), oursprivacy.TrackEventParams{
		Token:      token,
		Event:      "Example Signup", // must be an allowed event on your Source
		ExternalID: oursprivacy.String("example-user-123"),
		// Reuse the same DistinctID when you resend an event so it deduplicates.
		DistinctID: oursprivacy.String(fmt.Sprintf("example-signup-%d", time.Now().UnixNano())),
		EventProperties: map[string]oursprivacy.TrackEventParamsEventPropertyUnion{
			"plan": {OfString: oursprivacy.String("example-plan")},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tracked:", response.Success)
}
