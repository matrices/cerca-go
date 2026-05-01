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

// ToolService contains methods and other services that help with interacting with
// the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolService] method instead.
type ToolService struct {
	Options []option.RequestOption
}

// NewToolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewToolService(opts ...option.RequestOption) (r *ToolService) {
	r = &ToolService{}
	r.Options = opts
	return
}

// Create tool
func (r *ToolService) New(ctx context.Context, fleetID string, body ToolNewParams, opts ...option.RequestOption) (res *ToolSource, err error) {
	opts = slices.Concat(r.Options, opts)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("fleets/%s/tools", fleetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve tool source
func (r *ToolService) Get(ctx context.Context, fleetID string, sourceID string, opts ...option.RequestOption) (res *ToolSource, err error) {
	opts = slices.Concat(r.Options, opts)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return nil, err
	}
	if sourceID == "" {
		err = errors.New("missing required sourceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("fleets/%s/tools/%s", fleetID, sourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update tool source
func (r *ToolService) Update(ctx context.Context, fleetID string, sourceID string, body ToolUpdateParams, opts ...option.RequestOption) (res *ToolSource, err error) {
	opts = slices.Concat(r.Options, opts)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return nil, err
	}
	if sourceID == "" {
		err = errors.New("missing required sourceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("fleets/%s/tools/%s", fleetID, sourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// List tools
func (r *ToolService) List(ctx context.Context, fleetID string, query ToolListParams, opts ...option.RequestOption) (res *pagination.SourcesCursorPage[ToolSource], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("fleets/%s/tools", fleetID)
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

// List tools
func (r *ToolService) ListAutoPaging(ctx context.Context, fleetID string, query ToolListParams, opts ...option.RequestOption) *pagination.SourcesCursorPageAutoPager[ToolSource] {
	return pagination.NewSourcesCursorPageAutoPager(r.List(ctx, fleetID, query, opts...))
}

// Delete tool source
func (r *ToolService) Delete(ctx context.Context, fleetID string, sourceID string, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("fleets/%s/tools/%s", fleetID, sourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Where an API key or OAuth access token should be injected into outgoing
// tool-source requests.
type AuthInjection struct {
	Location AuthInjectionLocation `json:"location" api:"required"`
	Name     string                `json:"name" api:"required"`
	// Optional value template. For OAuth connection auth, use values such as
	// `Bearer ${token}`.
	Value string            `json:"value"`
	JSON  authInjectionJSON `json:"-"`
	union AuthInjectionUnion
}

// authInjectionJSON contains the JSON metadata for the struct [AuthInjection]
type authInjectionJSON struct {
	Location    apijson.Field
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authInjectionJSON) RawJSON() string {
	return r.raw
}

func (r *AuthInjection) UnmarshalJSON(data []byte) (err error) {
	*r = AuthInjection{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthInjectionUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [AuthInjectionHeaderAuthInjection],
// [AuthInjectionQueryAuthInjection].
func (r AuthInjection) AsUnion() AuthInjectionUnion {
	return r.union
}

// Where an API key or OAuth access token should be injected into outgoing
// tool-source requests.
//
// Union satisfied by [AuthInjectionHeaderAuthInjection] or
// [AuthInjectionQueryAuthInjection].
type AuthInjectionUnion interface {
	implementsAuthInjection()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthInjectionUnion)(nil)).Elem(),
		"location",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AuthInjectionHeaderAuthInjection{}),
			DiscriminatorValue: "header",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AuthInjectionQueryAuthInjection{}),
			DiscriminatorValue: "query",
		},
	)
}

type AuthInjectionHeaderAuthInjection struct {
	Location AuthInjectionHeaderAuthInjectionLocation `json:"location" api:"required"`
	Name     string                                   `json:"name" api:"required"`
	// Optional value template. For OAuth connection auth, use values such as
	// `Bearer ${token}`.
	Value string                               `json:"value"`
	JSON  authInjectionHeaderAuthInjectionJSON `json:"-"`
}

// authInjectionHeaderAuthInjectionJSON contains the JSON metadata for the struct
// [AuthInjectionHeaderAuthInjection]
type authInjectionHeaderAuthInjectionJSON struct {
	Location    apijson.Field
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthInjectionHeaderAuthInjection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authInjectionHeaderAuthInjectionJSON) RawJSON() string {
	return r.raw
}

func (r AuthInjectionHeaderAuthInjection) implementsAuthInjection() {}

type AuthInjectionHeaderAuthInjectionLocation string

const (
	AuthInjectionHeaderAuthInjectionLocationHeader AuthInjectionHeaderAuthInjectionLocation = "header"
)

func (r AuthInjectionHeaderAuthInjectionLocation) IsKnown() bool {
	switch r {
	case AuthInjectionHeaderAuthInjectionLocationHeader:
		return true
	}
	return false
}

type AuthInjectionQueryAuthInjection struct {
	Location AuthInjectionQueryAuthInjectionLocation `json:"location" api:"required"`
	Name     string                                  `json:"name" api:"required"`
	JSON     authInjectionQueryAuthInjectionJSON     `json:"-"`
}

// authInjectionQueryAuthInjectionJSON contains the JSON metadata for the struct
// [AuthInjectionQueryAuthInjection]
type authInjectionQueryAuthInjectionJSON struct {
	Location    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthInjectionQueryAuthInjection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authInjectionQueryAuthInjectionJSON) RawJSON() string {
	return r.raw
}

func (r AuthInjectionQueryAuthInjection) implementsAuthInjection() {}

type AuthInjectionQueryAuthInjectionLocation string

const (
	AuthInjectionQueryAuthInjectionLocationQuery AuthInjectionQueryAuthInjectionLocation = "query"
)

func (r AuthInjectionQueryAuthInjectionLocation) IsKnown() bool {
	switch r {
	case AuthInjectionQueryAuthInjectionLocationQuery:
		return true
	}
	return false
}

type AuthInjectionLocation string

const (
	AuthInjectionLocationHeader AuthInjectionLocation = "header"
	AuthInjectionLocationQuery  AuthInjectionLocation = "query"
)

func (r AuthInjectionLocation) IsKnown() bool {
	switch r {
	case AuthInjectionLocationHeader, AuthInjectionLocationQuery:
		return true
	}
	return false
}

// Where an API key or OAuth access token should be injected into outgoing
// tool-source requests.
type AuthInjectionParam struct {
	Location param.Field[AuthInjectionLocation] `json:"location" api:"required"`
	Name     param.Field[string]                `json:"name" api:"required"`
	// Optional value template. For OAuth connection auth, use values such as
	// `Bearer ${token}`.
	Value param.Field[string] `json:"value"`
}

func (r AuthInjectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthInjectionParam) implementsAuthInjectionUnionParam() {}

// Where an API key or OAuth access token should be injected into outgoing
// tool-source requests.
//
// Satisfied by [AuthInjectionHeaderAuthInjectionParam],
// [AuthInjectionQueryAuthInjectionParam], [AuthInjectionParam].
type AuthInjectionUnionParam interface {
	implementsAuthInjectionUnionParam()
}

type AuthInjectionHeaderAuthInjectionParam struct {
	Location param.Field[AuthInjectionHeaderAuthInjectionLocation] `json:"location" api:"required"`
	Name     param.Field[string]                                   `json:"name" api:"required"`
	// Optional value template. For OAuth connection auth, use values such as
	// `Bearer ${token}`.
	Value param.Field[string] `json:"value"`
}

func (r AuthInjectionHeaderAuthInjectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthInjectionHeaderAuthInjectionParam) implementsAuthInjectionUnionParam() {}

type AuthInjectionQueryAuthInjectionParam struct {
	Location param.Field[AuthInjectionQueryAuthInjectionLocation] `json:"location" api:"required"`
	Name     param.Field[string]                                  `json:"name" api:"required"`
}

func (r AuthInjectionQueryAuthInjectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthInjectionQueryAuthInjectionParam) implementsAuthInjectionUnionParam() {}

// HTTP endpoint invoked when the tool is called.
type HTTPEndpoint struct {
	Method  HTTPEndpointMethod `json:"method" api:"required"`
	URL     string             `json:"url" api:"required"`
	Body    HTTPEndpointBody   `json:"body"`
	Headers map[string]string  `json:"headers"`
	Path    map[string]string  `json:"path"`
	Query   map[string]string  `json:"query"`
	JSON    httpEndpointJSON   `json:"-"`
}

// httpEndpointJSON contains the JSON metadata for the struct [HTTPEndpoint]
type httpEndpointJSON struct {
	Method      apijson.Field
	URL         apijson.Field
	Body        apijson.Field
	Headers     apijson.Field
	Path        apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpEndpointJSON) RawJSON() string {
	return r.raw
}

type HTTPEndpointMethod string

const (
	HTTPEndpointMethodGet    HTTPEndpointMethod = "GET"
	HTTPEndpointMethodPost   HTTPEndpointMethod = "POST"
	HTTPEndpointMethodPut    HTTPEndpointMethod = "PUT"
	HTTPEndpointMethodPatch  HTTPEndpointMethod = "PATCH"
	HTTPEndpointMethodDelete HTTPEndpointMethod = "DELETE"
)

func (r HTTPEndpointMethod) IsKnown() bool {
	switch r {
	case HTTPEndpointMethodGet, HTTPEndpointMethodPost, HTTPEndpointMethodPut, HTTPEndpointMethodPatch, HTTPEndpointMethodDelete:
		return true
	}
	return false
}

type HTTPEndpointBody string

const (
	HTTPEndpointBodyJsonParams HTTPEndpointBody = "json_params"
)

func (r HTTPEndpointBody) IsKnown() bool {
	switch r {
	case HTTPEndpointBodyJsonParams:
		return true
	}
	return false
}

// HTTP endpoint invoked when the tool is called.
type HTTPEndpointParam struct {
	Method  param.Field[HTTPEndpointMethod] `json:"method" api:"required"`
	URL     param.Field[string]             `json:"url" api:"required"`
	Body    param.Field[HTTPEndpointBody]   `json:"body"`
	Headers param.Field[map[string]string]  `json:"headers"`
	Path    param.Field[map[string]string]  `json:"path"`
	Query   param.Field[map[string]string]  `json:"query"`
}

func (r HTTPEndpointParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Definition for a single HTTP tool exposed by a tool source.
type HTTPToolDefinition struct {
	Description string `json:"description" api:"required"`
	// HTTP endpoint invoked when the tool is called.
	Endpoint HTTPEndpoint `json:"endpoint" api:"required"`
	// JSON Schema object describing tool input parameters.
	InputSchema map[string]interface{} `json:"inputSchema" api:"required"`
	Name        string                 `json:"name" api:"required"`
	Approval    ToolApprovalMode       `json:"approval"`
	// HTTP tool execution retry and timeout policy.
	Execution HTTPToolExecutionPolicy `json:"execution"`
	// How the HTTP response should be normalized for the agent.
	Response ResponseNormalizationHint `json:"response"`
	JSON     httpToolDefinitionJSON    `json:"-"`
}

// httpToolDefinitionJSON contains the JSON metadata for the struct
// [HTTPToolDefinition]
type httpToolDefinitionJSON struct {
	Description apijson.Field
	Endpoint    apijson.Field
	InputSchema apijson.Field
	Name        apijson.Field
	Approval    apijson.Field
	Execution   apijson.Field
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPToolDefinition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpToolDefinitionJSON) RawJSON() string {
	return r.raw
}

// Definition for a single HTTP tool exposed by a tool source.
type HTTPToolDefinitionParam struct {
	Description param.Field[string] `json:"description" api:"required"`
	// HTTP endpoint invoked when the tool is called.
	Endpoint param.Field[HTTPEndpointParam] `json:"endpoint" api:"required"`
	// JSON Schema object describing tool input parameters.
	InputSchema param.Field[map[string]interface{}] `json:"inputSchema" api:"required"`
	Name        param.Field[string]                 `json:"name" api:"required"`
	Approval    param.Field[ToolApprovalMode]       `json:"approval"`
	// HTTP tool execution retry and timeout policy.
	Execution param.Field[HTTPToolExecutionPolicyParam] `json:"execution"`
	// How the HTTP response should be normalized for the agent.
	Response param.Field[ResponseNormalizationHintParam] `json:"response"`
}

func (r HTTPToolDefinitionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// HTTP tool execution retry and timeout policy.
type HTTPToolExecutionPolicy struct {
	IdempotencyKeyHeader string                           `json:"idempotencyKeyHeader"`
	MaxAttempts          float64                          `json:"maxAttempts"`
	RetryMode            HTTPToolExecutionPolicyRetryMode `json:"retryMode"`
	TimeoutMs            float64                          `json:"timeoutMs"`
	JSON                 httpToolExecutionPolicyJSON      `json:"-"`
}

// httpToolExecutionPolicyJSON contains the JSON metadata for the struct
// [HTTPToolExecutionPolicy]
type httpToolExecutionPolicyJSON struct {
	IdempotencyKeyHeader apijson.Field
	MaxAttempts          apijson.Field
	RetryMode            apijson.Field
	TimeoutMs            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HTTPToolExecutionPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpToolExecutionPolicyJSON) RawJSON() string {
	return r.raw
}

type HTTPToolExecutionPolicyRetryMode string

const (
	HTTPToolExecutionPolicyRetryModeDisabled HTTPToolExecutionPolicyRetryMode = "disabled"
	HTTPToolExecutionPolicyRetryModeSafeOnly HTTPToolExecutionPolicyRetryMode = "safe_only"
	HTTPToolExecutionPolicyRetryModeEnabled  HTTPToolExecutionPolicyRetryMode = "enabled"
)

func (r HTTPToolExecutionPolicyRetryMode) IsKnown() bool {
	switch r {
	case HTTPToolExecutionPolicyRetryModeDisabled, HTTPToolExecutionPolicyRetryModeSafeOnly, HTTPToolExecutionPolicyRetryModeEnabled:
		return true
	}
	return false
}

// HTTP tool execution retry and timeout policy.
type HTTPToolExecutionPolicyParam struct {
	IdempotencyKeyHeader param.Field[string]                           `json:"idempotencyKeyHeader"`
	MaxAttempts          param.Field[float64]                          `json:"maxAttempts"`
	RetryMode            param.Field[HTTPToolExecutionPolicyRetryMode] `json:"retryMode"`
	TimeoutMs            param.Field[float64]                          `json:"timeoutMs"`
}

func (r HTTPToolExecutionPolicyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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

// MCP discovery and execution retry/timeout policy.
type McpToolExecutionPolicy struct {
	DiscoveryTimeoutMs float64                         `json:"discoveryTimeoutMs"`
	ExecutionTimeoutMs float64                         `json:"executionTimeoutMs"`
	MaxAttempts        float64                         `json:"maxAttempts"`
	RetryMode          McpToolExecutionPolicyRetryMode `json:"retryMode"`
	JSON               mcpToolExecutionPolicyJSON      `json:"-"`
}

// mcpToolExecutionPolicyJSON contains the JSON metadata for the struct
// [McpToolExecutionPolicy]
type mcpToolExecutionPolicyJSON struct {
	DiscoveryTimeoutMs apijson.Field
	ExecutionTimeoutMs apijson.Field
	MaxAttempts        apijson.Field
	RetryMode          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *McpToolExecutionPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpToolExecutionPolicyJSON) RawJSON() string {
	return r.raw
}

type McpToolExecutionPolicyRetryMode string

const (
	McpToolExecutionPolicyRetryModeDisabled    McpToolExecutionPolicyRetryMode = "disabled"
	McpToolExecutionPolicyRetryModeConnectOnly McpToolExecutionPolicyRetryMode = "connect_only"
	McpToolExecutionPolicyRetryModeEnabled     McpToolExecutionPolicyRetryMode = "enabled"
)

func (r McpToolExecutionPolicyRetryMode) IsKnown() bool {
	switch r {
	case McpToolExecutionPolicyRetryModeDisabled, McpToolExecutionPolicyRetryModeConnectOnly, McpToolExecutionPolicyRetryModeEnabled:
		return true
	}
	return false
}

// MCP discovery and execution retry/timeout policy.
type McpToolExecutionPolicyParam struct {
	DiscoveryTimeoutMs param.Field[float64]                         `json:"discoveryTimeoutMs"`
	ExecutionTimeoutMs param.Field[float64]                         `json:"executionTimeoutMs"`
	MaxAttempts        param.Field[float64]                         `json:"maxAttempts"`
	RetryMode          param.Field[McpToolExecutionPolicyRetryMode] `json:"retryMode"`
}

func (r McpToolExecutionPolicyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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

// Source-owned OAuth exchange configuration. Public schemas use connection IDs for
// stored client secrets and assertions.
type OAuthExchangeConfig struct {
	GrantType                OAuthExchangeConfigGrantType `json:"grantType" api:"required"`
	TokenURL                 string                       `json:"tokenUrl" api:"required"`
	AssertionConnectionID    string                       `json:"assertionConnectionId"`
	Audience                 string                       `json:"audience"`
	ClientID                 string                       `json:"clientId"`
	ClientSecretConnectionID string                       `json:"clientSecretConnectionId"`
	// This field can have the runtime type of [map[string]string].
	ExtraParams interface{} `json:"extraParams"`
	Issuer      string      `json:"issuer"`
	// This field can have the runtime type of [[]string].
	Scopes  interface{}             `json:"scopes"`
	Subject string                  `json:"subject"`
	JSON    oauthExchangeConfigJSON `json:"-"`
	union   OAuthExchangeConfigUnion
}

// oauthExchangeConfigJSON contains the JSON metadata for the struct
// [OAuthExchangeConfig]
type oauthExchangeConfigJSON struct {
	GrantType                apijson.Field
	TokenURL                 apijson.Field
	AssertionConnectionID    apijson.Field
	Audience                 apijson.Field
	ClientID                 apijson.Field
	ClientSecretConnectionID apijson.Field
	ExtraParams              apijson.Field
	Issuer                   apijson.Field
	Scopes                   apijson.Field
	Subject                  apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r oauthExchangeConfigJSON) RawJSON() string {
	return r.raw
}

func (r *OAuthExchangeConfig) UnmarshalJSON(data []byte) (err error) {
	*r = OAuthExchangeConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [OAuthExchangeConfigUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [OAuthExchangeConfigClientCredentialsOAuthExchange],
// [OAuthExchangeConfigJwtBearerOAuthExchange].
func (r OAuthExchangeConfig) AsUnion() OAuthExchangeConfigUnion {
	return r.union
}

// Source-owned OAuth exchange configuration. Public schemas use connection IDs for
// stored client secrets and assertions.
//
// Union satisfied by [OAuthExchangeConfigClientCredentialsOAuthExchange] or
// [OAuthExchangeConfigJwtBearerOAuthExchange].
type OAuthExchangeConfigUnion interface {
	implementsOAuthExchangeConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OAuthExchangeConfigUnion)(nil)).Elem(),
		"grantType",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(OAuthExchangeConfigClientCredentialsOAuthExchange{}),
			DiscriminatorValue: "client_credentials",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(OAuthExchangeConfigJwtBearerOAuthExchange{}),
			DiscriminatorValue: "jwt_bearer",
		},
	)
}

type OAuthExchangeConfigClientCredentialsOAuthExchange struct {
	ClientID                 string                                                     `json:"clientId" api:"required"`
	ClientSecretConnectionID string                                                     `json:"clientSecretConnectionId" api:"required"`
	GrantType                OAuthExchangeConfigClientCredentialsOAuthExchangeGrantType `json:"grantType" api:"required"`
	TokenURL                 string                                                     `json:"tokenUrl" api:"required"`
	Audience                 string                                                     `json:"audience"`
	ExtraParams              map[string]string                                          `json:"extraParams"`
	Scopes                   []string                                                   `json:"scopes"`
	JSON                     oauthExchangeConfigClientCredentialsOAuthExchangeJSON      `json:"-"`
}

// oauthExchangeConfigClientCredentialsOAuthExchangeJSON contains the JSON metadata
// for the struct [OAuthExchangeConfigClientCredentialsOAuthExchange]
type oauthExchangeConfigClientCredentialsOAuthExchangeJSON struct {
	ClientID                 apijson.Field
	ClientSecretConnectionID apijson.Field
	GrantType                apijson.Field
	TokenURL                 apijson.Field
	Audience                 apijson.Field
	ExtraParams              apijson.Field
	Scopes                   apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *OAuthExchangeConfigClientCredentialsOAuthExchange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthExchangeConfigClientCredentialsOAuthExchangeJSON) RawJSON() string {
	return r.raw
}

func (r OAuthExchangeConfigClientCredentialsOAuthExchange) implementsOAuthExchangeConfig() {}

type OAuthExchangeConfigClientCredentialsOAuthExchangeGrantType string

const (
	OAuthExchangeConfigClientCredentialsOAuthExchangeGrantTypeClientCredentials OAuthExchangeConfigClientCredentialsOAuthExchangeGrantType = "client_credentials"
)

func (r OAuthExchangeConfigClientCredentialsOAuthExchangeGrantType) IsKnown() bool {
	switch r {
	case OAuthExchangeConfigClientCredentialsOAuthExchangeGrantTypeClientCredentials:
		return true
	}
	return false
}

type OAuthExchangeConfigJwtBearerOAuthExchange struct {
	AssertionConnectionID string                                             `json:"assertionConnectionId" api:"required"`
	Audience              string                                             `json:"audience" api:"required"`
	GrantType             OAuthExchangeConfigJwtBearerOAuthExchangeGrantType `json:"grantType" api:"required"`
	Issuer                string                                             `json:"issuer" api:"required"`
	TokenURL              string                                             `json:"tokenUrl" api:"required"`
	ExtraParams           map[string]string                                  `json:"extraParams"`
	Scopes                []string                                           `json:"scopes"`
	Subject               string                                             `json:"subject"`
	JSON                  oauthExchangeConfigJwtBearerOAuthExchangeJSON      `json:"-"`
}

// oauthExchangeConfigJwtBearerOAuthExchangeJSON contains the JSON metadata for the
// struct [OAuthExchangeConfigJwtBearerOAuthExchange]
type oauthExchangeConfigJwtBearerOAuthExchangeJSON struct {
	AssertionConnectionID apijson.Field
	Audience              apijson.Field
	GrantType             apijson.Field
	Issuer                apijson.Field
	TokenURL              apijson.Field
	ExtraParams           apijson.Field
	Scopes                apijson.Field
	Subject               apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *OAuthExchangeConfigJwtBearerOAuthExchange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthExchangeConfigJwtBearerOAuthExchangeJSON) RawJSON() string {
	return r.raw
}

func (r OAuthExchangeConfigJwtBearerOAuthExchange) implementsOAuthExchangeConfig() {}

type OAuthExchangeConfigJwtBearerOAuthExchangeGrantType string

const (
	OAuthExchangeConfigJwtBearerOAuthExchangeGrantTypeJwtBearer OAuthExchangeConfigJwtBearerOAuthExchangeGrantType = "jwt_bearer"
)

func (r OAuthExchangeConfigJwtBearerOAuthExchangeGrantType) IsKnown() bool {
	switch r {
	case OAuthExchangeConfigJwtBearerOAuthExchangeGrantTypeJwtBearer:
		return true
	}
	return false
}

type OAuthExchangeConfigGrantType string

const (
	OAuthExchangeConfigGrantTypeClientCredentials OAuthExchangeConfigGrantType = "client_credentials"
	OAuthExchangeConfigGrantTypeJwtBearer         OAuthExchangeConfigGrantType = "jwt_bearer"
)

func (r OAuthExchangeConfigGrantType) IsKnown() bool {
	switch r {
	case OAuthExchangeConfigGrantTypeClientCredentials, OAuthExchangeConfigGrantTypeJwtBearer:
		return true
	}
	return false
}

// Source-owned OAuth exchange configuration. Public schemas use connection IDs for
// stored client secrets and assertions.
type OAuthExchangeConfigParam struct {
	GrantType                param.Field[OAuthExchangeConfigGrantType] `json:"grantType" api:"required"`
	TokenURL                 param.Field[string]                       `json:"tokenUrl" api:"required"`
	AssertionConnectionID    param.Field[string]                       `json:"assertionConnectionId"`
	Audience                 param.Field[string]                       `json:"audience"`
	ClientID                 param.Field[string]                       `json:"clientId"`
	ClientSecretConnectionID param.Field[string]                       `json:"clientSecretConnectionId"`
	ExtraParams              param.Field[interface{}]                  `json:"extraParams"`
	Issuer                   param.Field[string]                       `json:"issuer"`
	Scopes                   param.Field[interface{}]                  `json:"scopes"`
	Subject                  param.Field[string]                       `json:"subject"`
}

func (r OAuthExchangeConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OAuthExchangeConfigParam) implementsOAuthExchangeConfigUnionParam() {}

// Source-owned OAuth exchange configuration. Public schemas use connection IDs for
// stored client secrets and assertions.
//
// Satisfied by [OAuthExchangeConfigClientCredentialsOAuthExchangeParam],
// [OAuthExchangeConfigJwtBearerOAuthExchangeParam], [OAuthExchangeConfigParam].
type OAuthExchangeConfigUnionParam interface {
	implementsOAuthExchangeConfigUnionParam()
}

type OAuthExchangeConfigClientCredentialsOAuthExchangeParam struct {
	ClientID                 param.Field[string]                                                     `json:"clientId" api:"required"`
	ClientSecretConnectionID param.Field[string]                                                     `json:"clientSecretConnectionId" api:"required"`
	GrantType                param.Field[OAuthExchangeConfigClientCredentialsOAuthExchangeGrantType] `json:"grantType" api:"required"`
	TokenURL                 param.Field[string]                                                     `json:"tokenUrl" api:"required"`
	Audience                 param.Field[string]                                                     `json:"audience"`
	ExtraParams              param.Field[map[string]string]                                          `json:"extraParams"`
	Scopes                   param.Field[[]string]                                                   `json:"scopes"`
}

func (r OAuthExchangeConfigClientCredentialsOAuthExchangeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OAuthExchangeConfigClientCredentialsOAuthExchangeParam) implementsOAuthExchangeConfigUnionParam() {
}

type OAuthExchangeConfigJwtBearerOAuthExchangeParam struct {
	AssertionConnectionID param.Field[string]                                             `json:"assertionConnectionId" api:"required"`
	Audience              param.Field[string]                                             `json:"audience" api:"required"`
	GrantType             param.Field[OAuthExchangeConfigJwtBearerOAuthExchangeGrantType] `json:"grantType" api:"required"`
	Issuer                param.Field[string]                                             `json:"issuer" api:"required"`
	TokenURL              param.Field[string]                                             `json:"tokenUrl" api:"required"`
	ExtraParams           param.Field[map[string]string]                                  `json:"extraParams"`
	Scopes                param.Field[[]string]                                           `json:"scopes"`
	Subject               param.Field[string]                                             `json:"subject"`
}

func (r OAuthExchangeConfigJwtBearerOAuthExchangeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OAuthExchangeConfigJwtBearerOAuthExchangeParam) implementsOAuthExchangeConfigUnionParam() {}

// How the HTTP response should be normalized for the agent.
type ResponseNormalizationHint struct {
	Mode ResponseNormalizationHintMode `json:"mode" api:"required"`
	JSON responseNormalizationHintJSON `json:"-"`
}

// responseNormalizationHintJSON contains the JSON metadata for the struct
// [ResponseNormalizationHint]
type responseNormalizationHintJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseNormalizationHint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseNormalizationHintJSON) RawJSON() string {
	return r.raw
}

type ResponseNormalizationHintMode string

const (
	ResponseNormalizationHintModeAuto     ResponseNormalizationHintMode = "auto"
	ResponseNormalizationHintModeJson     ResponseNormalizationHintMode = "json"
	ResponseNormalizationHintModeMarkdown ResponseNormalizationHintMode = "markdown"
	ResponseNormalizationHintModeBlob     ResponseNormalizationHintMode = "blob"
)

func (r ResponseNormalizationHintMode) IsKnown() bool {
	switch r {
	case ResponseNormalizationHintModeAuto, ResponseNormalizationHintModeJson, ResponseNormalizationHintModeMarkdown, ResponseNormalizationHintModeBlob:
		return true
	}
	return false
}

// How the HTTP response should be normalized for the agent.
type ResponseNormalizationHintParam struct {
	Mode param.Field[ResponseNormalizationHintMode] `json:"mode" api:"required"`
}

func (r ResponseNormalizationHintParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

// Tool source authentication configuration. The `kind` field selects one of
// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
type ToolSourceAuth struct {
	Kind         ToolSourceAuthKind `json:"kind" api:"required"`
	ConnectionID string             `json:"connectionId"`
	// Source-owned OAuth exchange configuration. Public schemas use connection IDs for
	// stored client secrets and assertions.
	Exchange OAuthExchangeConfig `json:"exchange"`
	// Where an API key or OAuth access token should be injected into outgoing
	// tool-source requests.
	Inject   AuthInjection `json:"inject"`
	Provider string        `json:"provider"`
	// This field can have the runtime type of [[]string].
	RequiredScopes interface{}        `json:"requiredScopes"`
	JSON           toolSourceAuthJSON `json:"-"`
	union          ToolSourceAuthUnion
}

// toolSourceAuthJSON contains the JSON metadata for the struct [ToolSourceAuth]
type toolSourceAuthJSON struct {
	Kind           apijson.Field
	ConnectionID   apijson.Field
	Exchange       apijson.Field
	Inject         apijson.Field
	Provider       apijson.Field
	RequiredScopes apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r toolSourceAuthJSON) RawJSON() string {
	return r.raw
}

func (r *ToolSourceAuth) UnmarshalJSON(data []byte) (err error) {
	*r = ToolSourceAuth{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ToolSourceAuthUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ToolSourceAuthNoToolSourceAuth],
// [ToolSourceAuthAPIKeyToolSourceAuth],
// [ToolSourceAuthOAuthExchangeToolSourceAuth],
// [ToolSourceAuthOAuthConnectionToolSourceAuth].
func (r ToolSourceAuth) AsUnion() ToolSourceAuthUnion {
	return r.union
}

// Tool source authentication configuration. The `kind` field selects one of
// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
//
// Union satisfied by [ToolSourceAuthNoToolSourceAuth],
// [ToolSourceAuthAPIKeyToolSourceAuth],
// [ToolSourceAuthOAuthExchangeToolSourceAuth] or
// [ToolSourceAuthOAuthConnectionToolSourceAuth].
type ToolSourceAuthUnion interface {
	implementsToolSourceAuth()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ToolSourceAuthUnion)(nil)).Elem(),
		"kind",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ToolSourceAuthNoToolSourceAuth{}),
			DiscriminatorValue: "none",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ToolSourceAuthAPIKeyToolSourceAuth{}),
			DiscriminatorValue: "api_key",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ToolSourceAuthOAuthExchangeToolSourceAuth{}),
			DiscriminatorValue: "oauth_exchange",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ToolSourceAuthOAuthConnectionToolSourceAuth{}),
			DiscriminatorValue: "oauth_connection",
		},
	)
}

