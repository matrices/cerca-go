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

// ConnectionService contains methods and other services that help with interacting
// with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectionService] method instead.
type ConnectionService struct {
	Options []option.RequestOption
}

// NewConnectionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectionService(opts ...option.RequestOption) (r *ConnectionService) {
	r = &ConnectionService{}
	r.Options = opts
	return
}

func (r *ConnectionService) List(ctx context.Context, agentID string, opts ...option.RequestOption) (res *ConnectionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/connections", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *ConnectionService) Attach(ctx context.Context, agentID string, body ConnectionAttachParams, opts ...option.RequestOption) (res *AttachedConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/connections", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ConnectionService) Detach(ctx context.Context, agentID string, connectionID string, opts ...option.RequestOption) (res *ConnectionDetachResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if connectionID == "" {
		err = errors.New("missing required connectionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/connections/%s", agentID, connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type AttachedConnection struct {
	AttachedAt string                 `json:"attachedAt" api:"required"`
	Metadata   ToolConnectionMetadata `json:"metadata"`
	JSON       attachedConnectionJSON `json:"-"`
	Connection
}

// attachedConnectionJSON contains the JSON metadata for the struct
// [AttachedConnection]
type attachedConnectionJSON struct {
	AttachedAt  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttachedConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attachedConnectionJSON) RawJSON() string {
	return r.raw
}

type ToolConnectionMetadata map[string]string

type ToolConnectionMetadataParam map[string]string

type ConnectionListResponse struct {
	Connections []AttachedConnection       `json:"connections" api:"required"`
	Cursor      string                     `json:"cursor" api:"required,nullable"`
	HasMore     bool                       `json:"hasMore" api:"required"`
	JSON        connectionListResponseJSON `json:"-"`
}

// connectionListResponseJSON contains the JSON metadata for the struct
// [ConnectionListResponse]
type connectionListResponseJSON struct {
	Connections apijson.Field
	Cursor      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionListResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectionDetachResponse struct {
	Success ConnectionDetachResponseSuccess `json:"success" api:"required"`
	JSON    connectionDetachResponseJSON    `json:"-"`
}

// connectionDetachResponseJSON contains the JSON metadata for the struct
// [ConnectionDetachResponse]
type connectionDetachResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionDetachResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionDetachResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectionDetachResponseSuccess bool

const (
	ConnectionDetachResponseSuccessTrue ConnectionDetachResponseSuccess = true
)

func (r ConnectionDetachResponseSuccess) IsKnown() bool {
	switch r {
	case ConnectionDetachResponseSuccessTrue:
		return true
	}
	return false
}

type ConnectionAttachParams struct {
	ConnectionID param.Field[string]                      `json:"connectionId" api:"required"`
	Metadata     param.Field[ToolConnectionMetadataParam] `json:"metadata"`
}

func (r ConnectionAttachParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
