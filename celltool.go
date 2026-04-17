// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cercago

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/matrices/cerca-go/internal/apijson"
	"github.com/matrices/cerca-go/internal/requestconfig"
	"github.com/matrices/cerca-go/option"
)

// CellToolService contains methods and other services that help with interacting
// with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellToolService] method instead.
type CellToolService struct {
	Options []option.RequestOption
}

// NewCellToolService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCellToolService(opts ...option.RequestOption) (r *CellToolService) {
	r = &CellToolService{}
	r.Options = opts
	return
}

func (r *CellToolService) List(ctx context.Context, cellID string, opts ...option.RequestOption) (res *CellToolListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/tools", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type RuntimeDiscoveredTool struct {
	Approval    RuntimeDiscoveredToolApproval `json:"approval" api:"required"`
	Description string                        `json:"description" api:"required"`
	// JSON Schema object describing tool input parameters.
	InputSchema        interface{}                 `json:"inputSchema" api:"required"`
	Name               string                      `json:"name" api:"required"`
	Origin             RuntimeDiscoveredToolOrigin `json:"origin" api:"required"`
	Source             string                      `json:"source" api:"required"`
	AccountLabel       string                      `json:"accountLabel"`
	ConnectionID       string                      `json:"connectionId"`
	ConnectionMetadata ToolConnectionMetadata      `json:"connectionMetadata"`
	SourceID           string                      `json:"sourceId"`
	SourceVersion      float64                     `json:"sourceVersion"`
	JSON               runtimeDiscoveredToolJSON   `json:"-"`
}

// runtimeDiscoveredToolJSON contains the JSON metadata for the struct
// [RuntimeDiscoveredTool]
type runtimeDiscoveredToolJSON struct {
	Approval           apijson.Field
	Description        apijson.Field
	InputSchema        apijson.Field
	Name               apijson.Field
	Origin             apijson.Field
	Source             apijson.Field
	AccountLabel       apijson.Field
	ConnectionID       apijson.Field
	ConnectionMetadata apijson.Field
	SourceID           apijson.Field
	SourceVersion      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RuntimeDiscoveredTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeDiscoveredToolJSON) RawJSON() string {
	return r.raw
}

type RuntimeDiscoveredToolApproval string

const (
	RuntimeDiscoveredToolApprovalAlways RuntimeDiscoveredToolApproval = "always"
	RuntimeDiscoveredToolApprovalNever  RuntimeDiscoveredToolApproval = "never"
)

func (r RuntimeDiscoveredToolApproval) IsKnown() bool {
	switch r {
	case RuntimeDiscoveredToolApprovalAlways, RuntimeDiscoveredToolApprovalNever:
		return true
	}
	return false
}

type RuntimeDiscoveredToolOrigin string

const (
	RuntimeDiscoveredToolOriginBuiltin    RuntimeDiscoveredToolOrigin = "builtin"
	RuntimeDiscoveredToolOriginToolSource RuntimeDiscoveredToolOrigin = "tool_source"
)

func (r RuntimeDiscoveredToolOrigin) IsKnown() bool {
	switch r {
	case RuntimeDiscoveredToolOriginBuiltin, RuntimeDiscoveredToolOriginToolSource:
		return true
	}
	return false
}

type RuntimeSourceWarning struct {
	Message  string                   `json:"message" api:"required"`
	Source   string                   `json:"source" api:"required"`
	SourceID string                   `json:"sourceId"`
	JSON     runtimeSourceWarningJSON `json:"-"`
}

// runtimeSourceWarningJSON contains the JSON metadata for the struct
// [RuntimeSourceWarning]
type runtimeSourceWarningJSON struct {
	Message     apijson.Field
	Source      apijson.Field
	SourceID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeSourceWarning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeSourceWarningJSON) RawJSON() string {
	return r.raw
}

type RuntimeToolDescriptor struct {
	Description string                    `json:"description" api:"required"`
	InputSchema interface{}               `json:"inputSchema" api:"required"`
	Name        RuntimeToolName           `json:"name" api:"required"`
	JSON        runtimeToolDescriptorJSON `json:"-"`
}

