package main

import haas "iamrobertyoung/manutd-ticket-checker/v2/internal/home-assistant"

type UnitedEventListPage struct {
	*UnitedPage

	haas_api *haas.HomeAssistantAPI
}

func (c *UnitedEventListPage) FindEvents(premier_league_only bool) []*UnitedEventItem {
	events := c.MustElements("#eventsList .dataItem")
	var event_list []*UnitedEventItem

	for _, element := range events {
		event := UnitedEventItem{
			Element:  element,
			haas_api: c.haas_api,
		}

		is_premier_league := premier_league_only && event.IsPremierLeagueEvent()

		if !premier_league_only || premier_league_only && is_premier_league {
			event_list = append(event_list, &event)
		}
	}

	return event_list
}
