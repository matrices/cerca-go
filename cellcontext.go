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

// CellContextService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellContextService] method instead.
type CellContextService struct {
	Options []option.RequestOption
	Entries *CellContextEntryService
}

// NewCellContextService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCellContextService(opts ...option.RequestOption) (r *CellContextService) {
	r = &CellContextService{}
	r.Options = opts
	r.Entries = NewCellContextEntryService(opts...)
	return
}

func (r *CellContextService) List(ctx context.Context, cellID string, query CellContextListParams, opts ...option.RequestOption) (res *CellContextListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/context", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *CellContextService) Search(ctx context.Context, cellID string, query CellContextSearchParams, opts ...option.RequestOption) (res *CellContextSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/context/search", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ContextEntry struct {
	Content   string           `json:"content" api:"required"`
	Key       string           `json:"key" api:"required"`
	UpdatedAt string           `json:"updatedAt" api:"required"`
	Version   float64          `json:"version" api:"required"`
	MimeType  string           `json:"mimeType"`
	JSON      contextEntryJSON `json:"-"`
}

// contextEntryJSON contains the JSON metadata for the struct [ContextEntry]
type contextEntryJSON struct {
	Content     apijson.Field
	Key         apijson.Field
	UpdatedAt   apijson.Field
	Version     apijson.Field
	MimeType    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContextEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contextEntryJSON) RawJSON() string {
	return r.raw
}

type ContextEntrySummary struct {
	ByteLength float64                 `json:"byteLength" api:"required"`
	Key        string                  `json:"key" api:"required"`
	UpdatedAt  string                  `json:"updatedAt" api:"required"`
	Version    float64                 `json:"version" api:"required"`
	MimeType   string                  `json:"mimeType"`
	JSON       contextEntrySummaryJSON `json:"-"`
}

// contextEntrySummaryJSON contains the JSON metadata for the struct
// [ContextEntrySummary]
type contextEntrySummaryJSON struct {
	ByteLength  apijson.Field
	Key         apijson.Field
	UpdatedAt   apijson.Field
	Version     apijson.Field
	MimeType    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContextEntrySummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contextEntrySummaryJSON) RawJSON() string {
	return r.raw
}

type ContextSearchResult struct {
	ByteLength float64                 `json:"byteLength" api:"required"`
	Key        string                  `json:"key" api:"required"`
	Score      float64                 `json:"score" api:"required"`
	Snippet    string                  `json:"snippet" api:"required"`
	UpdatedAt  string                  `json:"updatedAt" api:"required"`
	Version    float64                 `json:"version" api:"required"`
	MimeType   string                  `json:"mimeType"`
	JSON       contextSearchResultJSON `json:"-"`
}

// contextSearchResultJSON contains the JSON metadata for the struct
// [ContextSearchResult]
type contextSearchResultJSON struct {
	ByteLength  apijson.Field
	Key         apijson.Field
	Score       apijson.Field
	Snippet     apijson.Field
	UpdatedAt   apijson.Field
	Version     apijson.Field
	MimeType    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContextSearchResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contextSearchResultJSON) RawJSON() string {
	return r.raw
}

type CellContextListResponse struct {
	Cursor  string                      `json:"cursor" api:"required,nullable"`
	Entries []ContextEntrySummary       `json:"entries" api:"required"`
	HasMore bool                        `json:"hasMore" api:"required"`
	JSON    cellContextListResponseJSON `json:"-"`
}

// cellContextListResponseJSON contains the JSON metadata for the struct
// [CellContextListResponse]
type cellContextListResponseJSON struct {
	Cursor      apijson.Field
	Entries     apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellContextListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellContextListResponseJSON) RawJSON() string {
	return r.raw
}

type CellContextSearchResponse struct {
	Cursor  string                        `json:"cursor" api:"required,nullable"`
	HasMore bool                          `json:"hasMore" api:"required"`
	Results []ContextSearchResult         `json:"results" api:"required"`
	JSON    cellContextSearchResponseJSON `json:"-"`
}

// cellContextSearchResponseJSON contains the JSON metadata for the struct
// [CellContextSearchResponse]
type cellContextSearchResponseJSON struct {
	Cursor      apijson.Field
	HasMore     apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellContextSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellContextSearchResponseJSON) RawJSON() string {
	return r.raw
}

type CellContextListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
	// Optional key prefix filter.
	Prefix param.Field[string] `query:"prefix"`
}

// URLQuery serializes [CellContextListParams]'s query parameters as `url.Values`.
func (r CellContextListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CellContextSearchParams struct {
	// Search query.
	Q param.Field[string] `query:"q" api:"required"`
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
	// Optional key prefix filter.
	Prefix param.Field[string] `query:"prefix"`
}

// URLQuery serializes [CellContextSearchParams]'s query parameters as
// `url.Values`.
func (r CellContextSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
