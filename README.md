# Manchester United Ticket Checker

A application to scrape the Manchested United ticket site to find available matches and the prices.

## Getting Started

```sh
# run application
> go run .

# only show premier league matches
> go run . --premier-league-only
Finding matches for premier league only
Checking Manchester United v Newcastle United... prices found: £320 -> £320
Checking Manchester United v Tottenham Hotspur... prices found: £440 -> £440
Checking Manchester United v West Ham United... prices found: £320 -> £320
Checking Manchester United v Nottingham Forest... prices found: £50 -> £270
Checking Manchester United v A.F.C. Bournemouth... prices found: £270 -> £270
Checking Manchester United v Manchester City... prices found: £534 -> £534
Checking Manchester United v Crystal Palace... prices found: £270 -> £270
Checking Manchester United v Leicester City... prices found: £270 -> £270

# cli usage
> go run . -h
NAME:
   manutd-ticket-checker - finds available manchester united tickets

USAGE:
   manutd-ticket-checker [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h                    show help (default: false)
   --premier-league-only, --plo  (default: false)
   --rod value                   rod specific arguments, eg. https://go-rod.github.io/#/get-started/README?id=slow-motion-and-visual-trace
```
