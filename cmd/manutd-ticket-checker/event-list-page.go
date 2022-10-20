package mutc

import (
	"log"

	haas "github.com/robertyoung/manutd-ticket-checker/v2/pkg/home-assistant"
)

type UnitedEventListPage struct {
	*UnitedPage

	config   *Config
	haas_api *haas.HomeAssistantAPI
}

func (c *UnitedEventListPage) FindEvents(premier_league_only bool) []*UnitedEventItem {
	events, err := c.Elements("#eventsList .dataItem")

	if err != nil {
		log.Print("failed to find event items: ", err)
		return []*UnitedEventItem{}
	}

	var event_list []*UnitedEventItem

	for _, element := range events {
		event := UnitedEventItem{
			Element:  element,
			haas_api: c.haas_api,
			config:   c.config,
		}

		is_premier_league := premier_league_only && event.IsPremierLeagueEvent()

		if !premier_league_only || premier_league_only && is_premier_league {
			event_list = append(event_list, &event)
		}
	}

	return event_list
}
