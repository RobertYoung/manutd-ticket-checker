package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

func main() {
	page := rod.New().MustConnect().MustPage("https://tickets.manutd.com/en-GB/categories/home-tickets")

	tint_overlay, err := page.Timeout(5 * time.Second).Element("#tint-overlay")

	if err == nil {
		tint_overlay.Remove()
	}

	cookie_messages, err := page.Timeout(5 * time.Second).Element("#cookie-messages")

	if err == nil {
		cookie_messages.Remove()
	}

	events := page.MustElements("#eventsList .dataItem")

	for _, event := range events {
		name, err := event.Attribute("data-name")

		if err == nil {
			fmt.Println(*name)
		}
	}
}
