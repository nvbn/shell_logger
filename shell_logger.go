package main

import (
	"github.com/nvbn/shell_logger/shell"
	"flag"
	"os"
	"github.com/nvbn/shell_logger/configurator"
	"github.com/nvbn/shell_logger/wrapper"
	"github.com/nvbn/shell_logger/client"
)

func main() {
	sh, err := shell.Get()
	if err != nil {
		panic(err)
	}

	clientPath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	var configure bool
	var wrap bool
	var startListening bool
	var stopListening bool

	flag.BoolVar(&configure, "configure", false, "Setup shell logger")
	flag.BoolVar(&wrap, "wrap", false, "Wrap shell")
	flag.BoolVar(&startListening, "start-listening", false, "Start listenning to command")
	flag.BoolVar(&stopListening, "stop-listening", false, "Stop listenning to command")

	flag.Parse()

	if (configure) {
		configurator.Configure(clientPath, sh)
	} else if (wrap) {
		wrapper.Wrap(sh)
	} else if (startListening) {
		client.StartListening()
	} else if (stopListening) {
		client.StopListening()
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
