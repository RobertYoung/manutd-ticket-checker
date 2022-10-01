package mutc

import (
	"fmt"
	"log"
	"time"

	models "github.com/robertyoung/manutd-ticket-checker/v2/cmd/manutd-ticket-checker/models"
	haas "github.com/robertyoung/manutd-ticket-checker/v2/pkg/home-assistant"

	"github.com/go-rod/rod"
	"golang.org/x/exp/slices"
)

const UNITED_PREMIER_IMAGE_ID = "1000284.png"
const UNITED_BUY_BUTTON_TEXT = "BUY NOW"
const UNITED_EVENT_PAGE = "https://tickets.manutd.com/en-GB/categories/home-tickets"

type UnitedChecker struct {
	browser *rod.Browser
	store   *Store

	config *Config

	event_list          *UnitedEventListPage
	haas_api            *haas.HomeAssistantAPI
	events              []*UnitedEventItem
	available_events    []*UnitedEventItem
	notification_events []*UnitedEventItem
}

func (c *UnitedChecker) Check() {
	c.browser = rod.New()
	c.LoadEventListPage()
	c.event_list.DeleteCookieOverlay()
	c.events = c.event_list.FindEvents(c.config.PremierLeagueOnly)

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
	c.EventsAvailable()
	c.ReadStore()
	c.NotificationEvents()
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
		config:   c.config,
	}
}

func (c *UnitedChecker) EventsAvailable() []*UnitedEventItem {
	c.available_events = nil

	for _, event := range c.events {
		if event.State() != "available" {
			continue
		}

		if event.MaxPrice == 0 {
			continue
		}

		log.Printf("%s available!\n", event.Name())

		c.available_events = append(c.available_events, event)
	}

	return c.available_events
}

func (c *UnitedChecker) NotificationEvents() []*UnitedEventItem {
	if c.available_events == nil {
		panic("available events unavailable")
	}

	for _, available_event := range c.available_events {
		if available_event.Model == nil {
			c.notification_events = append(c.notification_events, available_event)
			continue
		}

		refresh_time := time.Now().Add(-time.Minute * time.Duration(c.config.HaasNotificationThrottle))

		if available_event.Model.NotificationSentAt.Before(refresh_time) {
			c.notification_events = append(c.notification_events, available_event)
		}
	}

	return c.notification_events
}

func (c *UnitedChecker) SendNotification() {
	if c.haas_api == nil {
		return
	}

	count := len(c.notification_events)

	if count == 0 {
		return
	}

	for _, event := range c.notification_events {
		request := haas.HomeAssistantNotifyRequest{
			Title:   "Manchester United",
			Message: fmt.Sprintf("Tickets available for %s (£%d -> £%d)! 🔴⚽", event.Name(), event.MinPrice, event.MaxPrice),
		}
		c.haas_api.Notify(c.config.HaasNotifyDevice, request)

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
	var records = c.store.Read()

	for _, available_event := range c.available_events {
		index := slices.IndexFunc(records, func(model models.EventModel) bool {
			return model.Uuid == available_event.Uuid()
		})

		if index >= 0 {
			model := records[index]
			available_event.Model = &model
		}
	}

	return records
}
