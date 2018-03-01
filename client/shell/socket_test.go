package shell

import (
	"io"
	"log"
	"net"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestSetInvalidSocketEnv(t *testing.T) {
	// Test that no environment variable throws an error
	os.Setenv(SocketEnv, "")

	done := make(chan os.Signal, 1)
	result := SetUpUnixSocket(done)
	if result == nil {
		t.Fatalf("Expected nil but got %s", result)
	}
}

func readFromSocket(done chan os.Signal, r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		if n != 1 {
			log.Fatalf("Expected 1 byte but got %d", n)
			return
		}

		// If the client didn't get back 0, then it's a fail
		expected := byte(0)
		if buf[0] != expected {
			log.Fatalf("Expected %v but got %v", expected, buf[0])
		}
		break
	}
}

func TestSetUpUnixSocket(t *testing.T) {
	DBPath = "test.db"
	ClearData()
	SetupDatabase(DBPath)

	socketFile := "/tmp/test_socket"
	os.Setenv(SocketEnv, socketFile)
	done := make(chan os.Signal, 1)
	go SetUpUnixSocket(done)
	time.Sleep(500 * time.Millisecond)

	c, err := net.Dial("unix", socketFile)
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer c.Close()

	go readFromSocket(done, c)
	for {
		msg := "command"
		n, err := c.Write([]byte(msg))
		if err != nil {
			log.Fatal("Write error:", err)
			break
		}
		if n == len(msg) {
			break
		}
	}
	done <- syscall.SIGINT
	// Wait a second to make sure socket file is cleaned up
	time.Sleep(1 * time.Second)
}
