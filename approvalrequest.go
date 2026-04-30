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

// ApprovalRequestService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApprovalRequestService] method instead.
type ApprovalRequestService struct {
	Options []option.RequestOption
}

// NewApprovalRequestService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewApprovalRequestService(opts ...option.RequestOption) (r *ApprovalRequestService) {
	r = &ApprovalRequestService{}
	r.Options = opts
	return
}

// Approvals
func (r *ApprovalRequestService) List(ctx context.Context, agentID string, query ApprovalRequestListParams, opts ...option.RequestOption) (res *pagination.ApprovalRequestsCursorPage[ApprovalRequest], err error) {
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

// Approvals
func (r *ApprovalRequestService) ListAutoPaging(ctx context.Context, agentID string, query ApprovalRequestListParams, opts ...option.RequestOption) *pagination.ApprovalRequestsCursorPageAutoPager[ApprovalRequest] {
	return pagination.NewApprovalRequestsCursorPageAutoPager(r.List(ctx, agentID, query, opts...))
}

// Approval
func (r *ApprovalRequestService) Resolve(ctx context.Context, agentID string, threadID string, approvalID string, body ApprovalRequestResolveParams, opts ...option.RequestOption) (res *ApprovalRequest, err error) {
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
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Parsed JSON tool input from the original tool call. Generated SDKs may expose
	// this as unknown or Any.
	Input             interface{}           `json:"input" api:"required"`
	ResolvedAt        string                `json:"resolvedAt" api:"required,nullable"`
	RuntimeToolName   shared.ToolName       `json:"runtimeToolName" api:"required"`
	Status            ApprovalRequestStatus `json:"status" api:"required"`
	ThreadID          string                `json:"threadId" api:"required"`
	TimeoutAt         string                `json:"timeoutAt" api:"required,nullable"`
	TimeoutMs         float64               `json:"timeoutMs" api:"required,nullable"`
	ToolIndex         float64               `json:"toolIndex" api:"required"`
	ToolName          string                `json:"toolName" api:"required"`
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
	Input             apijson.Field
	ResolvedAt        apijson.Field
	RuntimeToolName   apijson.Field
	Status            apijson.Field
	ThreadID          apijson.Field
	TimeoutAt         apijson.Field
	TimeoutMs         apijson.Field
	ToolIndex         apijson.Field
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

type ApprovalRequestListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
	// Optional thread id filter.
	ThreadID param.Field[string] `query:"threadId"`
}

// URLQuery serializes [ApprovalRequestListParams]'s query parameters as
// `url.Values`.
func (r ApprovalRequestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ApprovalRequestResolveParams struct {
	Decision param.Field[ApprovalRequestResolveParamsDecision] `json:"decision" api:"required"`
	Grant    param.Field[ApprovalRequestResolveParamsGrant]    `json:"grant"`
}

func (r ApprovalRequestResolveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ApprovalRequestResolveParamsDecision string

const (
	ApprovalRequestResolveParamsDecisionApprove ApprovalRequestResolveParamsDecision = "approve"
	ApprovalRequestResolveParamsDecisionDeny    ApprovalRequestResolveParamsDecision = "deny"
	ApprovalRequestResolveParamsDecisionCancel  ApprovalRequestResolveParamsDecision = "cancel"
)

func (r ApprovalRequestResolveParamsDecision) IsKnown() bool {
	switch r {
	case ApprovalRequestResolveParamsDecisionApprove, ApprovalRequestResolveParamsDecisionDeny, ApprovalRequestResolveParamsDecisionCancel:
		return true
	}
	return false
}

type ApprovalRequestResolveParamsGrant string

const (
	ApprovalRequestResolveParamsGrantThread ApprovalRequestResolveParamsGrant = "thread"
	ApprovalRequestResolveParamsGrantAgent  ApprovalRequestResolveParamsGrant = "agent"
)

func (r ApprovalRequestResolveParamsGrant) IsKnown() bool {
	switch r {
	case ApprovalRequestResolveParamsGrantThread, ApprovalRequestResolveParamsGrantAgent:
		return true
	}
	return false
}
