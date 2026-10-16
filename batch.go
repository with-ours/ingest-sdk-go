// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/with-ours/ingest-sdk-go/internal/apijson"
	"github.com/with-ours/ingest-sdk-go/internal/requestconfig"
	"github.com/with-ours/ingest-sdk-go/option"
	"github.com/with-ours/ingest-sdk-go/packages/param"
	"github.com/with-ours/ingest-sdk-go/packages/respjson"
)

// BatchService contains methods and other services that help with interacting with
// the ours-privacy API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBatchService] method instead.
type BatchService struct {
	Options []option.RequestOption
}

// NewBatchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBatchService(opts ...option.RequestOption) (r BatchService) {
	r = BatchService{}
	r.Options = opts
	return
}

// Send multiple `/track`-shaped events in a single request. This endpoint is
// intended for replay, backfill, and asynchronous bulk delivery. The top-level
// token is authorized once for the full batch. Each batch row must include
// `distinctId`, and mixed validation or queue outcomes are reported per row.
func (r *BatchService) New(ctx context.Context, body BatchNewParams, opts ...option.RequestOption) (res *BatchNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.oursprivacy.com/api/v1/")}, opts...)
	path := "batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type BatchNewResponse struct {
	Accepted int64                         `json:"accepted" api:"required"`
	Failed   int64                         `json:"failed" api:"required"`
	Results  []BatchNewResponseResultUnion `json:"results" api:"required"`
	Success  bool                          `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accepted    respjson.Field
		Failed      respjson.Field
		Results     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchNewResponseResultUnion contains all possible properties and values from
// [BatchNewResponseResultObject], [BatchNewResponseResultObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BatchNewResponseResultUnion struct {
	Index   int64 `json:"index"`
	Success bool  `json:"success"`
	// This field is from variant [BatchNewResponseResultObject2].
	Code string `json:"code"`
	// This field is from variant [BatchNewResponseResultObject2].
	Message string `json:"message"`
	JSON    struct {
		Index   respjson.Field
		Success respjson.Field
		Code    respjson.Field
		Message respjson.Field
		raw     string
	} `json:"-"`
}

func (u BatchNewResponseResultUnion) AsBatchNewResponseResultObject() (v BatchNewResponseResultObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BatchNewResponseResultUnion) AsBatchNewResponseResultObject2() (v BatchNewResponseResultObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BatchNewResponseResultUnion) RawJSON() string { return u.JSON.raw }

func (r *BatchNewResponseResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchNewResponseResultObject struct {
	Index int64 `json:"index" api:"required"`
	// Any of true.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Index       respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponseResultObject) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponseResultObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchNewResponseResultObject2 struct {
	// Any of "invalid_event", "queue_failed".
	Code    string `json:"code" api:"required"`
	Index   int64  `json:"index" api:"required"`
	Message string `json:"message" api:"required"`
	// Any of false.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Index       respjson.Field
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponseResultObject2) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponseResultObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchNewParams struct {
	// The token for your Source. You can find this in the dashboard. Leading and
	// trailing whitespace is removed before validation.
	Token string `json:"token" api:"required"`
	// A list of `/track`-shaped events to validate and enqueue together.
	Events []BatchNewParamsEvent `json:"events,omitzero" api:"required"`
	paramObj
}

func (r BatchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DistinctID, Event are required.
type BatchNewParamsEvent struct {
	// A unique identifier for this event used for deduplication. Highly recommended —
	// if omitted, Ours will generate one for you, but supplying your own gives you
	// stronger idempotency guarantees (e.g. a Stripe payment intent ID or your
	// internal order ID). Leading and trailing whitespace is removed before
	// validation.
	DistinctID string `json:"distinctId" api:"required"`
	// The name of the event you're tracking. This must be whitelisted in the Ours
	// dashboard. Leading and trailing whitespace is removed before validation.
	Event string `json:"event" api:"required"`
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
	DefaultProperties BatchNewParamsEventDefaultProperties `json:"defaultProperties,omitzero"`
	// Any additional event properties you want to pass along. Use at most 200 nonempty
	// property keys and five object levels. Array contents do not count toward the
	// object-depth limit. Empty keys are removed recursively, arrays are truncated to
	// 250 entries, and top-level numbers outside the safe integer range become
	// strings. After normalization, the serialized JSON must be shorter than 1,048,576
	// UTF-16 code units.
	EventProperties map[string]BatchNewParamsEventEventPropertyUnion `json:"eventProperties,omitzero"`
	// End-user network context for server-side calls. Required for probabilistic
	// identity resolution when the caller is a backend server rather than an end-user
	// browser.
	IdentityContext BatchNewParamsEventIdentityContext `json:"identityContext,omitzero"`
	// Properties to set on the visitor. (optional) You can also update these
	// properties via the identify endpoint.
	UserProperties BatchNewParamsEventUserProperties `json:"userProperties,omitzero"`
	paramObj
}

func (r BatchNewParamsEvent) MarshalJSON() (data []byte, err error) {
	type shadow BatchNewParamsEvent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchNewParamsEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are used throughout the Ours app to pass known values onto
// destinations
type BatchNewParamsEventDefaultProperties struct {
	// The Everflow affiliate Click (Transaction) ID, captured from the
	// `_ef_transaction_id` URL parameter. Ex: ef_click_abc123 Leading and trailing
	// whitespace is removed before validation.
	EfTransactionID param.Opt[string] `json:"_ef_transaction_id,omitzero"`
	// The active time in milliseconds that the user had this tab active
	ActiveDuration param.Opt[float64] `json:"activeDuration,omitzero"`
	// The ad id for detected in the session. This is set by the web sdk automatically.
	// Leading and trailing whitespace is removed before validation.
	AdID param.Opt[string] `json:"ad_id,omitzero"`
	// The Admitad (Mitgo) affiliate Click ID. Ex: admitad_uid_abc123 Leading and
	// trailing whitespace is removed before validation.
	AdmitadUid param.Opt[string] `json:"admitad_uid,omitzero"`
	// The adset id for detected in the session. This is set by the web sdk
	// automatically. Leading and trailing whitespace is removed before validation.
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
	// The language of the browser. Ex: en-US Leading and trailing whitespace is
	// removed before validation.
	BrowserLanguage param.Opt[string] `json:"browser_language,omitzero"`
	// The name of the browser. Ex: Chrome Leading and trailing whitespace is removed
	// before validation.
	BrowserName param.Opt[string] `json:"browser_name,omitzero"`
	// The version of the browser. Ex: 114.0 Leading and trailing whitespace is removed
	// before validation.
	BrowserVersion param.Opt[string] `json:"browser_version,omitzero"`
	// The campaign id for detected in the session. This is set by the web sdk
	// automatically. Leading and trailing whitespace is removed before validation.
	CampaignID param.Opt[string] `json:"campaign_id,omitzero"`
	// The Click ID. Ex: clickid123 Leading and trailing whitespace is removed before
	// validation.
	Clickid param.Opt[string] `json:"clickid,omitzero"`
	// The Generic Click ID. Ex: clid123 Leading and trailing whitespace is removed
	// before validation.
	Clid param.Opt[string] `json:"clid,omitzero"`
	// The architecture of the CPU. Ex: x64 Leading and trailing whitespace is removed
	// before validation.
	CPUArchitecture param.Opt[string] `json:"cpu_architecture,omitzero"`
	// The full url (including query params) of the current page Leading and trailing
	// whitespace is removed before validation.
	CurrentURL param.Opt[string] `json:"current_url,omitzero"`
	// The DoubleClick Click ID. Ex: dclid123 Leading and trailing whitespace is
	// removed before validation.
	Dclid param.Opt[string] `json:"dclid,omitzero"`
	// The model of the device. Ex: iPhone 13 Leading and trailing whitespace is
	// removed before validation.
	DeviceModel param.Opt[string] `json:"device_model,omitzero"`
	// The type of device the user is using. Ex: mobile Leading and trailing whitespace
	// is removed before validation.
	DeviceType param.Opt[string] `json:"device_type,omitzero"`
	// The vendor of the device. Ex: Apple Leading and trailing whitespace is removed
	// before validation.
	DeviceVendor param.Opt[string] `json:"device_vendor,omitzero"`
	// The time in milliseconds since the page was loaded // script was loaded
	Duration param.Opt[float64] `json:"duration,omitzero"`
	// The browsers encoding. Ex: UTF-8 Leading and trailing whitespace is removed
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
	// The time the event was received by an Ours server in ISO format
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
	// The SnapChat Click ID. Ex: sccid123 Leading and trailing whitespace is removed
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
	// Leading and trailing whitespace is removed before validation.
	Sr param.Opt[string] `json:"sr,omitzero"`
	// The title of the current page Leading and trailing whitespace is removed before
	// validation.
	Title param.Opt[string] `json:"title,omitzero"`
	// The TikTok Click ID. Ex: ttclid123 Leading and trailing whitespace is removed
	// before validation.
	Ttclid param.Opt[string] `json:"ttclid,omitzero"`
	// The Twitter Click ID. Ex: twclid123 Leading and trailing whitespace is removed
	// before validation.
	Twclid param.Opt[string] `json:"twclid,omitzero"`
	// User agent as a full list of strings. Leading and trailing whitespace is removed
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
	// Any JSON value: string, number, boolean, array, object, or null.
	IsBot BatchNewParamsEventDefaultPropertiesIsBotUnion `json:"is_bot,omitzero"`
	paramObj
}

func (r BatchNewParamsEventDefaultProperties) MarshalJSON() (data []byte, err error) {
	type shadow BatchNewParamsEventDefaultProperties
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchNewParamsEventDefaultProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchNewParamsEventDefaultPropertiesIsBotUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventDefaultPropertiesIsBotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventDefaultPropertiesIsBotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventDefaultPropertiesIsBotUnion) asAny() any {
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
type BatchNewParamsEventEventPropertyUnion struct {
	OfString                            param.Opt[string]                                       `json:",omitzero,inline"`
	OfFloat                             param.Opt[float64]                                      `json:",omitzero,inline"`
	OfBool                              param.Opt[bool]                                         `json:",omitzero,inline"`
	OfBatchNewsEventEventPropertyArray  []BatchNewParamsEventEventPropertyArrayItemUnion        `json:",omitzero,inline"`
	OfBatchNewsEventEventPropertyMapMap map[string]BatchNewParamsEventEventPropertyMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventEventPropertyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfBatchNewsEventEventPropertyArray,
		u.OfBatchNewsEventEventPropertyMapMap)
}
func (u *BatchNewParamsEventEventPropertyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventEventPropertyUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfBatchNewsEventEventPropertyArray) {
		return &u.OfBatchNewsEventEventPropertyArray
	} else if !param.IsOmitted(u.OfBatchNewsEventEventPropertyMapMap) {
		return &u.OfBatchNewsEventEventPropertyMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchNewParamsEventEventPropertyArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventEventPropertyArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventEventPropertyArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventEventPropertyArrayItemUnion) asAny() any {
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
type BatchNewParamsEventEventPropertyMapItemUnion struct {
	OfString                                   param.Opt[string]                                              `json:",omitzero,inline"`
	OfFloat                                    param.Opt[float64]                                             `json:",omitzero,inline"`
	OfBool                                     param.Opt[bool]                                                `json:",omitzero,inline"`
	OfBatchNewsEventEventPropertyMapItemArray  []BatchNewParamsEventEventPropertyMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfBatchNewsEventEventPropertyMapItemMapMap map[string]BatchNewParamsEventEventPropertyMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventEventPropertyMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfBatchNewsEventEventPropertyMapItemArray,
		u.OfBatchNewsEventEventPropertyMapItemMapMap)
}
func (u *BatchNewParamsEventEventPropertyMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventEventPropertyMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfBatchNewsEventEventPropertyMapItemArray) {
		return &u.OfBatchNewsEventEventPropertyMapItemArray
	} else if !param.IsOmitted(u.OfBatchNewsEventEventPropertyMapItemMapMap) {
		return &u.OfBatchNewsEventEventPropertyMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchNewParamsEventEventPropertyMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventEventPropertyMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventEventPropertyMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventEventPropertyMapItemArrayItemUnion) asAny() any {
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
type BatchNewParamsEventEventPropertyMapItemMapItemUnion struct {
	OfString                                          param.Opt[string]                                                     `json:",omitzero,inline"`
	OfFloat                                           param.Opt[float64]                                                    `json:",omitzero,inline"`
	OfBool                                            param.Opt[bool]                                                       `json:",omitzero,inline"`
	OfBatchNewsEventEventPropertyMapItemMapItemArray  []BatchNewParamsEventEventPropertyMapItemMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfBatchNewsEventEventPropertyMapItemMapItemMapMap map[string]BatchNewParamsEventEventPropertyMapItemMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventEventPropertyMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfBatchNewsEventEventPropertyMapItemMapItemArray,
		u.OfBatchNewsEventEventPropertyMapItemMapItemMapMap)
}
func (u *BatchNewParamsEventEventPropertyMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventEventPropertyMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfBatchNewsEventEventPropertyMapItemMapItemArray) {
		return &u.OfBatchNewsEventEventPropertyMapItemMapItemArray
	} else if !param.IsOmitted(u.OfBatchNewsEventEventPropertyMapItemMapItemMapMap) {
		return &u.OfBatchNewsEventEventPropertyMapItemMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchNewParamsEventEventPropertyMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventEventPropertyMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventEventPropertyMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventEventPropertyMapItemMapItemArrayItemUnion) asAny() any {
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
type BatchNewParamsEventEventPropertyMapItemMapItemMapItemUnion struct {
	OfString                                                 param.Opt[string]                                                            `json:",omitzero,inline"`
	OfFloat                                                  param.Opt[float64]                                                           `json:",omitzero,inline"`
	OfBool                                                   param.Opt[bool]                                                              `json:",omitzero,inline"`
	OfBatchNewsEventEventPropertyMapItemMapItemMapItemArray  []BatchNewParamsEventEventPropertyMapItemMapItemMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfBatchNewsEventEventPropertyMapItemMapItemMapItemMapMap map[string]BatchNewParamsEventEventPropertyMapItemMapItemMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventEventPropertyMapItemMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfBatchNewsEventEventPropertyMapItemMapItemMapItemArray,
		u.OfBatchNewsEventEventPropertyMapItemMapItemMapItemMapMap)
}
func (u *BatchNewParamsEventEventPropertyMapItemMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventEventPropertyMapItemMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfBatchNewsEventEventPropertyMapItemMapItemMapItemArray) {
		return &u.OfBatchNewsEventEventPropertyMapItemMapItemMapItemArray
	} else if !param.IsOmitted(u.OfBatchNewsEventEventPropertyMapItemMapItemMapItemMapMap) {
		return &u.OfBatchNewsEventEventPropertyMapItemMapItemMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchNewParamsEventEventPropertyMapItemMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventEventPropertyMapItemMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventEventPropertyMapItemMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventEventPropertyMapItemMapItemMapItemArrayItemUnion) asAny() any {
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
type BatchNewParamsEventEventPropertyMapItemMapItemMapItemMapItemUnion struct {
	OfString                                                       param.Opt[string]                                                            `json:",omitzero,inline"`
	OfFloat                                                        param.Opt[float64]                                                           `json:",omitzero,inline"`
	OfBool                                                         param.Opt[bool]                                                              `json:",omitzero,inline"`
	OfBatchNewsEventEventPropertyMapItemMapItemMapItemMapItemArray []BatchNewParamsEventEventPropertyMapItemMapItemMapItemMapItemArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventEventPropertyMapItemMapItemMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfBatchNewsEventEventPropertyMapItemMapItemMapItemMapItemArray)
}
func (u *BatchNewParamsEventEventPropertyMapItemMapItemMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventEventPropertyMapItemMapItemMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfBatchNewsEventEventPropertyMapItemMapItemMapItemMapItemArray) {
		return &u.OfBatchNewsEventEventPropertyMapItemMapItemMapItemMapItemArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchNewParamsEventEventPropertyMapItemMapItemMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventEventPropertyMapItemMapItemMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventEventPropertyMapItemMapItemMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventEventPropertyMapItemMapItemMapItemMapItemArrayItemUnion) asAny() any {
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
type BatchNewParamsEventIdentityContext struct {
	// The end-user IP address (not the server IP).
	IP string `json:"ip" api:"required"`
	// The end-user User-Agent string (not the server UA).
	UserAgent string `json:"userAgent" api:"required"`
	paramObj
}

func (r BatchNewParamsEventIdentityContext) MarshalJSON() (data []byte, err error) {
	type shadow BatchNewParamsEventIdentityContext
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchNewParamsEventIdentityContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Properties to set on the visitor. (optional) You can also update these
// properties via the identify endpoint.
type BatchNewParamsEventUserProperties struct {
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
	State             param.Opt[string] `json:"state,omitzero"`
	Ttclid            param.Opt[string] `json:"ttclid,omitzero"`
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
	// Use at most 200 nonempty property keys and five object levels. Array contents do
	// not count toward the object-depth limit. Empty keys are removed recursively,
	// arrays are truncated to 250 entries, and top-level numbers outside the safe
	// integer range become strings. After normalization, the serialized JSON must be
	// shorter than 1,048,576 UTF-16 code units.
	Consent map[string]BatchNewParamsEventUserPropertiesConsentUnion `json:"consent,omitzero"`
	// Use at most 200 nonempty property keys and five object levels. Array contents do
	// not count toward the object-depth limit. Empty keys are removed recursively,
	// arrays are truncated to 250 entries, and top-level numbers outside the safe
	// integer range become strings. After normalization, the serialized JSON must be
	// shorter than 1,048,576 UTF-16 code units.
	CustomProperties map[string]BatchNewParamsEventUserPropertiesCustomPropertyUnion `json:"custom_properties,omitzero"`
	// Any JSON value: string, number, boolean, array, object, or null.
	IsBot BatchNewParamsEventUserPropertiesIsBotUnion `json:"is_bot,omitzero"`
	// Any JSON value: string, number, boolean, array, object, or null.
	PhoneNumber BatchNewParamsEventUserPropertiesPhoneNumberUnion `json:"phone_number,omitzero"`
	// Any JSON value: string, number, boolean, array, object, or null.
	Zip BatchNewParamsEventUserPropertiesZipUnion `json:"zip,omitzero"`
	paramObj
}

func (r BatchNewParamsEventUserProperties) MarshalJSON() (data []byte, err error) {
	type shadow BatchNewParamsEventUserProperties
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchNewParamsEventUserProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchNewParamsEventUserPropertiesConsentUnion struct {
	OfString                                    param.Opt[string]                                               `json:",omitzero,inline"`
	OfFloat                                     param.Opt[float64]                                              `json:",omitzero,inline"`
	OfBool                                      param.Opt[bool]                                                 `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesConsentArray  []BatchNewParamsEventUserPropertiesConsentArrayItemUnion        `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesConsentMapMap map[string]BatchNewParamsEventUserPropertiesConsentMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesConsentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfBatchNewsEventUserPropertiesConsentArray,
		u.OfBatchNewsEventUserPropertiesConsentMapMap)
}
func (u *BatchNewParamsEventUserPropertiesConsentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesConsentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesConsentArray) {
		return &u.OfBatchNewsEventUserPropertiesConsentArray
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesConsentMapMap) {
		return &u.OfBatchNewsEventUserPropertiesConsentMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchNewParamsEventUserPropertiesConsentArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesConsentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventUserPropertiesConsentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesConsentArrayItemUnion) asAny() any {
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
type BatchNewParamsEventUserPropertiesConsentMapItemUnion struct {
	OfString                                           param.Opt[string]                                                      `json:",omitzero,inline"`
	OfFloat                                            param.Opt[float64]                                                     `json:",omitzero,inline"`
	OfBool                                             param.Opt[bool]                                                        `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesConsentMapItemArray  []BatchNewParamsEventUserPropertiesConsentMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesConsentMapItemMapMap map[string]BatchNewParamsEventUserPropertiesConsentMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesConsentMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfBatchNewsEventUserPropertiesConsentMapItemArray,
		u.OfBatchNewsEventUserPropertiesConsentMapItemMapMap)
}
func (u *BatchNewParamsEventUserPropertiesConsentMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesConsentMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesConsentMapItemArray) {
		return &u.OfBatchNewsEventUserPropertiesConsentMapItemArray
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesConsentMapItemMapMap) {
		return &u.OfBatchNewsEventUserPropertiesConsentMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchNewParamsEventUserPropertiesConsentMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesConsentMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventUserPropertiesConsentMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesConsentMapItemArrayItemUnion) asAny() any {
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
type BatchNewParamsEventUserPropertiesConsentMapItemMapItemUnion struct {
	OfString                                                  param.Opt[string]                                                             `json:",omitzero,inline"`
	OfFloat                                                   param.Opt[float64]                                                            `json:",omitzero,inline"`
	OfBool                                                    param.Opt[bool]                                                               `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesConsentMapItemMapItemArray  []BatchNewParamsEventUserPropertiesConsentMapItemMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesConsentMapItemMapItemMapMap map[string]BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesConsentMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfBatchNewsEventUserPropertiesConsentMapItemMapItemArray,
		u.OfBatchNewsEventUserPropertiesConsentMapItemMapItemMapMap)
}
func (u *BatchNewParamsEventUserPropertiesConsentMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesConsentMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesConsentMapItemMapItemArray) {
		return &u.OfBatchNewsEventUserPropertiesConsentMapItemMapItemArray
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesConsentMapItemMapItemMapMap) {
		return &u.OfBatchNewsEventUserPropertiesConsentMapItemMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchNewParamsEventUserPropertiesConsentMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesConsentMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventUserPropertiesConsentMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesConsentMapItemMapItemArrayItemUnion) asAny() any {
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
type BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemUnion struct {
	OfString                                                         param.Opt[string]                                                                    `json:",omitzero,inline"`
	OfFloat                                                          param.Opt[float64]                                                                   `json:",omitzero,inline"`
	OfBool                                                           param.Opt[bool]                                                                      `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesConsentMapItemMapItemMapItemArray  []BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesConsentMapItemMapItemMapItemMapMap map[string]BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfBatchNewsEventUserPropertiesConsentMapItemMapItemMapItemArray,
		u.OfBatchNewsEventUserPropertiesConsentMapItemMapItemMapItemMapMap)
}
func (u *BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesConsentMapItemMapItemMapItemArray) {
		return &u.OfBatchNewsEventUserPropertiesConsentMapItemMapItemMapItemArray
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesConsentMapItemMapItemMapItemMapMap) {
		return &u.OfBatchNewsEventUserPropertiesConsentMapItemMapItemMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemArrayItemUnion) asAny() any {
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
type BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemMapItemUnion struct {
	OfString                                                               param.Opt[string]                                                                    `json:",omitzero,inline"`
	OfFloat                                                                param.Opt[float64]                                                                   `json:",omitzero,inline"`
	OfBool                                                                 param.Opt[bool]                                                                      `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesConsentMapItemMapItemMapItemMapItemArray []BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemMapItemArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfBatchNewsEventUserPropertiesConsentMapItemMapItemMapItemMapItemArray)
}
func (u *BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesConsentMapItemMapItemMapItemMapItemArray) {
		return &u.OfBatchNewsEventUserPropertiesConsentMapItemMapItemMapItemMapItemArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesConsentMapItemMapItemMapItemMapItemArrayItemUnion) asAny() any {
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
type BatchNewParamsEventUserPropertiesCustomPropertyUnion struct {
	OfString                                           param.Opt[string]                                                      `json:",omitzero,inline"`
	OfFloat                                            param.Opt[float64]                                                     `json:",omitzero,inline"`
	OfBool                                             param.Opt[bool]                                                        `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesCustomPropertyArray  []BatchNewParamsEventUserPropertiesCustomPropertyArrayItemUnion        `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesCustomPropertyMapMap map[string]BatchNewParamsEventUserPropertiesCustomPropertyMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesCustomPropertyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfBatchNewsEventUserPropertiesCustomPropertyArray,
		u.OfBatchNewsEventUserPropertiesCustomPropertyMapMap)
}
func (u *BatchNewParamsEventUserPropertiesCustomPropertyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesCustomPropertyUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesCustomPropertyArray) {
		return &u.OfBatchNewsEventUserPropertiesCustomPropertyArray
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesCustomPropertyMapMap) {
		return &u.OfBatchNewsEventUserPropertiesCustomPropertyMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchNewParamsEventUserPropertiesCustomPropertyArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesCustomPropertyArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventUserPropertiesCustomPropertyArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesCustomPropertyArrayItemUnion) asAny() any {
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
type BatchNewParamsEventUserPropertiesCustomPropertyMapItemUnion struct {
	OfString                                                  param.Opt[string]                                                             `json:",omitzero,inline"`
	OfFloat                                                   param.Opt[float64]                                                            `json:",omitzero,inline"`
	OfBool                                                    param.Opt[bool]                                                               `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesCustomPropertyMapItemArray  []BatchNewParamsEventUserPropertiesCustomPropertyMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapMap map[string]BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesCustomPropertyMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemArray,
		u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapMap)
}
func (u *BatchNewParamsEventUserPropertiesCustomPropertyMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesCustomPropertyMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemArray) {
		return &u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemArray
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapMap) {
		return &u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchNewParamsEventUserPropertiesCustomPropertyMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesCustomPropertyMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventUserPropertiesCustomPropertyMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesCustomPropertyMapItemArrayItemUnion) asAny() any {
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
type BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemUnion struct {
	OfString                                                         param.Opt[string]                                                                    `json:",omitzero,inline"`
	OfFloat                                                          param.Opt[float64]                                                                   `json:",omitzero,inline"`
	OfBool                                                           param.Opt[bool]                                                                      `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemArray  []BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemMapMap map[string]BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemArray,
		u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemMapMap)
}
func (u *BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemArray) {
		return &u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemArray
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemMapMap) {
		return &u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemArrayItemUnion) asAny() any {
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
type BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemUnion struct {
	OfString                                                                param.Opt[string]                                                                           `json:",omitzero,inline"`
	OfFloat                                                                 param.Opt[float64]                                                                          `json:",omitzero,inline"`
	OfBool                                                                  param.Opt[bool]                                                                             `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemMapItemArray  []BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapMap map[string]BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemMapItemArray,
		u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapMap)
}
func (u *BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemMapItemArray) {
		return &u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemMapItemArray
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapMap) {
		return &u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemArrayItemUnion) asAny() any {
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
type BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemUnion struct {
	OfString                                                                      param.Opt[string]                                                                           `json:",omitzero,inline"`
	OfFloat                                                                       param.Opt[float64]                                                                          `json:",omitzero,inline"`
	OfBool                                                                        param.Opt[bool]                                                                             `json:",omitzero,inline"`
	OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArray []BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArray)
}
func (u *BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArray) {
		return &u.OfBatchNewsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArrayItemUnion) asAny() any {
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
type BatchNewParamsEventUserPropertiesIsBotUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesIsBotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventUserPropertiesIsBotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesIsBotUnion) asAny() any {
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
type BatchNewParamsEventUserPropertiesPhoneNumberUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesPhoneNumberUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventUserPropertiesPhoneNumberUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesPhoneNumberUnion) asAny() any {
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
type BatchNewParamsEventUserPropertiesZipUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchNewParamsEventUserPropertiesZipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *BatchNewParamsEventUserPropertiesZipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BatchNewParamsEventUserPropertiesZipUnion) asAny() any {
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
