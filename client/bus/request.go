package bus

import (
	"encoding/json"
)

// Start listening to shell logger:
//	{"type": "startListening", "time": 1525212510}
type StartListeningRequest struct {
	Type string `json:"type"`
	Time int `json:"time"`
}

const startListeningType = "startListening"

func startListeningRequest(time int) []byte {
	bytes, _ := json.Marshal(&StartListeningRequest{startListeningType, time})

	return bytes
}

// Stop listening to shell logger:
//	{"command": "ls", "type": "stopListening", "returnCode": 2, "time": 1525212510}
type StopListeningRequest struct {
	Type       string `json:"type"`
	Command    string `json:"command"`
	ReturnCode int    `json:"returnCode"`
	Time       int    `json:"time"`
}

const stopListeningType = "stopListening"

func stopListeningRequest(command string, returnCode int, time int) []byte {
	bytes, _ := json.Marshal(&StopListeningRequest{stopListeningType, command, returnCode, time})

	return bytes
}
