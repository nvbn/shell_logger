package main

import (
	"fmt"
	"github.com/nvbn/shell_logger/wrapper/bus"
	"github.com/nvbn/shell_logger/wrapper/logger"
	"github.com/nvbn/shell_logger/wrapper/storage"
	"os"
	"os/exec"
)

func main() {
	command := exec.Command("bash")
	buffer := make(chan []byte)

	go func() {
		err := logger.Wrap(command, buffer)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	store := storage.NewInMemory(buffer)

	bus.Start("/tmp/test", store)
}
