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

// SandboxService contains methods and other services that help with interacting
// with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxService] method instead.
type SandboxService struct {
	Options []option.RequestOption
}

// NewSandboxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSandboxService(opts ...option.RequestOption) (r *SandboxService) {
	r = &SandboxService{}
	r.Options = opts
	return
}

func (r *SandboxService) Exec(ctx context.Context, agentID string, body SandboxExecParams, opts ...option.RequestOption) (res *ExecResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/sandbox/exec", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *SandboxService) Read(ctx context.Context, agentID string, body SandboxReadParams, opts ...option.RequestOption) (res *ReadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/sandbox/read", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *SandboxService) Write(ctx context.Context, agentID string, body SandboxWriteParams, opts ...option.RequestOption) (res *SandboxWriteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/sandbox/write", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ExecResult struct {
	ExitCode     float64         `json:"exitCode" api:"required,nullable"`
	Handle       string          `json:"handle" api:"required"`
	State        ExecResultState `json:"state" api:"required"`
	Stderr       string          `json:"stderr" api:"required"`
	StderrOffset float64         `json:"stderrOffset" api:"required"`
	Stdout       string          `json:"stdout" api:"required"`
	StdoutOffset float64         `json:"stdoutOffset" api:"required"`
	JSON         execResultJSON  `json:"-"`
}

// execResultJSON contains the JSON metadata for the struct [ExecResult]
type execResultJSON struct {
	ExitCode     apijson.Field
	Handle       apijson.Field
	State        apijson.Field
	Stderr       apijson.Field
	StderrOffset apijson.Field
	Stdout       apijson.Field
	StdoutOffset apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ExecResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r execResultJSON) RawJSON() string {
	return r.raw
}

type ExecResultState string

const (
	ExecResultStateRunning ExecResultState = "running"
	ExecResultStateExited  ExecResultState = "exited"
)

func (r ExecResultState) IsKnown() bool {
	switch r {
	case ExecResultStateRunning, ExecResultStateExited:
		return true
	}
	return false
}

type ReadResponse struct {
	BytesRead float64          `json:"bytesRead" api:"required"`
	Content   string           `json:"content" api:"required"`
	JSON      readResponseJSON `json:"-"`
}

// readResponseJSON contains the JSON metadata for the struct [ReadResponse]
type readResponseJSON struct {
	BytesRead   apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r readResponseJSON) RawJSON() string {
	return r.raw
}

type SandboxWriteResponse struct {
	Success SandboxWriteResponseSuccess `json:"success" api:"required"`
	JSON    sandboxWriteResponseJSON    `json:"-"`
}

// sandboxWriteResponseJSON contains the JSON metadata for the struct
// [SandboxWriteResponse]
type sandboxWriteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SandboxWriteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxWriteResponseJSON) RawJSON() string {
	return r.raw
}

type SandboxWriteResponseSuccess bool

const (
	SandboxWriteResponseSuccessTrue SandboxWriteResponseSuccess = true
)

func (r SandboxWriteResponseSuccess) IsKnown() bool {
	switch r {
	case SandboxWriteResponseSuccessTrue:
		return true
	}
	return false
}

type SandboxExecParams struct {
	Command   param.Field[string]  `json:"command" api:"required"`
	MaxBuffer param.Field[float64] `json:"maxBuffer"`
	// Timeout in seconds. Runtime converts this to milliseconds.
	ExecutionTimeout param.Field[float64] `json:"timeout"`
	// Optional sandbox working directory.
	Workdir param.Field[string] `json:"workdir"`
}

func (r SandboxExecParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxReadParams struct {
	Path   param.Field[string]  `json:"path" api:"required"`
	Limit  param.Field[float64] `json:"limit"`
	Offset param.Field[float64] `json:"offset"`
}

func (r SandboxReadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxWriteParams struct {
	Content param.Field[string] `json:"content" api:"required"`
	Path    param.Field[string] `json:"path" api:"required"`
}

func (r SandboxWriteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
