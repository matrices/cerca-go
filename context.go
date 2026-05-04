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
)

// ContextService contains methods and other services that help with interacting
// with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContextService] method instead.
type ContextService struct {
	Options []option.RequestOption
}

// NewContextService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewContextService(opts ...option.RequestOption) (r *ContextService) {
	r = &ContextService{}
	r.Options = opts
	return
}

// Retrieve context entry
func (r *ContextService) Get(ctx context.Context, agentID string, query ContextGetParams, opts ...option.RequestOption) (res *Entry, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/context/entry", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List context entries
func (r *ContextService) List(ctx context.Context, agentID string, query ContextListParams, opts ...option.RequestOption) (res *pagination.EntriesCursorPage[EntrySummary], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/context", agentID)
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

// List context entries
func (r *ContextService) ListAutoPaging(ctx context.Context, agentID string, query ContextListParams, opts ...option.RequestOption) *pagination.EntriesCursorPageAutoPager[EntrySummary] {
	return pagination.NewEntriesCursorPageAutoPager(r.List(ctx, agentID, query, opts...))
}

// Delete context entry
func (r *ContextService) Delete(ctx context.Context, agentID string, body ContextDeleteParams, opts ...option.RequestOption) (res *ContextDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/context/entry", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Search context
func (r *ContextService) Search(ctx context.Context, agentID string, query ContextSearchParams, opts ...option.RequestOption) (res *pagination.ResultsCursorPage[SearchResult], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/context/search", agentID)
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

// Search context
func (r *ContextService) SearchAutoPaging(ctx context.Context, agentID string, query ContextSearchParams, opts ...option.RequestOption) *pagination.ResultsCursorPageAutoPager[SearchResult] {
	return pagination.NewResultsCursorPageAutoPager(r.Search(ctx, agentID, query, opts...))
}

// Update context entry
func (r *ContextService) Write(ctx context.Context, agentID string, body ContextWriteParams, opts ...option.RequestOption) (res *Entry, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/context/entry", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type Entry struct {
	Content   string    `json:"content" api:"required"`
	Key       string    `json:"key" api:"required"`
	UpdatedAt string    `json:"updatedAt" api:"required"`
	Version   float64   `json:"version" api:"required"`
	MimeType  string    `json:"mimeType"`
	JSON      entryJSON `json:"-"`
}

// entryJSON contains the JSON metadata for the struct [Entry]
type entryJSON struct {
	Content     apijson.Field
	Key         apijson.Field
	UpdatedAt   apijson.Field
	Version     apijson.Field
	MimeType    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Entry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entryJSON) RawJSON() string {
	return r.raw
}

type EntrySummary struct {
	ByteLength float64          `json:"byteLength" api:"required"`
	Key        string           `json:"key" api:"required"`
	UpdatedAt  string           `json:"updatedAt" api:"required"`
	Version    float64          `json:"version" api:"required"`
	MimeType   string           `json:"mimeType"`
	JSON       entrySummaryJSON `json:"-"`
}

// entrySummaryJSON contains the JSON metadata for the struct [EntrySummary]
type entrySummaryJSON struct {
	ByteLength  apijson.Field
	Key         apijson.Field
	UpdatedAt   apijson.Field
	Version     apijson.Field
	MimeType    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntrySummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entrySummaryJSON) RawJSON() string {
	return r.raw
}

type SearchResult struct {
	ByteLength float64          `json:"byteLength" api:"required"`
	Key        string           `json:"key" api:"required"`
	Score      float64          `json:"score" api:"required"`
	Snippet    string           `json:"snippet" api:"required"`
	UpdatedAt  string           `json:"updatedAt" api:"required"`
	Version    float64          `json:"version" api:"required"`
	MimeType   string           `json:"mimeType"`
	JSON       searchResultJSON `json:"-"`
}

// searchResultJSON contains the JSON metadata for the struct [SearchResult]
type searchResultJSON struct {
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

func (r *SearchResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchResultJSON) RawJSON() string {
	return r.raw
}

type ContextDeleteResponse struct {
	Success ContextDeleteResponseSuccess `json:"success" api:"required"`
	JSON    contextDeleteResponseJSON    `json:"-"`
}

// contextDeleteResponseJSON contains the JSON metadata for the struct
// [ContextDeleteResponse]
type contextDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContextDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contextDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ContextDeleteResponseSuccess bool

const (
	ContextDeleteResponseSuccessTrue ContextDeleteResponseSuccess = true
)

func (r ContextDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ContextDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ContextGetParams struct {
	// Context entry key.
	Key param.Field[string] `query:"key" api:"required"`
}

// URLQuery serializes [ContextGetParams]'s query parameters as `url.Values`.
func (r ContextGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ContextListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
	// Optional key prefix filter.
	Prefix param.Field[string] `query:"prefix"`
}

// URLQuery serializes [ContextListParams]'s query parameters as `url.Values`.
func (r ContextListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ContextDeleteParams struct {
	// Context entry key.
	Key param.Field[string] `query:"key" api:"required"`
}

// URLQuery serializes [ContextDeleteParams]'s query parameters as `url.Values`.
func (r ContextDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ContextSearchParams struct {
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

// URLQuery serializes [ContextSearchParams]'s query parameters as `url.Values`.
func (r ContextSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ContextWriteParams struct {
	Content         param.Field[string]  `json:"content" api:"required"`
	Key             param.Field[string]  `json:"key" api:"required"`
	ExpectedVersion param.Field[float64] `json:"expectedVersion"`
	MimeType        param.Field[string]  `json:"mimeType"`
}

func (r ContextWriteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
