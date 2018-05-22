package main

import (
	"flag"
	"github.com/nvbn/shell_logger/client"
	"github.com/nvbn/shell_logger/configurator"
	"github.com/nvbn/shell_logger/shell"
	"github.com/nvbn/shell_logger/wrapper"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var configure bool
	var wrap bool
	var startListening bool
	var stopListening bool

	flag.BoolVar(&configure, "configure", false, "Setup shell logger")
	flag.BoolVar(&wrap, "wrap", false, "Wrap shell")
	flag.BoolVar(&startListening, "start-listening", false, "Start listenning to command")
	flag.BoolVar(&stopListening, "stop-listening", false, "Stop listenning to command")

	flag.Parse()

	if os.Getenv(shell.DebugEnv) == "true" {
		log.SetOutput(os.Stderr)
	} else {
		log.SetOutput(ioutil.Discard)
	}

	sh, err := shell.Get()
	if err != nil {
		panic(err)
	}

	clientPath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	if configure {
		configurator.Configure(clientPath, sh)
	} else if wrap {
		wrapper.Wrap(sh)
	} else if startListening {
		client.StartListening(sh)
	} else if stopListening {
		client.StopListening(sh)
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
