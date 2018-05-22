package shell

import (
	"errors"
	"os"
	"path/filepath"
)

const DebugEnv = "SHELL_LOGGER_DEBUG"

type Shell interface {
	// Returns shell path
	GetPath() string
	// Returns shell specific code for starting our wrapper
	SetupWrapper(clientPath string) string
	// Returns true when shell is already in wrapper
	InWrapper() bool
	// Returns wrapper socket path
	GetSocketPath() string
	// Set socket path for child processes
	SetSocketPath(socketPath string)
	// Returns shell specific code for pre/post command hooks
	SetupHooks(clientPath string) string
	// Returns sanitized start time
	GetStartTime() (int, error)
	// Returns sanitized previous command
	GetCommand() string
	// Returns sanitized return code
	GetReturnCode() (int, error)
	// Returns sanitized end time
	GetEndTime() (int, error)
}

// Returns current shell or error
func Get() (Shell, error) {
	shellPath := os.Getenv("SHELL")
	if shellPath == "" {
		return nil, errors.New("Shell can't be identified")
	}

	_, shellName := filepath.Split(shellPath)

	switch shellName {
	case "bash":
		return &bash{shellPath}, nil
	case "zsh":
		return &zsh{shellPath}, nil
	case "fish":
		return &fish{shellPath}, nil
	default:
		return nil, errors.New("Shell is not supported")
	}
}
