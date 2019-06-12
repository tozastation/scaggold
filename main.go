package main

import (
	"errors"
	"github.com/tozastation/GoGoEnv/Initialize"
	"log"
	"os"

	"github.com/urfave/cli"
)

const (
	// N0_ARGUMENTS_COMMAND is has not arguments your command
	N0ArgumentsCommand = ""
	InitializeCommand = "init"
)

func main() {
	app := cli.NewApp()
	app.Name = "GoGoEnv"
	app.Usage = "GoGoEnv is Create Boiler Template"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		command := c.Args().Get(0)
		switch command {
		case N0ArgumentsCommand:
			return errors.New("[Enough Arguments]: please type one argument at least")
		case InitializeCommand:
			subCommand := c.Args().Get(1)
			if subCommand == "" {
				return errors.New("[Enough Arguments]: init + {Your Application Name}")
			}
			Initialize.Initialize(subCommand)
			return nil
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
