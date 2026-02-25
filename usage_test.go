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
	t.Skip("Mock server tests are disabled")
	response, err := client.Track.Event(context.TODO(), oursprivacy.TrackEventParams{
		Token: "REPLACE_ME",
		Event: "REPLACE_ME",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.Success)
}
