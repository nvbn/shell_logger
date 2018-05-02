package client

import (
	"github.com/nvbn/shell_logger/shell"
	"github.com/nvbn/shell_logger/client/bus"
	"os"
	"strconv"
)

func StartListening() {
	socketPath := os.Getenv(shell.SocketEnv)
	startTime, err := strconv.Atoi(os.Getenv(shell.StartTimeEnv))
	if err != nil {
		panic(err)
	}

	bus.StartListening(socketPath, startTime)
}

func StopListening() {
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
