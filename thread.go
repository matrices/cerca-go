// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cercago

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/matrices/cerca-go/internal/apijson"
	"github.com/matrices/cerca-go/internal/apiquery"
	"github.com/matrices/cerca-go/internal/param"
	"github.com/matrices/cerca-go/internal/requestconfig"
	"github.com/matrices/cerca-go/option"
	"github.com/matrices/cerca-go/packages/pagination"
	"github.com/matrices/cerca-go/shared"
	"github.com/tidwall/gjson"
)

// ThreadService contains methods and other services that help with interacting
// with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreadService] method instead.
type ThreadService struct {
	Options []option.RequestOption
}

// NewThreadService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewThreadService(opts ...option.RequestOption) (r *ThreadService) {
	r = &ThreadService{}
	r.Options = opts
	return
}

func (r *ThreadService) New(ctx context.Context, agentID string, body ThreadNewParams, opts ...option.RequestOption) (res *Thread, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/threads", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ThreadService) Get(ctx context.Context, agentID string, threadID string, query ThreadGetParams, opts ...option.RequestOption) (res *Thread, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/threads/%s", agentID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *ThreadService) List(ctx context.Context, agentID string, query ThreadListParams, opts ...option.RequestOption) (res *pagination.ThreadsCursorPage[ThreadSummary], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/threads", agentID)
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

func (r *ThreadService) ListAutoPaging(ctx context.Context, agentID string, query ThreadListParams, opts ...option.RequestOption) *pagination.ThreadsCursorPageAutoPager[ThreadSummary] {
	return pagination.NewThreadsCursorPageAutoPager(r.List(ctx, agentID, query, opts...))
}

// Cancel a running or awaiting thread. The underlying runtime treats repeat
// cancellation as an idempotent lifecycle operation when possible.
func (r *ThreadService) Cancel(ctx context.Context, agentID string, threadID string, opts ...option.RequestOption) (res *Thread, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/threads/%s/cancel", agentID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Close an idle thread. Closing a running, awaiting, or already-closed thread
// returns a lifecycle conflict.
func (r *ThreadService) Close(ctx context.Context, agentID string, threadID string, opts ...option.RequestOption) (res *Thread, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/threads/%s/close", agentID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Force context compaction for an idle thread. Compacting a running thread returns
// a lifecycle conflict.
func (r *ThreadService) Compact(ctx context.Context, agentID string, threadID string, opts ...option.RequestOption) (res *Thread, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/threads/%s/compact", agentID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

func (r *ThreadService) StartTurn(ctx context.Context, agentID string, threadID string, body ThreadStartTurnParams, opts ...option.RequestOption) (res *Turn, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/threads/%s/turns", agentID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Steer a thread with another user message. Steering a closed thread returns a
// conflict; steering a running or awaiting thread queues the message.
func (r *ThreadService) Steer(ctx context.Context, agentID string, threadID string, body ThreadSteerParams, opts ...option.RequestOption) (res *SteerResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/threads/%s/steer", agentID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type CompiledContext struct {
	EnabledTools []shared.ToolName   `json:"enabledTools" api:"required,nullable"`
	SystemPrompt string              `json:"systemPrompt" api:"required,nullable"`
	TurnID       string              `json:"turnId" api:"required,nullable"`
	JSON         compiledContextJSON `json:"-"`
}

// compiledContextJSON contains the JSON metadata for the struct [CompiledContext]
type compiledContextJSON struct {
	EnabledTools apijson.Field
	SystemPrompt apijson.Field
	TurnID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CompiledContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r compiledContextJSON) RawJSON() string {
	return r.raw
}

type ContentBlock map[string]interface{}

type ContextWindow struct {
	Compacted            bool                          `json:"compacted" api:"required"`
	CompactionThroughSeq float64                       `json:"compactionThroughSeq" api:"required,nullable"`
	ContextWindowTokens  float64                       `json:"contextWindowTokens" api:"required"`
	EstimationMethod     ContextWindowEstimationMethod `json:"estimationMethod" api:"required"`
	LlmEstimatedTokens   float64                       `json:"llmEstimatedTokens" api:"required"`
	Model                string                        `json:"model" api:"required"`
	RawEstimatedTokens   float64                       `json:"rawEstimatedTokens" api:"required"`
	JSON                 contextWindowJSON             `json:"-"`
}

// contextWindowJSON contains the JSON metadata for the struct [ContextWindow]
type contextWindowJSON struct {
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

func (r *ContextWindow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contextWindowJSON) RawJSON() string {
	return r.raw
}

type ContextWindowEstimationMethod string

const (
	ContextWindowEstimationMethodHeuristicRequestV1 ContextWindowEstimationMethod = "heuristic_request_v1"
)

func (r ContextWindowEstimationMethod) IsKnown() bool {
	switch r {
	case ContextWindowEstimationMethodHeuristicRequestV1:
		return true
	}
	return false
}

type Message struct {
	Content   []ContentBlock `json:"content" api:"required"`
	CreatedAt string         `json:"createdAt" api:"required"`
	Role      MessageRole    `json:"role" api:"required"`
	Seq       float64        `json:"seq" api:"required"`
	JSON      messageJSON    `json:"-"`
}

// messageJSON contains the JSON metadata for the struct [Message]
type messageJSON struct {
	Content     apijson.Field
	CreatedAt   apijson.Field
	Role        apijson.Field
	Seq         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageJSON) RawJSON() string {
	return r.raw
}

type MessageRole string

const (
	MessageRoleUser      MessageRole = "user"
	MessageRoleAssistant MessageRole = "assistant"
)

func (r MessageRole) IsKnown() bool {
	switch r {
	case MessageRoleUser, MessageRoleAssistant:
		return true
	}
	return false
}

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

// `idle` threads can accept a new turn or be closed. `running` threads have an
// active turn. `awaiting` threads are paused on external input such as approvals.
// `closed` threads are terminal.
type Status string

const (
	StatusIdle     Status = "idle"
	StatusRunning  Status = "running"
	StatusAwaiting Status = "awaiting"
	StatusClosed   Status = "closed"
)

func (r Status) IsKnown() bool {
	switch r {
	case StatusIdle, StatusRunning, StatusAwaiting, StatusClosed:
		return true
	}
	return false
}

type SteerResult struct {
	Status SteerResultStatus `json:"status" api:"required"`
	// `idle` threads can accept a new turn or be closed. `running` threads have an
	// active turn. `awaiting` threads are paused on external input such as approvals.
	// `closed` threads are terminal.
	ThreadStatus Status          `json:"threadStatus" api:"required"`
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
	ID                   string               `json:"id" api:"required"`
	AgentID              string               `json:"agentId" api:"required"`
	CompactionSummary    string               `json:"compactionSummary" api:"required,nullable"`
	CompactionThroughSeq float64              `json:"compactionThroughSeq" api:"required,nullable"`
	CompiledContext      CompiledContext      `json:"compiledContext" api:"required,nullable"`
	CompletedAt          string               `json:"completedAt" api:"required,nullable"`
	ComposedSystemPrompt string               `json:"composedSystemPrompt" api:"required,nullable"`
	ContextWindow        ContextWindow        `json:"contextWindow" api:"required"`
	CreatedAt            string               `json:"createdAt" api:"required"`
	Depth                float64              `json:"depth" api:"required"`
	Error                string               `json:"error" api:"required,nullable"`
	Instructions         string               `json:"instructions" api:"required,nullable"`
	LastTurnStatus       ThreadLastTurnStatus `json:"lastTurnStatus" api:"required,nullable"`
	Messages             []Message            `json:"messages" api:"required"`
	Model                string               `json:"model" api:"required"`
	ParentThreadID       string               `json:"parentThreadId" api:"required,nullable"`
	PendingSubThreads    []PendingSubThread   `json:"pendingSubThreads" api:"required"`
	Result               string               `json:"result" api:"required,nullable"`
	ScheduleID           string               `json:"scheduleId" api:"required,nullable"`
	ScheduleSeq          float64              `json:"scheduleSeq" api:"required,nullable"`
	// `idle` threads can accept a new turn or be closed. `running` threads have an
	// active turn. `awaiting` threads are paused on external input such as approvals.
	// `closed` threads are terminal.
	Status                 Status                            `json:"status" api:"required"`
	SubThreads             []ThreadSummary                   `json:"subThreads" api:"required"`
	ToolSnapshot           []shared.ToolName                 `json:"toolSnapshot" api:"required"`
	Turns                  []Turn                            `json:"turns" api:"required"`
	UpdatedAt              string                            `json:"updatedAt" api:"required"`
	UserMessage            string                            `json:"userMessage" api:"required"`
	ExternalToolNamespaces ThreadExternalToolNamespacesUnion `json:"externalToolNamespaces"`
	JSON                   threadJSON                        `json:"-"`
}

// threadJSON contains the JSON metadata for the struct [Thread]
type threadJSON struct {
	ID                     apijson.Field
	AgentID                apijson.Field
	CompactionSummary      apijson.Field
	CompactionThroughSeq   apijson.Field
	CompiledContext        apijson.Field
	CompletedAt            apijson.Field
	ComposedSystemPrompt   apijson.Field
	ContextWindow          apijson.Field
	CreatedAt              apijson.Field
	Depth                  apijson.Field
	Error                  apijson.Field
	Instructions           apijson.Field
	LastTurnStatus         apijson.Field
	Messages               apijson.Field
	Model                  apijson.Field
	ParentThreadID         apijson.Field
	PendingSubThreads      apijson.Field
	Result                 apijson.Field
	ScheduleID             apijson.Field
	ScheduleSeq            apijson.Field
	Status                 apijson.Field
	SubThreads             apijson.Field
	ToolSnapshot           apijson.Field
	Turns                  apijson.Field
	UpdatedAt              apijson.Field
	UserMessage            apijson.Field
	ExternalToolNamespaces apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
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

// Union satisfied by [ThreadExternalToolNamespacesString] or
// [ThreadExternalToolNamespacesArray].
type ThreadExternalToolNamespacesUnion interface {
	implementsThreadExternalToolNamespacesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ThreadExternalToolNamespacesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ThreadExternalToolNamespacesString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ThreadExternalToolNamespacesArray{}),
		},
	)
}

type ThreadExternalToolNamespacesString string

const (
	ThreadExternalToolNamespacesStringAll ThreadExternalToolNamespacesString = "all"
)

func (r ThreadExternalToolNamespacesString) IsKnown() bool {
	switch r {
	case ThreadExternalToolNamespacesStringAll:
		return true
	}
	return false
}

func (r ThreadExternalToolNamespacesString) implementsThreadExternalToolNamespacesUnion() {}

type ThreadExternalToolNamespacesArray []string

func (r ThreadExternalToolNamespacesArray) implementsThreadExternalToolNamespacesUnion() {}

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
	Status    Status            `json:"status" api:"required"`
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

type ThreadNewParams struct {
	Instructions param.Field[string] `json:"instructions"`
	Model        param.Field[string] `json:"model"`
	// Deprecated alias for `instructions`; accepted for backwards compatibility.
	SystemPrompt param.Field[string]                 `json:"systemPrompt"`
	Tools        param.Field[[]shared.ToolSpecParam] `json:"tools"`
	UserMessage  param.Field[string]                 `json:"userMessage"`
}

func (r ThreadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreadGetParams struct {
	// When true, includes debug-only compiled context fields.
	Debug param.Field[ThreadGetParamsDebug] `query:"debug"`
	// When true, includes message content in the thread detail.
	IncludeMessages param.Field[ThreadGetParamsIncludeMessages] `query:"includeMessages"`
}

// URLQuery serializes [ThreadGetParams]'s query parameters as `url.Values`.
func (r ThreadGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// When true, includes debug-only compiled context fields.
type ThreadGetParamsDebug string

const (
	ThreadGetParamsDebugTrue  ThreadGetParamsDebug = "true"
	ThreadGetParamsDebugFalse ThreadGetParamsDebug = "false"
)

func (r ThreadGetParamsDebug) IsKnown() bool {
	switch r {
	case ThreadGetParamsDebugTrue, ThreadGetParamsDebugFalse:
		return true
	}
	return false
}

// When true, includes message content in the thread detail.
type ThreadGetParamsIncludeMessages string

const (
	ThreadGetParamsIncludeMessagesTrue  ThreadGetParamsIncludeMessages = "true"
	ThreadGetParamsIncludeMessagesFalse ThreadGetParamsIncludeMessages = "false"
)

func (r ThreadGetParamsIncludeMessages) IsKnown() bool {
	switch r {
	case ThreadGetParamsIncludeMessagesTrue, ThreadGetParamsIncludeMessagesFalse:
		return true
	}
	return false
}

type ThreadListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
	// Optional schedule id filter.
	ScheduleID param.Field[string] `query:"scheduleId"`
	// Optional thread status filter.
	Status param.Field[Status] `query:"status"`
}

// URLQuery serializes [ThreadListParams]'s query parameters as `url.Values`.
func (r ThreadListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ThreadStartTurnParams struct {
	UserMessage param.Field[string]                 `json:"userMessage" api:"required"`
	Model       param.Field[string]                 `json:"model"`
	Tools       param.Field[[]shared.ToolSpecParam] `json:"tools"`
}

func (r ThreadStartTurnParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreadSteerParams struct {
	Message param.Field[string] `json:"message" api:"required"`
}

func (r ThreadSteerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
