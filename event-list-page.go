package main

import haas "iamrobertyoung/manutd-ticket-checker/v2/internal/home-assistant"

type UnitedEventListPage struct {
	*UnitedPage

	haas_api *haas.HomeAssistantAPI
}

func (c *UnitedEventListPage) FindAvailableEvents(premier_league_only bool) []*UnitedEventItem {
	events := c.MustElements("#eventsList .dataItem")
	var availableEvents []*UnitedEventItem

	for _, element := range events {
		event := UnitedEventItem{
			Element:  element,
			haas_api: c.haas_api,
		}

		_, err := event.FindBuyButton()
		is_premier_league := premier_league_only && event.IsPremierLeagueEvent()

		if err == nil && (!premier_league_only || premier_league_only && is_premier_league) {
			availableEvents = append(availableEvents, &event)
		}
	}

	return availableEvents
}
