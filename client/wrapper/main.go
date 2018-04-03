package main

import (
	"os/exec"
	"os"
	"os/signal"
	"syscall"
	"github.com/kr/pty"
	"golang.org/x/crypto/ssh/terminal"
	"io"
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

func handleStdout(ptmx *os.File) {
	buf := make([]byte, 1)
	for {
		n, err := ptmx.Read(buf)
		if err != nil {
			panic(err)
		}

		if n == 0 {
			os.Exit(0)
		}

		os.Stdout.Write(buf)

	}
}

func main() {
	cmd := exec.Command("bash")

	ptmx, err := pty.Start(cmd)
	if err != nil {
		panic(err)
	}

	defer ptmx.Close()

	handleResize(ptmx)

	oldState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(int(os.Stdin.Fd()), oldState)

	go handleStdin(ptmx)
	handleStdout(ptmx)
}
