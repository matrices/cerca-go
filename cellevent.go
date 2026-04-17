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

// CellEventService contains methods and other services that help with interacting
// with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellEventService] method instead.
type CellEventService struct {
	Options []option.RequestOption
}

// NewCellEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCellEventService(opts ...option.RequestOption) (r *CellEventService) {
	r = &CellEventService{}
	r.Options = opts
	return
}

func (r *CellEventService) List(ctx context.Context, cellID string, query CellEventListParams, opts ...option.RequestOption) (res *CellEventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/events", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// WebSocket upgrade endpoint. Set
// `Sec-WebSocket-Protocol: cell-v1, cell-auth-<API_KEY>` so the runtime can
// authenticate the stream while preserving the public subprotocol. HTTP clients
// that cannot upgrade should use `/cells/{cellId}/events` as the polling analog.
func (r *CellEventService) Subscribe(ctx context.Context, cellID string, opts ...option.RequestOption) (res *CellEventSubscribeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/events/subscribe", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CellEventListResponse struct {
	Cursor  string                    `json:"cursor" api:"required,nullable"`
	Events  []SubscriptionEvent       `json:"events" api:"required"`
	HasMore bool                      `json:"hasMore" api:"required"`
	JSON    cellEventListResponseJSON `json:"-"`
}

// cellEventListResponseJSON contains the JSON metadata for the struct
// [CellEventListResponse]
type cellEventListResponseJSON struct {
	Cursor      apijson.Field
	Events      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellEventListResponseJSON) RawJSON() string {
	return r.raw
}

type CellEventSubscribeResponse struct {
	Error string                         `json:"error" api:"required"`
	JSON  cellEventSubscribeResponseJSON `json:"-"`
}

// cellEventSubscribeResponseJSON contains the JSON metadata for the struct
// [CellEventSubscribeResponse]
type cellEventSubscribeResponseJSON struct {
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellEventSubscribeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellEventSubscribeResponseJSON) RawJSON() string {
	return r.raw
}

type CellEventListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Comma-separated event type filter.
	Events param.Field[string] `query:"events"`
	// When true, starts from the beginning of the retained buffer.
	History param.Field[CellEventListParamsHistory] `query:"history"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [CellEventListParams]'s query parameters as `url.Values`.
func (r CellEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// When true, starts from the beginning of the retained buffer.
type CellEventListParamsHistory string

const (
	CellEventListParamsHistoryTrue  CellEventListParamsHistory = "true"
	CellEventListParamsHistoryFalse CellEventListParamsHistory = "false"
)

func (r CellEventListParamsHistory) IsKnown() bool {
	switch r {
	case CellEventListParamsHistoryTrue, CellEventListParamsHistoryFalse:
		return true
	}
	return false
}
