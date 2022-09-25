# Manchester United Ticket Checker

A application to scrape the Manchested United ticket site to find available matches, prices, and pushes the data to Home Assistant.

## Installation 

```
go install github.com/robertyoung/manutd-ticket-checker/v2@latest
```

Or download the latest release from [Github](https://github.com/RobertYoung/manutd-ticket-checker/releases/latest)

## Usage

```sh
# run application
> ./manutd-ticket-checker

# only show premier league matches
> ./manutd-ticket-checker --premier-league-only

# push state to home assistant and send notification
> ./manutd-ticket-checker --premier-league-only --haas-url ${HA_URL} --haas-token ${HA_TOKEN}

# optionally, a file can be specified to configure the cli
> ./manutd-ticket-checker --env-file env.yml

# cli usage
> ./manutd-ticket-checker -h
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

## Development

```sh
# run application
> go run .

# create a snapshot with goreleaser
> goreleaser release --snapshot --rm-dist
```

## Home Assistsant

![Home Assistant Dashboard Example](/assets/img/haas_dashboard.png "Home Assistant Dashboard Example")
![Home Assistant Notification Example](/assets/img/haas_notification.jpeg "Home Assistant Notification Example")

## Docker

```sh
# docker build
> docker build -t robertyoung/manutd-ticket-checker:dev .

# m1 macbook local build
> docker buildx build --platform linux/arm64 -t robertyoung/manutd-ticket-checker:dev --load .

# run locally
> docker run -it -v $(pwd)/.cache:/root/.cache robertyoung/manutd-ticket-checker:dev

# build and push to dockerhub
> docker buildx build --platform linux/arm64,linux/amd64 -t robertyoung/manutd-ticket-checker:dev --push .
```