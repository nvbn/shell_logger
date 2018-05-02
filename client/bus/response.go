package bus

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Base response:
//	{"status": "everything"}
type BaseResponse struct {
	Status string `json:"status"`
}

const statusOk = "ok"

func isOk(body []byte) bool {
	var response BaseResponse

	err := json.Unmarshal(body, &response)
	if err != nil {
		return false
	}

	return response.Status == statusOk
}

// Error response:
//	{"status": "error", "error": "everything"}
type ErrorResponse struct {
	Error string `json:"error"`
}

func getError(body []byte) error {
	var response ErrorResponse

	err := json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	return errors.New(fmt.Sprintf("Wrapper error: %s", response.Error))
}
