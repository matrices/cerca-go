// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cercago

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/matrices/cerca-go/internal/apijson"
	"github.com/matrices/cerca-go/internal/apiquery"
	"github.com/matrices/cerca-go/internal/param"
	"github.com/matrices/cerca-go/internal/requestconfig"
	"github.com/matrices/cerca-go/option"
	"github.com/tidwall/gjson"
)

// CellService contains methods and other services that help with interacting with
// the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellService] method instead.
type CellService struct {
	Options        []option.RequestOption
	Config         *CellConfigService
	Metadata       *CellMetadataService
	Tools          *CellToolService
	ToolSources    *CellToolSourceService
	Context        *CellContextService
	Connections    *CellConnectionService
	Schedules      *CellScheduleService
	Threads        *CellThreadService
	Approvals      *CellApprovalService
	ApprovalGrants *CellApprovalGrantService
	Logs           *CellLogService
	Sandbox        *CellSandboxService
	Events         *CellEventService
}

// NewCellService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCellService(opts ...option.RequestOption) (r *CellService) {
	r = &CellService{}
	r.Options = opts
	r.Config = NewCellConfigService(opts...)
	r.Metadata = NewCellMetadataService(opts...)
	r.Tools = NewCellToolService(opts...)
	r.ToolSources = NewCellToolSourceService(opts...)
	r.Context = NewCellContextService(opts...)
	r.Connections = NewCellConnectionService(opts...)
	r.Schedules = NewCellScheduleService(opts...)
	r.Threads = NewCellThreadService(opts...)
	r.Approvals = NewCellApprovalService(opts...)
	r.ApprovalGrants = NewCellApprovalGrantService(opts...)
	r.Logs = NewCellLogService(opts...)
	r.Sandbox = NewCellSandboxService(opts...)
	r.Events = NewCellEventService(opts...)
	return
}

func (r *CellService) New(ctx context.Context, body CellNewParams, opts ...option.RequestOption) (res *Cell, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cells"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *CellService) Get(ctx context.Context, cellID string, opts ...option.RequestOption) (res *Cell, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *CellService) Update(ctx context.Context, cellID string, body CellUpdateParams, opts ...option.RequestOption) (res *Cell, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

func (r *CellService) List(ctx context.Context, query CellListParams, opts ...option.RequestOption) (res *CellListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cells"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *CellService) Delete(ctx context.Context, cellID string, opts ...option.RequestOption) (res *CellDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type Cell struct {
	ID                 string             `json:"id" api:"required"`
	Configuration      CellConfiguration  `json:"configuration" api:"required"`
	CreatedAt          string             `json:"createdAt" api:"required"`
	EnvironmentID      string             `json:"environmentId" api:"required"`
	ExecutionPrincipal ExecutionPrincipal `json:"executionPrincipal" api:"required,nullable"`
	// Arbitrary string metadata stored on a cell. Runtime enforces maximum key and
	// value sizes.
	Metadata  CellMetadata `json:"metadata" api:"required"`
	OrgID     string       `json:"orgId" api:"required"`
	UpdatedAt string       `json:"updatedAt" api:"required"`
	UserID    string       `json:"userId" api:"required"`
	JSON      cellJSON     `json:"-"`
}

// cellJSON contains the JSON metadata for the struct [Cell]
type cellJSON struct {
	ID                 apijson.Field
	Configuration      apijson.Field
	CreatedAt          apijson.Field
	EnvironmentID      apijson.Field
	ExecutionPrincipal apijson.Field
	Metadata           apijson.Field
	OrgID              apijson.Field
	UpdatedAt          apijson.Field
	UserID             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Cell) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellJSON) RawJSON() string {
	return r.raw
}

type CellConfiguration struct {
	ApprovalRequired      []RuntimeToolName                          `json:"approvalRequired"`
	ApprovalTimeoutMs     float64                                    `json:"approvalTimeoutMs"`
	DefaultModel          string                                     `json:"defaultModel"`
	ExcludedToolSourceIDs []string                                   `json:"excludedToolSourceIds"`
	Features              []CellConfigurationFeature                 `json:"features"`
	Instructions          string                                     `json:"instructions"`
	ToolApprovalOverrides map[string]RuntimeToolApprovalOverrideMode `json:"toolApprovalOverrides"`
	ToolSourceMode        CellConfigurationToolSourceMode            `json:"toolSourceMode"`
	UserContext           string                                     `json:"userContext"`
	JSON                  cellConfigurationJSON                      `json:"-"`
}

