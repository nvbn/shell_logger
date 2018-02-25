package shell

import (
	"fmt"
	"text/template"
)

type zsh struct{}

func (z *zsh) SetupWrapper(clientPath string) string {
	return fmt.Sprintf("%[1]s --mode=wrapper", clientPath)
    return fmt.Sprintf("%[1]s --mode=wrapper", clientPath)
}

var zshHooksTmpl = `
precmd() {
    export __SHELL_LOGGER_RETURN_CODE=$?
    export __SHELL_LOGGER_SOCKET=<socket>
    export __SHELL_LOGGER_COMMAND=$(fc -ln -1)
    export __SHELL_LOGGER_FAILED_COMMAND=$(fc -ln -3 | head -n 1)
    export __SHELL_LOGGER_CLIENT_PATH=$GOPATH/src/github.com/nvbn/shell_logger/client
    export __SHELL_LOGGER_FUCK_CMD=$(fc -ln -2 | head -n 1)
    [ "$__SHELL_LOGGER_FUCK_CMD" == "fuck" ] && $__SHELL_LOGGER_CLIENT_PATH/client -mode=submit
}
`

func (z *zsh) SetupHooks(clientPath string) string {
	tmpl, err := template.New("zsh-hook").Parse(zshHooksTmpl)
	if err != nil {
		panic(err)
	}

	return renderHooks(tmpl, clientPath)
}
