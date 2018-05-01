package bus

import (
	"encoding/json"
	"errors"
)

// Start listening to shell logger:
//	{"type": "startListening", "time": 1525212510}
type StartListeningRequest struct {
	Time int `json:"time"`
}

const startListeningType = "startListening"

func newStartListeningRequest(bytes []byte) (*StartListeningRequest, error) {
	var request StartListeningRequest

	err := json.Unmarshal(bytes, &request)
	if err != nil {
		return nil, err
	}

	return &request, nil
}

// Stop listening to shell logger:
//	{"command": "ls", "type": "stopListening", "returnCode": 2, "time": 1525212510}
type StopListeningRequest struct {
	Command    string `json:"command"`
	ReturnCode int    `json:"returnCode"`
	Time       int    `json:"time"`
}

const stopListeningType = "stopListening"

func newStopListeningRequest(bytes []byte) (*StopListeningRequest, error) {
	var request StopListeningRequest

	err := json.Unmarshal(bytes, &request)
	if err != nil {
		return nil, err
	}

	if request.Command == "" {
		return nil, errors.New("Malformed request")
	}

	return &request, nil
}

// List logged commands:
//	{"type": "list", "count": 10}
type ListRequest struct {
	Count int `json:"count"`
}

const listType = "list"

func newListRequest(bytes []byte) (*ListRequest, error) {
	var request ListRequest

	err := json.Unmarshal(bytes, &request)
	if err != nil {
		return nil, err
	}

	if request.Count <= 0 {
		return nil, errors.New("Malformed request")
	}

	return &request, nil
}

// Base request used to parse actual type:
//	{"type": "whatever"}
type BaseRequest struct {
	Type string `json:"type"`
}

func requestType(bytes []byte) string {
	var baseRequest BaseRequest
	json.Unmarshal(bytes, &baseRequest)

	return baseRequest.Type
}
