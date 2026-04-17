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

// EnvironmentWebhookService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentWebhookService] method instead.
type EnvironmentWebhookService struct {
	Options []option.RequestOption
}

// NewEnvironmentWebhookService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEnvironmentWebhookService(opts ...option.RequestOption) (r *EnvironmentWebhookService) {
	r = &EnvironmentWebhookService{}
	r.Options = opts
	return
}

func (r *EnvironmentWebhookService) New(ctx context.Context, environmentID string, body EnvironmentWebhookNewParams, opts ...option.RequestOption) (res *WebhookSubscriptionCreated, err error) {
	opts = slices.Concat(r.Options, opts)
	if environmentID == "" {
		err = errors.New("missing required environmentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("environments/%s/webhooks", environmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *EnvironmentWebhookService) Get(ctx context.Context, environmentID string, webhookID string, opts ...option.RequestOption) (res *WebhookSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if environmentID == "" {
		err = errors.New("missing required environmentId parameter")
		return nil, err
	}
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("environments/%s/webhooks/%s", environmentID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *EnvironmentWebhookService) Update(ctx context.Context, environmentID string, webhookID string, body EnvironmentWebhookUpdateParams, opts ...option.RequestOption) (res *WebhookSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if environmentID == "" {
		err = errors.New("missing required environmentId parameter")
		return nil, err
	}
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("environments/%s/webhooks/%s", environmentID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

func (r *EnvironmentWebhookService) List(ctx context.Context, environmentID string, query EnvironmentWebhookListParams, opts ...option.RequestOption) (res *EnvironmentWebhookListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if environmentID == "" {
		err = errors.New("missing required environmentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("environments/%s/webhooks", environmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *EnvironmentWebhookService) Delete(ctx context.Context, environmentID string, webhookID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if environmentID == "" {
		err = errors.New("missing required environmentId parameter")
		return err
	}
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return err
	}
	path := fmt.Sprintf("environments/%s/webhooks/%s", environmentID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *EnvironmentWebhookService) Rotate(ctx context.Context, environmentID string, webhookID string, opts ...option.RequestOption) (res *WebhookSubscriptionCreated, err error) {
	opts = slices.Concat(r.Options, opts)
	if environmentID == "" {
		err = errors.New("missing required environmentId parameter")
		return nil, err
	}
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("environments/%s/webhooks/%s/rotate", environmentID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

func (r *EnvironmentWebhookService) Test(ctx context.Context, environmentID string, webhookID string, opts ...option.RequestOption) (res *EnvironmentWebhookTestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if environmentID == "" {
		err = errors.New("missing required environmentId parameter")
		return nil, err
	}
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("environments/%s/webhooks/%s/test", environmentID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type WebhookEventType string

const (
	WebhookEventTypeCellCreated         WebhookEventType = "cell.created"
	WebhookEventTypeCellUpdated         WebhookEventType = "cell.updated"
	WebhookEventTypeCellDeleted         WebhookEventType = "cell.deleted"
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
	WebhookEventTypeScheduleCreated     WebhookEventType = "schedule.created"
	WebhookEventTypeScheduleDeleted     WebhookEventType = "schedule.deleted"
	WebhookEventTypeScheduleTriggered   WebhookEventType = "schedule.triggered"
	WebhookEventTypeConnectionAttached  WebhookEventType = "connection.attached"
	WebhookEventTypeConnectionDetached  WebhookEventType = "connection.detached"
	WebhookEventTypeWebhookTest         WebhookEventType = "webhook.test"
)

func (r WebhookEventType) IsKnown() bool {
	switch r {
	case WebhookEventTypeCellCreated, WebhookEventTypeCellUpdated, WebhookEventTypeCellDeleted, WebhookEventTypeThreadCreated, WebhookEventTypeThreadStatusChanged, WebhookEventTypeThreadCompleted, WebhookEventTypeThreadFailed, WebhookEventTypeTurnCreated, WebhookEventTypeTurnCompleted, WebhookEventTypeTurnFailed, WebhookEventTypeMessageCreated, WebhookEventTypeApprovalRequested, WebhookEventTypeApprovalResolved, WebhookEventTypeScheduleCreated, WebhookEventTypeScheduleDeleted, WebhookEventTypeScheduleTriggered, WebhookEventTypeConnectionAttached, WebhookEventTypeConnectionDetached, WebhookEventTypeWebhookTest:
		return true
	}
	return false
}

// Webhook subscription metadata returned by list, get, and update. The one-time
// signing secret is not returned here.
type WebhookSubscription struct {
	ID            string                  `json:"id" api:"required"`
	CreatedAt     string                  `json:"createdAt" api:"required"`
	Enabled       bool                    `json:"enabled" api:"required"`
	EnvironmentID string                  `json:"environmentId" api:"required"`
	Events        []WebhookEventType      `json:"events" api:"required"`
	UpdatedAt     string                  `json:"updatedAt" api:"required"`
	URL           string                  `json:"url" api:"required"`
	JSON          webhookSubscriptionJSON `json:"-"`
}

// webhookSubscriptionJSON contains the JSON metadata for the struct
// [WebhookSubscription]
type webhookSubscriptionJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	EnvironmentID apijson.Field
	Events        apijson.Field
	UpdatedAt     apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	ID            string             `json:"id" api:"required"`
	CreatedAt     string             `json:"createdAt" api:"required"`
	Enabled       bool               `json:"enabled" api:"required"`
	EnvironmentID string             `json:"environmentId" api:"required"`
	Events        []WebhookEventType `json:"events" api:"required"`
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
	ID            apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	EnvironmentID apijson.Field
	Events        apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	WebhookSubscriptionEventTypeCellCreated         WebhookSubscriptionEventType = "cell.created"
	WebhookSubscriptionEventTypeCellUpdated         WebhookSubscriptionEventType = "cell.updated"
	WebhookSubscriptionEventTypeCellDeleted         WebhookSubscriptionEventType = "cell.deleted"
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
	WebhookSubscriptionEventTypeScheduleCreated     WebhookSubscriptionEventType = "schedule.created"
	WebhookSubscriptionEventTypeScheduleDeleted     WebhookSubscriptionEventType = "schedule.deleted"
	WebhookSubscriptionEventTypeScheduleTriggered   WebhookSubscriptionEventType = "schedule.triggered"
	WebhookSubscriptionEventTypeConnectionAttached  WebhookSubscriptionEventType = "connection.attached"
	WebhookSubscriptionEventTypeConnectionDetached  WebhookSubscriptionEventType = "connection.detached"
)

func (r WebhookSubscriptionEventType) IsKnown() bool {
	switch r {
	case WebhookSubscriptionEventTypeCellCreated, WebhookSubscriptionEventTypeCellUpdated, WebhookSubscriptionEventTypeCellDeleted, WebhookSubscriptionEventTypeThreadCreated, WebhookSubscriptionEventTypeThreadStatusChanged, WebhookSubscriptionEventTypeThreadCompleted, WebhookSubscriptionEventTypeThreadFailed, WebhookSubscriptionEventTypeTurnCreated, WebhookSubscriptionEventTypeTurnCompleted, WebhookSubscriptionEventTypeTurnFailed, WebhookSubscriptionEventTypeMessageCreated, WebhookSubscriptionEventTypeApprovalRequested, WebhookSubscriptionEventTypeApprovalResolved, WebhookSubscriptionEventTypeScheduleCreated, WebhookSubscriptionEventTypeScheduleDeleted, WebhookSubscriptionEventTypeScheduleTriggered, WebhookSubscriptionEventTypeConnectionAttached, WebhookSubscriptionEventTypeConnectionDetached:
		return true
	}
	return false
}

type EnvironmentWebhookListResponse struct {
	Cursor        string                             `json:"cursor" api:"required,nullable"`
	HasMore       bool                               `json:"hasMore" api:"required"`
	Subscriptions []WebhookSubscription              `json:"subscriptions" api:"required"`
	JSON          environmentWebhookListResponseJSON `json:"-"`
}

// environmentWebhookListResponseJSON contains the JSON metadata for the struct
// [EnvironmentWebhookListResponse]
type environmentWebhookListResponseJSON struct {
	Cursor        apijson.Field
	HasMore       apijson.Field
	Subscriptions apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentWebhookListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentWebhookListResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentWebhookTestResponse struct {
	Error      string                             `json:"error" api:"required,nullable"`
	StatusCode int64                              `json:"statusCode" api:"required,nullable"`
	Success    bool                               `json:"success" api:"required"`
	JSON       environmentWebhookTestResponseJSON `json:"-"`
}

// environmentWebhookTestResponseJSON contains the JSON metadata for the struct
// [EnvironmentWebhookTestResponse]
type environmentWebhookTestResponseJSON struct {
	Error       apijson.Field
	StatusCode  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentWebhookTestResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentWebhookTestResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentWebhookNewParams struct {
	// HTTPS endpoint that will receive webhook deliveries.
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
	// Event names to deliver. Omit to subscribe to all non-test events.
	Events param.Field[[]WebhookSubscriptionEventType] `json:"events"`
}

func (r EnvironmentWebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentWebhookUpdateParams struct {
	// Whether deliveries are enabled for this subscription.
	Enabled param.Field[bool] `json:"enabled"`
	// Event names to deliver. Omit to subscribe to all non-test events.
	Events param.Field[[]WebhookSubscriptionEventType] `json:"events"`
	// HTTPS endpoint that will receive webhook deliveries.
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r EnvironmentWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentWebhookListParams struct {
	// Opaque pagination cursor returned by a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of items to return. Defaults to 20 and preserves parseInt
	// semantics.
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [EnvironmentWebhookListParams]'s query parameters as
// `url.Values`.
func (r EnvironmentWebhookListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