type ToolSourceAuthNoToolSourceAuth struct {
	Kind ToolSourceAuthNoToolSourceAuthKind `json:"kind" api:"required"`
	JSON toolSourceAuthNoToolSourceAuthJSON `json:"-"`
}

// toolSourceAuthNoToolSourceAuthJSON contains the JSON metadata for the struct
// [ToolSourceAuthNoToolSourceAuth]
type toolSourceAuthNoToolSourceAuthJSON struct {
	Kind        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSourceAuthNoToolSourceAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSourceAuthNoToolSourceAuthJSON) RawJSON() string {
	return r.raw
}

func (r ToolSourceAuthNoToolSourceAuth) implementsToolSourceAuth() {}

type ToolSourceAuthNoToolSourceAuthKind string

const (
	ToolSourceAuthNoToolSourceAuthKindNone ToolSourceAuthNoToolSourceAuthKind = "none"
)

func (r ToolSourceAuthNoToolSourceAuthKind) IsKnown() bool {
	switch r {
	case ToolSourceAuthNoToolSourceAuthKindNone:
		return true
	}
	return false
}

type ToolSourceAuthAPIKeyToolSourceAuth struct {
	ConnectionID string `json:"connectionId" api:"required"`
	// Where an API key or OAuth access token should be injected into outgoing
	// tool-source requests.
	Inject AuthInjection                          `json:"inject" api:"required"`
	Kind   ToolSourceAuthAPIKeyToolSourceAuthKind `json:"kind" api:"required"`
	JSON   toolSourceAuthAPIKeyToolSourceAuthJSON `json:"-"`
}

