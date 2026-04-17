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

// CellThreadLogService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellThreadLogService] method instead.
type CellThreadLogService struct {
	Options []option.RequestOption
}

// NewCellThreadLogService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCellThreadLogService(opts ...option.RequestOption) (r *CellThreadLogService) {
	r = &CellThreadLogService{}
	r.Options = opts
	return
}

func (r *CellThreadLogService) List(ctx context.Context, cellID string, threadID string, query CellThreadLogListParams, opts ...option.RequestOption) (res *CellThreadLogListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/threads/%s/logs", cellID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ThreadLog struct {
	ID              string          `json:"id" api:"required"`
	Activity        string          `json:"activity" api:"required,nullable"`
	CreatedAt       string          `json:"createdAt" api:"required"`
	InputTokens     float64         `json:"inputTokens" api:"required"`
	MessageCount    float64         `json:"messageCount" api:"required"`
	NextStep        string          `json:"nextStep" api:"required,nullable"`
	OutputTokens    float64         `json:"outputTokens" api:"required"`
	Status          ThreadLogStatus `json:"status" api:"required"`
	ThreadID        string          `json:"threadId" api:"required"`
	TurnCompletedAt string          `json:"turnCompletedAt" api:"required"`
	TurnSeq         float64         `json:"turnSeq" api:"required"`
	JSON            threadLogJSON   `json:"-"`
}

// threadLogJSON contains the JSON metadata for the struct [ThreadLog]
type threadLogJSON struct {
	ID              apijson.Field
	Activity        apijson.Field
	CreatedAt       apijson.Field
	InputTokens     apijson.Field
	MessageCount    apijson.Field
	NextStep        apijson.Field
	OutputTokens    apijson.Field
	Status          apijson.Field
	ThreadID        apijson.Field
	TurnCompletedAt apijson.Field
	TurnSeq         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ThreadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadLogJSON) RawJSON() string {
	return r.raw
}

type ThreadLogStatus string

const (
	ThreadLogStatusCompleted ThreadLogStatus = "completed"
	ThreadLogStatusFailed    ThreadLogStatus = "failed"
)

func (r ThreadLogStatus) IsKnown() bool {
	switch r {
	case ThreadLogStatusCompleted, ThreadLogStatusFailed:
		return true
	}
	return false
}

type CellThreadLogListResponse struct {
	Cursor  string                        `json:"cursor" api:"required,nullable"`
	HasMore bool                          `json:"hasMore" api:"required"`
	Logs    []ThreadLog                   `json:"logs" api:"required"`
	JSON    cellThreadLogListResponseJSON `json:"-"`
}

// cellThreadLogListResponseJSON contains the JSON metadata for the struct
// [CellThreadLogListResponse]
type cellThreadLogListResponseJSON struct {
	Cursor      apijson.Field
	HasMore     apijson.Field
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellThreadLogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellThreadLogListResponseJSON) RawJSON() string {
	return r.raw
}

type CellThreadLogListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [CellThreadLogListParams]'s query parameters as
// `url.Values`.
func (r CellThreadLogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
