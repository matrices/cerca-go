// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type ApprovalMode string

const (
	ApprovalModeAlways ApprovalMode = "always"
	ApprovalModeNever  ApprovalMode = "never"
)

func (r ApprovalMode) IsKnown() bool {
	switch r {
	case ApprovalModeAlways, ApprovalModeNever:
		return true
	}
	return false
}

type ToolName string

const (
	ToolNameGetTime                   ToolName = "get_time"
	ToolNameSubThread                 ToolName = "sub_thread"
	ToolNameWait                      ToolName = "wait"
	ToolNameToolCall                  ToolName = "tool.call"
	ToolNameToolDiscover              ToolName = "tool.discover"
	ToolNameArtifactRead              ToolName = "artifact.read"
	ToolNameAgentThreadsList          ToolName = "agent.threads.list"
	ToolNameAgentThreadsGet           ToolName = "agent.threads.get"
	ToolNameAgentApprovalsCancel      ToolName = "agent.approvals.cancel"
	ToolNameAgentApprovalsGrantThread ToolName = "agent.approvals.grant_thread"
	ToolNameAgentApprovalsGrantAgent  ToolName = "agent.approvals.grant_agent"
	ToolNameAgentCreate               ToolName = "agent.create"
	ToolNameAgentApprovalsUpdate      ToolName = "agent.approvals.update"
	ToolNameToolConnect               ToolName = "tool.connect"
	ToolNameSandboxExec               ToolName = "sandbox.exec"
	ToolNameSandboxRead               ToolName = "sandbox.read"
	ToolNameSandboxWriteFile          ToolName = "sandbox.write_file"
	ToolNameSandboxReadFile           ToolName = "sandbox.read_file"
	ToolNameSandboxSpawn              ToolName = "sandbox.spawn"
	ToolNameSandboxStdin              ToolName = "sandbox.stdin"
	ToolNameSandboxSessionRead        ToolName = "sandbox.session_read"
	ToolNameSandboxKill               ToolName = "sandbox.kill"
	ToolNameSandboxListSessions       ToolName = "sandbox.list_sessions"
	ToolNameSandboxSyncArtifact       ToolName = "sandbox.sync_artifact"
	ToolNameMemoryRead                ToolName = "memory.read"
	ToolNameMemoryList                ToolName = "memory.list"
	ToolNameMemorySearch              ToolName = "memory.search"
	ToolNameMemoryWrite               ToolName = "memory.write"
	ToolNameMemoryDelete              ToolName = "memory.delete"
	ToolNameDBQuery                   ToolName = "db.query"
	ToolNameWebSearch                 ToolName = "web.search"
	ToolNameWebFetch                  ToolName = "web.fetch"
	ToolNameAgentScheduleList         ToolName = "agent.schedule.list"
	ToolNameAgentScheduleCreate       ToolName = "agent.schedule.create"
	ToolNameAgentScheduleUpdate       ToolName = "agent.schedule.update"
	ToolNameAgentScheduleDelete       ToolName = "agent.schedule.delete"
	ToolNameAgentScheduleTrigger      ToolName = "agent.schedule.trigger"
)

func (r ToolName) IsKnown() bool {
	switch r {
	case ToolNameGetTime, ToolNameSubThread, ToolNameWait, ToolNameToolCall, ToolNameToolDiscover, ToolNameArtifactRead, ToolNameAgentThreadsList, ToolNameAgentThreadsGet, ToolNameAgentApprovalsCancel, ToolNameAgentApprovalsGrantThread, ToolNameAgentApprovalsGrantAgent, ToolNameAgentCreate, ToolNameAgentApprovalsUpdate, ToolNameToolConnect, ToolNameSandboxExec, ToolNameSandboxRead, ToolNameSandboxWriteFile, ToolNameSandboxReadFile, ToolNameSandboxSpawn, ToolNameSandboxStdin, ToolNameSandboxSessionRead, ToolNameSandboxKill, ToolNameSandboxListSessions, ToolNameSandboxSyncArtifact, ToolNameMemoryRead, ToolNameMemoryList, ToolNameMemorySearch, ToolNameMemoryWrite, ToolNameMemoryDelete, ToolNameDBQuery, ToolNameWebSearch, ToolNameWebFetch, ToolNameAgentScheduleList, ToolNameAgentScheduleCreate, ToolNameAgentScheduleUpdate, ToolNameAgentScheduleDelete, ToolNameAgentScheduleTrigger:
		return true
	}
	return false
}

type ToolSpec = string

type ToolSpecParam = string