// toolSourceAuthAPIKeyToolSourceAuthJSON contains the JSON metadata for the struct
// [ToolSourceAuthAPIKeyToolSourceAuth]
type toolSourceAuthAPIKeyToolSourceAuthJSON struct {
	ConnectionID apijson.Field
	Inject       apijson.Field
	Kind         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ToolSourceAuthAPIKeyToolSourceAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSourceAuthAPIKeyToolSourceAuthJSON) RawJSON() string {
	return r.raw
}

func (r ToolSourceAuthAPIKeyToolSourceAuth) implementsToolSourceAuth() {}

type ToolSourceAuthAPIKeyToolSourceAuthKind string

const (
	ToolSourceAuthAPIKeyToolSourceAuthKindAPIKey ToolSourceAuthAPIKeyToolSourceAuthKind = "api_key"
)

func (r ToolSourceAuthAPIKeyToolSourceAuthKind) IsKnown() bool {
	switch r {
	case ToolSourceAuthAPIKeyToolSourceAuthKindAPIKey:
		return true
	}
	return false
}

type ToolSourceAuthOAuthExchangeToolSourceAuth struct {
	// Source-owned OAuth exchange configuration. Public schemas use connection IDs for
	// stored client secrets and assertions.
	Exchange OAuthExchangeConfig `json:"exchange" api:"required"`
	// Where an API key or OAuth access token should be injected into outgoing
	// tool-source requests.
	Inject AuthInjection                                 `json:"inject" api:"required"`
	Kind   ToolSourceAuthOAuthExchangeToolSourceAuthKind `json:"kind" api:"required"`
	JSON   toolSourceAuthOAuthExchangeToolSourceAuthJSON `json:"-"`
}

