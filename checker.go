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
}

func (c *UnitedChecker) Check() {
	c.browser = rod.New()
	c.LoadEventListPage()
	c.event_list.DeleteCookieOverlay()

	events := c.event_list.FindAvailableEvents(c.premier_league_only)

	for _, event := range events {
		fmt.Printf("Checking %s...", *event.Name())

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
		event_detail_page.FindPrices()
		event_detail_page.Close()

		fmt.Printf(" prices found: £%d -> £%d \n", event_detail_page.min_price, event_detail_page.max_price)
	}
}

func (c *UnitedChecker) LoadEventListPage() {
	c.event_list = &UnitedEventListPage{
		&UnitedPage{
			c.browser.MustConnect().MustPage(UNITED_EVENT_PAGE),
		},
	}
}

func (c *UnitedChecker) SendNotification(device string) {
	if c.haas_api == nil {
		return
	}

	request := haas.HomeAssistantNotifyRequest{
		Title:   "Manchester United",
		Message: "Tickets available! 🔴⚽",
	}
	c.haas_api.Notify(device, request)
}

func (c *UnitedChecker) UpdateState(entity, state string, min_price, max_price uint16) {
	if c.haas_api == nil {
		return
	}

	request := haas.HomeAssistantStateUpdateRequest{
		State: state,
		Attribute: HomeAssistantMatchStateAttributes{
			MinPrice: min_price,
			MaxPrice: max_price,
		},
	}
	c.haas_api.StateUpdate(entity, request)
}
