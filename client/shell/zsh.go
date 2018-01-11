package shell

import (
	"fmt"
	"text/template"
)

type zsh struct{}

func (z *zsh) SetupWrapper(clientPath string) string {
	return fmt.Sprintf("%[1]s --mode=wrapper", clientPath)
}

var zshHooksTmpl = `
preexec () {
	{{.StartTimeEnv}}=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
}

precmd () {
	{{.ReturnCodeEnv}}=$?
	{{.CommandEnv}}=$(fc -ln -1)
	{{.ClientPath}} --mode=submit
}
`

func (z *zsh) SetupHooks(clientPath string) string {
	tmpl, err := template.New("zsh-hook").Parse(zshHooksTmpl)
	if err != nil {
		panic(err)
	}

	return renderHooks(tmpl, clientPath)
}
