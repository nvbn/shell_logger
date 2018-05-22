package client

import (
	"github.com/nvbn/shell_logger/client/bus"
	"github.com/nvbn/shell_logger/shell"
	"log"
)

const logPrefix = "CLIENT: "

func StartListening(sh shell.Shell) {
	log.SetPrefix(logPrefix)

	socketPath := sh.GetSocketPath()
	startTime, err := sh.GetStartTime()
	if err != nil {
		panic(err)
	}

	bus.StartListening(socketPath, startTime)
}

func StopListening(sh shell.Shell) {
	log.SetPrefix(logPrefix)

	socketPath := sh.GetSocketPath()
	command := sh.GetCommand()
	returnCode, err := sh.GetReturnCode()
	if err != nil {
		panic(err)
	}

	endTime, err := sh.GetEndTime()
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
