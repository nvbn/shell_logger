package bus

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"github.com/nvbn/shell_logger/wrapper/storage"
)

func handleSocketConnection(connection net.Conn, store *storage.Storage) {
	reader := bufio.NewReader(connection)

	for {
		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			return
		}

		var baseRequest Request
		json.Unmarshal(bytes, &baseRequest)

		switch baseRequest.Type {
		case "start":
			var request StartListenRequest
			json.Unmarshal(bytes, &request)
			fmt.Println("start:", request)
			store.StartListen(request.Time)
		case "stop":
			var request StopListenRequest
			json.Unmarshal(bytes, &request)
			fmt.Println("stop:", request)
			store.StopListen(request.Command, request.ReturnCode, request.Time)
		case "get":
			var request GetRequest
			json.Unmarshal(bytes, &request)
			fmt.Println("get:", request)
			fmt.Println("val:", store.Get(request.Number))
		default:
			fmt.Println("Not supported request:", baseRequest.Type)
		}
	}
}
