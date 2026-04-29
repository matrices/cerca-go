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
	"github.com/matrices/cerca-go/packages/pagination"
	"github.com/matrices/cerca-go/shared"
	"github.com/tidwall/gjson"
)

// AgentService contains methods and other services that help with interacting with
// the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentService] method instead.
type AgentService struct {
	Options []option.RequestOption
}

// NewAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentService(opts ...option.RequestOption) (r *AgentService) {
	r = &AgentService{}
	r.Options = opts
	return
}

func (r *AgentService) New(ctx context.Context, body AgentNewParams, opts ...option.RequestOption) (res *Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *AgentService) Get(ctx context.Context, agentID string, opts ...option.RequestOption) (res *Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *AgentService) Update(ctx context.Context, agentID string, body AgentUpdateParams, opts ...option.RequestOption) (res *Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

func (r *AgentService) List(ctx context.Context, query AgentListParams, opts ...option.RequestOption) (res *pagination.AgentsCursorPage[AgentSummary], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "agents"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *AgentService) ListAutoPaging(ctx context.Context, query AgentListParams, opts ...option.RequestOption) *pagination.AgentsCursorPageAutoPager[AgentSummary] {
	return pagination.NewAgentsCursorPageAutoPager(r.List(ctx, query, opts...))
}

func (r *AgentService) Delete(ctx context.Context, agentID string, opts ...option.RequestOption) (res *AgentDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

func (r *AgentService) GetConfig(ctx context.Context, agentID string, opts ...option.RequestOption) (res *EffectiveConfiguration, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/config", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *AgentService) UpdateMetadata(ctx context.Context, agentID string, body AgentUpdateMetadataParams, opts ...option.RequestOption) (res *Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/metadata", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type Agent struct {
	ID                 string             `json:"id" api:"required"`
	Configuration      Configuration      `json:"configuration" api:"required"`
	CreatedAt          string             `json:"createdAt" api:"required"`
	ExecutionPrincipal ExecutionPrincipal `json:"executionPrincipal" api:"required,nullable"`
	FleetID            string             `json:"fleetId" api:"required"`
	// Arbitrary string metadata stored on an agent. Runtime enforces maximum key and
	// value sizes.
	Metadata  Metadata               `json:"metadata" api:"required"`
	OrgID     string                 `json:"orgId" api:"required"`
	UpdatedAt string                 `json:"updatedAt" api:"required"`
	UserID    string                 `json:"userId" api:"required"`
	Effective EffectiveConfiguration `json:"effective"`
	JSON      agentJSON              `json:"-"`
}

// agentJSON contains the JSON metadata for the struct [Agent]
type agentJSON struct {
	ID                 apijson.Field
	Configuration      apijson.Field
	CreatedAt          apijson.Field
	ExecutionPrincipal apijson.Field
	FleetID            apijson.Field
	Metadata           apijson.Field
	OrgID              apijson.Field
	UpdatedAt          apijson.Field
	UserID             apijson.Field
	Effective          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Agent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentJSON) RawJSON() string {
	return r.raw
}

type AgentSummary struct {
	ID                  string  `json:"id" api:"required"`
	ActiveThreadCount   float64 `json:"activeThreadCount" api:"required"`
	AwaitingThreadCount float64 `json:"awaitingThreadCount" api:"required"`
	CreatedAt           string  `json:"createdAt" api:"required"`
	DefaultModel        string  `json:"defaultModel" api:"required,nullable"`
	FleetID             string  `json:"fleetId" api:"required"`
	// Arbitrary string metadata stored on an agent. Runtime enforces maximum key and
	// value sizes.
	Metadata  Metadata         `json:"metadata" api:"required"`
	OrgID     string           `json:"orgId" api:"required"`
	UpdatedAt string           `json:"updatedAt" api:"required"`
	UserID    string           `json:"userId" api:"required"`
	JSON      agentSummaryJSON `json:"-"`
}

// agentSummaryJSON contains the JSON metadata for the struct [AgentSummary]
type agentSummaryJSON struct {
	ID                  apijson.Field
	ActiveThreadCount   apijson.Field
	AwaitingThreadCount apijson.Field
	CreatedAt           apijson.Field
	DefaultModel        apijson.Field
	FleetID             apijson.Field
	Metadata            apijson.Field
	OrgID               apijson.Field
	UpdatedAt           apijson.Field
	UserID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AgentSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentSummaryJSON) RawJSON() string {
	return r.raw
}

type ApprovalPolicy struct {
	TimeoutMs float64                        `json:"timeoutMs"`
	Tools     map[string]shared.ApprovalMode `json:"tools"`
	JSON      approvalPolicyJSON             `json:"-"`
}

// approvalPolicyJSON contains the JSON metadata for the struct [ApprovalPolicy]
type approvalPolicyJSON struct {
	TimeoutMs   apijson.Field
	Tools       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApprovalPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r approvalPolicyJSON) RawJSON() string {
	return r.raw
}

type ApprovalPolicyParam struct {
	TimeoutMs param.Field[float64]                        `json:"timeoutMs"`
	Tools     param.Field[map[string]shared.ApprovalMode] `json:"tools"`
}

func (r ApprovalPolicyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Configuration struct {
	Approvals    ApprovalPolicy    `json:"approvals"`
	DefaultModel string            `json:"defaultModel"`
	Instructions string            `json:"instructions"`
	Tools        []shared.ToolSpec `json:"tools"`
	JSON         configurationJSON `json:"-"`
}

// configurationJSON contains the JSON metadata for the struct [Configuration]
type configurationJSON struct {
	Approvals    apijson.Field
	DefaultModel apijson.Field
	Instructions apijson.Field
	Tools        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationJSON) RawJSON() string {
	return r.raw
}

type ConfigurationParam struct {
	Approvals    param.Field[ApprovalPolicyParam]    `json:"approvals"`
	DefaultModel param.Field[string]                 `json:"defaultModel"`
	Instructions param.Field[string]                 `json:"instructions"`
	Tools        param.Field[[]shared.ToolSpecParam] `json:"tools"`
}

func (r ConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigurationFieldName string

const (
	ConfigurationFieldNameDefaultModel ConfigurationFieldName = "defaultModel"
	ConfigurationFieldNameInstructions ConfigurationFieldName = "instructions"
	ConfigurationFieldNameTools        ConfigurationFieldName = "tools"
	ConfigurationFieldNameApprovals    ConfigurationFieldName = "approvals"
)

func (r ConfigurationFieldName) IsKnown() bool {
	switch r {
	case ConfigurationFieldNameDefaultModel, ConfigurationFieldNameInstructions, ConfigurationFieldNameTools, ConfigurationFieldNameApprovals:
		return true
	}
	return false
}

type EffectiveConfiguration struct {
	Approvals                EffectiveConfigurationApprovals `json:"approvals" api:"required"`
	ApprovalsWritableByAgent bool                            `json:"approvalsWritableByAgent" api:"required"`
	DefaultModel             string                          `json:"defaultModel" api:"required"`
	LockedByFleet            []ConfigurationFieldName        `json:"lockedByFleet" api:"required"`
	ResolvedFromFleet        []ConfigurationFieldName        `json:"resolvedFromFleet" api:"required"`
	Tools                    []shared.ToolName               `json:"tools" api:"required"`
	Instructions             string                          `json:"instructions"`
	JSON                     effectiveConfigurationJSON      `json:"-"`
}

// effectiveConfigurationJSON contains the JSON metadata for the struct
// [EffectiveConfiguration]
type effectiveConfigurationJSON struct {
	Approvals                apijson.Field
	ApprovalsWritableByAgent apijson.Field
	DefaultModel             apijson.Field
	LockedByFleet            apijson.Field
	ResolvedFromFleet        apijson.Field
	Tools                    apijson.Field
	Instructions             apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EffectiveConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r effectiveConfigurationJSON) RawJSON() string {
	return r.raw
}

type EffectiveConfigurationApprovals struct {
	TimeoutMs float64                             `json:"timeoutMs" api:"required"`
	Tools     map[string]shared.ApprovalMode      `json:"tools" api:"required"`
	JSON      effectiveConfigurationApprovalsJSON `json:"-"`
}

// effectiveConfigurationApprovalsJSON contains the JSON metadata for the struct
// [EffectiveConfigurationApprovals]
type effectiveConfigurationApprovalsJSON struct {
	TimeoutMs   apijson.Field
	Tools       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EffectiveConfigurationApprovals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r effectiveConfigurationApprovalsJSON) RawJSON() string {
	return r.raw
}

type ExecutionPrincipal struct {
	Kind    ExecutionPrincipalKind `json:"kind" api:"required"`
	FleetID string                 `json:"fleetId" api:"nullable"`
	KeyID   string                 `json:"keyId"`
	UserID  string                 `json:"userId"`
	JSON    executionPrincipalJSON `json:"-"`
	union   ExecutionPrincipalUnion
}

// executionPrincipalJSON contains the JSON metadata for the struct
// [ExecutionPrincipal]
type executionPrincipalJSON struct {
	Kind        apijson.Field
	FleetID     apijson.Field
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
// Possible runtime types of the union are
// [ExecutionPrincipalUserExecutionPrincipal],
// [ExecutionPrincipalAPIKeyExecutionPrincipal].
func (r ExecutionPrincipal) AsUnion() ExecutionPrincipalUnion {
	return r.union
}

// Union satisfied by [ExecutionPrincipalUserExecutionPrincipal] or
// [ExecutionPrincipalAPIKeyExecutionPrincipal].
type ExecutionPrincipalUnion interface {
	implementsExecutionPrincipal()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ExecutionPrincipalUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExecutionPrincipalUserExecutionPrincipal{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExecutionPrincipalAPIKeyExecutionPrincipal{}),
		},
	)
}

type ExecutionPrincipalUserExecutionPrincipal struct {
	Kind   ExecutionPrincipalUserExecutionPrincipalKind `json:"kind" api:"required"`
	UserID string                                       `json:"userId" api:"required"`
	JSON   executionPrincipalUserExecutionPrincipalJSON `json:"-"`
}

// executionPrincipalUserExecutionPrincipalJSON contains the JSON metadata for the
// struct [ExecutionPrincipalUserExecutionPrincipal]
type executionPrincipalUserExecutionPrincipalJSON struct {
	Kind        apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionPrincipalUserExecutionPrincipal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionPrincipalUserExecutionPrincipalJSON) RawJSON() string {
	return r.raw
}

func (r ExecutionPrincipalUserExecutionPrincipal) implementsExecutionPrincipal() {}

type ExecutionPrincipalUserExecutionPrincipalKind string

const (
	ExecutionPrincipalUserExecutionPrincipalKindUser ExecutionPrincipalUserExecutionPrincipalKind = "user"
)

func (r ExecutionPrincipalUserExecutionPrincipalKind) IsKnown() bool {
	switch r {
	case ExecutionPrincipalUserExecutionPrincipalKindUser:
		return true
	}
	return false
}

type ExecutionPrincipalAPIKeyExecutionPrincipal struct {
	FleetID string                                         `json:"fleetId" api:"required,nullable"`
	KeyID   string                                         `json:"keyId" api:"required"`
	Kind    ExecutionPrincipalAPIKeyExecutionPrincipalKind `json:"kind" api:"required"`
	JSON    executionPrincipalAPIKeyExecutionPrincipalJSON `json:"-"`
}

// executionPrincipalAPIKeyExecutionPrincipalJSON contains the JSON metadata for
// the struct [ExecutionPrincipalAPIKeyExecutionPrincipal]
type executionPrincipalAPIKeyExecutionPrincipalJSON struct {
	FleetID     apijson.Field
	KeyID       apijson.Field
	Kind        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionPrincipalAPIKeyExecutionPrincipal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionPrincipalAPIKeyExecutionPrincipalJSON) RawJSON() string {
	return r.raw
}

func (r ExecutionPrincipalAPIKeyExecutionPrincipal) implementsExecutionPrincipal() {}

type ExecutionPrincipalAPIKeyExecutionPrincipalKind string

const (
	ExecutionPrincipalAPIKeyExecutionPrincipalKindAPIKey ExecutionPrincipalAPIKeyExecutionPrincipalKind = "apiKey"
)

func (r ExecutionPrincipalAPIKeyExecutionPrincipalKind) IsKnown() bool {
	switch r {
	case ExecutionPrincipalAPIKeyExecutionPrincipalKindAPIKey:
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

type Metadata map[string]string

type MetadataParam map[string]string

type AgentDeleteResponse struct {
	Success AgentDeleteResponseSuccess `json:"success" api:"required"`
	JSON    agentDeleteResponseJSON    `json:"-"`
}

// agentDeleteResponseJSON contains the JSON metadata for the struct
// [AgentDeleteResponse]
type agentDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AgentDeleteResponseSuccess bool

const (
	AgentDeleteResponseSuccessTrue AgentDeleteResponseSuccess = true
)

func (r AgentDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AgentDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AgentNewParams struct {
	Configuration param.Field[ConfigurationParam] `json:"configuration"`
	FleetID       param.Field[string]             `json:"fleetId"`
	// Arbitrary string metadata stored on an agent. Runtime enforces maximum key and
	// value sizes.
	Metadata param.Field[MetadataParam] `json:"metadata"`
	UserID   param.Field[string]        `json:"userId"`
}

func (r AgentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentUpdateParams struct {
	Configuration param.Field[ConfigurationParam] `json:"configuration" api:"required"`
}

func (r AgentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentListParams struct {
	// When set to true, lists only agents with active or awaiting threads.
	Active param.Field[AgentListParamsActive] `query:"active"`
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Fleet to list agents from. Defaults to the API key's default fleet when omitted.
	FleetID param.Field[string] `query:"fleetId"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [AgentListParams]'s query parameters as `url.Values`.
func (r AgentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// When set to true, lists only agents with active or awaiting threads.
type AgentListParamsActive string

const (
	AgentListParamsActiveTrue  AgentListParamsActive = "true"
	AgentListParamsActiveFalse AgentListParamsActive = "false"
)

func (r AgentListParamsActive) IsKnown() bool {
	switch r {
	case AgentListParamsActiveTrue, AgentListParamsActiveFalse:
		return true
	}
	return false
}

type AgentUpdateMetadataParams struct {
	// Arbitrary string metadata stored on an agent. Runtime enforces maximum key and
	// value sizes.
	Metadata param.Field[MetadataParam] `json:"metadata" api:"required"`
}

func (r AgentUpdateMetadataParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
