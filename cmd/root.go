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
		CustomAppHelpTemplate: `
		NAME:
		{{.Name}}{{if .Usage}} - {{.Usage}}{{end}}

		USAGE:
		{{.Name}} 

		VERSION:
		{{.Version}}

		AUTHOR:
		{{if .Authors}}
		{{range $index, $author := .Authors}}{{if $index}}, {{end}}{{.Name}}{{end}}
		{{end}}

		COMMANDS:
		{{range .Commands}}
		{{.Name}}{{if .Usage}} - {{.Usage}}{{end}}
		{{end}}

		GLOBAL OPTIONS:
		{{range .VisibleFlags}}
		{{.}}
		{{end}}

		EXAMPLES:
		go run cmd/root.go fetch --number 23 --type math
		go run cmd/root.go fetch --random
		go run cmd/root.go fetch --random --type trivia
		`,
	}

	if err := app.Run(os.Args); err != nil {
		// log.Fatal(err)
		log.Fatalf("App run error: %v", err)
	}
}
