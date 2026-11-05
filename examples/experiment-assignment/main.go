// Choose an experiment variant on the server and fail open: if the assignment
// errors or times out, serve the default experience and carry on.
//
//	OURS_PRIVACY_EXPERIMENT_TOKEN=exp_... go run ./examples/experiment-assignment
//
// Experiment calls need the experiment token (exp_*). They don't accept a Source
// token. Set OURS_PRIVACY_BASE_URL to send to a host other than production. All
// values are samples.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	oursprivacy "github.com/with-ours/ingest-sdk-go"
)

const defaultVariant = "default"

func variantFor(client oursprivacy.Client, token, visitorID string) string {
	// Bound the call so a slow assignment can't hold up the page.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	assignment, err := client.Experiments.Assignment(ctx, "example_experiment", oursprivacy.ExperimentAssignmentParams{
		Token:     token,
		VisitorID: visitorID,
		Context: oursprivacy.ExperimentAssignmentParamsContext{
			URL: oursprivacy.String("https://example.com/pricing"),
		},
	})
	if err != nil {
		log.Printf("assignment unavailable, serving %q: %v", defaultVariant, err)
		return defaultVariant
	}
	if assignment.InExperiment && assignment.VariantName != "" {
		return assignment.VariantName
	}
	return defaultVariant
}

func main() {
	token := os.Getenv("OURS_PRIVACY_EXPERIMENT_TOKEN")
	if token == "" {
		log.Fatal("OURS_PRIVACY_EXPERIMENT_TOKEN is required: set it to your experiment token (exp_*)")
	}
	fmt.Println("variant:", variantFor(oursprivacy.NewClient(), token, "example-user-123"))
}
