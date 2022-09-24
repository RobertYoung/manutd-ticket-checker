# Manchester United Ticket Checker

A application to scrape the Manchested United ticket site to find available matches and the prices.

## Getting Started

```sh
# run application
> go run .

# only show premier league matches
> go run . --premier-league-only

# push state to home assistant and send notification
> go run . --premier-league-only --haas-url ${HA_URL} --haas-token ${HA_TOKEN}

# optionally, a file can be specified to configure the cli
> go run . --env-file env.yml

# cli usage
> go run . -h
NAME:
   manutd-ticket-checker - finds available manchester united tickets

USAGE:
   manutd-ticket-checker [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --env-file value
   --haas-notify-device value, --hnd value  device in home assistant to send the notification to
   --haas-token value, --ht value           token for home assistant to authenticate to the api
   --haas-url value, --hu value             url of home assistant to push state and messages to
   --help, -h                               show help (default: false)
   --premier-league-only, --plo             (default: false)
   --rod value                              rod specific arguments, eg. https://go-rod.github.io/#/get-started/README?id=slow-motion-and-visual-trace
```