// cellConfigurationJSON contains the JSON metadata for the struct
// [CellConfiguration]
type cellConfigurationJSON struct {
	ApprovalRequired      apijson.Field
	ApprovalTimeoutMs     apijson.Field
	DefaultModel          apijson.Field
	ExcludedToolSourceIDs apijson.Field
	Features              apijson.Field
	Instructions          apijson.Field
	ToolApprovalOverrides apijson.Field
	ToolSourceMode        apijson.Field
	UserContext           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CellConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellConfigurationJSON) RawJSON() string {
	return r.raw
}

type CellConfigurationFeature string

const (
	CellConfigurationFeatureMemory       CellConfigurationFeature = "memory"
	CellConfigurationFeatureSandbox      CellConfigurationFeature = "sandbox"
	CellConfigurationFeatureWeb          CellConfigurationFeature = "web"
	CellConfigurationFeatureConnections  CellConfigurationFeature = "connections"
	CellConfigurationFeatureOAuthConnect CellConfigurationFeature = "oauth_connect"
)

func (r CellConfigurationFeature) IsKnown() bool {
	switch r {
	case CellConfigurationFeatureMemory, CellConfigurationFeatureSandbox, CellConfigurationFeatureWeb, CellConfigurationFeatureConnections, CellConfigurationFeatureOAuthConnect:
		return true
	}
	return false
}

type CellConfigurationToolSourceMode string

const (
	CellConfigurationToolSourceModeInherit  CellConfigurationToolSourceMode = "inherit"
	CellConfigurationToolSourceModeExplicit CellConfigurationToolSourceMode = "explicit"
)

func (r CellConfigurationToolSourceMode) IsKnown() bool {
	switch r {
	case CellConfigurationToolSourceModeInherit, CellConfigurationToolSourceModeExplicit:
		return true
	}
	return false
}

type CellConfigurationParam struct {
	ApprovalRequired      param.Field[[]RuntimeToolName]                          `json:"approvalRequired"`
	ApprovalTimeoutMs     param.Field[float64]                                    `json:"approvalTimeoutMs"`
	DefaultModel          param.Field[string]                                     `json:"defaultModel"`
	ExcludedToolSourceIDs param.Field[[]string]                                   `json:"excludedToolSourceIds"`
	Features              param.Field[[]CellConfigurationFeature]                 `json:"features"`
	Instructions          param.Field[string]                                     `json:"instructions"`
	ToolApprovalOverrides param.Field[map[string]RuntimeToolApprovalOverrideMode] `json:"toolApprovalOverrides"`
	ToolSourceMode        param.Field[CellConfigurationToolSourceMode]            `json:"toolSourceMode"`
	UserContext           param.Field[string]                                     `json:"userContext"`
}

func (r CellConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CellMetadata map[string]string

type CellMetadataParam map[string]string

type CellSummary struct {
	ID                  string  `json:"id" api:"required"`
	ActiveThreadCount   float64 `json:"activeThreadCount" api:"required"`
	AwaitingThreadCount float64 `json:"awaitingThreadCount" api:"required"`
	CreatedAt           string  `json:"createdAt" api:"required"`
	DefaultModel        string  `json:"defaultModel" api:"required,nullable"`
	EnvironmentID       string  `json:"environmentId" api:"required"`
	// Arbitrary string metadata stored on a cell. Runtime enforces maximum key and
	// value sizes.
	Metadata  CellMetadata    `json:"metadata" api:"required"`
	OrgID     string          `json:"orgId" api:"required"`
	UpdatedAt string          `json:"updatedAt" api:"required"`
	UserID    string          `json:"userId" api:"required"`
	JSON      cellSummaryJSON `json:"-"`
}

// cellSummaryJSON contains the JSON metadata for the struct [CellSummary]
type cellSummaryJSON struct {
	ID                  apijson.Field
	ActiveThreadCount   apijson.Field
	AwaitingThreadCount apijson.Field
	CreatedAt           apijson.Field
	DefaultModel        apijson.Field
	EnvironmentID       apijson.Field
	Metadata            apijson.Field
	OrgID               apijson.Field
	UpdatedAt           apijson.Field
	UserID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CellSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellSummaryJSON) RawJSON() string {
	return r.raw
}

type ExecutionPrincipal struct {
	Kind   ExecutionPrincipalKind `json:"kind" api:"required"`
	EnvID  string                 `json:"envId" api:"nullable"`
	KeyID  string                 `json:"keyId"`
	UserID string                 `json:"userId"`
	JSON   executionPrincipalJSON `json:"-"`
	union  ExecutionPrincipalUnion
}

// executionPrincipalJSON contains the JSON metadata for the struct
// [ExecutionPrincipal]
type executionPrincipalJSON struct {
	Kind        apijson.Field
	EnvID       apijson.Field
	KeyID       apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r executionPrincipalJSON) RawJSON() string {
	return r.raw
}

func (r *ExecutionPrincipal) UnmarshalJSON(data []byte) (err error) {
	*r = ExecutionPrincipal{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ExecutionPrincipalUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ExecutionPrincipalObject],
// [ExecutionPrincipalObject].
func (r ExecutionPrincipal) AsUnion() ExecutionPrincipalUnion {
	return r.union
}

// Union satisfied by [ExecutionPrincipalObject] or [ExecutionPrincipalObject].
type ExecutionPrincipalUnion interface {
	implementsExecutionPrincipal()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ExecutionPrincipalUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExecutionPrincipalObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExecutionPrincipalObject{}),
		},
	)
}

