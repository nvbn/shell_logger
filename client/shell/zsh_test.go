package shell

import "testing"

func TestZsh_SetupWrapper(t *testing.T) {
	z := zsh{}

	result := z.SetupWrapper("shell_logger")
	expected := "shell_logger --mode=wrapper"

	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}

func TestZsh_SetupHooks(t *testing.T) {
	z := zsh{}

	result := z.SetupHooks("shell_logger")
	expected := `
preexec () {
	__SHELL_LOGGER_START_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
}

precmd () {
	__SHELL_LOGGER_RETURN_CODE=$?
	__SHELL_LOGGER_COMMAND=$(fc -ln -1)
	shell_logger --mode=submit
}
`
	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}
