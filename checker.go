package main

import "github.com/go-rod/rod"

type UnitedChecker struct {
	browser    *rod.Browser
	event_list *UnitedEventListPage
}

func (c *UnitedChecker) Check() {
	c.browser = rod.New()
	c.LoadEventListPage()
	c.event_list.DeleteCookieOverlay()

	events := c.event_list.FindAvailableEvents()

	for _, event := range events {
		c.LoadEventDetailPage(event)
	}
}

func (c *UnitedChecker) LoadEventListPage() {
	c.event_list = &UnitedEventListPage{
		&UnitedPage{
			c.browser.MustConnect().MustPage("https://tickets.manutd.com/en-GB/categories/home-tickets"),
		},
	}
}
