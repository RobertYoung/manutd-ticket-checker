package main

import "github.com/go-rod/rod"

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
	return e.element.Element("div.addToBasket:not([style*='display: none']) > a")
}

func (e UnitedEventItem) BuyButton() *rod.Element {
	return e.element.MustElement("div.addToBasket > a")
}
