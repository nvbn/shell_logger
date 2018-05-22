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
