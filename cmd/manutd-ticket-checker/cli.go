package mutc

import (
	"log"
	"os"

	haas "github.com/robertyoung/manutd-ticket-checker/v2/pkg/home-assistant"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func Cli() {
	flags := []cli.Flag{
		&cli.StringFlag{Name: "env-file"},
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:    "premier-league-only",
			Aliases: []string{"plo"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:     "haas-url",
			Usage:    "url of home assistant to push state and messages to",
			Aliases:  []string{"hu"},
			Required: false,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:     "haas-token",
			Usage:    "token for home assistant to authenticate to the api",
			Aliases:  []string{"ht"},
			Required: false,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:     "haas-notify-device",
			Usage:    "device in home assistant to send the notification to",
			Aliases:  []string{"hnd"},
			Required: false,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "rod",
			Usage: "rod specific arguments, eg. https://go-rod.github.io/#/get-started/README?id=slow-motion-and-visual-trace",
		}),
	}
	app := &cli.App{
		Name:                   "manutd-ticket-checker",
		Usage:                  "finds available manchester united tickets",
		UseShortOptionHandling: true,
		Flags:                  flags,
		Before:                 altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("env-file")),
		Action: func(cCtx *cli.Context) error {
			premier_league_only := cCtx.Bool("premier-league-only")
			haas_url := cCtx.String("haas-url")
			haas_token := cCtx.String("haas-token")
			haas_notify_device := cCtx.String("haas-notify-device")

			if premier_league_only {
				log.Println("Finding matches for premier league only")
			} else {
				log.Println("Finding matches")
			}

			var haas_api *haas.HomeAssistantAPI = nil

			if haas_url != "" && haas_token != "" {
				haas_api = haas.NewHomeAssistantAPI(haas_url, haas_token)
			} else {
				log.Println("home assistant integration disabled due to missing url and token")
			}

			checker := UnitedChecker{
				premier_league_only: premier_league_only,
				haas_api:            haas_api,
				haas_notify_device:  haas_notify_device,
			}

			checker.Check()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
