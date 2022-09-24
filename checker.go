package main

import "github.com/go-rod/rod"

type UnitedChecker struct {
	browser    *rod.Browser
	event_list *UnitedEventListPage
}

func (c *UnitedChecker) Check() {
	c.browser = rod.New()
	c.event_list.DeleteCookieOverlay()

	events := c.event_list.FindAvailableEvents()

	for _, event := range events {
		c.FindTicketPrices(event)
	}
}

func (c *UnitedChecker) LoadEventListPage() {
	c.event_list = &UnitedEventListPage{
		&UnitedPage{
			c.browser.MustConnect().MustPage("https://tickets.manutd.com/en-GB/categories/home-tickets"),
		},
	}
}

func (c *UnitedChecker) FindTicketPrices(event *UnitedEventItem) {
	buy_button := event.BuyButton()
	buy_button.MustEval(`() => this.target="_blank"`)
	buy_button.MustClick()
}
