package shell

import (
	"errors"
	"os"
	"path/filepath"
	"net"
	"log"
)

const ReturnCodeEnv = "__SHELL_LOGGER_RETURN_CODE"

const SocketEnv = "__SHELL_LOGGER_SOCKET"

const CommandEnv = "__SHELL_LOGGER_COMMAND"

const FailedCommandEnv = "__SHELL_LOGGER_FAILED_COMMAND"

const StartTimeEnv = "__SHELL_LOGGER_START_TIME"

const FuckCommand = "__SHELL_LOGGER_FUCK_CMD"

type Shell interface {
	// Returns shell specific code for starting our wrapper
	SetupWrapper(clientPath string) string
	// Returns shell specific code for pre/post command hooks
	SetupHooks(clientPath string) string
}

func handleSocketConnection(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]

		// TODO Query DB

		_, err = c.Write(data)
		if err != nil {
			log.Fatal("Write: ", err)
		}
	}
}

// Sets up unix socket to receive info
func SetUpUnixSocket () (error) {
	if (!InWrapper()) {
		var err = errors.New("Set environment variable " + SocketEnv)
		return err
	}
	unixAddr, err := net.ResolveUnixAddr("unix", os.Getenv(SocketEnv))

	if err != nil {
		log.Fatal(err)
		return err
	}

	unixLn, err := net.ListenUnix("unix", unixAddr )

	if err != nil {
		log.Fatal(err)
		return err
	}

	for {
		unixConn, err := unixLn.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		go handleSocketConnection(unixConn)
	}
	defer unixLn.Close()
	return nil
}

// Returns current shell or error
func Get() (Shell, error) {
	shellPath := os.Getenv("SHELL")
	if shellPath == "" {
		return nil, errors.New("Shell can't be identified")
	}

	_, shellName := filepath.Split(shellPath)

	switch shellName {
	case "zsh":
		return &zsh{}, nil
	case "fish":
		return &fish{}, nil
	default:
		return nil, errors.New("Shell is not supported")
	}
}

// Returns true when client runs inside the wrapper
func InWrapper() bool {
	return os.Getenv(SocketEnv) != ""
}

func GetFailedCommand() string {
    return os.Getenv(FailedCommandEnv)
}

func GetSuccessfulCommand() string {
    return os.Getenv(CommandEnv)
}