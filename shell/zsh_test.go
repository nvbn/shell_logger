package shell

import "testing"

func TestZsh_SetupWrapper(t *testing.T) {
	z := zsh{}

	result := z.SetupWrapper("shell_logger")
	expected := "shell_logger --wrap"

	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}

func TestZsh_SetupHooks(t *testing.T) {
	z := zsh{}

	result := z.SetupHooks("shell_logger")
	expected := `
function shell_logger_preexec () {
	export __SHELL_LOGGER_START_TIME=$(date +%s);
	shell_logger --start-listening;
};

function shell_logger_precmd () {
	export __SHELL_LOGGER_RETURN_CODE=$?;
	export __SHELL_LOGGER_COMMAND=$(fc -ln -1);
	export __SHELL_LOGGER_END_TIME=$(date +%s);
	shell_logger --stop-listening
};

autoload -Uz add-zsh-hook;
add-zsh-hook -Uz preexec shell_logger_preexec;
add-zsh-hook -Uz precmd shell_logger_precmd;
`
	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}
