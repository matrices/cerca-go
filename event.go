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
	"github.com/matrices/cerca-go/packages/ssestream"
	"github.com/matrices/cerca-go/shared"
	"github.com/tidwall/gjson"
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

// List events
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

// List events
func (r *EventService) ListForAgentAutoPaging(ctx context.Context, agentID string, query EventListForAgentParams, opts ...option.RequestOption) *pagination.EventsCursorPageAutoPager[SubscriptionEvent] {
	return pagination.NewEventsCursorPageAutoPager(r.ListForAgent(ctx, agentID, query, opts...))
}

// List events
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

// List events
func (r *EventService) ListForFleetAutoPaging(ctx context.Context, fleetID string, query EventListForFleetParams, opts ...option.RequestOption) *pagination.EventsCursorPageAutoPager[SubscriptionEvent] {
	return pagination.NewEventsCursorPageAutoPager(r.ListForFleet(ctx, fleetID, query, opts...))
}

// List events
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

// List events
func (r *EventService) ListForThreadAutoPaging(ctx context.Context, agentID string, threadID string, query EventListForThreadParams, opts ...option.RequestOption) *pagination.EventsCursorPageAutoPager[SubscriptionEvent] {
	return pagination.NewEventsCursorPageAutoPager(r.ListForThread(ctx, agentID, threadID, query, opts...))
}

