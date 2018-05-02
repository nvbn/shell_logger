package shell

import (
	"bytes"
	"text/template"
)

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
		ReturnCodeEnv: ReturnCodeEnv,
		CommandEnv:    CommandEnv,
		StartTimeEnv:  StartTimeEnv,
		EndTimeEnv:    EndTimeEnv,
	}

	var hook bytes.Buffer
	tmpl.Execute(&hook, hookOptions)

	return hook.String()
}
