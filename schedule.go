// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cercago

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/matrices/cerca-go/internal/apijson"
	"github.com/matrices/cerca-go/internal/param"
	"github.com/matrices/cerca-go/internal/requestconfig"
	"github.com/matrices/cerca-go/option"
	"github.com/matrices/cerca-go/shared"
)

// ScheduleService contains methods and other services that help with interacting
// with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScheduleService] method instead.
type ScheduleService struct {
	Options []option.RequestOption
}

// NewScheduleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScheduleService(opts ...option.RequestOption) (r *ScheduleService) {
	r = &ScheduleService{}
	r.Options = opts
	return
}

func (r *ScheduleService) New(ctx context.Context, agentID string, body ScheduleNewParams, opts ...option.RequestOption) (res *ScheduledThread, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/schedules", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ScheduleService) Update(ctx context.Context, agentID string, scheduleID string, body ScheduleUpdateParams, opts ...option.RequestOption) (res *ScheduledThread, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if scheduleID == "" {
		err = errors.New("missing required scheduleId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/schedules/%s", agentID, scheduleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

func (r *ScheduleService) List(ctx context.Context, agentID string, opts ...option.RequestOption) (res *ScheduleListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/schedules", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *ScheduleService) Delete(ctx context.Context, agentID string, scheduleID string, opts ...option.RequestOption) (res *ScheduleDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if scheduleID == "" {
		err = errors.New("missing required scheduleId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/schedules/%s", agentID, scheduleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

func (r *ScheduleService) Trigger(ctx context.Context, agentID string, scheduleID string, opts ...option.RequestOption) (res *ScheduleTriggerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if scheduleID == "" {
		err = errors.New("missing required scheduleId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/schedules/%s/trigger", agentID, scheduleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type ScheduledThread struct {
	ID           string              `json:"id" api:"required"`
	CreatedAt    string              `json:"createdAt" api:"required"`
	Enabled      bool                `json:"enabled" api:"required"`
	Name         string              `json:"name" api:"required"`
	NextAt       string              `json:"nextAt" api:"required,nullable"`
	Prompt       string              `json:"prompt" api:"required"`
	Timezone     string              `json:"timezone" api:"required"`
	UpdatedAt    string              `json:"updatedAt" api:"required"`
	Cron         string              `json:"cron"`
	Instructions string              `json:"instructions"`
	Model        string              `json:"model"`
	RunAt        time.Time           `json:"runAt" format:"date-time"`
	Tools        []shared.ToolSpec   `json:"tools"`
	JSON         scheduledThreadJSON `json:"-"`
}

// scheduledThreadJSON contains the JSON metadata for the struct [ScheduledThread]
type scheduledThreadJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Enabled      apijson.Field
	Name         apijson.Field
	NextAt       apijson.Field
	Prompt       apijson.Field
	Timezone     apijson.Field
	UpdatedAt    apijson.Field
	Cron         apijson.Field
	Instructions apijson.Field
	Model        apijson.Field
	RunAt        apijson.Field
	Tools        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScheduledThread) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduledThreadJSON) RawJSON() string {
	return r.raw
}

type ScheduleListResponse struct {
	Schedules []ScheduledThread        `json:"schedules" api:"required"`
	JSON      scheduleListResponseJSON `json:"-"`
}

// scheduleListResponseJSON contains the JSON metadata for the struct
// [ScheduleListResponse]
type scheduleListResponseJSON struct {
	Schedules   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleListResponseJSON) RawJSON() string {
	return r.raw
}

type ScheduleDeleteResponse struct {
	Success ScheduleDeleteResponseSuccess `json:"success" api:"required"`
	JSON    scheduleDeleteResponseJSON    `json:"-"`
}

// scheduleDeleteResponseJSON contains the JSON metadata for the struct
// [ScheduleDeleteResponse]
type scheduleDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ScheduleDeleteResponseSuccess bool

const (
	ScheduleDeleteResponseSuccessTrue ScheduleDeleteResponseSuccess = true
)

func (r ScheduleDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ScheduleDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ScheduleTriggerResponse struct {
	ThreadID string                      `json:"threadId" api:"required"`
	JSON     scheduleTriggerResponseJSON `json:"-"`
}

// scheduleTriggerResponseJSON contains the JSON metadata for the struct
// [ScheduleTriggerResponse]
type scheduleTriggerResponseJSON struct {
	ThreadID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleTriggerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleTriggerResponseJSON) RawJSON() string {
	return r.raw
}

type ScheduleNewParams struct {
	Name         param.Field[string]                 `json:"name" api:"required"`
	Prompt       param.Field[string]                 `json:"prompt" api:"required"`
	Cron         param.Field[string]                 `json:"cron"`
	Instructions param.Field[string]                 `json:"instructions"`
	Model        param.Field[string]                 `json:"model"`
	RunAt        param.Field[time.Time]              `json:"runAt" format:"date-time"`
	Timezone     param.Field[string]                 `json:"timezone"`
	Tools        param.Field[[]shared.ToolSpecParam] `json:"tools"`
}

func (r ScheduleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScheduleUpdateParams struct {
	Cron         param.Field[string]                 `json:"cron"`
	Enabled      param.Field[bool]                   `json:"enabled"`
	Instructions param.Field[string]                 `json:"instructions"`
	Model        param.Field[string]                 `json:"model"`
	Name         param.Field[string]                 `json:"name"`
	Prompt       param.Field[string]                 `json:"prompt"`
	RunAt        param.Field[time.Time]              `json:"runAt" format:"date-time"`
	Timezone     param.Field[string]                 `json:"timezone"`
	Tools        param.Field[[]shared.ToolSpecParam] `json:"tools"`
}

func (r ScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
