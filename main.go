package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                   "manutd-ticket-checker",
		Usage:                  "finds available manchester united tickets",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "premier-league-only",
				Aliases: []string{"plo"},
			},
			&cli.StringFlag{
				Name:  "rod",
				Usage: "rod specific arguments, eg. https://go-rod.github.io/#/get-started/README?id=slow-motion-and-visual-trace",
			},
		},
		Action: func(cCtx *cli.Context) error {
			premier_league_only := cCtx.Bool("premier-league-only")

			fmt.Print("Finding matches")
			if premier_league_only {
				fmt.Print(" for premier league only")
			}
			fmt.Println()

			checker := UnitedChecker{
				premier_league_only: premier_league_only,
			}

			checker.Check()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
