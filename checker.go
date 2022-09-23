package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

type Checker struct {
	browser *rod.Browser
	page    *rod.Page
}

func (c *Checker) LoadTicketPage() {
	c.page = c.browser.MustConnect().MustPage("https://tickets.manutd.com/en-GB/categories/home-tickets")
}

func (c *Checker) DeleteCookieOverlay() {
	tint_overlay, err := c.page.Timeout(5 * time.Second).Element("#tint-overlay")

	if err == nil {
		tint_overlay.Remove()
	}

	cookie_messages, err := c.page.Timeout(5 * time.Second).Element("#cookie-messages")

	if err == nil {
		cookie_messages.Remove()
	}
}

func (c *Checker) FindAvailableEvents() []*Event {
	events := c.page.MustElements("#eventsList .dataItem")
	var availableEvents []*Event

	for _, element := range events {
		event := Event{element}

		fmt.Println(*event.Name())

		_, err := event.FindBuyButton()

		if err == nil {
			availableEvents = append(availableEvents, &event)
		}
	}

	return availableEvents
}

func (c *Checker) Check() {
	events := c.FindAvailableEvents()

	for _, event := range events {
		c.FindTicketPrices(event)
	}
}

func (c *Checker) FindTicketPrices(event *Event) {
	buy_button := event.BuyButton()
	buy_button.MustEval(`() => this.target="_blank"`)
	buy_button.MustClick()
}
