package logger

import (
	"os"
	"os/signal"
	"syscall"
	"github.com/kr/pty"
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

