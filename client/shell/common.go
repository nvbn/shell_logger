package shell

import (
	"bytes"
	"text/template"
)

type hooksOptions struct {
	ClientPath    string
	ReturnCodeEnv string
	CommandEnv    string
	StartTimeEnv  string
}

func renderHooks(tmpl *template.Template, clientPath string) string {
	hookOptions := hooksOptions{
		ClientPath:    clientPath,
		ReturnCodeEnv: ReturnCodeEnv,
		CommandEnv:    CommandEnv,
		StartTimeEnv:  StartTimeEnv,
	}

	var hook bytes.Buffer
	tmpl.Execute(&hook, hookOptions)

	return hook.String()
}
