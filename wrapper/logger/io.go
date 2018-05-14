package logger

import (
	"io"
	"os"
)

func handleStdin(ptmx *os.File) {
	io.Copy(ptmx, os.Stdin)
}

func handleStdout(ptmx *os.File, output chan<- []byte) error {
	readBuffer := make([]byte, 1)
	sendBuffer := make([]byte, 0)

	for {
		n, err := ptmx.Read(readBuffer)
		if err != nil {
			return err
		}

		if n == 0 {
			break
		}

		os.Stdout.Write(readBuffer)
		sendBuffer = append(sendBuffer, readBuffer[0])

		if readBuffer[0] == '\n' {
			output <- sendBuffer
			sendBuffer = make([]byte, 0)
		}
	}

	return nil
}
