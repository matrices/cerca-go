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

// CellToolSourceService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellToolSourceService] method instead.
type CellToolSourceService struct {
	Options []option.RequestOption
}

// NewCellToolSourceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCellToolSourceService(opts ...option.RequestOption) (r *CellToolSourceService) {
	r = &CellToolSourceService{}
	r.Options = opts
	return
}

func (r *CellToolSourceService) List(ctx context.Context, cellID string, opts ...option.RequestOption) (res *CellToolSourceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/tool-sources", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CellToolSourceListResponse struct {
	Cursor  string                         `json:"cursor" api:"required,nullable"`
	HasMore bool                           `json:"hasMore" api:"required"`
	Sources []ToolSource                   `json:"sources" api:"required"`
	JSON    cellToolSourceListResponseJSON `json:"-"`
}

// cellToolSourceListResponseJSON contains the JSON metadata for the struct
// [CellToolSourceListResponse]
type cellToolSourceListResponseJSON struct {
	Cursor      apijson.Field
	HasMore     apijson.Field
	Sources     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellToolSourceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellToolSourceListResponseJSON) RawJSON() string {
	return r.raw
}
