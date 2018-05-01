package logger

import (
	"github.com/kr/pty"
	"os"
	"os/signal"
	"syscall"
)

func handleResize(ptmx *os.File) {
	resize := make(chan os.Signal, 1)
	signal.Notify(resize, syscall.SIGWINCH)

	go func() {
		for range resize {
			pty.InheritSize(os.Stdin, ptmx)
		}
	}()

	resize <- syscall.SIGWINCH
}
