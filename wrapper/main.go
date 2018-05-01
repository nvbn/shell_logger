package main

import (
	"os/exec"
	"github.com/nvbn/shell_logger/wrapper/logger"
	"github.com/nvbn/shell_logger/wrapper/storage"
	"github.com/nvbn/shell_logger/wrapper/bus"
)

func main() {
	command := exec.Command("bash")
	buffer := make(chan []byte)
	go logger.Wrap(command, buffer)
	store := storage.New(buffer)
	bus.Start("/tmp/test", store)
}
