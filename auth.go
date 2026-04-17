// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cercago

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/matrices/cerca-go/internal/apijson"
	"github.com/matrices/cerca-go/internal/apiquery"
	"github.com/matrices/cerca-go/internal/param"
	"github.com/matrices/cerca-go/internal/requestconfig"
	"github.com/matrices/cerca-go/option"
)

// AuthService contains methods and other services that help with interacting with
// the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthService] method instead.
type AuthService struct {
	Options []option.RequestOption
}

// NewAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthService(opts ...option.RequestOption) (r *AuthService) {
	r = &AuthService{}
	r.Options = opts
	return
}

func (r *AuthService) Context(ctx context.Context, opts ...option.RequestOption) (res *AuthContextResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/context"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *AuthService) ListEnvironments(ctx context.Context, query AuthListEnvironmentsParams, opts ...option.RequestOption) (res *AuthListEnvironmentsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/environments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AuthContextResponse struct {
	KeyID string                  `json:"keyId" api:"required"`
	OrgID string                  `json:"orgId" api:"required"`
	JSON  authContextResponseJSON `json:"-"`
}

// authContextResponseJSON contains the JSON metadata for the struct
// [AuthContextResponse]
type authContextResponseJSON struct {
	KeyID       apijson.Field
	OrgID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthContextResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authContextResponseJSON) RawJSON() string {
	return r.raw
}

type AuthListEnvironmentsResponse struct {
	Cursor       string                           `json:"cursor" api:"required,nullable"`
	Environments []Environment                    `json:"environments" api:"required"`
	HasMore      bool                             `json:"hasMore" api:"required"`
	JSON         authListEnvironmentsResponseJSON `json:"-"`
}

// authListEnvironmentsResponseJSON contains the JSON metadata for the struct
// [AuthListEnvironmentsResponse]
type authListEnvironmentsResponseJSON struct {
	Cursor       apijson.Field
	Environments apijson.Field
	HasMore      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AuthListEnvironmentsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authListEnvironmentsResponseJSON) RawJSON() string {
	return r.raw
}

type AuthListEnvironmentsParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [AuthListEnvironmentsParams]'s query parameters as
// `url.Values`.
func (r AuthListEnvironmentsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