// runtimeToolDescriptorJSON contains the JSON metadata for the struct
// [RuntimeToolDescriptor]
type runtimeToolDescriptorJSON struct {
	Description apijson.Field
	InputSchema apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeToolDescriptor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeToolDescriptorJSON) RawJSON() string {
	return r.raw
}

type RuntimeToolName string

const (
	RuntimeToolNameWebSearch        RuntimeToolName = "web_search"
	RuntimeToolNameWebFetch         RuntimeToolName = "web_fetch"
	RuntimeToolNameSandboxExec      RuntimeToolName = "sandbox_exec"
	RuntimeToolNameSandboxRead      RuntimeToolName = "sandbox_read"
	RuntimeToolNameSandboxWriteFile RuntimeToolName = "sandbox_write_file"
	RuntimeToolNameSandboxReadFile  RuntimeToolName = "sandbox_read_file"
	RuntimeToolNameCellDB           RuntimeToolName = "cell_db"
	RuntimeToolNameContextRead      RuntimeToolName = "context_read"
	RuntimeToolNameContextList      RuntimeToolName = "context_list"
	RuntimeToolNameContextSearch    RuntimeToolName = "context_search"
	RuntimeToolNameBlobSync         RuntimeToolName = "blob_sync"
	RuntimeToolNameContextSync      RuntimeToolName = "context_sync"
	RuntimeToolNameContextWrite     RuntimeToolName = "context_write"
	RuntimeToolNameContextDelete    RuntimeToolName = "context_delete"
	RuntimeToolNameToolDiscover     RuntimeToolName = "tool_discover"
	RuntimeToolNameToolCall         RuntimeToolName = "tool_call"
	RuntimeToolNameArtifactRead     RuntimeToolName = "artifact_read"
	RuntimeToolNameArtifactSync     RuntimeToolName = "artifact_sync"
	RuntimeToolNameToolConnect      RuntimeToolName = "tool_connect"
	RuntimeToolNameSubThread        RuntimeToolName = "sub_thread"
	RuntimeToolNameWait             RuntimeToolName = "wait"
	RuntimeToolNameGetCurrentTime   RuntimeToolName = "get_current_time"
)

func (r RuntimeToolName) IsKnown() bool {
	switch r {
	case RuntimeToolNameWebSearch, RuntimeToolNameWebFetch, RuntimeToolNameSandboxExec, RuntimeToolNameSandboxRead, RuntimeToolNameSandboxWriteFile, RuntimeToolNameSandboxReadFile, RuntimeToolNameCellDB, RuntimeToolNameContextRead, RuntimeToolNameContextList, RuntimeToolNameContextSearch, RuntimeToolNameBlobSync, RuntimeToolNameContextSync, RuntimeToolNameContextWrite, RuntimeToolNameContextDelete, RuntimeToolNameToolDiscover, RuntimeToolNameToolCall, RuntimeToolNameArtifactRead, RuntimeToolNameArtifactSync, RuntimeToolNameToolConnect, RuntimeToolNameSubThread, RuntimeToolNameWait, RuntimeToolNameGetCurrentTime:
		return true
	}
	return false
}

type CellToolListResponse struct {
	DiscoveredTools []RuntimeDiscoveredTool  `json:"discoveredTools" api:"required"`
	Tools           []RuntimeToolDescriptor  `json:"tools" api:"required"`
	SourceWarnings  []RuntimeSourceWarning   `json:"sourceWarnings"`
	JSON            cellToolListResponseJSON `json:"-"`
}

// cellToolListResponseJSON contains the JSON metadata for the struct
// [CellToolListResponse]
type cellToolListResponseJSON struct {
	DiscoveredTools apijson.Field
	Tools           apijson.Field
	SourceWarnings  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CellToolListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellToolListResponseJSON) RawJSON() string {
	return r.raw
}
