package main

import (
	"errors"
	"strings"

	"github.com/go-rod/rod"
)

type UnitedEventItem struct {
	element *rod.Element
}

func (e UnitedEventItem) Name() *string {
	default_name := "-"
	name, err := e.element.Attribute("data-name")

	if err != nil {
		return &default_name
	}

	return name
}

func (e UnitedEventItem) FindBuyButton() (*rod.Element, error) {
	element, err := e.element.Element("div.addToBasket:not([style*='display: none']) > a")

	if err != nil {
		return nil, err
	}

	if strings.ToUpper(element.MustText()) == "BUY NOW" {
		return element, err
	}

	return nil, errors.New("buy button not found")
}

func (e UnitedEventItem) BuyButton() *rod.Element {
	return e.element.MustElement("div.addToBasket > a")
}

func (c *UnitedChecker) LoadEventDetailPage(event *UnitedEventItem) {
	buy_button := event.BuyButton()
	buy_button.MustEval(`() => this.target="_blank"`)
	buy_button.MustClick()
}
