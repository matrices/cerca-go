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

// CellApprovalGrantService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellApprovalGrantService] method instead.
type CellApprovalGrantService struct {
	Options []option.RequestOption
}

// NewCellApprovalGrantService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCellApprovalGrantService(opts ...option.RequestOption) (r *CellApprovalGrantService) {
	r = &CellApprovalGrantService{}
	r.Options = opts
	return
}

func (r *CellApprovalGrantService) List(ctx context.Context, cellID string, query CellApprovalGrantListParams, opts ...option.RequestOption) (res *CellApprovalGrantListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/approval-grants", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *CellApprovalGrantService) Delete(ctx context.Context, cellID string, grantID string, opts ...option.RequestOption) (res *CellApprovalGrantDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	if grantID == "" {
		err = errors.New("missing required grantId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/approval-grants/%s", cellID, grantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type ApprovalGrant struct {
	ID        string             `json:"id" api:"required"`
	CreatedAt string             `json:"createdAt" api:"required"`
	CreatedBy string             `json:"createdBy" api:"required,nullable"`
	ExpiresAt string             `json:"expiresAt" api:"required,nullable"`
	GrantKey  string             `json:"grantKey" api:"required"`
	Scope     ApprovalGrantScope `json:"scope" api:"required"`
	JSON      approvalGrantJSON  `json:"-"`
}

// approvalGrantJSON contains the JSON metadata for the struct [ApprovalGrant]
type approvalGrantJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	CreatedBy   apijson.Field
	ExpiresAt   apijson.Field
	GrantKey    apijson.Field
	Scope       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApprovalGrant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r approvalGrantJSON) RawJSON() string {
	return r.raw
}

type ApprovalGrantScope string

const (
	ApprovalGrantScopeThread ApprovalGrantScope = "thread"
	ApprovalGrantScopeCell   ApprovalGrantScope = "cell"
)

func (r ApprovalGrantScope) IsKnown() bool {
	switch r {
	case ApprovalGrantScopeThread, ApprovalGrantScopeCell:
		return true
	}
	return false
}

type CellApprovalGrantListResponse struct {
	Cursor  string                            `json:"cursor" api:"required,nullable"`
	Grants  []ApprovalGrant                   `json:"grants" api:"required"`
	HasMore bool                              `json:"hasMore" api:"required"`
	JSON    cellApprovalGrantListResponseJSON `json:"-"`
}

// cellApprovalGrantListResponseJSON contains the JSON metadata for the struct
// [CellApprovalGrantListResponse]
type cellApprovalGrantListResponseJSON struct {
	Cursor      apijson.Field
	Grants      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellApprovalGrantListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellApprovalGrantListResponseJSON) RawJSON() string {
	return r.raw
}

type CellApprovalGrantDeleteResponse struct {
	Ok   CellApprovalGrantDeleteResponseOk   `json:"ok" api:"required"`
	JSON cellApprovalGrantDeleteResponseJSON `json:"-"`
}

// cellApprovalGrantDeleteResponseJSON contains the JSON metadata for the struct
// [CellApprovalGrantDeleteResponse]
type cellApprovalGrantDeleteResponseJSON struct {
	Ok          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellApprovalGrantDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellApprovalGrantDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CellApprovalGrantDeleteResponseOk bool

const (
	CellApprovalGrantDeleteResponseOkTrue CellApprovalGrantDeleteResponseOk = true
)

func (r CellApprovalGrantDeleteResponseOk) IsKnown() bool {
	switch r {
	case CellApprovalGrantDeleteResponseOkTrue:
		return true
	}
	return false
}

type CellApprovalGrantListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [CellApprovalGrantListParams]'s query parameters as
// `url.Values`.
func (r CellApprovalGrantListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
