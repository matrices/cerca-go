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

// WebhookService contains methods and other services that help with interacting
// with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	return
}

func (r *WebhookService) New(ctx context.Context, fleetID string, body WebhookNewParams, opts ...option.RequestOption) (res *WebhookSubscriptionCreated, err error) {
	opts = slices.Concat(r.Options, opts)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("fleets/%s/webhooks", fleetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *WebhookService) Get(ctx context.Context, fleetID string, webhookID string, opts ...option.RequestOption) (res *WebhookSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return nil, err
	}
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("fleets/%s/webhooks/%s", fleetID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *WebhookService) Update(ctx context.Context, fleetID string, webhookID string, body WebhookUpdateParams, opts ...option.RequestOption) (res *WebhookSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return nil, err
	}
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("fleets/%s/webhooks/%s", fleetID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

func (r *WebhookService) List(ctx context.Context, fleetID string, query WebhookListParams, opts ...option.RequestOption) (res *pagination.SubscriptionsCursorPage[WebhookSubscription], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("fleets/%s/webhooks", fleetID)
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

func (r *WebhookService) ListAutoPaging(ctx context.Context, fleetID string, query WebhookListParams, opts ...option.RequestOption) *pagination.SubscriptionsCursorPageAutoPager[WebhookSubscription] {
	return pagination.NewSubscriptionsCursorPageAutoPager(r.List(ctx, fleetID, query, opts...))
}

func (r *WebhookService) Delete(ctx context.Context, fleetID string, webhookID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return err
	}
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return err
	}
	path := fmt.Sprintf("fleets/%s/webhooks/%s", fleetID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *WebhookService) Rotate(ctx context.Context, fleetID string, webhookID string, opts ...option.RequestOption) (res *WebhookSubscriptionCreated, err error) {
	opts = slices.Concat(r.Options, opts)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return nil, err
	}
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("fleets/%s/webhooks/%s/rotate", fleetID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

func (r *WebhookService) Test(ctx context.Context, fleetID string, webhookID string, opts ...option.RequestOption) (res *WebhookTestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if fleetID == "" {
		err = errors.New("missing required fleetId parameter")
		return nil, err
	}
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("fleets/%s/webhooks/%s/test", fleetID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type WebhookEventType string

const (
	WebhookEventTypeAgentCreated        WebhookEventType = "agent.created"
	WebhookEventTypeAgentUpdated        WebhookEventType = "agent.updated"
	WebhookEventTypeAgentDeleted        WebhookEventType = "agent.deleted"
	WebhookEventTypeThreadCreated       WebhookEventType = "thread.created"
	WebhookEventTypeThreadStatusChanged WebhookEventType = "thread.status.changed"
	WebhookEventTypeThreadCompleted     WebhookEventType = "thread.completed"
	WebhookEventTypeThreadFailed        WebhookEventType = "thread.failed"
	WebhookEventTypeTurnCreated         WebhookEventType = "turn.created"
	WebhookEventTypeTurnCompleted       WebhookEventType = "turn.completed"
	WebhookEventTypeTurnFailed          WebhookEventType = "turn.failed"
	WebhookEventTypeMessageCreated      WebhookEventType = "message.created"
	WebhookEventTypeApprovalRequested   WebhookEventType = "approval.requested"
	WebhookEventTypeApprovalResolved    WebhookEventType = "approval.resolved"
	WebhookEventTypeApprovalGranted     WebhookEventType = "approval.granted"
	WebhookEventTypeScheduleCreated     WebhookEventType = "schedule.created"
	WebhookEventTypeScheduleDeleted     WebhookEventType = "schedule.deleted"
	WebhookEventTypeScheduleTriggered   WebhookEventType = "schedule.triggered"
	WebhookEventTypeConnectionAttached  WebhookEventType = "connection.attached"
	WebhookEventTypeConnectionDetached  WebhookEventType = "connection.detached"
	WebhookEventTypeWebhookTest         WebhookEventType = "webhook.test"
)

func (r WebhookEventType) IsKnown() bool {
	switch r {
	case WebhookEventTypeAgentCreated, WebhookEventTypeAgentUpdated, WebhookEventTypeAgentDeleted, WebhookEventTypeThreadCreated, WebhookEventTypeThreadStatusChanged, WebhookEventTypeThreadCompleted, WebhookEventTypeThreadFailed, WebhookEventTypeTurnCreated, WebhookEventTypeTurnCompleted, WebhookEventTypeTurnFailed, WebhookEventTypeMessageCreated, WebhookEventTypeApprovalRequested, WebhookEventTypeApprovalResolved, WebhookEventTypeApprovalGranted, WebhookEventTypeScheduleCreated, WebhookEventTypeScheduleDeleted, WebhookEventTypeScheduleTriggered, WebhookEventTypeConnectionAttached, WebhookEventTypeConnectionDetached, WebhookEventTypeWebhookTest:
		return true
	}
	return false
}

// Webhook subscription metadata returned by list, get, and update. The one-time
// signing secret is not returned here.
type WebhookSubscription struct {
	ID        string                  `json:"id" api:"required"`
	CreatedAt string                  `json:"createdAt" api:"required"`
	Enabled   bool                    `json:"enabled" api:"required"`
	Events    []WebhookEventType      `json:"events" api:"required"`
	FleetID   string                  `json:"fleetId" api:"required"`
	UpdatedAt string                  `json:"updatedAt" api:"required"`
	URL       string                  `json:"url" api:"required"`
	JSON      webhookSubscriptionJSON `json:"-"`
}

// webhookSubscriptionJSON contains the JSON metadata for the struct
// [WebhookSubscription]
type webhookSubscriptionJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Events      apijson.Field
	FleetID     apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookSubscriptionJSON) RawJSON() string {
	return r.raw
}

// Webhook subscription response for create and rotate. Includes the one-time
// `secret` field.
type WebhookSubscriptionCreated struct {
	ID        string             `json:"id" api:"required"`
	CreatedAt string             `json:"createdAt" api:"required"`
	Enabled   bool               `json:"enabled" api:"required"`
	Events    []WebhookEventType `json:"events" api:"required"`
	FleetID   string             `json:"fleetId" api:"required"`
	// One-time webhook signing secret returned only by create and rotate responses.
	// The field name is `secret`.
	Secret    string                         `json:"secret" api:"required"`
	UpdatedAt string                         `json:"updatedAt" api:"required"`
	URL       string                         `json:"url" api:"required"`
	JSON      webhookSubscriptionCreatedJSON `json:"-"`
}

// webhookSubscriptionCreatedJSON contains the JSON metadata for the struct
// [WebhookSubscriptionCreated]
type webhookSubscriptionCreatedJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Events      apijson.Field
	FleetID     apijson.Field
	Secret      apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookSubscriptionCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookSubscriptionCreatedJSON) RawJSON() string {
	return r.raw
}

// `webhook.test` is reserved for the test endpoint and cannot be subscribed to.
type WebhookSubscriptionEventType string

const (
	WebhookSubscriptionEventTypeAgentCreated        WebhookSubscriptionEventType = "agent.created"
	WebhookSubscriptionEventTypeAgentUpdated        WebhookSubscriptionEventType = "agent.updated"
	WebhookSubscriptionEventTypeAgentDeleted        WebhookSubscriptionEventType = "agent.deleted"
	WebhookSubscriptionEventTypeThreadCreated       WebhookSubscriptionEventType = "thread.created"
	WebhookSubscriptionEventTypeThreadStatusChanged WebhookSubscriptionEventType = "thread.status.changed"
	WebhookSubscriptionEventTypeThreadCompleted     WebhookSubscriptionEventType = "thread.completed"
	WebhookSubscriptionEventTypeThreadFailed        WebhookSubscriptionEventType = "thread.failed"
	WebhookSubscriptionEventTypeTurnCreated         WebhookSubscriptionEventType = "turn.created"
	WebhookSubscriptionEventTypeTurnCompleted       WebhookSubscriptionEventType = "turn.completed"
	WebhookSubscriptionEventTypeTurnFailed          WebhookSubscriptionEventType = "turn.failed"
	WebhookSubscriptionEventTypeMessageCreated      WebhookSubscriptionEventType = "message.created"
	WebhookSubscriptionEventTypeApprovalRequested   WebhookSubscriptionEventType = "approval.requested"
	WebhookSubscriptionEventTypeApprovalResolved    WebhookSubscriptionEventType = "approval.resolved"
	WebhookSubscriptionEventTypeApprovalGranted     WebhookSubscriptionEventType = "approval.granted"
	WebhookSubscriptionEventTypeScheduleCreated     WebhookSubscriptionEventType = "schedule.created"
	WebhookSubscriptionEventTypeScheduleDeleted     WebhookSubscriptionEventType = "schedule.deleted"
	WebhookSubscriptionEventTypeScheduleTriggered   WebhookSubscriptionEventType = "schedule.triggered"
	WebhookSubscriptionEventTypeConnectionAttached  WebhookSubscriptionEventType = "connection.attached"
	WebhookSubscriptionEventTypeConnectionDetached  WebhookSubscriptionEventType = "connection.detached"
)

func (r WebhookSubscriptionEventType) IsKnown() bool {
	switch r {
	case WebhookSubscriptionEventTypeAgentCreated, WebhookSubscriptionEventTypeAgentUpdated, WebhookSubscriptionEventTypeAgentDeleted, WebhookSubscriptionEventTypeThreadCreated, WebhookSubscriptionEventTypeThreadStatusChanged, WebhookSubscriptionEventTypeThreadCompleted, WebhookSubscriptionEventTypeThreadFailed, WebhookSubscriptionEventTypeTurnCreated, WebhookSubscriptionEventTypeTurnCompleted, WebhookSubscriptionEventTypeTurnFailed, WebhookSubscriptionEventTypeMessageCreated, WebhookSubscriptionEventTypeApprovalRequested, WebhookSubscriptionEventTypeApprovalResolved, WebhookSubscriptionEventTypeApprovalGranted, WebhookSubscriptionEventTypeScheduleCreated, WebhookSubscriptionEventTypeScheduleDeleted, WebhookSubscriptionEventTypeScheduleTriggered, WebhookSubscriptionEventTypeConnectionAttached, WebhookSubscriptionEventTypeConnectionDetached:
		return true
	}
	return false
}

type WebhookTestResponse struct {
	Error      string                  `json:"error" api:"required,nullable"`
	StatusCode int64                   `json:"statusCode" api:"required,nullable"`
	Success    bool                    `json:"success" api:"required"`
	JSON       webhookTestResponseJSON `json:"-"`
}

// webhookTestResponseJSON contains the JSON metadata for the struct
// [WebhookTestResponse]
type webhookTestResponseJSON struct {
	Error       apijson.Field
	StatusCode  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookTestResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookTestResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookNewParams struct {
	// HTTPS endpoint that will receive webhook deliveries.
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
	// Event names to deliver. Omit to subscribe to all non-test events.
	Events param.Field[[]WebhookSubscriptionEventType] `json:"events"`
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookUpdateParams struct {
	// Whether deliveries are enabled for this subscription.
	Enabled param.Field[bool] `json:"enabled"`
	// Event names to deliver. Omit to subscribe to all non-test events.
	Events param.Field[[]WebhookSubscriptionEventType] `json:"events"`
	// HTTPS endpoint that will receive webhook deliveries.
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [WebhookListParams]'s query parameters as `url.Values`.
func (r WebhookListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
