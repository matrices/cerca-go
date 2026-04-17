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

// CellThreadEventService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellThreadEventService] method instead.
type CellThreadEventService struct {
	Options []option.RequestOption
}

// NewCellThreadEventService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCellThreadEventService(opts ...option.RequestOption) (r *CellThreadEventService) {
	r = &CellThreadEventService{}
	r.Options = opts
	return
}

func (r *CellThreadEventService) List(ctx context.Context, cellID string, threadID string, query CellThreadEventListParams, opts ...option.RequestOption) (res *CellThreadEventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/threads/%s/events", cellID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// WebSocket upgrade endpoint. Set
// `Sec-WebSocket-Protocol: cell-v1, cell-auth-<API_KEY>` so the runtime can
// authenticate the stream while preserving the public subprotocol. HTTP clients
// that cannot upgrade should use `/cells/{cellId}/threads/{threadId}/events` as
// the polling analog.
func (r *CellThreadEventService) Subscribe(ctx context.Context, cellID string, threadID string, opts ...option.RequestOption) (res *CellThreadEventSubscribeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/threads/%s/events/subscribe", cellID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CellThreadEventListResponse struct {
	Cursor  string                          `json:"cursor" api:"required,nullable"`
	Events  []SubscriptionEvent             `json:"events" api:"required"`
	HasMore bool                            `json:"hasMore" api:"required"`
	JSON    cellThreadEventListResponseJSON `json:"-"`
}

// cellThreadEventListResponseJSON contains the JSON metadata for the struct
// [CellThreadEventListResponse]
type cellThreadEventListResponseJSON struct {
	Cursor      apijson.Field
	Events      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellThreadEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellThreadEventListResponseJSON) RawJSON() string {
	return r.raw
}

type CellThreadEventSubscribeResponse struct {
	Error string                               `json:"error" api:"required"`
	JSON  cellThreadEventSubscribeResponseJSON `json:"-"`
}

// cellThreadEventSubscribeResponseJSON contains the JSON metadata for the struct
// [CellThreadEventSubscribeResponse]
type cellThreadEventSubscribeResponseJSON struct {
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellThreadEventSubscribeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellThreadEventSubscribeResponseJSON) RawJSON() string {
	return r.raw
}

type CellThreadEventListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Comma-separated event type filter.
	Events param.Field[string] `query:"events"`
	// When true, starts from the beginning of the retained buffer.
	History param.Field[CellThreadEventListParamsHistory] `query:"history"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [CellThreadEventListParams]'s query parameters as
// `url.Values`.
func (r CellThreadEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// When true, starts from the beginning of the retained buffer.
type CellThreadEventListParamsHistory string

const (
	CellThreadEventListParamsHistoryTrue  CellThreadEventListParamsHistory = "true"
	CellThreadEventListParamsHistoryFalse CellThreadEventListParamsHistory = "false"
)

func (r CellThreadEventListParamsHistory) IsKnown() bool {
	switch r {
	case CellThreadEventListParamsHistoryTrue, CellThreadEventListParamsHistoryFalse:
		return true
	}
	return false
}
