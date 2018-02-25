package shell

import (
	"os"
	"reflect"
	"testing"
)

func TestGetZsh(t *testing.T) {
	os.Setenv("SHELL", "zsh")

	sh, err := Get()
	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}

	result := reflect.TypeOf(sh).String()
	expected := "*shell.zsh"

	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}

func TestGetFish(t *testing.T) {
	os.Setenv("SHELL", "fish")

	sh, err := Get()
	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}

	result := reflect.TypeOf(sh).String()
	expected := "*shell.fish"

	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}

func TestGetUnsupportedShell(t *testing.T) {
	os.Setenv("SHELL", "-")

	sh, err := Get()
	if sh != nil {
		t.Fatalf("Expected error got %s", reflect.TypeOf(sh).String())
	}

	if err == nil {
		t.Fatal("Not expected error")
	}
}

func TestInWrapper(t *testing.T) {
	os.Setenv(SocketEnv, "/tmp/sock")

	result := InWrapper()
	expected := true

	if result != expected {
		t.Fatalf("Expected %t but got %t", expected, result)
	}
}

func TestInWrapperFalse(t *testing.T) {
	os.Setenv(SocketEnv, "")

	result := InWrapper()
	expected := false

	if result != expected {
		t.Fatalf("Expected %t but got %t", expected, result)
	}
}

func TestSetUpUnixSocket(t *testing.T) {
	// Test that no environment variable throws an error
	os.Setenv(SocketEnv, "")

	result := SetUpUnixSocket()
	if result == nil {
		t.Fatalf("Expected nil but got %s", result)
	}
}

func TestGetFailedCommand(t *testing.T) {
    expected := "Failed Command"
    os.Setenv(FailedCommandEnv, expected)
    result := os.Getenv(FailedCommandEnv)

    if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
    }
}

func TestGetSuccessfulCommand(t *testing.T) {
    expected := "Successful Command"
    os.Setenv(CommandEnv, expected)
    result := os.Getenv(CommandEnv)

    if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
    }
}