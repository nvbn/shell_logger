package bus

import (
	"bufio"
	"log"
	"net"
)

func communicate(socketPath string, body []byte) ([]byte, error) {
	connection, err := net.Dial("unix", socketPath)
	if err != nil {
		return nil, err
	}
	defer connection.Close()

	writer := bufio.NewWriter(connection)
	log.Println("Sending:", string(body))
	_, err = writer.Write(append(body, '\n'))
	if err != nil {
		return nil, err
	}

	err = writer.Flush()
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(connection)
	bytes, err := reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	log.Println("Received:", string(bytes))

	return bytes, nil
}
