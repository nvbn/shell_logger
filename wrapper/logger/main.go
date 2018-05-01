package logger

import (
	"github.com/kr/pty"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"os/exec"
	"strings"
)

// Wraps command and logs output in passed channel.
func Wrap(cmd *exec.Cmd, output chan<- []byte) error {
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

	err = handleStdout(ptmx, output)
	// It's expected case when ptmx is already closed
	if err != nil && !strings.HasPrefix(err.Error(), "read") {
		return err
	}

	return nil
}
