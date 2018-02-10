package shell

import (
	"fmt"
	"text/template"
)

type fish struct{}

func (f *fish) SetupWrapper(clientPath string) string {
	return fmt.Sprintf("%[1]s --mode=wrapper", clientPath)
}

var fishHooksTmpl = `
function __shell_logger_preexec -e fish_preexec
  env \
    {{.StartTimeEnv}}=(date -u +"%Y-%m-%dT%H:%M:%SZ") \
    {{.ReturnCodeEnv}}=$status \
    {{.CommandEnv}}=$history[1] \
    {{.ClientPath}} --mode=submit
end
`

func (f *fish) SetupHooks(clientPath string) string {
	tmpl, err := template.New("fish-hook").Parse(fishHooksTmpl)
	if err != nil {
		panic(err)
	}

	return renderHooks(tmpl, clientPath)
}
