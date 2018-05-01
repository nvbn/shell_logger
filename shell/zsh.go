package shell

import (
	"fmt"
	"text/template"
)

type zsh struct{
	path string
}

func (z *zsh) GetPath() string {
	return z.path
}

func (z *zsh) SetupWrapper(clientPath string) string {
	return fmt.Sprintf("%[1]s --mode=wrapper", clientPath)
}

var zshHooksTmpl = `
preexec () {
	export {{.StartTimeEnv}}=$(date +%s)
}
precmd () {
	export {{.ReturnCodeEnv}}=$?
	export {{.CommandEnv}}=$(fc -ln -1)
	export {{.FailedCommandEnv}}=$(fc -ln -3 | head -n 1)
	export {{.EndTimeEnv}}=$(date +%s)
	shell_logger --mode=submit
}
`

func (z *zsh) SetupHooks(clientPath string) string {
	tmpl, err := template.New("zsh-hook").Parse(zshHooksTmpl)
	if err != nil {
		panic(err)
	}

	return renderHooks(tmpl, clientPath)
}
