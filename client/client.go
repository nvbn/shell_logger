package main

import (
	"flag"
	"fmt"
	"os"
    // "log"
	"github.com/nvbn/shell_logger/client/shell"
	"log"
)

func configure() {
	clientPath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	sh, err := shell.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(sh.SetupHooks(clientPath))
}

func main() {
	mode := flag.String("mode", "", "configure|daemon|wrapper|submit")

	flag.Parse()

	switch *mode {
	case "configure":
		configure()
	case "daemon":
		shell.SetupDatabase()
		shell.SetUpUnixSocket()
	case "wrapper":
		fmt.Println("wrapper")
	case "submit":
        var successfulCommand string = os.Getenv(shell.CommandEnv)
        var failedCommand string = os.Getenv(shell.FailedCommandEnv)

        err := shell.Insert([]byte(successfulCommand), []byte(failedCommand))
        if err != nil {
            log.Fatal(err)
        }
	default:
		flag.Usage()
		os.Exit(2)
	}
}
