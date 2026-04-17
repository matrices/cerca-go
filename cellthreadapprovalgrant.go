// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cercago

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/matrices/cerca-go/internal/apijson"
	"github.com/matrices/cerca-go/internal/requestconfig"
	"github.com/matrices/cerca-go/option"
)

// CellThreadApprovalGrantService contains methods and other services that help
// with interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellThreadApprovalGrantService] method instead.
type CellThreadApprovalGrantService struct {
	Options []option.RequestOption
}

// NewCellThreadApprovalGrantService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCellThreadApprovalGrantService(opts ...option.RequestOption) (r *CellThreadApprovalGrantService) {
	r = &CellThreadApprovalGrantService{}
	r.Options = opts
	return
}

func (r *CellThreadApprovalGrantService) List(ctx context.Context, cellID string, threadID string, opts ...option.RequestOption) (res *[]ApprovalGrant, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/threads/%s/approval-grants", cellID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *CellThreadApprovalGrantService) Delete(ctx context.Context, cellID string, threadID string, grantID string, opts ...option.RequestOption) (res *CellThreadApprovalGrantDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
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
	path := fmt.Sprintf("cells/%s/threads/%s/approval-grants/%s", cellID, threadID, grantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type CellThreadApprovalGrantDeleteResponse struct {
	Ok   CellThreadApprovalGrantDeleteResponseOk   `json:"ok" api:"required"`
	JSON cellThreadApprovalGrantDeleteResponseJSON `json:"-"`
}

// cellThreadApprovalGrantDeleteResponseJSON contains the JSON metadata for the
// struct [CellThreadApprovalGrantDeleteResponse]
type cellThreadApprovalGrantDeleteResponseJSON struct {
	Ok          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellThreadApprovalGrantDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellThreadApprovalGrantDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CellThreadApprovalGrantDeleteResponseOk bool

const (
	CellThreadApprovalGrantDeleteResponseOkTrue CellThreadApprovalGrantDeleteResponseOk = true
)

func (r CellThreadApprovalGrantDeleteResponseOk) IsKnown() bool {
	switch r {
	case CellThreadApprovalGrantDeleteResponseOkTrue:
		return true
	}
	return false
}
