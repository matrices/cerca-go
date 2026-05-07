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

// Create thread
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

// Retrieve thread
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

// List threads
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

// List threads
func (r *ThreadService) ListAutoPaging(ctx context.Context, agentID string, query ThreadListParams, opts ...option.RequestOption) *pagination.ThreadsCursorPageAutoPager[ThreadSummary] {
	return pagination.NewThreadsCursorPageAutoPager(r.List(ctx, agentID, query, opts...))
}

// Fetch compact current and recent activity for a thread without returning
// transcript content or runtime debug state.
func (r *ThreadService) Activity(ctx context.Context, agentID string, threadID string, query ThreadActivityParams, opts ...option.RequestOption) (res *ActivityDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/threads/%s/activity", agentID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
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

// List a bounded page of transcript messages for a thread, newest first. Use the
// returned `cursor` to page older messages.
func (r *ThreadService) ListMessages(ctx context.Context, agentID string, threadID string, query ThreadListMessagesParams, opts ...option.RequestOption) (res *pagination.ThreadMessagesCursorPage[Message], err error) {
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
	path := fmt.Sprintf("agents/%s/threads/%s/messages", agentID, threadID)
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

// List a bounded page of transcript messages for a thread, newest first. Use the
// returned `cursor` to page older messages.
func (r *ThreadService) ListMessagesAutoPaging(ctx context.Context, agentID string, threadID string, query ThreadListMessagesParams, opts ...option.RequestOption) *pagination.ThreadMessagesCursorPageAutoPager[Message] {
	return pagination.NewThreadMessagesCursorPageAutoPager(r.ListMessages(ctx, agentID, threadID, query, opts...))
}

// Create turn
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

type ActivityDetail struct {
	Error       string              `json:"error" api:"required,nullable"`
	RecentTurns []ThreadTurnSummary `json:"recentTurns" api:"required"`
	JSON        activityDetailJSON  `json:"-"`
	ActivitySummary
}

// activityDetailJSON contains the JSON metadata for the struct [ActivityDetail]
type activityDetailJSON struct {
	Error       apijson.Field
	RecentTurns apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivityDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activityDetailJSON) RawJSON() string {
	return r.raw
}

type ActivitySummary struct {
	ID             string  `json:"id" api:"required"`
	CompletedAt    string  `json:"completedAt" api:"required,nullable"`
	CreatedAt      string  `json:"createdAt" api:"required"`
	Goal           string  `json:"goal" api:"required"`
	LatestActivity string  `json:"latestActivity" api:"required,nullable"`
	MessageCount   float64 `json:"messageCount" api:"required"`
	Model          string  `json:"model" api:"required"`
	NextStep       string  `json:"nextStep" api:"required,nullable"`
	ParentThreadID string  `json:"parentThreadId" api:"required,nullable"`
	Result         string  `json:"result" api:"required,nullable"`
	ScheduleID     string  `json:"scheduleId" api:"required,nullable"`
	// `idle` threads can accept a new turn or be closed. `running` threads have an
	// active turn. `awaiting` threads are paused on external input such as approvals.
	// `closed` threads are terminal.
	Status    Status              `json:"status" api:"required"`
	StepCount float64             `json:"stepCount" api:"required"`
	UpdatedAt string              `json:"updatedAt" api:"required"`
	JSON      activitySummaryJSON `json:"-"`
}

// activitySummaryJSON contains the JSON metadata for the struct [ActivitySummary]
type activitySummaryJSON struct {
	ID             apijson.Field
	CompletedAt    apijson.Field
	CreatedAt      apijson.Field
	Goal           apijson.Field
	LatestActivity apijson.Field
	MessageCount   apijson.Field
	Model          apijson.Field
	NextStep       apijson.Field
	ParentThreadID apijson.Field
	Result         apijson.Field
	ScheduleID     apijson.Field
	Status         apijson.Field
	StepCount      apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ActivitySummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activitySummaryJSON) RawJSON() string {
	return r.raw
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

// Message content block. The `type` field distinguishes text, tool use, tool
// result, server tool use, and web search result blocks.
type ContentBlock struct {
	Type ContentBlockType `json:"type" api:"required"`
	ID   string           `json:"id"`
	// This field can have the runtime type of [string], [interface{}].
	Content      interface{} `json:"content"`
	DeniedByUser bool        `json:"deniedByUser"`
	// This field can have the runtime type of [interface{}].
	Input   interface{}     `json:"input"`
	IsError bool            `json:"isError"`
	Name    shared.ToolName `json:"name"`
	// This field can have the runtime type of [map[string]map[string]string].
	ProviderMetadata interface{}      `json:"providerMetadata"`
	Text             string           `json:"text"`
	ToolUseID        string           `json:"toolUseId"`
	JSON             contentBlockJSON `json:"-"`
	union            ContentBlockUnion
}

// contentBlockJSON contains the JSON metadata for the struct [ContentBlock]
type contentBlockJSON struct {
	Type             apijson.Field
	ID               apijson.Field
	Content          apijson.Field
	DeniedByUser     apijson.Field
	Input            apijson.Field
	IsError          apijson.Field
	Name             apijson.Field
	ProviderMetadata apijson.Field
	Text             apijson.Field
	ToolUseID        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r contentBlockJSON) RawJSON() string {
	return r.raw
}

func (r *ContentBlock) UnmarshalJSON(data []byte) (err error) {
	*r = ContentBlock{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ContentBlockUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ContentBlockTextContentBlock],
// [ContentBlockToolUseContentBlock], [ContentBlockToolResultContentBlock],
// [ContentBlockServerToolUseContentBlock],
// [ContentBlockWebSearchToolResultContentBlock].
func (r ContentBlock) AsUnion() ContentBlockUnion {
	return r.union
}

// Message content block. The `type` field distinguishes text, tool use, tool
// result, server tool use, and web search result blocks.
//
// Union satisfied by [ContentBlockTextContentBlock],
// [ContentBlockToolUseContentBlock], [ContentBlockToolResultContentBlock],
// [ContentBlockServerToolUseContentBlock] or
// [ContentBlockWebSearchToolResultContentBlock].
type ContentBlockUnion interface {
	implementsContentBlock()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ContentBlockUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ContentBlockTextContentBlock{}),
			DiscriminatorValue: "text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ContentBlockToolUseContentBlock{}),
			DiscriminatorValue: "tool_use",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ContentBlockToolResultContentBlock{}),
			DiscriminatorValue: "tool_result",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ContentBlockServerToolUseContentBlock{}),
			DiscriminatorValue: "server_tool_use",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ContentBlockWebSearchToolResultContentBlock{}),
			DiscriminatorValue: "web_search_tool_result",
		},
	)
}

