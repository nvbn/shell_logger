package main

import (
	"flag"
	"fmt"
	"os"
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

	fmt.Println(sh.SetupHooks(clientPath, shell.DBPath))
}

func inspect(key *string) {
	shell.SetupDatabase(shell.DBPath)
	goodCommands, err := shell.GetGoodCommands([]byte(*key))
	if err != nil {
		log.Fatal(err)
	}
	for _, goodCommand := range goodCommands {
		fmt.Println(string(goodCommand))
	}
}

func main() {
	mode := flag.String("mode", "", "configure|daemon|wrapper|submit")
	key := flag.String("key", "", "key to inspect")

	flag.Parse()

	shell.DBPath = shell.GetDBPath()
	if shell.DBPath == "" {
		log.Fatal("Database path should be specified")
	}

	switch *mode {
	case "configure":
		configure()
	case "inspect":
		inspect(key)
	case "daemon":
		shell.SetupDatabase(shell.DBPath)
		done := make(chan os.Signal, 1)
		shell.SetUpUnixSocket(done)
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
