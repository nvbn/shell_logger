package bus

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/nvbn/shell_logger/wrapper/storage"
	"net"
)

func handleSocketConnection(connection net.Conn, store storage.Storage) {
	reader := bufio.NewReader(connection)
	writer := bufio.NewWriter(connection)

	for {
		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			return
		}

		fmt.Println("Request: ", string(bytes))

		var response []byte
		type_ := requestType(bytes)

		switch type_ {
		case startListeningType:
			response = startListening(store, bytes)
		case stopListeningType:
			response = stopListening(store, bytes)
		case listType:
			response = list(store, bytes)
		default:
			response = errorResponse(errors.New(
				fmt.Sprintf("Unsupported type: %s", type_),
			))
		}

		_, err = writer.Write(append(response, '\n'))
		if err != nil {
			return
		}

		err = writer.Flush()
		if err != nil {
			return
		}

		fmt.Println("Response: ", string(response))
	}
}

func listenAndServe(unixLn *net.UnixListener, store storage.Storage) {
	for {
		unixConn, err := unixLn.Accept()
		if err != nil {
			break
		}
		go handleSocketConnection(unixConn, store)
	}
}
