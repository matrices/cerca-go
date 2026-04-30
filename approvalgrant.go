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
	"github.com/matrices/cerca-go/packages/pagination"
)

// ApprovalGrantService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApprovalGrantService] method instead.
type ApprovalGrantService struct {
	Options []option.RequestOption
}

// NewApprovalGrantService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewApprovalGrantService(opts ...option.RequestOption) (r *ApprovalGrantService) {
	r = &ApprovalGrantService{}
	r.Options = opts
	return
}

// List approval grants
func (r *ApprovalGrantService) List(ctx context.Context, agentID string, query ApprovalGrantListParams, opts ...option.RequestOption) (res *pagination.GrantsCursorPage[ApprovalGrant], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/approval-grants", agentID)
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

// List approval grants
func (r *ApprovalGrantService) ListAutoPaging(ctx context.Context, agentID string, query ApprovalGrantListParams, opts ...option.RequestOption) *pagination.GrantsCursorPageAutoPager[ApprovalGrant] {
	return pagination.NewGrantsCursorPageAutoPager(r.List(ctx, agentID, query, opts...))
}

// Delete approval grant
func (r *ApprovalGrantService) Delete(ctx context.Context, agentID string, grantID string, opts ...option.RequestOption) (res *ApprovalGrantDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if grantID == "" {
		err = errors.New("missing required grantId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/approval-grants/%s", agentID, grantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Delete approval grant
func (r *ApprovalGrantService) DeleteForThread(ctx context.Context, agentID string, threadID string, grantID string, opts ...option.RequestOption) (res *ApprovalGrantDeleteForThreadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	if grantID == "" {
		err = errors.New("missing required grantId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/threads/%s/approval-grants/%s", agentID, threadID, grantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// List approval grants
func (r *ApprovalGrantService) ListForThread(ctx context.Context, agentID string, threadID string, opts ...option.RequestOption) (res *[]ApprovalGrant, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/threads/%s/approval-grants", agentID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
	ApprovalGrantScopeAgent  ApprovalGrantScope = "agent"
)

func (r ApprovalGrantScope) IsKnown() bool {
	switch r {
	case ApprovalGrantScopeThread, ApprovalGrantScopeAgent:
		return true
	}
	return false
}

type ApprovalGrantDeleteResponse struct {
	Ok   ApprovalGrantDeleteResponseOk   `json:"ok" api:"required"`
	JSON approvalGrantDeleteResponseJSON `json:"-"`
}

// approvalGrantDeleteResponseJSON contains the JSON metadata for the struct
// [ApprovalGrantDeleteResponse]
type approvalGrantDeleteResponseJSON struct {
	Ok          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApprovalGrantDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r approvalGrantDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ApprovalGrantDeleteResponseOk bool

const (
	ApprovalGrantDeleteResponseOkTrue ApprovalGrantDeleteResponseOk = true
)

func (r ApprovalGrantDeleteResponseOk) IsKnown() bool {
	switch r {
	case ApprovalGrantDeleteResponseOkTrue:
		return true
	}
	return false
}

type ApprovalGrantDeleteForThreadResponse struct {
	Ok   ApprovalGrantDeleteForThreadResponseOk   `json:"ok" api:"required"`
	JSON approvalGrantDeleteForThreadResponseJSON `json:"-"`
}

// approvalGrantDeleteForThreadResponseJSON contains the JSON metadata for the
// struct [ApprovalGrantDeleteForThreadResponse]
type approvalGrantDeleteForThreadResponseJSON struct {
	Ok          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApprovalGrantDeleteForThreadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r approvalGrantDeleteForThreadResponseJSON) RawJSON() string {
	return r.raw
}

type ApprovalGrantDeleteForThreadResponseOk bool

const (
	ApprovalGrantDeleteForThreadResponseOkTrue ApprovalGrantDeleteForThreadResponseOk = true
)

func (r ApprovalGrantDeleteForThreadResponseOk) IsKnown() bool {
	switch r {
	case ApprovalGrantDeleteForThreadResponseOkTrue:
		return true
	}
	return false
}

type ApprovalGrantListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [ApprovalGrantListParams]'s query parameters as
// `url.Values`.
func (r ApprovalGrantListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
