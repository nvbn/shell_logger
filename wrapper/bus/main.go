package bus

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

type Command struct {
	Name string `json:"name"`
}

type StartListenCommand struct {
	Time string `json:"time"`
}

type StopListenCommand struct {
	Command    string `json:"command"`
	ReturnCode int    `json:"returnCode"`
	Time       string `json:"time"`
}

type GetCommand struct {
	Number int `json:"number"`
}

func handleSocketConnection(connection net.Conn) {
	reader := bufio.NewReader(connection)

	for {
		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			return
		}

		var baseCommand Command
		json.Unmarshal(bytes, &baseCommand)

		switch baseCommand.Name {
		case "start":
			var command StartListenCommand
			json.Unmarshal(bytes, &command)
			fmt.Println("start:", command)
		case "stop":
			var command StopListenCommand
			json.Unmarshal(bytes, &command)
			fmt.Println("stop:", command)
		case "get":
			var command GetCommand
			json.Unmarshal(bytes, &command)
			fmt.Println("get:", command)
		default:
			fmt.Println("Not supported command ", baseCommand.Name)
		}
	}
}

func ListenAndServe(unixLn *net.UnixListener) {
	for {
		unixConn, err := unixLn.Accept()
		if err != nil {
			break
		}
		go handleSocketConnection(unixConn)
	}
}

func Start(socketPath string) error {
	unixAddr, err := net.ResolveUnixAddr("unix", socketPath)
	if err != nil {
		return err
	}

	unixLn, err := net.ListenUnix("unix", unixAddr)
	if err != nil {
		return err
	}

	defer unixLn.Close()
	ListenAndServe(unixLn)

	return nil
}
