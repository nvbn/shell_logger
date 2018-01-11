package shell

import (
	"errors"
	"os"
	"path/filepath"
)

const ReturnCodeEnv = "__SHELL_LOGGER_RETURN_CODE"

const SocketEnv = "__SHELL_LOGGER_SOCKET"

const CommandEnv = "__SHELL_LOGGER_COMMAND"

const StartTimeEnv = "__SHELL_LOGGER_START_TIME"

type Shell interface {
	// Returns shell specific code for starting our wrapper
	SetupWrapper(clientPath string) string
	// Returns shell specific code for pre/post command hooks
	SetupHooks(clientPath string) string
}

// Returns current shell or error
func Get() (Shell, error) {
	shellPath := os.Getenv("SHELL")
	if shellPath == "" {
		return nil, errors.New("Shell can't be identified")
	}

	_, shellName := filepath.Split(shellPath)

	switch shellName {
	case "zsh":
		return &zsh{}, nil
	default:
		return nil, errors.New("Shell is not supported")
	}
}

// Returns true when client runs inside the wrapper
func InWrapper() bool {
	return os.Getenv(SocketEnv) != ""
}
