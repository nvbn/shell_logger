package shell

import (
	"errors"
	"os"
	"path/filepath"
)

const ReturnCodeEnv = "__SHELL_LOGGER_RETURN_CODE"

const SocketEnv = "__SHELL_LOGGER_SOCKET"

const CommandEnv = "__SHELL_LOGGER_COMMAND"

const FailedCommandEnv = "__SHELL_LOGGER_FAILED_COMMAND"

const StartTimeEnv = "__SHELL_LOGGER_START_TIME"

const FuckCommand = "__SHELL_LOGGER_FUCK_CMD"

const DBPathEnv = "__SHELL_LOGGER_DB_PATH"


var DBPath string

type Shell interface {
	// Returns shell specific code for starting our wrapper
	SetupWrapper(clientPath string) string
	// Returns shell specific code for pre/post command hooks
	SetupHooks(clientPath string, dbPath string) string
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
	case "fish":
		return &fish{}, nil
	default:
		return nil, errors.New("Shell is not supported")
	}
}

// Returns true when client runs inside the wrapper
func InWrapper() bool {
	return os.Getenv(SocketEnv) != ""
}

func GetFailedCommand() string {
    return os.Getenv(FailedCommandEnv)
}

func GetSuccessfulCommand() string {
    return os.Getenv(CommandEnv)
}

func GetDBPath() string {
	return os.Getenv(DBPathEnv)
}