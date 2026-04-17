// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cercago

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/matrices/cerca-go/internal/apijson"
	"github.com/matrices/cerca-go/internal/param"
	"github.com/matrices/cerca-go/internal/requestconfig"
	"github.com/matrices/cerca-go/option"
)

// CellScheduleService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellScheduleService] method instead.
type CellScheduleService struct {
	Options []option.RequestOption
}

// NewCellScheduleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCellScheduleService(opts ...option.RequestOption) (r *CellScheduleService) {
	r = &CellScheduleService{}
	r.Options = opts
	return
}

func (r *CellScheduleService) New(ctx context.Context, cellID string, body CellScheduleNewParams, opts ...option.RequestOption) (res *ScheduledThread, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/schedules", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *CellScheduleService) Update(ctx context.Context, cellID string, scheduleID string, body CellScheduleUpdateParams, opts ...option.RequestOption) (res *ScheduledThread, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	if scheduleID == "" {
		err = errors.New("missing required scheduleId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/schedules/%s", cellID, scheduleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

func (r *CellScheduleService) List(ctx context.Context, cellID string, opts ...option.RequestOption) (res *CellScheduleListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/schedules", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *CellScheduleService) Delete(ctx context.Context, cellID string, scheduleID string, opts ...option.RequestOption) (res *CellScheduleDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	if scheduleID == "" {
		err = errors.New("missing required scheduleId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/schedules/%s", cellID, scheduleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

func (r *CellScheduleService) Trigger(ctx context.Context, cellID string, scheduleID string, opts ...option.RequestOption) (res *CellScheduleTriggerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	if scheduleID == "" {
		err = errors.New("missing required scheduleId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/schedules/%s/trigger", cellID, scheduleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type ScheduledThread struct {
	ID           string                   `json:"id" api:"required"`
	CreatedAt    string                   `json:"createdAt" api:"required"`
	Cron         string                   `json:"cron" api:"required"`
	Enabled      bool                     `json:"enabled" api:"required"`
	Name         string                   `json:"name" api:"required"`
	NextAt       string                   `json:"nextAt" api:"required,nullable"`
	Prompt       string                   `json:"prompt" api:"required"`
	Timezone     string                   `json:"timezone" api:"required"`
	UpdatedAt    string                   `json:"updatedAt" api:"required"`
	Features     []ScheduledThreadFeature `json:"features"`
	Instructions string                   `json:"instructions"`
	Model        string                   `json:"model"`
	JSON         scheduledThreadJSON      `json:"-"`
}

// scheduledThreadJSON contains the JSON metadata for the struct [ScheduledThread]
type scheduledThreadJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Cron         apijson.Field
	Enabled      apijson.Field
	Name         apijson.Field
	NextAt       apijson.Field
	Prompt       apijson.Field
	Timezone     apijson.Field
	UpdatedAt    apijson.Field
	Features     apijson.Field
	Instructions apijson.Field
	Model        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScheduledThread) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduledThreadJSON) RawJSON() string {
	return r.raw
}

type ScheduledThreadFeature string

const (
	ScheduledThreadFeatureMemory       ScheduledThreadFeature = "memory"
	ScheduledThreadFeatureSandbox      ScheduledThreadFeature = "sandbox"
	ScheduledThreadFeatureWeb          ScheduledThreadFeature = "web"
	ScheduledThreadFeatureConnections  ScheduledThreadFeature = "connections"
	ScheduledThreadFeatureOAuthConnect ScheduledThreadFeature = "oauth_connect"
)

func (r ScheduledThreadFeature) IsKnown() bool {
	switch r {
	case ScheduledThreadFeatureMemory, ScheduledThreadFeatureSandbox, ScheduledThreadFeatureWeb, ScheduledThreadFeatureConnections, ScheduledThreadFeatureOAuthConnect:
		return true
	}
	return false
}

type CellScheduleListResponse struct {
	Schedules []ScheduledThread            `json:"schedules" api:"required"`
	JSON      cellScheduleListResponseJSON `json:"-"`
}

// cellScheduleListResponseJSON contains the JSON metadata for the struct
// [CellScheduleListResponse]
type cellScheduleListResponseJSON struct {
	Schedules   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellScheduleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellScheduleListResponseJSON) RawJSON() string {
	return r.raw
}

type CellScheduleDeleteResponse struct {
	Success CellScheduleDeleteResponseSuccess `json:"success" api:"required"`
	JSON    cellScheduleDeleteResponseJSON    `json:"-"`
}

// cellScheduleDeleteResponseJSON contains the JSON metadata for the struct
// [CellScheduleDeleteResponse]
type cellScheduleDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellScheduleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellScheduleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CellScheduleDeleteResponseSuccess bool

const (
	CellScheduleDeleteResponseSuccessTrue CellScheduleDeleteResponseSuccess = true
)

func (r CellScheduleDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case CellScheduleDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type CellScheduleTriggerResponse struct {
	ThreadID string                          `json:"threadId" api:"required"`
	JSON     cellScheduleTriggerResponseJSON `json:"-"`
}

// cellScheduleTriggerResponseJSON contains the JSON metadata for the struct
// [CellScheduleTriggerResponse]
type cellScheduleTriggerResponseJSON struct {
	ThreadID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellScheduleTriggerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellScheduleTriggerResponseJSON) RawJSON() string {
	return r.raw
}

type CellScheduleNewParams struct {
	Cron         param.Field[string]                         `json:"cron" api:"required"`
	Name         param.Field[string]                         `json:"name" api:"required"`
	Prompt       param.Field[string]                         `json:"prompt" api:"required"`
	Features     param.Field[[]CellScheduleNewParamsFeature] `json:"features"`
	Instructions param.Field[string]                         `json:"instructions"`
	Model        param.Field[string]                         `json:"model"`
	Timezone     param.Field[string]                         `json:"timezone"`
}

func (r CellScheduleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CellScheduleNewParamsFeature string

const (
	CellScheduleNewParamsFeatureMemory       CellScheduleNewParamsFeature = "memory"
	CellScheduleNewParamsFeatureSandbox      CellScheduleNewParamsFeature = "sandbox"
	CellScheduleNewParamsFeatureWeb          CellScheduleNewParamsFeature = "web"
	CellScheduleNewParamsFeatureConnections  CellScheduleNewParamsFeature = "connections"
	CellScheduleNewParamsFeatureOAuthConnect CellScheduleNewParamsFeature = "oauth_connect"
)

func (r CellScheduleNewParamsFeature) IsKnown() bool {
	switch r {
	case CellScheduleNewParamsFeatureMemory, CellScheduleNewParamsFeatureSandbox, CellScheduleNewParamsFeatureWeb, CellScheduleNewParamsFeatureConnections, CellScheduleNewParamsFeatureOAuthConnect:
		return true
	}
	return false
}

type CellScheduleUpdateParams struct {
	Cron         param.Field[string]                            `json:"cron"`
	Enabled      param.Field[bool]                              `json:"enabled"`
	Features     param.Field[[]CellScheduleUpdateParamsFeature] `json:"features"`
	Instructions param.Field[string]                            `json:"instructions"`
	Model        param.Field[string]                            `json:"model"`
	Name         param.Field[string]                            `json:"name"`
	Prompt       param.Field[string]                            `json:"prompt"`
	Timezone     param.Field[string]                            `json:"timezone"`
}

func (r CellScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CellScheduleUpdateParamsFeature string

const (
	CellScheduleUpdateParamsFeatureMemory       CellScheduleUpdateParamsFeature = "memory"
	CellScheduleUpdateParamsFeatureSandbox      CellScheduleUpdateParamsFeature = "sandbox"
	CellScheduleUpdateParamsFeatureWeb          CellScheduleUpdateParamsFeature = "web"
	CellScheduleUpdateParamsFeatureConnections  CellScheduleUpdateParamsFeature = "connections"
	CellScheduleUpdateParamsFeatureOAuthConnect CellScheduleUpdateParamsFeature = "oauth_connect"
)

func (r CellScheduleUpdateParamsFeature) IsKnown() bool {
	switch r {
	case CellScheduleUpdateParamsFeatureMemory, CellScheduleUpdateParamsFeatureSandbox, CellScheduleUpdateParamsFeatureWeb, CellScheduleUpdateParamsFeatureConnections, CellScheduleUpdateParamsFeatureOAuthConnect:
		return true
	}
	return false
}
