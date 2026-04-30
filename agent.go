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

// Agents
func (r *AgentService) New(ctx context.Context, body AgentNewParams, opts ...option.RequestOption) (res *Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Agent
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

// Agent
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

// Agents
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

// Agents
func (r *AgentService) ListAutoPaging(ctx context.Context, query AgentListParams, opts ...option.RequestOption) *pagination.AgentsCursorPageAutoPager[AgentSummary] {
	return pagination.NewAgentsCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Agent
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

// Tools
func (r *AgentService) ListTools(ctx context.Context, agentID string, opts ...option.RequestOption) (res *AgentListToolsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/tools", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Config
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

// Metadata
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
	Approvals    ApprovalPolicy `json:"approvals"`
	DefaultModel string         `json:"defaultModel"`
	Instructions string         `json:"instructions"`
	// Agent tool allowlist. These tools are subject to fleet defaults and locks, and
	// thread or turn requests may only narrow the resulting effective tools.
	Tools []shared.ToolSpec `json:"tools"`
	JSON  configurationJSON `json:"-"`
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
	Approvals    param.Field[ApprovalPolicyParam] `json:"approvals"`
	DefaultModel param.Field[string]              `json:"defaultModel"`
	Instructions param.Field[string]              `json:"instructions"`
	// Agent tool allowlist. These tools are subject to fleet defaults and locks, and
	// thread or turn requests may only narrow the resulting effective tools.
	Tools param.Field[[]shared.ToolSpecParam] `json:"tools"`
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

// An external provider or tool-source tool currently available to the agent.
type DiscoveredTool struct {
	Approval    DiscoveredToolApproval `json:"approval" api:"required"`
	Category    DiscoveredToolCategory `json:"category" api:"required"`
	Description string                 `json:"description" api:"required"`
	// JSON Schema object describing tool input parameters.
	InputSchema        map[string]interface{} `json:"inputSchema" api:"required"`
	Name               string                 `json:"name" api:"required"`
	Origin             DiscoveredToolOrigin   `json:"origin" api:"required"`
	Source             string                 `json:"source" api:"required"`
	AccountLabel       string                 `json:"accountLabel"`
	ConnectionID       string                 `json:"connectionId"`
	ConnectionMetadata ToolConnectionMetadata `json:"connectionMetadata"`
	SourceID           string                 `json:"sourceId"`
	SourceVersion      float64                `json:"sourceVersion"`
	JSON               discoveredToolJSON     `json:"-"`
}

// discoveredToolJSON contains the JSON metadata for the struct [DiscoveredTool]
type discoveredToolJSON struct {
	Approval           apijson.Field
	Category           apijson.Field
	Description        apijson.Field
	InputSchema        apijson.Field
	Name               apijson.Field
	Origin             apijson.Field
	Source             apijson.Field
	AccountLabel       apijson.Field
	ConnectionID       apijson.Field
	ConnectionMetadata apijson.Field
	SourceID           apijson.Field
	SourceVersion      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DiscoveredTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveredToolJSON) RawJSON() string {
	return r.raw
}

func (r DiscoveredTool) implementsTool() {}

type DiscoveredToolApproval string

const (
	DiscoveredToolApprovalAlways DiscoveredToolApproval = "always"
	DiscoveredToolApprovalNever  DiscoveredToolApproval = "never"
)

func (r DiscoveredToolApproval) IsKnown() bool {
	switch r {
	case DiscoveredToolApprovalAlways, DiscoveredToolApprovalNever:
		return true
	}
	return false
}

type DiscoveredToolCategory string

const (
	DiscoveredToolCategoryExternal DiscoveredToolCategory = "external"
)

func (r DiscoveredToolCategory) IsKnown() bool {
	switch r {
	case DiscoveredToolCategoryExternal:
		return true
	}
	return false
}

type DiscoveredToolOrigin string

const (
	DiscoveredToolOriginBuiltin    DiscoveredToolOrigin = "builtin"
	DiscoveredToolOriginToolSource DiscoveredToolOrigin = "tool_source"
)

func (r DiscoveredToolOrigin) IsKnown() bool {
	switch r {
	case DiscoveredToolOriginBuiltin, DiscoveredToolOriginToolSource:
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
	// Effective internal runtime tools after applying fleet defaults, fleet locks, and
	// the agent allowlist.
	Tools        []shared.ToolName          `json:"tools" api:"required"`
	Instructions string                     `json:"instructions"`
	JSON         effectiveConfigurationJSON `json:"-"`
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
		"kind",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecutionPrincipalUserExecutionPrincipal{}),
			DiscriminatorValue: "user",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecutionPrincipalAPIKeyExecutionPrincipal{}),
			DiscriminatorValue: "apiKey",
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

type SourceWarning struct {
	Message  string            `json:"message" api:"required"`
	Source   string            `json:"source" api:"required"`
	SourceID string            `json:"sourceId"`
	JSON     sourceWarningJSON `json:"-"`
}

// sourceWarningJSON contains the JSON metadata for the struct [SourceWarning]
type sourceWarningJSON struct {
	Message     apijson.Field
	Source      apijson.Field
	SourceID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SourceWarning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sourceWarningJSON) RawJSON() string {
	return r.raw
}

// A unified available tool entry. Runtime tools include requiresApproval; external
// tools include approval.
type Tool struct {
	Category    ToolCategory `json:"category" api:"required"`
	Description string       `json:"description" api:"required"`
	// This field can have the runtime type of [shared.JsonObject],
	// [map[string]interface{}].
	InputSchema        interface{}            `json:"inputSchema" api:"required"`
	Name               shared.ToolName        `json:"name" api:"required"`
	AccountLabel       string                 `json:"accountLabel"`
	Approval           ToolApproval           `json:"approval"`
	ApprovalSource     ToolApprovalSource     `json:"approvalSource"`
	ConnectionID       string                 `json:"connectionId"`
	ConnectionMetadata ToolConnectionMetadata `json:"connectionMetadata"`
	Origin             ToolOrigin             `json:"origin"`
	RequiresApproval   bool                   `json:"requiresApproval"`
	Source             string                 `json:"source"`
	SourceID           string                 `json:"sourceId"`
	SourceVersion      float64                `json:"sourceVersion"`
	JSON               toolJSON               `json:"-"`
	union              ToolUnion
}

// toolJSON contains the JSON metadata for the struct [Tool]
type toolJSON struct {
	Category           apijson.Field
	Description        apijson.Field
	InputSchema        apijson.Field
	Name               apijson.Field
	AccountLabel       apijson.Field
	Approval           apijson.Field
	ApprovalSource     apijson.Field
	ConnectionID       apijson.Field
	ConnectionMetadata apijson.Field
	Origin             apijson.Field
	RequiresApproval   apijson.Field
	Source             apijson.Field
	SourceID           apijson.Field
	SourceVersion      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r toolJSON) RawJSON() string {
	return r.raw
}

func (r *Tool) UnmarshalJSON(data []byte) (err error) {
	*r = Tool{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ToolUnion] interface which you can cast to the specific types
// for more type safety.
//
// Possible runtime types of the union are [ToolDescriptor], [DiscoveredTool].
func (r Tool) AsUnion() ToolUnion {
	return r.union
}

// A unified available tool entry. Runtime tools include requiresApproval; external
// tools include approval.
//
// Union satisfied by [ToolDescriptor] or [DiscoveredTool].
type ToolUnion interface {
	implementsTool()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ToolUnion)(nil)).Elem(),
		"category",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ToolDescriptor{}),
			DiscriminatorValue: "builtIn",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ToolDescriptor{}),
			DiscriminatorValue: "configurable",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DiscoveredTool{}),
			DiscriminatorValue: "external",
		},
	)
}

