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
