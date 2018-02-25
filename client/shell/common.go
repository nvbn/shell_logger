package shell

import (
	"bytes"
	"text/template"
)

type hooksOptions struct {
	ReturnCodeEnv string
	CommandEnv    string
	StartTimeEnv  string
	FailedCommandEnv string
	FuckCommand   string
	BinaryPath	  string
}

func renderHooks(tmpl *template.Template, clientPath string) string {
	hookOptions := hooksOptions{
		BinaryPath:    clientPath,
		ReturnCodeEnv:    ReturnCodeEnv,
		FailedCommandEnv: FailedCommandEnv,
		CommandEnv:    CommandEnv,
		StartTimeEnv:  StartTimeEnv,
		FuckCommand:   FuckCommand,
	}

	var hook bytes.Buffer
	tmpl.Execute(&hook, hookOptions)

	return hook.String()
}
