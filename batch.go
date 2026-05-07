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

// Send multiple `/track`-shaped events in a single request. The top-level token is
// authorized once for the full batch. Each batch row must include `distinctId`,
// and mixed validation or queue outcomes are reported per row.
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
	// The token for your Source. You can find this in the dashboard.
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
	// A unique identifier for the event. This helps prevent duplicate events.
	DistinctID string `json:"distinctId" api:"required"`
	// The name of the event you're tracking. This must be whitelisted in the Ours
	// dashboard.
	Event string `json:"event" api:"required"`
	// The email address of a user. We will associate this event with the user or
	// create a user. Used for lookup if externalId and userId are not included in the
	// request.
	Email param.Opt[string] `json:"email,omitzero"`
	// The externalId (the ID in your system) of a user. We will associate this event
	// with the user or create a user. If included in the request, email lookup is
	// ignored.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// The time at which the event occurred in milliseconds since UTC epoch. The time
	// must be in the past and within the last 7 days.
	Time param.Opt[float64] `json:"time,omitzero"`
	// The Ours user id stored in local storage and cookies on your web properties. If
	// userId is included in the request, we do not lookup the user by email or
	// externalId.
	UserID param.Opt[string] `json:"userId,omitzero"`
	// These properties are used throughout the Ours app to pass known values onto
	// destinations
	DefaultProperties BatchNewParamsEventDefaultProperties `json:"defaultProperties,omitzero"`
	// Any additional event properties you want to pass along.
	EventProperties map[string]string `json:"eventProperties,omitzero"`
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
	// The active time in milliseconds that the user had this tab active
	ActiveDuration param.Opt[float64] `json:"activeDuration,omitzero"`
	// The ad id for detected in the session. This is set by the web sdk automatically.
	AdID param.Opt[string] `json:"ad_id,omitzero"`
	// The Admitad (Mitgo) affiliate Click ID. Ex: admitad_uid_abc123
	AdmitadUid param.Opt[string] `json:"admitad_uid,omitzero"`
	// The adset id for detected in the session. This is set by the web sdk
	// automatically.
	AdsetID param.Opt[string] `json:"adset_id,omitzero"`
	// The AppLovin alart query parameter. Ex: alart123
	Alart param.Opt[string] `json:"alart,omitzero"`
	// The AppLovin aleid query parameter. Ex: aleid123
	Aleid param.Opt[string] `json:"aleid,omitzero"`
	// The AppLovin Axon pixel cookie value (\_axwrt). Web-only.
	Axwrt param.Opt[string] `json:"axwrt,omitzero"`
	// The Basis DSP Click ID. Ex: basis_cid123
	BasisCid param.Opt[string] `json:"basis_cid,omitzero"`
	// The language of the browser. Ex: en-US
	BrowserLanguage param.Opt[string] `json:"browser_language,omitzero"`
	// The name of the browser. Ex: Chrome
	BrowserName param.Opt[string] `json:"browser_name,omitzero"`
	// The version of the browser. Ex: 114.0
	BrowserVersion param.Opt[string] `json:"browser_version,omitzero"`
	// The campaign id for detected in the session. This is set by the web sdk
	// automatically.
	CampaignID param.Opt[string] `json:"campaign_id,omitzero"`
	// The Click ID. Ex: clickid123
	Clickid param.Opt[string] `json:"clickid,omitzero"`
	// The Generic Click ID. Ex: clid123
	Clid param.Opt[string] `json:"clid,omitzero"`
	// The architecture of the CPU. Ex: x64
	CPUArchitecture param.Opt[string] `json:"cpu_architecture,omitzero"`
	// The full url (including query params) of the current page
	CurrentURL param.Opt[string] `json:"current_url,omitzero"`
	// The DoubleClick Click ID. Ex: dclid123
	Dclid param.Opt[string] `json:"dclid,omitzero"`
	// The model of the device. Ex: iPhone 13
	DeviceModel param.Opt[string] `json:"device_model,omitzero"`
	// The type of device the user is using. Ex: mobile
	DeviceType param.Opt[string] `json:"device_type,omitzero"`
	// The vendor of the device. Ex: Apple
	DeviceVendor param.Opt[string] `json:"device_vendor,omitzero"`
	// The time in milliseconds since the page was loaded // script was loaded
	Duration param.Opt[float64] `json:"duration,omitzero"`
	// The browsers encoding. Ex: UTF-8
	Encoding param.Opt[string] `json:"encoding,omitzero"`
	// The name of the browser engine. Ex: Blink
	EngineName param.Opt[string] `json:"engine_name,omitzero"`
	// The version of the browser engine. Ex: 114.0
	EngineVersion param.Opt[string] `json:"engine_version,omitzero"`
	// The Pinterest Click ID. Ex: epik456
	Epik param.Opt[string] `json:"epik,omitzero"`
	// Facebook Click ID with prefix format for Conversions API tracking. Ex:
	// fb.1.1554763741205.AbCdEfGhIjKlMnOpQrStUvWxYz1234567890
	Fbc param.Opt[string] `json:"fbc,omitzero"`
	// Raw Facebook Click ID query parameter without prefix from ad clicks. Ex:
	// AbCdEfGhIjKlMnOpQrStUvWxYz1234567890
	Fbclid param.Opt[string] `json:"fbclid,omitzero"`
	// Facebook Browser ID parameter for identifying browsers and attributing events.
	// Ex: fb.1.1554763741205.1098115397
	Fbp param.Opt[string] `json:"fbp,omitzero"`
	// Deprecated
	Fv param.Opt[bool] `json:"fv,omitzero"`
	// The Google Ad Source. Ex: google
	GadSource param.Opt[string] `json:"gad_source,omitzero"`
	// The Google Braid ID. Ex: gbraid123
	Gbraid param.Opt[string] `json:"gbraid,omitzero"`
	// The Google Click ID. Ex: gclid123
	Gclid param.Opt[string] `json:"gclid,omitzero"`
	// The host of the current page. Ex: example.com
	Host param.Opt[string] `json:"host,omitzero"`
	// Whether the user is in an iframe. Ex: true
	Iframe param.Opt[bool] `json:"iframe,omitzero"`
	// The Impact Click ID reference. Ex: im_ref123
	ImRef param.Opt[string] `json:"im_ref,omitzero"`
	// The IP address of the user. Ex: 127.0.0.1
	IP param.Opt[string] `json:"ip,omitzero"`
	// The Impact Click ID. Ex: irclickid123
	Irclickid param.Opt[string] `json:"irclickid,omitzero"`
	// Whether we have detected that the user is a bot. This is set automatically by
	// the Ours server primarily for events tracked through the web SDK.
	IsBot param.Opt[string] `json:"is_bot,omitzero"`
	// The LinkedIn Click ID. Ex: li_fat_id123
	LiFatID param.Opt[string] `json:"li_fat_id,omitzero"`
	// The Microsoft Click ID. Ex: msclkid123
	Msclkid param.Opt[string] `json:"msclkid,omitzero"`
	// The NextDoor Click ID. Ex: ndclid123
	Ndclid param.Opt[string] `json:"ndclid,omitzero"`
	// Deprecated
	NewS param.Opt[bool] `json:"new_s,omitzero"`
	// The name of the operating system. Ex: Windows
	OsName param.Opt[string] `json:"os_name,omitzero"`
	// The version of the operating system. Ex: 10.0
	OsVersion param.Opt[string] `json:"os_version,omitzero"`
	// A random set of numbers for the page load
	PageHash param.Opt[float64] `json:"page_hash,omitzero"`
	// The pathname of the current page. Ex: /home
	Pathname param.Opt[string] `json:"pathname,omitzero"`
	// The Quora Click ID. Ex: qclid123
	Qclid param.Opt[string] `json:"qclid,omitzero"`
	// The Reddit Click ID. Ex: rdt_cid123
	RdtCid param.Opt[string] `json:"rdt_cid,omitzero"`
	// The time the event was received by an Ours server in ISO format
	ReceivedAt param.Opt[string] `json:"received_at,omitzero"`
	// The referrer URL of the current page
	Referrer param.Opt[string] `json:"referrer,omitzero"`
	// The referring domain of the current page
	ReferringDomain param.Opt[string] `json:"referring_domain,omitzero"`
	// The StackAdapt Tracking ID. Ex: sacid123
	Sacid param.Opt[string] `json:"sacid,omitzero"`
	// The SnapChat Click ID. Ex: sccid123
	Sccid param.Opt[string] `json:"sccid,omitzero"`
	// The height of the screen. Ex: 1080
	ScreenHeight param.Opt[float64] `json:"screen_height,omitzero"`
	// The width of the screen. Ex: 1920
	ScreenWidth param.Opt[float64] `json:"screen_width,omitzero"`
	// The number of sessions the user has had. Ex: 3
	SessionCount param.Opt[float64] `json:"sessionCount,omitzero"`
	// The session ID as assigned automatically by the web SDK. This is required for
	// session replay
	Sid param.Opt[string] `json:"sid,omitzero"`
	Sr  param.Opt[string] `json:"sr,omitzero"`
	// The title of the current page
	Title param.Opt[string] `json:"title,omitzero"`
	// The TikTok Click ID. Ex: ttclid123
	Ttclid param.Opt[string] `json:"ttclid,omitzero"`
	// The Twitter Click ID. Ex: twclid123
	Twclid param.Opt[string] `json:"twclid,omitzero"`
	// User agent as a full list of strings.
	Uafvl param.Opt[string] `json:"uafvl,omitzero"`
	// The user agent of the browser
	UserAgent param.Opt[string] `json:"user_agent,omitzero"`
	// The UTM Campaign. The web SDK automatically captures this from the query params.
	UtmCampaign param.Opt[string] `json:"utm_campaign,omitzero"`
	// The UTM Content. The web SDK automatically captures this from the query params.
	UtmContent param.Opt[string] `json:"utm_content,omitzero"`
	// The UTM Medium. The web SDK automatically captures this from the query params.
	UtmMedium param.Opt[string] `json:"utm_medium,omitzero"`
	// The UTM Name. The web SDK automatically captures this from the query params.
	UtmName param.Opt[string] `json:"utm_name,omitzero"`
	// The UTM Source. The web SDK automatically captures this from the query params.
	UtmSource param.Opt[string] `json:"utm_source,omitzero"`
	// The UTM Term. The web SDK automatically captures this from the query params.
	UtmTerm param.Opt[string] `json:"utm_term,omitzero"`
	// The version of the web SDK
	Version param.Opt[string] `json:"version,omitzero"`
	// The WBRAID Identifier. The web SDK automatically captures this from the query
	// params.
	Wbraid param.Opt[string] `json:"wbraid,omitzero"`
	// Whether the user is in a webview. Ex: true
	Webview param.Opt[bool] `json:"webview,omitzero"`
	paramObj
}

