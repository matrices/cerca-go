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

// CellApprovalService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellApprovalService] method instead.
type CellApprovalService struct {
	Options []option.RequestOption
}

// NewCellApprovalService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCellApprovalService(opts ...option.RequestOption) (r *CellApprovalService) {
	r = &CellApprovalService{}
	r.Options = opts
	return
}

func (r *CellApprovalService) List(ctx context.Context, cellID string, query CellApprovalListParams, opts ...option.RequestOption) (res *CellApprovalListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/approvals", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ApprovalRequest struct {
	ID                string                `json:"id" api:"required"`
	CreatedAt         string                `json:"createdAt" api:"required"`
	RequestedToolName string                `json:"requestedToolName" api:"required"`
	ResolvedAt        string                `json:"resolvedAt" api:"required,nullable"`
	Status            ApprovalRequestStatus `json:"status" api:"required"`
	ThreadID          string                `json:"threadId" api:"required"`
	TimeoutAt         string                `json:"timeoutAt" api:"required,nullable"`
	TimeoutMs         float64               `json:"timeoutMs" api:"required,nullable"`
	ToolIndex         float64               `json:"toolIndex" api:"required"`
	ToolInput         string                `json:"toolInput" api:"required"`
	ToolName          RuntimeToolName       `json:"toolName" api:"required"`
	ToolUseID         string                `json:"toolUseId" api:"required"`
	TurnID            string                `json:"turnId" api:"required"`
	ToolSourceID      string                `json:"toolSourceId"`
	ToolSourceVersion float64               `json:"toolSourceVersion"`
	JSON              approvalRequestJSON   `json:"-"`
}

// approvalRequestJSON contains the JSON metadata for the struct [ApprovalRequest]
type approvalRequestJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	RequestedToolName apijson.Field
	ResolvedAt        apijson.Field
	Status            apijson.Field
	ThreadID          apijson.Field
	TimeoutAt         apijson.Field
	TimeoutMs         apijson.Field
	ToolIndex         apijson.Field
	ToolInput         apijson.Field
	ToolName          apijson.Field
	ToolUseID         apijson.Field
	TurnID            apijson.Field
	ToolSourceID      apijson.Field
	ToolSourceVersion apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ApprovalRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r approvalRequestJSON) RawJSON() string {
	return r.raw
}

type ApprovalRequestStatus string

const (
	ApprovalRequestStatusPending   ApprovalRequestStatus = "pending"
	ApprovalRequestStatusApproved  ApprovalRequestStatus = "approved"
	ApprovalRequestStatusDenied    ApprovalRequestStatus = "denied"
	ApprovalRequestStatusCancelled ApprovalRequestStatus = "cancelled"
	ApprovalRequestStatusTimedOut  ApprovalRequestStatus = "timed_out"
)

func (r ApprovalRequestStatus) IsKnown() bool {
	switch r {
	case ApprovalRequestStatusPending, ApprovalRequestStatusApproved, ApprovalRequestStatusDenied, ApprovalRequestStatusCancelled, ApprovalRequestStatusTimedOut:
		return true
	}
	return false
}

type CellApprovalListResponse struct {
	Approvals []ApprovalRequest            `json:"approvals" api:"required"`
	Cursor    string                       `json:"cursor" api:"required,nullable"`
	HasMore   bool                         `json:"hasMore" api:"required"`
	JSON      cellApprovalListResponseJSON `json:"-"`
}

// cellApprovalListResponseJSON contains the JSON metadata for the struct
// [CellApprovalListResponse]
type cellApprovalListResponseJSON struct {
	Approvals   apijson.Field
	Cursor      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellApprovalListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellApprovalListResponseJSON) RawJSON() string {
	return r.raw
}

type CellApprovalListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
	// Optional thread id filter.
	ThreadID param.Field[string] `query:"threadId"`
}

// URLQuery serializes [CellApprovalListParams]'s query parameters as `url.Values`.
func (r CellApprovalListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
