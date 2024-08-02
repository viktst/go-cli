package main

import (
	"go-cli/commands"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "Number Facts CLI App",
		Usage:    "Fetch interesting facts about numbers",
		Compiled: time.Now(),
		Version:  "v1.0.0",
		Authors: []*cli.Author{
			{
				Name: "Me",
			},
		},
		Commands: []*cli.Command{
			commands.FetchCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		// log.Fatal(err)
		log.Fatalf("App run error: %v", err)
	}
}
