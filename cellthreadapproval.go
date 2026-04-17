// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cercago

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/matrices/cerca-go/internal/apijson"
	"github.com/matrices/cerca-go/internal/param"
	"github.com/matrices/cerca-go/internal/requestconfig"
	"github.com/matrices/cerca-go/option"
)

// CellThreadApprovalService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellThreadApprovalService] method instead.
type CellThreadApprovalService struct {
	Options []option.RequestOption
}

// NewCellThreadApprovalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCellThreadApprovalService(opts ...option.RequestOption) (r *CellThreadApprovalService) {
	r = &CellThreadApprovalService{}
	r.Options = opts
	return
}

func (r *CellThreadApprovalService) Resolve(ctx context.Context, cellID string, threadID string, approvalID string, body CellThreadApprovalResolveParams, opts ...option.RequestOption) (res *ApprovalRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
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
	path := fmt.Sprintf("cells/%s/threads/%s/approvals/%s", cellID, threadID, approvalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type CellThreadApprovalResolveParams struct {
	Decision param.Field[CellThreadApprovalResolveParamsDecision] `json:"decision" api:"required"`
	Grant    param.Field[CellThreadApprovalResolveParamsGrant]    `json:"grant"`
}

func (r CellThreadApprovalResolveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CellThreadApprovalResolveParamsDecision string

const (
	CellThreadApprovalResolveParamsDecisionApprove CellThreadApprovalResolveParamsDecision = "approve"
	CellThreadApprovalResolveParamsDecisionDeny    CellThreadApprovalResolveParamsDecision = "deny"
	CellThreadApprovalResolveParamsDecisionCancel  CellThreadApprovalResolveParamsDecision = "cancel"
)

func (r CellThreadApprovalResolveParamsDecision) IsKnown() bool {
	switch r {
	case CellThreadApprovalResolveParamsDecisionApprove, CellThreadApprovalResolveParamsDecisionDeny, CellThreadApprovalResolveParamsDecisionCancel:
		return true
	}
	return false
}

type CellThreadApprovalResolveParamsGrant string

const (
	CellThreadApprovalResolveParamsGrantThread CellThreadApprovalResolveParamsGrant = "thread"
	CellThreadApprovalResolveParamsGrantCell   CellThreadApprovalResolveParamsGrant = "cell"
)

func (r CellThreadApprovalResolveParamsGrant) IsKnown() bool {
	switch r {
	case CellThreadApprovalResolveParamsGrantThread, CellThreadApprovalResolveParamsGrantCell:
		return true
	}
	return false
}
