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

// CellConnectionService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellConnectionService] method instead.
type CellConnectionService struct {
	Options []option.RequestOption
}

// NewCellConnectionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCellConnectionService(opts ...option.RequestOption) (r *CellConnectionService) {
	r = &CellConnectionService{}
	r.Options = opts
	return
}

func (r *CellConnectionService) List(ctx context.Context, cellID string, opts ...option.RequestOption) (res *CellConnectionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/connections", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *CellConnectionService) Attach(ctx context.Context, cellID string, body CellConnectionAttachParams, opts ...option.RequestOption) (res *AttachedConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/connections", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *CellConnectionService) Detach(ctx context.Context, cellID string, connectionID string, opts ...option.RequestOption) (res *CellConnectionDetachResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	if connectionID == "" {
		err = errors.New("missing required connectionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/connections/%s", cellID, connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type AttachedConnection struct {
	AttachedAt string                 `json:"attachedAt" api:"required"`
	Metadata   ToolConnectionMetadata `json:"metadata"`
	JSON       attachedConnectionJSON `json:"-"`
	Connection
}

// attachedConnectionJSON contains the JSON metadata for the struct
// [AttachedConnection]
type attachedConnectionJSON struct {
	AttachedAt  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttachedConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attachedConnectionJSON) RawJSON() string {
	return r.raw
}

type ToolConnectionMetadata map[string]string

type ToolConnectionMetadataParam map[string]string

type CellConnectionListResponse struct {
	Connections []AttachedConnection           `json:"connections" api:"required"`
	Cursor      string                         `json:"cursor" api:"required,nullable"`
	HasMore     bool                           `json:"hasMore" api:"required"`
	JSON        cellConnectionListResponseJSON `json:"-"`
}

// cellConnectionListResponseJSON contains the JSON metadata for the struct
// [CellConnectionListResponse]
type cellConnectionListResponseJSON struct {
	Connections apijson.Field
	Cursor      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellConnectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellConnectionListResponseJSON) RawJSON() string {
	return r.raw
}

type CellConnectionDetachResponse struct {
	Success CellConnectionDetachResponseSuccess `json:"success" api:"required"`
	JSON    cellConnectionDetachResponseJSON    `json:"-"`
}

// cellConnectionDetachResponseJSON contains the JSON metadata for the struct
// [CellConnectionDetachResponse]
type cellConnectionDetachResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellConnectionDetachResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellConnectionDetachResponseJSON) RawJSON() string {
	return r.raw
}

type CellConnectionDetachResponseSuccess bool

const (
	CellConnectionDetachResponseSuccessTrue CellConnectionDetachResponseSuccess = true
)

func (r CellConnectionDetachResponseSuccess) IsKnown() bool {
	switch r {
	case CellConnectionDetachResponseSuccessTrue:
		return true
	}
	return false
}

type CellConnectionAttachParams struct {
	ConnectionID param.Field[string]                      `json:"connectionId" api:"required"`
	Metadata     param.Field[ToolConnectionMetadataParam] `json:"metadata"`
}

func (r CellConnectionAttachParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
