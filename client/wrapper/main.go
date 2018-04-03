package main

import (
	"bufio"
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

func handleStdout(ptmx *os.File, ch chan<- byte) {
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
		ch <- buf[0]
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

	ch := make(chan byte)
	go handleStdout(ptmx, ch)

	f, err := os.Create("/tmp/shell_output")
	w := bufio.NewWriter(f)
	for {
		b := <-ch
		w.WriteByte(b)
		w.Flush()
		f.Sync()
	}
}