type ExecutionPrincipalObject struct {
	Kind   ExecutionPrincipalObjectKind `json:"kind" api:"required"`
	UserID string                       `json:"userId" api:"required"`
	JSON   executionPrincipalObjectJSON `json:"-"`
}

// executionPrincipalObjectJSON contains the JSON metadata for the struct
// [ExecutionPrincipalObject]
type executionPrincipalObjectJSON struct {
	Kind        apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionPrincipalObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionPrincipalObjectJSON) RawJSON() string {
	return r.raw
}

func (r ExecutionPrincipalObject) implementsExecutionPrincipal() {}

type ExecutionPrincipalObjectKind string

const (
	ExecutionPrincipalObjectKindUser ExecutionPrincipalObjectKind = "user"
)

func (r ExecutionPrincipalObjectKind) IsKnown() bool {
	switch r {
	case ExecutionPrincipalObjectKindUser:
		return true
	}
	return false
}

type ExecutionPrincipalKind string

const (
	ExecutionPrincipalKindUser   ExecutionPrincipalKind = "user"
	ExecutionPrincipalKindAPIKey ExecutionPrincipalKind = "apiKey"
)

func (r ExecutionPrincipalKind) IsKnown() bool {
	switch r {
	case ExecutionPrincipalKindUser, ExecutionPrincipalKindAPIKey:
		return true
	}
	return false
}

type RuntimeToolApprovalOverrideMode string

const (
	RuntimeToolApprovalOverrideModeAlways RuntimeToolApprovalOverrideMode = "always"
	RuntimeToolApprovalOverrideModeNever  RuntimeToolApprovalOverrideMode = "never"
)

func (r RuntimeToolApprovalOverrideMode) IsKnown() bool {
	switch r {
	case RuntimeToolApprovalOverrideModeAlways, RuntimeToolApprovalOverrideModeNever:
		return true
	}
	return false
}

type CellListResponse struct {
	Cells   []CellSummary        `json:"cells" api:"required"`
	Cursor  string               `json:"cursor" api:"required,nullable"`
	HasMore bool                 `json:"hasMore" api:"required"`
	JSON    cellListResponseJSON `json:"-"`
}

// cellListResponseJSON contains the JSON metadata for the struct
// [CellListResponse]
type cellListResponseJSON struct {
	Cells       apijson.Field
	Cursor      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellListResponseJSON) RawJSON() string {
	return r.raw
}

type CellDeleteResponse struct {
	Success CellDeleteResponseSuccess `json:"success" api:"required"`
	JSON    cellDeleteResponseJSON    `json:"-"`
}

// cellDeleteResponseJSON contains the JSON metadata for the struct
// [CellDeleteResponse]
type cellDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CellDeleteResponseSuccess bool

const (
	CellDeleteResponseSuccessTrue CellDeleteResponseSuccess = true
)

func (r CellDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case CellDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type CellNewParams struct {
	UserID        param.Field[string]                 `json:"userId" api:"required"`
	Configuration param.Field[CellConfigurationParam] `json:"configuration"`
	EnvironmentID param.Field[string]                 `json:"environmentId"`
	// Arbitrary string metadata stored on a cell. Runtime enforces maximum key and
	// value sizes.
	Metadata param.Field[CellMetadataParam] `json:"metadata"`
}

func (r CellNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CellUpdateParams struct {
	Configuration param.Field[CellConfigurationParam] `json:"configuration" api:"required"`
}

func (r CellUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CellListParams struct {
	// When set to true, lists only cells with active or awaiting threads.
	Active param.Field[CellListParamsActive] `query:"active"`
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Environment to list cells from. Defaults to the API key's default environment
	// when omitted.
	EnvironmentID param.Field[string] `query:"environmentId"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [CellListParams]'s query parameters as `url.Values`.
func (r CellListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// When set to true, lists only cells with active or awaiting threads.
type CellListParamsActive string

const (
	CellListParamsActiveTrue  CellListParamsActive = "true"
	CellListParamsActiveFalse CellListParamsActive = "false"
)

func (r CellListParamsActive) IsKnown() bool {
	switch r {
	case CellListParamsActiveTrue, CellListParamsActiveFalse:
		return true
	}
	return false
}
