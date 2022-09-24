package main

import (
	"fmt"
	haas "iamrobertyoung/manutd-ticket-checker/v2/internal/home-assistant"

	"github.com/go-rod/rod"
)

const UNITED_PREMIER_IMAGE_ID = "1000284.png"
const UNITED_BUY_BUTTON_TEXT = "BUY NOW"
const UNITED_EVENT_PAGE = "https://tickets.manutd.com/en-GB/categories/home-tickets"

type UnitedChecker struct {
	browser             *rod.Browser
	event_list          *UnitedEventListPage
	premier_league_only bool
	haas_api            *haas.HomeAssistantAPI
	available_events    []*UnitedEventItem
}

func (c *UnitedChecker) Check() {
	c.browser = rod.New()
	c.LoadEventListPage()
	c.event_list.DeleteCookieOverlay()
	c.available_events = c.event_list.FindAvailableEvents(c.premier_league_only)

	for _, event := range c.available_events {
		fmt.Printf("Checking %s...", event.Name())

		event.LoadEventDetailPage(event)

		pages, err := c.browser.Pages()

		if err != nil {
			panic(err)
		}

		event_detail_page := UnitedEventDetailPage{
			UnitedPage: &UnitedPage{
				pages.MustFindByURL("/events/"),
			},
		}
		event_detail_page.WaitLoad()
		event_detail_page.DeleteCookieOverlay()

		min_price, max_price := event_detail_page.FindPrices()
		event.min_price = min_price
		event.max_price = max_price

		fmt.Printf(" prices found: £%d -> £%d \n", min_price, max_price)

		event_detail_page.Close()
	}

	c.UpdateHaasState()
	// c.SendNotification()
}

func (c *UnitedChecker) UpdateHaasState() {
	if c.haas_api == nil {
		return
	}

	for _, event := range c.available_events {
		event.UpdateState()
	}
}

func (c *UnitedChecker) LoadEventListPage() {
	c.event_list = &UnitedEventListPage{
		UnitedPage: &UnitedPage{
			c.browser.MustConnect().MustPage(UNITED_EVENT_PAGE),
		},
		haas_api: c.haas_api,
	}
}

func (c *UnitedChecker) SendNotification(device string) {
	request := haas.HomeAssistantNotifyRequest{
		Title:   "Manchester United",
		Message: "Tickets available! 🔴⚽",
	}
	c.haas_api.Notify(device, request)
}
