// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/with-ours/ingest-sdk-go"
	"github.com/with-ours/ingest-sdk-go/internal/testutil"
	"github.com/with-ours/ingest-sdk-go/option"
)

func TestExperimentAssignmentWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Experiments.Assignment(
		context.TODO(),
		"experiment_key",
		oursprivacy.ExperimentAssignmentParams{
			Token:     "token",
			VisitorID: "x",
			Context: oursprivacy.ExperimentAssignmentParamsContext{
				Search: oursprivacy.String("search"),
				URL:    oursprivacy.String("url"),
			},
			TrackImpression: oursprivacy.Bool(true),
		},
	)
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentPersonalization(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Experiments.Personalization(context.TODO(), oursprivacy.ExperimentPersonalizationParams{
		Token:     "token",
		VisitorID: "x",
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