// toolSourceAuthOAuthExchangeToolSourceAuthJSON contains the JSON metadata for the
// struct [ToolSourceAuthOAuthExchangeToolSourceAuth]
type toolSourceAuthOAuthExchangeToolSourceAuthJSON struct {
	Exchange    apijson.Field
	Inject      apijson.Field
	Kind        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSourceAuthOAuthExchangeToolSourceAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSourceAuthOAuthExchangeToolSourceAuthJSON) RawJSON() string {
	return r.raw
}

func (r ToolSourceAuthOAuthExchangeToolSourceAuth) implementsToolSourceAuth() {}

type ToolSourceAuthOAuthExchangeToolSourceAuthKind string

const (
	ToolSourceAuthOAuthExchangeToolSourceAuthKindOAuthExchange ToolSourceAuthOAuthExchangeToolSourceAuthKind = "oauth_exchange"
)

func (r ToolSourceAuthOAuthExchangeToolSourceAuthKind) IsKnown() bool {
	switch r {
	case ToolSourceAuthOAuthExchangeToolSourceAuthKindOAuthExchange:
		return true
	}
	return false
}

type ToolSourceAuthOAuthConnectionToolSourceAuth struct {
	Kind     ToolSourceAuthOAuthConnectionToolSourceAuthKind `json:"kind" api:"required"`
	Provider string                                          `json:"provider" api:"required"`
	// Where an API key or OAuth access token should be injected into outgoing
	// tool-source requests.
	Inject         AuthInjection                                   `json:"inject"`
	RequiredScopes []string                                        `json:"requiredScopes"`
	JSON           toolSourceAuthOAuthConnectionToolSourceAuthJSON `json:"-"`
}

