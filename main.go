package main

import (
	"errors"
	"github.com/tozastation/GoGoEnv/Initialize"
	"log"
	"os"

	"github.com/urfave/cli"
)

const (
	N0_ARGUMENTS_COMMAND = ""
	INITIALIZE_COMMAND = "init"
)

func main() {
	app := cli.NewApp()
	app.Name = "GoGoEnv"
	app.Usage = "GoGoEnv is Create Boiler Template"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		command := c.Args().Get(0)
		switch command {
		case N0_ARGUMENTS_COMMAND:
			return errors.New("Enough Arguments")
		case INITIALIZE_COMMAND:
			Initialize.Initialize(c.Args().Get(1))
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
