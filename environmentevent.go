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

// EnvironmentEventService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentEventService] method instead.
type EnvironmentEventService struct {
	Options []option.RequestOption
}

// NewEnvironmentEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEnvironmentEventService(opts ...option.RequestOption) (r *EnvironmentEventService) {
	r = &EnvironmentEventService{}
	r.Options = opts
	return
}

func (r *EnvironmentEventService) List(ctx context.Context, envID string, query EnvironmentEventListParams, opts ...option.RequestOption) (res *EnvironmentEventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if envID == "" {
		err = errors.New("missing required envId parameter")
		return nil, err
	}
	path := fmt.Sprintf("environments/%s/events", envID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// WebSocket upgrade endpoint. Set
// `Sec-WebSocket-Protocol: cell-v1, cell-auth-<API_KEY>` so the runtime can
// authenticate the stream while preserving the public subprotocol. HTTP clients
// that cannot upgrade should use `/environments/{envId}/events` as the polling
// analog.
func (r *EnvironmentEventService) Subscribe(ctx context.Context, envID string, opts ...option.RequestOption) (res *EnvironmentEventSubscribeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if envID == "" {
		err = errors.New("missing required envId parameter")
		return nil, err
	}
	path := fmt.Sprintf("environments/%s/events/subscribe", envID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type RuntimeWebhookEvent struct {
	ID            string                  `json:"id" api:"required"`
	CellID        string                  `json:"cellId" api:"required"`
	Data          map[string]interface{}  `json:"data" api:"required"`
	EnvironmentID string                  `json:"environmentId" api:"required"`
	Event         RuntimeWebhookEventType `json:"event" api:"required"`
	OrgID         string                  `json:"orgId" api:"required"`
	ThreadID      string                  `json:"threadId" api:"required,nullable"`
	Timestamp     string                  `json:"timestamp" api:"required"`
	JSON          runtimeWebhookEventJSON `json:"-"`
}

// runtimeWebhookEventJSON contains the JSON metadata for the struct
// [RuntimeWebhookEvent]
type runtimeWebhookEventJSON struct {
	ID            apijson.Field
	CellID        apijson.Field
	Data          apijson.Field
	EnvironmentID apijson.Field
	Event         apijson.Field
	OrgID         apijson.Field
	ThreadID      apijson.Field
	Timestamp     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RuntimeWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeWebhookEventJSON) RawJSON() string {
	return r.raw
}

type RuntimeWebhookEventType string

const (
	RuntimeWebhookEventTypeCellCreated         RuntimeWebhookEventType = "cell.created"
	RuntimeWebhookEventTypeCellUpdated         RuntimeWebhookEventType = "cell.updated"
	RuntimeWebhookEventTypeCellDeleted         RuntimeWebhookEventType = "cell.deleted"
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
	RuntimeWebhookEventTypeScheduleCreated     RuntimeWebhookEventType = "schedule.created"
	RuntimeWebhookEventTypeScheduleDeleted     RuntimeWebhookEventType = "schedule.deleted"
	RuntimeWebhookEventTypeScheduleTriggered   RuntimeWebhookEventType = "schedule.triggered"
	RuntimeWebhookEventTypeConnectionAttached  RuntimeWebhookEventType = "connection.attached"
	RuntimeWebhookEventTypeConnectionDetached  RuntimeWebhookEventType = "connection.detached"
	RuntimeWebhookEventTypeWebhookTest         RuntimeWebhookEventType = "webhook.test"
)

func (r RuntimeWebhookEventType) IsKnown() bool {
	switch r {
	case RuntimeWebhookEventTypeCellCreated, RuntimeWebhookEventTypeCellUpdated, RuntimeWebhookEventTypeCellDeleted, RuntimeWebhookEventTypeThreadCreated, RuntimeWebhookEventTypeThreadStatusChanged, RuntimeWebhookEventTypeThreadCompleted, RuntimeWebhookEventTypeThreadFailed, RuntimeWebhookEventTypeTurnCreated, RuntimeWebhookEventTypeTurnCompleted, RuntimeWebhookEventTypeTurnFailed, RuntimeWebhookEventTypeMessageCreated, RuntimeWebhookEventTypeApprovalRequested, RuntimeWebhookEventTypeApprovalResolved, RuntimeWebhookEventTypeScheduleCreated, RuntimeWebhookEventTypeScheduleDeleted, RuntimeWebhookEventTypeScheduleTriggered, RuntimeWebhookEventTypeConnectionAttached, RuntimeWebhookEventTypeConnectionDetached, RuntimeWebhookEventTypeWebhookTest:
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

type EnvironmentEventListResponse struct {
	Cursor  string                           `json:"cursor" api:"required,nullable"`
	Events  []SubscriptionEvent              `json:"events" api:"required"`
	HasMore bool                             `json:"hasMore" api:"required"`
	JSON    environmentEventListResponseJSON `json:"-"`
}

// environmentEventListResponseJSON contains the JSON metadata for the struct
// [EnvironmentEventListResponse]
type environmentEventListResponseJSON struct {
	Cursor      apijson.Field
	Events      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentEventListResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentEventSubscribeResponse struct {
	Error string                                `json:"error" api:"required"`
	JSON  environmentEventSubscribeResponseJSON `json:"-"`
}

// environmentEventSubscribeResponseJSON contains the JSON metadata for the struct
// [EnvironmentEventSubscribeResponse]
type environmentEventSubscribeResponseJSON struct {
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentEventSubscribeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentEventSubscribeResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentEventListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Comma-separated event type filter.
	Events param.Field[string] `query:"events"`
	// When true, starts from the beginning of the retained buffer.
	History param.Field[EnvironmentEventListParamsHistory] `query:"history"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [EnvironmentEventListParams]'s query parameters as
// `url.Values`.
func (r EnvironmentEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// When true, starts from the beginning of the retained buffer.
type EnvironmentEventListParamsHistory string

const (
	EnvironmentEventListParamsHistoryTrue  EnvironmentEventListParamsHistory = "true"
	EnvironmentEventListParamsHistoryFalse EnvironmentEventListParamsHistory = "false"
)

func (r EnvironmentEventListParamsHistory) IsKnown() bool {
	switch r {
	case EnvironmentEventListParamsHistoryTrue, EnvironmentEventListParamsHistoryFalse:
		return true
	}
	return false
}
