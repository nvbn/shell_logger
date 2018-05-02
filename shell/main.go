package shell

import (
	"errors"
	"os"
	"path/filepath"
)

const ReturnCodeEnv = "__SHELL_LOGGER_RETURN_CODE"

const SocketEnv = "SHELL_LOGGER_SOCKET"

const CommandEnv = "__SHELL_LOGGER_COMMAND"

const StartTimeEnv = "__SHELL_LOGGER_START_TIME"

const EndTimeEnv = "__SHELL_LOGGER_END_TIME"


type Shell interface {
	// Returns shell path
	GetPath() string
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
		return &zsh{shellPath}, nil
	case "fish":
		return &fish{shellPath}, nil
	default:
		return nil, errors.New("Shell is not supported")
	}
}

// Returns true when client runs inside the wrapper
func InWrapper() bool {
	return os.Getenv(SocketEnv) != ""
}
