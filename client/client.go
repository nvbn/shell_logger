package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nvbn/shell_logger/client/shell"
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
		shell.SetUpUnixSocket()
		// TODO setup DB first time
	case "wrapper":
		fmt.Println("wrapper")
	case "submit":
		fmt.Println("submit")
	default:
		flag.Usage()
		os.Exit(2)
	}
}