// toolSourceAuthOAuthConnectionToolSourceAuthJSON contains the JSON metadata for
// the struct [ToolSourceAuthOAuthConnectionToolSourceAuth]
type toolSourceAuthOAuthConnectionToolSourceAuthJSON struct {
	Kind           apijson.Field
	Provider       apijson.Field
	Inject         apijson.Field
	RequiredScopes apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ToolSourceAuthOAuthConnectionToolSourceAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSourceAuthOAuthConnectionToolSourceAuthJSON) RawJSON() string {
	return r.raw
}

func (r ToolSourceAuthOAuthConnectionToolSourceAuth) implementsToolSourceAuth() {}

type ToolSourceAuthOAuthConnectionToolSourceAuthKind string

const (
	ToolSourceAuthOAuthConnectionToolSourceAuthKindOAuthConnection ToolSourceAuthOAuthConnectionToolSourceAuthKind = "oauth_connection"
)

func (r ToolSourceAuthOAuthConnectionToolSourceAuthKind) IsKnown() bool {
	switch r {
	case ToolSourceAuthOAuthConnectionToolSourceAuthKindOAuthConnection:
		return true
	}
	return false
}

type ToolSourceAuthKind string

const (
	ToolSourceAuthKindNone            ToolSourceAuthKind = "none"
	ToolSourceAuthKindAPIKey          ToolSourceAuthKind = "api_key"
	ToolSourceAuthKindOAuthExchange   ToolSourceAuthKind = "oauth_exchange"
	ToolSourceAuthKindOAuthConnection ToolSourceAuthKind = "oauth_connection"
)

