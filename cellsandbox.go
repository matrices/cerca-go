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

// CellSandboxService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellSandboxService] method instead.
type CellSandboxService struct {
	Options []option.RequestOption
}

// NewCellSandboxService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCellSandboxService(opts ...option.RequestOption) (r *CellSandboxService) {
	r = &CellSandboxService{}
	r.Options = opts
	return
}

func (r *CellSandboxService) Exec(ctx context.Context, cellID string, body CellSandboxExecParams, opts ...option.RequestOption) (res *SandboxExecResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/sandbox/exec", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *CellSandboxService) Read(ctx context.Context, cellID string, body CellSandboxReadParams, opts ...option.RequestOption) (res *SandboxReadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/sandbox/read", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *CellSandboxService) Write(ctx context.Context, cellID string, body CellSandboxWriteParams, opts ...option.RequestOption) (res *CellSandboxWriteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/sandbox/write", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SandboxExecResult struct {
	ExitCode     float64                `json:"exitCode" api:"required,nullable"`
	Handle       string                 `json:"handle" api:"required"`
	State        SandboxExecResultState `json:"state" api:"required"`
	Stderr       string                 `json:"stderr" api:"required"`
	StderrOffset float64                `json:"stderrOffset" api:"required"`
	Stdout       string                 `json:"stdout" api:"required"`
	StdoutOffset float64                `json:"stdoutOffset" api:"required"`
	JSON         sandboxExecResultJSON  `json:"-"`
}

// sandboxExecResultJSON contains the JSON metadata for the struct
// [SandboxExecResult]
type sandboxExecResultJSON struct {
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

func (r *SandboxExecResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxExecResultJSON) RawJSON() string {
	return r.raw
}

type SandboxExecResultState string

const (
	SandboxExecResultStateRunning SandboxExecResultState = "running"
	SandboxExecResultStateExited  SandboxExecResultState = "exited"
)

func (r SandboxExecResultState) IsKnown() bool {
	switch r {
	case SandboxExecResultStateRunning, SandboxExecResultStateExited:
		return true
	}
	return false
}

type SandboxReadResponse struct {
	BytesRead float64                 `json:"bytesRead" api:"required"`
	Content   string                  `json:"content" api:"required"`
	JSON      sandboxReadResponseJSON `json:"-"`
}

// sandboxReadResponseJSON contains the JSON metadata for the struct
// [SandboxReadResponse]
type sandboxReadResponseJSON struct {
	BytesRead   apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SandboxReadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxReadResponseJSON) RawJSON() string {
	return r.raw
}

type CellSandboxWriteResponse struct {
	Success CellSandboxWriteResponseSuccess `json:"success" api:"required"`
	JSON    cellSandboxWriteResponseJSON    `json:"-"`
}

// cellSandboxWriteResponseJSON contains the JSON metadata for the struct
// [CellSandboxWriteResponse]
type cellSandboxWriteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CellSandboxWriteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cellSandboxWriteResponseJSON) RawJSON() string {
	return r.raw
}

type CellSandboxWriteResponseSuccess bool

const (
	CellSandboxWriteResponseSuccessTrue CellSandboxWriteResponseSuccess = true
)

func (r CellSandboxWriteResponseSuccess) IsKnown() bool {
	switch r {
	case CellSandboxWriteResponseSuccessTrue:
		return true
	}
	return false
}

type CellSandboxExecParams struct {
	Command   param.Field[string]  `json:"command" api:"required"`
	MaxBuffer param.Field[float64] `json:"maxBuffer"`
	// Timeout in seconds. Runtime converts this to milliseconds.
	Timeout param.Field[float64] `json:"timeout"`
	// Optional sandbox working directory.
	Workdir param.Field[string] `json:"workdir"`
}

func (r CellSandboxExecParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CellSandboxReadParams struct {
	Path   param.Field[string]  `json:"path" api:"required"`
	Limit  param.Field[float64] `json:"limit"`
	Offset param.Field[float64] `json:"offset"`
}

func (r CellSandboxReadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CellSandboxWriteParams struct {
	Content param.Field[string] `json:"content" api:"required"`
	Path    param.Field[string] `json:"path" api:"required"`
}

func (r CellSandboxWriteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
