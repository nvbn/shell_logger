package shell

import (
	"bytes"
	"text/template"
)

const returnCodeEnv = "__SHELL_LOGGER_RETURN_CODE"

const socketEnv = "SHELL_LOGGER_SOCKET"

const commandEnv = "__SHELL_LOGGER_COMMAND"

const startTimeEnv = "__SHELL_LOGGER_START_TIME"

const endTimeEnv = "__SHELL_LOGGER_END_TIME"

type hooksOptions struct {
	ReturnCodeEnv string
	CommandEnv    string
	StartTimeEnv  string
	EndTimeEnv    string
	BinaryPath    string
}

func renderHooks(tmpl *template.Template, clientPath string) string {
	hookOptions := hooksOptions{
		BinaryPath:    clientPath,
		ReturnCodeEnv: returnCodeEnv,
		CommandEnv:    commandEnv,
		StartTimeEnv:  startTimeEnv,
		EndTimeEnv:    endTimeEnv,
	}

	var hook bytes.Buffer
	tmpl.Execute(&hook, hookOptions)

	return hook.String()
}
