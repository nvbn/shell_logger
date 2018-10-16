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
		t.Fatalf("Expected %#v but got %#v", expected, result)
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
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestBash_InWrapper(t *testing.T) {
	b := bash{}
	os.Setenv(socketEnv, "/tmp/socket")

	result := b.InWrapper()
	expected := true

	if result != expected {
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestBash_InWrapper_Not(t *testing.T) {
	b := bash{}
	os.Setenv(socketEnv, "")

	result := b.InWrapper()
	expected := false

	if result != expected {
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestBash_GetSocketPath(t *testing.T) {
	b := bash{}
	os.Setenv(socketEnv, "/tmp/socket")

	result := b.GetSocketPath()
	expected := "/tmp/socket"

	if result != expected {
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestBash_SetSocketPath(t *testing.T) {
	b := bash{}
	b.SetSocketPath("/tmp/new-socket")

	result := os.Getenv(socketEnv)
	expected := "/tmp/new-socket"

	if result != expected {
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestBash_GetStartTime(t *testing.T) {
	b := bash{}
	os.Setenv(startTimeEnv, "123")

	result, err := b.GetStartTime()
	if err != nil {
		t.Fatalf("Unexpected error %#v", err)
	}

	expected := 123

	if result != expected {
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestBash_GetStartTime_NotANumber(t *testing.T) {
	b := bash{}
	os.Setenv(startTimeEnv, "test")

	result, err := b.GetStartTime()

	if err == nil {
		t.Fatalf("Expected error but got %#v", result)
	}
}

func TestBash_GetCommand(t *testing.T) {
	b := bash{}
	os.Setenv(commandEnv, "    1  ls ~/.local/bin")

	result := b.GetCommand()
	expected := "ls ~/.local/bin"

	if result != expected {
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestBash_GetReturnCode(t *testing.T) {
	b := bash{}
	os.Setenv(returnCodeEnv, "-1")

	result, err := b.GetReturnCode()
	if err != nil {
		t.Fatalf("Unexpected error %#v", err)
	}

	expected := -1

	if result != expected {
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestBash_GetReturnCode_NotANumber(t *testing.T) {
	b := bash{}
	os.Setenv(returnCodeEnv, "test")

	result, err := b.GetReturnCode()

	if err == nil {
		t.Fatalf("Expected error but got %#v", result)
	}
}

func TestBash_GetEndTime(t *testing.T) {
	b := bash{}
	os.Setenv(endTimeEnv, "123")

	result, err := b.GetEndTime()
	if err != nil {
		t.Fatalf("Unexpected error %#v", err)
	}

	expected := 123

	if result != expected {
		t.Fatalf("Expected %#v but got %#v", expected, result)
	}
}

func TestBash_GetEndTime_NotANumber(t *testing.T) {
	b := bash{}
	os.Setenv(endTimeEnv, "test")

	result, err := b.GetEndTime()

	if err == nil {
		t.Fatalf("Expected error but got %#v", result)
	}
}
