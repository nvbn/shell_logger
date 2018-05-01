package logger

import (
	"os"
	"io"
)

func handleStdin(ptmx *os.File) {
	io.Copy(ptmx, os.Stdin)
}

func handleStdout(ptmx *os.File, ch chan<- []byte) {
	readBuffer := make([]byte, 1)
	sendBuffer := make([]byte, 0)

	for {
		n, err := ptmx.Read(readBuffer)
		if err != nil {
			panic(err)
		}

		if n == 0 {
			os.Exit(0)
		}

		os.Stdout.Write(readBuffer)
		sendBuffer = append(sendBuffer, readBuffer[0])

		if readBuffer[0] == '\n' {
			ch <- sendBuffer
		}
	}
}
