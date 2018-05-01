package main

import (
	"github.com/nvbn/shell_logger/wrapper/bus"
	"github.com/nvbn/shell_logger/wrapper/logger"
	"github.com/nvbn/shell_logger/wrapper/storage"
	"os/exec"
)

func main() {
	command := exec.Command("bash")
	buffer := make(chan []byte)
	go logger.Wrap(command, buffer)
	store := storage.NewInMemory(buffer)
	bus.Start("/tmp/test", store)
}
