// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cercago

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/matrices/cerca-go/internal/apijson"
	"github.com/matrices/cerca-go/internal/apiquery"
	"github.com/matrices/cerca-go/internal/param"
	"github.com/matrices/cerca-go/internal/requestconfig"
	"github.com/matrices/cerca-go/option"
)

// CredentialService contains methods and other services that help with interacting
// with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCredentialService] method instead.
type CredentialService struct {
	Options []option.RequestOption
}

// NewCredentialService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCredentialService(opts ...option.RequestOption) (r *CredentialService) {
	r = &CredentialService{}
	r.Options = opts
	return
}

func (r *CredentialService) List(ctx context.Context, scope string, query CredentialListParams, opts ...option.RequestOption) (res *CredentialListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if scope == "" {
		err = errors.New("missing required scope parameter")
		return nil, err
	}
	path := fmt.Sprintf("credentials/%s", scope)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *CredentialService) Delete(ctx context.Context, scope string, connectionID string, opts ...option.RequestOption) (res *CredentialDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if scope == "" {
		err = errors.New("missing required scope parameter")
		return nil, err
	}
	if connectionID == "" {
		err = errors.New("missing required connectionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("credentials/%s/%s", scope, connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

func (r *CredentialService) NewAPIKey(ctx context.Context, scope string, body CredentialNewAPIKeyParams, opts ...option.RequestOption) (res *Connection, err error) {
	opts = slices.Concat(r.Options, opts)
	if scope == "" {
		err = errors.New("missing required scope parameter")
		return nil, err
	}
	path := fmt.Sprintf("credentials/%s/api-keys", scope)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type Connection struct {
	ID           string             `json:"id" api:"required"`
	AccountLabel string             `json:"accountLabel" api:"required,nullable"`
	CreatedAt    string             `json:"createdAt" api:"required"`
	Provider     CredentialProvider `json:"provider" api:"required"`
	// Reusable credential connection scope.
	Scope     string         `json:"scope" api:"required"`
	Scopes    []string       `json:"scopes" api:"required"`
	Type      CredentialType `json:"type" api:"required"`
	UpdatedAt string         `json:"updatedAt" api:"required"`
	JSON      connectionJSON `json:"-"`
}

// connectionJSON contains the JSON metadata for the struct [Connection]
type connectionJSON struct {
	ID           apijson.Field
	AccountLabel apijson.Field
	CreatedAt    apijson.Field
	Provider     apijson.Field
	Scope        apijson.Field
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

type CredentialListResponse struct {
	Connections []Connection               `json:"connections" api:"required"`
	Cursor      string                     `json:"cursor" api:"required,nullable"`
	HasMore     bool                       `json:"hasMore" api:"required"`
	JSON        credentialListResponseJSON `json:"-"`
}

// credentialListResponseJSON contains the JSON metadata for the struct
// [CredentialListResponse]
type credentialListResponseJSON struct {
	Connections apijson.Field
	Cursor      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CredentialListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r credentialListResponseJSON) RawJSON() string {
	return r.raw
}

type CredentialDeleteResponse struct {
	Success CredentialDeleteResponseSuccess `json:"success" api:"required"`
	JSON    credentialDeleteResponseJSON    `json:"-"`
}

// credentialDeleteResponseJSON contains the JSON metadata for the struct
// [CredentialDeleteResponse]
type credentialDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CredentialDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r credentialDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CredentialDeleteResponseSuccess bool

const (
	CredentialDeleteResponseSuccessTrue CredentialDeleteResponseSuccess = true
)

func (r CredentialDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case CredentialDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type CredentialListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [CredentialListParams]'s query parameters as `url.Values`.
func (r CredentialListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CredentialNewAPIKeyParams struct {
	// API key secret. It is stored securely and is not returned.
	APIKey param.Field[string] `json:"apiKey" api:"required"`
	// Credential provider to store an API key for.
	Provider param.Field[string] `json:"provider" api:"required"`
	// Optional human-readable account label.
	AccountLabel param.Field[string] `json:"accountLabel"`
}

func (r CredentialNewAPIKeyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