func (r ToolSourceAuthKind) IsKnown() bool {
	switch r {
	case ToolSourceAuthKindNone, ToolSourceAuthKindAPIKey, ToolSourceAuthKindOAuthExchange, ToolSourceAuthKindOAuthConnection:
		return true
	}
	return false
}

// Tool source authentication configuration. The `kind` field selects one of
// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
type ToolSourceAuthParam struct {
	Kind         param.Field[ToolSourceAuthKind] `json:"kind" api:"required"`
	ConnectionID param.Field[string]             `json:"connectionId"`
	// Source-owned OAuth exchange configuration. Public schemas use connection IDs for
	// stored client secrets and assertions.
	Exchange param.Field[OAuthExchangeConfigUnionParam] `json:"exchange"`
	// Where an API key or OAuth access token should be injected into outgoing
	// tool-source requests.
	Inject         param.Field[AuthInjectionUnionParam] `json:"inject"`
	Provider       param.Field[string]                  `json:"provider"`
	RequiredScopes param.Field[interface{}]             `json:"requiredScopes"`
}

func (r ToolSourceAuthParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolSourceAuthParam) implementsToolSourceAuthUnionParam() {}

// Tool source authentication configuration. The `kind` field selects one of
// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
//
// Satisfied by [ToolSourceAuthNoToolSourceAuthParam],
// [ToolSourceAuthAPIKeyToolSourceAuthParam],
// [ToolSourceAuthOAuthExchangeToolSourceAuthParam],
// [ToolSourceAuthOAuthConnectionToolSourceAuthParam], [ToolSourceAuthParam].
type ToolSourceAuthUnionParam interface {
	implementsToolSourceAuthUnionParam()
}

