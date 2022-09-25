package mutc

import (
	"fmt"

	models "github.com/robertyoung/manutd-ticket-checker/v2/cmd/manutd-ticket-checker/models"
	haas "github.com/robertyoung/manutd-ticket-checker/v2/pkg/home-assistant"

	"log"

	"github.com/go-rod/rod"
)

const UNITED_PREMIER_IMAGE_ID = "1000284.png"
const UNITED_BUY_BUTTON_TEXT = "BUY NOW"
const UNITED_EVENT_PAGE = "https://tickets.manutd.com/en-GB/categories/home-tickets"
const UNITED_MAX_PRICE = 100

type UnitedChecker struct {
	browser             *rod.Browser
	store               *Store
	event_list          *UnitedEventListPage
	premier_league_only bool
	haas_api            *haas.HomeAssistantAPI
	events              []*UnitedEventItem
	available_events    []*UnitedEventItem
	haas_notify_device  string
}

func (c *UnitedChecker) Check() {
	c.browser = rod.New()
	c.LoadEventListPage()
	c.event_list.DeleteCookieOverlay()
	c.events = c.event_list.FindEvents(c.premier_league_only)

	for _, event := range c.events {
		_, err := event.FindBuyButton()

		if err != nil {
			continue
		}

		name := event.Name()
		log.Printf("checking %s...", name)

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
		event.MinPrice = min_price
		event.MaxPrice = max_price

		log.Printf("found %s prices: £%d -> £%d \n", name, min_price, max_price)

		event_detail_page.Close()
	}

	c.UpdateHaasState()
	c.SendNotification()
	c.UpdateStore()
}

func (c *UnitedChecker) UpdateHaasState() {
	if c.haas_api == nil {
		return
	}

	for _, event := range c.events {
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

func (c *UnitedChecker) CountEventsAvailable() (int, []*UnitedEventItem) {
	var count int = 0
	var available_events []*UnitedEventItem

	for _, event := range c.events {
		if event.State() != "available" {
			continue
		}

		log.Printf("%s available!\n", event.Name())

		count += 1
		available_events = append(available_events, event)
	}

	return count, available_events
}

func (c *UnitedChecker) SendNotification() {
	if c.haas_api == nil {
		return
	}

	count, available_events := c.CountEventsAvailable()

	if count == 0 {
		return
	}

	request := haas.HomeAssistantNotifyRequest{
		Title:   "Manchester United",
		Message: fmt.Sprintf("Tickets available (%d)! 🔴⚽", count),
	}
	c.haas_api.Notify(c.haas_notify_device, request)

	for _, event := range available_events {
		event.NotificationSent()
	}
}

func (c *UnitedChecker) UpdateStore() {
	var event_models []models.EventModel

	for _, event := range c.events {
		event_models = append(event_models, event.ToEventModel())
	}

	c.store.Save(event_models)
}

func (c *UnitedChecker) ReadStore() []models.EventModel {
	return c.store.Read()
}
