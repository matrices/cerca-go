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

// LogService contains methods and other services that help with interacting with
// the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogService] method instead.
type LogService struct {
	Options []option.RequestOption
}

// NewLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLogService(opts ...option.RequestOption) (r *LogService) {
	r = &LogService{}
	r.Options = opts
	return
}

func (r *LogService) ListForAgent(ctx context.Context, agentID string, query LogListForAgentParams, opts ...option.RequestOption) (res *pagination.LogsCursorPage[ThreadLog], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/logs", agentID)
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

func (r *LogService) ListForAgentAutoPaging(ctx context.Context, agentID string, query LogListForAgentParams, opts ...option.RequestOption) *pagination.LogsCursorPageAutoPager[ThreadLog] {
	return pagination.NewLogsCursorPageAutoPager(r.ListForAgent(ctx, agentID, query, opts...))
}

func (r *LogService) ListForThread(ctx context.Context, agentID string, threadID string, query LogListForThreadParams, opts ...option.RequestOption) (res *pagination.LogsCursorPage[ThreadLog], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/threads/%s/logs", agentID, threadID)
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

func (r *LogService) ListForThreadAutoPaging(ctx context.Context, agentID string, threadID string, query LogListForThreadParams, opts ...option.RequestOption) *pagination.LogsCursorPageAutoPager[ThreadLog] {
	return pagination.NewLogsCursorPageAutoPager(r.ListForThread(ctx, agentID, threadID, query, opts...))
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

type LogListForAgentParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [LogListForAgentParams]'s query parameters as `url.Values`.
func (r LogListForAgentParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LogListForThreadParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [LogListForThreadParams]'s query parameters as `url.Values`.
func (r LogListForThreadParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
