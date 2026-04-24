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
	"github.com/matrices/cerca-go/shared"
)

// ApprovalService contains methods and other services that help with interacting
// with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApprovalService] method instead.
type ApprovalService struct {
	Options []option.RequestOption
}

// NewApprovalService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewApprovalService(opts ...option.RequestOption) (r *ApprovalService) {
	r = &ApprovalService{}
	r.Options = opts
	return
}

func (r *ApprovalService) List(ctx context.Context, agentID string, query ApprovalListParams, opts ...option.RequestOption) (res *pagination.ApprovalsCursorPage[ApprovalRequest], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/approvals", agentID)
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

func (r *ApprovalService) ListAutoPaging(ctx context.Context, agentID string, query ApprovalListParams, opts ...option.RequestOption) *pagination.ApprovalsCursorPageAutoPager[ApprovalRequest] {
	return pagination.NewApprovalsCursorPageAutoPager(r.List(ctx, agentID, query, opts...))
}

func (r *ApprovalService) Resolve(ctx context.Context, agentID string, threadID string, approvalID string, body ApprovalResolveParams, opts ...option.RequestOption) (res *ApprovalRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	if approvalID == "" {
		err = errors.New("missing required approvalId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/threads/%s/approvals/%s", agentID, threadID, approvalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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
	ToolName          shared.ToolName       `json:"toolName" api:"required"`
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

type ApprovalListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
	// Optional thread id filter.
	ThreadID param.Field[string] `query:"threadId"`
}

// URLQuery serializes [ApprovalListParams]'s query parameters as `url.Values`.
func (r ApprovalListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ApprovalResolveParams struct {
	Decision param.Field[ApprovalResolveParamsDecision] `json:"decision" api:"required"`
	Grant    param.Field[ApprovalResolveParamsGrant]    `json:"grant"`
}

func (r ApprovalResolveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ApprovalResolveParamsDecision string

const (
	ApprovalResolveParamsDecisionApprove ApprovalResolveParamsDecision = "approve"
	ApprovalResolveParamsDecisionDeny    ApprovalResolveParamsDecision = "deny"
	ApprovalResolveParamsDecisionCancel  ApprovalResolveParamsDecision = "cancel"
)

func (r ApprovalResolveParamsDecision) IsKnown() bool {
	switch r {
	case ApprovalResolveParamsDecisionApprove, ApprovalResolveParamsDecisionDeny, ApprovalResolveParamsDecisionCancel:
		return true
	}
	return false
}

type ApprovalResolveParamsGrant string

const (
	ApprovalResolveParamsGrantThread ApprovalResolveParamsGrant = "thread"
	ApprovalResolveParamsGrantAgent  ApprovalResolveParamsGrant = "agent"
)

func (r ApprovalResolveParamsGrant) IsKnown() bool {
	switch r {
	case ApprovalResolveParamsGrantThread, ApprovalResolveParamsGrantAgent:
		return true
	}
	return false
}