type ContentBlockTextContentBlock struct {
	Text             string                           `json:"text" api:"required"`
	Type             ContentBlockTextContentBlockType `json:"type" api:"required"`
	ProviderMetadata map[string]map[string]string     `json:"providerMetadata"`
	JSON             contentBlockTextContentBlockJSON `json:"-"`
}

// contentBlockTextContentBlockJSON contains the JSON metadata for the struct
// [ContentBlockTextContentBlock]
type contentBlockTextContentBlockJSON struct {
	Text             apijson.Field
	Type             apijson.Field
	ProviderMetadata apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ContentBlockTextContentBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentBlockTextContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r ContentBlockTextContentBlock) implementsContentBlock() {}

type ContentBlockTextContentBlockType string

const (
	ContentBlockTextContentBlockTypeText ContentBlockTextContentBlockType = "text"
)

func (r ContentBlockTextContentBlockType) IsKnown() bool {
	switch r {
	case ContentBlockTextContentBlockTypeText:
		return true
	}
	return false
}

type ContentBlockToolUseContentBlock struct {
	ID string `json:"id" api:"required"`
	// Any JSON value. Generated SDKs may expose this value boundary as unknown or Any.
	Input            interface{}                         `json:"input" api:"required"`
	Name             shared.ToolName                     `json:"name" api:"required"`
	Type             ContentBlockToolUseContentBlockType `json:"type" api:"required"`
	ProviderMetadata map[string]map[string]string        `json:"providerMetadata"`
	JSON             contentBlockToolUseContentBlockJSON `json:"-"`
}

