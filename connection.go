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

// ConnectionService contains methods and other services that help with interacting
// with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectionService] method instead.
type ConnectionService struct {
	Options []option.RequestOption
}

// NewConnectionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectionService(opts ...option.RequestOption) (r *ConnectionService) {
	r = &ConnectionService{}
	r.Options = opts
	return
}

// API Keys
func (r *ConnectionService) New(ctx context.Context, body ConnectionNewParams, opts ...option.RequestOption) (res *Connection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "connections/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Connections
func (r *ConnectionService) List(ctx context.Context, query ConnectionListParams, opts ...option.RequestOption) (res *pagination.ConnectionsCursorPage[Connection], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "connections"
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

// Connections
func (r *ConnectionService) ListAutoPaging(ctx context.Context, query ConnectionListParams, opts ...option.RequestOption) *pagination.ConnectionsCursorPageAutoPager[Connection] {
	return pagination.NewConnectionsCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Connection
func (r *ConnectionService) Delete(ctx context.Context, connectionID string, body ConnectionDeleteParams, opts ...option.RequestOption) (res *ConnectionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if connectionID == "" {
		err = errors.New("missing required connectionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("connections/%s", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Connections
func (r *ConnectionService) Attach(ctx context.Context, agentID string, body ConnectionAttachParams, opts ...option.RequestOption) (res *AttachedConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/connections", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Connection
func (r *ConnectionService) Detach(ctx context.Context, agentID string, connectionID string, opts ...option.RequestOption) (res *ConnectionDetachResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if connectionID == "" {
		err = errors.New("missing required connectionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/connections/%s", agentID, connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Connections
func (r *ConnectionService) ListForAgent(ctx context.Context, agentID string, opts ...option.RequestOption) (res *ConnectionListForAgentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/connections", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AttachedConnection struct {
	AttachedAt string                 `json:"attachedAt" api:"required"`
	Metadata   ToolConnectionMetadata `json:"metadata"`
	JSON       attachedConnectionJSON `json:"-"`
	Connection
}

// attachedConnectionJSON contains the JSON metadata for the struct
// [AttachedConnection]
type attachedConnectionJSON struct {
	AttachedAt  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttachedConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attachedConnectionJSON) RawJSON() string {
	return r.raw
}

type Connection struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	AccountLabel string `json:"accountLabel" api:"required,nullable"`
	CreatedAt    string `json:"createdAt" api:"required"`
	// Public owner for a reusable connection. Organization owners use the
	// authenticated organization; fleet owners add a fleetId.
	Owner     ConnectionOwner    `json:"owner" api:"required"`
	Provider  CredentialProvider `json:"provider" api:"required"`
	Scopes    []string           `json:"scopes" api:"required"`
	Type      CredentialType     `json:"type" api:"required"`
	UpdatedAt string             `json:"updatedAt" api:"required"`
	JSON      connectionJSON     `json:"-"`
}

// connectionJSON contains the JSON metadata for the struct [Connection]
type connectionJSON struct {
	ID           apijson.Field
	AccountLabel apijson.Field
	CreatedAt    apijson.Field
	Owner        apijson.Field
	Provider     apijson.Field
	Scopes       apijson.Field
	Type         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Connection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionJSON) RawJSON() string {
	return r.raw
}

// Public owner for a reusable connection. Organization owners use the
// authenticated organization; fleet owners add a fleetId.
type ConnectionOwner struct {
	Type    ConnectionOwnerType `json:"type" api:"required"`
	FleetID string              `json:"fleetId"`
	JSON    connectionOwnerJSON `json:"-"`
	union   ConnectionOwnerUnion
}

// connectionOwnerJSON contains the JSON metadata for the struct [ConnectionOwner]
type connectionOwnerJSON struct {
	Type        apijson.Field
	FleetID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r connectionOwnerJSON) RawJSON() string {
	return r.raw
}

func (r *ConnectionOwner) UnmarshalJSON(data []byte) (err error) {
	*r = ConnectionOwner{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConnectionOwnerUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [ConnectionOwnerOrganizationConnectionOwner],
// [ConnectionOwnerFleetConnectionOwner].
func (r ConnectionOwner) AsUnion() ConnectionOwnerUnion {
	return r.union
}

// Public owner for a reusable connection. Organization owners use the
// authenticated organization; fleet owners add a fleetId.
//
// Union satisfied by [ConnectionOwnerOrganizationConnectionOwner] or
// [ConnectionOwnerFleetConnectionOwner].
type ConnectionOwnerUnion interface {
	implementsConnectionOwner()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConnectionOwnerUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectionOwnerOrganizationConnectionOwner{}),
			DiscriminatorValue: "organization",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectionOwnerFleetConnectionOwner{}),
			DiscriminatorValue: "fleet",
		},
	)
}

type ConnectionOwnerOrganizationConnectionOwner struct {
	Type ConnectionOwnerOrganizationConnectionOwnerType `json:"type" api:"required"`
	JSON connectionOwnerOrganizationConnectionOwnerJSON `json:"-"`
}

// connectionOwnerOrganizationConnectionOwnerJSON contains the JSON metadata for
// the struct [ConnectionOwnerOrganizationConnectionOwner]
type connectionOwnerOrganizationConnectionOwnerJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionOwnerOrganizationConnectionOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionOwnerOrganizationConnectionOwnerJSON) RawJSON() string {
	return r.raw
}

func (r ConnectionOwnerOrganizationConnectionOwner) implementsConnectionOwner() {}

type ConnectionOwnerOrganizationConnectionOwnerType string

const (
	ConnectionOwnerOrganizationConnectionOwnerTypeOrganization ConnectionOwnerOrganizationConnectionOwnerType = "organization"
)

func (r ConnectionOwnerOrganizationConnectionOwnerType) IsKnown() bool {
	switch r {
	case ConnectionOwnerOrganizationConnectionOwnerTypeOrganization:
		return true
	}
	return false
}

type ConnectionOwnerFleetConnectionOwner struct {
	FleetID string                                  `json:"fleetId" api:"required"`
	Type    ConnectionOwnerFleetConnectionOwnerType `json:"type" api:"required"`
	JSON    connectionOwnerFleetConnectionOwnerJSON `json:"-"`
}

// connectionOwnerFleetConnectionOwnerJSON contains the JSON metadata for the
// struct [ConnectionOwnerFleetConnectionOwner]
type connectionOwnerFleetConnectionOwnerJSON struct {
	FleetID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionOwnerFleetConnectionOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionOwnerFleetConnectionOwnerJSON) RawJSON() string {
	return r.raw
}

func (r ConnectionOwnerFleetConnectionOwner) implementsConnectionOwner() {}

type ConnectionOwnerFleetConnectionOwnerType string

const (
	ConnectionOwnerFleetConnectionOwnerTypeFleet ConnectionOwnerFleetConnectionOwnerType = "fleet"
)

func (r ConnectionOwnerFleetConnectionOwnerType) IsKnown() bool {
	switch r {
	case ConnectionOwnerFleetConnectionOwnerTypeFleet:
		return true
	}
	return false
}

type ConnectionOwnerType string

const (
	ConnectionOwnerTypeOrganization ConnectionOwnerType = "organization"
	ConnectionOwnerTypeFleet        ConnectionOwnerType = "fleet"
)

func (r ConnectionOwnerType) IsKnown() bool {
	switch r {
	case ConnectionOwnerTypeOrganization, ConnectionOwnerTypeFleet:
		return true
	}
	return false
}

// Public owner for a reusable connection. Organization owners use the
// authenticated organization; fleet owners add a fleetId.
type ConnectionOwnerParam struct {
	Type    param.Field[ConnectionOwnerType] `json:"type" api:"required"`
	FleetID param.Field[string]              `json:"fleetId"`
}

func (r ConnectionOwnerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConnectionOwnerParam) implementsConnectionOwnerUnionParam() {}

// Public owner for a reusable connection. Organization owners use the
// authenticated organization; fleet owners add a fleetId.
//
// Satisfied by [ConnectionOwnerOrganizationConnectionOwnerParam],
// [ConnectionOwnerFleetConnectionOwnerParam], [ConnectionOwnerParam].
type ConnectionOwnerUnionParam interface {
	implementsConnectionOwnerUnionParam()
}

type ConnectionOwnerOrganizationConnectionOwnerParam struct {
	Type param.Field[ConnectionOwnerOrganizationConnectionOwnerType] `json:"type" api:"required"`
}

func (r ConnectionOwnerOrganizationConnectionOwnerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConnectionOwnerOrganizationConnectionOwnerParam) implementsConnectionOwnerUnionParam() {}

type ConnectionOwnerFleetConnectionOwnerParam struct {
	FleetID param.Field[string]                                  `json:"fleetId" api:"required"`
	Type    param.Field[ConnectionOwnerFleetConnectionOwnerType] `json:"type" api:"required"`
}

func (r ConnectionOwnerFleetConnectionOwnerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConnectionOwnerFleetConnectionOwnerParam) implementsConnectionOwnerUnionParam() {}

type CredentialProvider string

const (
	CredentialProviderGoogle  CredentialProvider = "google"
	CredentialProviderGitHub  CredentialProvider = "github"
	CredentialProviderSlack   CredentialProvider = "slack"
	CredentialProviderLinear  CredentialProvider = "linear"
	CredentialProviderNotion  CredentialProvider = "notion"
	CredentialProviderCustom  CredentialProvider = "custom"
	CredentialProviderWebhook CredentialProvider = "webhook"
)

func (r CredentialProvider) IsKnown() bool {
	switch r {
	case CredentialProviderGoogle, CredentialProviderGitHub, CredentialProviderSlack, CredentialProviderLinear, CredentialProviderNotion, CredentialProviderCustom, CredentialProviderWebhook:
		return true
	}
	return false
}

type CredentialType string

const (
	CredentialTypeOAuth  CredentialType = "oauth"
	CredentialTypeAPIKey CredentialType = "api_key"
)

func (r CredentialType) IsKnown() bool {
	switch r {
	case CredentialTypeOAuth, CredentialTypeAPIKey:
		return true
	}
	return false
}

type ToolConnectionMetadata map[string]string

type ToolConnectionMetadataParam map[string]string

type ConnectionDeleteResponse struct {
	Success ConnectionDeleteResponseSuccess `json:"success" api:"required"`
	JSON    connectionDeleteResponseJSON    `json:"-"`
}

// connectionDeleteResponseJSON contains the JSON metadata for the struct
// [ConnectionDeleteResponse]
type connectionDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectionDeleteResponseSuccess bool

const (
	ConnectionDeleteResponseSuccessTrue ConnectionDeleteResponseSuccess = true
)

func (r ConnectionDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ConnectionDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ConnectionDetachResponse struct {
	Success ConnectionDetachResponseSuccess `json:"success" api:"required"`
	JSON    connectionDetachResponseJSON    `json:"-"`
}

// connectionDetachResponseJSON contains the JSON metadata for the struct
// [ConnectionDetachResponse]
type connectionDetachResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionDetachResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionDetachResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectionDetachResponseSuccess bool

const (
	ConnectionDetachResponseSuccessTrue ConnectionDetachResponseSuccess = true
)

func (r ConnectionDetachResponseSuccess) IsKnown() bool {
	switch r {
	case ConnectionDetachResponseSuccessTrue:
		return true
	}
	return false
}

type ConnectionListForAgentResponse struct {
	Connections []AttachedConnection               `json:"connections" api:"required"`
	Cursor      string                             `json:"cursor" api:"required,nullable"`
	HasMore     bool                               `json:"hasMore" api:"required"`
	JSON        connectionListForAgentResponseJSON `json:"-"`
}

// connectionListForAgentResponseJSON contains the JSON metadata for the struct
// [ConnectionListForAgentResponse]
type connectionListForAgentResponseJSON struct {
	Connections apijson.Field
	Cursor      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionListForAgentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionListForAgentResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectionNewParams struct {
	// API key secret. It is stored securely and is not returned.
	APIKey param.Field[string] `json:"apiKey" api:"required"`
	// Public owner for a reusable connection. Organization owners use the
	// authenticated organization; fleet owners add a fleetId.
	Owner param.Field[ConnectionOwnerUnionParam] `json:"owner" api:"required"`
	// Credential provider to store an API key for.
	Provider param.Field[string] `json:"provider" api:"required"`
	// Optional human-readable account label.
	AccountLabel param.Field[string] `json:"accountLabel"`
}

func (r ConnectionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectionListParams struct {
	// Public connection owner type.
	OwnerType param.Field[ConnectionListParamsOwnerType] `query:"ownerType" api:"required"`
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Required when ownerType is fleet.
	FleetID param.Field[string] `query:"fleetId"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [ConnectionListParams]'s query parameters as `url.Values`.
func (r ConnectionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Public connection owner type.
type ConnectionListParamsOwnerType string

const (
	ConnectionListParamsOwnerTypeOrganization ConnectionListParamsOwnerType = "organization"
	ConnectionListParamsOwnerTypeFleet        ConnectionListParamsOwnerType = "fleet"
)

func (r ConnectionListParamsOwnerType) IsKnown() bool {
	switch r {
	case ConnectionListParamsOwnerTypeOrganization, ConnectionListParamsOwnerTypeFleet:
		return true
	}
	return false
}

type ConnectionDeleteParams struct {
	// Public connection owner type.
	OwnerType param.Field[ConnectionDeleteParamsOwnerType] `query:"ownerType" api:"required"`
	// Required when ownerType is fleet.
	FleetID param.Field[string] `query:"fleetId"`
}

// URLQuery serializes [ConnectionDeleteParams]'s query parameters as `url.Values`.
func (r ConnectionDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Public connection owner type.
type ConnectionDeleteParamsOwnerType string

const (
	ConnectionDeleteParamsOwnerTypeOrganization ConnectionDeleteParamsOwnerType = "organization"
	ConnectionDeleteParamsOwnerTypeFleet        ConnectionDeleteParamsOwnerType = "fleet"
)

func (r ConnectionDeleteParamsOwnerType) IsKnown() bool {
	switch r {
	case ConnectionDeleteParamsOwnerTypeOrganization, ConnectionDeleteParamsOwnerTypeFleet:
		return true
	}
	return false
}

type ConnectionAttachParams struct {
	ConnectionID param.Field[string]                      `json:"connectionId" api:"required" format:"uuid"`
	Metadata     param.Field[ToolConnectionMetadataParam] `json:"metadata"`
}

func (r ConnectionAttachParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
