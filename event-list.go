package main

import (
	"fmt"
)

type UnitedEventListPage struct {
	*UnitedPage
}

func (c *UnitedEventListPage) FindAvailableEvents() []*UnitedEventItem {
	events := c.MustElements("#eventsList .dataItem")
	var availableEvents []*UnitedEventItem

	for _, element := range events {
		event := UnitedEventItem{element}

		fmt.Println(*event.Name())

		_, err := event.FindBuyButton()

		if err == nil {
			availableEvents = append(availableEvents, &event)
		}
	}

	return availableEvents
}