// contentBlockToolUseContentBlockJSON contains the JSON metadata for the struct
// [ContentBlockToolUseContentBlock]
type contentBlockToolUseContentBlockJSON struct {
	ID               apijson.Field
	Input            apijson.Field
	Name             apijson.Field
	Type             apijson.Field
	ProviderMetadata apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ContentBlockToolUseContentBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentBlockToolUseContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r ContentBlockToolUseContentBlock) implementsContentBlock() {}

type ContentBlockToolUseContentBlockType string

const (
	ContentBlockToolUseContentBlockTypeToolUse ContentBlockToolUseContentBlockType = "tool_use"
)

func (r ContentBlockToolUseContentBlockType) IsKnown() bool {
	switch r {
	case ContentBlockToolUseContentBlockTypeToolUse:
		return true
	}
	return false
}

type ContentBlockToolResultContentBlock struct {
	Content          string                                 `json:"content" api:"required"`
	IsError          bool                                   `json:"isError" api:"required"`
	ToolUseID        string                                 `json:"toolUseId" api:"required"`
	Type             ContentBlockToolResultContentBlockType `json:"type" api:"required"`
	DeniedByUser     bool                                   `json:"deniedByUser"`
	ProviderMetadata map[string]map[string]string           `json:"providerMetadata"`
	JSON             contentBlockToolResultContentBlockJSON `json:"-"`
}

// contentBlockToolResultContentBlockJSON contains the JSON metadata for the struct
// [ContentBlockToolResultContentBlock]
type contentBlockToolResultContentBlockJSON struct {
	Content          apijson.Field
	IsError          apijson.Field
	ToolUseID        apijson.Field
	Type             apijson.Field
	DeniedByUser     apijson.Field
	ProviderMetadata apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ContentBlockToolResultContentBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentBlockToolResultContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r ContentBlockToolResultContentBlock) implementsContentBlock() {}

type ContentBlockToolResultContentBlockType string

const (
	ContentBlockToolResultContentBlockTypeToolResult ContentBlockToolResultContentBlockType = "tool_result"
)

func (r ContentBlockToolResultContentBlockType) IsKnown() bool {
	switch r {
	case ContentBlockToolResultContentBlockTypeToolResult:
		return true
	}
	return false
}

type ContentBlockServerToolUseContentBlock struct {
	ID string `json:"id" api:"required"`
	// Any JSON value. Generated SDKs may expose this value boundary as unknown or Any.
	Input            interface{}                               `json:"input" api:"required"`
	Name             string                                    `json:"name" api:"required"`
	Type             ContentBlockServerToolUseContentBlockType `json:"type" api:"required"`
	ProviderMetadata map[string]map[string]string              `json:"providerMetadata"`
	JSON             contentBlockServerToolUseContentBlockJSON `json:"-"`
}

// contentBlockServerToolUseContentBlockJSON contains the JSON metadata for the
// struct [ContentBlockServerToolUseContentBlock]
type contentBlockServerToolUseContentBlockJSON struct {
	ID               apijson.Field
	Input            apijson.Field
	Name             apijson.Field
	Type             apijson.Field
	ProviderMetadata apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ContentBlockServerToolUseContentBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentBlockServerToolUseContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r ContentBlockServerToolUseContentBlock) implementsContentBlock() {}

type ContentBlockServerToolUseContentBlockType string

const (
	ContentBlockServerToolUseContentBlockTypeServerToolUse ContentBlockServerToolUseContentBlockType = "server_tool_use"
)

func (r ContentBlockServerToolUseContentBlockType) IsKnown() bool {
	switch r {
	case ContentBlockServerToolUseContentBlockTypeServerToolUse:
		return true
	}
	return false
}