func (r BatchNewParamsEventDefaultProperties) MarshalJSON() (data []byte, err error) {
	type shadow BatchNewParamsEventDefaultProperties
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchNewParamsEventDefaultProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	AdID        param.Opt[string] `json:"ad_id,omitzero"`
	AdmitadUid  param.Opt[string] `json:"admitad_uid,omitzero"`
	AdsetID     param.Opt[string] `json:"adset_id,omitzero"`
	Alart       param.Opt[string] `json:"alart,omitzero"`
	Aleid       param.Opt[string] `json:"aleid,omitzero"`
	Axwrt       param.Opt[string] `json:"axwrt,omitzero"`
	BasisCid    param.Opt[string] `json:"basis_cid,omitzero"`
	CampaignID  param.Opt[string] `json:"campaign_id,omitzero"`
	City        param.Opt[string] `json:"city,omitzero"`
	Clickid     param.Opt[string] `json:"clickid,omitzero"`
	Clid        param.Opt[string] `json:"clid,omitzero"`
	CompanyName param.Opt[string] `json:"company_name,omitzero"`
	Country     param.Opt[string] `json:"country,omitzero"`
	DateOfBirth param.Opt[string] `json:"date_of_birth,omitzero"`
	Dclid       param.Opt[string] `json:"dclid,omitzero"`
	Email       param.Opt[string] `json:"email,omitzero"`
	Epik        param.Opt[string] `json:"epik,omitzero"`
	ExternalID  param.Opt[string] `json:"external_id,omitzero"`
	Fbc         param.Opt[string] `json:"fbc,omitzero"`
	Fbclid      param.Opt[string] `json:"fbclid,omitzero"`
	Fbp         param.Opt[string] `json:"fbp,omitzero"`
	FirstName   param.Opt[string] `json:"first_name,omitzero"`
	GadSource   param.Opt[string] `json:"gad_source,omitzero"`
	Gbraid      param.Opt[string] `json:"gbraid,omitzero"`
	Gclid       param.Opt[string] `json:"gclid,omitzero"`
	Gender      param.Opt[string] `json:"gender,omitzero"`
	ImRef       param.Opt[string] `json:"im_ref,omitzero"`
	// The IP address of the user
	IP                param.Opt[string] `json:"ip,omitzero"`
	Irclickid         param.Opt[string] `json:"irclickid,omitzero"`
	IsBot             param.Opt[string] `json:"is_bot,omitzero"`
	JobTitle          param.Opt[string] `json:"job_title,omitzero"`
	LastName          param.Opt[string] `json:"last_name,omitzero"`
	LiFatID           param.Opt[string] `json:"li_fat_id,omitzero"`
	Msclkid           param.Opt[string] `json:"msclkid,omitzero"`
	Ndclid            param.Opt[string] `json:"ndclid,omitzero"`
	PhoneNumber       param.Opt[string] `json:"phone_number,omitzero"`
	Qclid             param.Opt[string] `json:"qclid,omitzero"`
	RdtCid            param.Opt[string] `json:"rdt_cid,omitzero"`
	Referrer          param.Opt[string] `json:"referrer,omitzero"`
	ReferringDomain   param.Opt[string] `json:"referring_domain,omitzero"`
	Sacid             param.Opt[string] `json:"sacid,omitzero"`
	Sccid             param.Opt[string] `json:"sccid,omitzero"`
	Sid               param.Opt[string] `json:"sid,omitzero"`
	State             param.Opt[string] `json:"state,omitzero"`
	Ttclid            param.Opt[string] `json:"ttclid,omitzero"`
	Twclid            param.Opt[string] `json:"twclid,omitzero"`
	UserAgent         param.Opt[string] `json:"user_agent,omitzero"`
	UserAgentFullList param.Opt[string] `json:"user_agent_full_list,omitzero"`
	UtmCampaign       param.Opt[string] `json:"utm_campaign,omitzero"`
	UtmContent        param.Opt[string] `json:"utm_content,omitzero"`
	UtmMedium         param.Opt[string] `json:"utm_medium,omitzero"`
	UtmName           param.Opt[string] `json:"utm_name,omitzero"`
	UtmSource         param.Opt[string] `json:"utm_source,omitzero"`
	UtmTerm           param.Opt[string] `json:"utm_term,omitzero"`
	Wbraid            param.Opt[string] `json:"wbraid,omitzero"`
	Zip               param.Opt[string] `json:"zip,omitzero"`
	Consent           map[string]string `json:"consent,omitzero"`
	CustomProperties  map[string]string `json:"custom_properties,omitzero"`
	paramObj
}

func (r BatchNewParamsEventUserProperties) MarshalJSON() (data []byte, err error) {
	type shadow BatchNewParamsEventUserProperties
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchNewParamsEventUserProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
