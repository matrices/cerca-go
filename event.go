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
	"github.com/matrices/cerca-go/shared"
)

// EventService contains methods and other services that help with interacting with
// the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	return
}

func (r *EventService) ListForAgent(ctx context.Context, agentID string, query EventListForAgentParams, opts ...option.RequestOption) (res *pagination.EventsCursorPage[SubscriptionEvent], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/events", agentID)
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

func (r *EventService) ListForAgentAutoPaging(ctx context.Context, agentID string, query EventListForAgentParams, opts ...option.RequestOption) *pagination.EventsCursorPageAutoPager[SubscriptionEvent] {
	return pagination.NewEventsCursorPageAutoPager(r.ListForAgent(ctx, agentID, query, opts...))
}

func (r *EventService) ListForFleet(ctx context.Context, fleetID string, query EventListForFleetParams, opts ...option.RequestOption) (res *pagination.EventsCursorPage[SubscriptionEvent], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("fleets/%s/events", fleetID)
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

func (r *EventService) ListForFleetAutoPaging(ctx context.Context, fleetID string, query EventListForFleetParams, opts ...option.RequestOption) *pagination.EventsCursorPageAutoPager[SubscriptionEvent] {
	return pagination.NewEventsCursorPageAutoPager(r.ListForFleet(ctx, fleetID, query, opts...))
}

func (r *EventService) ListForThread(ctx context.Context, agentID string, threadID string, query EventListForThreadParams, opts ...option.RequestOption) (res *pagination.EventsCursorPage[SubscriptionEvent], err error) {
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
	path := fmt.Sprintf("agents/%s/threads/%s/events", agentID, threadID)
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

func (r *EventService) ListForThreadAutoPaging(ctx context.Context, agentID string, threadID string, query EventListForThreadParams, opts ...option.RequestOption) *pagination.EventsCursorPageAutoPager[SubscriptionEvent] {
	return pagination.NewEventsCursorPageAutoPager(r.ListForThread(ctx, agentID, threadID, query, opts...))
}

// WebSocket upgrade endpoint. Set
// `Sec-WebSocket-Protocol: agent-v1, agent-auth-<API_KEY>` so the runtime can
// authenticate the stream while preserving the public subprotocol. HTTP clients
// that cannot upgrade should use `/agents/{agentId}/events` as the polling analog.
func (r *EventService) SubscribeAgent(ctx context.Context, agentID string, opts ...option.RequestOption) (res *shared.ErrorResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/events/subscribe", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// WebSocket upgrade endpoint. Set
// `Sec-WebSocket-Protocol: agent-v1, agent-auth-<API_KEY>` so the runtime can
// authenticate the stream while preserving the public subprotocol. HTTP clients
// that cannot upgrade should use `/fleets/{fleetId}/events` as the polling analog.
func (r *EventService) SubscribeFleet(ctx context.Context, fleetID string, opts ...option.RequestOption) (res *shared.ErrorResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("fleets/%s/events/subscribe", fleetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// WebSocket upgrade endpoint. Set
// `Sec-WebSocket-Protocol: agent-v1, agent-auth-<API_KEY>` so the runtime can
// authenticate the stream while preserving the public subprotocol. HTTP clients
// that cannot upgrade should use `/agents/{agentId}/threads/{threadId}/events` as
// the polling analog.
func (r *EventService) SubscribeThread(ctx context.Context, agentID string, threadID string, opts ...option.RequestOption) (res *shared.ErrorResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/threads/%s/events/subscribe", agentID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type RuntimeWebhookEvent struct {
	ID        string                  `json:"id" api:"required"`
	AgentID   string                  `json:"agentId" api:"required"`
	Data      map[string]interface{}  `json:"data" api:"required"`
	Event     RuntimeWebhookEventType `json:"event" api:"required"`
	FleetID   string                  `json:"fleetId" api:"required"`
	OrgID     string                  `json:"orgId" api:"required"`
	ThreadID  string                  `json:"threadId" api:"required,nullable"`
	Timestamp string                  `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventJSON `json:"-"`
}

// runtimeWebhookEventJSON contains the JSON metadata for the struct
// [RuntimeWebhookEvent]
type runtimeWebhookEventJSON struct {
	ID          apijson.Field
	AgentID     apijson.Field
	Data        apijson.Field
	Event       apijson.Field
	FleetID     apijson.Field
	OrgID       apijson.Field
	ThreadID    apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventType string

const (
	RuntimeWebhookEventTypeAgentCreated        RuntimeWebhookEventType = "agent.created"
	RuntimeWebhookEventTypeAgentUpdated        RuntimeWebhookEventType = "agent.updated"
	RuntimeWebhookEventTypeAgentDeleted        RuntimeWebhookEventType = "agent.deleted"
	RuntimeWebhookEventTypeThreadCreated       RuntimeWebhookEventType = "thread.created"
	RuntimeWebhookEventTypeThreadStatusChanged RuntimeWebhookEventType = "thread.status.changed"
	RuntimeWebhookEventTypeThreadCompleted     RuntimeWebhookEventType = "thread.completed"
	RuntimeWebhookEventTypeThreadFailed        RuntimeWebhookEventType = "thread.failed"
	RuntimeWebhookEventTypeTurnCreated         RuntimeWebhookEventType = "turn.created"
	RuntimeWebhookEventTypeTurnCompleted       RuntimeWebhookEventType = "turn.completed"
	RuntimeWebhookEventTypeTurnFailed          RuntimeWebhookEventType = "turn.failed"
	RuntimeWebhookEventTypeMessageCreated      RuntimeWebhookEventType = "message.created"
	RuntimeWebhookEventTypeApprovalRequested   RuntimeWebhookEventType = "approval.requested"
	RuntimeWebhookEventTypeApprovalResolved    RuntimeWebhookEventType = "approval.resolved"
	RuntimeWebhookEventTypeApprovalGranted     RuntimeWebhookEventType = "approval.granted"
	RuntimeWebhookEventTypeScheduleCreated     RuntimeWebhookEventType = "schedule.created"
	RuntimeWebhookEventTypeScheduleDeleted     RuntimeWebhookEventType = "schedule.deleted"
	RuntimeWebhookEventTypeScheduleTriggered   RuntimeWebhookEventType = "schedule.triggered"
	RuntimeWebhookEventTypeConnectionAttached  RuntimeWebhookEventType = "connection.attached"
	RuntimeWebhookEventTypeConnectionDetached  RuntimeWebhookEventType = "connection.detached"
	RuntimeWebhookEventTypeWebhookTest         RuntimeWebhookEventType = "webhook.test"
)

func (r RuntimeWebhookEventType) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventTypeAgentCreated, RuntimeWebhookEventTypeAgentUpdated, RuntimeWebhookEventTypeAgentDeleted, RuntimeWebhookEventTypeThreadCreated, RuntimeWebhookEventTypeThreadStatusChanged, RuntimeWebhookEventTypeThreadCompleted, RuntimeWebhookEventTypeThreadFailed, RuntimeWebhookEventTypeTurnCreated, RuntimeWebhookEventTypeTurnCompleted, RuntimeWebhookEventTypeTurnFailed, RuntimeWebhookEventTypeMessageCreated, RuntimeWebhookEventTypeApprovalRequested, RuntimeWebhookEventTypeApprovalResolved, RuntimeWebhookEventTypeApprovalGranted, RuntimeWebhookEventTypeScheduleCreated, RuntimeWebhookEventTypeScheduleDeleted, RuntimeWebhookEventTypeScheduleTriggered, RuntimeWebhookEventTypeConnectionAttached, RuntimeWebhookEventTypeConnectionDetached, RuntimeWebhookEventTypeWebhookTest:
		return true
	}
	return false
}

type SubscriptionEvent struct {
	Event RuntimeWebhookEvent   `json:"event" api:"required"`
	Seq   float64               `json:"seq" api:"required"`
	JSON  subscriptionEventJSON `json:"-"`
}

// subscriptionEventJSON contains the JSON metadata for the struct
// [SubscriptionEvent]
type subscriptionEventJSON struct {
	Event       apijson.Field
	Seq         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionEventJSON) RawJSON() string {
	return r.raw
}

type EventListForAgentParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Comma-separated event type filter.
	Events param.Field[string] `query:"events"`
	// When true, starts from the beginning of the retained buffer.
	History param.Field[EventListForAgentParamsHistory] `query:"history"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [EventListForAgentParams]'s query parameters as
// `url.Values`.
func (r EventListForAgentParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// When true, starts from the beginning of the retained buffer.
type EventListForAgentParamsHistory string

const (
	EventListForAgentParamsHistoryTrue  EventListForAgentParamsHistory = "true"
	EventListForAgentParamsHistoryFalse EventListForAgentParamsHistory = "false"
)

func (r EventListForAgentParamsHistory) IsKnown() bool {
	switch r {
	case EventListForAgentParamsHistoryTrue, EventListForAgentParamsHistoryFalse:
		return true
	}
	return false
}

type EventListForFleetParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Comma-separated event type filter.
	Events param.Field[string] `query:"events"`
	// When true, starts from the beginning of the retained buffer.
	History param.Field[EventListForFleetParamsHistory] `query:"history"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [EventListForFleetParams]'s query parameters as
// `url.Values`.
func (r EventListForFleetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// When true, starts from the beginning of the retained buffer.
type EventListForFleetParamsHistory string

const (
	EventListForFleetParamsHistoryTrue  EventListForFleetParamsHistory = "true"
	EventListForFleetParamsHistoryFalse EventListForFleetParamsHistory = "false"
)

func (r EventListForFleetParamsHistory) IsKnown() bool {
	switch r {
	case EventListForFleetParamsHistoryTrue, EventListForFleetParamsHistoryFalse:
		return true
	}
	return false
}

type EventListForThreadParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Comma-separated event type filter.
	Events param.Field[string] `query:"events"`
	// When true, starts from the beginning of the retained buffer.
	History param.Field[EventListForThreadParamsHistory] `query:"history"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [EventListForThreadParams]'s query parameters as
// `url.Values`.
func (r EventListForThreadParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// When true, starts from the beginning of the retained buffer.
type EventListForThreadParamsHistory string

const (
	EventListForThreadParamsHistoryTrue  EventListForThreadParamsHistory = "true"
	EventListForThreadParamsHistoryFalse EventListForThreadParamsHistory = "false"
)

func (r EventListForThreadParamsHistory) IsKnown() bool {
	switch r {
	case EventListForThreadParamsHistoryTrue, EventListForThreadParamsHistoryFalse:
		return true
	}
	return false
}
