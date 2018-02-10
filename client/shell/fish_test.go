package shell

import "testing"

func TestFish_SetupWrapper(t *testing.T) {
	f := fish{}

	result := f.SetupWrapper("shell_logger")
	expected := "shell_logger --mode=wrapper"

	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}

func TestFish_SetupHooks(t *testing.T) {
	f := fish{}

	result := f.SetupHooks("shell_logger")
	expected := `
function __shell_logger_preexec -e fish_preexec
  env \
    __SHELL_LOGGER_START_TIME=(date -u +"%Y-%m-%dT%H:%M:%SZ") \
    __SHELL_LOGGER_RETURN_CODE=$status \
    __SHELL_LOGGER_COMMAND=$history[1] \
    shell_logger --mode=submit
end
`
	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}
