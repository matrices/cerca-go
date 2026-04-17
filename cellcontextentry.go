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

// CellContextEntryService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellContextEntryService] method instead.
type CellContextEntryService struct {
	Options []option.RequestOption
}

// NewCellContextEntryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCellContextEntryService(opts ...option.RequestOption) (r *CellContextEntryService) {
	r = &CellContextEntryService{}
	r.Options = opts
	return
}

func (r *CellContextEntryService) Get(ctx context.Context, cellID string, query CellContextEntryGetParams, opts ...option.RequestOption) (res *ContextEntry, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/context/entry", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *CellContextEntryService) Delete(ctx context.Context, cellID string, body CellContextEntryDeleteParams, opts ...option.RequestOption) (res *CellContextEntryDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/context/entry", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

func (r *CellContextEntryService) Put(ctx context.Context, cellID string, body CellContextEntryPutParams, opts ...option.RequestOption) (res *ContextEntry, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/context/entry", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type CellContextEntryDeleteResponse struct {
	Success CellContextEntryDeleteResponseSuccess `json:"success" api:"required"`
	JSON    cellContextEntryDeleteResponseJSON    `json:"-"`
}

// cellContextEntryDeleteResponseJSON contains the JSON metadata for the struct
// [CellContextEntryDeleteResponse]
type cellContextEntryDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellContextEntryDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellContextEntryDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CellContextEntryDeleteResponseSuccess bool

const (
	CellContextEntryDeleteResponseSuccessTrue CellContextEntryDeleteResponseSuccess = true
)

func (r CellContextEntryDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case CellContextEntryDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type CellContextEntryGetParams struct {
	// Context entry key.
	Key param.Field[string] `query:"key" api:"required"`
}

// URLQuery serializes [CellContextEntryGetParams]'s query parameters as
// `url.Values`.
func (r CellContextEntryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CellContextEntryDeleteParams struct {
	// Context entry key.
	Key param.Field[string] `query:"key" api:"required"`
}

// URLQuery serializes [CellContextEntryDeleteParams]'s query parameters as
// `url.Values`.
func (r CellContextEntryDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CellContextEntryPutParams struct {
	Content         param.Field[string]  `json:"content" api:"required"`
	Key             param.Field[string]  `json:"key" api:"required"`
	ExpectedVersion param.Field[float64] `json:"expectedVersion"`
	MimeType        param.Field[string]  `json:"mimeType"`
}

func (r CellContextEntryPutParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
