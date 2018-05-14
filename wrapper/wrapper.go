package wrapper

import (
	"fmt"
	"github.com/nvbn/shell_logger/shell"
	"github.com/nvbn/shell_logger/wrapper/bus"
	"github.com/nvbn/shell_logger/wrapper/logger"
	"github.com/nvbn/shell_logger/wrapper/storage"
	"github.com/satori/go.uuid"
	"log"
	"os"
	"os/exec"
)

func generateSocketPath() string {
	id, _ := uuid.NewV4()
	return fmt.Sprintf("/tmp/shell-logger-%s", id)
}

func wrapShell(sh shell.Shell) chan []byte {
	command := exec.Command(sh.GetPath())
	output := make(chan []byte)

	go func() {
		err := logger.Wrap(command, output)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	return output
}

func Wrap(sh shell.Shell) {
	socketPath := generateSocketPath()
	os.Setenv(shell.SocketEnv, socketPath)
	output := wrapShell(sh)
	store := storage.NewInMemory(output)
	log.Println("Wrapper started on", socketPath)
	bus.Start(socketPath, store)
}
