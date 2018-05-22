package shell

import (
	"fmt"
	"os"
	"strconv"
	"text/template"
)

type zsh struct {
	path string
}

func (z *zsh) GetPath() string {
	return z.path
}

func (z *zsh) SetupWrapper(clientPath string) string {
	return fmt.Sprintf("%[1]s --wrap", clientPath)
}

var zshHooksTmpl = `
function shell_logger_preexec () {
	export {{.StartTimeEnv}}=$(date +%s);
	{{.BinaryPath}} --start-listening;
};

function shell_logger_precmd () {
	export {{.ReturnCodeEnv}}=$?;
	export {{.CommandEnv}}=$(fc -ln -1);
	export {{.EndTimeEnv}}=$(date +%s);
	{{.BinaryPath}} --stop-listening
};

autoload -Uz add-zsh-hook;
add-zsh-hook -Uz preexec shell_logger_preexec;
add-zsh-hook -Uz precmd shell_logger_precmd;
`

func (z *zsh) SetupHooks(clientPath string) string {
	tmpl, err := template.New("zsh-hook").Parse(zshHooksTmpl)
	if err != nil {
		panic(err)
	}

	return renderHooks(tmpl, clientPath)
}

func (z *zsh) InWrapper() bool {
	return z.GetSocketPath() != ""
}

func (z *zsh) GetSocketPath() string {
	return os.Getenv(socketEnv)
}

func (z *zsh) SetSocketPath(socketPath string) {
	os.Setenv(socketEnv, socketPath)
}

func (z *zsh) GetStartTime() (int, error) {
	return strconv.Atoi(os.Getenv(startTimeEnv))
}

func (z *zsh) GetCommand() string {
	return os.Getenv(commandEnv)
}

func (z *zsh) GetReturnCode() (int, error) {
	return strconv.Atoi(os.Getenv(returnCodeEnv))
}

func (z *zsh) GetEndTime() (int, error) {
	return strconv.Atoi(os.Getenv(endTimeEnv))
}