type ContentBlockWebSearchToolResultContentBlock struct {
	// Web search result payload. The runtime returns either an array of web search
	// results or an error object.
	Content   interface{}                                     `json:"content" api:"required"`
	ToolUseID string                                          `json:"toolUseId" api:"required"`
	Type      ContentBlockWebSearchToolResultContentBlockType `json:"type" api:"required"`
	JSON      contentBlockWebSearchToolResultContentBlockJSON `json:"-"`
}

// contentBlockWebSearchToolResultContentBlockJSON contains the JSON metadata for
// the struct [ContentBlockWebSearchToolResultContentBlock]
type contentBlockWebSearchToolResultContentBlockJSON struct {
	Content     apijson.Field
	ToolUseID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentBlockWebSearchToolResultContentBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentBlockWebSearchToolResultContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r ContentBlockWebSearchToolResultContentBlock) implementsContentBlock() {}

type ContentBlockWebSearchToolResultContentBlockType string

const (
	ContentBlockWebSearchToolResultContentBlockTypeWebSearchToolResult ContentBlockWebSearchToolResultContentBlockType = "web_search_tool_result"
)

func (r ContentBlockWebSearchToolResultContentBlockType) IsKnown() bool {
	switch r {
	case ContentBlockWebSearchToolResultContentBlockTypeWebSearchToolResult:
		return true
	}
	return false
}

type ContentBlockType string

const (
	ContentBlockTypeText                ContentBlockType = "text"
	ContentBlockTypeToolUse             ContentBlockType = "tool_use"
	ContentBlockTypeToolResult          ContentBlockType = "tool_result"
	ContentBlockTypeServerToolUse       ContentBlockType = "server_tool_use"
	ContentBlockTypeWebSearchToolResult ContentBlockType = "web_search_tool_result"
)

