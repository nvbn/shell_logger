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

	if shell.InWrapper() {
		fmt.Println(sh.SetupHooks(clientPath))
	} else {
		fmt.Println(sh.SetupWrapper(clientPath))
	}
}

func main() {
	mode := flag.String("mode", "", "configure|wrapper|submit")

	flag.Parse()

	switch *mode {
	case "configure":
		configure()
	case "wrapper":
		fmt.Println("wrapper")
	case "submit":
		fmt.Println("submit")
	default:
		flag.Usage()
		os.Exit(2)
	}

}
