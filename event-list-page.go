package main

type UnitedEventListPage struct {
	*UnitedPage
}

func (c *UnitedEventListPage) FindAvailableEvents(premier_league_only bool) []*UnitedEventItem {
	events := c.MustElements("#eventsList .dataItem")
	var availableEvents []*UnitedEventItem

	for _, element := range events {
		event := UnitedEventItem{element}

		_, err := event.FindBuyButton()
		is_premier_league := premier_league_only && event.IsPremierLeagueEvent()

		if err == nil && (!premier_league_only || premier_league_only && is_premier_league) {
			availableEvents = append(availableEvents, &event)
		}
	}

	return availableEvents
}
