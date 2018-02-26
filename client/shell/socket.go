package shell

import (
	"log"
	"errors"
	"net"
	"os"
	"os/signal"
	"fmt"
)

func handleSocketConnection(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		command := buf[0:nr]
		fmt.Println(command)
		firstCommand := GetFirstCommand(command)
		if firstCommand == nil {
			c.Write([]byte{0})
			return
		}
		recommendedCommands, err := GetTopThreeCommands(firstCommand)
		if err != nil || len(recommendedCommands) == 0 {
			c.Write([]byte{0})
			return
		}

		// Write the length of the command, and the command itself
		c.Write([]byte{byte(len(recommendedCommands))})
		for i := 0; i < len(recommendedCommands); i++ {
			_, err = c.Write([]byte{byte(len(recommendedCommands[i]))})
			c.Write(recommendedCommands[i])
		}
		if err != nil {
			log.Fatal("Write: ", err)
		}
	}
}

// Sets up unix socket to receive info
func SetUpUnixSocket(done chan os.Signal) (error) {
	if !InWrapper() {
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

	signal.Notify(done, os.Interrupt)
	go ListenAndServe(unixLn)
	<-done
	unixLn.Close()

	return nil
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