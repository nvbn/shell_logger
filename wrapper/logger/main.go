package logger

import (
	"github.com/kr/pty"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func handleResize(ptmx *os.File) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)

	go func() {
		for range ch {
			pty.InheritSize(os.Stdin, ptmx)
		}
	}()

	ch <- syscall.SIGWINCH
}

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

func Wrap(cmd *exec.Cmd, output <-chan []byte) error {
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return err
	}

	defer ptmx.Close()

	handleResize(ptmx)

	oldState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}

	defer terminal.Restore(int(os.Stdin.Fd()), oldState)

	go handleStdin(ptmx)

	handleStdout(ptmx, output)

	return nil
}
