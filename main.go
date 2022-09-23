package main

import (
	"github.com/go-rod/rod"
)

func main() {
	ticket := Checker{browser: rod.New()}

	ticket.LoadTicketPage()
	ticket.DeleteCookieOverlay()
	ticket.Check()
}
