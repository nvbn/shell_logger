package bus

import (
	"encoding/json"
	"github.com/nvbn/shell_logger/wrapper/storage"
)

// Base response:
//	{"status": "everything"}
type Response struct {
	Status string `json:"status"`
}

func response(status string) []byte {
	bytes, _ := json.Marshal(&Response{status})

	return bytes
}

// OK response:
//	{"status": "ok"}
var okResponse = response("ok")

// Error response:
//	{"status": "error", "error": "everything"}
type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func errorResponse(err error) []byte {
	bytes, _ := json.Marshal(&ErrorResponse{"error", err.Error()})

	return bytes
}

// Commands list response:
//	{"status":"ok","commands":[{"command":"ls","output":"​","returnCode":2}]}
type ListResponse struct {
	Status   string             `json:"status"`
	Commands []*storage.Command `json:"commands"`
}

func listResponse(commands []*storage.Command) []byte {
	bytes, _ := json.Marshal(&ListResponse{"ok", commands})

	return bytes
}
