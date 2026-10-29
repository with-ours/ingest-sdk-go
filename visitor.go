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

// VisitorService contains methods and other services that help with interacting
// with the ours-privacy API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVisitorService] method instead.
type VisitorService struct {
	Options []option.RequestOption
}

// NewVisitorService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVisitorService(opts ...option.RequestOption) (r VisitorService) {
	r = VisitorService{}
	r.Options = opts
	return
}

// Set or update properties on an existing visitor, or create a new visitor if no
// match is found. This fires a $identify event, making the call visible in the
// event stream. Identity resolution runs in priority order: userId (direct, no
// lookup) → externalId → email → userProperties.phone_number. Each lookup runs
// only if earlier identifiers did not resolve a visitor. Phone matching ignores
// common formatting and an international + or 00 prefix, but does not infer a
// country code. When a visitor is found, their Ours Visitor ID is used going
// forward so all future events are attached to the same profile. For top-level
// visitor properties: null clears the existing value, while undefined, omitted
// fields, and empty strings are ignored. For entries inside custom_properties:
// null, undefined, and empty strings are all ignored (custom_properties use merge
// semantics). See https://docs.oursprivacy.com/docs/data-types for details and
// common pitfalls.
func (r *VisitorService) Upsert(ctx context.Context, body VisitorUpsertParams, opts ...option.RequestOption) (res *VisitorUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.oursprivacy.com/api/v1/")}, opts...)
	path := "identify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type VisitorUpsertResponse struct {
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
func (r VisitorUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *VisitorUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VisitorUpsertParams struct {
	// The token for your Source. You can find this in the dashboard. Leading and
	// trailing whitespace is removed before validation.
	Token string `json:"token" api:"required"`
	// User properties to associate with this user. The existing user properties will
	// be updated. And all future events will have these properties associated with
	// them.
	UserProperties VisitorUpsertParamsUserProperties `json:"userProperties,omitzero" api:"required"`
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
	// The Ours Visitor ID stored in local storage and cookies on your web properties.
	// When present, this is used directly — no lookup by externalId, email, or phone
	// is performed. If you have both a userId and an externalId, send both so the
	// event is attached to the right visitor without any lookup overhead. Leading and
	// trailing whitespace is removed before validation.
	UserID param.Opt[string] `json:"userId,omitzero"`
	// These properties are used throughout the Ours app to pass known values onto
	// destinations
	DefaultProperties VisitorUpsertParamsDefaultProperties `json:"defaultProperties,omitzero"`
	// End-user network context for server-side calls. Required for probabilistic
	// identity resolution when the caller is a backend server rather than an end-user
	// browser.
	IdentityContext VisitorUpsertParamsIdentityContext `json:"identityContext,omitzero"`
	paramObj
}

func (r VisitorUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow VisitorUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VisitorUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User properties to associate with this user. The existing user properties will
// be updated. And all future events will have these properties associated with
// them.
type VisitorUpsertParamsUserProperties struct {
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
	Consent map[string]VisitorUpsertParamsUserPropertiesConsentUnion `json:"consent,omitzero"`
	// Use at most 200 nonempty property keys and five object levels. Array contents do
	// not count toward the object-depth limit. Empty keys are removed recursively,
	// arrays are truncated to 250 entries, and top-level numbers outside the safe
	// integer range become strings. After normalization, the serialized JSON must be
	// shorter than 1,048,576 UTF-16 code units.
	CustomProperties map[string]VisitorUpsertParamsUserPropertiesCustomPropertyUnion `json:"custom_properties,omitzero"`
	// Any JSON value: string, number, boolean, array, object, or null.
	IsBot VisitorUpsertParamsUserPropertiesIsBotUnion `json:"is_bot,omitzero"`
	// Any JSON value: string, number, boolean, array, object, or null.
	PhoneNumber VisitorUpsertParamsUserPropertiesPhoneNumberUnion `json:"phone_number,omitzero"`
	// Any JSON value: string, number, boolean, array, object, or null.
	Zip VisitorUpsertParamsUserPropertiesZipUnion `json:"zip,omitzero"`
	paramObj
}

func (r VisitorUpsertParamsUserProperties) MarshalJSON() (data []byte, err error) {
	type shadow VisitorUpsertParamsUserProperties
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VisitorUpsertParamsUserProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VisitorUpsertParamsUserPropertiesConsentUnion struct {
	OfString                                    param.Opt[string]                                               `json:",omitzero,inline"`
	OfFloat                                     param.Opt[float64]                                              `json:",omitzero,inline"`
	OfBool                                      param.Opt[bool]                                                 `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesConsentArray  []VisitorUpsertParamsUserPropertiesConsentArrayItemUnion        `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesConsentMapMap map[string]VisitorUpsertParamsUserPropertiesConsentMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesConsentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfVisitorUpsertsUserPropertiesConsentArray,
		u.OfVisitorUpsertsUserPropertiesConsentMapMap)
}
func (u *VisitorUpsertParamsUserPropertiesConsentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesConsentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesConsentArray) {
		return &u.OfVisitorUpsertsUserPropertiesConsentArray
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesConsentMapMap) {
		return &u.OfVisitorUpsertsUserPropertiesConsentMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VisitorUpsertParamsUserPropertiesConsentArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesConsentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *VisitorUpsertParamsUserPropertiesConsentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesConsentArrayItemUnion) asAny() any {
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
type VisitorUpsertParamsUserPropertiesConsentMapItemUnion struct {
	OfString                                           param.Opt[string]                                                      `json:",omitzero,inline"`
	OfFloat                                            param.Opt[float64]                                                     `json:",omitzero,inline"`
	OfBool                                             param.Opt[bool]                                                        `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesConsentMapItemArray  []VisitorUpsertParamsUserPropertiesConsentMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesConsentMapItemMapMap map[string]VisitorUpsertParamsUserPropertiesConsentMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesConsentMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfVisitorUpsertsUserPropertiesConsentMapItemArray,
		u.OfVisitorUpsertsUserPropertiesConsentMapItemMapMap)
}
func (u *VisitorUpsertParamsUserPropertiesConsentMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesConsentMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesConsentMapItemArray) {
		return &u.OfVisitorUpsertsUserPropertiesConsentMapItemArray
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesConsentMapItemMapMap) {
		return &u.OfVisitorUpsertsUserPropertiesConsentMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VisitorUpsertParamsUserPropertiesConsentMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesConsentMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *VisitorUpsertParamsUserPropertiesConsentMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesConsentMapItemArrayItemUnion) asAny() any {
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
type VisitorUpsertParamsUserPropertiesConsentMapItemMapItemUnion struct {
	OfString                                                  param.Opt[string]                                                             `json:",omitzero,inline"`
	OfFloat                                                   param.Opt[float64]                                                            `json:",omitzero,inline"`
	OfBool                                                    param.Opt[bool]                                                               `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesConsentMapItemMapItemArray  []VisitorUpsertParamsUserPropertiesConsentMapItemMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesConsentMapItemMapItemMapMap map[string]VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesConsentMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfVisitorUpsertsUserPropertiesConsentMapItemMapItemArray,
		u.OfVisitorUpsertsUserPropertiesConsentMapItemMapItemMapMap)
}
func (u *VisitorUpsertParamsUserPropertiesConsentMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesConsentMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesConsentMapItemMapItemArray) {
		return &u.OfVisitorUpsertsUserPropertiesConsentMapItemMapItemArray
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesConsentMapItemMapItemMapMap) {
		return &u.OfVisitorUpsertsUserPropertiesConsentMapItemMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VisitorUpsertParamsUserPropertiesConsentMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesConsentMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *VisitorUpsertParamsUserPropertiesConsentMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesConsentMapItemMapItemArrayItemUnion) asAny() any {
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
type VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemUnion struct {
	OfString                                                         param.Opt[string]                                                                    `json:",omitzero,inline"`
	OfFloat                                                          param.Opt[float64]                                                                   `json:",omitzero,inline"`
	OfBool                                                           param.Opt[bool]                                                                      `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesConsentMapItemMapItemMapItemArray  []VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesConsentMapItemMapItemMapItemMapMap map[string]VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfVisitorUpsertsUserPropertiesConsentMapItemMapItemMapItemArray,
		u.OfVisitorUpsertsUserPropertiesConsentMapItemMapItemMapItemMapMap)
}
func (u *VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesConsentMapItemMapItemMapItemArray) {
		return &u.OfVisitorUpsertsUserPropertiesConsentMapItemMapItemMapItemArray
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesConsentMapItemMapItemMapItemMapMap) {
		return &u.OfVisitorUpsertsUserPropertiesConsentMapItemMapItemMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemArrayItemUnion) asAny() any {
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
type VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemMapItemUnion struct {
	OfString                                                               param.Opt[string]                                                                    `json:",omitzero,inline"`
	OfFloat                                                                param.Opt[float64]                                                                   `json:",omitzero,inline"`
	OfBool                                                                 param.Opt[bool]                                                                      `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesConsentMapItemMapItemMapItemMapItemArray []VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemMapItemArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfVisitorUpsertsUserPropertiesConsentMapItemMapItemMapItemMapItemArray)
}
func (u *VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesConsentMapItemMapItemMapItemMapItemArray) {
		return &u.OfVisitorUpsertsUserPropertiesConsentMapItemMapItemMapItemMapItemArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesConsentMapItemMapItemMapItemMapItemArrayItemUnion) asAny() any {
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
type VisitorUpsertParamsUserPropertiesCustomPropertyUnion struct {
	OfString                                           param.Opt[string]                                                      `json:",omitzero,inline"`
	OfFloat                                            param.Opt[float64]                                                     `json:",omitzero,inline"`
	OfBool                                             param.Opt[bool]                                                        `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesCustomPropertyArray  []VisitorUpsertParamsUserPropertiesCustomPropertyArrayItemUnion        `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesCustomPropertyMapMap map[string]VisitorUpsertParamsUserPropertiesCustomPropertyMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesCustomPropertyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfVisitorUpsertsUserPropertiesCustomPropertyArray,
		u.OfVisitorUpsertsUserPropertiesCustomPropertyMapMap)
}
func (u *VisitorUpsertParamsUserPropertiesCustomPropertyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesCustomPropertyUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesCustomPropertyArray) {
		return &u.OfVisitorUpsertsUserPropertiesCustomPropertyArray
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesCustomPropertyMapMap) {
		return &u.OfVisitorUpsertsUserPropertiesCustomPropertyMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VisitorUpsertParamsUserPropertiesCustomPropertyArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesCustomPropertyArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *VisitorUpsertParamsUserPropertiesCustomPropertyArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesCustomPropertyArrayItemUnion) asAny() any {
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
type VisitorUpsertParamsUserPropertiesCustomPropertyMapItemUnion struct {
	OfString                                                  param.Opt[string]                                                             `json:",omitzero,inline"`
	OfFloat                                                   param.Opt[float64]                                                            `json:",omitzero,inline"`
	OfBool                                                    param.Opt[bool]                                                               `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesCustomPropertyMapItemArray  []VisitorUpsertParamsUserPropertiesCustomPropertyMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapMap map[string]VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesCustomPropertyMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemArray,
		u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapMap)
}
func (u *VisitorUpsertParamsUserPropertiesCustomPropertyMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesCustomPropertyMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemArray) {
		return &u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemArray
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapMap) {
		return &u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VisitorUpsertParamsUserPropertiesCustomPropertyMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesCustomPropertyMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *VisitorUpsertParamsUserPropertiesCustomPropertyMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesCustomPropertyMapItemArrayItemUnion) asAny() any {
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
type VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemUnion struct {
	OfString                                                         param.Opt[string]                                                                    `json:",omitzero,inline"`
	OfFloat                                                          param.Opt[float64]                                                                   `json:",omitzero,inline"`
	OfBool                                                           param.Opt[bool]                                                                      `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemArray  []VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemMapMap map[string]VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemArray,
		u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemMapMap)
}
func (u *VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemArray) {
		return &u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemArray
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemMapMap) {
		return &u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemArrayItemUnion) asAny() any {
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
type VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemUnion struct {
	OfString                                                                param.Opt[string]                                                                           `json:",omitzero,inline"`
	OfFloat                                                                 param.Opt[float64]                                                                          `json:",omitzero,inline"`
	OfBool                                                                  param.Opt[bool]                                                                             `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemMapItemArray  []VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemMapItemMapMap map[string]VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemMapItemArray,
		u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemMapItemMapMap)
}
func (u *VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemMapItemArray) {
		return &u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemMapItemArray
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemMapItemMapMap) {
		return &u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemArrayItemUnion) asAny() any {
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
type VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemUnion struct {
	OfString                                                                      param.Opt[string]                                                                           `json:",omitzero,inline"`
	OfFloat                                                                       param.Opt[float64]                                                                          `json:",omitzero,inline"`
	OfBool                                                                        param.Opt[bool]                                                                             `json:",omitzero,inline"`
	OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArray []VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArray)
}
func (u *VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArray) {
		return &u.OfVisitorUpsertsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesCustomPropertyMapItemMapItemMapItemMapItemArrayItemUnion) asAny() any {
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
type VisitorUpsertParamsUserPropertiesIsBotUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesIsBotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *VisitorUpsertParamsUserPropertiesIsBotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesIsBotUnion) asAny() any {
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
type VisitorUpsertParamsUserPropertiesPhoneNumberUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesPhoneNumberUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *VisitorUpsertParamsUserPropertiesPhoneNumberUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesPhoneNumberUnion) asAny() any {
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
type VisitorUpsertParamsUserPropertiesZipUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsUserPropertiesZipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *VisitorUpsertParamsUserPropertiesZipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsUserPropertiesZipUnion) asAny() any {
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

// These properties are used throughout the Ours app to pass known values onto
// destinations
type VisitorUpsertParamsDefaultProperties struct {
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
	IsBot VisitorUpsertParamsDefaultPropertiesIsBotUnion `json:"is_bot,omitzero"`
	paramObj
}

func (r VisitorUpsertParamsDefaultProperties) MarshalJSON() (data []byte, err error) {
	type shadow VisitorUpsertParamsDefaultProperties
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VisitorUpsertParamsDefaultProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VisitorUpsertParamsDefaultPropertiesIsBotUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u VisitorUpsertParamsDefaultPropertiesIsBotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *VisitorUpsertParamsDefaultPropertiesIsBotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VisitorUpsertParamsDefaultPropertiesIsBotUnion) asAny() any {
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
type VisitorUpsertParamsIdentityContext struct {
	// The end-user IP address (not the server IP).
	IP string `json:"ip" api:"required"`
	// The end-user User-Agent string (not the server UA).
	UserAgent string `json:"userAgent" api:"required"`
	paramObj
}

func (r VisitorUpsertParamsIdentityContext) MarshalJSON() (data []byte, err error) {
	type shadow VisitorUpsertParamsIdentityContext
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VisitorUpsertParamsIdentityContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
