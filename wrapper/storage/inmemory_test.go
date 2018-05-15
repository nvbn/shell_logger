// +build race

package storage

import (
	"reflect"
	"testing"
	"time"
)

func TestInMemoryStorage_NewInMemory(t *testing.T) {
	buffer := make(chan []byte)
	defer close(buffer)

	store := NewInMemory(buffer)

	result := reflect.TypeOf(store).String()
	expected := "*storage.inMemoryStorage"

	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}

func TestInMemoryStorage_StartListening(t *testing.T) {
	buffer := make(chan []byte)
	defer close(buffer)

	store := NewInMemory(buffer)
	buffer <- []byte("first line\n")

	store.StartListening(0)

	result := store.List(10)
	expected := []*Command{}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}

func TestInMemoryStorage_StopListening(t *testing.T) {
	buffer := make(chan []byte)
	defer close(buffer)

	store := NewInMemory(buffer)
	time.Sleep(100)
	buffer <- []byte("first line\n")
	time.Sleep(100)

	store.StartListening(0)
	time.Sleep(100)
	buffer <- []byte("second line\n")
	time.Sleep(100)

	store.StopListening("ls", 0, 1)
	time.Sleep(100)
	buffer <- []byte("third line\n")
	time.Sleep(100)

	result := store.List(10)
	expected := []*Command{
		{"ls", "second line\n", 0, 0, 1},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}