func (r ContentBlockType) IsKnown() bool {
	switch r {
	case ContentBlockTypeText, ContentBlockTypeToolUse, ContentBlockTypeToolResult, ContentBlockTypeServerToolUse, ContentBlockTypeWebSearchToolResult:
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

type MessagePage struct {
	Cursor   string          `json:"cursor" api:"required,nullable"`
	HasMore  bool            `json:"hasMore" api:"required"`
	Messages []Message       `json:"messages" api:"required"`
	JSON     messagePageJSON `json:"-"`
}

// messagePageJSON contains the JSON metadata for the struct [MessagePage]
type messagePageJSON struct {
	Cursor      apijson.Field
	HasMore     apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagePage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagePageJSON) RawJSON() string {
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

type SubThreadSummary struct {
	ID             string  `json:"id" api:"required"`
	CompletedAt    string  `json:"completedAt" api:"required,nullable"`
	CreatedAt      string  `json:"createdAt" api:"required"`
	MessageCount   float64 `json:"messageCount" api:"required"`
	Model          string  `json:"model" api:"required"`
	ParentThreadID string  `json:"parentThreadId" api:"required,nullable"`
	Result         string  `json:"result" api:"required,nullable"`
	ScheduleID     string  `json:"scheduleId" api:"required,nullable"`
	ScheduleSeq    float64 `json:"scheduleSeq" api:"required,nullable"`
	// `pending` means the parent thread is still waiting on this sub-thread. Other
	// values are the child thread lifecycle state.
	State     SubThreadSummaryState `json:"state" api:"required"`
	StepCount float64               `json:"stepCount" api:"required"`
	ToolUseID string                `json:"toolUseId" api:"required,nullable"`
	JSON      subThreadSummaryJSON  `json:"-"`
}

// subThreadSummaryJSON contains the JSON metadata for the struct
// [SubThreadSummary]
type subThreadSummaryJSON struct {
	ID             apijson.Field
	CompletedAt    apijson.Field
	CreatedAt      apijson.Field
	MessageCount   apijson.Field
	Model          apijson.Field
	ParentThreadID apijson.Field
	Result         apijson.Field
	ScheduleID     apijson.Field
	ScheduleSeq    apijson.Field
	State          apijson.Field
	StepCount      apijson.Field
	ToolUseID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubThreadSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subThreadSummaryJSON) RawJSON() string {
	return r.raw
}

// `pending` means the parent thread is still waiting on this sub-thread. Other
// values are the child thread lifecycle state.
type SubThreadSummaryState string

const (
	SubThreadSummaryStatePending  SubThreadSummaryState = "pending"
	SubThreadSummaryStateIdle     SubThreadSummaryState = "idle"
	SubThreadSummaryStateRunning  SubThreadSummaryState = "running"
	SubThreadSummaryStateAwaiting SubThreadSummaryState = "awaiting"
	SubThreadSummaryStateClosed   SubThreadSummaryState = "closed"
)

func (r SubThreadSummaryState) IsKnown() bool {
	switch r {
	case SubThreadSummaryStatePending, SubThreadSummaryStateIdle, SubThreadSummaryStateRunning, SubThreadSummaryStateAwaiting, SubThreadSummaryStateClosed:
		return true
	}
	return false
}

type Thread struct {
	ID              string               `json:"id" api:"required"`
	AgentID         string               `json:"agentId" api:"required"`
	CompiledContext CompiledContext      `json:"compiledContext" api:"required,nullable"`
	CompletedAt     string               `json:"completedAt" api:"required,nullable"`
	CreatedAt       string               `json:"createdAt" api:"required"`
	Depth           float64              `json:"depth" api:"required"`
	Error           string               `json:"error" api:"required,nullable"`
	HasMoreMessages bool                 `json:"hasMoreMessages" api:"required"`
	Instructions    string               `json:"instructions" api:"required,nullable"`
	LastTurnStatus  ThreadLastTurnStatus `json:"lastTurnStatus" api:"required,nullable"`
	Message         string               `json:"message" api:"required"`
	MessageCursor   float64              `json:"messageCursor" api:"required,nullable"`
	Messages        []Message            `json:"messages" api:"required"`
	Model           string               `json:"model" api:"required"`
	ParentThreadID  string               `json:"parentThreadId" api:"required,nullable"`
	Result          string               `json:"result" api:"required,nullable"`
	ScheduleID      string               `json:"scheduleId" api:"required,nullable"`
	ScheduleSeq     float64              `json:"scheduleSeq" api:"required,nullable"`
	// `idle` threads can accept a new turn or be closed. `running` threads have an
	// active turn. `awaiting` threads are paused on external input such as approvals.
	// `closed` threads are terminal.
	Status                 Status                            `json:"status" api:"required"`
	SubThreads             []SubThreadSummary                `json:"subThreads" api:"required"`
	Tools                  []shared.ToolName                 `json:"tools" api:"required"`
	Turns                  []Turn                            `json:"turns" api:"required"`
	UpdatedAt              string                            `json:"updatedAt" api:"required"`
	ExternalToolNamespaces ThreadExternalToolNamespacesUnion `json:"externalToolNamespaces"`
	JSON                   threadJSON                        `json:"-"`
}

// threadJSON contains the JSON metadata for the struct [Thread]
type threadJSON struct {
	ID                     apijson.Field
	AgentID                apijson.Field
	CompiledContext        apijson.Field
	CompletedAt            apijson.Field
	CreatedAt              apijson.Field
	Depth                  apijson.Field
	Error                  apijson.Field
	HasMoreMessages        apijson.Field
	Instructions           apijson.Field
	LastTurnStatus         apijson.Field
	Message                apijson.Field
	MessageCursor          apijson.Field
	Messages               apijson.Field
	Model                  apijson.Field
	ParentThreadID         apijson.Field
	Result                 apijson.Field
	ScheduleID             apijson.Field
	ScheduleSeq            apijson.Field
	Status                 apijson.Field
	SubThreads             apijson.Field
	Tools                  apijson.Field
	Turns                  apijson.Field
	UpdatedAt              apijson.Field
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

type ThreadTurnSummary struct {
	Activity     string                  `json:"activity" api:"required,nullable"`
	CompletedAt  string                  `json:"completedAt" api:"required"`
	MessageCount float64                 `json:"messageCount" api:"required"`
	NextStep     string                  `json:"nextStep" api:"required,nullable"`
	Status       ThreadTurnSummaryStatus `json:"status" api:"required"`
	TurnSeq      float64                 `json:"turnSeq" api:"required"`
	JSON         threadTurnSummaryJSON   `json:"-"`
}

// threadTurnSummaryJSON contains the JSON metadata for the struct
// [ThreadTurnSummary]
type threadTurnSummaryJSON struct {
	Activity     apijson.Field
	CompletedAt  apijson.Field
	MessageCount apijson.Field
	NextStep     apijson.Field
	Status       apijson.Field
	TurnSeq      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreadTurnSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadTurnSummaryJSON) RawJSON() string {
	return r.raw
}

type ThreadTurnSummaryStatus string

const (
	ThreadTurnSummaryStatusCompleted ThreadTurnSummaryStatus = "completed"
	ThreadTurnSummaryStatusFailed    ThreadTurnSummaryStatus = "failed"
)

func (r ThreadTurnSummaryStatus) IsKnown() bool {
	switch r {
	case ThreadTurnSummaryStatusCompleted, ThreadTurnSummaryStatusFailed:
		return true
	}
	return false
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
	Message         string     `json:"message" api:"required"`
	Result          string     `json:"result" api:"required,nullable"`
	Seq             float64    `json:"seq" api:"required"`
	StartedAt       string     `json:"startedAt" api:"required"`
	StartMessageSeq float64    `json:"startMessageSeq" api:"required"`
	Status          TurnStatus `json:"status" api:"required"`
	ThreadID        string     `json:"threadId" api:"required"`
	TokenUsage      TokenUsage `json:"tokenUsage" api:"required,nullable"`
	JSON            turnJSON   `json:"-"`
}

// turnJSON contains the JSON metadata for the struct [Turn]
type turnJSON struct {
	ID              apijson.Field
	CompletedAt     apijson.Field
	EndMessageSeq   apijson.Field
	Error           apijson.Field
	Message         apijson.Field
	Result          apijson.Field
	Seq             apijson.Field
	StartedAt       apijson.Field
	StartMessageSeq apijson.Field
	Status          apijson.Field
	ThreadID        apijson.Field
	TokenUsage      apijson.Field
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
	Message      param.Field[string] `json:"message"`
	Model        param.Field[string] `json:"model"`
	// Deprecated alias for `instructions`; accepted for backwards compatibility.
	SystemPrompt param.Field[string] `json:"systemPrompt"`
	// Per-thread tool subset. Omit to inherit the agent's full effective tools; pass
	// [] to run with no configurable tools. Provided entries can only narrow the
	// agent's effective tools.
	Tools param.Field[[]shared.ToolSpecParam] `json:"tools"`
}

func (r ThreadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreadGetParams struct {
	// When true, includes debug-only compiled context fields.
	Debug param.Field[ThreadGetParamsDebug] `query:"debug"`
	// Deprecated compatibility flag. Thread detail includes a bounded recent message
	// page by default; pass `false` only to opt out when no message pagination params
	// are present.
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

// Deprecated compatibility flag. Thread detail includes a bounded recent message
// page by default; pass `false` only to opt out when no message pagination params
// are present.
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

type ThreadActivityParams struct {
	// Optional fleet id for index-backed authorization.
	FleetID param.Field[string] `query:"fleetId"`
}

// URLQuery serializes [ThreadActivityParams]'s query parameters as `url.Values`.
func (r ThreadActivityParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ThreadListMessagesParams struct {
	// Cursor returned by a previous thread messages response.
	Cursor param.Field[string] `query:"cursor"`
	// Optional fleet id for index-backed authorization.
	FleetID param.Field[string] `query:"fleetId"`
	// Maximum number of messages to include, capped at 500.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [ThreadListMessagesParams]'s query parameters as
// `url.Values`.
func (r ThreadListMessagesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ThreadStartTurnParams struct {
	Message param.Field[string] `json:"message" api:"required"`
	Model   param.Field[string] `json:"model"`
	// Per-turn tool subset. Omit to inherit the thread's current tools; pass [] to run
	// this turn with no configurable tools. Provided entries can only narrow the
	// agent/thread effective tools.
	Tools param.Field[[]shared.ToolSpecParam] `json:"tools"`
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
