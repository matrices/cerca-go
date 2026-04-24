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
	"github.com/tidwall/gjson"
)

// ToolSourceService contains methods and other services that help with interacting
// with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolSourceService] method instead.
type ToolSourceService struct {
	Options []option.RequestOption
}

// NewToolSourceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewToolSourceService(opts ...option.RequestOption) (r *ToolSourceService) {
	r = &ToolSourceService{}
	r.Options = opts
	return
}

func (r *ToolSourceService) New(ctx context.Context, fleetID string, body ToolSourceNewParams, opts ...option.RequestOption) (res *ToolSource, err error) {
	opts = slices.Concat(r.Options, opts)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("fleets/%s/tool-sources", fleetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ToolSourceService) Get(ctx context.Context, fleetID string, sourceID string, opts ...option.RequestOption) (res *ToolSource, err error) {
	opts = slices.Concat(r.Options, opts)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return nil, err
	}
	if sourceID == "" {
		err = errors.New("missing required sourceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("fleets/%s/tool-sources/%s", fleetID, sourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *ToolSourceService) Update(ctx context.Context, fleetID string, sourceID string, body ToolSourceUpdateParams, opts ...option.RequestOption) (res *ToolSource, err error) {
	opts = slices.Concat(r.Options, opts)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return nil, err
	}
	if sourceID == "" {
		err = errors.New("missing required sourceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("fleets/%s/tool-sources/%s", fleetID, sourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

func (r *ToolSourceService) List(ctx context.Context, fleetID string, query ToolSourceListParams, opts ...option.RequestOption) (res *pagination.SourcesCursorPage[ToolSource], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("fleets/%s/tool-sources", fleetID)
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

func (r *ToolSourceService) ListAutoPaging(ctx context.Context, fleetID string, query ToolSourceListParams, opts ...option.RequestOption) *pagination.SourcesCursorPageAutoPager[ToolSource] {
	return pagination.NewSourcesCursorPageAutoPager(r.List(ctx, fleetID, query, opts...))
}

func (r *ToolSourceService) Delete(ctx context.Context, fleetID string, sourceID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return err
	}
	if sourceID == "" {
		err = errors.New("missing required sourceId parameter")
		return err
	}
	path := fmt.Sprintf("fleets/%s/tool-sources/%s", fleetID, sourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *ToolSourceService) ListForAgent(ctx context.Context, agentID string, opts ...option.RequestOption) (res *ToolSourceListForAgentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/tool-sources", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type HTTPToolDefinition map[string]interface{}

type HTTPToolDefinitionParam map[string]interface{}

type HTTPToolExecutionPolicy map[string]interface{}

type HTTPToolExecutionPolicyParam map[string]interface{}

type HTTPToolSource struct {
	ID string `json:"id" api:"required"`
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      ToolSourceAuth       `json:"auth" api:"required"`
	CreatedAt string               `json:"createdAt" api:"required"`
	Enabled   bool                 `json:"enabled" api:"required"`
	FleetID   string               `json:"fleetId" api:"required"`
	Namespace string               `json:"namespace" api:"required"`
	Tools     []HTTPToolDefinition `json:"tools" api:"required"`
	Type      HTTPToolSourceType   `json:"type" api:"required"`
	UpdatedAt string               `json:"updatedAt" api:"required"`
	Version   float64              `json:"version" api:"required"`
	Approval  ToolApprovalMode     `json:"approval"`
	// HTTP tool execution retry and timeout policy.
	Execution HTTPToolExecutionPolicy `json:"execution"`
	JSON      httpToolSourceJSON      `json:"-"`
}

// httpToolSourceJSON contains the JSON metadata for the struct [HTTPToolSource]
type httpToolSourceJSON struct {
	ID          apijson.Field
	Auth        apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	FleetID     apijson.Field
	Namespace   apijson.Field
	Tools       apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Version     apijson.Field
	Approval    apijson.Field
	Execution   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPToolSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpToolSourceJSON) RawJSON() string {
	return r.raw
}

func (r HTTPToolSource) implementsToolSource() {}

type HTTPToolSourceType string

const (
	HTTPToolSourceTypeHTTP HTTPToolSourceType = "http"
)

func (r HTTPToolSourceType) IsKnown() bool {
	switch r {
	case HTTPToolSourceTypeHTTP:
		return true
	}
	return false
}

type McpToolExecutionPolicy map[string]interface{}

type McpToolExecutionPolicyParam map[string]interface{}

type McpToolSource struct {
	ID string `json:"id" api:"required"`
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      ToolSourceAuth    `json:"auth" api:"required"`
	CreatedAt string            `json:"createdAt" api:"required"`
	Enabled   bool              `json:"enabled" api:"required"`
	FleetID   string            `json:"fleetId" api:"required"`
	Namespace string            `json:"namespace" api:"required"`
	Type      McpToolSourceType `json:"type" api:"required"`
	UpdatedAt string            `json:"updatedAt" api:"required"`
	URL       string            `json:"url" api:"required"`
	Version   float64           `json:"version" api:"required"`
	Approval  ToolApprovalMode  `json:"approval"`
	// MCP discovery and execution retry/timeout policy.
	Execution McpToolExecutionPolicy `json:"execution"`
	JSON      mcpToolSourceJSON      `json:"-"`
}

// mcpToolSourceJSON contains the JSON metadata for the struct [McpToolSource]
type mcpToolSourceJSON struct {
	ID          apijson.Field
	Auth        apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	FleetID     apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	Version     apijson.Field
	Approval    apijson.Field
	Execution   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpToolSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpToolSourceJSON) RawJSON() string {
	return r.raw
}

func (r McpToolSource) implementsToolSource() {}

type McpToolSourceType string

const (
	McpToolSourceTypeMcp McpToolSourceType = "mcp"
)

func (r McpToolSourceType) IsKnown() bool {
	switch r {
	case McpToolSourceTypeMcp:
		return true
	}
	return false
}

type ToolApprovalMode string

const (
	ToolApprovalModeAlways ToolApprovalMode = "always"
	ToolApprovalModeNever  ToolApprovalMode = "never"
)

func (r ToolApprovalMode) IsKnown() bool {
	switch r {
	case ToolApprovalModeAlways, ToolApprovalModeNever:
		return true
	}
	return false
}

type ToolSource struct {
	ID string `json:"id" api:"required"`
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      ToolSourceAuth   `json:"auth" api:"required"`
	CreatedAt string           `json:"createdAt" api:"required"`
	Enabled   bool             `json:"enabled" api:"required"`
	FleetID   string           `json:"fleetId" api:"required"`
	Namespace string           `json:"namespace" api:"required"`
	Type      ToolSourceType   `json:"type" api:"required"`
	UpdatedAt string           `json:"updatedAt" api:"required"`
	Version   float64          `json:"version" api:"required"`
	Approval  ToolApprovalMode `json:"approval"`
	// This field can have the runtime type of [HTTPToolExecutionPolicy],
	// [McpToolExecutionPolicy].
	Execution interface{} `json:"execution"`
	// This field can have the runtime type of [[]HTTPToolDefinition].
	Tools interface{}    `json:"tools"`
	URL   string         `json:"url"`
	JSON  toolSourceJSON `json:"-"`
	union ToolSourceUnion
}

// toolSourceJSON contains the JSON metadata for the struct [ToolSource]
type toolSourceJSON struct {
	ID          apijson.Field
	Auth        apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	FleetID     apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Version     apijson.Field
	Approval    apijson.Field
	Execution   apijson.Field
	Tools       apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r toolSourceJSON) RawJSON() string {
	return r.raw
}

func (r *ToolSource) UnmarshalJSON(data []byte) (err error) {
	*r = ToolSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ToolSourceUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [HTTPToolSource], [McpToolSource].
func (r ToolSource) AsUnion() ToolSourceUnion {
	return r.union
}

// Union satisfied by [HTTPToolSource] or [McpToolSource].
type ToolSourceUnion interface {
	implementsToolSource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ToolSourceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(HTTPToolSource{}),
			DiscriminatorValue: "http",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(McpToolSource{}),
			DiscriminatorValue: "mcp",
		},
	)
}

type ToolSourceType string

const (
	ToolSourceTypeHTTP ToolSourceType = "http"
	ToolSourceTypeMcp  ToolSourceType = "mcp"
)

func (r ToolSourceType) IsKnown() bool {
	switch r {
	case ToolSourceTypeHTTP, ToolSourceTypeMcp:
		return true
	}
	return false
}

type ToolSourceAuth map[string]interface{}

type ToolSourceAuthParam map[string]interface{}

type ToolSourceListForAgentResponse struct {
	Cursor  string                             `json:"cursor" api:"required,nullable"`
	HasMore bool                               `json:"hasMore" api:"required"`
	Sources []ToolSource                       `json:"sources" api:"required"`
	JSON    toolSourceListForAgentResponseJSON `json:"-"`
}

// toolSourceListForAgentResponseJSON contains the JSON metadata for the struct
// [ToolSourceListForAgentResponse]
type toolSourceListForAgentResponseJSON struct {
	Cursor      apijson.Field
	HasMore     apijson.Field
	Sources     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSourceListForAgentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSourceListForAgentResponseJSON) RawJSON() string {
	return r.raw
}

type ToolSourceNewParams struct {
	Source ToolSourceNewParamsSourceUnion `json:"source" api:"required"`
}

func (r ToolSourceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Source)
}

type ToolSourceNewParamsSource struct {
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      param.Field[ToolSourceAuthParam]           `json:"auth" api:"required"`
	Namespace param.Field[string]                        `json:"namespace" api:"required"`
	Type      param.Field[ToolSourceNewParamsSourceType] `json:"type" api:"required"`
	Approval  param.Field[ToolApprovalMode]              `json:"approval"`
	Enabled   param.Field[bool]                          `json:"enabled"`
	Execution param.Field[interface{}]                   `json:"execution"`
	Tools     param.Field[interface{}]                   `json:"tools"`
	URL       param.Field[string]                        `json:"url"`
}

func (r ToolSourceNewParamsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolSourceNewParamsSource) implementsToolSourceNewParamsSourceUnion() {}

// Satisfied by [ToolSourceNewParamsSourceCreateHTTPToolSourceRequest],
// [ToolSourceNewParamsSourceCreateMcpToolSourceRequest],
// [ToolSourceNewParamsSource].
type ToolSourceNewParamsSourceUnion interface {
	implementsToolSourceNewParamsSourceUnion()
}

type ToolSourceNewParamsSourceCreateHTTPToolSourceRequest struct {
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      param.Field[ToolSourceAuthParam]                                      `json:"auth" api:"required"`
	Namespace param.Field[string]                                                   `json:"namespace" api:"required"`
	Tools     param.Field[[]HTTPToolDefinitionParam]                                `json:"tools" api:"required"`
	Type      param.Field[ToolSourceNewParamsSourceCreateHTTPToolSourceRequestType] `json:"type" api:"required"`
	Approval  param.Field[ToolApprovalMode]                                         `json:"approval"`
	Enabled   param.Field[bool]                                                     `json:"enabled"`
	// HTTP tool execution retry and timeout policy.
	Execution   param.Field[HTTPToolExecutionPolicyParam] `json:"execution"`
	ExtraFields map[string]interface{}                    `json:"-,extras"`
}

func (r ToolSourceNewParamsSourceCreateHTTPToolSourceRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolSourceNewParamsSourceCreateHTTPToolSourceRequest) implementsToolSourceNewParamsSourceUnion() {
}

type ToolSourceNewParamsSourceCreateHTTPToolSourceRequestType string

const (
	ToolSourceNewParamsSourceCreateHTTPToolSourceRequestTypeHTTP ToolSourceNewParamsSourceCreateHTTPToolSourceRequestType = "http"
)

func (r ToolSourceNewParamsSourceCreateHTTPToolSourceRequestType) IsKnown() bool {
	switch r {
	case ToolSourceNewParamsSourceCreateHTTPToolSourceRequestTypeHTTP:
		return true
	}
	return false
}

type ToolSourceNewParamsSourceCreateMcpToolSourceRequest struct {
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      param.Field[ToolSourceAuthParam]                                     `json:"auth" api:"required"`
	Namespace param.Field[string]                                                  `json:"namespace" api:"required"`
	Type      param.Field[ToolSourceNewParamsSourceCreateMcpToolSourceRequestType] `json:"type" api:"required"`
	URL       param.Field[string]                                                  `json:"url" api:"required"`
	Approval  param.Field[ToolApprovalMode]                                        `json:"approval"`
	Enabled   param.Field[bool]                                                    `json:"enabled"`
	// MCP discovery and execution retry/timeout policy.
	Execution   param.Field[McpToolExecutionPolicyParam] `json:"execution"`
	ExtraFields map[string]interface{}                   `json:"-,extras"`
}

func (r ToolSourceNewParamsSourceCreateMcpToolSourceRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolSourceNewParamsSourceCreateMcpToolSourceRequest) implementsToolSourceNewParamsSourceUnion() {
}

type ToolSourceNewParamsSourceCreateMcpToolSourceRequestType string

const (
	ToolSourceNewParamsSourceCreateMcpToolSourceRequestTypeMcp ToolSourceNewParamsSourceCreateMcpToolSourceRequestType = "mcp"
)

func (r ToolSourceNewParamsSourceCreateMcpToolSourceRequestType) IsKnown() bool {
	switch r {
	case ToolSourceNewParamsSourceCreateMcpToolSourceRequestTypeMcp:
		return true
	}
	return false
}

type ToolSourceNewParamsSourceType string

const (
	ToolSourceNewParamsSourceTypeHTTP ToolSourceNewParamsSourceType = "http"
	ToolSourceNewParamsSourceTypeMcp  ToolSourceNewParamsSourceType = "mcp"
)

func (r ToolSourceNewParamsSourceType) IsKnown() bool {
	switch r {
	case ToolSourceNewParamsSourceTypeHTTP, ToolSourceNewParamsSourceTypeMcp:
		return true
	}
	return false
}

type ToolSourceUpdateParams struct {
	Source ToolSourceUpdateParamsSourceUnion `json:"source" api:"required"`
}

func (r ToolSourceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Source)
}

type ToolSourceUpdateParamsSource struct {
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      param.Field[ToolSourceAuthParam]              `json:"auth" api:"required"`
	Namespace param.Field[string]                           `json:"namespace" api:"required"`
	Type      param.Field[ToolSourceUpdateParamsSourceType] `json:"type" api:"required"`
	Approval  param.Field[ToolApprovalMode]                 `json:"approval"`
	Enabled   param.Field[bool]                             `json:"enabled"`
	Execution param.Field[interface{}]                      `json:"execution"`
	Tools     param.Field[interface{}]                      `json:"tools"`
	URL       param.Field[string]                           `json:"url"`
}

func (r ToolSourceUpdateParamsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolSourceUpdateParamsSource) implementsToolSourceUpdateParamsSourceUnion() {}

// Satisfied by [ToolSourceUpdateParamsSourceUpdateHTTPToolSourceRequest],
// [ToolSourceUpdateParamsSourceUpdateMcpToolSourceRequest],
// [ToolSourceUpdateParamsSource].
type ToolSourceUpdateParamsSourceUnion interface {
	implementsToolSourceUpdateParamsSourceUnion()
}

type ToolSourceUpdateParamsSourceUpdateHTTPToolSourceRequest struct {
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      param.Field[ToolSourceAuthParam]                                         `json:"auth" api:"required"`
	Namespace param.Field[string]                                                      `json:"namespace" api:"required"`
	Tools     param.Field[[]HTTPToolDefinitionParam]                                   `json:"tools" api:"required"`
	Type      param.Field[ToolSourceUpdateParamsSourceUpdateHTTPToolSourceRequestType] `json:"type" api:"required"`
	Approval  param.Field[ToolApprovalMode]                                            `json:"approval"`
	Enabled   param.Field[bool]                                                        `json:"enabled"`
	// HTTP tool execution retry and timeout policy.
	Execution   param.Field[HTTPToolExecutionPolicyParam] `json:"execution"`
	ExtraFields map[string]interface{}                    `json:"-,extras"`
}

func (r ToolSourceUpdateParamsSourceUpdateHTTPToolSourceRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolSourceUpdateParamsSourceUpdateHTTPToolSourceRequest) implementsToolSourceUpdateParamsSourceUnion() {
}

type ToolSourceUpdateParamsSourceUpdateHTTPToolSourceRequestType string

const (
	ToolSourceUpdateParamsSourceUpdateHTTPToolSourceRequestTypeHTTP ToolSourceUpdateParamsSourceUpdateHTTPToolSourceRequestType = "http"
)

func (r ToolSourceUpdateParamsSourceUpdateHTTPToolSourceRequestType) IsKnown() bool {
	switch r {
	case ToolSourceUpdateParamsSourceUpdateHTTPToolSourceRequestTypeHTTP:
		return true
	}
	return false
}

type ToolSourceUpdateParamsSourceUpdateMcpToolSourceRequest struct {
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      param.Field[ToolSourceAuthParam]                                        `json:"auth" api:"required"`
	Namespace param.Field[string]                                                     `json:"namespace" api:"required"`
	Type      param.Field[ToolSourceUpdateParamsSourceUpdateMcpToolSourceRequestType] `json:"type" api:"required"`
	URL       param.Field[string]                                                     `json:"url" api:"required"`
	Approval  param.Field[ToolApprovalMode]                                           `json:"approval"`
	Enabled   param.Field[bool]                                                       `json:"enabled"`
	// MCP discovery and execution retry/timeout policy.
	Execution   param.Field[McpToolExecutionPolicyParam] `json:"execution"`
	ExtraFields map[string]interface{}                   `json:"-,extras"`
}

func (r ToolSourceUpdateParamsSourceUpdateMcpToolSourceRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolSourceUpdateParamsSourceUpdateMcpToolSourceRequest) implementsToolSourceUpdateParamsSourceUnion() {
}

type ToolSourceUpdateParamsSourceUpdateMcpToolSourceRequestType string

const (
	ToolSourceUpdateParamsSourceUpdateMcpToolSourceRequestTypeMcp ToolSourceUpdateParamsSourceUpdateMcpToolSourceRequestType = "mcp"
)

func (r ToolSourceUpdateParamsSourceUpdateMcpToolSourceRequestType) IsKnown() bool {
	switch r {
	case ToolSourceUpdateParamsSourceUpdateMcpToolSourceRequestTypeMcp:
		return true
	}
	return false
}

type ToolSourceUpdateParamsSourceType string

const (
	ToolSourceUpdateParamsSourceTypeHTTP ToolSourceUpdateParamsSourceType = "http"
	ToolSourceUpdateParamsSourceTypeMcp  ToolSourceUpdateParamsSourceType = "mcp"
)

func (r ToolSourceUpdateParamsSourceType) IsKnown() bool {
	switch r {
	case ToolSourceUpdateParamsSourceTypeHTTP, ToolSourceUpdateParamsSourceTypeMcp:
		return true
	}
	return false
}

type ToolSourceListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [ToolSourceListParams]'s query parameters as `url.Values`.
func (r ToolSourceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
