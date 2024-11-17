package c2api

import (
	"fyne.io/fyne/v2"
)

type HttpScan struct {
	ID       string `json:"id"`
	ReqID    int64  `json:"reqId"`
	AgentIDs string `json:"agentIds"`
}

type HttpScanTask struct {
	ScanID      string `json:"scanId"`
	AgentID     string `json:"agentId"`
	ReqID       int64  `json:"reqId"`
	Payload     string `json:"payload"`
	RawRequest  string `json:"rawRequest"`
	RawResponse string `json:"rawResponse"`
}

type Attack struct {
	ID       string `json:"id"`
	AgentID  string `json:"agentId"`
	TargetID string `json:"targetId"`
	Comment  string `json:"comment"`
}

type Agent struct {
	ID        string `json:"id"`
	Hostname  string `json:"hostname"`
	PubKeyPEM string `json:"pubKeyPem"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
	IpAddress string `json:"ipAddress"`
}

type Message struct {
	ID                  string `json:"id"`
	AgentID             string `json:"agentId"`
	PipelineExecutionID string `json:"pipelineExecutionId"`
	FriendlyTitle       string `json:"friendlyTitle"`
	Request             string `json:"request"`
	Response            string `json:"response"`
	CreatedAt           int64  `json:"createdAt"`
}

type Pipeline struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"description"`
	Category string `json:"category"`
	Settings string `json:"settings"` // JSON stringified "PipelineSettings" variable.
}

type PipelineRun struct {
	ID               string `json:"id"`
	PipelineID       string `json:"pipelineId"`
	FinishedPipeline string `json:"finishedPipeline"` // JSON stringified "Pipeline" variable.
	FinishedAt       int64  `json:"finishedAt"`
	FinishedAtLabel  string `json:"finishedAtLabel"`
}

type PipelineSettings struct {
	Input map[string]string       `json:"input"`
	Steps map[string]PipelineStep `json:"steps"`
}

type PipelineStep struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	Position fyne.Position `json:"position"`
	Tool     Tool          `json:"tool"`
	LinkedTo []string      `json:"linkedTo"` // ID of a steps it is linked towards.
}

type Tool struct {
	ID               string                `json:"id"`
	Name             string                `json:"name"`
	Description      string                `json:"description"`
	ToolCategoryName string                `json:"toolCategoryName"`
	Inputs           map[string]ToolInput  `json:"inputs"`
	Outputs          map[string]ToolOutput `json:"outputs"`
}

type ToolInput struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Value       string `json:"value"`
}

type ToolOutput struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type FileRecord struct {
	Name         string `json:"name"`
	AbsolutePath string `json:"absolutePath"`
	IsDir        bool   `json:"isDir"`
}

type FileBrowserCtx struct {
	WorkingDir string       `json:"wd"`
	Records    []FileRecord `json:"records"`
	Err        string       `json:"err"`
}

// Refers to ToolInput.Type and ToolOutput.Type
const TOOL_IO_TYPE_STRING = "STRING"
const TOOL_IO_TYPE_FILE = "FILE"

type Target struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type File struct {
	ID                string `json:"id"`
	UploadedByAgentID string `json:"uploadedByAgentId"`
	OriginalName      string `json:"originalName"`
	StorageName       string `json:"storageName"`
	Uploaded          bool   `json:"uploaded"`
	UploadedAt        string `json:"uploadedAt"`
}
