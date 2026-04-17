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

// CellLogService contains methods and other services that help with interacting
// with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellLogService] method instead.
type CellLogService struct {
	Options []option.RequestOption
}

// NewCellLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCellLogService(opts ...option.RequestOption) (r *CellLogService) {
	r = &CellLogService{}
	r.Options = opts
	return
}

func (r *CellLogService) List(ctx context.Context, cellID string, query CellLogListParams, opts ...option.RequestOption) (res *CellLogListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/logs", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type CellLogListResponse struct {
	Cursor  string                  `json:"cursor" api:"required,nullable"`
	HasMore bool                    `json:"hasMore" api:"required"`
	Logs    []ThreadLog             `json:"logs" api:"required"`
	JSON    cellLogListResponseJSON `json:"-"`
}

// cellLogListResponseJSON contains the JSON metadata for the struct
// [CellLogListResponse]
type cellLogListResponseJSON struct {
	Cursor      apijson.Field
	HasMore     apijson.Field
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellLogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellLogListResponseJSON) RawJSON() string {
	return r.raw
}

type CellLogListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [CellLogListParams]'s query parameters as `url.Values`.
func (r CellLogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
