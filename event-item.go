package main

import (
	"errors"
	"strings"

	"github.com/go-rod/rod"
)

type UnitedEventItem struct {
	*rod.Element
}

func (e UnitedEventItem) Name() *string {
	default_name := "-"
	name, err := e.Attribute("data-name")

	if err != nil {
		return &default_name
	}

	return name
}

func (e UnitedEventItem) FindBuyButton() (*rod.Element, error) {
	element, err := e.Element.Element("div.addToBasket:not([style*='display: none']) > a")

	if err != nil {
		return nil, err
	}

	if strings.ToUpper(element.MustText()) == UNITED_BUY_BUTTON_TEXT {
		return element, err
	}

	return nil, errors.New("buy button not found")
}

func (e UnitedEventItem) IsPremierLeagueEvent() bool {
	button, err := e.Element.Element("img.item_image.otherImageMediumImageUrl")

	if err != nil {
		return false
	}

	if strings.Contains(*button.MustAttribute("src"), UNITED_PREMIER_IMAGE_ID) {
		return true
	}

	return false
}

func (e UnitedEventItem) BuyButton() *rod.Element {
	return e.MustElement("div.addToBasket > a")
}

func (e UnitedEventItem) LoadEventDetailPage(event *UnitedEventItem) {
	buy_button := event.BuyButton()
	buy_button.MustEval(`() => this.target="_blank"`)
	buy_button.MustClick()
}
