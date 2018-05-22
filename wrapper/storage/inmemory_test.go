package storage

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestInMemoryStorage_Empty(t *testing.T) {
	buffer := make(chan []byte)
	defer close(buffer)

	store := NewInMemory(buffer)

	result := store.List(10)
	expected := []*Command{}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}

func TestInMemoryStorage_SingleCommand(t *testing.T) {
	buffer := make(chan []byte)
	defer close(buffer)

	store := NewInMemory(buffer)

	startTime := 1000
	output := "output"
	command := "ls"
	returnCode := 1
	endTime := 2000

	store.StartListening(startTime)
	time.Sleep(10 * time.Millisecond)

	buffer <- []byte(output)
	time.Sleep(10 * time.Millisecond)

	store.StopListening(command, returnCode, endTime)

	result := store.List(10)
	expected := []*Command{
		{command, output, returnCode, startTime, endTime},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}

func TestInMemoryStorage_MoreThenLimit(t *testing.T) {
	buffer := make(chan []byte)
	defer close(buffer)

	store := NewInMemory(buffer)

	expected := []*Command{}

	for i := 0; i <= inMemoryStorageSize+10; i++ {
		startTime := 1000 + i
		output := fmt.Sprintf("output %d", i)
		command := fmt.Sprintf("ls %d", i)
		returnCode := i
		endTime := 2000 + i

		store.StartListening(startTime)
		time.Sleep(10 * time.Millisecond)

		buffer <- []byte(output)
		time.Sleep(10 * time.Millisecond)

		store.StopListening(command, returnCode, endTime)

		result := store.List(10)
		expected = append([]*Command{
			{command, output, returnCode, startTime, endTime},
		}, expected...)

		if i >= 10 {
			expected = expected[:10]
		}

		if !reflect.DeepEqual(result, expected) {
			t.Fatalf("Expected %s but got %s", expected, result)
		}
	}
}
