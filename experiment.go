// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/with-ours/ingest-sdk-go/internal/apijson"
	"github.com/with-ours/ingest-sdk-go/internal/requestconfig"
	"github.com/with-ours/ingest-sdk-go/option"
	"github.com/with-ours/ingest-sdk-go/packages/param"
	"github.com/with-ours/ingest-sdk-go/packages/respjson"
)

// ExperimentService contains methods and other services that help with interacting
// with the ours-privacy API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperimentService] method instead.
type ExperimentService struct {
	Options []option.RequestOption
}

// NewExperimentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExperimentService(opts ...option.RequestOption) (r ExperimentService) {
	r = ExperimentService{}
	r.Options = opts
	return
}

// Return the server-side variant assignment for a visitor in a single A/B or
// multivariate experiment, identified by its experiment key. Bucketing is
// deterministic on `visitor_id` and is sticky across calls. Tracking an impression
// is the default — pass `track_impression: false` to read without recording. The
// browser SDK and this endpoint converge on the same variant for the same visitor.
func (r *ExperimentService) Assignment(ctx context.Context, experimentKey string, body ExperimentAssignmentParams, opts ...option.RequestOption) (res *ExperimentAssignmentResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.oursprivacy.com/api/v1/")}, opts...)
	if experimentKey == "" {
		err = errors.New("missing required experiment_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("experiments/assignments/%s", experimentKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Return the active personalization assignments for a visitor. Read-only and never
// records an impression. Personalizations are populated by the event-driven rule
// engine — until that ships, this endpoint returns an empty list for every
// visitor, which is the correct fail-closed behavior (no false positives).
func (r *ExperimentService) Personalization(ctx context.Context, body ExperimentPersonalizationParams, opts ...option.RequestOption) (res *ExperimentPersonalizationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.oursprivacy.com/api/v1/")}, opts...)
	path := "experiments/personalization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// ExperimentAssignmentResponseUnion contains all possible properties and values
// from [ExperimentAssignmentResponseObject],
// [ExperimentAssignmentResponseObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ExperimentAssignmentResponseUnion struct {
	// This field is from variant [ExperimentAssignmentResponseObject].
	ExperimentID string `json:"experiment_id"`
	InExperiment bool   `json:"in_experiment"`
	Success      bool   `json:"success"`
	// This field is from variant [ExperimentAssignmentResponseObject].
	VariantID string `json:"variant_id"`
	// This field is from variant [ExperimentAssignmentResponseObject].
	ExperimentKey string `json:"experiment_key"`
	// This field is from variant [ExperimentAssignmentResponseObject].
	ExperimentName string `json:"experiment_name"`
	// This field is from variant [ExperimentAssignmentResponseObject].
	IsControl bool `json:"is_control"`
	// This field is from variant [ExperimentAssignmentResponseObject].
	Type string `json:"type"`
	// This field is from variant [ExperimentAssignmentResponseObject].
	VariantName string `json:"variant_name"`
	JSON        struct {
		ExperimentID   respjson.Field
		InExperiment   respjson.Field
		Success        respjson.Field
		VariantID      respjson.Field
		ExperimentKey  respjson.Field
		ExperimentName respjson.Field
		IsControl      respjson.Field
		Type           respjson.Field
		VariantName    respjson.Field
		raw            string
	} `json:"-"`
}

func (u ExperimentAssignmentResponseUnion) AsExperimentAssignmentResponseObject() (v ExperimentAssignmentResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExperimentAssignmentResponseUnion) AsExperimentAssignmentResponseObject2() (v ExperimentAssignmentResponseObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExperimentAssignmentResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ExperimentAssignmentResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentAssignmentResponseObject struct {
	ExperimentID string `json:"experiment_id" api:"required"`
	// Any of true.
	InExperiment bool `json:"in_experiment" api:"required"`
	// Any of true.
	Success        bool   `json:"success" api:"required"`
	VariantID      string `json:"variant_id" api:"required"`
	ExperimentKey  string `json:"experiment_key" api:"nullable"`
	ExperimentName string `json:"experiment_name" api:"nullable"`
	IsControl      bool   `json:"is_control" api:"nullable"`
	Type           string `json:"type" api:"nullable"`
	VariantName    string `json:"variant_name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExperimentID   respjson.Field
		InExperiment   respjson.Field
		Success        respjson.Field
		VariantID      respjson.Field
		ExperimentKey  respjson.Field
		ExperimentName respjson.Field
		IsControl      respjson.Field
		Type           respjson.Field
		VariantName    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentAssignmentResponseObject) RawJSON() string { return r.JSON.raw }
func (r *ExperimentAssignmentResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentAssignmentResponseObject2 struct {
	// Any of false.
	InExperiment bool `json:"in_experiment" api:"required"`
	// Any of true.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InExperiment respjson.Field
		Success      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentAssignmentResponseObject2) RawJSON() string { return r.JSON.raw }
func (r *ExperimentAssignmentResponseObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentPersonalizationResponse struct {
	Personalizations []ExperimentPersonalizationResponsePersonalization `json:"personalizations" api:"required"`
	// Any of true.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Personalizations respjson.Field
		Success          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentPersonalizationResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentPersonalizationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentPersonalizationResponsePersonalization struct {
	AssignedAt     float64 `json:"assigned_at" api:"required"`
	ExperimentID   string  `json:"experiment_id" api:"required"`
	VariantID      string  `json:"variant_id" api:"required"`
	ExperimentKey  string  `json:"experiment_key" api:"nullable"`
	ExperimentName string  `json:"experiment_name" api:"nullable"`
	VariantName    string  `json:"variant_name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssignedAt     respjson.Field
		ExperimentID   respjson.Field
		VariantID      respjson.Field
		ExperimentKey  respjson.Field
		ExperimentName respjson.Field
		VariantName    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentPersonalizationResponsePersonalization) RawJSON() string { return r.JSON.raw }
func (r *ExperimentPersonalizationResponsePersonalization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentAssignmentParams struct {
	// The experiment token (`exp_*`) for the experiment settings holding this
	// experiment. Available from the dashboard.
	Token string `json:"token" api:"required"`
	// Stable identifier for the visitor — typically the Ours visitor id from your
	// browser cookie, or your own server-side user id if you keep the same id
	// consistent across browser and server.
	VisitorID string `json:"visitor_id" api:"required"`
	// When true (default), an `$experiment_impression` event is enqueued and the
	// visitor's `experiment_assignments` map is updated. Set to false to read the
	// assignment without recording an impression — useful for in-test diagnostics.
	TrackImpression param.Opt[bool] `json:"track_impression,omitzero"`
	// Optional page context for URL + query-param eligibility. Variant bucketing is
	// deterministic on `visitor_id` regardless of context.
	Context ExperimentAssignmentParamsContext `json:"context,omitzero"`
	paramObj
}

func (r ExperimentAssignmentParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentAssignmentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentAssignmentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional page context for URL + query-param eligibility. Variant bucketing is
// deterministic on `visitor_id` regardless of context.
type ExperimentAssignmentParamsContext struct {
	// The current query string (e.g. `?utm_source=meta`). When provided, the
	// experiment's query-param conditions are evaluated for eligibility. If omitted,
	// the query string is parsed from `url`.
	Search param.Opt[string] `json:"search,omitzero"`
	// The current page URL. When provided, the experiment's URL patterns are evaluated
	// for eligibility — visitors on non-matching URLs are returned
	// `in_experiment: false`. Omit when the caller is pre-gating the request.
	URL param.Opt[string] `json:"url,omitzero"`
	paramObj
}

func (r ExperimentAssignmentParamsContext) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentAssignmentParamsContext
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentAssignmentParamsContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentPersonalizationParams struct {
	// The experiment token (`exp_*`).
	Token     string `json:"token" api:"required"`
	VisitorID string `json:"visitor_id" api:"required"`
	paramObj
}

func (r ExperimentPersonalizationParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentPersonalizationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentPersonalizationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