type ToolSourceAuthNoToolSourceAuthParam struct {
	Kind param.Field[ToolSourceAuthNoToolSourceAuthKind] `json:"kind" api:"required"`
}

func (r ToolSourceAuthNoToolSourceAuthParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolSourceAuthNoToolSourceAuthParam) implementsToolSourceAuthUnionParam() {}

type ToolSourceAuthAPIKeyToolSourceAuthParam struct {
	ConnectionID param.Field[string] `json:"connectionId" api:"required"`
	// Where an API key or OAuth access token should be injected into outgoing
	// tool-source requests.
	Inject param.Field[AuthInjectionUnionParam]                `json:"inject" api:"required"`
	Kind   param.Field[ToolSourceAuthAPIKeyToolSourceAuthKind] `json:"kind" api:"required"`
}

func (r ToolSourceAuthAPIKeyToolSourceAuthParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolSourceAuthAPIKeyToolSourceAuthParam) implementsToolSourceAuthUnionParam() {}

type ToolSourceAuthOAuthExchangeToolSourceAuthParam struct {
	// Source-owned OAuth exchange configuration. Public schemas use connection IDs for
	// stored client secrets and assertions.
	Exchange param.Field[OAuthExchangeConfigUnionParam] `json:"exchange" api:"required"`
	// Where an API key or OAuth access token should be injected into outgoing
	// tool-source requests.
	Inject param.Field[AuthInjectionUnionParam]                       `json:"inject" api:"required"`
	Kind   param.Field[ToolSourceAuthOAuthExchangeToolSourceAuthKind] `json:"kind" api:"required"`
}

func (r ToolSourceAuthOAuthExchangeToolSourceAuthParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolSourceAuthOAuthExchangeToolSourceAuthParam) implementsToolSourceAuthUnionParam() {}

type ToolSourceAuthOAuthConnectionToolSourceAuthParam struct {
	Kind     param.Field[ToolSourceAuthOAuthConnectionToolSourceAuthKind] `json:"kind" api:"required"`
	Provider param.Field[string]                                          `json:"provider" api:"required"`
	// Where an API key or OAuth access token should be injected into outgoing
	// tool-source requests.
	Inject         param.Field[AuthInjectionUnionParam] `json:"inject"`
	RequiredScopes param.Field[[]string]                `json:"requiredScopes"`
}

func (r ToolSourceAuthOAuthConnectionToolSourceAuthParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolSourceAuthOAuthConnectionToolSourceAuthParam) implementsToolSourceAuthUnionParam() {}

type ToolNewParams struct {
	Tool ToolNewParamsToolUnion `json:"tool" api:"required"`
}

func (r ToolNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Tool)
}

type ToolNewParamsTool struct {
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      param.Field[ToolSourceAuthUnionParam] `json:"auth" api:"required"`
	Namespace param.Field[string]                   `json:"namespace" api:"required"`
	Type      param.Field[ToolNewParamsToolType]    `json:"type" api:"required"`
	Approval  param.Field[ToolApprovalMode]         `json:"approval"`
	Enabled   param.Field[bool]                     `json:"enabled"`
	Execution param.Field[interface{}]              `json:"execution"`
	Tools     param.Field[interface{}]              `json:"tools"`
	URL       param.Field[string]                   `json:"url"`
}

func (r ToolNewParamsTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolNewParamsTool) implementsToolNewParamsToolUnion() {}

// Satisfied by [ToolNewParamsToolCreateHTTPToolSourceRequest],
// [ToolNewParamsToolCreateMcpToolSourceRequest], [ToolNewParamsTool].
type ToolNewParamsToolUnion interface {
	implementsToolNewParamsToolUnion()
}

type ToolNewParamsToolCreateHTTPToolSourceRequest struct {
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      param.Field[ToolSourceAuthUnionParam]                         `json:"auth" api:"required"`
	Namespace param.Field[string]                                           `json:"namespace" api:"required"`
	Tools     param.Field[[]HTTPToolDefinitionParam]                        `json:"tools" api:"required"`
	Type      param.Field[ToolNewParamsToolCreateHTTPToolSourceRequestType] `json:"type" api:"required"`
	Approval  param.Field[ToolApprovalMode]                                 `json:"approval"`
	Enabled   param.Field[bool]                                             `json:"enabled"`
	// HTTP tool execution retry and timeout policy.
	Execution   param.Field[HTTPToolExecutionPolicyParam] `json:"execution"`
	ExtraFields map[string]interface{}                    `json:"-,extras"`
}

func (r ToolNewParamsToolCreateHTTPToolSourceRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolNewParamsToolCreateHTTPToolSourceRequest) implementsToolNewParamsToolUnion() {}

type ToolNewParamsToolCreateHTTPToolSourceRequestType string

const (
	ToolNewParamsToolCreateHTTPToolSourceRequestTypeHTTP ToolNewParamsToolCreateHTTPToolSourceRequestType = "http"
)

func (r ToolNewParamsToolCreateHTTPToolSourceRequestType) IsKnown() bool {
	switch r {
	case ToolNewParamsToolCreateHTTPToolSourceRequestTypeHTTP:
		return true
	}
	return false
}