// Server-Sent Events stream. Each SSE data frame is JSON matching this response
// schema.
func (r *EventService) StreamForAgentStreaming(ctx context.Context, agentID string, params EventStreamForAgentParams, opts ...option.RequestOption) (stream *ssestream.Stream[SubscriptionEvent]) {
	var (
		raw *http.Response
		err error
	)
	if params.LastEventID.Present {
		opts = append(opts, option.WithHeader("Last-Event-ID", fmt.Sprintf("%v", params.LastEventID)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return ssestream.NewStream[SubscriptionEvent](nil, err)
	}
	path := fmt.Sprintf("agents/%s/events/stream", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &raw, opts...)
	return ssestream.NewStream[SubscriptionEvent](ssestream.NewDecoder(raw), err)
}

// Server-Sent Events stream. Each SSE data frame is JSON matching this response
// schema.
func (r *EventService) StreamForFleetStreaming(ctx context.Context, fleetID string, params EventStreamForFleetParams, opts ...option.RequestOption) (stream *ssestream.Stream[SubscriptionEvent]) {
	var (
		raw *http.Response
		err error
	)
	if params.LastEventID.Present {
		opts = append(opts, option.WithHeader("Last-Event-ID", fmt.Sprintf("%v", params.LastEventID)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return ssestream.NewStream[SubscriptionEvent](nil, err)
	}
	path := fmt.Sprintf("fleets/%s/events/stream", fleetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &raw, opts...)
	return ssestream.NewStream[SubscriptionEvent](ssestream.NewDecoder(raw), err)
}

// Server-Sent Events stream. Each SSE data frame is JSON matching this response
// schema.
func (r *EventService) StreamForThreadStreaming(ctx context.Context, agentID string, threadID string, opts ...option.RequestOption) (stream *ssestream.Stream[ThreadStreamEvent]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return ssestream.NewStream[ThreadStreamEvent](nil, err)
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return ssestream.NewStream[ThreadStreamEvent](nil, err)
	}
	path := fmt.Sprintf("agents/%s/threads/%s/stream", agentID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	return ssestream.NewStream[ThreadStreamEvent](ssestream.NewDecoder(raw), err)
}

// Server-Sent Events stream. Each SSE data frame is JSON matching this response
// schema.
func (r *EventService) StreamForThreadEventsStreaming(ctx context.Context, agentID string, threadID string, params EventStreamForThreadEventsParams, opts ...option.RequestOption) (stream *ssestream.Stream[SubscriptionEvent]) {
	var (
		raw *http.Response
		err error
	)
	if params.LastEventID.Present {
		opts = append(opts, option.WithHeader("Last-Event-ID", fmt.Sprintf("%v", params.LastEventID)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return ssestream.NewStream[SubscriptionEvent](nil, err)
	}
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return ssestream.NewStream[SubscriptionEvent](nil, err)
	}
	path := fmt.Sprintf("agents/%s/threads/%s/events/stream", agentID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &raw, opts...)
	return ssestream.NewStream[SubscriptionEvent](ssestream.NewDecoder(raw), err)
}

type RuntimeWebhookEvent struct {
	ID      string `json:"id" api:"required"`
	AgentID string `json:"agentId" api:"required"`
	// This field can have the runtime type of
	// [RuntimeWebhookEventAgentCreatedWebhookEventData],
	// [RuntimeWebhookEventAgentUpdatedWebhookEventData],
	// [RuntimeWebhookEventAgentDeletedWebhookEventData],
	// [RuntimeWebhookEventThreadCreatedWebhookEventData],
	// [RuntimeWebhookEventThreadStatusChangedWebhookEventData],
	// [RuntimeWebhookEventThreadCompletedWebhookEventData],
	// [RuntimeWebhookEventThreadFailedWebhookEventData],
	// [RuntimeWebhookEventTurnCreatedWebhookEventData],
	// [RuntimeWebhookEventTurnCompletedWebhookEventData],
	// [RuntimeWebhookEventTurnFailedWebhookEventData],
	// [RuntimeWebhookEventMessageCreatedWebhookEventData],
	// [RuntimeWebhookEventApprovalRequestedWebhookEventData],
	// [RuntimeWebhookEventApprovalResolvedWebhookEventData],
	// [RuntimeWebhookEventApprovalGrantedWebhookEventData],
	// [RuntimeWebhookEventScheduleCreatedWebhookEventData],
	// [RuntimeWebhookEventScheduleDeletedWebhookEventData],
	// [RuntimeWebhookEventScheduleTriggeredWebhookEventData],
	// [RuntimeWebhookEventConnectionAttachedWebhookEventData],
	// [RuntimeWebhookEventConnectionDetachedWebhookEventData],
	// [RuntimeWebhookEventWebhookTestWebhookEventData].
	Data      interface{}              `json:"data" api:"required"`
	Event     RuntimeWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                   `json:"fleetId" api:"required"`
	OrgID     string                   `json:"orgId" api:"required"`
	ThreadID  string                   `json:"threadId" api:"required,nullable"`
	Timestamp string                   `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventJSON  `json:"-"`
	union     RuntimeWebhookEventUnion
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

func (r runtimeWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r *RuntimeWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	*r = RuntimeWebhookEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RuntimeWebhookEventUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [RuntimeWebhookEventAgentCreatedWebhookEvent],
// [RuntimeWebhookEventAgentUpdatedWebhookEvent],
// [RuntimeWebhookEventAgentDeletedWebhookEvent],
// [RuntimeWebhookEventThreadCreatedWebhookEvent],
// [RuntimeWebhookEventThreadStatusChangedWebhookEvent],
// [RuntimeWebhookEventThreadCompletedWebhookEvent],
// [RuntimeWebhookEventThreadFailedWebhookEvent],
// [RuntimeWebhookEventTurnCreatedWebhookEvent],
// [RuntimeWebhookEventTurnCompletedWebhookEvent],
// [RuntimeWebhookEventTurnFailedWebhookEvent],
// [RuntimeWebhookEventMessageCreatedWebhookEvent],
// [RuntimeWebhookEventApprovalRequestedWebhookEvent],
// [RuntimeWebhookEventApprovalResolvedWebhookEvent],
// [RuntimeWebhookEventApprovalGrantedWebhookEvent],
// [RuntimeWebhookEventScheduleCreatedWebhookEvent],
// [RuntimeWebhookEventScheduleDeletedWebhookEvent],
// [RuntimeWebhookEventScheduleTriggeredWebhookEvent],
// [RuntimeWebhookEventConnectionAttachedWebhookEvent],
// [RuntimeWebhookEventConnectionDetachedWebhookEvent],
// [RuntimeWebhookEventWebhookTestWebhookEvent].
func (r RuntimeWebhookEvent) AsUnion() RuntimeWebhookEventUnion {
	return r.union
}

// Union satisfied by [RuntimeWebhookEventAgentCreatedWebhookEvent],
// [RuntimeWebhookEventAgentUpdatedWebhookEvent],
// [RuntimeWebhookEventAgentDeletedWebhookEvent],
// [RuntimeWebhookEventThreadCreatedWebhookEvent],
// [RuntimeWebhookEventThreadStatusChangedWebhookEvent],
// [RuntimeWebhookEventThreadCompletedWebhookEvent],
// [RuntimeWebhookEventThreadFailedWebhookEvent],
// [RuntimeWebhookEventTurnCreatedWebhookEvent],
// [RuntimeWebhookEventTurnCompletedWebhookEvent],
// [RuntimeWebhookEventTurnFailedWebhookEvent],
// [RuntimeWebhookEventMessageCreatedWebhookEvent],
// [RuntimeWebhookEventApprovalRequestedWebhookEvent],
// [RuntimeWebhookEventApprovalResolvedWebhookEvent],
// [RuntimeWebhookEventApprovalGrantedWebhookEvent],
// [RuntimeWebhookEventScheduleCreatedWebhookEvent],
// [RuntimeWebhookEventScheduleDeletedWebhookEvent],
// [RuntimeWebhookEventScheduleTriggeredWebhookEvent],
// [RuntimeWebhookEventConnectionAttachedWebhookEvent],
// [RuntimeWebhookEventConnectionDetachedWebhookEvent] or
// [RuntimeWebhookEventWebhookTestWebhookEvent].
type RuntimeWebhookEventUnion interface {
	implementsRuntimeWebhookEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuntimeWebhookEventUnion)(nil)).Elem(),
		"event",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventAgentCreatedWebhookEvent{}),
			DiscriminatorValue: "agent.created",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventAgentUpdatedWebhookEvent{}),
			DiscriminatorValue: "agent.updated",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventAgentDeletedWebhookEvent{}),
			DiscriminatorValue: "agent.deleted",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventThreadCreatedWebhookEvent{}),
			DiscriminatorValue: "thread.created",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventThreadStatusChangedWebhookEvent{}),
			DiscriminatorValue: "thread.status.changed",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventThreadCompletedWebhookEvent{}),
			DiscriminatorValue: "thread.completed",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventThreadFailedWebhookEvent{}),
			DiscriminatorValue: "thread.failed",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventTurnCreatedWebhookEvent{}),
			DiscriminatorValue: "turn.created",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventTurnCompletedWebhookEvent{}),
			DiscriminatorValue: "turn.completed",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventTurnFailedWebhookEvent{}),
			DiscriminatorValue: "turn.failed",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventMessageCreatedWebhookEvent{}),
			DiscriminatorValue: "message.created",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventApprovalRequestedWebhookEvent{}),
			DiscriminatorValue: "approval.requested",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventApprovalResolvedWebhookEvent{}),
			DiscriminatorValue: "approval.resolved",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventApprovalGrantedWebhookEvent{}),
			DiscriminatorValue: "approval.granted",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventScheduleCreatedWebhookEvent{}),
			DiscriminatorValue: "schedule.created",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventScheduleDeletedWebhookEvent{}),
			DiscriminatorValue: "schedule.deleted",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventScheduleTriggeredWebhookEvent{}),
			DiscriminatorValue: "schedule.triggered",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventConnectionAttachedWebhookEvent{}),
			DiscriminatorValue: "connection.attached",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventConnectionDetachedWebhookEvent{}),
			DiscriminatorValue: "connection.detached",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuntimeWebhookEventWebhookTestWebhookEvent{}),
			DiscriminatorValue: "webhook.test",
		},
	)
}

type RuntimeWebhookEventAgentCreatedWebhookEvent struct {
	ID        string                                           `json:"id" api:"required"`
	AgentID   string                                           `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventAgentCreatedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventAgentCreatedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                           `json:"fleetId" api:"required"`
	OrgID     string                                           `json:"orgId" api:"required"`
	ThreadID  string                                           `json:"threadId" api:"required,nullable"`
	Timestamp string                                           `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventAgentCreatedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventAgentCreatedWebhookEventJSON contains the JSON metadata for
// the struct [RuntimeWebhookEventAgentCreatedWebhookEvent]
type runtimeWebhookEventAgentCreatedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventAgentCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventAgentCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventAgentCreatedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventAgentCreatedWebhookEventData struct {
	Agent RuntimeWebhookEventAgentCreatedWebhookEventDataAgent `json:"agent" api:"required"`
	JSON  runtimeWebhookEventAgentCreatedWebhookEventDataJSON  `json:"-"`
}

// runtimeWebhookEventAgentCreatedWebhookEventDataJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventAgentCreatedWebhookEventData]
type runtimeWebhookEventAgentCreatedWebhookEventDataJSON struct {
	Agent       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeWebhookEventAgentCreatedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventAgentCreatedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventAgentCreatedWebhookEventDataAgent struct {
	ID            string        `json:"id" api:"required"`
	Configuration Configuration `json:"configuration" api:"required"`
	CreatedAt     string        `json:"createdAt" api:"required"`
	FleetID       string        `json:"fleetId" api:"required"`
	// Arbitrary string metadata stored on an agent. Runtime enforces maximum key and
	// value sizes.
	Metadata  Metadata                                                 `json:"metadata" api:"required"`
	OrgID     string                                                   `json:"orgId" api:"required"`
	UpdatedAt string                                                   `json:"updatedAt" api:"required"`
	UserID    string                                                   `json:"userId" api:"required"`
	Effective EffectiveConfiguration                                   `json:"effective"`
	JSON      runtimeWebhookEventAgentCreatedWebhookEventDataAgentJSON `json:"-"`
}

// runtimeWebhookEventAgentCreatedWebhookEventDataAgentJSON contains the JSON
// metadata for the struct [RuntimeWebhookEventAgentCreatedWebhookEventDataAgent]
type runtimeWebhookEventAgentCreatedWebhookEventDataAgentJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	CreatedAt     apijson.Field
	FleetID       apijson.Field
	Metadata      apijson.Field
	OrgID         apijson.Field
	UpdatedAt     apijson.Field
	UserID        apijson.Field
	Effective     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RuntimeWebhookEventAgentCreatedWebhookEventDataAgent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventAgentCreatedWebhookEventDataAgentJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventAgentCreatedWebhookEventEvent string

const (
	RuntimeWebhookEventAgentCreatedWebhookEventEventAgentCreated RuntimeWebhookEventAgentCreatedWebhookEventEvent = "agent.created"
)

func (r RuntimeWebhookEventAgentCreatedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventAgentCreatedWebhookEventEventAgentCreated:
		return true
	}
	return false
}

type RuntimeWebhookEventAgentUpdatedWebhookEvent struct {
	ID        string                                           `json:"id" api:"required"`
	AgentID   string                                           `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventAgentUpdatedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventAgentUpdatedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                           `json:"fleetId" api:"required"`
	OrgID     string                                           `json:"orgId" api:"required"`
	ThreadID  string                                           `json:"threadId" api:"required,nullable"`
	Timestamp string                                           `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventAgentUpdatedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventAgentUpdatedWebhookEventJSON contains the JSON metadata for
// the struct [RuntimeWebhookEventAgentUpdatedWebhookEvent]
type runtimeWebhookEventAgentUpdatedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventAgentUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventAgentUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventAgentUpdatedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventAgentUpdatedWebhookEventData struct {
	Agent   RuntimeWebhookEventAgentUpdatedWebhookEventDataAgent   `json:"agent" api:"required"`
	Trigger RuntimeWebhookEventAgentUpdatedWebhookEventDataTrigger `json:"trigger" api:"required"`
	JSON    runtimeWebhookEventAgentUpdatedWebhookEventDataJSON    `json:"-"`
}

// runtimeWebhookEventAgentUpdatedWebhookEventDataJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventAgentUpdatedWebhookEventData]
type runtimeWebhookEventAgentUpdatedWebhookEventDataJSON struct {
	Agent       apijson.Field
	Trigger     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeWebhookEventAgentUpdatedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventAgentUpdatedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventAgentUpdatedWebhookEventDataAgent struct {
	ID            string        `json:"id" api:"required"`
	Configuration Configuration `json:"configuration" api:"required"`
	CreatedAt     string        `json:"createdAt" api:"required"`
	FleetID       string        `json:"fleetId" api:"required"`
	// Arbitrary string metadata stored on an agent. Runtime enforces maximum key and
	// value sizes.
	Metadata  Metadata                                                 `json:"metadata" api:"required"`
	OrgID     string                                                   `json:"orgId" api:"required"`
	UpdatedAt string                                                   `json:"updatedAt" api:"required"`
	UserID    string                                                   `json:"userId" api:"required"`
	Effective EffectiveConfiguration                                   `json:"effective"`
	JSON      runtimeWebhookEventAgentUpdatedWebhookEventDataAgentJSON `json:"-"`
}

// runtimeWebhookEventAgentUpdatedWebhookEventDataAgentJSON contains the JSON
// metadata for the struct [RuntimeWebhookEventAgentUpdatedWebhookEventDataAgent]
type runtimeWebhookEventAgentUpdatedWebhookEventDataAgentJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	CreatedAt     apijson.Field
	FleetID       apijson.Field
	Metadata      apijson.Field
	OrgID         apijson.Field
	UpdatedAt     apijson.Field
	UserID        apijson.Field
	Effective     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RuntimeWebhookEventAgentUpdatedWebhookEventDataAgent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventAgentUpdatedWebhookEventDataAgentJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventAgentUpdatedWebhookEventDataTrigger string

const (
	RuntimeWebhookEventAgentUpdatedWebhookEventDataTriggerConfiguration RuntimeWebhookEventAgentUpdatedWebhookEventDataTrigger = "configuration"
	RuntimeWebhookEventAgentUpdatedWebhookEventDataTriggerMetadata      RuntimeWebhookEventAgentUpdatedWebhookEventDataTrigger = "metadata"
)

func (r RuntimeWebhookEventAgentUpdatedWebhookEventDataTrigger) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventAgentUpdatedWebhookEventDataTriggerConfiguration, RuntimeWebhookEventAgentUpdatedWebhookEventDataTriggerMetadata:
		return true
	}
	return false
}

type RuntimeWebhookEventAgentUpdatedWebhookEventEvent string

const (
	RuntimeWebhookEventAgentUpdatedWebhookEventEventAgentUpdated RuntimeWebhookEventAgentUpdatedWebhookEventEvent = "agent.updated"
)

func (r RuntimeWebhookEventAgentUpdatedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventAgentUpdatedWebhookEventEventAgentUpdated:
		return true
	}
	return false
}

type RuntimeWebhookEventAgentDeletedWebhookEvent struct {
	ID        string                                           `json:"id" api:"required"`
	AgentID   string                                           `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventAgentDeletedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventAgentDeletedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                           `json:"fleetId" api:"required"`
	OrgID     string                                           `json:"orgId" api:"required"`
	ThreadID  string                                           `json:"threadId" api:"required,nullable"`
	Timestamp string                                           `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventAgentDeletedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventAgentDeletedWebhookEventJSON contains the JSON metadata for
// the struct [RuntimeWebhookEventAgentDeletedWebhookEvent]
type runtimeWebhookEventAgentDeletedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventAgentDeletedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventAgentDeletedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventAgentDeletedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventAgentDeletedWebhookEventData struct {
	AgentID string                                              `json:"agentId" api:"required"`
	JSON    runtimeWebhookEventAgentDeletedWebhookEventDataJSON `json:"-"`
}

// runtimeWebhookEventAgentDeletedWebhookEventDataJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventAgentDeletedWebhookEventData]
type runtimeWebhookEventAgentDeletedWebhookEventDataJSON struct {
	AgentID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeWebhookEventAgentDeletedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventAgentDeletedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventAgentDeletedWebhookEventEvent string

const (
	RuntimeWebhookEventAgentDeletedWebhookEventEventAgentDeleted RuntimeWebhookEventAgentDeletedWebhookEventEvent = "agent.deleted"
)

func (r RuntimeWebhookEventAgentDeletedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventAgentDeletedWebhookEventEventAgentDeleted:
		return true
	}
	return false
}

type RuntimeWebhookEventThreadCreatedWebhookEvent struct {
	ID        string                                            `json:"id" api:"required"`
	AgentID   string                                            `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventThreadCreatedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventThreadCreatedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                            `json:"fleetId" api:"required"`
	OrgID     string                                            `json:"orgId" api:"required"`
	ThreadID  string                                            `json:"threadId" api:"required,nullable"`
	Timestamp string                                            `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventThreadCreatedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventThreadCreatedWebhookEventJSON contains the JSON metadata for
// the struct [RuntimeWebhookEventThreadCreatedWebhookEvent]
type runtimeWebhookEventThreadCreatedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventThreadCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventThreadCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventThreadCreatedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventThreadCreatedWebhookEventData struct {
	Message string                                               `json:"message" api:"required"`
	Model   string                                               `json:"model" api:"required"`
	JSON    runtimeWebhookEventThreadCreatedWebhookEventDataJSON `json:"-"`
}

// runtimeWebhookEventThreadCreatedWebhookEventDataJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventThreadCreatedWebhookEventData]
type runtimeWebhookEventThreadCreatedWebhookEventDataJSON struct {
	Message     apijson.Field
	Model       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeWebhookEventThreadCreatedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventThreadCreatedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventThreadCreatedWebhookEventEvent string

const (
	RuntimeWebhookEventThreadCreatedWebhookEventEventThreadCreated RuntimeWebhookEventThreadCreatedWebhookEventEvent = "thread.created"
)

func (r RuntimeWebhookEventThreadCreatedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventThreadCreatedWebhookEventEventThreadCreated:
		return true
	}
	return false
}

type RuntimeWebhookEventThreadStatusChangedWebhookEvent struct {
	ID        string                                                  `json:"id" api:"required"`
	AgentID   string                                                  `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventThreadStatusChangedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventThreadStatusChangedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                                  `json:"fleetId" api:"required"`
	OrgID     string                                                  `json:"orgId" api:"required"`
	ThreadID  string                                                  `json:"threadId" api:"required,nullable"`
	Timestamp string                                                  `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventThreadStatusChangedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventThreadStatusChangedWebhookEventJSON contains the JSON
// metadata for the struct [RuntimeWebhookEventThreadStatusChangedWebhookEvent]
type runtimeWebhookEventThreadStatusChangedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventThreadStatusChangedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventThreadStatusChangedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventThreadStatusChangedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventThreadStatusChangedWebhookEventData struct {
	Error string `json:"error" api:"required,nullable"`
	// `idle` threads can accept a new turn or be closed. `running` threads have an
	// active turn. `awaiting` threads are paused on external input such as approvals.
	// `closed` threads are terminal.
	PreviousStatus Status `json:"previousStatus" api:"required"`
	Result         string `json:"result" api:"required,nullable"`
	// `idle` threads can accept a new turn or be closed. `running` threads have an
	// active turn. `awaiting` threads are paused on external input such as approvals.
	// `closed` threads are terminal.
	Status Status                                                     `json:"status" api:"required"`
	JSON   runtimeWebhookEventThreadStatusChangedWebhookEventDataJSON `json:"-"`
}

// runtimeWebhookEventThreadStatusChangedWebhookEventDataJSON contains the JSON
// metadata for the struct [RuntimeWebhookEventThreadStatusChangedWebhookEventData]
type runtimeWebhookEventThreadStatusChangedWebhookEventDataJSON struct {
	Error          apijson.Field
	PreviousStatus apijson.Field
	Result         apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RuntimeWebhookEventThreadStatusChangedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventThreadStatusChangedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventThreadStatusChangedWebhookEventEvent string

const (
	RuntimeWebhookEventThreadStatusChangedWebhookEventEventThreadStatusChanged RuntimeWebhookEventThreadStatusChangedWebhookEventEvent = "thread.status.changed"
)

func (r RuntimeWebhookEventThreadStatusChangedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventThreadStatusChangedWebhookEventEventThreadStatusChanged:
		return true
	}
	return false
}

type RuntimeWebhookEventThreadCompletedWebhookEvent struct {
	ID        string                                              `json:"id" api:"required"`
	AgentID   string                                              `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventThreadCompletedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventThreadCompletedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                              `json:"fleetId" api:"required"`
	OrgID     string                                              `json:"orgId" api:"required"`
	ThreadID  string                                              `json:"threadId" api:"required,nullable"`
	Timestamp string                                              `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventThreadCompletedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventThreadCompletedWebhookEventJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventThreadCompletedWebhookEvent]
type runtimeWebhookEventThreadCompletedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventThreadCompletedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventThreadCompletedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventThreadCompletedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventThreadCompletedWebhookEventData struct {
	// `idle` threads can accept a new turn or be closed. `running` threads have an
	// active turn. `awaiting` threads are paused on external input such as approvals.
	// `closed` threads are terminal.
	PreviousStatus Status `json:"previousStatus" api:"required"`
	Result         string `json:"result" api:"required"`
	// `idle` threads can accept a new turn or be closed. `running` threads have an
	// active turn. `awaiting` threads are paused on external input such as approvals.
	// `closed` threads are terminal.
	Status Status                                                 `json:"status" api:"required"`
	JSON   runtimeWebhookEventThreadCompletedWebhookEventDataJSON `json:"-"`
}

// runtimeWebhookEventThreadCompletedWebhookEventDataJSON contains the JSON
// metadata for the struct [RuntimeWebhookEventThreadCompletedWebhookEventData]
type runtimeWebhookEventThreadCompletedWebhookEventDataJSON struct {
	PreviousStatus apijson.Field
	Result         apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RuntimeWebhookEventThreadCompletedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventThreadCompletedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventThreadCompletedWebhookEventEvent string

const (
	RuntimeWebhookEventThreadCompletedWebhookEventEventThreadCompleted RuntimeWebhookEventThreadCompletedWebhookEventEvent = "thread.completed"
)

func (r RuntimeWebhookEventThreadCompletedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventThreadCompletedWebhookEventEventThreadCompleted:
		return true
	}
	return false
}

type RuntimeWebhookEventThreadFailedWebhookEvent struct {
	ID        string                                           `json:"id" api:"required"`
	AgentID   string                                           `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventThreadFailedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventThreadFailedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                           `json:"fleetId" api:"required"`
	OrgID     string                                           `json:"orgId" api:"required"`
	ThreadID  string                                           `json:"threadId" api:"required,nullable"`
	Timestamp string                                           `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventThreadFailedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventThreadFailedWebhookEventJSON contains the JSON metadata for
// the struct [RuntimeWebhookEventThreadFailedWebhookEvent]
type runtimeWebhookEventThreadFailedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventThreadFailedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventThreadFailedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventThreadFailedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventThreadFailedWebhookEventData struct {
	Error string `json:"error" api:"required"`
	// `idle` threads can accept a new turn or be closed. `running` threads have an
	// active turn. `awaiting` threads are paused on external input such as approvals.
	// `closed` threads are terminal.
	PreviousStatus Status `json:"previousStatus" api:"required"`
	// `idle` threads can accept a new turn or be closed. `running` threads have an
	// active turn. `awaiting` threads are paused on external input such as approvals.
	// `closed` threads are terminal.
	Status Status                                              `json:"status" api:"required"`
	JSON   runtimeWebhookEventThreadFailedWebhookEventDataJSON `json:"-"`
}

// runtimeWebhookEventThreadFailedWebhookEventDataJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventThreadFailedWebhookEventData]
type runtimeWebhookEventThreadFailedWebhookEventDataJSON struct {
	Error          apijson.Field
	PreviousStatus apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RuntimeWebhookEventThreadFailedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventThreadFailedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventThreadFailedWebhookEventEvent string

const (
	RuntimeWebhookEventThreadFailedWebhookEventEventThreadFailed RuntimeWebhookEventThreadFailedWebhookEventEvent = "thread.failed"
)

func (r RuntimeWebhookEventThreadFailedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventThreadFailedWebhookEventEventThreadFailed:
		return true
	}
	return false
}

type RuntimeWebhookEventTurnCreatedWebhookEvent struct {
	ID        string                                          `json:"id" api:"required"`
	AgentID   string                                          `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventTurnCreatedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventTurnCreatedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                          `json:"fleetId" api:"required"`
	OrgID     string                                          `json:"orgId" api:"required"`
	ThreadID  string                                          `json:"threadId" api:"required,nullable"`
	Timestamp string                                          `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventTurnCreatedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventTurnCreatedWebhookEventJSON contains the JSON metadata for
// the struct [RuntimeWebhookEventTurnCreatedWebhookEvent]
type runtimeWebhookEventTurnCreatedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventTurnCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventTurnCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventTurnCreatedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventTurnCreatedWebhookEventData struct {
	Message   string                                             `json:"message" api:"required"`
	Seq       float64                                            `json:"seq" api:"required"`
	StartedAt string                                             `json:"startedAt" api:"required"`
	TurnID    string                                             `json:"turnId" api:"required"`
	JSON      runtimeWebhookEventTurnCreatedWebhookEventDataJSON `json:"-"`
}

// runtimeWebhookEventTurnCreatedWebhookEventDataJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventTurnCreatedWebhookEventData]
type runtimeWebhookEventTurnCreatedWebhookEventDataJSON struct {
	Message     apijson.Field
	Seq         apijson.Field
	StartedAt   apijson.Field
	TurnID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeWebhookEventTurnCreatedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventTurnCreatedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventTurnCreatedWebhookEventEvent string

const (
	RuntimeWebhookEventTurnCreatedWebhookEventEventTurnCreated RuntimeWebhookEventTurnCreatedWebhookEventEvent = "turn.created"
)

func (r RuntimeWebhookEventTurnCreatedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventTurnCreatedWebhookEventEventTurnCreated:
		return true
	}
	return false
}

type RuntimeWebhookEventTurnCompletedWebhookEvent struct {
	ID        string                                            `json:"id" api:"required"`
	AgentID   string                                            `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventTurnCompletedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventTurnCompletedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                            `json:"fleetId" api:"required"`
	OrgID     string                                            `json:"orgId" api:"required"`
	ThreadID  string                                            `json:"threadId" api:"required,nullable"`
	Timestamp string                                            `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventTurnCompletedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventTurnCompletedWebhookEventJSON contains the JSON metadata for
// the struct [RuntimeWebhookEventTurnCompletedWebhookEvent]
type runtimeWebhookEventTurnCompletedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventTurnCompletedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventTurnCompletedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventTurnCompletedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventTurnCompletedWebhookEventData struct {
	CompletedAt string                                                 `json:"completedAt" api:"required"`
	Error       string                                                 `json:"error" api:"required,nullable"`
	Result      string                                                 `json:"result" api:"required,nullable"`
	Seq         float64                                                `json:"seq" api:"required"`
	Status      RuntimeWebhookEventTurnCompletedWebhookEventDataStatus `json:"status" api:"required"`
	TokenUsage  TokenUsage                                             `json:"tokenUsage" api:"required,nullable"`
	TurnID      string                                                 `json:"turnId" api:"required"`
	JSON        runtimeWebhookEventTurnCompletedWebhookEventDataJSON   `json:"-"`
}

// runtimeWebhookEventTurnCompletedWebhookEventDataJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventTurnCompletedWebhookEventData]
type runtimeWebhookEventTurnCompletedWebhookEventDataJSON struct {
	CompletedAt apijson.Field
	Error       apijson.Field
	Result      apijson.Field
	Seq         apijson.Field
	Status      apijson.Field
	TokenUsage  apijson.Field
	TurnID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeWebhookEventTurnCompletedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventTurnCompletedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventTurnCompletedWebhookEventDataStatus string

const (
	RuntimeWebhookEventTurnCompletedWebhookEventDataStatusCompleted RuntimeWebhookEventTurnCompletedWebhookEventDataStatus = "completed"
)

func (r RuntimeWebhookEventTurnCompletedWebhookEventDataStatus) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventTurnCompletedWebhookEventDataStatusCompleted:
		return true
	}
	return false
}

type RuntimeWebhookEventTurnCompletedWebhookEventEvent string

const (
	RuntimeWebhookEventTurnCompletedWebhookEventEventTurnCompleted RuntimeWebhookEventTurnCompletedWebhookEventEvent = "turn.completed"
)

func (r RuntimeWebhookEventTurnCompletedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventTurnCompletedWebhookEventEventTurnCompleted:
		return true
	}
	return false
}

type RuntimeWebhookEventTurnFailedWebhookEvent struct {
	ID        string                                         `json:"id" api:"required"`
	AgentID   string                                         `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventTurnFailedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventTurnFailedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                         `json:"fleetId" api:"required"`
	OrgID     string                                         `json:"orgId" api:"required"`
	ThreadID  string                                         `json:"threadId" api:"required,nullable"`
	Timestamp string                                         `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventTurnFailedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventTurnFailedWebhookEventJSON contains the JSON metadata for the
// struct [RuntimeWebhookEventTurnFailedWebhookEvent]
type runtimeWebhookEventTurnFailedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventTurnFailedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventTurnFailedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventTurnFailedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventTurnFailedWebhookEventData struct {
	CompletedAt string                                              `json:"completedAt" api:"required"`
	Error       string                                              `json:"error" api:"required,nullable"`
	Result      string                                              `json:"result" api:"required,nullable"`
	Seq         float64                                             `json:"seq" api:"required"`
	Status      RuntimeWebhookEventTurnFailedWebhookEventDataStatus `json:"status" api:"required"`
	TokenUsage  TokenUsage                                          `json:"tokenUsage" api:"required,nullable"`
	TurnID      string                                              `json:"turnId" api:"required"`
	JSON        runtimeWebhookEventTurnFailedWebhookEventDataJSON   `json:"-"`
}

// runtimeWebhookEventTurnFailedWebhookEventDataJSON contains the JSON metadata for
// the struct [RuntimeWebhookEventTurnFailedWebhookEventData]
type runtimeWebhookEventTurnFailedWebhookEventDataJSON struct {
	CompletedAt apijson.Field
	Error       apijson.Field
	Result      apijson.Field
	Seq         apijson.Field
	Status      apijson.Field
	TokenUsage  apijson.Field
	TurnID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeWebhookEventTurnFailedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventTurnFailedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventTurnFailedWebhookEventDataStatus string

const (
	RuntimeWebhookEventTurnFailedWebhookEventDataStatusFailed RuntimeWebhookEventTurnFailedWebhookEventDataStatus = "failed"
)

func (r RuntimeWebhookEventTurnFailedWebhookEventDataStatus) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventTurnFailedWebhookEventDataStatusFailed:
		return true
	}
	return false
}

type RuntimeWebhookEventTurnFailedWebhookEventEvent string

const (
	RuntimeWebhookEventTurnFailedWebhookEventEventTurnFailed RuntimeWebhookEventTurnFailedWebhookEventEvent = "turn.failed"
)

func (r RuntimeWebhookEventTurnFailedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventTurnFailedWebhookEventEventTurnFailed:
		return true
	}
	return false
}

type RuntimeWebhookEventMessageCreatedWebhookEvent struct {
	ID        string                                             `json:"id" api:"required"`
	AgentID   string                                             `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventMessageCreatedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventMessageCreatedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                             `json:"fleetId" api:"required"`
	OrgID     string                                             `json:"orgId" api:"required"`
	ThreadID  string                                             `json:"threadId" api:"required,nullable"`
	Timestamp string                                             `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventMessageCreatedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventMessageCreatedWebhookEventJSON contains the JSON metadata for
// the struct [RuntimeWebhookEventMessageCreatedWebhookEvent]
type runtimeWebhookEventMessageCreatedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventMessageCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventMessageCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventMessageCreatedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventMessageCreatedWebhookEventData struct {
	CreatedAt string                                                `json:"createdAt" api:"required"`
	Role      RuntimeWebhookEventMessageCreatedWebhookEventDataRole `json:"role" api:"required"`
	Seq       float64                                               `json:"seq" api:"required"`
	TurnID    string                                                `json:"turnId" api:"required,nullable"`
	JSON      runtimeWebhookEventMessageCreatedWebhookEventDataJSON `json:"-"`
}

// runtimeWebhookEventMessageCreatedWebhookEventDataJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventMessageCreatedWebhookEventData]
type runtimeWebhookEventMessageCreatedWebhookEventDataJSON struct {
	CreatedAt   apijson.Field
	Role        apijson.Field
	Seq         apijson.Field
	TurnID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeWebhookEventMessageCreatedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventMessageCreatedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventMessageCreatedWebhookEventDataRole string

const (
	RuntimeWebhookEventMessageCreatedWebhookEventDataRoleUser      RuntimeWebhookEventMessageCreatedWebhookEventDataRole = "user"
	RuntimeWebhookEventMessageCreatedWebhookEventDataRoleAssistant RuntimeWebhookEventMessageCreatedWebhookEventDataRole = "assistant"
)

func (r RuntimeWebhookEventMessageCreatedWebhookEventDataRole) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventMessageCreatedWebhookEventDataRoleUser, RuntimeWebhookEventMessageCreatedWebhookEventDataRoleAssistant:
		return true
	}
	return false
}

type RuntimeWebhookEventMessageCreatedWebhookEventEvent string

const (
	RuntimeWebhookEventMessageCreatedWebhookEventEventMessageCreated RuntimeWebhookEventMessageCreatedWebhookEventEvent = "message.created"
)

func (r RuntimeWebhookEventMessageCreatedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventMessageCreatedWebhookEventEventMessageCreated:
		return true
	}
	return false
}

type RuntimeWebhookEventApprovalRequestedWebhookEvent struct {
	ID        string                                                `json:"id" api:"required"`
	AgentID   string                                                `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventApprovalRequestedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventApprovalRequestedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                                `json:"fleetId" api:"required"`
	OrgID     string                                                `json:"orgId" api:"required"`
	ThreadID  string                                                `json:"threadId" api:"required,nullable"`
	Timestamp string                                                `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventApprovalRequestedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventApprovalRequestedWebhookEventJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventApprovalRequestedWebhookEvent]
type runtimeWebhookEventApprovalRequestedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventApprovalRequestedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventApprovalRequestedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventApprovalRequestedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventApprovalRequestedWebhookEventData struct {
	Approval RuntimeWebhookEventApprovalRequestedWebhookEventDataApproval `json:"approval" api:"required"`
	JSON     runtimeWebhookEventApprovalRequestedWebhookEventDataJSON     `json:"-"`
}

// runtimeWebhookEventApprovalRequestedWebhookEventDataJSON contains the JSON
// metadata for the struct [RuntimeWebhookEventApprovalRequestedWebhookEventData]
type runtimeWebhookEventApprovalRequestedWebhookEventDataJSON struct {
	Approval    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeWebhookEventApprovalRequestedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventApprovalRequestedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventApprovalRequestedWebhookEventDataApproval struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// JSON value payload. Generated SDKs may expose this as unknown or Any.
	Input             interface{}                                                        `json:"input" api:"required"`
	ResolvedAt        string                                                             `json:"resolvedAt" api:"required,nullable"`
	RuntimeToolName   shared.ToolName                                                    `json:"runtimeToolName" api:"required"`
	Status            RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatus `json:"status" api:"required"`
	ThreadID          string                                                             `json:"threadId" api:"required,nullable"`
	TimeoutAt         string                                                             `json:"timeoutAt" api:"required,nullable"`
	TimeoutMs         float64                                                            `json:"timeoutMs" api:"required,nullable"`
	ToolIndex         float64                                                            `json:"toolIndex" api:"required"`
	ToolName          string                                                             `json:"toolName" api:"required"`
	ToolUseID         string                                                             `json:"toolUseId" api:"required"`
	TurnID            string                                                             `json:"turnId" api:"required"`
	ToolSourceID      string                                                             `json:"toolSourceId"`
	ToolSourceVersion float64                                                            `json:"toolSourceVersion"`
	JSON              runtimeWebhookEventApprovalRequestedWebhookEventDataApprovalJSON   `json:"-"`
}

// runtimeWebhookEventApprovalRequestedWebhookEventDataApprovalJSON contains the
// JSON metadata for the struct
// [RuntimeWebhookEventApprovalRequestedWebhookEventDataApproval]
type runtimeWebhookEventApprovalRequestedWebhookEventDataApprovalJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	Input             apijson.Field
	ResolvedAt        apijson.Field
	RuntimeToolName   apijson.Field
	Status            apijson.Field
	ThreadID          apijson.Field
	TimeoutAt         apijson.Field
	TimeoutMs         apijson.Field
	ToolIndex         apijson.Field
	ToolName          apijson.Field
	ToolUseID         apijson.Field
	TurnID            apijson.Field
	ToolSourceID      apijson.Field
	ToolSourceVersion apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RuntimeWebhookEventApprovalRequestedWebhookEventDataApproval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventApprovalRequestedWebhookEventDataApprovalJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatus string

const (
	RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatusPending   RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatus = "pending"
	RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatusApproved  RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatus = "approved"
	RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatusDenied    RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatus = "denied"
	RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatusCancelled RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatus = "cancelled"
	RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatusTimedOut  RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatus = "timed_out"
)

func (r RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatus) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatusPending, RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatusApproved, RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatusDenied, RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatusCancelled, RuntimeWebhookEventApprovalRequestedWebhookEventDataApprovalStatusTimedOut:
		return true
	}
	return false
}

type RuntimeWebhookEventApprovalRequestedWebhookEventEvent string

const (
	RuntimeWebhookEventApprovalRequestedWebhookEventEventApprovalRequested RuntimeWebhookEventApprovalRequestedWebhookEventEvent = "approval.requested"
)

func (r RuntimeWebhookEventApprovalRequestedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventApprovalRequestedWebhookEventEventApprovalRequested:
		return true
	}
	return false
}

type RuntimeWebhookEventApprovalResolvedWebhookEvent struct {
	ID        string                                               `json:"id" api:"required"`
	AgentID   string                                               `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventApprovalResolvedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventApprovalResolvedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                               `json:"fleetId" api:"required"`
	OrgID     string                                               `json:"orgId" api:"required"`
	ThreadID  string                                               `json:"threadId" api:"required,nullable"`
	Timestamp string                                               `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventApprovalResolvedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventApprovalResolvedWebhookEventJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventApprovalResolvedWebhookEvent]
type runtimeWebhookEventApprovalResolvedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventApprovalResolvedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventApprovalResolvedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventApprovalResolvedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventApprovalResolvedWebhookEventData struct {
	ApprovalID string                                                      `json:"approvalId" api:"required"`
	Decision   RuntimeWebhookEventApprovalResolvedWebhookEventDataDecision `json:"decision" api:"required"`
	JSON       runtimeWebhookEventApprovalResolvedWebhookEventDataJSON     `json:"-"`
}

// runtimeWebhookEventApprovalResolvedWebhookEventDataJSON contains the JSON
// metadata for the struct [RuntimeWebhookEventApprovalResolvedWebhookEventData]
type runtimeWebhookEventApprovalResolvedWebhookEventDataJSON struct {
	ApprovalID  apijson.Field
	Decision    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeWebhookEventApprovalResolvedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventApprovalResolvedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventApprovalResolvedWebhookEventDataDecision string

const (
	RuntimeWebhookEventApprovalResolvedWebhookEventDataDecisionApproved  RuntimeWebhookEventApprovalResolvedWebhookEventDataDecision = "approved"
	RuntimeWebhookEventApprovalResolvedWebhookEventDataDecisionDenied    RuntimeWebhookEventApprovalResolvedWebhookEventDataDecision = "denied"
	RuntimeWebhookEventApprovalResolvedWebhookEventDataDecisionTimedOut  RuntimeWebhookEventApprovalResolvedWebhookEventDataDecision = "timed_out"
	RuntimeWebhookEventApprovalResolvedWebhookEventDataDecisionCancelled RuntimeWebhookEventApprovalResolvedWebhookEventDataDecision = "cancelled"
)

func (r RuntimeWebhookEventApprovalResolvedWebhookEventDataDecision) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventApprovalResolvedWebhookEventDataDecisionApproved, RuntimeWebhookEventApprovalResolvedWebhookEventDataDecisionDenied, RuntimeWebhookEventApprovalResolvedWebhookEventDataDecisionTimedOut, RuntimeWebhookEventApprovalResolvedWebhookEventDataDecisionCancelled:
		return true
	}
	return false
}

type RuntimeWebhookEventApprovalResolvedWebhookEventEvent string

const (
	RuntimeWebhookEventApprovalResolvedWebhookEventEventApprovalResolved RuntimeWebhookEventApprovalResolvedWebhookEventEvent = "approval.resolved"
)

func (r RuntimeWebhookEventApprovalResolvedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventApprovalResolvedWebhookEventEventApprovalResolved:
		return true
	}
	return false
}

type RuntimeWebhookEventApprovalGrantedWebhookEvent struct {
	ID        string                                              `json:"id" api:"required"`
	AgentID   string                                              `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventApprovalGrantedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventApprovalGrantedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                              `json:"fleetId" api:"required"`
	OrgID     string                                              `json:"orgId" api:"required"`
	ThreadID  string                                              `json:"threadId" api:"required,nullable"`
	Timestamp string                                              `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventApprovalGrantedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventApprovalGrantedWebhookEventJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventApprovalGrantedWebhookEvent]
type runtimeWebhookEventApprovalGrantedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventApprovalGrantedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventApprovalGrantedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventApprovalGrantedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventApprovalGrantedWebhookEventData struct {
	GrantID  string                                                  `json:"grantId" api:"required"`
	Scope    RuntimeWebhookEventApprovalGrantedWebhookEventDataScope `json:"scope" api:"required"`
	ToolName string                                                  `json:"toolName" api:"required"`
	ThreadID string                                                  `json:"threadId"`
	JSON     runtimeWebhookEventApprovalGrantedWebhookEventDataJSON  `json:"-"`
}

// runtimeWebhookEventApprovalGrantedWebhookEventDataJSON contains the JSON
// metadata for the struct [RuntimeWebhookEventApprovalGrantedWebhookEventData]
type runtimeWebhookEventApprovalGrantedWebhookEventDataJSON struct {
	GrantID     apijson.Field
	Scope       apijson.Field
	ToolName    apijson.Field
	ThreadID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeWebhookEventApprovalGrantedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventApprovalGrantedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventApprovalGrantedWebhookEventDataScope string

const (
	RuntimeWebhookEventApprovalGrantedWebhookEventDataScopeThread RuntimeWebhookEventApprovalGrantedWebhookEventDataScope = "thread"
	RuntimeWebhookEventApprovalGrantedWebhookEventDataScopeAgent  RuntimeWebhookEventApprovalGrantedWebhookEventDataScope = "agent"
)

func (r RuntimeWebhookEventApprovalGrantedWebhookEventDataScope) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventApprovalGrantedWebhookEventDataScopeThread, RuntimeWebhookEventApprovalGrantedWebhookEventDataScopeAgent:
		return true
	}
	return false
}

type RuntimeWebhookEventApprovalGrantedWebhookEventEvent string

const (
	RuntimeWebhookEventApprovalGrantedWebhookEventEventApprovalGranted RuntimeWebhookEventApprovalGrantedWebhookEventEvent = "approval.granted"
)

func (r RuntimeWebhookEventApprovalGrantedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventApprovalGrantedWebhookEventEventApprovalGranted:
		return true
	}
	return false
}

type RuntimeWebhookEventScheduleCreatedWebhookEvent struct {
	ID        string                                              `json:"id" api:"required"`
	AgentID   string                                              `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventScheduleCreatedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventScheduleCreatedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                              `json:"fleetId" api:"required"`
	OrgID     string                                              `json:"orgId" api:"required"`
	ThreadID  string                                              `json:"threadId" api:"required,nullable"`
	Timestamp string                                              `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventScheduleCreatedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventScheduleCreatedWebhookEventJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventScheduleCreatedWebhookEvent]
type runtimeWebhookEventScheduleCreatedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventScheduleCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventScheduleCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventScheduleCreatedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventScheduleCreatedWebhookEventData struct {
	Name       string                                                 `json:"name" api:"required"`
	ScheduleID string                                                 `json:"scheduleId" api:"required"`
	Cron       string                                                 `json:"cron"`
	RunAt      string                                                 `json:"runAt"`
	JSON       runtimeWebhookEventScheduleCreatedWebhookEventDataJSON `json:"-"`
}

// runtimeWebhookEventScheduleCreatedWebhookEventDataJSON contains the JSON
// metadata for the struct [RuntimeWebhookEventScheduleCreatedWebhookEventData]
type runtimeWebhookEventScheduleCreatedWebhookEventDataJSON struct {
	Name        apijson.Field
	ScheduleID  apijson.Field
	Cron        apijson.Field
	RunAt       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeWebhookEventScheduleCreatedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventScheduleCreatedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventScheduleCreatedWebhookEventEvent string

const (
	RuntimeWebhookEventScheduleCreatedWebhookEventEventScheduleCreated RuntimeWebhookEventScheduleCreatedWebhookEventEvent = "schedule.created"
)

func (r RuntimeWebhookEventScheduleCreatedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventScheduleCreatedWebhookEventEventScheduleCreated:
		return true
	}
	return false
}

type RuntimeWebhookEventScheduleDeletedWebhookEvent struct {
	ID        string                                              `json:"id" api:"required"`
	AgentID   string                                              `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventScheduleDeletedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventScheduleDeletedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                              `json:"fleetId" api:"required"`
	OrgID     string                                              `json:"orgId" api:"required"`
	ThreadID  string                                              `json:"threadId" api:"required,nullable"`
	Timestamp string                                              `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventScheduleDeletedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventScheduleDeletedWebhookEventJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventScheduleDeletedWebhookEvent]
type runtimeWebhookEventScheduleDeletedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventScheduleDeletedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventScheduleDeletedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventScheduleDeletedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventScheduleDeletedWebhookEventData struct {
	ScheduleID string                                                 `json:"scheduleId" api:"required"`
	JSON       runtimeWebhookEventScheduleDeletedWebhookEventDataJSON `json:"-"`
}

// runtimeWebhookEventScheduleDeletedWebhookEventDataJSON contains the JSON
// metadata for the struct [RuntimeWebhookEventScheduleDeletedWebhookEventData]
type runtimeWebhookEventScheduleDeletedWebhookEventDataJSON struct {
	ScheduleID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeWebhookEventScheduleDeletedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventScheduleDeletedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventScheduleDeletedWebhookEventEvent string

const (
	RuntimeWebhookEventScheduleDeletedWebhookEventEventScheduleDeleted RuntimeWebhookEventScheduleDeletedWebhookEventEvent = "schedule.deleted"
)

func (r RuntimeWebhookEventScheduleDeletedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventScheduleDeletedWebhookEventEventScheduleDeleted:
		return true
	}
	return false
}

type RuntimeWebhookEventScheduleTriggeredWebhookEvent struct {
	ID        string                                                `json:"id" api:"required"`
	AgentID   string                                                `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventScheduleTriggeredWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventScheduleTriggeredWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                                `json:"fleetId" api:"required"`
	OrgID     string                                                `json:"orgId" api:"required"`
	ThreadID  string                                                `json:"threadId" api:"required,nullable"`
	Timestamp string                                                `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventScheduleTriggeredWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventScheduleTriggeredWebhookEventJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventScheduleTriggeredWebhookEvent]
type runtimeWebhookEventScheduleTriggeredWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventScheduleTriggeredWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventScheduleTriggeredWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventScheduleTriggeredWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventScheduleTriggeredWebhookEventData struct {
	ScheduleID string                                                   `json:"scheduleId" api:"required"`
	ThreadID   string                                                   `json:"threadId" api:"required"`
	JSON       runtimeWebhookEventScheduleTriggeredWebhookEventDataJSON `json:"-"`
}

// runtimeWebhookEventScheduleTriggeredWebhookEventDataJSON contains the JSON
// metadata for the struct [RuntimeWebhookEventScheduleTriggeredWebhookEventData]
type runtimeWebhookEventScheduleTriggeredWebhookEventDataJSON struct {
	ScheduleID  apijson.Field
	ThreadID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeWebhookEventScheduleTriggeredWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventScheduleTriggeredWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventScheduleTriggeredWebhookEventEvent string

const (
	RuntimeWebhookEventScheduleTriggeredWebhookEventEventScheduleTriggered RuntimeWebhookEventScheduleTriggeredWebhookEventEvent = "schedule.triggered"
)

func (r RuntimeWebhookEventScheduleTriggeredWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventScheduleTriggeredWebhookEventEventScheduleTriggered:
		return true
	}
	return false
}

type RuntimeWebhookEventConnectionAttachedWebhookEvent struct {
	ID        string                                                 `json:"id" api:"required"`
	AgentID   string                                                 `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventConnectionAttachedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventConnectionAttachedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                                 `json:"fleetId" api:"required"`
	OrgID     string                                                 `json:"orgId" api:"required"`
	ThreadID  string                                                 `json:"threadId" api:"required,nullable"`
	Timestamp string                                                 `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventConnectionAttachedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventConnectionAttachedWebhookEventJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventConnectionAttachedWebhookEvent]
type runtimeWebhookEventConnectionAttachedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventConnectionAttachedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventConnectionAttachedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventConnectionAttachedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventConnectionAttachedWebhookEventData struct {
	ConnectionID string                                                    `json:"connectionId" api:"required"`
	Provider     string                                                    `json:"provider" api:"required"`
	JSON         runtimeWebhookEventConnectionAttachedWebhookEventDataJSON `json:"-"`
}

// runtimeWebhookEventConnectionAttachedWebhookEventDataJSON contains the JSON
// metadata for the struct [RuntimeWebhookEventConnectionAttachedWebhookEventData]
type runtimeWebhookEventConnectionAttachedWebhookEventDataJSON struct {
	ConnectionID apijson.Field
	Provider     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RuntimeWebhookEventConnectionAttachedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventConnectionAttachedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventConnectionAttachedWebhookEventEvent string

const (
	RuntimeWebhookEventConnectionAttachedWebhookEventEventConnectionAttached RuntimeWebhookEventConnectionAttachedWebhookEventEvent = "connection.attached"
)

func (r RuntimeWebhookEventConnectionAttachedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventConnectionAttachedWebhookEventEventConnectionAttached:
		return true
	}
	return false
}

type RuntimeWebhookEventConnectionDetachedWebhookEvent struct {
	ID        string                                                 `json:"id" api:"required"`
	AgentID   string                                                 `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventConnectionDetachedWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventConnectionDetachedWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                                 `json:"fleetId" api:"required"`
	OrgID     string                                                 `json:"orgId" api:"required"`
	ThreadID  string                                                 `json:"threadId" api:"required,nullable"`
	Timestamp string                                                 `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventConnectionDetachedWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventConnectionDetachedWebhookEventJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventConnectionDetachedWebhookEvent]
type runtimeWebhookEventConnectionDetachedWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventConnectionDetachedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventConnectionDetachedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventConnectionDetachedWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventConnectionDetachedWebhookEventData struct {
	ConnectionID string                                                    `json:"connectionId" api:"required"`
	Provider     string                                                    `json:"provider" api:"required"`
	JSON         runtimeWebhookEventConnectionDetachedWebhookEventDataJSON `json:"-"`
}

// runtimeWebhookEventConnectionDetachedWebhookEventDataJSON contains the JSON
// metadata for the struct [RuntimeWebhookEventConnectionDetachedWebhookEventData]
type runtimeWebhookEventConnectionDetachedWebhookEventDataJSON struct {
	ConnectionID apijson.Field
	Provider     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RuntimeWebhookEventConnectionDetachedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventConnectionDetachedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventConnectionDetachedWebhookEventEvent string

const (
	RuntimeWebhookEventConnectionDetachedWebhookEventEventConnectionDetached RuntimeWebhookEventConnectionDetachedWebhookEventEvent = "connection.detached"
)

func (r RuntimeWebhookEventConnectionDetachedWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventConnectionDetachedWebhookEventEventConnectionDetached:
		return true
	}
	return false
}

type RuntimeWebhookEventWebhookTestWebhookEvent struct {
	ID        string                                          `json:"id" api:"required"`
	AgentID   string                                          `json:"agentId" api:"required"`
	Data      RuntimeWebhookEventWebhookTestWebhookEventData  `json:"data" api:"required"`
	Event     RuntimeWebhookEventWebhookTestWebhookEventEvent `json:"event" api:"required"`
	FleetID   string                                          `json:"fleetId" api:"required"`
	OrgID     string                                          `json:"orgId" api:"required"`
	ThreadID  string                                          `json:"threadId" api:"required,nullable"`
	Timestamp string                                          `json:"timestamp" api:"required"`
	JSON      runtimeWebhookEventWebhookTestWebhookEventJSON  `json:"-"`
}

// runtimeWebhookEventWebhookTestWebhookEventJSON contains the JSON metadata for
// the struct [RuntimeWebhookEventWebhookTestWebhookEvent]
type runtimeWebhookEventWebhookTestWebhookEventJSON struct {
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

func (r *RuntimeWebhookEventWebhookTestWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventWebhookTestWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RuntimeWebhookEventWebhookTestWebhookEvent) implementsRuntimeWebhookEvent() {}

type RuntimeWebhookEventWebhookTestWebhookEventData struct {
	JSON runtimeWebhookEventWebhookTestWebhookEventDataJSON `json:"-"`
}

// runtimeWebhookEventWebhookTestWebhookEventDataJSON contains the JSON metadata
// for the struct [RuntimeWebhookEventWebhookTestWebhookEventData]
type runtimeWebhookEventWebhookTestWebhookEventDataJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeWebhookEventWebhookTestWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventWebhookTestWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventWebhookTestWebhookEventEvent string

const (
	RuntimeWebhookEventWebhookTestWebhookEventEventWebhookTest RuntimeWebhookEventWebhookTestWebhookEventEvent = "webhook.test"
)

func (r RuntimeWebhookEventWebhookTestWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventWebhookTestWebhookEventEventWebhookTest:
		return true
	}
	return false
}

type RuntimeWebhookEventEvent string

const (
	RuntimeWebhookEventEventAgentCreated        RuntimeWebhookEventEvent = "agent.created"
	RuntimeWebhookEventEventAgentUpdated        RuntimeWebhookEventEvent = "agent.updated"
	RuntimeWebhookEventEventAgentDeleted        RuntimeWebhookEventEvent = "agent.deleted"
	RuntimeWebhookEventEventThreadCreated       RuntimeWebhookEventEvent = "thread.created"
	RuntimeWebhookEventEventThreadStatusChanged RuntimeWebhookEventEvent = "thread.status.changed"
	RuntimeWebhookEventEventThreadCompleted     RuntimeWebhookEventEvent = "thread.completed"
	RuntimeWebhookEventEventThreadFailed        RuntimeWebhookEventEvent = "thread.failed"
	RuntimeWebhookEventEventTurnCreated         RuntimeWebhookEventEvent = "turn.created"
	RuntimeWebhookEventEventTurnCompleted       RuntimeWebhookEventEvent = "turn.completed"
	RuntimeWebhookEventEventTurnFailed          RuntimeWebhookEventEvent = "turn.failed"
	RuntimeWebhookEventEventMessageCreated      RuntimeWebhookEventEvent = "message.created"
	RuntimeWebhookEventEventApprovalRequested   RuntimeWebhookEventEvent = "approval.requested"
	RuntimeWebhookEventEventApprovalResolved    RuntimeWebhookEventEvent = "approval.resolved"
	RuntimeWebhookEventEventApprovalGranted     RuntimeWebhookEventEvent = "approval.granted"
	RuntimeWebhookEventEventScheduleCreated     RuntimeWebhookEventEvent = "schedule.created"
	RuntimeWebhookEventEventScheduleDeleted     RuntimeWebhookEventEvent = "schedule.deleted"
	RuntimeWebhookEventEventScheduleTriggered   RuntimeWebhookEventEvent = "schedule.triggered"
	RuntimeWebhookEventEventConnectionAttached  RuntimeWebhookEventEvent = "connection.attached"
	RuntimeWebhookEventEventConnectionDetached  RuntimeWebhookEventEvent = "connection.detached"
	RuntimeWebhookEventEventWebhookTest         RuntimeWebhookEventEvent = "webhook.test"
)

func (r RuntimeWebhookEventEvent) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventEventAgentCreated, RuntimeWebhookEventEventAgentUpdated, RuntimeWebhookEventEventAgentDeleted, RuntimeWebhookEventEventThreadCreated, RuntimeWebhookEventEventThreadStatusChanged, RuntimeWebhookEventEventThreadCompleted, RuntimeWebhookEventEventThreadFailed, RuntimeWebhookEventEventTurnCreated, RuntimeWebhookEventEventTurnCompleted, RuntimeWebhookEventEventTurnFailed, RuntimeWebhookEventEventMessageCreated, RuntimeWebhookEventEventApprovalRequested, RuntimeWebhookEventEventApprovalResolved, RuntimeWebhookEventEventApprovalGranted, RuntimeWebhookEventEventScheduleCreated, RuntimeWebhookEventEventScheduleDeleted, RuntimeWebhookEventEventScheduleTriggered, RuntimeWebhookEventEventConnectionAttached, RuntimeWebhookEventEventConnectionDetached, RuntimeWebhookEventEventWebhookTest:
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

type ThreadStreamEvent struct {
	EventSeq float64               `json:"eventSeq" api:"required"`
	TurnID   string                `json:"turnId" api:"required"`
	Type     ThreadStreamEventType `json:"type" api:"required"`
	// This field can have the runtime type of
	// [ThreadStreamEventThreadStreamApprovalRequestedMessageApproval].
	Approval               interface{}               `json:"approval"`
	ApprovalID             string                    `json:"approvalId"`
	CumulativeInputTokens  float64                   `json:"cumulativeInputTokens"`
	CumulativeOutputTokens float64                   `json:"cumulativeOutputTokens"`
	Decision               ThreadStreamEventDecision `json:"decision"`
	DurationMs             float64                   `json:"durationMs"`
	Error                  string                    `json:"error" api:"nullable"`
	InputTokens            float64                   `json:"inputTokens"`
	IsError                bool                      `json:"isError"`
	ItemID                 string                    `json:"itemId"`
	ItemType               ThreadStreamEventItemType `json:"itemType"`
	// This field can have the runtime type of
	// [ThreadStreamEventThreadStreamContentMessageMessage].
	Message           interface{} `json:"message"`
	MessageCount      float64     `json:"messageCount"`
	OutputTokens      float64     `json:"outputTokens"`
	RequestedToolName string      `json:"requestedToolName"`
	Result            string      `json:"result" api:"nullable"`
	// `idle` threads can accept a new turn or be closed. `running` threads have an
	// active turn. `awaiting` threads are paused on external input such as approvals.
	// `closed` threads are terminal.
	Status      Status                `json:"status"`
	StepCount   float64               `json:"stepCount"`
	StepSeq     float64               `json:"stepSeq"`
	SubThreadID string                `json:"subThreadId"`
	ToolName    shared.ToolName       `json:"toolName"`
	TurnSeq     float64               `json:"turnSeq"`
	JSON        threadStreamEventJSON `json:"-"`
	union       ThreadStreamEventUnion
}

// threadStreamEventJSON contains the JSON metadata for the struct
// [ThreadStreamEvent]
type threadStreamEventJSON struct {
	EventSeq               apijson.Field
	TurnID                 apijson.Field
	Type                   apijson.Field
	Approval               apijson.Field
	ApprovalID             apijson.Field
	CumulativeInputTokens  apijson.Field
	CumulativeOutputTokens apijson.Field
	Decision               apijson.Field
	DurationMs             apijson.Field
	Error                  apijson.Field
	InputTokens            apijson.Field
	IsError                apijson.Field
	ItemID                 apijson.Field
	ItemType               apijson.Field
	Message                apijson.Field
	MessageCount           apijson.Field
	OutputTokens           apijson.Field
	RequestedToolName      apijson.Field
	Result                 apijson.Field
	Status                 apijson.Field
	StepCount              apijson.Field
	StepSeq                apijson.Field
	SubThreadID            apijson.Field
	ToolName               apijson.Field
	TurnSeq                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r threadStreamEventJSON) RawJSON() string {
	return r.raw
}

func (r *ThreadStreamEvent) UnmarshalJSON(data []byte) (err error) {
	*r = ThreadStreamEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ThreadStreamEventUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [ThreadStreamEventThreadStreamStatusMessage],
// [ThreadStreamEventThreadStreamContentMessage],
// [ThreadStreamEventThreadStreamProgressMessage],
// [ThreadStreamEventThreadStreamTurnStartedMessage],
// [ThreadStreamEventThreadStreamTurnCompletedMessage],
// [ThreadStreamEventThreadStreamItemStartedMessage],
// [ThreadStreamEventThreadStreamItemCompletedMessage],
// [ThreadStreamEventThreadStreamTokenUsageMessage],
// [ThreadStreamEventThreadStreamApprovalRequestedMessage],
// [ThreadStreamEventThreadStreamApprovalResolvedMessage].
func (r ThreadStreamEvent) AsUnion() ThreadStreamEventUnion {
	return r.union
}

// Union satisfied by [ThreadStreamEventThreadStreamStatusMessage],
// [ThreadStreamEventThreadStreamContentMessage],
// [ThreadStreamEventThreadStreamProgressMessage],
// [ThreadStreamEventThreadStreamTurnStartedMessage],
// [ThreadStreamEventThreadStreamTurnCompletedMessage],
// [ThreadStreamEventThreadStreamItemStartedMessage],
// [ThreadStreamEventThreadStreamItemCompletedMessage],
// [ThreadStreamEventThreadStreamTokenUsageMessage],
// [ThreadStreamEventThreadStreamApprovalRequestedMessage] or
// [ThreadStreamEventThreadStreamApprovalResolvedMessage].
type ThreadStreamEventUnion interface {
	implementsThreadStreamEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ThreadStreamEventUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ThreadStreamEventThreadStreamStatusMessage{}),
			DiscriminatorValue: "status",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ThreadStreamEventThreadStreamContentMessage{}),
			DiscriminatorValue: "message",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ThreadStreamEventThreadStreamProgressMessage{}),
			DiscriminatorValue: "progress",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ThreadStreamEventThreadStreamTurnStartedMessage{}),
			DiscriminatorValue: "turn/started",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ThreadStreamEventThreadStreamTurnCompletedMessage{}),
			DiscriminatorValue: "turn/completed",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ThreadStreamEventThreadStreamItemStartedMessage{}),
			DiscriminatorValue: "item/started",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ThreadStreamEventThreadStreamItemCompletedMessage{}),
			DiscriminatorValue: "item/completed",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ThreadStreamEventThreadStreamTokenUsageMessage{}),
			DiscriminatorValue: "token_usage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ThreadStreamEventThreadStreamApprovalRequestedMessage{}),
			DiscriminatorValue: "approval/requested",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ThreadStreamEventThreadStreamApprovalResolvedMessage{}),
			DiscriminatorValue: "approval/resolved",
		},
	)
}

type ThreadStreamEventThreadStreamStatusMessage struct {
	Error    string  `json:"error" api:"required,nullable"`
	EventSeq float64 `json:"eventSeq" api:"required"`
	Result   string  `json:"result" api:"required,nullable"`
	// `idle` threads can accept a new turn or be closed. `running` threads have an
	// active turn. `awaiting` threads are paused on external input such as approvals.
	// `closed` threads are terminal.
	Status  Status                                         `json:"status" api:"required"`
	StepSeq float64                                        `json:"stepSeq" api:"required"`
	TurnID  string                                         `json:"turnId" api:"required"`
	Type    ThreadStreamEventThreadStreamStatusMessageType `json:"type" api:"required"`
	JSON    threadStreamEventThreadStreamStatusMessageJSON `json:"-"`
}

// threadStreamEventThreadStreamStatusMessageJSON contains the JSON metadata for
// the struct [ThreadStreamEventThreadStreamStatusMessage]
type threadStreamEventThreadStreamStatusMessageJSON struct {
	Error       apijson.Field
	EventSeq    apijson.Field
	Result      apijson.Field
	Status      apijson.Field
	StepSeq     apijson.Field
	TurnID      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreadStreamEventThreadStreamStatusMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadStreamEventThreadStreamStatusMessageJSON) RawJSON() string {
	return r.raw
}

func (r ThreadStreamEventThreadStreamStatusMessage) implementsThreadStreamEvent() {}

type ThreadStreamEventThreadStreamStatusMessageType string

const (
	ThreadStreamEventThreadStreamStatusMessageTypeStatus ThreadStreamEventThreadStreamStatusMessageType = "status"
)

func (r ThreadStreamEventThreadStreamStatusMessageType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamStatusMessageTypeStatus:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamContentMessage struct {
	EventSeq float64                                            `json:"eventSeq" api:"required"`
	Message  ThreadStreamEventThreadStreamContentMessageMessage `json:"message" api:"required"`
	StepSeq  float64                                            `json:"stepSeq" api:"required"`
	TurnID   string                                             `json:"turnId" api:"required"`
	Type     ThreadStreamEventThreadStreamContentMessageType    `json:"type" api:"required"`
	JSON     threadStreamEventThreadStreamContentMessageJSON    `json:"-"`
}

// threadStreamEventThreadStreamContentMessageJSON contains the JSON metadata for
// the struct [ThreadStreamEventThreadStreamContentMessage]
type threadStreamEventThreadStreamContentMessageJSON struct {
	EventSeq    apijson.Field
	Message     apijson.Field
	StepSeq     apijson.Field
	TurnID      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreadStreamEventThreadStreamContentMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadStreamEventThreadStreamContentMessageJSON) RawJSON() string {
	return r.raw
}

func (r ThreadStreamEventThreadStreamContentMessage) implementsThreadStreamEvent() {}

type ThreadStreamEventThreadStreamContentMessageMessage struct {
	Content   []ThreadStreamEventThreadStreamContentMessageMessageContent `json:"content" api:"required"`
	CreatedAt string                                                      `json:"createdAt" api:"required"`
	Role      ThreadStreamEventThreadStreamContentMessageMessageRole      `json:"role" api:"required"`
	Seq       float64                                                     `json:"seq" api:"required"`
	JSON      threadStreamEventThreadStreamContentMessageMessageJSON      `json:"-"`
}

// threadStreamEventThreadStreamContentMessageMessageJSON contains the JSON
// metadata for the struct [ThreadStreamEventThreadStreamContentMessageMessage]
type threadStreamEventThreadStreamContentMessageMessageJSON struct {
	Content     apijson.Field
	CreatedAt   apijson.Field
	Role        apijson.Field
	Seq         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreadStreamEventThreadStreamContentMessageMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadStreamEventThreadStreamContentMessageMessageJSON) RawJSON() string {
	return r.raw
}

type ThreadStreamEventThreadStreamContentMessageMessageContent struct {
	Type ThreadStreamEventThreadStreamContentMessageMessageContentType `json:"type" api:"required"`
	ID   string                                                        `json:"id"`
	// This field can have the runtime type of [string], [interface{}].
	Content      interface{} `json:"content"`
	DeniedByUser bool        `json:"deniedByUser"`
	// JSON-serialized tool input payload. Parse as JSON before reading tool-specific
	// fields.
	Input   string          `json:"input"`
	IsError bool            `json:"isError"`
	Name    shared.ToolName `json:"name"`
	// This field can have the runtime type of [map[string]map[string]string].
	ProviderMetadata interface{}                                                   `json:"providerMetadata"`
	Text             string                                                        `json:"text"`
	ToolUseID        string                                                        `json:"toolUseId"`
	JSON             threadStreamEventThreadStreamContentMessageMessageContentJSON `json:"-"`
	union            ThreadStreamEventThreadStreamContentMessageMessageContentUnion
}

// threadStreamEventThreadStreamContentMessageMessageContentJSON contains the JSON
// metadata for the struct
// [ThreadStreamEventThreadStreamContentMessageMessageContent]
type threadStreamEventThreadStreamContentMessageMessageContentJSON struct {
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

func (r threadStreamEventThreadStreamContentMessageMessageContentJSON) RawJSON() string {
	return r.raw
}

func (r *ThreadStreamEventThreadStreamContentMessageMessageContent) UnmarshalJSON(data []byte) (err error) {
	*r = ThreadStreamEventThreadStreamContentMessageMessageContent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [ThreadStreamEventThreadStreamContentMessageMessageContentUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamTextContentBlock],
// [ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolUseContentBlock],
// [ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolResultContentBlock],
// [ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamServerToolUseContentBlock],
// [ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamWebSearchToolResultContentBlock].
func (r ThreadStreamEventThreadStreamContentMessageMessageContent) AsUnion() ThreadStreamEventThreadStreamContentMessageMessageContentUnion {
	return r.union
}

// Union satisfied by
// [ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamTextContentBlock],
// [ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolUseContentBlock],
// [ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolResultContentBlock],
// [ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamServerToolUseContentBlock]
// or
// [ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamWebSearchToolResultContentBlock].
type ThreadStreamEventThreadStreamContentMessageMessageContentUnion interface {
	implementsThreadStreamEventThreadStreamContentMessageMessageContent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ThreadStreamEventThreadStreamContentMessageMessageContentUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamTextContentBlock{}),
			DiscriminatorValue: "text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolUseContentBlock{}),
			DiscriminatorValue: "tool_use",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolResultContentBlock{}),
			DiscriminatorValue: "tool_result",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamServerToolUseContentBlock{}),
			DiscriminatorValue: "server_tool_use",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamWebSearchToolResultContentBlock{}),
			DiscriminatorValue: "web_search_tool_result",
		},
	)
}

type ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamTextContentBlock struct {
	Text             string                                                                                    `json:"text" api:"required"`
	Type             ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamTextContentBlockType `json:"type" api:"required"`
	ProviderMetadata map[string]map[string]string                                                              `json:"providerMetadata"`
	JSON             threadStreamEventThreadStreamContentMessageMessageContentThreadStreamTextContentBlockJSON `json:"-"`
}

// threadStreamEventThreadStreamContentMessageMessageContentThreadStreamTextContentBlockJSON
// contains the JSON metadata for the struct
// [ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamTextContentBlock]
type threadStreamEventThreadStreamContentMessageMessageContentThreadStreamTextContentBlockJSON struct {
	Text             apijson.Field
	Type             apijson.Field
	ProviderMetadata apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamTextContentBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadStreamEventThreadStreamContentMessageMessageContentThreadStreamTextContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamTextContentBlock) implementsThreadStreamEventThreadStreamContentMessageMessageContent() {
}

type ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamTextContentBlockType string

const (
	ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamTextContentBlockTypeText ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamTextContentBlockType = "text"
)

func (r ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamTextContentBlockType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamTextContentBlockTypeText:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolUseContentBlock struct {
	ID string `json:"id" api:"required"`
	// JSON-serialized tool input payload. Parse as JSON before reading tool-specific
	// fields.
	Input            string                                                                                       `json:"input" api:"required"`
	Name             shared.ToolName                                                                              `json:"name" api:"required"`
	Type             ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolUseContentBlockType `json:"type" api:"required"`
	ProviderMetadata map[string]map[string]string                                                                 `json:"providerMetadata"`
	JSON             threadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolUseContentBlockJSON `json:"-"`
}

// threadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolUseContentBlockJSON
// contains the JSON metadata for the struct
// [ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolUseContentBlock]
type threadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolUseContentBlockJSON struct {
	ID               apijson.Field
	Input            apijson.Field
	Name             apijson.Field
	Type             apijson.Field
	ProviderMetadata apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolUseContentBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolUseContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolUseContentBlock) implementsThreadStreamEventThreadStreamContentMessageMessageContent() {
}

type ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolUseContentBlockType string

const (
	ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolUseContentBlockTypeToolUse ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolUseContentBlockType = "tool_use"
)

func (r ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolUseContentBlockType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolUseContentBlockTypeToolUse:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolResultContentBlock struct {
	Content          string                                                                                          `json:"content" api:"required"`
	IsError          bool                                                                                            `json:"isError" api:"required"`
	ToolUseID        string                                                                                          `json:"toolUseId" api:"required"`
	Type             ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolResultContentBlockType `json:"type" api:"required"`
	DeniedByUser     bool                                                                                            `json:"deniedByUser"`
	ProviderMetadata map[string]map[string]string                                                                    `json:"providerMetadata"`
	JSON             threadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolResultContentBlockJSON `json:"-"`
}

// threadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolResultContentBlockJSON
// contains the JSON metadata for the struct
// [ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolResultContentBlock]
type threadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolResultContentBlockJSON struct {
	Content          apijson.Field
	IsError          apijson.Field
	ToolUseID        apijson.Field
	Type             apijson.Field
	DeniedByUser     apijson.Field
	ProviderMetadata apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolResultContentBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolResultContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolResultContentBlock) implementsThreadStreamEventThreadStreamContentMessageMessageContent() {
}

type ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolResultContentBlockType string

const (
	ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolResultContentBlockTypeToolResult ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolResultContentBlockType = "tool_result"
)

func (r ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolResultContentBlockType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamToolResultContentBlockTypeToolResult:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamServerToolUseContentBlock struct {
	ID string `json:"id" api:"required"`
	// JSON-serialized tool input payload. Parse as JSON before reading tool-specific
	// fields.
	Input            string                                                                                             `json:"input" api:"required"`
	Name             string                                                                                             `json:"name" api:"required"`
	Type             ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamServerToolUseContentBlockType `json:"type" api:"required"`
	ProviderMetadata map[string]map[string]string                                                                       `json:"providerMetadata"`
	JSON             threadStreamEventThreadStreamContentMessageMessageContentThreadStreamServerToolUseContentBlockJSON `json:"-"`
}

// threadStreamEventThreadStreamContentMessageMessageContentThreadStreamServerToolUseContentBlockJSON
// contains the JSON metadata for the struct
// [ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamServerToolUseContentBlock]
type threadStreamEventThreadStreamContentMessageMessageContentThreadStreamServerToolUseContentBlockJSON struct {
	ID               apijson.Field
	Input            apijson.Field
	Name             apijson.Field
	Type             apijson.Field
	ProviderMetadata apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamServerToolUseContentBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadStreamEventThreadStreamContentMessageMessageContentThreadStreamServerToolUseContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamServerToolUseContentBlock) implementsThreadStreamEventThreadStreamContentMessageMessageContent() {
}

type ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamServerToolUseContentBlockType string

const (
	ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamServerToolUseContentBlockTypeServerToolUse ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamServerToolUseContentBlockType = "server_tool_use"
)

func (r ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamServerToolUseContentBlockType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamServerToolUseContentBlockTypeServerToolUse:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamWebSearchToolResultContentBlock struct {
	// Web search result payload. The runtime returns either an array of web search
	// results or an error object.
	Content   interface{}                                                                                              `json:"content" api:"required"`
	ToolUseID string                                                                                                   `json:"toolUseId" api:"required"`
	Type      ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamWebSearchToolResultContentBlockType `json:"type" api:"required"`
	JSON      threadStreamEventThreadStreamContentMessageMessageContentThreadStreamWebSearchToolResultContentBlockJSON `json:"-"`
}

// threadStreamEventThreadStreamContentMessageMessageContentThreadStreamWebSearchToolResultContentBlockJSON
// contains the JSON metadata for the struct
// [ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamWebSearchToolResultContentBlock]
type threadStreamEventThreadStreamContentMessageMessageContentThreadStreamWebSearchToolResultContentBlockJSON struct {
	Content     apijson.Field
	ToolUseID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamWebSearchToolResultContentBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadStreamEventThreadStreamContentMessageMessageContentThreadStreamWebSearchToolResultContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamWebSearchToolResultContentBlock) implementsThreadStreamEventThreadStreamContentMessageMessageContent() {
}

type ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamWebSearchToolResultContentBlockType string

const (
	ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamWebSearchToolResultContentBlockTypeWebSearchToolResult ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamWebSearchToolResultContentBlockType = "web_search_tool_result"
)

func (r ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamWebSearchToolResultContentBlockType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamContentMessageMessageContentThreadStreamWebSearchToolResultContentBlockTypeWebSearchToolResult:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamContentMessageMessageContentType string

const (
	ThreadStreamEventThreadStreamContentMessageMessageContentTypeText                ThreadStreamEventThreadStreamContentMessageMessageContentType = "text"
	ThreadStreamEventThreadStreamContentMessageMessageContentTypeToolUse             ThreadStreamEventThreadStreamContentMessageMessageContentType = "tool_use"
	ThreadStreamEventThreadStreamContentMessageMessageContentTypeToolResult          ThreadStreamEventThreadStreamContentMessageMessageContentType = "tool_result"
	ThreadStreamEventThreadStreamContentMessageMessageContentTypeServerToolUse       ThreadStreamEventThreadStreamContentMessageMessageContentType = "server_tool_use"
	ThreadStreamEventThreadStreamContentMessageMessageContentTypeWebSearchToolResult ThreadStreamEventThreadStreamContentMessageMessageContentType = "web_search_tool_result"
)

func (r ThreadStreamEventThreadStreamContentMessageMessageContentType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamContentMessageMessageContentTypeText, ThreadStreamEventThreadStreamContentMessageMessageContentTypeToolUse, ThreadStreamEventThreadStreamContentMessageMessageContentTypeToolResult, ThreadStreamEventThreadStreamContentMessageMessageContentTypeServerToolUse, ThreadStreamEventThreadStreamContentMessageMessageContentTypeWebSearchToolResult:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamContentMessageMessageRole string

const (
	ThreadStreamEventThreadStreamContentMessageMessageRoleUser      ThreadStreamEventThreadStreamContentMessageMessageRole = "user"
	ThreadStreamEventThreadStreamContentMessageMessageRoleAssistant ThreadStreamEventThreadStreamContentMessageMessageRole = "assistant"
)

func (r ThreadStreamEventThreadStreamContentMessageMessageRole) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamContentMessageMessageRoleUser, ThreadStreamEventThreadStreamContentMessageMessageRoleAssistant:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamContentMessageType string

const (
	ThreadStreamEventThreadStreamContentMessageTypeMessage ThreadStreamEventThreadStreamContentMessageType = "message"
)

func (r ThreadStreamEventThreadStreamContentMessageType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamContentMessageTypeMessage:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamProgressMessage struct {
	EventSeq     float64                                          `json:"eventSeq" api:"required"`
	MessageCount float64                                          `json:"messageCount" api:"required"`
	StepCount    float64                                          `json:"stepCount" api:"required"`
	StepSeq      float64                                          `json:"stepSeq" api:"required"`
	TurnID       string                                           `json:"turnId" api:"required"`
	Type         ThreadStreamEventThreadStreamProgressMessageType `json:"type" api:"required"`
	JSON         threadStreamEventThreadStreamProgressMessageJSON `json:"-"`
}

// threadStreamEventThreadStreamProgressMessageJSON contains the JSON metadata for
// the struct [ThreadStreamEventThreadStreamProgressMessage]
type threadStreamEventThreadStreamProgressMessageJSON struct {
	EventSeq     apijson.Field
	MessageCount apijson.Field
	StepCount    apijson.Field
	StepSeq      apijson.Field
	TurnID       apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreadStreamEventThreadStreamProgressMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadStreamEventThreadStreamProgressMessageJSON) RawJSON() string {
	return r.raw
}

func (r ThreadStreamEventThreadStreamProgressMessage) implementsThreadStreamEvent() {}

type ThreadStreamEventThreadStreamProgressMessageType string

const (
	ThreadStreamEventThreadStreamProgressMessageTypeProgress ThreadStreamEventThreadStreamProgressMessageType = "progress"
)

func (r ThreadStreamEventThreadStreamProgressMessageType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamProgressMessageTypeProgress:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamTurnStartedMessage struct {
	EventSeq float64                                             `json:"eventSeq" api:"required"`
	TurnID   string                                              `json:"turnId" api:"required"`
	TurnSeq  float64                                             `json:"turnSeq" api:"required"`
	Type     ThreadStreamEventThreadStreamTurnStartedMessageType `json:"type" api:"required"`
	JSON     threadStreamEventThreadStreamTurnStartedMessageJSON `json:"-"`
}

// threadStreamEventThreadStreamTurnStartedMessageJSON contains the JSON metadata
// for the struct [ThreadStreamEventThreadStreamTurnStartedMessage]
type threadStreamEventThreadStreamTurnStartedMessageJSON struct {
	EventSeq    apijson.Field
	TurnID      apijson.Field
	TurnSeq     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreadStreamEventThreadStreamTurnStartedMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadStreamEventThreadStreamTurnStartedMessageJSON) RawJSON() string {
	return r.raw
}

func (r ThreadStreamEventThreadStreamTurnStartedMessage) implementsThreadStreamEvent() {}

type ThreadStreamEventThreadStreamTurnStartedMessageType string

const (
	ThreadStreamEventThreadStreamTurnStartedMessageTypeTurnStarted ThreadStreamEventThreadStreamTurnStartedMessageType = "turn/started"
)

func (r ThreadStreamEventThreadStreamTurnStartedMessageType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamTurnStartedMessageTypeTurnStarted:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamTurnCompletedMessage struct {
	Error    string                                                  `json:"error" api:"required,nullable"`
	EventSeq float64                                                 `json:"eventSeq" api:"required"`
	Result   string                                                  `json:"result" api:"required,nullable"`
	Status   ThreadStreamEventThreadStreamTurnCompletedMessageStatus `json:"status" api:"required"`
	TurnID   string                                                  `json:"turnId" api:"required"`
	TurnSeq  float64                                                 `json:"turnSeq" api:"required"`
	Type     ThreadStreamEventThreadStreamTurnCompletedMessageType   `json:"type" api:"required"`
	JSON     threadStreamEventThreadStreamTurnCompletedMessageJSON   `json:"-"`
}

// threadStreamEventThreadStreamTurnCompletedMessageJSON contains the JSON metadata
// for the struct [ThreadStreamEventThreadStreamTurnCompletedMessage]
type threadStreamEventThreadStreamTurnCompletedMessageJSON struct {
	Error       apijson.Field
	EventSeq    apijson.Field
	Result      apijson.Field
	Status      apijson.Field
	TurnID      apijson.Field
	TurnSeq     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreadStreamEventThreadStreamTurnCompletedMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadStreamEventThreadStreamTurnCompletedMessageJSON) RawJSON() string {
	return r.raw
}

func (r ThreadStreamEventThreadStreamTurnCompletedMessage) implementsThreadStreamEvent() {}

type ThreadStreamEventThreadStreamTurnCompletedMessageStatus string

const (
	ThreadStreamEventThreadStreamTurnCompletedMessageStatusCompleted ThreadStreamEventThreadStreamTurnCompletedMessageStatus = "completed"
	ThreadStreamEventThreadStreamTurnCompletedMessageStatusFailed    ThreadStreamEventThreadStreamTurnCompletedMessageStatus = "failed"
)

func (r ThreadStreamEventThreadStreamTurnCompletedMessageStatus) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamTurnCompletedMessageStatusCompleted, ThreadStreamEventThreadStreamTurnCompletedMessageStatusFailed:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamTurnCompletedMessageType string

const (
	ThreadStreamEventThreadStreamTurnCompletedMessageTypeTurnCompleted ThreadStreamEventThreadStreamTurnCompletedMessageType = "turn/completed"
)

func (r ThreadStreamEventThreadStreamTurnCompletedMessageType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamTurnCompletedMessageTypeTurnCompleted:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamItemStartedMessage struct {
	EventSeq          float64                                                 `json:"eventSeq" api:"required"`
	ItemID            string                                                  `json:"itemId" api:"required"`
	ItemType          ThreadStreamEventThreadStreamItemStartedMessageItemType `json:"itemType" api:"required"`
	TurnID            string                                                  `json:"turnId" api:"required"`
	Type              ThreadStreamEventThreadStreamItemStartedMessageType     `json:"type" api:"required"`
	RequestedToolName string                                                  `json:"requestedToolName"`
	SubThreadID       string                                                  `json:"subThreadId"`
	ToolName          shared.ToolName                                         `json:"toolName"`
	JSON              threadStreamEventThreadStreamItemStartedMessageJSON     `json:"-"`
}

// threadStreamEventThreadStreamItemStartedMessageJSON contains the JSON metadata
// for the struct [ThreadStreamEventThreadStreamItemStartedMessage]
type threadStreamEventThreadStreamItemStartedMessageJSON struct {
	EventSeq          apijson.Field
	ItemID            apijson.Field
	ItemType          apijson.Field
	TurnID            apijson.Field
	Type              apijson.Field
	RequestedToolName apijson.Field
	SubThreadID       apijson.Field
	ToolName          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ThreadStreamEventThreadStreamItemStartedMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadStreamEventThreadStreamItemStartedMessageJSON) RawJSON() string {
	return r.raw
}

func (r ThreadStreamEventThreadStreamItemStartedMessage) implementsThreadStreamEvent() {}

type ThreadStreamEventThreadStreamItemStartedMessageItemType string

const (
	ThreadStreamEventThreadStreamItemStartedMessageItemTypeToolCall     ThreadStreamEventThreadStreamItemStartedMessageItemType = "tool_call"
	ThreadStreamEventThreadStreamItemStartedMessageItemTypeAgentMessage ThreadStreamEventThreadStreamItemStartedMessageItemType = "agent_message"
	ThreadStreamEventThreadStreamItemStartedMessageItemTypeLlmCall      ThreadStreamEventThreadStreamItemStartedMessageItemType = "llm_call"
	ThreadStreamEventThreadStreamItemStartedMessageItemTypeSubThread    ThreadStreamEventThreadStreamItemStartedMessageItemType = "sub_thread"
)

func (r ThreadStreamEventThreadStreamItemStartedMessageItemType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamItemStartedMessageItemTypeToolCall, ThreadStreamEventThreadStreamItemStartedMessageItemTypeAgentMessage, ThreadStreamEventThreadStreamItemStartedMessageItemTypeLlmCall, ThreadStreamEventThreadStreamItemStartedMessageItemTypeSubThread:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamItemStartedMessageType string

const (
	ThreadStreamEventThreadStreamItemStartedMessageTypeItemStarted ThreadStreamEventThreadStreamItemStartedMessageType = "item/started"
)

func (r ThreadStreamEventThreadStreamItemStartedMessageType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamItemStartedMessageTypeItemStarted:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamItemCompletedMessage struct {
	DurationMs        float64                                                   `json:"durationMs" api:"required"`
	EventSeq          float64                                                   `json:"eventSeq" api:"required"`
	ItemID            string                                                    `json:"itemId" api:"required"`
	ItemType          ThreadStreamEventThreadStreamItemCompletedMessageItemType `json:"itemType" api:"required"`
	TurnID            string                                                    `json:"turnId" api:"required"`
	Type              ThreadStreamEventThreadStreamItemCompletedMessageType     `json:"type" api:"required"`
	IsError           bool                                                      `json:"isError"`
	RequestedToolName string                                                    `json:"requestedToolName"`
	JSON              threadStreamEventThreadStreamItemCompletedMessageJSON     `json:"-"`
}

// threadStreamEventThreadStreamItemCompletedMessageJSON contains the JSON metadata
// for the struct [ThreadStreamEventThreadStreamItemCompletedMessage]
type threadStreamEventThreadStreamItemCompletedMessageJSON struct {
	DurationMs        apijson.Field
	EventSeq          apijson.Field
	ItemID            apijson.Field
	ItemType          apijson.Field
	TurnID            apijson.Field
	Type              apijson.Field
	IsError           apijson.Field
	RequestedToolName apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ThreadStreamEventThreadStreamItemCompletedMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadStreamEventThreadStreamItemCompletedMessageJSON) RawJSON() string {
	return r.raw
}

func (r ThreadStreamEventThreadStreamItemCompletedMessage) implementsThreadStreamEvent() {}

type ThreadStreamEventThreadStreamItemCompletedMessageItemType string

const (
	ThreadStreamEventThreadStreamItemCompletedMessageItemTypeToolCall     ThreadStreamEventThreadStreamItemCompletedMessageItemType = "tool_call"
	ThreadStreamEventThreadStreamItemCompletedMessageItemTypeAgentMessage ThreadStreamEventThreadStreamItemCompletedMessageItemType = "agent_message"
	ThreadStreamEventThreadStreamItemCompletedMessageItemTypeLlmCall      ThreadStreamEventThreadStreamItemCompletedMessageItemType = "llm_call"
	ThreadStreamEventThreadStreamItemCompletedMessageItemTypeSubThread    ThreadStreamEventThreadStreamItemCompletedMessageItemType = "sub_thread"
)

func (r ThreadStreamEventThreadStreamItemCompletedMessageItemType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamItemCompletedMessageItemTypeToolCall, ThreadStreamEventThreadStreamItemCompletedMessageItemTypeAgentMessage, ThreadStreamEventThreadStreamItemCompletedMessageItemTypeLlmCall, ThreadStreamEventThreadStreamItemCompletedMessageItemTypeSubThread:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamItemCompletedMessageType string

const (
	ThreadStreamEventThreadStreamItemCompletedMessageTypeItemCompleted ThreadStreamEventThreadStreamItemCompletedMessageType = "item/completed"
)

func (r ThreadStreamEventThreadStreamItemCompletedMessageType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamItemCompletedMessageTypeItemCompleted:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamTokenUsageMessage struct {
	CumulativeInputTokens  float64                                            `json:"cumulativeInputTokens" api:"required"`
	CumulativeOutputTokens float64                                            `json:"cumulativeOutputTokens" api:"required"`
	EventSeq               float64                                            `json:"eventSeq" api:"required"`
	InputTokens            float64                                            `json:"inputTokens" api:"required"`
	OutputTokens           float64                                            `json:"outputTokens" api:"required"`
	TurnID                 string                                             `json:"turnId" api:"required"`
	Type                   ThreadStreamEventThreadStreamTokenUsageMessageType `json:"type" api:"required"`
	JSON                   threadStreamEventThreadStreamTokenUsageMessageJSON `json:"-"`
}

// threadStreamEventThreadStreamTokenUsageMessageJSON contains the JSON metadata
// for the struct [ThreadStreamEventThreadStreamTokenUsageMessage]
type threadStreamEventThreadStreamTokenUsageMessageJSON struct {
	CumulativeInputTokens  apijson.Field
	CumulativeOutputTokens apijson.Field
	EventSeq               apijson.Field
	InputTokens            apijson.Field
	OutputTokens           apijson.Field
	TurnID                 apijson.Field
	Type                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ThreadStreamEventThreadStreamTokenUsageMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadStreamEventThreadStreamTokenUsageMessageJSON) RawJSON() string {
	return r.raw
}

func (r ThreadStreamEventThreadStreamTokenUsageMessage) implementsThreadStreamEvent() {}

type ThreadStreamEventThreadStreamTokenUsageMessageType string

const (
	ThreadStreamEventThreadStreamTokenUsageMessageTypeTokenUsage ThreadStreamEventThreadStreamTokenUsageMessageType = "token_usage"
)

func (r ThreadStreamEventThreadStreamTokenUsageMessageType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamTokenUsageMessageTypeTokenUsage:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamApprovalRequestedMessage struct {
	Approval ThreadStreamEventThreadStreamApprovalRequestedMessageApproval `json:"approval" api:"required"`
	EventSeq float64                                                       `json:"eventSeq" api:"required"`
	TurnID   string                                                        `json:"turnId" api:"required"`
	Type     ThreadStreamEventThreadStreamApprovalRequestedMessageType     `json:"type" api:"required"`
	JSON     threadStreamEventThreadStreamApprovalRequestedMessageJSON     `json:"-"`
}

// threadStreamEventThreadStreamApprovalRequestedMessageJSON contains the JSON
// metadata for the struct [ThreadStreamEventThreadStreamApprovalRequestedMessage]
type threadStreamEventThreadStreamApprovalRequestedMessageJSON struct {
	Approval    apijson.Field
	EventSeq    apijson.Field
	TurnID      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreadStreamEventThreadStreamApprovalRequestedMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadStreamEventThreadStreamApprovalRequestedMessageJSON) RawJSON() string {
	return r.raw
}

func (r ThreadStreamEventThreadStreamApprovalRequestedMessage) implementsThreadStreamEvent() {}

type ThreadStreamEventThreadStreamApprovalRequestedMessageApproval struct {
	ID                string                                                              `json:"id" api:"required"`
	CreatedAt         string                                                              `json:"createdAt" api:"required"`
	RequestedToolName string                                                              `json:"requestedToolName" api:"required"`
	ResolvedAt        string                                                              `json:"resolvedAt" api:"required,nullable"`
	Status            ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatus `json:"status" api:"required"`
	ThreadID          string                                                              `json:"threadId" api:"required"`
	TimeoutAt         string                                                              `json:"timeoutAt" api:"required,nullable"`
	TimeoutMs         float64                                                             `json:"timeoutMs" api:"required,nullable"`
	ToolIndex         float64                                                             `json:"toolIndex" api:"required"`
	// JSON-serialized tool input payload. Parse as JSON before reading tool-specific
	// fields.
	ToolInput         string                                                            `json:"toolInput" api:"required"`
	ToolName          shared.ToolName                                                   `json:"toolName" api:"required"`
	ToolUseID         string                                                            `json:"toolUseId" api:"required"`
	TurnID            string                                                            `json:"turnId" api:"required"`
	ToolSourceID      string                                                            `json:"toolSourceId"`
	ToolSourceVersion float64                                                           `json:"toolSourceVersion"`
	JSON              threadStreamEventThreadStreamApprovalRequestedMessageApprovalJSON `json:"-"`
}

// threadStreamEventThreadStreamApprovalRequestedMessageApprovalJSON contains the
// JSON metadata for the struct
// [ThreadStreamEventThreadStreamApprovalRequestedMessageApproval]
type threadStreamEventThreadStreamApprovalRequestedMessageApprovalJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	RequestedToolName apijson.Field
	ResolvedAt        apijson.Field
	Status            apijson.Field
	ThreadID          apijson.Field
	TimeoutAt         apijson.Field
	TimeoutMs         apijson.Field
	ToolIndex         apijson.Field
	ToolInput         apijson.Field
	ToolName          apijson.Field
	ToolUseID         apijson.Field
	TurnID            apijson.Field
	ToolSourceID      apijson.Field
	ToolSourceVersion apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ThreadStreamEventThreadStreamApprovalRequestedMessageApproval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadStreamEventThreadStreamApprovalRequestedMessageApprovalJSON) RawJSON() string {
	return r.raw
}

type ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatus string

const (
	ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatusPending   ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatus = "pending"
	ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatusApproved  ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatus = "approved"
	ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatusDenied    ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatus = "denied"
	ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatusCancelled ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatus = "cancelled"
	ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatusTimedOut  ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatus = "timed_out"
)

func (r ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatus) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatusPending, ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatusApproved, ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatusDenied, ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatusCancelled, ThreadStreamEventThreadStreamApprovalRequestedMessageApprovalStatusTimedOut:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamApprovalRequestedMessageType string

const (
	ThreadStreamEventThreadStreamApprovalRequestedMessageTypeApprovalRequested ThreadStreamEventThreadStreamApprovalRequestedMessageType = "approval/requested"
)

func (r ThreadStreamEventThreadStreamApprovalRequestedMessageType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamApprovalRequestedMessageTypeApprovalRequested:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamApprovalResolvedMessage struct {
	ApprovalID string                                                       `json:"approvalId" api:"required"`
	Decision   ThreadStreamEventThreadStreamApprovalResolvedMessageDecision `json:"decision" api:"required"`
	EventSeq   float64                                                      `json:"eventSeq" api:"required"`
	TurnID     string                                                       `json:"turnId" api:"required"`
	Type       ThreadStreamEventThreadStreamApprovalResolvedMessageType     `json:"type" api:"required"`
	JSON       threadStreamEventThreadStreamApprovalResolvedMessageJSON     `json:"-"`
}

// threadStreamEventThreadStreamApprovalResolvedMessageJSON contains the JSON
// metadata for the struct [ThreadStreamEventThreadStreamApprovalResolvedMessage]
type threadStreamEventThreadStreamApprovalResolvedMessageJSON struct {
	ApprovalID  apijson.Field
	Decision    apijson.Field
	EventSeq    apijson.Field
	TurnID      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreadStreamEventThreadStreamApprovalResolvedMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadStreamEventThreadStreamApprovalResolvedMessageJSON) RawJSON() string {
	return r.raw
}

func (r ThreadStreamEventThreadStreamApprovalResolvedMessage) implementsThreadStreamEvent() {}

type ThreadStreamEventThreadStreamApprovalResolvedMessageDecision string

const (
	ThreadStreamEventThreadStreamApprovalResolvedMessageDecisionApprove ThreadStreamEventThreadStreamApprovalResolvedMessageDecision = "approve"
	ThreadStreamEventThreadStreamApprovalResolvedMessageDecisionDeny    ThreadStreamEventThreadStreamApprovalResolvedMessageDecision = "deny"
	ThreadStreamEventThreadStreamApprovalResolvedMessageDecisionCancel  ThreadStreamEventThreadStreamApprovalResolvedMessageDecision = "cancel"
	ThreadStreamEventThreadStreamApprovalResolvedMessageDecisionTimeout ThreadStreamEventThreadStreamApprovalResolvedMessageDecision = "timeout"
)

func (r ThreadStreamEventThreadStreamApprovalResolvedMessageDecision) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamApprovalResolvedMessageDecisionApprove, ThreadStreamEventThreadStreamApprovalResolvedMessageDecisionDeny, ThreadStreamEventThreadStreamApprovalResolvedMessageDecisionCancel, ThreadStreamEventThreadStreamApprovalResolvedMessageDecisionTimeout:
		return true
	}
	return false
}

type ThreadStreamEventThreadStreamApprovalResolvedMessageType string

const (
	ThreadStreamEventThreadStreamApprovalResolvedMessageTypeApprovalResolved ThreadStreamEventThreadStreamApprovalResolvedMessageType = "approval/resolved"
)

func (r ThreadStreamEventThreadStreamApprovalResolvedMessageType) IsKnown() bool {
	switch r {
	case ThreadStreamEventThreadStreamApprovalResolvedMessageTypeApprovalResolved:
		return true
	}
	return false
}

type ThreadStreamEventType string

const (
	ThreadStreamEventTypeStatus            ThreadStreamEventType = "status"
	ThreadStreamEventTypeMessage           ThreadStreamEventType = "message"
	ThreadStreamEventTypeProgress          ThreadStreamEventType = "progress"
	ThreadStreamEventTypeTurnStarted       ThreadStreamEventType = "turn/started"
	ThreadStreamEventTypeTurnCompleted     ThreadStreamEventType = "turn/completed"
	ThreadStreamEventTypeItemStarted       ThreadStreamEventType = "item/started"
	ThreadStreamEventTypeItemCompleted     ThreadStreamEventType = "item/completed"
	ThreadStreamEventTypeTokenUsage        ThreadStreamEventType = "token_usage"
	ThreadStreamEventTypeApprovalRequested ThreadStreamEventType = "approval/requested"
	ThreadStreamEventTypeApprovalResolved  ThreadStreamEventType = "approval/resolved"
)

func (r ThreadStreamEventType) IsKnown() bool {
	switch r {
	case ThreadStreamEventTypeStatus, ThreadStreamEventTypeMessage, ThreadStreamEventTypeProgress, ThreadStreamEventTypeTurnStarted, ThreadStreamEventTypeTurnCompleted, ThreadStreamEventTypeItemStarted, ThreadStreamEventTypeItemCompleted, ThreadStreamEventTypeTokenUsage, ThreadStreamEventTypeApprovalRequested, ThreadStreamEventTypeApprovalResolved:
		return true
	}
	return false
}

type ThreadStreamEventDecision string

const (
	ThreadStreamEventDecisionApprove ThreadStreamEventDecision = "approve"
	ThreadStreamEventDecisionDeny    ThreadStreamEventDecision = "deny"
	ThreadStreamEventDecisionCancel  ThreadStreamEventDecision = "cancel"
	ThreadStreamEventDecisionTimeout ThreadStreamEventDecision = "timeout"
)

func (r ThreadStreamEventDecision) IsKnown() bool {
	switch r {
	case ThreadStreamEventDecisionApprove, ThreadStreamEventDecisionDeny, ThreadStreamEventDecisionCancel, ThreadStreamEventDecisionTimeout:
		return true
	}
	return false
}

type ThreadStreamEventItemType string

const (
	ThreadStreamEventItemTypeToolCall     ThreadStreamEventItemType = "tool_call"
	ThreadStreamEventItemTypeAgentMessage ThreadStreamEventItemType = "agent_message"
	ThreadStreamEventItemTypeLlmCall      ThreadStreamEventItemType = "llm_call"
	ThreadStreamEventItemTypeSubThread    ThreadStreamEventItemType = "sub_thread"
)

func (r ThreadStreamEventItemType) IsKnown() bool {
	switch r {
	case ThreadStreamEventItemTypeToolCall, ThreadStreamEventItemTypeAgentMessage, ThreadStreamEventItemTypeLlmCall, ThreadStreamEventItemTypeSubThread:
		return true
	}
	return false
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

type EventStreamForAgentParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Comma-separated event type filter.
	Events param.Field[string] `query:"events"`
	// When true, starts from the beginning of the retained buffer.
	History param.Field[EventStreamForAgentParamsHistory] `query:"history"`
	// Resume an event-log stream after the last received SSE event id.
	LastEventID param.Field[string] `header:"Last-Event-ID"`
}

// URLQuery serializes [EventStreamForAgentParams]'s query parameters as
// `url.Values`.
func (r EventStreamForAgentParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// When true, starts from the beginning of the retained buffer.
type EventStreamForAgentParamsHistory string

const (
	EventStreamForAgentParamsHistoryTrue  EventStreamForAgentParamsHistory = "true"
	EventStreamForAgentParamsHistoryFalse EventStreamForAgentParamsHistory = "false"
)

func (r EventStreamForAgentParamsHistory) IsKnown() bool {
	switch r {
	case EventStreamForAgentParamsHistoryTrue, EventStreamForAgentParamsHistoryFalse:
		return true
	}
	return false
}

type EventStreamForFleetParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Comma-separated event type filter.
	Events param.Field[string] `query:"events"`
	// When true, starts from the beginning of the retained buffer.
	History param.Field[EventStreamForFleetParamsHistory] `query:"history"`
	// Resume an event-log stream after the last received SSE event id.
	LastEventID param.Field[string] `header:"Last-Event-ID"`
}

// URLQuery serializes [EventStreamForFleetParams]'s query parameters as
// `url.Values`.
func (r EventStreamForFleetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// When true, starts from the beginning of the retained buffer.
type EventStreamForFleetParamsHistory string

const (
	EventStreamForFleetParamsHistoryTrue  EventStreamForFleetParamsHistory = "true"
	EventStreamForFleetParamsHistoryFalse EventStreamForFleetParamsHistory = "false"
)

func (r EventStreamForFleetParamsHistory) IsKnown() bool {
	switch r {
	case EventStreamForFleetParamsHistoryTrue, EventStreamForFleetParamsHistoryFalse:
		return true
	}
	return false
}

type EventStreamForThreadEventsParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Comma-separated event type filter.
	Events param.Field[string] `query:"events"`
	// When true, starts from the beginning of the retained buffer.
	History param.Field[EventStreamForThreadEventsParamsHistory] `query:"history"`
	// Resume an event-log stream after the last received SSE event id.
	LastEventID param.Field[string] `header:"Last-Event-ID"`
}

// URLQuery serializes [EventStreamForThreadEventsParams]'s query parameters as
// `url.Values`.
func (r EventStreamForThreadEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// When true, starts from the beginning of the retained buffer.
type EventStreamForThreadEventsParamsHistory string

const (
	EventStreamForThreadEventsParamsHistoryTrue  EventStreamForThreadEventsParamsHistory = "true"
	EventStreamForThreadEventsParamsHistoryFalse EventStreamForThreadEventsParamsHistory = "false"
)

func (r EventStreamForThreadEventsParamsHistory) IsKnown() bool {
	switch r {
	case EventStreamForThreadEventsParamsHistoryTrue, EventStreamForThreadEventsParamsHistoryFalse:
		return true
	}
	return false
}
