package bus

import (
	"net"
	"github.com/nvbn/shell_logger/wrapper/storage"
)

func listenAndServe(unixLn *net.UnixListener, store *storage.Storage) {
	for {
		unixConn, err := unixLn.Accept()
		if err != nil {
			break
		}
		go handleSocketConnection(unixConn, store)
	}
}

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
