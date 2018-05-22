package shell

import (
	"os"
	"testing"
)

func TestBash_SetupWrapper(t *testing.T) {
	b := bash{}

	result := b.SetupWrapper("shell_logger")
	expected := "shell_logger --wrap"

	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}

func TestBash_SetupHooks(t *testing.T) {
	s := bash{}

	result := s.SetupHooks("shell_logger")
	expected := `
function shell_logger_preexec () {
	export __SHELL_LOGGER_START_TIME=$(date +%s);
	shell_logger --start-listening;
};

function shell_logger_precmd () {
	export __SHELL_LOGGER_COMMAND=$(builtin history 1);
	export __SHELL_LOGGER_RETURN_CODE=$?;
	export __SHELL_LOGGER_END_TIME=$(date +%s);
	shell_logger --stop-listening;
};

preexec_functions+=(shell_logger_preexec)
precmd_functions+=(shell_logger_precmd)
`
	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}

func TestBash_GetCommand(t *testing.T) {
	b := bash{}
	os.Setenv(commandEnv, "    1  ls ~/.local/bin")

	result := b.GetCommand()
	expected := "ls ~/.local/bin"

	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}