type ToolCategory string

const (
	ToolCategoryBuiltIn      ToolCategory = "builtIn"
	ToolCategoryConfigurable ToolCategory = "configurable"
	ToolCategoryExternal     ToolCategory = "external"
)

func (r ToolCategory) IsKnown() bool {
	switch r {
	case ToolCategoryBuiltIn, ToolCategoryConfigurable, ToolCategoryExternal:
		return true
	}
	return false
}

type ToolApproval string

const (
	ToolApprovalAlways ToolApproval = "always"
	ToolApprovalNever  ToolApproval = "never"
)

func (r ToolApproval) IsKnown() bool {
	switch r {
	case ToolApprovalAlways, ToolApprovalNever:
		return true
	}
	return false
}

type ToolApprovalSource string

const (
	ToolApprovalSourceAgent   ToolApprovalSource = "agent"
	ToolApprovalSourceEnv     ToolApprovalSource = "env"
	ToolApprovalSourceCatalog ToolApprovalSource = "catalog"
)

func (r ToolApprovalSource) IsKnown() bool {
	switch r {
	case ToolApprovalSourceAgent, ToolApprovalSourceEnv, ToolApprovalSourceCatalog:
		return true
	}
	return false
}

type ToolOrigin string

const (
	ToolOriginBuiltin    ToolOrigin = "builtin"
	ToolOriginToolSource ToolOrigin = "tool_source"
)

