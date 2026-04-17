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

// EnvironmentToolSourceService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentToolSourceService] method instead.
type EnvironmentToolSourceService struct {
	Options []option.RequestOption
}

// NewEnvironmentToolSourceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEnvironmentToolSourceService(opts ...option.RequestOption) (r *EnvironmentToolSourceService) {
	r = &EnvironmentToolSourceService{}
	r.Options = opts
	return
}

func (r *EnvironmentToolSourceService) New(ctx context.Context, environmentID string, body EnvironmentToolSourceNewParams, opts ...option.RequestOption) (res *ToolSource, err error) {
	opts = slices.Concat(r.Options, opts)
	if environmentID == "" {
		err = errors.New("missing required environmentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("environments/%s/tool-sources", environmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *EnvironmentToolSourceService) Get(ctx context.Context, environmentID string, sourceID string, opts ...option.RequestOption) (res *ToolSource, err error) {
	opts = slices.Concat(r.Options, opts)
	if environmentID == "" {
		err = errors.New("missing required environmentId parameter")
		return nil, err
	}
	if sourceID == "" {
		err = errors.New("missing required sourceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("environments/%s/tool-sources/%s", environmentID, sourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *EnvironmentToolSourceService) Update(ctx context.Context, environmentID string, sourceID string, body EnvironmentToolSourceUpdateParams, opts ...option.RequestOption) (res *ToolSource, err error) {
	opts = slices.Concat(r.Options, opts)
	if environmentID == "" {
		err = errors.New("missing required environmentId parameter")
		return nil, err
	}
	if sourceID == "" {
		err = errors.New("missing required sourceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("environments/%s/tool-sources/%s", environmentID, sourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

func (r *EnvironmentToolSourceService) List(ctx context.Context, environmentID string, query EnvironmentToolSourceListParams, opts ...option.RequestOption) (res *EnvironmentToolSourceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if environmentID == "" {
		err = errors.New("missing required environmentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("environments/%s/tool-sources", environmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *EnvironmentToolSourceService) Delete(ctx context.Context, environmentID string, sourceID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if environmentID == "" {
		err = errors.New("missing required environmentId parameter")
		return err
	}
	if sourceID == "" {
		err = errors.New("missing required sourceId parameter")
		return err
	}
	path := fmt.Sprintf("environments/%s/tool-sources/%s", environmentID, sourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type HTTPToolDefinition = interface{}

type HTTPToolExecutionPolicy = interface{}

type HTTPToolSource struct {
	ID string `json:"id" api:"required"`
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth          ToolSourceAuth       `json:"auth" api:"required"`
	CreatedAt     string               `json:"createdAt" api:"required"`
	Enabled       bool                 `json:"enabled" api:"required"`
	EnvironmentID string               `json:"environmentId" api:"required"`
	Namespace     string               `json:"namespace" api:"required"`
	Tools         []HTTPToolDefinition `json:"tools" api:"required"`
	Type          HTTPToolSourceType   `json:"type" api:"required"`
	UpdatedAt     string               `json:"updatedAt" api:"required"`
	Version       float64              `json:"version" api:"required"`
	Approval      ToolApprovalMode     `json:"approval"`
	// HTTP tool execution retry and timeout policy.
	Execution HTTPToolExecutionPolicy `json:"execution"`
	JSON      httpToolSourceJSON      `json:"-"`
}

// httpToolSourceJSON contains the JSON metadata for the struct [HTTPToolSource]
type httpToolSourceJSON struct {
	ID            apijson.Field
	Auth          apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	EnvironmentID apijson.Field
	Namespace     apijson.Field
	Tools         apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	Version       apijson.Field
	Approval      apijson.Field
	Execution     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

type McpToolExecutionPolicy = interface{}

type McpToolSource struct {
	ID string `json:"id" api:"required"`
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth          ToolSourceAuth    `json:"auth" api:"required"`
	CreatedAt     string            `json:"createdAt" api:"required"`
	Enabled       bool              `json:"enabled" api:"required"`
	EnvironmentID string            `json:"environmentId" api:"required"`
	Namespace     string            `json:"namespace" api:"required"`
	Type          McpToolSourceType `json:"type" api:"required"`
	UpdatedAt     string            `json:"updatedAt" api:"required"`
	URL           string            `json:"url" api:"required"`
	Version       float64           `json:"version" api:"required"`
	Approval      ToolApprovalMode  `json:"approval"`
	// MCP discovery and execution retry/timeout policy.
	Execution McpToolExecutionPolicy `json:"execution"`
	JSON      mcpToolSourceJSON      `json:"-"`
}

// mcpToolSourceJSON contains the JSON metadata for the struct [McpToolSource]
type mcpToolSourceJSON struct {
	ID            apijson.Field
	Auth          apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	EnvironmentID apijson.Field
	Namespace     apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	URL           apijson.Field
	Version       apijson.Field
	Approval      apijson.Field
	Execution     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// This field can have the runtime type of [ToolSourceAuth].
	Auth          ToolSourceAuth   `json:"auth" api:"required"`
	CreatedAt     string           `json:"createdAt" api:"required"`
	Enabled       bool             `json:"enabled" api:"required"`
	EnvironmentID string           `json:"environmentId" api:"required"`
	Namespace     string           `json:"namespace" api:"required"`
	Type          ToolSourceType   `json:"type" api:"required"`
	UpdatedAt     string           `json:"updatedAt" api:"required"`
	Version       float64          `json:"version" api:"required"`
	Approval      ToolApprovalMode `json:"approval"`
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
	ID            apijson.Field
	Auth          apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	EnvironmentID apijson.Field
	Namespace     apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	Version       apijson.Field
	Approval      apijson.Field
	Execution     apijson.Field
	Tools         apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

type ToolSourceAuth = interface{}

type EnvironmentToolSourceListResponse struct {
	Cursor  string                                `json:"cursor" api:"required,nullable"`
	HasMore bool                                  `json:"hasMore" api:"required"`
	Sources []ToolSource                          `json:"sources" api:"required"`
	JSON    environmentToolSourceListResponseJSON `json:"-"`
}

// environmentToolSourceListResponseJSON contains the JSON metadata for the struct
// [EnvironmentToolSourceListResponse]
type environmentToolSourceListResponseJSON struct {
	Cursor      apijson.Field
	HasMore     apijson.Field
	Sources     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentToolSourceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentToolSourceListResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentToolSourceNewParams struct {
	Body EnvironmentToolSourceNewParamsBodyUnion `json:"body" api:"required"`
}

func (r EnvironmentToolSourceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type EnvironmentToolSourceNewParamsBody struct {
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      param.Field[ToolSourceAuthParam]                    `json:"auth" api:"required"`
	Namespace param.Field[string]                                 `json:"namespace" api:"required"`
	Type      param.Field[EnvironmentToolSourceNewParamsBodyType] `json:"type" api:"required"`
	Approval  param.Field[ToolApprovalMode]                       `json:"approval"`
	Enabled   param.Field[bool]                                   `json:"enabled"`
	Execution param.Field[interface{}]                            `json:"execution"`
	Tools     param.Field[interface{}]                            `json:"tools"`
	URL       param.Field[string]                                 `json:"url"`
}

func (r EnvironmentToolSourceNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentToolSourceNewParamsBody) implementsEnvironmentToolSourceNewParamsBodyUnion() {}

// Satisfied by [EnvironmentToolSourceNewParamsBodyObject],
// [EnvironmentToolSourceNewParamsBodyObject],
// [EnvironmentToolSourceNewParamsBody].
type EnvironmentToolSourceNewParamsBodyUnion interface {
	implementsEnvironmentToolSourceNewParamsBodyUnion()
}

type EnvironmentToolSourceNewParamsBodyObject struct {
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      param.Field[ToolSourceAuthParam]                          `json:"auth" api:"required"`
	Namespace param.Field[string]                                       `json:"namespace" api:"required"`
	Tools     param.Field[[]HTTPToolDefinitionParam]                    `json:"tools" api:"required"`
	Type      param.Field[EnvironmentToolSourceNewParamsBodyObjectType] `json:"type" api:"required"`
	Approval  param.Field[ToolApprovalMode]                             `json:"approval"`
	Enabled   param.Field[bool]                                         `json:"enabled"`
	// HTTP tool execution retry and timeout policy.
	Execution   param.Field[HTTPToolExecutionPolicyParam] `json:"execution"`
	ExtraFields map[string]interface{}                    `json:"-,extras"`
}

func (r EnvironmentToolSourceNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentToolSourceNewParamsBodyObject) implementsEnvironmentToolSourceNewParamsBodyUnion() {
}

type EnvironmentToolSourceNewParamsBodyObjectType string

const (
	EnvironmentToolSourceNewParamsBodyObjectTypeHTTP EnvironmentToolSourceNewParamsBodyObjectType = "http"
)

func (r EnvironmentToolSourceNewParamsBodyObjectType) IsKnown() bool {
	switch r {
	case EnvironmentToolSourceNewParamsBodyObjectTypeHTTP:
		return true
	}
	return false
}

type EnvironmentToolSourceNewParamsBodyType string

const (
	EnvironmentToolSourceNewParamsBodyTypeHTTP EnvironmentToolSourceNewParamsBodyType = "http"
	EnvironmentToolSourceNewParamsBodyTypeMcp  EnvironmentToolSourceNewParamsBodyType = "mcp"
)

func (r EnvironmentToolSourceNewParamsBodyType) IsKnown() bool {
	switch r {
	case EnvironmentToolSourceNewParamsBodyTypeHTTP, EnvironmentToolSourceNewParamsBodyTypeMcp:
		return true
	}
	return false
}

type EnvironmentToolSourceUpdateParams struct {
	Body EnvironmentToolSourceUpdateParamsBodyUnion `json:"body" api:"required"`
}

func (r EnvironmentToolSourceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type EnvironmentToolSourceUpdateParamsBody struct {
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      param.Field[ToolSourceAuthParam]                       `json:"auth" api:"required"`
	Namespace param.Field[string]                                    `json:"namespace" api:"required"`
	Type      param.Field[EnvironmentToolSourceUpdateParamsBodyType] `json:"type" api:"required"`
	Approval  param.Field[ToolApprovalMode]                          `json:"approval"`
	Enabled   param.Field[bool]                                      `json:"enabled"`
	Execution param.Field[interface{}]                               `json:"execution"`
	Tools     param.Field[interface{}]                               `json:"tools"`
	URL       param.Field[string]                                    `json:"url"`
}

func (r EnvironmentToolSourceUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentToolSourceUpdateParamsBody) implementsEnvironmentToolSourceUpdateParamsBodyUnion() {
}

// Satisfied by [EnvironmentToolSourceUpdateParamsBodyObject],
// [EnvironmentToolSourceUpdateParamsBodyObject],
// [EnvironmentToolSourceUpdateParamsBody].
type EnvironmentToolSourceUpdateParamsBodyUnion interface {
	implementsEnvironmentToolSourceUpdateParamsBodyUnion()
}

type EnvironmentToolSourceUpdateParamsBodyObject struct {
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      param.Field[ToolSourceAuthParam]                             `json:"auth" api:"required"`
	Namespace param.Field[string]                                          `json:"namespace" api:"required"`
	Tools     param.Field[[]HTTPToolDefinitionParam]                       `json:"tools" api:"required"`
	Type      param.Field[EnvironmentToolSourceUpdateParamsBodyObjectType] `json:"type" api:"required"`
	Approval  param.Field[ToolApprovalMode]                                `json:"approval"`
	Enabled   param.Field[bool]                                            `json:"enabled"`
	// HTTP tool execution retry and timeout policy.
	Execution   param.Field[HTTPToolExecutionPolicyParam] `json:"execution"`
	ExtraFields map[string]interface{}                    `json:"-,extras"`
}

func (r EnvironmentToolSourceUpdateParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentToolSourceUpdateParamsBodyObject) implementsEnvironmentToolSourceUpdateParamsBodyUnion() {
}

type EnvironmentToolSourceUpdateParamsBodyObjectType string

const (
	EnvironmentToolSourceUpdateParamsBodyObjectTypeHTTP EnvironmentToolSourceUpdateParamsBodyObjectType = "http"
)

func (r EnvironmentToolSourceUpdateParamsBodyObjectType) IsKnown() bool {
	switch r {
	case EnvironmentToolSourceUpdateParamsBodyObjectTypeHTTP:
		return true
	}
	return false
}

type EnvironmentToolSourceUpdateParamsBodyType string

const (
	EnvironmentToolSourceUpdateParamsBodyTypeHTTP EnvironmentToolSourceUpdateParamsBodyType = "http"
	EnvironmentToolSourceUpdateParamsBodyTypeMcp  EnvironmentToolSourceUpdateParamsBodyType = "mcp"
)

func (r EnvironmentToolSourceUpdateParamsBodyType) IsKnown() bool {
	switch r {
	case EnvironmentToolSourceUpdateParamsBodyTypeHTTP, EnvironmentToolSourceUpdateParamsBodyTypeMcp:
		return true
	}
	return false
}

type EnvironmentToolSourceListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [EnvironmentToolSourceListParams]'s query parameters as
// `url.Values`.
func (r EnvironmentToolSourceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
