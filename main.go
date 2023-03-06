package main

import (
	"os"

	"github.com/alehechka/poker/cmd"
	log "github.com/sirupsen/logrus"
)

// Version of application
var Version = "dev"

func main() {
	if err := cmd.App(Version).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
