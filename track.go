// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"net/http"
	"slices"

	"github.com/with-ours/ingest-sdk-go/internal/apijson"
	"github.com/with-ours/ingest-sdk-go/internal/requestconfig"
	"github.com/with-ours/ingest-sdk-go/option"
	"github.com/with-ours/ingest-sdk-go/packages/param"
	"github.com/with-ours/ingest-sdk-go/packages/respjson"
)

// TrackService contains methods and other services that help with interacting with
// the ours-privacy API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTrackService] method instead.
type TrackService struct {
	Options []option.RequestOption
}

// NewTrackService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTrackService(opts ...option.RequestOption) (r TrackService) {
	r = TrackService{}
	r.Options = opts
	return
}

// Track events from your server. Include userId, externalId, email, or
// userProperties.phone_number to associate the event with an existing visitor.
// Identity resolution runs in priority order: userId (direct, no lookup) →
// externalId → email → userProperties.phone_number. Each lookup runs only if
// earlier identifiers did not resolve a visitor. Phone matching ignores common
// formatting and an international + or 00 prefix, but does not infer a country
// code. If you know both userId and externalId, send both. For top-level visitor
// properties: null clears the existing value, while undefined, omitted fields, and
// empty strings are ignored. For entries inside custom_properties: null,
// undefined, and empty strings are all ignored (custom_properties use merge
// semantics). See https://docs.oursprivacy.com/docs/data-types for details and
// common pitfalls.
func (r *TrackService) Event(ctx context.Context, body TrackEventParams, opts ...option.RequestOption) (res *TrackEventResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.oursprivacy.com/api/v1/")}, opts...)
	path := "track"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type TrackEventResponse struct {
	// Any of true.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackEventResponse) RawJSON() string { return r.JSON.raw }
func (r *TrackEventResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrackEventParams struct {
	// The token for your Source. You can find this in the dashboard. Leading and
	// trailing whitespace is removed before validation.
	Token string `json:"token" api:"required"`
	// The name of the event you're tracking. This must be whitelisted in the Ours
	// dashboard. Leading and trailing whitespace is removed before validation.
	Event string `json:"event" api:"required"`
	// A unique identifier for this event used for deduplication. Highly recommended —
	// if omitted, Ours will generate one for you, but supplying your own gives you
	// stronger idempotency guarantees (e.g. a Stripe payment intent ID or your
	// internal order ID). Non-string input is converted to text before validation;
	// clients should send a string. Leading and trailing whitespace is removed before
	// validation.
	DistinctID param.Opt[string] `json:"distinctId,omitzero"`
	// The email address of a user. When userId is absent and externalId does not
	// resolve a visitor, we search your account for a visitor with this email. If no
	// match is found, we try userProperties.phone_number before creating a new
	// visitor. Leading and trailing whitespace is removed before validation.
	Email param.Opt[string] `json:"email,omitzero"`
	// Your system's unique identifier for this user. When userId is absent, we search
	// your account for an existing visitor with this externalId. If no match is found,
	// we try email and then userProperties.phone_number before creating a new visitor.
	// If you also have the userId from cookies or local storage, send both — it
	// removes the lookup round-trip. Non-string input is converted to text before
	// validation; clients should send a string. Leading and trailing whitespace is
	// removed before validation.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// The time at which the event occurred in milliseconds since UTC epoch. The time
	// must be in the past and within the last 7 days.
	Time param.Opt[float64] `json:"time,omitzero"`
	// The Ours Visitor ID stored in local storage and cookies on your web properties.
	// When present, this is used directly — no lookup by externalId, email, or phone
	// is performed. If you have both a userId and an externalId, send both so the
	// event is attached to the right visitor without any lookup overhead. Leading and
	// trailing whitespace is removed before validation.
	UserID param.Opt[string] `json:"userId,omitzero"`
	// These properties are used throughout the Ours app to pass known values onto
	// destinations
	DefaultProperties TrackEventParamsDefaultProperties `json:"defaultProperties,omitzero"`
	// Any additional event properties you want to pass along. Use at most 200 nonempty
	// property keys and five object levels. Array contents do not count toward the
	// object-depth limit. Empty keys are removed recursively, arrays are truncated to
	// 250 entries, and top-level numbers outside the safe integer range become
	// strings. After normalization, the serialized JSON must be shorter than 1,048,576
	// UTF-16 code units.
	EventProperties map[string]TrackEventParamsEventPropertyUnion `json:"eventProperties,omitzero"`
	// End-user network context for server-side calls. Required for probabilistic
	// identity resolution when the caller is a backend server rather than an end-user
	// browser.
	IdentityContext TrackEventParamsIdentityContext `json:"identityContext,omitzero"`
	// Properties to set on the visitor. (optional) You can also update these
	// properties via the identify endpoint.
	UserProperties TrackEventParamsUserProperties `json:"userProperties,omitzero"`
	paramObj
}

func (r TrackEventParams) MarshalJSON() (data []byte, err error) {
	type shadow TrackEventParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TrackEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are used throughout the Ours app to pass known values onto
// destinations
type TrackEventParamsDefaultProperties struct {
	// The Everflow affiliate Click (Transaction) ID, captured from the
	// `_ef_transaction_id` URL parameter. Ex: ef_click_abc123 Leading and trailing
	// whitespace is removed before validation.
	EfTransactionID param.Opt[string] `json:"_ef_transaction_id,omitzero"`
	// The total time in milliseconds this tab has been visible since the page loaded
	ActiveDuration param.Opt[float64] `json:"activeDuration,omitzero"`
	// The ad id detected in the session. This is set by the web sdk automatically.
	// Leading and trailing whitespace is removed before validation.
	AdID param.Opt[string] `json:"ad_id,omitzero"`
	// The Admitad (Mitgo) affiliate Click ID. Ex: admitad_uid_abc123 Leading and
	// trailing whitespace is removed before validation.
	AdmitadUid param.Opt[string] `json:"admitad_uid,omitzero"`
	// The adset id detected in the session. This is set by the web sdk automatically.
	// Leading and trailing whitespace is removed before validation.
	AdsetID param.Opt[string] `json:"adset_id,omitzero"`
	// The AppLovin alart query parameter. Ex: alart123 Leading and trailing whitespace
	// is removed before validation.
	Alart param.Opt[string] `json:"alart,omitzero"`
	// The AppLovin aleid query parameter. Ex: aleid123 Leading and trailing whitespace
	// is removed before validation.
	Aleid param.Opt[string] `json:"aleid,omitzero"`
	// The AppLovin pixel cookie value (\_axwrt). Web-only. Leading and trailing
	// whitespace is removed before validation.
	Axwrt param.Opt[string] `json:"axwrt,omitzero"`
	// The Basis DSP Click ID. Ex: basis_cid123 Leading and trailing whitespace is
	// removed before validation.
	BasisCid param.Opt[string] `json:"basis_cid,omitzero"`
	// The Beeswax (FreeWheel Buyer Cloud) auction ID, captured from the
	// `{{AUCTION_ID}}` macro on creative click URLs. Ex: bx-auc-abc123 Leading and
	// trailing whitespace is removed before validation.
	BeeswaxAuctionID param.Opt[string] `json:"beeswax_auction_id,omitzero"`
	// The language of the browser, lowercased. Ex: en-us Leading and trailing
	// whitespace is removed before validation.
	BrowserLanguage param.Opt[string] `json:"browser_language,omitzero"`
	// The name of the browser. Ex: Chrome Leading and trailing whitespace is removed
	// before validation.
	BrowserName param.Opt[string] `json:"browser_name,omitzero"`
	// The version of the browser. Ex: 114.0 Leading and trailing whitespace is removed
	// before validation.
	BrowserVersion param.Opt[string] `json:"browser_version,omitzero"`
	// The campaign id detected in the session. This is set by the web sdk
	// automatically. Leading and trailing whitespace is removed before validation.
	CampaignID param.Opt[string] `json:"campaign_id,omitzero"`
	// The Click ID. Ex: clickid123 Leading and trailing whitespace is removed before
	// validation.
	Clickid param.Opt[string] `json:"clickid,omitzero"`
	// The Generic Click ID. Ex: clid123 Leading and trailing whitespace is removed
	// before validation.
	Clid param.Opt[string] `json:"clid,omitzero"`
	// The architecture of the CPU. Ex: amd64 Leading and trailing whitespace is
	// removed before validation.
	CPUArchitecture param.Opt[string] `json:"cpu_architecture,omitzero"`
	// The full url (including query params) of the current page Leading and trailing
	// whitespace is removed before validation.
	CurrentURL param.Opt[string] `json:"current_url,omitzero"`
	// The DoubleClick Click ID. Ex: dclid123 Leading and trailing whitespace is
	// removed before validation.
	Dclid param.Opt[string] `json:"dclid,omitzero"`
	// The model of the device. Ex: iPhone Leading and trailing whitespace is removed
	// before validation.
	DeviceModel param.Opt[string] `json:"device_model,omitzero"`
	// The type of device the user is using. Ex: mobile Leading and trailing whitespace
	// is removed before validation.
	DeviceType param.Opt[string] `json:"device_type,omitzero"`
	// The vendor of the device. Ex: Apple Leading and trailing whitespace is removed
	// before validation.
	DeviceVendor param.Opt[string] `json:"device_vendor,omitzero"`
	// The time in milliseconds since the previous event on this page, or since the web
	// SDK loaded for the first event. Restarts when a hidden tab becomes visible
	// again.
	Duration param.Opt[float64] `json:"duration,omitzero"`
	// The browser's encoding. Ex: UTF-8 Leading and trailing whitespace is removed
	// before validation.
	Encoding param.Opt[string] `json:"encoding,omitzero"`
	// The name of the browser engine. Ex: Blink Leading and trailing whitespace is
	// removed before validation.
	EngineName param.Opt[string] `json:"engine_name,omitzero"`
	// The version of the browser engine. Ex: 114.0 Leading and trailing whitespace is
	// removed before validation.
	EngineVersion param.Opt[string] `json:"engine_version,omitzero"`
	// The Pinterest Click ID. Ex: epik456 Leading and trailing whitespace is removed
	// before validation.
	Epik param.Opt[string] `json:"epik,omitzero"`
	// Facebook Click ID with prefix format for Conversions API tracking. Ex:
	// fb.1.1554763741205.AbCdEfGhIjKlMnOpQrStUvWxYz1234567890 Leading and trailing
	// whitespace is removed before validation.
	Fbc param.Opt[string] `json:"fbc,omitzero"`
	// Raw Facebook Click ID query parameter without prefix from ad clicks. Ex:
	// AbCdEfGhIjKlMnOpQrStUvWxYz1234567890 Leading and trailing whitespace is removed
	// before validation.
	Fbclid param.Opt[string] `json:"fbclid,omitzero"`
	// Facebook Browser ID parameter for identifying browsers and attributing events.
	// Ex: fb.1.1554763741205.1098115397 Leading and trailing whitespace is removed
	// before validation.
	Fbp param.Opt[string] `json:"fbp,omitzero"`
	// Deprecated
	Fv param.Opt[bool] `json:"fv,omitzero"`
	// The Google Ad Source. Ex: google Leading and trailing whitespace is removed
	// before validation.
	GadSource param.Opt[string] `json:"gad_source,omitzero"`
	// The Google Braid ID. Ex: gbraid123 Leading and trailing whitespace is removed
	// before validation.
	Gbraid param.Opt[string] `json:"gbraid,omitzero"`
	// The Google Click ID. Ex: gclid123 Leading and trailing whitespace is removed
	// before validation.
	Gclid param.Opt[string] `json:"gclid,omitzero"`
	// The Gravity attribution ID, captured from the `graid` URL parameter on the
	// landing page. Gravity resolves it to the campaign, publisher, and ad IDs
	// server-side. Ex: graid-abc123 Leading and trailing whitespace is removed before
	// validation.
	Graid param.Opt[string] `json:"graid,omitzero"`
	// The Gravity click ID, captured from the `grclid` URL parameter that Gravity
	// (trygravity.com) appends to ad click-through URLs on the landing page. Ex:
	// grclid-abc123 Leading and trailing whitespace is removed before validation.
	Grclid param.Opt[string] `json:"grclid,omitzero"`
	// The host of the current page. Ex: example.com Leading and trailing whitespace is
	// removed before validation.
	Host param.Opt[string] `json:"host,omitzero"`
	// Whether the user is in an iframe. Ex: true
	Iframe param.Opt[bool] `json:"iframe,omitzero"`
	// The Impact Click ID reference. Ex: im_ref123 Leading and trailing whitespace is
	// removed before validation.
	ImRef param.Opt[string] `json:"im_ref,omitzero"`
	// The IP address of the user. Ex: 127.0.0.1 Leading and trailing whitespace is
	// removed before validation.
	IP param.Opt[string] `json:"ip,omitzero"`
	// The Impact Click ID. Ex: irclickid123 Leading and trailing whitespace is removed
	// before validation.
	Irclickid param.Opt[string] `json:"irclickid,omitzero"`
	// The LinkedIn Click ID. Ex: li_fat_id123 Leading and trailing whitespace is
	// removed before validation.
	LiFatID param.Opt[string] `json:"li_fat_id,omitzero"`
	// The Microsoft Click ID. Ex: msclkid123 Leading and trailing whitespace is
	// removed before validation.
	Msclkid param.Opt[string] `json:"msclkid,omitzero"`
	// The NextDoor Click ID. Ex: ndclid123 Leading and trailing whitespace is removed
	// before validation.
	Ndclid param.Opt[string] `json:"ndclid,omitzero"`
	// Deprecated
	NewS param.Opt[bool] `json:"new_s,omitzero"`
	// The Outbrain click ID, captured from the `ob_click_id` URL parameter (Outbrain
	// `{{ob_click_id}}` macro) on the landing page. Ex: ob_click_abc123 Leading and
	// trailing whitespace is removed before validation.
	ObClickID param.Opt[string] `json:"ob_click_id,omitzero"`
	// The OpenAI Ads privacy-preserving reference, captured from the `oppref` URL
	// parameter on landing pages (the OpenAI Pixel also stores it in a `__oppref`
	// cookie). Sent to OpenAI Ads on Conversions API events for attribution. Ex:
	// oppref_abc Leading and trailing whitespace is removed before validation.
	Oppref param.Opt[string] `json:"oppref,omitzero"`
	// The name of the operating system. Ex: Windows Leading and trailing whitespace is
	// removed before validation.
	OsName param.Opt[string] `json:"os_name,omitzero"`
	// The version of the operating system. Ex: 10.0 Leading and trailing whitespace is
	// removed before validation.
	OsVersion param.Opt[string] `json:"os_version,omitzero"`
	// A random set of numbers for the page load
	PageHash param.Opt[float64] `json:"page_hash,omitzero"`
	// The pathname of the current page. Ex: /home Leading and trailing whitespace is
	// removed before validation.
	Pathname param.Opt[string] `json:"pathname,omitzero"`
	// The Quora Click ID. Ex: qclid123 Leading and trailing whitespace is removed
	// before validation.
	Qclid param.Opt[string] `json:"qclid,omitzero"`
	// The Reddit Click ID. Ex: rdt_cid123 Leading and trailing whitespace is removed
	// before validation.
	RdtCid param.Opt[string] `json:"rdt_cid,omitzero"`
	// The time an Ours server received the event, in ISO 8601 UTC
	ReceivedAt param.Opt[string] `json:"received_at,omitzero"`
	// The referrer URL of the current page Leading and trailing whitespace is removed
	// before validation.
	Referrer param.Opt[string] `json:"referrer,omitzero"`
	// The referring domain of the current page Leading and trailing whitespace is
	// removed before validation.
	ReferringDomain param.Opt[string] `json:"referring_domain,omitzero"`
	// The StackAdapt Tracking ID. Ex: sacid123 Leading and trailing whitespace is
	// removed before validation.
	Sacid param.Opt[string] `json:"sacid,omitzero"`
	// The Snapchat Click ID. Ex: sccid123 Leading and trailing whitespace is removed
	// before validation.
	Sccid param.Opt[string] `json:"sccid,omitzero"`
	// The height of the screen. Ex: 1080
	ScreenHeight param.Opt[float64] `json:"screen_height,omitzero"`
	// The width of the screen. Ex: 1920
	ScreenWidth param.Opt[float64] `json:"screen_width,omitzero"`
	// The number of sessions the user has had. Ex: 3
	SessionCount param.Opt[float64] `json:"sessionCount,omitzero"`
	// The session ID as assigned automatically by the web SDK. This is required for
	// session replay Leading and trailing whitespace is removed before validation.
	Sid param.Opt[string] `json:"sid,omitzero"`
	// The screen resolution as width x height in pixels. Ex: 1920x1080 Leading and
	// trailing whitespace is removed before validation.
	Sr param.Opt[string] `json:"sr,omitzero"`
	// The title of the current page Leading and trailing whitespace is removed before
	// validation.
	Title param.Opt[string] `json:"title,omitzero"`
	// The TikTok Click ID. Ex: ttclid123 Leading and trailing whitespace is removed
	// before validation.
	Ttclid param.Opt[string] `json:"ttclid,omitzero"`
	// The Trade Desk impression ID, captured from a configured ttd_impression_id URL
	// parameter using the %%TTD_IMPRESSIONID%% creative macro. Sent as imp on
	// conversions. Leading and trailing whitespace is removed before validation.
	TtdImpressionID param.Opt[string] `json:"ttd_impression_id,omitzero"`
	// The Twitter Click ID. Ex: twclid123 Leading and trailing whitespace is removed
	// before validation.
	Twclid param.Opt[string] `json:"twclid,omitzero"`
	// The browser brands and full versions from User-Agent Client Hints, as
	// brand;version pairs separated by |. Leading and trailing whitespace is removed
	// before validation.
	Uafvl param.Opt[string] `json:"uafvl,omitzero"`
	// The user agent of the browser Leading and trailing whitespace is removed before
	// validation.
	UserAgent param.Opt[string] `json:"user_agent,omitzero"`
	// The UTM Campaign. The web SDK automatically captures this from the query params.
	// Leading and trailing whitespace is removed before validation.
	UtmCampaign param.Opt[string] `json:"utm_campaign,omitzero"`
	// The UTM Content. The web SDK automatically captures this from the query params.
	// Leading and trailing whitespace is removed before validation.
	UtmContent param.Opt[string] `json:"utm_content,omitzero"`
	// The UTM Medium. The web SDK automatically captures this from the query params.
	// Leading and trailing whitespace is removed before validation.
	UtmMedium param.Opt[string] `json:"utm_medium,omitzero"`
	// The UTM Name. The web SDK automatically captures this from the query params.
	// Leading and trailing whitespace is removed before validation.
	UtmName param.Opt[string] `json:"utm_name,omitzero"`
	// The UTM Source. The web SDK automatically captures this from the query params.
	// Leading and trailing whitespace is removed before validation.
	UtmSource param.Opt[string] `json:"utm_source,omitzero"`
	// The UTM Term. The web SDK automatically captures this from the query params.
	// Leading and trailing whitespace is removed before validation.
	UtmTerm param.Opt[string] `json:"utm_term,omitzero"`
	// The SDK version (e.g., web SDK or ingest-sdk-\* via generated SDK headers)
	// Leading and trailing whitespace is removed before validation.
	Version param.Opt[string] `json:"version,omitzero"`
	// The Viant (Adelphic) Click ID, captured from the `viant_click_id` URL parameter
	// (Viant `${ADELPHIC_CLICKID}` macro). Sent as `xid` on Viant postbacks. Ex:
	// viant_click_abc123 Leading and trailing whitespace is removed before validation.
	ViantClickID param.Opt[string] `json:"viant_click_id,omitzero"`
	// The Viant (Adelphic) Impression ID, captured from the `viant_impression_id` URL
	// parameter (Viant `${ADELPHIC_IMPRESSIONID}` macro). Sent as `imp_id` on Viant
	// postbacks for post-view attribution. Ex: viant_imp_abc123 Leading and trailing
	// whitespace is removed before validation.
	ViantImpressionID param.Opt[string] `json:"viant_impression_id,omitzero"`
	// The WBRAID Identifier. The web SDK automatically captures this from the query
	// params. Leading and trailing whitespace is removed before validation.
	Wbraid param.Opt[string] `json:"wbraid,omitzero"`
	// Whether the user is in a webview. Ex: true
	Webview param.Opt[bool] `json:"webview,omitzero"`
	// The Yelp Click ID from the landing URL, sent unchanged and unhashed.
	Ylpcid param.Opt[string] `json:"ylpcid,omitzero"`
	// Any JSON value: string, number, boolean, array, object, or null.
	IsBot TrackEventParamsDefaultPropertiesIsBotUnion `json:"is_bot,omitzero"`
	paramObj
}

func (r TrackEventParamsDefaultProperties) MarshalJSON() (data []byte, err error) {
	type shadow TrackEventParamsDefaultProperties
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TrackEventParamsDefaultProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsDefaultPropertiesIsBotUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsDefaultPropertiesIsBotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsDefaultPropertiesIsBotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsDefaultPropertiesIsBotUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsEventPropertyUnion struct {
	OfString                         param.Opt[string]                                    `json:",omitzero,inline"`
	OfFloat                          param.Opt[float64]                                   `json:",omitzero,inline"`
	OfBool                           param.Opt[bool]                                      `json:",omitzero,inline"`
	OfTrackEventsEventPropertyArray  []TrackEventParamsEventPropertyArrayItemUnion        `json:",omitzero,inline"`
	OfTrackEventsEventPropertyMapMap map[string]TrackEventParamsEventPropertyMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsEventPropertyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfTrackEventsEventPropertyArray,
		u.OfTrackEventsEventPropertyMapMap)
}
func (u *TrackEventParamsEventPropertyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsEventPropertyUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfTrackEventsEventPropertyArray) {
		return &u.OfTrackEventsEventPropertyArray
	} else if !param.IsOmitted(u.OfTrackEventsEventPropertyMapMap) {
		return &u.OfTrackEventsEventPropertyMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsEventPropertyArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsEventPropertyArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsEventPropertyArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsEventPropertyArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsEventPropertyMapItemUnion struct {
	OfString                                param.Opt[string]                                           `json:",omitzero,inline"`
	OfFloat                                 param.Opt[float64]                                          `json:",omitzero,inline"`
	OfBool                                  param.Opt[bool]                                             `json:",omitzero,inline"`
	OfTrackEventsEventPropertyMapItemArray  []TrackEventParamsEventPropertyMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfTrackEventsEventPropertyMapItemMapMap map[string]TrackEventParamsEventPropertyMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsEventPropertyMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfTrackEventsEventPropertyMapItemArray,
		u.OfTrackEventsEventPropertyMapItemMapMap)
}
func (u *TrackEventParamsEventPropertyMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsEventPropertyMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfTrackEventsEventPropertyMapItemArray) {
		return &u.OfTrackEventsEventPropertyMapItemArray
	} else if !param.IsOmitted(u.OfTrackEventsEventPropertyMapItemMapMap) {
		return &u.OfTrackEventsEventPropertyMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsEventPropertyMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsEventPropertyMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsEventPropertyMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsEventPropertyMapItemArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsEventPropertyMapItemMapItemUnion struct {
	OfString                                       param.Opt[string]                                                  `json:",omitzero,inline"`
	OfFloat                                        param.Opt[float64]                                                 `json:",omitzero,inline"`
	OfBool                                         param.Opt[bool]                                                    `json:",omitzero,inline"`
	OfTrackEventsEventPropertyMapItemMapItemArray  []TrackEventParamsEventPropertyMapItemMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfTrackEventsEventPropertyMapItemMapItemMapMap map[string]TrackEventParamsEventPropertyMapItemMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsEventPropertyMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfTrackEventsEventPropertyMapItemMapItemArray,
		u.OfTrackEventsEventPropertyMapItemMapItemMapMap)
}
func (u *TrackEventParamsEventPropertyMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsEventPropertyMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfTrackEventsEventPropertyMapItemMapItemArray) {
		return &u.OfTrackEventsEventPropertyMapItemMapItemArray
	} else if !param.IsOmitted(u.OfTrackEventsEventPropertyMapItemMapItemMapMap) {
		return &u.OfTrackEventsEventPropertyMapItemMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsEventPropertyMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsEventPropertyMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsEventPropertyMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsEventPropertyMapItemMapItemArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsEventPropertyMapItemMapItemMapItemUnion struct {
	OfString                                              param.Opt[string]                                                         `json:",omitzero,inline"`
	OfFloat                                               param.Opt[float64]                                                        `json:",omitzero,inline"`
	OfBool                                                param.Opt[bool]                                                           `json:",omitzero,inline"`
	OfTrackEventsEventPropertyMapItemMapItemMapItemArray  []TrackEventParamsEventPropertyMapItemMapItemMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfTrackEventsEventPropertyMapItemMapItemMapItemMapMap map[string]TrackEventParamsEventPropertyMapItemMapItemMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsEventPropertyMapItemMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfTrackEventsEventPropertyMapItemMapItemMapItemArray,
		u.OfTrackEventsEventPropertyMapItemMapItemMapItemMapMap)
}
func (u *TrackEventParamsEventPropertyMapItemMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsEventPropertyMapItemMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfTrackEventsEventPropertyMapItemMapItemMapItemArray) {
		return &u.OfTrackEventsEventPropertyMapItemMapItemMapItemArray
	} else if !param.IsOmitted(u.OfTrackEventsEventPropertyMapItemMapItemMapItemMapMap) {
		return &u.OfTrackEventsEventPropertyMapItemMapItemMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsEventPropertyMapItemMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsEventPropertyMapItemMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsEventPropertyMapItemMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsEventPropertyMapItemMapItemMapItemArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsEventPropertyMapItemMapItemMapItemMapItemUnion struct {
	OfString                                                    param.Opt[string]                                                         `json:",omitzero,inline"`
	OfFloat                                                     param.Opt[float64]                                                        `json:",omitzero,inline"`
	OfBool                                                      param.Opt[bool]                                                           `json:",omitzero,inline"`
	OfTrackEventsEventPropertyMapItemMapItemMapItemMapItemArray []TrackEventParamsEventPropertyMapItemMapItemMapItemMapItemArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsEventPropertyMapItemMapItemMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfTrackEventsEventPropertyMapItemMapItemMapItemMapItemArray)
}
func (u *TrackEventParamsEventPropertyMapItemMapItemMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsEventPropertyMapItemMapItemMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfTrackEventsEventPropertyMapItemMapItemMapItemMapItemArray) {
		return &u.OfTrackEventsEventPropertyMapItemMapItemMapItemMapItemArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsEventPropertyMapItemMapItemMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsEventPropertyMapItemMapItemMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsEventPropertyMapItemMapItemMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsEventPropertyMapItemMapItemMapItemMapItemArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// End-user network context for server-side calls. Required for probabilistic
// identity resolution when the caller is a backend server rather than an end-user
// browser.
//
// The properties IP, UserAgent are required.
type TrackEventParamsIdentityContext struct {
	// The end-user IP address (not the server IP).
	IP string `json:"ip" api:"required"`
	// The end-user User-Agent string (not the server UA).
	UserAgent string `json:"userAgent" api:"required"`
	paramObj
}

func (r TrackEventParamsIdentityContext) MarshalJSON() (data []byte, err error) {
	type shadow TrackEventParamsIdentityContext
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TrackEventParamsIdentityContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Properties to set on the visitor. (optional) You can also update these
// properties via the identify endpoint.
type TrackEventParamsUserProperties struct {
	// Leading and trailing whitespace is removed before validation.
	EfTransactionID param.Opt[string] `json:"_ef_transaction_id,omitzero"`
	AdID            param.Opt[string] `json:"ad_id,omitzero"`
	// Leading and trailing whitespace is removed before validation.
	AdmitadUid param.Opt[string] `json:"admitad_uid,omitzero"`
	AdsetID    param.Opt[string] `json:"adset_id,omitzero"`
	Alart      param.Opt[string] `json:"alart,omitzero"`
	Aleid      param.Opt[string] `json:"aleid,omitzero"`
	Axwrt      param.Opt[string] `json:"axwrt,omitzero"`
	// Leading and trailing whitespace is removed before validation.
	BasisCid param.Opt[string] `json:"basis_cid,omitzero"`
	// Leading and trailing whitespace is removed before validation.
	BeeswaxAuctionID param.Opt[string] `json:"beeswax_auction_id,omitzero"`
	CampaignID       param.Opt[string] `json:"campaign_id,omitzero"`
	// Non-string input is converted to text before validation; clients should send a
	// string. Leading and trailing whitespace is removed before validation.
	City    param.Opt[string] `json:"city,omitzero"`
	Clickid param.Opt[string] `json:"clickid,omitzero"`
	Clid    param.Opt[string] `json:"clid,omitzero"`
	// Leading and trailing whitespace is removed before validation.
	CompanyName param.Opt[string] `json:"company_name,omitzero"`
	// Non-string input is converted to text before validation; clients should send a
	// string. Leading and trailing whitespace is removed before validation.
	Country param.Opt[string] `json:"country,omitzero"`
	// Non-string input is converted to text before validation; clients should send a
	// string. Leading and trailing whitespace is removed before validation.
	DateOfBirth param.Opt[string] `json:"date_of_birth,omitzero"`
	Dclid       param.Opt[string] `json:"dclid,omitzero"`
	// Non-string input is converted to text before validation; clients should send a
	// string. Leading and trailing whitespace is removed before validation.
	Email param.Opt[string] `json:"email,omitzero"`
	// Leading and trailing whitespace is removed before validation.
	Epik param.Opt[string] `json:"epik,omitzero"`
	// Non-string input is converted to text before validation; clients should send a
	// string. Leading and trailing whitespace is removed before validation.
	ExternalID param.Opt[string] `json:"external_id,omitzero"`
	Fbc        param.Opt[string] `json:"fbc,omitzero"`
	Fbclid     param.Opt[string] `json:"fbclid,omitzero"`
	Fbp        param.Opt[string] `json:"fbp,omitzero"`
	// Non-string input is converted to text before validation; clients should send a
	// string. Leading and trailing whitespace is removed before validation.
	FirstName param.Opt[string] `json:"first_name,omitzero"`
	GadSource param.Opt[string] `json:"gad_source,omitzero"`
	Gbraid    param.Opt[string] `json:"gbraid,omitzero"`
	Gclid     param.Opt[string] `json:"gclid,omitzero"`
	// Non-string input is converted to text before validation; clients should send a
	// string. Leading and trailing whitespace is removed before validation.
	Gender param.Opt[string] `json:"gender,omitzero"`
	// Leading and trailing whitespace is removed before validation.
	Graid param.Opt[string] `json:"graid,omitzero"`
	// Leading and trailing whitespace is removed before validation.
	Grclid param.Opt[string] `json:"grclid,omitzero"`
	ImRef  param.Opt[string] `json:"im_ref,omitzero"`
	// The IP address of the user Leading and trailing whitespace is removed before
	// validation.
	IP        param.Opt[string] `json:"ip,omitzero"`
	Irclickid param.Opt[string] `json:"irclickid,omitzero"`
	// Leading and trailing whitespace is removed before validation.
	JobTitle param.Opt[string] `json:"job_title,omitzero"`
	// Non-string input is converted to text before validation; clients should send a
	// string. Leading and trailing whitespace is removed before validation.
	LastName param.Opt[string] `json:"last_name,omitzero"`
	LiFatID  param.Opt[string] `json:"li_fat_id,omitzero"`
	Msclkid  param.Opt[string] `json:"msclkid,omitzero"`
	Ndclid   param.Opt[string] `json:"ndclid,omitzero"`
	// Leading and trailing whitespace is removed before validation.
	ObClickID       param.Opt[string] `json:"ob_click_id,omitzero"`
	Oppref          param.Opt[string] `json:"oppref,omitzero"`
	Qclid           param.Opt[string] `json:"qclid,omitzero"`
	RdtCid          param.Opt[string] `json:"rdt_cid,omitzero"`
	Referrer        param.Opt[string] `json:"referrer,omitzero"`
	ReferringDomain param.Opt[string] `json:"referring_domain,omitzero"`
	// Leading and trailing whitespace is removed before validation.
	Sacid param.Opt[string] `json:"sacid,omitzero"`
	Sccid param.Opt[string] `json:"sccid,omitzero"`
	// Non-string input is converted to text before validation; clients should send a
	// string.
	Sid param.Opt[string] `json:"sid,omitzero"`
	// Non-string input is converted to text before validation; clients should send a
	// string. Leading and trailing whitespace is removed before validation.
	State  param.Opt[string] `json:"state,omitzero"`
	Ttclid param.Opt[string] `json:"ttclid,omitzero"`
	// Leading and trailing whitespace is removed before validation.
	TtdImpressionID   param.Opt[string] `json:"ttd_impression_id,omitzero"`
	Twclid            param.Opt[string] `json:"twclid,omitzero"`
	UserAgent         param.Opt[string] `json:"user_agent,omitzero"`
	UserAgentFullList param.Opt[string] `json:"user_agent_full_list,omitzero"`
	// Non-string input is converted to text before validation; clients should send a
	// string.
	UtmCampaign param.Opt[string] `json:"utm_campaign,omitzero"`
	// Non-string input is converted to text before validation; clients should send a
	// string.
	UtmContent param.Opt[string] `json:"utm_content,omitzero"`
	// Non-string input is converted to text before validation; clients should send a
	// string.
	UtmMedium param.Opt[string] `json:"utm_medium,omitzero"`
	// Non-string input is converted to text before validation; clients should send a
	// string.
	UtmName param.Opt[string] `json:"utm_name,omitzero"`
	// Non-string input is converted to text before validation; clients should send a
	// string.
	UtmSource param.Opt[string] `json:"utm_source,omitzero"`
	// Non-string input is converted to text before validation; clients should send a
	// string.
	UtmTerm param.Opt[string] `json:"utm_term,omitzero"`
	// Leading and trailing whitespace is removed before validation.
	ViantClickID param.Opt[string] `json:"viant_click_id,omitzero"`
	// Leading and trailing whitespace is removed before validation.
	ViantImpressionID param.Opt[string] `json:"viant_impression_id,omitzero"`
	Wbraid            param.Opt[string] `json:"wbraid,omitzero"`
	Ylpcid            param.Opt[string] `json:"ylpcid,omitzero"`
	// Use at most 200 nonempty property keys and five object levels. Array contents do
	// not count toward the object-depth limit. Empty keys are removed recursively,
	// arrays are truncated to 250 entries, and top-level numbers outside the safe
	// integer range become strings. After normalization, the serialized JSON must be
	// shorter than 1,048,576 UTF-16 code units.
	Consent map[string]TrackEventParamsUserPropertiesConsentUnion `json:"consent,omitzero"`
	// Use at most 200 nonempty property keys and five object levels. Array contents do
	// not count toward the object-depth limit. Empty keys are removed recursively,
	// arrays are truncated to 250 entries, and top-level numbers outside the safe
	// integer range become strings. After normalization, the serialized JSON must be
	// shorter than 1,048,576 UTF-16 code units.
	CustomProperties map[string]TrackEventParamsUserPropertiesCustomPropertyUnion `json:"custom_properties,omitzero"`
	// Any JSON value: string, number, boolean, array, object, or null.
	IsBot TrackEventParamsUserPropertiesIsBotUnion `json:"is_bot,omitzero"`
	// Any JSON value: string, number, boolean, array, object, or null.
	PhoneNumber TrackEventParamsUserPropertiesPhoneNumberUnion `json:"phone_number,omitzero"`
	// Any JSON value: string, number, boolean, array, object, or null.
	Zip TrackEventParamsUserPropertiesZipUnion `json:"zip,omitzero"`
	paramObj
}

func (r TrackEventParamsUserProperties) MarshalJSON() (data []byte, err error) {
	type shadow TrackEventParamsUserProperties
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TrackEventParamsUserProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesConsentUnion struct {
	OfString                                 param.Opt[string]                                            `json:",omitzero,inline"`
	OfFloat                                  param.Opt[float64]                                           `json:",omitzero,inline"`
	OfBool                                   param.Opt[bool]                                              `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesConsentArray  []TrackEventParamsUserPropertiesConsentArrayItemUnion        `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesConsentMapMap map[string]TrackEventParamsUserPropertiesConsentMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesConsentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfTrackEventsUserPropertiesConsentArray,
		u.OfTrackEventsUserPropertiesConsentMapMap)
}
func (u *TrackEventParamsUserPropertiesConsentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesConsentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesConsentArray) {
		return &u.OfTrackEventsUserPropertiesConsentArray
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesConsentMapMap) {
		return &u.OfTrackEventsUserPropertiesConsentMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesConsentArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesConsentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsUserPropertiesConsentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesConsentArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesConsentMapItemUnion struct {
	OfString                                        param.Opt[string]                                                   `json:",omitzero,inline"`
	OfFloat                                         param.Opt[float64]                                                  `json:",omitzero,inline"`
	OfBool                                          param.Opt[bool]                                                     `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesConsentMapItemArray  []TrackEventParamsUserPropertiesConsentMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesConsentMapItemMapMap map[string]TrackEventParamsUserPropertiesConsentMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesConsentMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfTrackEventsUserPropertiesConsentMapItemArray,
		u.OfTrackEventsUserPropertiesConsentMapItemMapMap)
}
func (u *TrackEventParamsUserPropertiesConsentMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesConsentMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesConsentMapItemArray) {
		return &u.OfTrackEventsUserPropertiesConsentMapItemArray
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesConsentMapItemMapMap) {
		return &u.OfTrackEventsUserPropertiesConsentMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesConsentMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesConsentMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsUserPropertiesConsentMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesConsentMapItemArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesConsentMapItemMapItemUnion struct {
	OfString                                               param.Opt[string]                                                          `json:",omitzero,inline"`
	OfFloat                                                param.Opt[float64]                                                         `json:",omitzero,inline"`
	OfBool                                                 param.Opt[bool]                                                            `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesConsentMapItemMapItemArray  []TrackEventParamsUserPropertiesConsentMapItemMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesConsentMapItemMapItemMapMap map[string]TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesConsentMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfTrackEventsUserPropertiesConsentMapItemMapItemArray,
		u.OfTrackEventsUserPropertiesConsentMapItemMapItemMapMap)
}
func (u *TrackEventParamsUserPropertiesConsentMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesConsentMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesConsentMapItemMapItemArray) {
		return &u.OfTrackEventsUserPropertiesConsentMapItemMapItemArray
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesConsentMapItemMapItemMapMap) {
		return &u.OfTrackEventsUserPropertiesConsentMapItemMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesConsentMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesConsentMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsUserPropertiesConsentMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesConsentMapItemMapItemArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemUnion struct {
	OfString                                                      param.Opt[string]                                                                 `json:",omitzero,inline"`
	OfFloat                                                       param.Opt[float64]                                                                `json:",omitzero,inline"`
	OfBool                                                        param.Opt[bool]                                                                   `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesConsentMapItemMapItemMapItemArray  []TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesConsentMapItemMapItemMapItemMapMap map[string]TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfTrackEventsUserPropertiesConsentMapItemMapItemMapItemArray,
		u.OfTrackEventsUserPropertiesConsentMapItemMapItemMapItemMapMap)
}
func (u *TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesConsentMapItemMapItemMapItemArray) {
		return &u.OfTrackEventsUserPropertiesConsentMapItemMapItemMapItemArray
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesConsentMapItemMapItemMapItemMapMap) {
		return &u.OfTrackEventsUserPropertiesConsentMapItemMapItemMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemMapItemUnion struct {
	OfString                                                            param.Opt[string]                                                                 `json:",omitzero,inline"`
	OfFloat                                                             param.Opt[float64]                                                                `json:",omitzero,inline"`
	OfBool                                                              param.Opt[bool]                                                                   `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesConsentMapItemMapItemMapItemMapItemArray []TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemMapItemArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfTrackEventsUserPropertiesConsentMapItemMapItemMapItemMapItemArray)
}
func (u *TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesConsentMapItemMapItemMapItemMapItemArray) {
		return &u.OfTrackEventsUserPropertiesConsentMapItemMapItemMapItemMapItemArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesConsentMapItemMapItemMapItemMapItemArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesCustomPropertyUnion struct {
	OfString                                        param.Opt[string]                                                   `json:",omitzero,inline"`
	OfFloat                                         param.Opt[float64]                                                  `json:",omitzero,inline"`
	OfBool                                          param.Opt[bool]                                                     `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesCustomPropertyArray  []TrackEventParamsUserPropertiesCustomPropertyArrayItemUnion        `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesCustomPropertyMapMap map[string]TrackEventParamsUserPropertiesCustomPropertyMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesCustomPropertyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfTrackEventsUserPropertiesCustomPropertyArray,
		u.OfTrackEventsUserPropertiesCustomPropertyMapMap)
}
func (u *TrackEventParamsUserPropertiesCustomPropertyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesCustomPropertyUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesCustomPropertyArray) {
		return &u.OfTrackEventsUserPropertiesCustomPropertyArray
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesCustomPropertyMapMap) {
		return &u.OfTrackEventsUserPropertiesCustomPropertyMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesCustomPropertyArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesCustomPropertyArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsUserPropertiesCustomPropertyArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesCustomPropertyArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesCustomPropertyMapItemUnion struct {
	OfString                                               param.Opt[string]                                                          `json:",omitzero,inline"`
	OfFloat                                                param.Opt[float64]                                                         `json:",omitzero,inline"`
	OfBool                                                 param.Opt[bool]                                                            `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesCustomPropertyMapItemArray  []TrackEventParamsUserPropertiesCustomPropertyMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesCustomPropertyMapItemMapMap map[string]TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesCustomPropertyMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfTrackEventsUserPropertiesCustomPropertyMapItemArray,
		u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapMap)
}
func (u *TrackEventParamsUserPropertiesCustomPropertyMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesCustomPropertyMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesCustomPropertyMapItemArray) {
		return &u.OfTrackEventsUserPropertiesCustomPropertyMapItemArray
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapMap) {
		return &u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesCustomPropertyMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesCustomPropertyMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsUserPropertiesCustomPropertyMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesCustomPropertyMapItemArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemUnion struct {
	OfString                                                      param.Opt[string]                                                                 `json:",omitzero,inline"`
	OfFloat                                                       param.Opt[float64]                                                                `json:",omitzero,inline"`
	OfBool                                                        param.Opt[bool]                                                                   `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemArray  []TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemMapMap map[string]TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemArray,
		u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemMapMap)
}
func (u *TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemArray) {
		return &u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemArray
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemMapMap) {
		return &u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemUnion struct {
	OfString                                                             param.Opt[string]                                                                        `json:",omitzero,inline"`
	OfFloat                                                              param.Opt[float64]                                                                       `json:",omitzero,inline"`
	OfBool                                                               param.Opt[bool]                                                                          `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemMapItemArray  []TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemMapItemMapMap map[string]TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemMapItemArray,
		u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemMapItemMapMap)
}
func (u *TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemMapItemArray) {
		return &u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemMapItemArray
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemMapItemMapMap) {
		return &u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemUnion struct {
	OfString                                                                   param.Opt[string]                                                                        `json:",omitzero,inline"`
	OfFloat                                                                    param.Opt[float64]                                                                       `json:",omitzero,inline"`
	OfBool                                                                     param.Opt[bool]                                                                          `json:",omitzero,inline"`
	OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArray []TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArray)
}
func (u *TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArray) {
		return &u.OfTrackEventsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesIsBotUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesIsBotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsUserPropertiesIsBotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesIsBotUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesPhoneNumberUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesPhoneNumberUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsUserPropertiesPhoneNumberUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesPhoneNumberUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrackEventParamsUserPropertiesZipUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TrackEventParamsUserPropertiesZipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TrackEventParamsUserPropertiesZipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrackEventParamsUserPropertiesZipUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}
