// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy_test

import (
	"context"
	"os"
	"testing"

	"github.com/with-ours/ingest-sdk-go"
	"github.com/with-ours/ingest-sdk-go/internal/testutil"
	"github.com/with-ours/ingest-sdk-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := oursprivacy.NewClient(
		option.WithBaseURL(baseURL),
	)
	response, err := client.Track.Event(context.TODO(), oursprivacy.TrackEventParams{
		Token:      "YOUR_SOURCE_TOKEN",
		Event:      "signup_completed",
		DistinctID: oursprivacy.String("evt_01J9ZK3Q7X"),
		EventProperties: map[string]oursprivacy.TrackEventParamsEventPropertyUnion{
			"plan": {
				OfString: oursprivacy.String("pro"),
			},
			"value": {
				OfFloat: oursprivacy.Float(49),
			},
			"currency": {
				OfString: oursprivacy.String("USD"),
			},
		},
		ExternalID: oursprivacy.String("crm_contact_1042"),
		IdentityContext: oursprivacy.TrackEventParamsIdentityContext{
			IP:        "203.0.113.10",
			UserAgent: "Mozilla/5.0",
		},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.Success)
}