func (r ToolOrigin) IsKnown() bool {
	switch r {
	case ToolOriginBuiltin, ToolOriginToolSource:
		return true
	}
	return false
}

// A built-in or configurable runtime tool currently available to the agent.
type ToolDescriptor struct {
	ApprovalSource   ToolDescriptorApprovalSource `json:"approvalSource" api:"required"`
	Category         ToolDescriptorCategory       `json:"category" api:"required"`
	Description      string                       `json:"description" api:"required"`
	InputSchema      shared.JsonObject            `json:"inputSchema" api:"required"`
	Name             shared.ToolName              `json:"name" api:"required"`
	RequiresApproval bool                         `json:"requiresApproval" api:"required"`
	JSON             toolDescriptorJSON           `json:"-"`
}

// toolDescriptorJSON contains the JSON metadata for the struct [ToolDescriptor]
type toolDescriptorJSON struct {
	ApprovalSource   apijson.Field
	Category         apijson.Field
	Description      apijson.Field
	InputSchema      apijson.Field
	Name             apijson.Field
	RequiresApproval apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ToolDescriptor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolDescriptorJSON) RawJSON() string {
	return r.raw
}

func (r ToolDescriptor) implementsTool() {}

type ToolDescriptorApprovalSource string

const (
	ToolDescriptorApprovalSourceAgent   ToolDescriptorApprovalSource = "agent"
	ToolDescriptorApprovalSourceEnv     ToolDescriptorApprovalSource = "env"
	ToolDescriptorApprovalSourceCatalog ToolDescriptorApprovalSource = "catalog"
)

func (r ToolDescriptorApprovalSource) IsKnown() bool {
	switch r {
	case ToolDescriptorApprovalSourceAgent, ToolDescriptorApprovalSourceEnv, ToolDescriptorApprovalSourceCatalog:
		return true
	}
	return false
}

type ToolDescriptorCategory string

const (
	ToolDescriptorCategoryBuiltIn      ToolDescriptorCategory = "builtIn"
	ToolDescriptorCategoryConfigurable ToolDescriptorCategory = "configurable"
)

func (r ToolDescriptorCategory) IsKnown() bool {
	switch r {
	case ToolDescriptorCategoryBuiltIn, ToolDescriptorCategoryConfigurable:
		return true
	}
	return false
}

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

// Response for GET /agents/{agentId}/tools. The tools field is an inspection list,
// not a configuration request field.
type AgentListToolsResponse struct {
	// Unified list of runtime and external tools currently available to the agent.
	Tools          []Tool                     `json:"tools" api:"required"`
	SourceWarnings []SourceWarning            `json:"sourceWarnings"`
	JSON           agentListToolsResponseJSON `json:"-"`
}

// agentListToolsResponseJSON contains the JSON metadata for the struct
// [AgentListToolsResponse]
type agentListToolsResponseJSON struct {
	Tools          apijson.Field
	SourceWarnings apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AgentListToolsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentListToolsResponseJSON) RawJSON() string {
	return r.raw
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
