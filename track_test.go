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

func TestTrackEventWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Track.Event(context.TODO(), oursprivacy.TrackEventParams{
		Token: "x",
		Event: "x",
		DefaultProperties: oursprivacy.TrackEventParamsDefaultProperties{
			ActiveDuration:  oursprivacy.Float(0),
			AdID:            oursprivacy.String("ad_id"),
			AdsetID:         oursprivacy.String("adset_id"),
			BasisCid:        oursprivacy.String("basis_cid"),
			BrowserLanguage: oursprivacy.String("browser_language"),
			BrowserName:     oursprivacy.String("browser_name"),
			BrowserVersion:  oursprivacy.String("browser_version"),
			CampaignID:      oursprivacy.String("campaign_id"),
			Clickid:         oursprivacy.String("clickid"),
			Clid:            oursprivacy.String("clid"),
			CPUArchitecture: oursprivacy.String("cpu_architecture"),
			CurrentURL:      oursprivacy.String("current_url"),
			Dclid:           oursprivacy.String("dclid"),
			DeviceModel:     oursprivacy.String("device_model"),
			DeviceType:      oursprivacy.String("device_type"),
			DeviceVendor:    oursprivacy.String("device_vendor"),
			Duration:        oursprivacy.Float(0),
			Encoding:        oursprivacy.String("encoding"),
			EngineName:      oursprivacy.String("engine_name"),
			EngineVersion:   oursprivacy.String("engine_version"),
			Epik:            oursprivacy.String("epik"),
			Fbc:             oursprivacy.String("fbc"),
			Fbclid:          oursprivacy.String("fbclid"),
			Fbp:             oursprivacy.String("fbp"),
			Fv:              oursprivacy.Bool(true),
			GadSource:       oursprivacy.String("gad_source"),
			Gbraid:          oursprivacy.String("gbraid"),
			Gclid:           oursprivacy.String("gclid"),
			Host:            oursprivacy.String("host"),
			Iframe:          oursprivacy.Bool(true),
			ImRef:           oursprivacy.String("im_ref"),
			IP:              oursprivacy.String("ip"),
			Irclickid:       oursprivacy.String("irclickid"),
			IsBot:           oursprivacy.String("is_bot"),
			LiFatID:         oursprivacy.String("li_fat_id"),
			Msclkid:         oursprivacy.String("msclkid"),
			Ndclid:          oursprivacy.String("ndclid"),
			NewS:            oursprivacy.Bool(true),
			OsName:          oursprivacy.String("os_name"),
			OsVersion:       oursprivacy.String("os_version"),
			PageHash:        oursprivacy.Float(0),
			Pathname:        oursprivacy.String("pathname"),
			Qclid:           oursprivacy.String("qclid"),
			RdtCid:          oursprivacy.String("rdt_cid"),
			ReceivedAt:      oursprivacy.String("received_at"),
			Referrer:        oursprivacy.String("referrer"),
			ReferringDomain: oursprivacy.String("referring_domain"),
			Sacid:           oursprivacy.String("sacid"),
			Sccid:           oursprivacy.String("sccid"),
			ScreenHeight:    oursprivacy.Float(0),
			ScreenWidth:     oursprivacy.Float(0),
			SessionCount:    oursprivacy.Float(0),
			Sid:             oursprivacy.String("sid"),
			Sr:              oursprivacy.String("sr"),
			Title:           oursprivacy.String("title"),
			Ttclid:          oursprivacy.String("ttclid"),
			Twclid:          oursprivacy.String("twclid"),
			Uafvl:           oursprivacy.String("uafvl"),
			UserAgent:       oursprivacy.String("user_agent"),
			UtmCampaign:     oursprivacy.String("utm_campaign"),
			UtmContent:      oursprivacy.String("utm_content"),
			UtmMedium:       oursprivacy.String("utm_medium"),
			UtmName:         oursprivacy.String("utm_name"),
			UtmSource:       oursprivacy.String("utm_source"),
			UtmTerm:         oursprivacy.String("utm_term"),
			Version:         oursprivacy.String("version"),
			Wbraid:          oursprivacy.String("wbraid"),
			Webview:         oursprivacy.Bool(true),
		},
		DistinctID: oursprivacy.String("x"),
		Email:      oursprivacy.String("x"),
		EventProperties: map[string]string{
			"foo": "string",
		},
		ExternalID: oursprivacy.String("x"),
		Time:       oursprivacy.Float(0),
		UserID:     oursprivacy.String("x"),
		UserProperties: oursprivacy.TrackEventParamsUserProperties{
			AdID:        oursprivacy.String("ad_id"),
			AdsetID:     oursprivacy.String("adset_id"),
			BasisCid:    oursprivacy.String("basis_cid"),
			CampaignID:  oursprivacy.String("campaign_id"),
			City:        oursprivacy.String("city"),
			Clickid:     oursprivacy.String("clickid"),
			Clid:        oursprivacy.String("clid"),
			CompanyName: oursprivacy.String("company_name"),
			Consent: map[string]string{
				"foo": "string",
			},
			Country: oursprivacy.String("country"),
			CustomProperties: map[string]string{
				"foo": "string",
			},
			DateOfBirth:       oursprivacy.String("date_of_birth"),
			Dclid:             oursprivacy.String("dclid"),
			Email:             oursprivacy.String("email"),
			Epik:              oursprivacy.String("epik"),
			ExternalID:        oursprivacy.String("external_id"),
			Fbc:               oursprivacy.String("fbc"),
			Fbclid:            oursprivacy.String("fbclid"),
			Fbp:               oursprivacy.String("fbp"),
			FirstName:         oursprivacy.String("first_name"),
			GadSource:         oursprivacy.String("gad_source"),
			Gbraid:            oursprivacy.String("gbraid"),
			Gclid:             oursprivacy.String("gclid"),
			Gender:            oursprivacy.String("gender"),
			ImRef:             oursprivacy.String("im_ref"),
			IP:                oursprivacy.String("ip"),
			Irclickid:         oursprivacy.String("irclickid"),
			IsBot:             oursprivacy.String("is_bot"),
			JobTitle:          oursprivacy.String("job_title"),
			LastName:          oursprivacy.String("last_name"),
			LiFatID:           oursprivacy.String("li_fat_id"),
			Msclkid:           oursprivacy.String("msclkid"),
			Ndclid:            oursprivacy.String("ndclid"),
			PhoneNumber:       oursprivacy.String("phone_number"),
			Qclid:             oursprivacy.String("qclid"),
			RdtCid:            oursprivacy.String("rdt_cid"),
			Referrer:          oursprivacy.String("referrer"),
			ReferringDomain:   oursprivacy.String("referring_domain"),
			Sacid:             oursprivacy.String("sacid"),
			Sccid:             oursprivacy.String("sccid"),
			Sid:               oursprivacy.String("sid"),
			State:             oursprivacy.String("state"),
			Ttclid:            oursprivacy.String("ttclid"),
			Twclid:            oursprivacy.String("twclid"),
			UserAgent:         oursprivacy.String("user_agent"),
			UserAgentFullList: oursprivacy.String("user_agent_full_list"),
			UtmCampaign:       oursprivacy.String("utm_campaign"),
			UtmContent:        oursprivacy.String("utm_content"),
			UtmMedium:         oursprivacy.String("utm_medium"),
			UtmName:           oursprivacy.String("utm_name"),
			UtmSource:         oursprivacy.String("utm_source"),
			UtmTerm:           oursprivacy.String("utm_term"),
			Wbraid:            oursprivacy.String("wbraid"),
			Zip:               oursprivacy.String("zip"),
		},
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
