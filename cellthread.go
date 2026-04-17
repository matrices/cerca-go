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

// CellThreadService contains methods and other services that help with interacting
// with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellThreadService] method instead.
type CellThreadService struct {
	Options        []option.RequestOption
	Approvals      *CellThreadApprovalService
	ApprovalGrants *CellThreadApprovalGrantService
	Logs           *CellThreadLogService
	Events         *CellThreadEventService
}

// NewCellThreadService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCellThreadService(opts ...option.RequestOption) (r *CellThreadService) {
	r = &CellThreadService{}
	r.Options = opts
	r.Approvals = NewCellThreadApprovalService(opts...)
	r.ApprovalGrants = NewCellThreadApprovalGrantService(opts...)
	r.Logs = NewCellThreadLogService(opts...)
	r.Events = NewCellThreadEventService(opts...)
	return
}

func (r *CellThreadService) New(ctx context.Context, cellID string, body CellThreadNewParams, opts ...option.RequestOption) (res *Thread, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/threads", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *CellThreadService) Get(ctx context.Context, cellID string, threadID string, query CellThreadGetParams, opts ...option.RequestOption) (res *Thread, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/threads/%s", cellID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *CellThreadService) List(ctx context.Context, cellID string, query CellThreadListParams, opts ...option.RequestOption) (res *CellThreadListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/threads", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Cancel a running or awaiting thread. The underlying runtime treats repeat
// cancellation as an idempotent lifecycle operation when possible.
func (r *CellThreadService) Cancel(ctx context.Context, cellID string, threadID string, opts ...option.RequestOption) (res *Thread, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/threads/%s/cancel", cellID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Close an idle thread. Closing a running, awaiting, or already-closed thread
// returns a lifecycle conflict.
func (r *CellThreadService) Close(ctx context.Context, cellID string, threadID string, opts ...option.RequestOption) (res *Thread, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/threads/%s/close", cellID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Force context compaction for an idle thread. Compacting a running thread returns
// a lifecycle conflict.
func (r *CellThreadService) Compact(ctx context.Context, cellID string, threadID string, opts ...option.RequestOption) (res *Thread, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/threads/%s/compact", cellID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Steer a thread with another user message. Steering a closed thread returns a
// conflict; steering a running or awaiting thread queues the message.
func (r *CellThreadService) Steer(ctx context.Context, cellID string, threadID string, body CellThreadSteerParams, opts ...option.RequestOption) (res *SteerResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/threads/%s/steer", cellID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *CellThreadService) Turn(ctx context.Context, cellID string, threadID string, body CellThreadTurnParams, opts ...option.RequestOption) (res *Turn, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/threads/%s/turns", cellID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ContentBlock = interface{}

type PendingSubThread struct {
	SubThreadID string               `json:"subThreadId" api:"required"`
	ToolUseID   string               `json:"toolUseId" api:"required"`
	JSON        pendingSubThreadJSON `json:"-"`
}

// pendingSubThreadJSON contains the JSON metadata for the struct
// [PendingSubThread]
type pendingSubThreadJSON struct {
	SubThreadID apijson.Field
	ToolUseID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PendingSubThread) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingSubThreadJSON) RawJSON() string {
	return r.raw
}

type SteerResult struct {
	Status SteerResultStatus `json:"status" api:"required"`
	// `idle` threads can accept a new turn or be closed. `running` threads have an
	// active turn. `awaiting` threads are paused on external input such as approvals.
	// `closed` threads are terminal.
	ThreadStatus ThreadStatus    `json:"threadStatus" api:"required"`
	TurnID       string          `json:"turnId" api:"required"`
	JSON         steerResultJSON `json:"-"`
}

// steerResultJSON contains the JSON metadata for the struct [SteerResult]
type steerResultJSON struct {
	Status       apijson.Field
	ThreadStatus apijson.Field
	TurnID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SteerResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r steerResultJSON) RawJSON() string {
	return r.raw
}

type SteerResultStatus string

const (
	SteerResultStatusEnqueued    SteerResultStatus = "enqueued"
	SteerResultStatusTurnStarted SteerResultStatus = "turn_started"
)

func (r SteerResultStatus) IsKnown() bool {
	switch r {
	case SteerResultStatusEnqueued, SteerResultStatusTurnStarted:
		return true
	}
	return false
}

type Thread struct {
	ID                   string                `json:"id" api:"required"`
	CellID               string                `json:"cellId" api:"required"`
	CompactionSummary    string                `json:"compactionSummary" api:"required,nullable"`
	CompactionThroughSeq float64               `json:"compactionThroughSeq" api:"required,nullable"`
	CompiledContext      ThreadCompiledContext `json:"compiledContext" api:"required,nullable"`
	CompletedAt          string                `json:"completedAt" api:"required,nullable"`
	ComposedSystemPrompt string                `json:"composedSystemPrompt" api:"required,nullable"`
	ContextWindow        ThreadContextWindow   `json:"contextWindow" api:"required"`
	CreatedAt            string                `json:"createdAt" api:"required"`
	Depth                float64               `json:"depth" api:"required"`
	Error                string                `json:"error" api:"required,nullable"`
	Instructions         string                `json:"instructions" api:"required,nullable"`
	LastTurnStatus       ThreadLastTurnStatus  `json:"lastTurnStatus" api:"required,nullable"`
	Messages             []ThreadMessage       `json:"messages" api:"required"`
	Model                string                `json:"model" api:"required"`
	ParentThreadID       string                `json:"parentThreadId" api:"required,nullable"`
	PendingSubThreads    []PendingSubThread    `json:"pendingSubThreads" api:"required"`
	Result               string                `json:"result" api:"required,nullable"`
	ScheduleID           string                `json:"scheduleId" api:"required,nullable"`
	ScheduleSeq          float64               `json:"scheduleSeq" api:"required,nullable"`
	// `idle` threads can accept a new turn or be closed. `running` threads have an
	// active turn. `awaiting` threads are paused on external input such as approvals.
	// `closed` threads are terminal.
	Status      ThreadStatus      `json:"status" api:"required"`
	SubThreads  []ThreadSummary   `json:"subThreads" api:"required"`
	Tools       []RuntimeToolName `json:"tools" api:"required"`
	Turns       []Turn            `json:"turns" api:"required"`
	UpdatedAt   string            `json:"updatedAt" api:"required"`
	UserMessage string            `json:"userMessage" api:"required"`
	JSON        threadJSON        `json:"-"`
}

// threadJSON contains the JSON metadata for the struct [Thread]
type threadJSON struct {
	ID                   apijson.Field
	CellID               apijson.Field
	CompactionSummary    apijson.Field
	CompactionThroughSeq apijson.Field
	CompiledContext      apijson.Field
	CompletedAt          apijson.Field
	ComposedSystemPrompt apijson.Field
	ContextWindow        apijson.Field
	CreatedAt            apijson.Field
	Depth                apijson.Field
	Error                apijson.Field
	Instructions         apijson.Field
	LastTurnStatus       apijson.Field
	Messages             apijson.Field
	Model                apijson.Field
	ParentThreadID       apijson.Field
	PendingSubThreads    apijson.Field
	Result               apijson.Field
	ScheduleID           apijson.Field
	ScheduleSeq          apijson.Field
	Status               apijson.Field
	SubThreads           apijson.Field
	Tools                apijson.Field
	Turns                apijson.Field
	UpdatedAt            apijson.Field
	UserMessage          apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Thread) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadJSON) RawJSON() string {
	return r.raw
}

type ThreadLastTurnStatus string

const (
	ThreadLastTurnStatusCompleted ThreadLastTurnStatus = "completed"
	ThreadLastTurnStatusFailed    ThreadLastTurnStatus = "failed"
)

func (r ThreadLastTurnStatus) IsKnown() bool {
	switch r {
	case ThreadLastTurnStatusCompleted, ThreadLastTurnStatusFailed:
		return true
	}
	return false
}

type ThreadCompiledContext struct {
	EnabledTools []RuntimeToolName         `json:"enabledTools" api:"required,nullable"`
	SystemPrompt string                    `json:"systemPrompt" api:"required,nullable"`
	TurnID       string                    `json:"turnId" api:"required,nullable"`
	JSON         threadCompiledContextJSON `json:"-"`
}

// threadCompiledContextJSON contains the JSON metadata for the struct
// [ThreadCompiledContext]
type threadCompiledContextJSON struct {
	EnabledTools apijson.Field
	SystemPrompt apijson.Field
	TurnID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreadCompiledContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadCompiledContextJSON) RawJSON() string {
	return r.raw
}

type ThreadContextWindow struct {
	Compacted            bool                                `json:"compacted" api:"required"`
	CompactionThroughSeq float64                             `json:"compactionThroughSeq" api:"required,nullable"`
	ContextWindowTokens  float64                             `json:"contextWindowTokens" api:"required"`
	EstimationMethod     ThreadContextWindowEstimationMethod `json:"estimationMethod" api:"required"`
	LlmEstimatedTokens   float64                             `json:"llmEstimatedTokens" api:"required"`
	Model                string                              `json:"model" api:"required"`
	RawEstimatedTokens   float64                             `json:"rawEstimatedTokens" api:"required"`
	JSON                 threadContextWindowJSON             `json:"-"`
}

// threadContextWindowJSON contains the JSON metadata for the struct
// [ThreadContextWindow]
type threadContextWindowJSON struct {
	Compacted            apijson.Field
	CompactionThroughSeq apijson.Field
	ContextWindowTokens  apijson.Field
	EstimationMethod     apijson.Field
	LlmEstimatedTokens   apijson.Field
	Model                apijson.Field
	RawEstimatedTokens   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ThreadContextWindow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadContextWindowJSON) RawJSON() string {
	return r.raw
}

type ThreadContextWindowEstimationMethod string

const (
	ThreadContextWindowEstimationMethodHeuristicRequestV1 ThreadContextWindowEstimationMethod = "heuristic_request_v1"
)

func (r ThreadContextWindowEstimationMethod) IsKnown() bool {
	switch r {
	case ThreadContextWindowEstimationMethodHeuristicRequestV1:
		return true
	}
	return false
}

type ThreadMessage struct {
	Content   []ContentBlock    `json:"content" api:"required"`
	CreatedAt string            `json:"createdAt" api:"required"`
	Role      ThreadMessageRole `json:"role" api:"required"`
	Seq       float64           `json:"seq" api:"required"`
	JSON      threadMessageJSON `json:"-"`
}

// threadMessageJSON contains the JSON metadata for the struct [ThreadMessage]
type threadMessageJSON struct {
	Content     apijson.Field
	CreatedAt   apijson.Field
	Role        apijson.Field
	Seq         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreadMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadMessageJSON) RawJSON() string {
	return r.raw
}

type ThreadMessageRole string

const (
	ThreadMessageRoleUser      ThreadMessageRole = "user"
	ThreadMessageRoleAssistant ThreadMessageRole = "assistant"
)

func (r ThreadMessageRole) IsKnown() bool {
	switch r {
	case ThreadMessageRoleUser, ThreadMessageRoleAssistant:
		return true
	}
	return false
}

// `idle` threads can accept a new turn or be closed. `running` threads have an
// active turn. `awaiting` threads are paused on external input such as approvals.
// `closed` threads are terminal.
type ThreadStatus string

const (
	ThreadStatusIdle     ThreadStatus = "idle"
	ThreadStatusRunning  ThreadStatus = "running"
	ThreadStatusAwaiting ThreadStatus = "awaiting"
	ThreadStatusClosed   ThreadStatus = "closed"
)

func (r ThreadStatus) IsKnown() bool {
	switch r {
	case ThreadStatusIdle, ThreadStatusRunning, ThreadStatusAwaiting, ThreadStatusClosed:
		return true
	}
	return false
}

type ThreadSummary struct {
	ID             string  `json:"id" api:"required"`
	CompletedAt    string  `json:"completedAt" api:"required,nullable"`
	CreatedAt      string  `json:"createdAt" api:"required"`
	MessageCount   float64 `json:"messageCount" api:"required"`
	Model          string  `json:"model" api:"required"`
	ParentThreadID string  `json:"parentThreadId" api:"required,nullable"`
	Result         string  `json:"result" api:"required,nullable"`
	ScheduleID     string  `json:"scheduleId" api:"required,nullable"`
	ScheduleSeq    float64 `json:"scheduleSeq" api:"required,nullable"`
	// `idle` threads can accept a new turn or be closed. `running` threads have an
	// active turn. `awaiting` threads are paused on external input such as approvals.
	// `closed` threads are terminal.
	Status    ThreadStatus      `json:"status" api:"required"`
	StepCount float64           `json:"stepCount" api:"required"`
	JSON      threadSummaryJSON `json:"-"`
}

// threadSummaryJSON contains the JSON metadata for the struct [ThreadSummary]
type threadSummaryJSON struct {
	ID             apijson.Field
	CompletedAt    apijson.Field
	CreatedAt      apijson.Field
	MessageCount   apijson.Field
	Model          apijson.Field
	ParentThreadID apijson.Field
	Result         apijson.Field
	ScheduleID     apijson.Field
	ScheduleSeq    apijson.Field
	Status         apijson.Field
	StepCount      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ThreadSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadSummaryJSON) RawJSON() string {
	return r.raw
}

type TokenUsage struct {
	InputTokens  float64        `json:"inputTokens" api:"required"`
	OutputTokens float64        `json:"outputTokens" api:"required"`
	JSON         tokenUsageJSON `json:"-"`
}

// tokenUsageJSON contains the JSON metadata for the struct [TokenUsage]
type tokenUsageJSON struct {
	InputTokens  apijson.Field
	OutputTokens apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TokenUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenUsageJSON) RawJSON() string {
	return r.raw
}

type Turn struct {
	ID              string     `json:"id" api:"required"`
	CompletedAt     string     `json:"completedAt" api:"required,nullable"`
	EndMessageSeq   float64    `json:"endMessageSeq" api:"required,nullable"`
	Error           string     `json:"error" api:"required,nullable"`
	Result          string     `json:"result" api:"required,nullable"`
	Seq             float64    `json:"seq" api:"required"`
	StartedAt       string     `json:"startedAt" api:"required"`
	StartMessageSeq float64    `json:"startMessageSeq" api:"required"`
	Status          TurnStatus `json:"status" api:"required"`
	ThreadID        string     `json:"threadId" api:"required"`
	TokenUsage      TokenUsage `json:"tokenUsage" api:"required,nullable"`
	UserMessage     string     `json:"userMessage" api:"required"`
	JSON            turnJSON   `json:"-"`
}

// turnJSON contains the JSON metadata for the struct [Turn]
type turnJSON struct {
	ID              apijson.Field
	CompletedAt     apijson.Field
	EndMessageSeq   apijson.Field
	Error           apijson.Field
	Result          apijson.Field
	Seq             apijson.Field
	StartedAt       apijson.Field
	StartMessageSeq apijson.Field
	Status          apijson.Field
	ThreadID        apijson.Field
	TokenUsage      apijson.Field
	UserMessage     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Turn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnJSON) RawJSON() string {
	return r.raw
}

type TurnStatus string

const (
	TurnStatusRunning   TurnStatus = "running"
	TurnStatusCompleted TurnStatus = "completed"
	TurnStatusFailed    TurnStatus = "failed"
)

func (r TurnStatus) IsKnown() bool {
	switch r {
	case TurnStatusRunning, TurnStatusCompleted, TurnStatusFailed:
		return true
	}
	return false
}

type CellThreadListResponse struct {
	Cursor  string                     `json:"cursor" api:"required,nullable"`
	HasMore bool                       `json:"hasMore" api:"required"`
	Threads []ThreadSummary            `json:"threads" api:"required"`
	JSON    cellThreadListResponseJSON `json:"-"`
}

// cellThreadListResponseJSON contains the JSON metadata for the struct
// [CellThreadListResponse]
type cellThreadListResponseJSON struct {
	Cursor      apijson.Field
	HasMore     apijson.Field
	Threads     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellThreadListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellThreadListResponseJSON) RawJSON() string {
	return r.raw
}

type CellThreadNewParams struct {
	Features     param.Field[[]CellThreadNewParamsFeature] `json:"features"`
	Instructions param.Field[string]                       `json:"instructions"`
	Model        param.Field[string]                       `json:"model"`
	// Deprecated alias for `instructions`; accepted for backwards compatibility.
	SystemPrompt param.Field[string] `json:"systemPrompt"`
	UserMessage  param.Field[string] `json:"userMessage"`
}

func (r CellThreadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CellThreadNewParamsFeature string

const (
	CellThreadNewParamsFeatureMemory       CellThreadNewParamsFeature = "memory"
	CellThreadNewParamsFeatureSandbox      CellThreadNewParamsFeature = "sandbox"
	CellThreadNewParamsFeatureWeb          CellThreadNewParamsFeature = "web"
	CellThreadNewParamsFeatureConnections  CellThreadNewParamsFeature = "connections"
	CellThreadNewParamsFeatureOAuthConnect CellThreadNewParamsFeature = "oauth_connect"
)

func (r CellThreadNewParamsFeature) IsKnown() bool {
	switch r {
	case CellThreadNewParamsFeatureMemory, CellThreadNewParamsFeatureSandbox, CellThreadNewParamsFeatureWeb, CellThreadNewParamsFeatureConnections, CellThreadNewParamsFeatureOAuthConnect:
		return true
	}
	return false
}

type CellThreadGetParams struct {
	// When true, includes debug-only compiled context fields.
	Debug param.Field[CellThreadGetParamsDebug] `query:"debug"`
	// When true, includes message content in the thread detail.
	IncludeMessages param.Field[CellThreadGetParamsIncludeMessages] `query:"includeMessages"`
}

// URLQuery serializes [CellThreadGetParams]'s query parameters as `url.Values`.
func (r CellThreadGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// When true, includes debug-only compiled context fields.
type CellThreadGetParamsDebug string

const (
	CellThreadGetParamsDebugTrue  CellThreadGetParamsDebug = "true"
	CellThreadGetParamsDebugFalse CellThreadGetParamsDebug = "false"
)

func (r CellThreadGetParamsDebug) IsKnown() bool {
	switch r {
	case CellThreadGetParamsDebugTrue, CellThreadGetParamsDebugFalse:
		return true
	}
	return false
}

// When true, includes message content in the thread detail.
type CellThreadGetParamsIncludeMessages string

const (
	CellThreadGetParamsIncludeMessagesTrue  CellThreadGetParamsIncludeMessages = "true"
	CellThreadGetParamsIncludeMessagesFalse CellThreadGetParamsIncludeMessages = "false"
)

func (r CellThreadGetParamsIncludeMessages) IsKnown() bool {
	switch r {
	case CellThreadGetParamsIncludeMessagesTrue, CellThreadGetParamsIncludeMessagesFalse:
		return true
	}
	return false
}

type CellThreadListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
	// Optional schedule id filter.
	ScheduleID param.Field[string] `query:"scheduleId"`
	// Optional thread status filter.
	Status param.Field[ThreadStatus] `query:"status"`
}

// URLQuery serializes [CellThreadListParams]'s query parameters as `url.Values`.
func (r CellThreadListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CellThreadSteerParams struct {
	Message param.Field[string] `json:"message" api:"required"`
}

func (r CellThreadSteerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CellThreadTurnParams struct {
	UserMessage param.Field[string]                        `json:"userMessage" api:"required"`
	Features    param.Field[[]CellThreadTurnParamsFeature] `json:"features"`
	Model       param.Field[string]                        `json:"model"`
}

func (r CellThreadTurnParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CellThreadTurnParamsFeature string

const (
	CellThreadTurnParamsFeatureMemory       CellThreadTurnParamsFeature = "memory"
	CellThreadTurnParamsFeatureSandbox      CellThreadTurnParamsFeature = "sandbox"
	CellThreadTurnParamsFeatureWeb          CellThreadTurnParamsFeature = "web"
	CellThreadTurnParamsFeatureConnections  CellThreadTurnParamsFeature = "connections"
	CellThreadTurnParamsFeatureOAuthConnect CellThreadTurnParamsFeature = "oauth_connect"
)

func (r CellThreadTurnParamsFeature) IsKnown() bool {
	switch r {
	case CellThreadTurnParamsFeatureMemory, CellThreadTurnParamsFeatureSandbox, CellThreadTurnParamsFeatureWeb, CellThreadTurnParamsFeatureConnections, CellThreadTurnParamsFeatureOAuthConnect:
		return true
	}
	return false
}