type ToolNewParamsToolCreateMcpToolSourceRequest struct {
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      param.Field[ToolSourceAuthUnionParam]                        `json:"auth" api:"required"`
	Namespace param.Field[string]                                          `json:"namespace" api:"required"`
	Type      param.Field[ToolNewParamsToolCreateMcpToolSourceRequestType] `json:"type" api:"required"`
	URL       param.Field[string]                                          `json:"url" api:"required"`
	Approval  param.Field[ToolApprovalMode]                                `json:"approval"`
	Enabled   param.Field[bool]                                            `json:"enabled"`
	// MCP discovery and execution retry/timeout policy.
	Execution   param.Field[McpToolExecutionPolicyParam] `json:"execution"`
	ExtraFields map[string]interface{}                   `json:"-,extras"`
}

func (r ToolNewParamsToolCreateMcpToolSourceRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolNewParamsToolCreateMcpToolSourceRequest) implementsToolNewParamsToolUnion() {}

type ToolNewParamsToolCreateMcpToolSourceRequestType string

const (
	ToolNewParamsToolCreateMcpToolSourceRequestTypeMcp ToolNewParamsToolCreateMcpToolSourceRequestType = "mcp"
)

func (r ToolNewParamsToolCreateMcpToolSourceRequestType) IsKnown() bool {
	switch r {
	case ToolNewParamsToolCreateMcpToolSourceRequestTypeMcp:
		return true
	}
	return false
}

type ToolNewParamsToolType string

const (
	ToolNewParamsToolTypeHTTP ToolNewParamsToolType = "http"
	ToolNewParamsToolTypeMcp  ToolNewParamsToolType = "mcp"
)

func (r ToolNewParamsToolType) IsKnown() bool {
	switch r {
	case ToolNewParamsToolTypeHTTP, ToolNewParamsToolTypeMcp:
		return true
	}
	return false
}

type ToolUpdateParams struct {
	Tool ToolUpdateParamsToolUnion `json:"tool" api:"required"`
}

func (r ToolUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Tool)
}

type ToolUpdateParamsTool struct {
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      param.Field[ToolSourceAuthUnionParam] `json:"auth" api:"required"`
	Namespace param.Field[string]                   `json:"namespace" api:"required"`
	Type      param.Field[ToolUpdateParamsToolType] `json:"type" api:"required"`
	Approval  param.Field[ToolApprovalMode]         `json:"approval"`
	Enabled   param.Field[bool]                     `json:"enabled"`
	Execution param.Field[interface{}]              `json:"execution"`
	Tools     param.Field[interface{}]              `json:"tools"`
	URL       param.Field[string]                   `json:"url"`
}

func (r ToolUpdateParamsTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolUpdateParamsTool) implementsToolUpdateParamsToolUnion() {}

// Satisfied by [ToolUpdateParamsToolUpdateHTTPToolSourceRequest],
// [ToolUpdateParamsToolUpdateMcpToolSourceRequest], [ToolUpdateParamsTool].
type ToolUpdateParamsToolUnion interface {
	implementsToolUpdateParamsToolUnion()
}

type ToolUpdateParamsToolUpdateHTTPToolSourceRequest struct {
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      param.Field[ToolSourceAuthUnionParam]                            `json:"auth" api:"required"`
	Namespace param.Field[string]                                              `json:"namespace" api:"required"`
	Tools     param.Field[[]HTTPToolDefinitionParam]                           `json:"tools" api:"required"`
	Type      param.Field[ToolUpdateParamsToolUpdateHTTPToolSourceRequestType] `json:"type" api:"required"`
	Approval  param.Field[ToolApprovalMode]                                    `json:"approval"`
	Enabled   param.Field[bool]                                                `json:"enabled"`
	// HTTP tool execution retry and timeout policy.
	Execution   param.Field[HTTPToolExecutionPolicyParam] `json:"execution"`
	ExtraFields map[string]interface{}                    `json:"-,extras"`
}

func (r ToolUpdateParamsToolUpdateHTTPToolSourceRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolUpdateParamsToolUpdateHTTPToolSourceRequest) implementsToolUpdateParamsToolUnion() {}

type ToolUpdateParamsToolUpdateHTTPToolSourceRequestType string

const (
	ToolUpdateParamsToolUpdateHTTPToolSourceRequestTypeHTTP ToolUpdateParamsToolUpdateHTTPToolSourceRequestType = "http"
)

func (r ToolUpdateParamsToolUpdateHTTPToolSourceRequestType) IsKnown() bool {
	switch r {
	case ToolUpdateParamsToolUpdateHTTPToolSourceRequestTypeHTTP:
		return true
	}
	return false
}

type ToolUpdateParamsToolUpdateMcpToolSourceRequest struct {
	// Tool source authentication configuration. The `kind` field selects one of
	// `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.
	Auth      param.Field[ToolSourceAuthUnionParam]                           `json:"auth" api:"required"`
	Namespace param.Field[string]                                             `json:"namespace" api:"required"`
	Type      param.Field[ToolUpdateParamsToolUpdateMcpToolSourceRequestType] `json:"type" api:"required"`
	URL       param.Field[string]                                             `json:"url" api:"required"`
	Approval  param.Field[ToolApprovalMode]                                   `json:"approval"`
	Enabled   param.Field[bool]                                               `json:"enabled"`
	// MCP discovery and execution retry/timeout policy.
	Execution   param.Field[McpToolExecutionPolicyParam] `json:"execution"`
	ExtraFields map[string]interface{}                   `json:"-,extras"`
}

func (r ToolUpdateParamsToolUpdateMcpToolSourceRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolUpdateParamsToolUpdateMcpToolSourceRequest) implementsToolUpdateParamsToolUnion() {}

type ToolUpdateParamsToolUpdateMcpToolSourceRequestType string

const (
	ToolUpdateParamsToolUpdateMcpToolSourceRequestTypeMcp ToolUpdateParamsToolUpdateMcpToolSourceRequestType = "mcp"
)

func (r ToolUpdateParamsToolUpdateMcpToolSourceRequestType) IsKnown() bool {
	switch r {
	case ToolUpdateParamsToolUpdateMcpToolSourceRequestTypeMcp:
		return true
	}
	return false
}

type ToolUpdateParamsToolType string

const (
	ToolUpdateParamsToolTypeHTTP ToolUpdateParamsToolType = "http"
	ToolUpdateParamsToolTypeMcp  ToolUpdateParamsToolType = "mcp"
)

func (r ToolUpdateParamsToolType) IsKnown() bool {
	switch r {
	case ToolUpdateParamsToolTypeHTTP, ToolUpdateParamsToolTypeMcp:
		return true
	}
	return false
}

type ToolListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [ToolListParams]'s query parameters as `url.Values`.
func (r ToolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
