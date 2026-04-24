// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cercago

import (
	"github.com/matrices/cerca-go/internal/apierror"
	"github.com/matrices/cerca-go/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type ApprovalMode = shared.ApprovalMode

// This is an alias to an internal value.
const ApprovalModeAlways = shared.ApprovalModeAlways

// This is an alias to an internal value.
const ApprovalModeNever = shared.ApprovalModeNever

// This is an alias to an internal type.
type ErrorResponse = shared.ErrorResponse

// This is an alias to an internal type.
type JsonObject = shared.JsonObject

// This is an alias to an internal type.
type ToolName = shared.ToolName

// This is an alias to an internal value.
const ToolNameGetTime = shared.ToolNameGetTime

// This is an alias to an internal value.
const ToolNameSubThread = shared.ToolNameSubThread

// This is an alias to an internal value.
const ToolNameWait = shared.ToolNameWait

// This is an alias to an internal value.
const ToolNameToolCall = shared.ToolNameToolCall

// This is an alias to an internal value.
const ToolNameToolDiscover = shared.ToolNameToolDiscover

// This is an alias to an internal value.
const ToolNameArtifactRead = shared.ToolNameArtifactRead

// This is an alias to an internal value.
const ToolNameAgentThreadsList = shared.ToolNameAgentThreadsList

// This is an alias to an internal value.
const ToolNameAgentThreadsGet = shared.ToolNameAgentThreadsGet

// This is an alias to an internal value.
const ToolNameAgentApprovalsCancel = shared.ToolNameAgentApprovalsCancel

// This is an alias to an internal value.
const ToolNameAgentApprovalsGrantThread = shared.ToolNameAgentApprovalsGrantThread

// This is an alias to an internal value.
const ToolNameAgentApprovalsGrantAgent = shared.ToolNameAgentApprovalsGrantAgent

// This is an alias to an internal value.
const ToolNameAgentCreate = shared.ToolNameAgentCreate

// This is an alias to an internal value.
const ToolNameAgentApprovalsUpdate = shared.ToolNameAgentApprovalsUpdate

// This is an alias to an internal value.
const ToolNameToolConnect = shared.ToolNameToolConnect

// This is an alias to an internal value.
const ToolNameSandboxExec = shared.ToolNameSandboxExec

// This is an alias to an internal value.
const ToolNameSandboxRead = shared.ToolNameSandboxRead

// This is an alias to an internal value.
const ToolNameSandboxWriteFile = shared.ToolNameSandboxWriteFile

// This is an alias to an internal value.
const ToolNameSandboxReadFile = shared.ToolNameSandboxReadFile

// This is an alias to an internal value.
const ToolNameSandboxSpawn = shared.ToolNameSandboxSpawn

// This is an alias to an internal value.
const ToolNameSandboxStdin = shared.ToolNameSandboxStdin

// This is an alias to an internal value.
const ToolNameSandboxSessionRead = shared.ToolNameSandboxSessionRead

// This is an alias to an internal value.
const ToolNameSandboxKill = shared.ToolNameSandboxKill

// This is an alias to an internal value.
const ToolNameSandboxListSessions = shared.ToolNameSandboxListSessions

// This is an alias to an internal value.
const ToolNameSandboxSyncArtifact = shared.ToolNameSandboxSyncArtifact

// This is an alias to an internal value.
const ToolNameMemoryRead = shared.ToolNameMemoryRead

// This is an alias to an internal value.
const ToolNameMemoryList = shared.ToolNameMemoryList

// This is an alias to an internal value.
const ToolNameMemorySearch = shared.ToolNameMemorySearch

// This is an alias to an internal value.
const ToolNameMemoryWrite = shared.ToolNameMemoryWrite

// This is an alias to an internal value.
const ToolNameMemoryDelete = shared.ToolNameMemoryDelete

// This is an alias to an internal value.
const ToolNameDBQuery = shared.ToolNameDBQuery

// This is an alias to an internal value.
const ToolNameWebSearch = shared.ToolNameWebSearch

// This is an alias to an internal value.
const ToolNameWebFetch = shared.ToolNameWebFetch

// This is an alias to an internal value.
const ToolNameAgentScheduleList = shared.ToolNameAgentScheduleList

// This is an alias to an internal value.
const ToolNameAgentScheduleCreate = shared.ToolNameAgentScheduleCreate

// This is an alias to an internal value.
const ToolNameAgentScheduleUpdate = shared.ToolNameAgentScheduleUpdate

// This is an alias to an internal value.
const ToolNameAgentScheduleDelete = shared.ToolNameAgentScheduleDelete

// This is an alias to an internal value.
const ToolNameAgentScheduleTrigger = shared.ToolNameAgentScheduleTrigger

// Configurable runtime tool name, configurable namespace wildcard such as
// sandbox._, or external namespace wildcard such as google._.
//
// This is an alias to an internal type.
type ToolSpec = shared.ToolSpec

// Configurable runtime tool name, configurable namespace wildcard such as
// sandbox._, or external namespace wildcard such as google._.
//
// This is an alias to an internal type.
type ToolSpecParam = shared.ToolSpecParam
