package client

import (
	"github.com/nvbn/shell_logger/client/bus"
	"github.com/nvbn/shell_logger/shell"
	"log"
	"os"
	"strconv"
)

const logPrefix = "CLIENT: "

func StartListening() {
	log.SetPrefix(logPrefix)

	socketPath := os.Getenv(shell.SocketEnv)
	startTime, err := strconv.Atoi(os.Getenv(shell.StartTimeEnv))
	if err != nil {
		panic(err)
	}

	bus.StartListening(socketPath, startTime)
}

func StopListening() {
	log.SetPrefix(logPrefix)

	socketPath := os.Getenv(shell.SocketEnv)
	command := os.Getenv(shell.CommandEnv)
	returnCode, err := strconv.Atoi(os.Getenv(shell.ReturnCodeEnv))
	if err != nil {
		panic(err)
	}

	endTime, err := strconv.Atoi(os.Getenv(shell.EndTimeEnv))
	if err != nil {
		panic(err)
	}

	bus.StopListening(
		socketPath,
		command,
		returnCode,
		endTime,
	)
}
