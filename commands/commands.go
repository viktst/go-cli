package commands

import (
	"fmt"
	"go-cli/fetchers"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

// Types
const (
	Math   = "math"
	Year   = "year"
	Date   = "date"
	Trivia = "trivia"
)

type FactType string

func (ft FactType) IsValid() bool {
	switch ft {
	case
		Trivia, Math, Date, Year:
		return true
	default:
		return false
	}
}

var FetchCommand = &cli.Command{
	Name:        "fetch",
	Usage:       "Fetch a fact about a number",
	Description: "..",
	Flags: []cli.Flag{
		&cli.IntFlag{Name: "number", Value: 0, Usage: "Number to fetch a fact about", Aliases: []string{"n"}},
		&cli.StringFlag{Name: "type", Value: Math, Usage: "Type of fact", Aliases: []string{"t"}},
		&cli.BoolFlag{Name: "random", Value: false, Usage: "Fetch a random fact", Aliases: []string{"r"}},
	},
	Action: func(c *cli.Context) error {
		errOutput := color.New(color.FgRed).SprintFunc()
		output := color.New(color.FgMagenta).SprintFunc()

		if c.Bool("random") && c.Int("number") != 0 {
			return fmt.Errorf(errOutput("err: --random and --number cannot be used together"))
		}

		if !c.Bool("random") && c.Int("number") == 0 {
			return fmt.Errorf(errOutput("err: --random or --number must be specified to fetch a fact"))
		}

		number := GetNumber(c)
		factType := FactType(c.String("type"))

		if !factType.IsValid() {
			return fmt.Errorf(errOutput("invalid fact type: %s", factType))
		}

		fact, err := fetchers.GetFacts(number, string(factType))
		if err != nil {
			return fmt.Errorf(errOutput("failed to fetch fact: %w", err))
		}

		fmt.Println(output(fact)) // Output
		return nil
	},
}

func GetNumber(c *cli.Context) interface{} {
	if c.Bool("random") {
		return "random"
	}
	return c.Int("number")
}
