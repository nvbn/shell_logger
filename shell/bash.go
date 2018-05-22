package shell

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"text/template"
)

type bash struct {
	path string
}

func (b *bash) GetPath() string {
	return b.path
}

func (b *bash) SetupWrapper(clientPath string) string {
	return fmt.Sprintf("%[1]s --wrap", clientPath)
}

var bashHooksTmpl = `
function shell_logger_preexec () {
	export {{.StartTimeEnv}}=$(date +%s);
	{{.BinaryPath}} --start-listening;
};

function shell_logger_precmd () {
	export {{.CommandEnv}}=$(builtin history 1);
	export {{.ReturnCodeEnv}}=$?;
	export {{.EndTimeEnv}}=$(date +%s);
	{{.BinaryPath}} --stop-listening;
};

preexec_functions+=(shell_logger_preexec)
precmd_functions+=(shell_logger_precmd)
`

func (b *bash) SetupHooks(clientPath string) string {
	tmpl, err := template.New("bash-hook").Parse(bashHooksTmpl)
	if err != nil {
		panic(err)
	}

	return renderHooks(tmpl, clientPath)
}

func (b *bash) InWrapper() bool {
	return b.GetSocketPath() != ""
}

func (b *bash) GetSocketPath() string {
	return os.Getenv(socketEnv)
}

func (b *bash) SetSocketPath(socketPath string) {
	os.Setenv(socketEnv, socketPath)
}

func (b *bash) GetStartTime() (int, error) {
	return strconv.Atoi(os.Getenv(startTimeEnv))
}

func (b *bash) GetCommand() string {
	historyEntry := os.Getenv(commandEnv)
	re := regexp.MustCompile(`^\s+\d+\s+`)
	return re.ReplaceAllString(historyEntry, "")
}

func (b *bash) GetReturnCode() (int, error) {
	return strconv.Atoi(os.Getenv(returnCodeEnv))
}

func (b *bash) GetEndTime() (int, error) {
	return strconv.Atoi(os.Getenv(endTimeEnv))
}
