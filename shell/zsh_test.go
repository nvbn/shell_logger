package shell

import (
	"os"
	"testing"
)

func TestZsh_SetupWrapper(t *testing.T) {
	z := zsh{}

	result := z.SetupWrapper("shell_logger")
	expected := "shell_logger --wrap"

	if result != expected {
		t.Fatalf("Expected %#v but got %#v", expected, result)
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
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestZsh_InWrapper(t *testing.T) {
	z := zsh{}
	os.Setenv(socketEnv, "/tmp/socket")

	result := z.InWrapper()
	expected := true

	if result != expected {
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestZsh_InWrapper_Not(t *testing.T) {
	z := zsh{}
	os.Setenv(socketEnv, "")

	result := z.InWrapper()
	expected := false

	if result != expected {
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestZsh_GetSocketPath(t *testing.T) {
	z := zsh{}
	os.Setenv(socketEnv, "/tmp/socket")

	result := z.GetSocketPath()
	expected := "/tmp/socket"

	if result != expected {
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestZsh_SetSocketPath(t *testing.T) {
	z := zsh{}
	z.SetSocketPath("/tmp/new-socket")

	result := os.Getenv(socketEnv)
	expected := "/tmp/new-socket"

	if result != expected {
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestZsh_GetStartTime(t *testing.T) {
	z := zsh{}
	os.Setenv(startTimeEnv, "123")

	result, err := z.GetStartTime()
	if err != nil {
		t.Fatalf("Unexpected error %#v", err)
	}

	expected := 123

	if result != expected {
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestZsh_GetStartTime_NotANumber(t *testing.T) {
	z := zsh{}
	os.Setenv(startTimeEnv, "test")

	result, err := z.GetStartTime()

	if err == nil {
		t.Fatalf("Expected error but got %#v", result)
	}
}

func TestZsh_GetCommand(t *testing.T) {
	z := zsh{}
	os.Setenv(commandEnv, "ls /")

	result := z.GetCommand()
	expected := "ls /"

	if result != expected {
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestZsh_GetReturnCode(t *testing.T) {
	z := zsh{}
	os.Setenv(returnCodeEnv, "-1")

	result, err := z.GetReturnCode()
	if err != nil {
		t.Fatalf("Unexpected error %#v", err)
	}

	expected := -1

	if result != expected {
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestZsh_GetReturnCode_NotANumber(t *testing.T) {
	z := zsh{}
	os.Setenv(returnCodeEnv, "test")

	result, err := z.GetReturnCode()

	if err == nil {
		t.Fatalf("Expected error but got %#v", result)
	}
}

func TestZsh_GetEndTime(t *testing.T) {
	z := zsh{}
	os.Setenv(endTimeEnv, "123")

	result, err := z.GetEndTime()
	if err != nil {
		t.Fatalf("Unexpected error %#v", err)
	}

	expected := 123

	if result != expected {
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestZsh_GetEndTime_NotANumber(t *testing.T) {
	z := zsh{}
	os.Setenv(endTimeEnv, "test")

	result, err := z.GetEndTime()

	if err == nil {
		t.Fatalf("Expected error but got %#v", result)
	}
}
