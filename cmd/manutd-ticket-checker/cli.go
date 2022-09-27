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
			Usage:   "filter premier league events only",
			Aliases: []string{"plo"},
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:    "max-price",
			Usage:   "the maximum price to mark an event as available",
			Aliases: []string{"mp"},
			Value:   100,
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
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:     "haas-notification-throttle",
			Usage:    "duration in minutes to wait before resending a notification",
			Value:    60,
			Aliases:  []string{"rnd"},
			Required: false,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:     "rod",
			Usage:    "rod specific arguments, eg. https://go-rod.github.io/#/get-started/README?id=slow-motion-and-visual-trace",
			Required: false,
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
			max_price := cCtx.Int("max-price")
			haas_url := cCtx.String("haas-url")
			haas_token := cCtx.String("haas-token")
			haas_notify_device := cCtx.String("haas-notify-device")
			haas_notification_throttle := cCtx.Int("haas-notification-throttle")

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
				haas_api: haas_api,
				config: &Config{
					PremierLeagueOnly:        premier_league_only,
					MaxPrice:                 max_price,
					HaasUrl:                  haas_url,
					HaasToken:                haas_token,
					HaasNotifyDevice:         haas_notify_device,
					HaasNotificationThrottle: haas_notification_throttle,
				},
			}

			checker.Check()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
