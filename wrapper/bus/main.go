package bus

import (
	"github.com/nvbn/shell_logger/wrapper/storage"
	"net"
)

func Start(socketPath string, store *storage.Storage) error {
	unixAddr, err := net.ResolveUnixAddr("unix", socketPath)
	if err != nil {
		return err
	}

	unixLn, err := net.ListenUnix("unix", unixAddr)
	if err != nil {
		return err
	}

	defer unixLn.Close()
	listenAndServe(unixLn, store)

	return nil
}
