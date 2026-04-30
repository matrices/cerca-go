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
	"github.com/matrices/cerca-go/packages/pagination"
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

// Retrieve context
func (r *AuthService) Context(ctx context.Context, opts ...option.RequestOption) (res *AuthContextResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/context"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List fleets
func (r *AuthService) ListFleets(ctx context.Context, query AuthListFleetsParams, opts ...option.RequestOption) (res *pagination.FleetsCursorPage[Fleet], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "auth/fleets"
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

// List fleets
func (r *AuthService) ListFleetsAutoPaging(ctx context.Context, query AuthListFleetsParams, opts ...option.RequestOption) *pagination.FleetsCursorPageAutoPager[Fleet] {
	return pagination.NewFleetsCursorPageAutoPager(r.ListFleets(ctx, query, opts...))
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

type AuthFleetsResponse struct {
	Cursor  string                 `json:"cursor" api:"required,nullable"`
	Fleets  []Fleet                `json:"fleets" api:"required"`
	HasMore bool                   `json:"hasMore" api:"required"`
	JSON    authFleetsResponseJSON `json:"-"`
}

// authFleetsResponseJSON contains the JSON metadata for the struct
// [AuthFleetsResponse]
type authFleetsResponseJSON struct {
	Cursor      apijson.Field
	Fleets      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthFleetsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authFleetsResponseJSON) RawJSON() string {
	return r.raw
}

type AuthListFleetsParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [AuthListFleetsParams]'s query parameters as `url.Values`.
func (r AuthListFleetsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
